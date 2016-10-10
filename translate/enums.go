package translate

import apb "github.com/asunaio/charon/gen-go/asuna"

func parseSeason(season string) apb.Season {
	switch season {
	case "PRESEASON3":
		return apb.Season_PRESEASON3
	case "SEASON3":
		return apb.Season_SEASON3
	case "PRESEASON2014":
		return apb.Season_PRESEASON2014
	case "SEASON2014":
		return apb.Season_SEASON2014
	case "PRESEASON2015":
		return apb.Season_PRESEASON2015
	case "SEASON2015":
		return apb.Season_SEASON2015
	case "PRESEASON2016":
		return apb.Season_PRESEASON2016
	case "SEASON2016":
		return apb.Season_SEASON2016
	default:
		return apb.Season_UNDEFINED_SEASON
	}
}

func parseRegion(region string) apb.Region {
	switch region {
	case "BR":
		return apb.Region_BR
	case "EUNE":
		return apb.Region_EUNE
	case "EUW":
		return apb.Region_EUW
	case "JP":
		return apb.Region_JP
	case "KR":
		return apb.Region_KR
	case "LAN":
		return apb.Region_LAN
	case "LAS":
		return apb.Region_LAS
	case "NA":
		return apb.Region_NA
	case "OCE":
		return apb.Region_OCE
	case "TR":
		return apb.Region_TR
	case "RU":
		return apb.Region_RU
	case "PBE":
		return apb.Region_PBE
	case "GLOBAL":
		return apb.Region_GLOBAL
	default:
		return apb.Region_UNDEFINED_REGION
	}
}

func parseMatchMode(mmode string) apb.MatchMode {
	switch mmode {
	case "CLASSIC":
		return apb.MatchMode_CLASSIC
	case "ODIN":
		return apb.MatchMode_ODIN
	case "ARAM":
		return apb.MatchMode_ARAM
	case "TUTORIAL":
		return apb.MatchMode_TUTORIAL
	case "ONEFORALL":
		return apb.MatchMode_ONEFORALL
	case "ASCENSION":
		return apb.MatchMode_ASCENSION
	case "FIRSTBLOOD":
		return apb.MatchMode_FIRSTBLOOD
	case "KINGPORO":
		return apb.MatchMode_KINGPORO
	default:
		return apb.MatchMode_UNDEFINED_MATCH_MODE
	}
}

func parseMatchType(mtype string) apb.MatchType {
	switch mtype {
	case "CUSTOM_GAME":
		return apb.MatchType_CUSTOM_GAME
	case "MATCHED_GAME":
		return apb.MatchType_MATCHED_GAME
	case "TUTORIAL_GAME":
		return apb.MatchType_TUTORIAL_GAME
	default:
		return apb.MatchType_UNDEFINED_MATCH_TYPE
	}
}

func parseQueueType(qtype string) apb.QueueType {
	switch qtype {
	case "CUSTOM":
		return apb.QueueType_CUSTOM
	case "NORMAL_5x5_BLIND":
		return apb.QueueType_NORMAL_5x5_BLIND
	case "RANKED_SOLO_5x5":
		return apb.QueueType_RANKED_SOLO_5x5
	case "RANKED_PREMADE_5x5":
		return apb.QueueType_RANKED_PREMADE_5x5
	case "BOT_5x5":
		return apb.QueueType_BOT_5x5
	case "NORMAL_3x3":
		return apb.QueueType_NORMAL_3x3
	case "RANKED_PREMADE_3x3":
		return apb.QueueType_RANKED_PREMADE_3x3
	case "NORMAL_5x5_DRAFT":
		return apb.QueueType_NORMAL_5x5_DRAFT
	case "ODIN_5x5_BLIND":
		return apb.QueueType_ODIN_5x5_BLIND
	case "ODIN_5x5_DRAFT":
		return apb.QueueType_ODIN_5x5_DRAFT
	case "BOT_ODIN_5x5":
		return apb.QueueType_BOT_ODIN_5x5
	case "BOT_5x5_INTRO":
		return apb.QueueType_BOT_5x5_INTRO
	case "BOT_5x5_BEGINNER":
		return apb.QueueType_BOT_5x5_BEGINNER
	case "BOT_5x5_INTERMEDIATE":
		return apb.QueueType_BOT_5x5_INTERMEDIATE
	case "RANKED_TEAM_3x3":
		return apb.QueueType_RANKED_TEAM_3x3
	case "RANKED_TEAM_5x5":
		return apb.QueueType_RANKED_TEAM_5x5
	case "BOT_TT_3x3":
		return apb.QueueType_BOT_TT_3x3
	case "GROUP_FINDER_5x5":
		return apb.QueueType_GROUP_FINDER_5x5
	case "ARAM_5x5":
		return apb.QueueType_ARAM_5x5
	case "ONEFORALL_5x5":
		return apb.QueueType_ONEFORALL_5x5
	case "FIRSTBLOOD_1x1":
		return apb.QueueType_FIRSTBLOOD_1x1
	case "FIRSTBLOOD_2x2":
		return apb.QueueType_FIRSTBLOOD_2x2
	case "SR_6x6":
		return apb.QueueType_SR_6x6
	case "URF_5x5":
		return apb.QueueType_URF_5x5
	case "ONEFORALL_MIRRORMODE_5x5":
		return apb.QueueType_ONEFORALL_MIRRORMODE_5x5
	case "BOT_URF_5x5":
		return apb.QueueType_BOT_URF_5x5
	case "NIGHTMARE_BOT_5x5_RANK1":
		return apb.QueueType_NIGHTMARE_BOT_5x5_RANK1
	case "NIGHTMARE_BOT_5x5_RANK2":
		return apb.QueueType_NIGHTMARE_BOT_5x5_RANK2
	case "NIGHTMARE_BOT_5x5_RANK5":
		return apb.QueueType_NIGHTMARE_BOT_5x5_RANK5
	case "ASCENSION_5x5":
		return apb.QueueType_ASCENSION_5x5
	case "HEXAKILL":
		return apb.QueueType_HEXAKILL
	case "BILGEWATER_ARAM_5x5":
		return apb.QueueType_BILGEWATER_ARAM_5x5
	case "KING_PORO_5x5":
		return apb.QueueType_KING_PORO_5x5
	case "COUNTER_PICK":
		return apb.QueueType_COUNTER_PICK
	case "BILGEWATER_5x5":
		return apb.QueueType_BILGEWATER_5x5
	case "TEAM_BUILDER_DRAFT_UNRANKED_5x5":
		return apb.QueueType_TEAM_BUILDER_DRAFT_UNRANKED_5x5
	case "TEAM_BUILDER_DRAFT_RANKED_5x5":
		return apb.QueueType_TEAM_BUILDER_DRAFT_RANKED_5x5
	default:
		return apb.QueueType_UNDEFINED_QUEUE_TYPE
	}
}
