package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/asunaio/charon/riot/models"
)

// Match gets match details
func (r *API) Match(matchID uint64) (*models.RiotMatch, error) {
	res, err := r.fetchWithParams(
		"match",
		fmt.Sprintf("%s/v2.2/match/%d", r.apiLol, matchID),
		url.Values{"includeTimeline": []string{"true"}},
		true,
	)

	if err != nil {
		return nil, err
	}

	var m models.RiotMatch
	if err = json.Unmarshal(res, &m); err != nil {
		return nil, fmt.Errorf("could not to unmarshal match response: %v", err)
	}

	return &m, nil
}
