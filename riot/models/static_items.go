package models

type StaticItemMap struct {
	Data map[string]struct {
		Colloq        string            `json:"colloq"`
		ConsumeOnFull bool              `json:"consumeOnFull"`
		Consumed      bool              `json:"consumed"`
		Depth         uint32            `json:"depth"`
		Description   string            `json:"description"`
		Effect        map[string]string `json:"effect"`
		From          []string          `json:"from"`
		Gold          struct {
			Base        uint32 `json:"base"`
			Purchasable bool   `json:"purchasable"`
			Sell        uint32 `json:"sell"`
			Total       uint32 `json:"total"`
		} `json:"gold"`
		Group                string          `json:"group"`
		HideFromAll          bool            `json:"hideFromAll"`
		Id                   uint32          `json:"id"`
		Image                StaticImage     `json:"image"`
		InStore              bool            `json:"inStore"`
		Into                 []string        `json:"into"`
		Maps                 map[string]bool `json:"maps"`
		Name                 string          `json:"name"`
		Plaintext            string          `json:"plaintext"`
		RequiredChampion     string          `json:"requiredChampion"`
		SanitizedDescription string          `json:"sanitizedDescription"`
		SpecialRecipe        uint32          `json:"specialRecipe"`
		Stacks               uint32          `json:"stacks"`
		Stats                struct {
			FlatArmor             float64 `json:"FlatArmorMod"`
			FlatAttackSpeed       float64 `json:"FlatAttackSpeedMod"`
			FlatBlock             float64 `json:"FlatBlockMod"`
			FlatCritChance        float64 `json:"FlatCritChanceMod"`
			FlatCritDamage        float64 `json:"FlatCritDamageMod"`
			FlatEXPBonus          float64 `json:"FlatEXPBonus"`
			FlatEnergyPool        float64 `json:"FlatEnergyPoolMod"`
			FlatEnergyRegen       float64 `json:"FlatEnergyRegenMod"`
			FlatHPPool            float64 `json:"FlatHPPoolMod"`
			FlatHPRegen           float64 `json:"FlatHPRegenMod"`
			FlatMPPool            float64 `json:"FlatMPPoolMod"`
			FlatMPRegen           float64 `json:"FlatMPRegenMod"`
			FlatMagicDamage       float64 `json:"FlatMagicDamageMod"`
			FlatMovementSpeed     float64 `json:"FlatMovementSpeedMod"`
			FlatPhysicalDamage    float64 `json:"FlatPhysicalDamageMod"`
			FlatSpellBlock        float64 `json:"FlatSpellBlockMod"`
			PercentArmor          float64 `json:"PercentArmorMod"`
			PercentAttackSpeed    float64 `json:"PercentAttackSpeedMod"`
			PercentBlock          float64 `json:"PercentBlockMod"`
			PercentCritChance     float64 `json:"PercentCritChanceMod"`
			PercentCritDamage     float64 `json:"PercentCritDamageMod"`
			PercentDodge          float64 `json:"PercentDodgeMod"`
			PercentEXPBonus       float64 `json:"PercentEXPBonus"`
			PercentHPPool         float64 `json:"PercentHPPoolMod"`
			PercentHPRegen        float64 `json:"PercentHPRegenMod"`
			PercentLifeSteal      float64 `json:"PercentLifeStealMod"`
			PercentMPPool         float64 `json:"PercentMPPoolMod"`
			PercentMPRegen        float64 `json:"PercentMPRegenMod"`
			PercentMagicDamage    float64 `json:"PercentMagicDamageMod"`
			PercentMovementSpeed  float64 `json:"PercentMovementSpeedMod"`
			PercentPhysicalDamage float64 `json:"PercentPhysicalDamageMod"`
			PercentSpellBlock     float64 `json:"PercentSpellBlockMod"`
			PercentSpellVamp      float64 `json:"PercentSpellVampMod"`
		} `json:"stats"`
		Tags []string `json:"tags"`
	} `json:"data"`
}
