package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func Match(match *models.RiotMatch) *apb.Charon_Match {
	return &apb.Charon_Match{
		MapId:           match.MapId,
		CreationTime:    match.MatchCreation,
		Duration:        match.MatchDuration,
		MatchId:         match.MatchId,
		Mode:            parseMatchMode(match.MatchMode),
		Type:            parseMatchType(match.MatchType),
		Version:         match.MatchVersion,
		PlatformId:      match.PlatformId,
		Region:          parseRegion(match.Region),
		QueueType:       parseQueueType(match.QueueType),
		Season:          parseSeason(match.Season),
		ParticipantInfo: generateParticipantInfo(match),
	}
}

func generateParticipantInfo(match *models.RiotMatch) []*apb.Charon_Match_ParticipantInfo {
	var info []*apb.Charon_Match_ParticipantInfo
	identityMap := mapParticipantIdentities(match.ParticipantIdentities)
	for _, p := range match.Participants {
		identity := identityMap[p.ParticipantId]
		info = append(info, &apb.Charon_Match_ParticipantInfo{
			ParticipantId: p.ParticipantId,
			TeamId:        p.TeamId,
			Identity: &apb.Charon_Match_ParticipantInfo_Identity{
				MatchHistoryUri: identity.MatchHistoryUri,
				ProfileIconId:   identity.ProfileIcon,
				SummonerId:      identity.SummonerId,
				SummonerName:    identity.SummonerName,
			},
			ChampionId: p.ChampionId,
			SummonerSpells: &apb.Charon_Match_ParticipantInfo_SummonerSpells{
				Spell1: p.Spell1Id,
				Spell2: p.Spell2Id,
			},
		})
	}
	return info
}

func mapParticipantIdentities(identities []models.ParticipantIdentity) map[uint32]models.Player {
	identityMap := map[uint32]models.Player{}
	for _, i := range identities {
		identityMap[i.ParticipantId] = i.Player
	}
	return identityMap
}
