package riot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/asunaio/charon/config"
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/metrics"
)

const (
	envRiot      = "RIOT_BASE"
	apiKeyParam  = "api_key"
	riotBaseTpl  = "https://%s.api.pvp.net"
	regionHeader = "riot-region"
)

// Client stores API clients.
type Client struct {
	Config  *config.AppConfig `inject:"t"`
	Metrics *metrics.Metrics  `inject:"t"`

	clients   map[apb.Region]*API
	clientsMu sync.RWMutex
}

// New creates a new Client.
func New() *Client {
	return &Client{
		clients: map[apb.Region]*API{},
	}
}

// Region gets an API client for the given region.
func (rc *Client) Region(region apb.Region) *API {
	rc.clientsMu.RLock()
	inst, ok := rc.clients[region]
	rg := strings.ToLower(apb.Region_name[int32(region)])
	rc.clientsMu.RUnlock()
	if !ok {
		base := fmt.Sprintf(riotBaseTpl, rg)
		inst = &API{
			Region:  rg,
			rl:      NewRateLimiter(rc.Config.MaxRate),
			apiBase: base,
			apiLol:  fmt.Sprintf("%s/api/lol/%s", base, rg),
			rc:      rc,
		}
		rc.clientsMu.Lock()
		rc.clients[region] = inst
		rc.clientsMu.Unlock()
	}
	return inst
}

// API is the Riot API interface
type API struct {
	Region  string
	rl      *RateLimiter
	apiBase string
	apiLol  string
	rc      *Client
}

// fetchWithParams fetches a path with the given parameters.
func (r *API) fetchWithParams(endpoint string, path string, params url.Values) ([]byte, error) {
	key := r.rc.Config.APIKey
	params.Set(apiKeyParam, key)
	url := fmt.Sprintf("%s?%s", path, params.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for {
		r.rl.Wait(context.TODO())
		resp, err := client.Do(req)
		r.rc.Metrics.Record(endpoint)
		if err != nil {
			return nil, err
		}
		// free memory
		defer resp.Body.Close()

		if resp.StatusCode != 429 {
			// we good
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("could not read body: %v", err)
			}
			return body, nil
		}

		// let's retry

		// check service level retry
		retryAfter := resp.Header.Get("Retry-After")
		if retryAfter != "" {
			seconds, err := strconv.Atoi(retryAfter)
			if err != nil {
				return nil, fmt.Errorf("could not parse Retry-After header: %v", err)
			}
			// extra 500ms for good measure
			r.rl.RetryAfter(time.Duration(seconds)*time.Second + 500*time.Millisecond)
		} else {
			// no retry header specified, try 1 second
			// https://developer.riotgames.com/docs/rate-limiting
			// TODO(igm): evaluate whether we give a shit about Riot rate limits.
			// if we remove the below code, we get 4x data.
			r.rl.RetryAfter(1 * time.Second)
		}
	}

	// unreachable
	return nil, nil
}

// fetch fetches a path via GET request.
func (r *API) fetch(endpoint string, path string) ([]byte, error) {
	return r.fetchWithParams(endpoint, path, url.Values{})
}