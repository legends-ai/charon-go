package models

type RiotStaticChampions struct {
	StaticChampions []StaticChampion
}

type StaticChampion struct {
	Allytips    []string              `json:"allytips"`
	Blurb       []string              `json:"blurb"`
	Enemytips   []string              `json:"enemytips"`
	Id          uint32                `json:"id"`
	Image       RiotStaticImage       `json:"image"`
	Info        ChampionInfo          `json:"info"`
	Key         string                `json:"key"`
	Lore        string                `json:"lore"`
	Name        string                `json:"name"`
	Partype     string                `json:"partype"`
	Passive     string                `json:"partype"`
	Recommended []ChampionRecommended `json:"recommended"`
}

type ChampionInfo struct {
	Attack     uint32 `json:"attack"`
	Defense    uint32 `json:"defense"`
	Difficulty uint32 `json:"difficulty"`
	Magic      uint32 `json:"magic"`
}

type ChampionRecommended struct {
	Blocks   []ChampionBlocks `json:"blocks"`
	Champion string           `json:"champion"`
	Map      string           `json:"map"`
	Mode     string           `json:"mode"`
	Priority bool             `json:"priority"`
	Title    string           `json:"string"`
	Type     string           `json:"string"`
}

type ChampionBlocks struct {
	Items   []ChampionBlockItem `json:"items"`
	RecMath bool                `json:"recMath"`
	Type    string              `json:"type"`
}

type ChampionBlockItem struct {
}
