package models

type StaticSpell struct {
	AltImages    []StaticImage `json:"altimages"` // Champion Only
	Cooldown     []float64     `json:"cooldown"`
	CooldownBurn string        `json:"cooldownBurn"`
	Cost         []uint32      `json:"cost"`
	CostBurn     string        `json:"costBurn"`
	CostType     string        `json:"costType"`
	Description  string        `json:"description"`
	Effect       [][]float64   `json:"effect"`
	EffectBurn   []string      `json:"effectBurn"`
	Id           uint32        `json:"id"` // Summoner Only
	Image        StaticImage   `json:"image"`
	Key          string        `json:"key"`
	LevelTip     struct {
		Effect []string `json:"effect"`
		Label  []string `json:"label"`
	} `json:"leveltip"`
	MaxRank              uint32            `json:"maxRank"`
	Modes                []string          `json:"modes"` // Summoner Only
	Name                 string            `json:"string"`
	Range                interface{}       `json:"range"`
	RangeBurn            string            `json:"rangeBurn"`
	Resource             string            `json:"resource"`
	SanitizedDescription string            `json:"sanitizedDescription"`
	SanitizedTooltip     string            `json:"sanitizedTooltip"`
	SummonerLevel        uint32            `json:"summonerLevel"` // Summoner Only
	Tooltip              string            `json:"tooltip"`
	Vars                 []StaticSpellVars `json:"vars"`
}

type StaticSpellVars struct {
	Coeff     []float64 `json:"coeff"`
	Dyn       string    `json:"dyn"`
	Key       string    `json:"key"`
	Link      string    `json:"link"`
	RanksWith string    `json:"ranksWith"`
}
