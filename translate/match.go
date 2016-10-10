package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func Match(match *models.RiotMatch) *apb.Charon_Match {
	return &apb.Charon_Match{
		MapId:        match.MapId,
		CreationTime: match.MatchCreation,
		Duration:     match.MatchDuration,
		MatchId:      match.MatchId,
		Mode:         parseMatchMode(match.MatchMode),
		Type:         parseMatchType(match.MatchType),
		Version:      match.MatchVersion,
		PlatformId:   match.PlatformId,
		Region:       parseRegion(match.Region),
		QueueType:    parseQueueType(match.QueueType),
		Season:       parseSeason(match.Season),
	}
}
