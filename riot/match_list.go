package riot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// MatchListResponse is the match list response.
type MatchListResponse struct {
	Matches []MatchListMatch `json:"matches"`
}

// MatchListMatch is a match in the match list.
type MatchListMatch struct {
	MatchId uint64 `json:"matchId"`
}

// MatchList gets match list
func (r *API) MatchList(summoner uint64, queues []string, seasons []string) (*MatchListResponse, error) {
	vals := url.Values{}
	if queues != nil {
		vals["rankedQueues"] = []string{strings.Join(queues, ",")}
	}

	if seasons != nil {
		vals["seasons"] = []string{strings.Join(seasons, ",")}
	}

	resp, err := r.fetchWithParams(
		fmt.Sprintf("%s/v2.2/matchlist/by-summoner/%d", r.apiLol, summoner), vals)
	if err != nil {
		return nil, err
	}

	var m MatchListResponse
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, fmt.Errorf("Could not unmarshal match response: %v", err)
	}

	return &m, nil
}
