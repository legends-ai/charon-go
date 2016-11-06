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
	MatchId   uint64 `json:"matchId"`
	Timestamp uint64 `json:"timestamp"`
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

	resp, err := r.fetchWithParams("matchlist",
		fmt.Sprintf("%s/v2.2/matchlist/by-summoner/%d", r.apiLol, summoner), vals, true)
	if err != nil {
		return nil, err
	}

	var m MatchListResponse
	if err = json.Unmarshal(resp, &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal match list response: %v", err)
	}
	return &m, nil
}
