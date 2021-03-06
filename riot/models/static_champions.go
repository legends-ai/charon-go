package models

type StaticChampionMap struct {
	Data map[string]struct {
		AllyTips  []string    `json:"allytips"`
		Blurb     string      `json:"blurb"`
		EnemyTips []string    `json:"enemytips"`
		Id        uint32      `json:"id"`
		Image     StaticImage `json:"image"`
		Info      struct {
			Attack     uint32 `json:"attack"`
			Defense    uint32 `json:"defense"`
			Difficulty uint32 `json:"difficulty"`
			Magic      uint32 `json:"magic"`
		} `json:"info"`
		Key     string `json:"key"`
		Lore    string `json:"lore"`
		Name    string `json:"name"`
		Partype string `json:"partype"`
		Passive struct {
			Description          string      `json:"description"`
			Image                StaticImage `json:"image"`
			Name                 string      `json:"name"`
			SanitizedDescription string      `json:"sanitizedDescription"`
		} `json:"passive"`
		Recommended []ChampionRecommended `json:"recommended"`
		Skins       []ChampionSkin        `json:"skins"`
		Spells      []StaticSpell         `json:"spells"`
		Stats       ChampionStats         `json:"stats"`
		Tags        []string              `json:"tags"`
		Title       string                `json:"title"`
	} `json:"data"`
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
