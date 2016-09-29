package riot

import (
	"encoding/json"
	"fmt"
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
	resp, err := r.fetchWithParams("match",
		fmt.Sprintf("%s/v2.2/match/%d", r.apiLol, matchID), url.Values{"includeTimeline": []string{"true"}})
	if err != nil {
		return nil, err
	}
	var m MatchResponse

	if err = json.Unmarshal(resp, &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal match response: %v", err)
	}

	m.RawJSON = string(resp)
	return &m, nil
}
