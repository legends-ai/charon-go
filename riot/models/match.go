package models

type RiotMatch struct {
	MapId                 uint32                `json:"mapId"`
	MatchCreation         uint64                `json:"matchCreation"`
	MatchDuration         uint64                `json:"matchDuration"`
	MatchId               uint64                `json:"matchId"`
	MatchMode             string                `json:"matchMode"`
	MatchType             string                `json:"matchType"`
	MatchVersion          string                `json:"matchVersion"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	PlatformId            string                `json:"platformId"`
	QueueType             string                `json:"queueType"`
	Region                string                `json:"region"`
	Season                string                `json:"season"`
	Teams                 []Team                `json:"teams"`
	Timeline              Timeline              `json:"timeline"`
}

type ParticipantIdentity struct {
	ParticipantId uint32 `json:"participantId"`
	Player        Player `json:"player"`
}

type Player struct {
	MatchHistoryUri string `json:"matchHistoryUri"`
	ProfileIcon     uint32 `json:"profileIcon"`
	SummonerId      uint64 `json:"summonerId"`
	SummonerName    string `json:"summonerName"`
}

type Participant struct {
	ChampionId                uint32              `json:"championId"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
	Masteries                 []Mastery           `json:"masteries"`
	ParticipantId             uint32              `json:"participantId"`
	Runes                     []Rune              `json:"runes"`
	Spell1Id                  uint32              `json:"spell1Id"`
	Spell2Id                  uint32              `json:"spell2Id"`
	Stats                     ParticipantStats    `json:"stats"`
	TeamId                    uint32              `json:"teamId"`
	Timeline                  ParticipantTimeline `json:"timeline"`
}

type Mastery struct {
	MasteryId uint64 `json:"masteryId"`
	Rank      uint64 `json:"rank"`
}

type Rune struct {
	Rank   uint64 `json:"rank"`
	RuneId uint64 `json:"runeId"`
}

type ParticipantStats struct {
	Assists                         uint64 `json:"assists"`
	ChampLevel                      uint64 `json:"champLevel"`
	CombatPlayerScore               uint64 `json:"combatPlayerScore"`
	Deaths                          uint64 `json:"deaths"`
	DoubleKills                     uint64 `json:"doubleKills"`
	FirstBloodAssist                bool   `json:"firstBloodAssist"`
	FirstBloodKill                  bool   `json:"firstBloodKill"`
	FirstInhibitorAssist            bool   `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool   `json:"firstInhibitorKill"`
	FirstTowerAssist                bool   `json:"firstTowerAssist"`
	FirstTowerKill                  bool   `json:"firstTowerKill"`
	GoldEarned                      uint64 `json:"goldEarned"`
	GoldSpent                       uint64 `json:"goldSpent"`
	InhibitorKills                  uint64 `json:"inhibitorKills"`
	Item0                           uint64 `json:"item0"`
	Item1                           uint64 `json:"item1"`
	Item2                           uint64 `json:"item2"`
	Item3                           uint64 `json:"item3"`
	Item4                           uint64 `json:"item4"`
	Item5                           uint64 `json:"item5"`
	Item6                           uint64 `json:"item6"`
	KillingSprees                   uint64 `json:"killingSprees"`
	Kills                           uint64 `json:"kills"`
	LargestCriticalStrike           uint64 `json:"largestCriticalStrike"`
	LargestKillingSpree             uint64 `json:"largestKillingSpree"`
	LargestMultiKill                uint64 `json:"largestMultiKill"`
	MagicDamageDealt                uint64 `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     uint64 `json:"magicDamageDealtToChampions"`
	MagicDamageTaken                uint64 `json:"magicDamageTaken"`
	MinionsKilled                   uint64 `json:"minionsKilled"`
	NeutralMinionsKilled            uint64 `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle uint64 `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  uint64 `json:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     uint64 `json:"nodeCapture"`
	NodeCaptureAssist               uint64 `json:"nodeCaptureAssist"`
	NodeNeutralize                  uint64 `json:"nodeNeutralize"`
	NodeNeutralizeAssist            uint64 `json:"nodeNeutralizeAssist"`
	ObjectivePlayerScore            uint64 `json:"objectivePlayerScore"`
	PentaKills                      uint64 `json:"pentaKills"`
	PhysicalDamageDealt             uint64 `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  uint64 `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             uint64 `json:"physicalDamageTaken"`
	QuadraKills                     uint64 `json:"quadraKills"`
	SightWardsBoughtInGame          uint64 `json:"sightWardsBoughtInGame"`
	TeamObjective                   uint64 `json:"teamObjective"`
	TotalDamageDealt                uint64 `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     uint64 `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                uint64 `json:"totalDamageTaken"`
	TotalHeal                       uint64 `json:"totalHeal"`
	TotalPlayerScore                uint64 `json:"totalPlayerScore"`
	TotalScoreRank                  uint64 `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      uint64 `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                uint64 `json:"totalUnitsHealed"`
	TowerKills                      uint64 `json:"towerKills"`
	TripleKills                     uint64 `json:"tripleKills"`
	TrueDamageDealt                 uint64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      uint64 `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 uint64 `json:"trueDamageTaken"`
	UnrealKills                     uint64 `json:"unrealKills"`
	VisionWardsBoughtInGame         uint64 `json:"visionWardsBoughtInGame"`
	WardsKilled                     uint64 `json:"wardsKilled"`
	WardsPlaced                     uint64 `json:"wardsPlaced"`
	Winner                          bool   `json:"winner"`
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        ParticipantTimelineData `json:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          ParticipantTimelineData `json:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              ParticipantTimelineData `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              ParticipantTimelineData `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         ParticipantTimelineData `json:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       ParticipantTimelineData `json:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         ParticipantTimelineData `json:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    ParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                ParticipantTimelineData `json:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      ParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`
	Lane                            string                  `json:"lane"`
	Role                            string                  `json:"role"`
	TowerAssistsPerMinCounts        ParticipantTimelineData `json:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          ParticipantTimelineData `json:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          ParticipantTimelineData `json:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      ParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        ParticipantTimelineData `json:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               ParticipantTimelineData `json:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              ParticipantTimelineData `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  ParticipantTimelineData `json:"xpPerMinDeltas"`
}

type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty"`
	ThirtyToEnd    float64 `json:"thirtyToEnd"`
	TwentyToThirty float64 `json:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen"`
}

type Team struct {
	Bans                 []BannedChampion `json:"bans"`
	BaronKills           uint32           `json:"baronKills"`
	DominionVictoryScore uint64           `json:"dominionVictoryScore"`
	DragonKills          uint32           `json:"dragonKills"`
	FirstBaron           bool             `json:"firstBaron"`
	FirstBlood           bool             `json:"firstBlood"`
	FirstDragon          bool             `json:"firstDragon"`
	FirstInhibitor       bool             `json:"firstInhibitor"`
	FirstRiftHerald      bool             `json:"firstRiftHerald"`
	FirstTower           bool             `json:"firstTower"`
	InhibitorKills       uint32           `json:"inhibitorKills"`
	RiftHeraldKills      uint32           `json:"riftHeraldKills"`
	TeamId               uint32           `json:"teamId"`
	TowerKills           uint32           `json:"towerKills"`
	VilemawKills         uint32           `json:"vilemawKills"`
	Winner               bool             `json:"winner"`
}

type BannedChampion struct {
	ChampionId uint32 `json:"championId"`
	PickTurn   uint32 `json:"pickTurn"`
}

type Timeline struct {
	FrameInterval uint64  `json:"frameInterval"`
	Frames        []Frame `json:"frames"`
}

type Frame struct {
	Events            []Event                     `json:"events"`
	ParticipantFrames map[string]ParticipantFrame `json:"participantFrames"`
	Timestamp         uint64                      `json:"timestamp"`
}

type Event struct {
	AscendedType            string   `json:"ascendedType"`
	AssistingParticipantIds []uint32 `json:"assistingParticipantIds"`
	BuildingType            string   `json:"buildingType"`
	CreatorId               uint32   `json:"creatorId"`
	EventType               string   `json:"eventType"`
	ItemAfter               uint32   `json:"itemAfter"`
	ItemBefore              uint32   `json:"itemBefore"`
	ItemId                  uint32   `json:"itemId"`
	KillerId                uint32   `json:"killerId"`
	LaneType                string   `json:"laneType"`
	LevelUpType             string   `json:"levelUpType"`
	MonsterType             string   `json:"monsterType"`
	ParticipantId           uint32   `json:"participantId"`
	PointCaptured           string   `json:"pointCaptured"`
	Position                Position `json:"position"`
	SkillSlot               uint32   `json:"skillSlot"`
	TeamId                  uint32   `json:"teamId"`
	Timestamp               uint64   `json:"timestamp"`
	TowerType               string   `json:"towerType"`
	VictimId                uint32   `json:"victimId"`
	WardType                string   `json:"wardType"`
}

type Position struct {
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
}

type ParticipantFrame struct {
	CurrentGold         int32    `json:"currentGold"`
	DominionScore       uint32   `json:"dominionScore"`
	JungleMinionsKilled uint32   `json:"jungleMinionsKilled"`
	Level               uint32   `json:"level"`
	MinionsKilled       uint32   `json:"minionsKilled"`
	ParticipantId       uint32   `json:"participantId"`
	Position            Position `json:"position"`
	TeamScore           uint32   `json:"teamScore"`
	TotalGold           uint32   `json:"totalGold"`
	Xp                  uint32   `json:"xp"`
}
