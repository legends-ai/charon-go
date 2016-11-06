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
	Skins       []ChampionSkin        `json:"skins"`
	Spells      []ChampionSpell       `json:"spells"`
	Stats       []ChampionStats       `json:"stats"`
	Tags        []string              `json:"tags"`
	Title       string                `json:"title"`
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
	Count uint32 `json:"count"`
	Id    uint32 `json:"id"`
}

type ChampionSkin struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
	Num  string `json:"num"`
}

type ChampionSpell struct {
	Altimages            []RiotStaticImage   `json:"altimages"`
	Cooldown             []float64           `json:"cooldown"`
	CooldownBurn         string              `json:"cooldownBurn"`
	Cost                 []uint32            `json:"cost"`
	CostBurn             []string            `json:"costBurn"`
	CostType             string              `json:"costType"`
	Description          string              `json:"description"`
	Effect               [][]float64         `json:"effect"`
	EffectBurn           []string            `json:"effectBurn"`
	Image                RiotStaticImage     `json:"image"`
	Key                  string              `json:"key"`
	LevelTip             ChampionLevelTip    `json:"levelTip"`
	MaxRank              uint32              `json:"maxRank"`
	Name                 string              `json:"string"`
	Range                []string            `json:"range"`
	RangeBurn            string              `json:"rangeBurn"`
	Resource             string              `json:"resource"`
	SanitizedDescription string              `json:"sanitizedDescription"`
	SanitizedTooltip     string              `json:"sanitizedTooltip"`
	Vars                 []ChampionSpellVars `json:"vars"`
}

type ChampionLevelTip struct {
	Effect []string `json:"effect"`
	Label  []string `json:"label"`
}

type ChampionSpellVars struct {
	Coeff     []float64 `json:"coeff"`
	Dyn       string    `json:"dyn"`
	Key       string    `json:"key"`
	Link      string    `json:"link"`
	RanksWith string    `json:"ranksWith"`
}

type ChampionStats struct {
	Armor                float64 `json:"armor"`
	ArmorPerLevel        float64 `json:"armorperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	AttackRange          float64 `json:"attackrange"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset"`
	Crit                 float64 `json:"crit"`
	HP                   float64 `json:"hp"`
	HPPerLevel           float64 `json:"hpperlevel"`
	HPRegen              float64 `json:"hpregen"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	MoveSpeed            float64 `json:"moveSpeed"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
}
