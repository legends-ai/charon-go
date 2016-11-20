package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func StaticRunes(srm *models.StaticRuneMap) map[uint32]*apb.CharonData_Static_Rune {
	srMap := map[uint32]*apb.CharonData_Static_Rune{}
	for _, sr := range srm.Data {
		srMap[sr.Id] = &apb.CharonData_Static_Rune{
			Description: &apb.CharonData_Static_TextPair{
				Raw:       sr.Description,
				Sanitized: sr.SanitizedDescription,
			},
			Id:    sr.Id,
			Image: parseImage(sr.Image),
			Name:  sr.Name,
			Stats: &apb.CharonData_Static_Rune_Stats{
				FlatArmorPerLevel:               sr.Stats.FlatArmorPerLevel,
				FlatArmorPenetration:            sr.Stats.FlatArmorPenetration,
				FlatArmorPenetrationPerLevel:    sr.Stats.FlatArmorPenetrationPerLevel,
				FlatCritChancePerLevel:          sr.Stats.FlatCritChancePerLevel,
				FlatCritDamagePerLevel:          sr.Stats.FlatCritDamagePerLevel,
				FlatDodge:                       sr.Stats.FlatDodge,
				FlatDodgePerLevel:               sr.Stats.FlatDodgePerLevel,
				FlatEnergyPerLevel:              sr.Stats.FlatEnergyPerLevel,
				FlatEnergyRegenPerLevel:         sr.Stats.FlatEnergyRegenPerLevel,
				FlatGoldPer10:                   sr.Stats.FlatGoldPer10,
				Flat_HPPerLevel:                 sr.Stats.FlatHPPerLevel,
				Flat_HPRegenPerLevel:            sr.Stats.FlatHPRegenPerLevel,
				Flat_MPPerLevel:                 sr.Stats.FlatMPPerLevel,
				Flat_MPRegenPerLevel:            sr.Stats.FlatMPRegenPerLevel,
				FlatMagicDamagePerLevel:         sr.Stats.FlatMagicDamagePerLevel,
				FlatMagicPenetration:            sr.Stats.FlatMagicPenetration,
				FlatMagicPenetrationPerLevel:    sr.Stats.FlatMagicPenetrationPerLevel,
				FlatMovementSpeedPerLevel:       sr.Stats.FlatMovementSpeedPerLevel,
				FlatPhysicalDamagePerLevel:      sr.Stats.FlatPhysicalDamagePerLevel,
				FlatSpellBlockPerLevel:          sr.Stats.FlatSpellBlockPerLevel,
				FlatTimeDead:                    sr.Stats.FlatTimeDead,
				FlatTimeDeadPerLevel:            sr.Stats.FlatTimeDeadPerLevel,
				PercentArmorPenetration:         sr.Stats.PercentArmorPenetration,
				PercentArmorPenetrationPerLevel: sr.Stats.PercentArmorPenetrationPerLevel,
				PercentAttackSpeedPerLevel:      sr.Stats.PercentAttackSpeedPerLevel,
				PercentCooldown:                 sr.Stats.PercentCooldown,
				PercentCooldownPerLevel:         sr.Stats.PercentCooldownPerLevel,
				PercentMagicPenetration:         sr.Stats.PercentMagicPenetration,
				PercentMagicPenetrationPerLevel: sr.Stats.PercentMagicPenetrationPerLevel,
				PercentMovementSpeedPerLevel:    sr.Stats.PercentMovementSpeedPerLevel,
				PercentTimeDead:                 sr.Stats.PercentTimeDead,
				PercentTimeDeadPerLevel:         sr.Stats.PercentTimeDeadPerLevel,
			},
			Tags: sr.Tags,
			Tier: sr.Tier,
			Type: sr.Type,
		}
	}
	return srMap
}
