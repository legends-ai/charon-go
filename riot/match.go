package riot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// MatchResponse is the match response
type MatchResponse struct {
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	MatchVersion          string                `json:"matchVersion"`
	QueueType             string                `json:"queueType"`
	RawJSON               string                `json:"-"`
}

// ParticipantIdentity is the identity of a participant
type ParticipantIdentity struct {
	Player MatchResponsePlayer `json:"player"`
}

// MatchResponsePlayer is a player of a match response
type MatchResponsePlayer struct {
	SummonerID uint64 `json:"summonerId"`
}

// Match gets match details
func (r *API) Match(matchID uint64) (*MatchResponse, error) {
	resp, err := r.fetchWithParams(
		fmt.Sprintf("%s/v2.2/match/%d", r.apiLol, matchID), url.Values{"includeTimeline": []string{"true"}})
	if err != nil {
		return nil, err
	}
	var m MatchResponse
	defer resp.Body.Close()

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read match response: %v", err)
	}

	if err = json.Unmarshal(s, &m); err != nil {
		return nil, fmt.Errorf("Could not unmarshal match response: %v", err)
	}

	m.RawJSON = string(s)
	return &m, nil
}
