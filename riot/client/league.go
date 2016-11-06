package riot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	QueueSolo5x5                   = "RANKED_SOLO_5x5"
	QueuePremade5x5                = "RANKED_PREMADE_5x5"
	QueueTeam3x3                   = "RANKED_TEAM_3x3"
	QueueTeam5x5                   = "RANKED_TEAM_5x5"
	QueueTeamBuilderDraftRanked5x5 = "TEAM_BUILDER_DRAFT_RANKED_5x5"
)

// LeagueResponse is the league response
type LeagueResponse map[string][]*LeagueDto

// LeagueDto contains league information.
type LeagueDto struct {
	Name    string            `json:"name"`
	Queue   string            `json:"queue"`
	Tier    string            `json:"tier"`
	Entries []*LeagueEntryDto `json:"entries"`
}

// LeagueEntryDto contains league participant information representing a summoner or team.
type LeagueEntryDto struct {
	PlayerOrTeamID   string `json:"playerOrTeamId"`
	PlayerOrTeamName string `json:"playerOrTeamName"`
	Division         string `json:"division"`
}

// League gets a league
func (r *API) League(ids []uint64) (LeagueResponse, error) {
	// adapt so we can use join
	var summonerIDs []string
	for _, id := range ids {
		summonerIDs = append(summonerIDs, strconv.FormatUint(id, 10))
	}
	idsStr := strings.Join(summonerIDs, ",")

	resp, err := r.fetch("league",
		fmt.Sprintf("%s/v2.5/league/by-summoner/%s", r.apiLol, idsStr), true)
	if err != nil {
		return nil, err
	}

	ret := LeagueResponse{}
	if strings.Contains(string(resp), "Not Found") {
		// none of the summoners in the list were found
		return ret, nil
	}

	if err = json.Unmarshal(resp, &ret); err != nil {
		return nil, fmt.Errorf("could not unmarshal league response: %v", err)
	}
	return ret, nil
}
