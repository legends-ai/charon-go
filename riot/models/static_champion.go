package models

type RiotStaticChampions struct {
	Data map[string]struct {
		AllyTips  []string        `json:"allytips"`
		Blurb     string          `json:"blurb"`
		EnemyTips []string        `json:"enemytips"`
		Id        uint32          `json:"id"`
		Image     RiotStaticImage `json:"image"`
		Info      struct {
			Attack     uint32 `json:"attack"`
			Defense    uint32 `json:"defense"`
			Difficulty uint32 `json:"difficulty"`
			Magic      uint32 `json:"magic"`
		}
		Key     string `json:"key"`
		Lore    string `json:"lore"`
		Name    string `json:"name"`
		Partype string `json:"partype"`
		Passive struct {
			Description          string          `json:"description"`
			Image                RiotStaticImage `json:"image"`
			Name                 string          `json:"name"`
			SanitizedDescription string          `json:"sanitizedDescription"`
		}
		Recommended []ChampionRecommended `json:"recommended"`
		Skins       []ChampionSkin        `json:"skins"`
		Spells      []ChampionSpell       `json:"spells"`
		Stats       ChampionStats         `json:"stats"`
		Tags        []string              `json:"tags"`
		Title       string                `json:"title"`
	}
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
	Num  uint32 `json:"num"`
}

type ChampionSpell struct {
	AltImages    []RiotStaticImage `json:"altimages"`
	Cooldown     []float64         `json:"cooldown"`
	CooldownBurn string            `json:"cooldownBurn"`
	Cost         []uint32          `json:"cost"`
	CostBurn     string            `json:"costBurn"`
	CostType     string            `json:"costType"`
	Description  string            `json:"description"`
	Effect       [][]float64       `json:"effect"`
	EffectBurn   []string          `json:"effectBurn"`
	Image        RiotStaticImage   `json:"image"`
	Key          string            `json:"key"`
	LevelTip     struct {
		Effect []string `json:"effect"`
		Label  []string `json:"label"`
	}
	MaxRank              uint32              `json:"maxRank"`
	Name                 string              `json:"string"`
	Range                interface{}         `json:"range"`
	RangeBurn            string              `json:"rangeBurn"`
	Resource             string              `json:"resource"`
	SanitizedDescription string              `json:"sanitizedDescription"`
	SanitizedTooltip     string              `json:"sanitizedTooltip"`
	Tooltip              string              `json:"tooltip"`
	Vars                 []ChampionSpellVars `json:"vars"`
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
	MovementSpeed        float64 `json:"moveSpeed"`
	MP                   float64 `json:"mp"`
	MPPerLevel           float64 `json:"mpperlevel"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
}
