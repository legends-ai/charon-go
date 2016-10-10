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
		Mode:            apb.MatchMode(apb.MatchMode_value[match.MatchMode]),
		Type:            apb.MatchType(apb.MatchType_value[match.MatchType]),
		Version:         match.MatchVersion,
		PlatformId:      match.PlatformId,
		Region:          apb.Region(apb.Region_value[match.Region]),
		QueueType:       apb.QueueType(apb.QueueType_value[match.QueueType]),
		Season:          apb.Season(apb.Season_value[match.Season]),
		ParticipantInfo: generateParticipantInfo(match),
		TeamInfo:        generateTeamInfo(match),
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
			ChampionId:        p.ChampionId,
			HighestSeasonTier: apb.Tier(apb.Tier_value[p.HighestAchievedSeasonTier]),
			Lane:              parseLane(p.Timeline.Lane),
			TeamRole:          apb.TeamRole(apb.TeamRole_value[p.Timeline.Role]),
			SummonerSpells: &apb.Charon_Match_ParticipantInfo_SummonerSpells{
				Spell1: p.Spell1Id,
				Spell2: p.Spell2Id,
			},
			Masteries: generateMasteries(p.Masteries),
			Runes:     generateRunes(p.Runes),
			Items: []uint64{
				p.Stats.Item0, p.Stats.Item1, p.Stats.Item2,
				p.Stats.Item3, p.Stats.Item4, p.Stats.Item5, p.Stats.Item6,
			},
			Stats: &apb.Charon_Match_ParticipantInfo_Stats{
				Winner:                          p.Stats.Winner,
				ChampionLevel:                   p.Stats.ChampLevel,
				Kills:                           p.Stats.Kills,
				Deaths:                          p.Stats.Deaths,
				Assists:                         p.Stats.Assists,
				DoubleKills:                     p.Stats.DoubleKills,
				TripleKills:                     p.Stats.TripleKills,
				QuadraKills:                     p.Stats.QuadraKills,
				PentaKills:                      p.Stats.PentaKills,
				UnrealKills:                     p.Stats.UnrealKills,
				TowerKills:                      p.Stats.TowerKills,
				InhibitorKills:                  p.Stats.InhibitorKills,
				FirstBloodAssist:                p.Stats.FirstBloodAssist,
				FirstBloodKill:                  p.Stats.FirstBloodKill,
				FirstInhibitorAssist:            p.Stats.FirstInhibitorAssist,
				FirstInhibitorKill:              p.Stats.FirstInhibitorKill,
				FirstTowerAssist:                p.Stats.FirstTowerAssist,
				FirstTowerKill:                  p.Stats.FirstTowerKill,
				GoldEarned:                      p.Stats.GoldEarned,
				GoldSpent:                       p.Stats.GoldSpent,
				SightWardsBought:                p.Stats.SightWardsBoughtInGame,
				VisionWardsBought:               p.Stats.VisionWardsBoughtInGame,
				WardsKilled:                     p.Stats.WardsKilled,
				WardsPlaced:                     p.Stats.WardsPlaced,
				MinionsKilled:                   p.Stats.MinionsKilled,
				NeutralMinionsKilled:            p.Stats.NeutralMinionsKilled,
				NeutralMinionsKilledEnemyJungle: p.Stats.NeutralMinionsKilledEnemyJungle,
				NeutralMinionsKilledTeamJungle:  p.Stats.NeutralMinionsKilledTeamJungle,
				PhysicalDamageDealt:             p.Stats.PhysicalDamageDealt,
				PhysicalDamageDealtToChampions:  p.Stats.PhysicalDamageDealtToChampions,
				PhysicalDamageTaken:             p.Stats.PhysicalDamageTaken,
				MagicDamageDealt:                p.Stats.MagicDamageDealt,
				MagicDamageDealtToChampions:     p.Stats.MagicDamageDealtToChampions,
				MagicDamageTaken:                p.Stats.MagicDamageTaken,
				TrueDamageDealt:                 p.Stats.TrueDamageDealt,
				TrueDamageDealtToChampions:      p.Stats.TrueDamageDealtToChampions,
				TrueDamageTaken:                 p.Stats.TrueDamageTaken,
				TotalDamageDealt:                p.Stats.TotalDamageDealt,
				TotalDamageDealtToChampions:     p.Stats.TotalDamageDealtToChampions,
				TotalDamageTaken:                p.Stats.TotalDamageTaken,
				KillingSprees:                   p.Stats.KillingSprees,
				LargestKillingSpree:             p.Stats.LargestKillingSpree,
				LargestCriticalStrike:           p.Stats.LargestCriticalStrike,
				LargestMultiKill:                p.Stats.LargestMultiKill,
				CombatPlayerScore:               p.Stats.CombatPlayerScore,
				ObjectivePlayerScore:            p.Stats.ObjectivePlayerScore,
				TotalPlayerScore:                p.Stats.TotalPlayerScore,
				TotalPlayerScoreRank:            p.Stats.TotalScoreRank,
				NodeCaptures:                    p.Stats.NodeCapture,
				NodeCaptureAssists:              p.Stats.NodeCaptureAssist,
				NodeNeutralizations:             p.Stats.NodeNeutralize,
				NodeNeutralizationAssists:       p.Stats.NodeNeutralizeAssist,
				TeamObjectivesCompleted:         p.Stats.TeamObjective,
			},
			Timeline: &apb.Charon_Match_ParticipantInfo_Timeline{
				CreepsPerMin:              generateDelta(p.Timeline.CreepsPerMinDeltas),
				DamageTakenPerMin:         generateDelta(p.Timeline.DamageTakenPerMinDeltas),
				GoldPerMin:                generateDelta(p.Timeline.GoldPerMinDeltas),
				WardsPerMin:               generateDelta(p.Timeline.WardsPerMinDeltas),
				XpPerMin:                  generateDelta(p.Timeline.XpPerMinDeltas),
				AncientGolemAssistsPerMin: generateDelta(p.Timeline.AncientGolemAssistsPerMinCounts),
				AncientGolemKillsPerMin:   generateDelta(p.Timeline.AncientGolemKillsPerMinCounts),
				AssistedLaneDeathsPerMin:  generateDelta(p.Timeline.AssistedLaneDeathsPerMinDeltas),
				AssistedLaneKillsPerMin:   generateDelta(p.Timeline.AssistedLaneKillsPerMinDeltas),
				BaronAssistsPerMin:        generateDelta(p.Timeline.BaronAssistsPerMinCounts),
				BaronKillsPerMin:          generateDelta(p.Timeline.BaronKillsPerMinCounts),
				DragonAssistsPerMin:       generateDelta(p.Timeline.DragonAssistsPerMinCounts),
				DragonKillsPerMin:         generateDelta(p.Timeline.DragonKillsPerMinCounts),
				ElderLizardAssistsPerMin:  generateDelta(p.Timeline.ElderLizardAssistsPerMinCounts),
				ElderLizardKillsPerMin:    generateDelta(p.Timeline.ElderLizardKillsPerMinCounts),
				InhibitorAssistsPerMin:    generateDelta(p.Timeline.InhibitorAssistsPerMinCounts),
				InhibitorKillsPerMin:      generateDelta(p.Timeline.InhibitorKillsPerMinCounts),
				TowerAssistsPerMin:        generateDelta(p.Timeline.TowerAssistsPerMinCounts),
				TowerKillsPerMin:          generateDelta(p.Timeline.TowerKillsPerMinCounts),
				TowerKillsPerMinDeltas:    generateDelta(p.Timeline.TowerKillsPerMinDeltas),
				VilemawAssistsPerMin:      generateDelta(p.Timeline.VilemawAssistsPerMinCounts),
				VilemawKillsPerMin:        generateDelta(p.Timeline.VilemawKillsPerMinCounts),
				CsDiffPerMin:              generateDelta(p.Timeline.CsDiffPerMinDeltas),
				DamageTakenDiffPerMin:     generateDelta(p.Timeline.DamageTakenPerMinDeltas),
				XpDiffPerMin:              generateDelta(p.Timeline.XpDiffPerMinDeltas),
			},
		})
	}
	return info
}

func generateTeamInfo(match *models.RiotMatch) []*apb.Charon_Match_TeamInfo {
	var info []*apb.Charon_Match_TeamInfo
	for _, t := range match.Teams {
		info = append(info, &apb.Charon_Match_TeamInfo{
			Winner:               t.Winner,
			Bans:                 generateBans(t.Bans),
			BaronKills:           t.BaronKills,
			DragonKills:          t.DragonKills,
			FirstBaron:           t.FirstBaron,
			FirstDragon:          t.FirstDragon,
			FirstInhibitor:       t.FirstInhibitor,
			FirstRiftHerald:      t.FirstRiftHerald,
			FirstTower:           t.FirstTower,
			InhibitorKills:       t.InhibitorKills,
			RiftHeraldKills:      t.RiftHeraldKills,
			VilemawKills:         t.VilemawKills,
			DominionVictoryScore: t.DominionVictoryScore,
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

func parseLane(lane string) apb.Lane {
	switch lane {
	case "TOP":
		return apb.Lane_LANE_TOP
	case "JUNGLE":
		return apb.Lane_LANE_JUNGLE
	case "MID", "MIDDLE":
		return apb.Lane_LANE_MIDDLE
	case "BOT", "BOTTOM":
		return apb.Lane_LANE_BOTTOM
	default:
		return apb.Lane_UNDEFINED_LANE
	}
}

func generateMasteries(masteries []models.Mastery) []*apb.Charon_Match_ParticipantInfo_Mastery {
	out := []*apb.Charon_Match_ParticipantInfo_Mastery{}
	for _, m := range masteries {
		out = append(out, &apb.Charon_Match_ParticipantInfo_Mastery{
			Id:   m.MasteryId,
			Rank: m.Rank,
		})
	}
	return out
}

func generateRunes(runes []models.Rune) []*apb.Charon_Match_ParticipantInfo_Rune {
	out := []*apb.Charon_Match_ParticipantInfo_Rune{}
	for _, m := range runes {
		out = append(out, &apb.Charon_Match_ParticipantInfo_Rune{
			Id:   m.RuneId,
			Rank: m.Rank,
		})
	}
	return out
}

func generateDelta(delta models.ParticipantTimelineData) *apb.Charon_Match_ParticipantInfo_Timeline_Delta {
	return &apb.Charon_Match_ParticipantInfo_Timeline_Delta{
		ZeroToTen:      delta.ZeroToTen,
		TenToTwenty:    delta.TenToTwenty,
		TwentyToThirty: delta.TwentyToThirty,
		ThirtyToEnd:    delta.ThirtyToEnd,
	}
}

func generateBans(bcs []models.BannedChampion) []*apb.Charon_Match_TeamInfo_BannedChampion {
	out := []*apb.Charon_Match_TeamInfo_BannedChampion{}
	for _, bc := range bcs {
		out = append(out, &apb.Charon_Match_TeamInfo_BannedChampion{
			ChampionId: bc.ChampionId,
			PickTurn:   bc.PickTurn,
		})
	}
	return out
}
