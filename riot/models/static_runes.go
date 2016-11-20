package models

type StaticRuneMap struct {
	Data map[string]struct {
		Description          string      `json:"description"`
		Id                   uint32      `json:"id"`
		Image                StaticImage `json:"image"`
		Name                 string      `json:"name"`
		SanitizedDescription string      `json:"sanitizedDescription"`
		Stats                struct {
			FlatArmorPerLevel               float64 `json:"rFlatArmorModPerLevel"`
			FlatArmorPenetration            float64 `json:"rFlatArmorPenetrationMod"`
			FlatArmorPenetrationPerLevel    float64 `json:"rFlatArmorPenetrationModPerLevel"`
			FlatCritChancePerLevel          float64 `json:"rFlatCritChanceModPerLevel"`
			FlatCritDamagePerLevel          float64 `json:"rFlatCritDamageModPerLevel"`
			FlatDodge                       float64 `json:"rFlatDodgeMod"`
			FlatDodgePerLevel               float64 `json:"rFlatDodgeModPerLevel"`
			FlatEnergyPerLevel              float64 `json:"rFlatEnergyModPerLevel"`
			FlatEnergyRegenPerLevel         float64 `json:"rFlatEnergyRegenModPerLevel"`
			FlatGoldPer10                   float64 `json:"rFlatGoldPer10Mod"`
			FlatHPPerLevel                  float64 `json:"rFlatHPModPerLevel"`
			FlatHPRegenPerLevel             float64 `json:"rFlatHPRegenModPerLevel"`
			FlatMPPerLevel                  float64 `json:"rFlatMPModPerLevel"`
			FlatMPRegenPerLevel             float64 `json:"rFlatMPRegenModPerLevel"`
			FlatMagicDamagePerLevel         float64 `json:"rFlatMagicDamageModPerLevel"`
			FlatMagicPenetration            float64 `json:"rFlatMagicPenetrationMod"`
			FlatMagicPenetrationPerLevel    float64 `json:"rFlatMagicPenetrationModPerLevel"`
			FlatMovementSpeedPerLevel       float64 `json:"rFlatMovementSpeedModPerLevel"`
			FlatPhysicalDamagePerLevel      float64 `json:"rFlatPhysicalDamageModPerLevel"`
			FlatSpellBlockPerLevel          float64 `json:"rFlatSpellBlockModPerLevel"`
			FlatTimeDead                    float64 `json:"rFlatTimeDeadMod"`
			FlatTimeDeadPerLevel            float64 `json:"rFlatTimeDeadModPerLevel"`
			PercentArmorPenetration         float64 `json:"rPercentArmorPenetrationMod"`
			PercentArmorPenetrationPerLevel float64 `json:"rPercentArmorPenetrationModPerLevel"`
			PercentAttackSpeedPerLevel      float64 `json:"rPercentAttackSpeedModPerLevel"`
			PercentCooldown                 float64 `json:"rPercentCooldownMod"`
			PercentCooldownPerLevel         float64 `json:"rPercentCooldownModPerLevel"`
			PercentMagicPenetration         float64 `json:"rPercentMagicPenetrationMod"`
			PercentMagicPenetrationPerLevel float64 `json:"rPercentMagicPenetrationModPerLevel"`
			PercentMovementSpeedPerLevel    float64 `json:"rPercentMovementSpeedModPerLevel"`
			PercentTimeDead                 float64 `json:"rPercentTimeDeadMod"`
			PercentTimeDeadPerLevel         float64 `json:"rPercentTimeDeadModPerLevel"`
		} `json:"stats"`
		Tags []string `json:"tags"`
		Tier string   `json:"tier"`
		Type string   `json:"type"`
	} `json:"data"`
}
