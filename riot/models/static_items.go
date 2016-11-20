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
		}
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
			FlatArmorMod             float64 `json:"FlatArmorMod"`
			FlatAttackSpeedMod       float64 `json:"FlatAttackSpeedMod"`
			FlatBlockMod             float64 `json:"FlatBlockMod"`
			FlatCritChanceMod        float64 `json:"FlatCritChanceMod"`
			FlatCritDamageMod        float64 `json:"FlatCritDamageMod"`
			FlatEXPBonus             float64 `json:"FlatEXPBonus"`
			FlatEnergyPoolMod        float64 `json:"FlatEnergyPoolMod"`
			FlatEnergyRegenMod       float64 `json:"FlatEnergyRegenMod"`
			FlatHPPoolMod            float64 `json:"FlatHPPoolMod"`
			FlatHPRegenMod           float64 `json:"FlatHPRegenMod"`
			FlatMPPoolMod            float64 `json:"FlatMPPoolMod"`
			FlatMPRegenMod           float64 `json:"FlatMPRegenMod"`
			FlatMagicDamageMod       float64 `json:"FlatMagicDamageMod"`
			FlatMovementSpeedMod     float64 `json:"FlatMovementSpeedMod"`
			FlatPhysicalDamageMod    float64 `json:"FlatPhysicalDamageMod"`
			FlatSpellBlockMod        float64 `json:"FlatSpellBlockMod"`
			PercentArmorMod          float64 `json:"PercentArmorMod"`
			PercentAttackSpeedMod    float64 `json:"PercentAttackSpeedMod"`
			PercentBlockMod          float64 `json:"PercentBlockMod"`
			PercentCritChanceMod     float64 `json:"PercentCritChanceMod"`
			PercentCritDamageMod     float64 `json:"PercentCritDamageMod"`
			PercentDodgeMod          float64 `json:"PercentDodgeMod"`
			PercentEXPBonus          float64 `json:"PercentEXPBonus"`
			PercentHPPoolMod         float64 `json:"PercentHPPoolMod"`
			PercentHPRegenMod        float64 `json:"PercentHPRegenMod"`
			PercentLifeStealMod      float64 `json:"PercentLifeStealMod"`
			PercentMPPoolMod         float64 `json:"PercentMPPoolMod"`
			PercentMPRegenMod        float64 `json:"PercentMPRegenMod"`
			PercentMagicDamageMod    float64 `json:"PercentMagicDamageMod"`
			PercentMovementSpeedMod  float64 `json:"PercentMovementSpeedMod"`
			PercentPhysicalDamageMod float64 `json:"PercentPhysicalDamageMod"`
			PercentSpellBlockMod     float64 `json:"PercentSpellBlockMod"`
			PercentSpellVampMod      float64 `json:"PercentSpellVampMod"`
		}
		Tags []string `json:"tags"`
	}
}
