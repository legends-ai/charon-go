package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func StaticItems(sim *models.StaticItemMap) map[uint32]*apb.CharonData_Static_Item {
	siMap := map[uint32]*apb.CharonData_Static_Item{}
	for _, si := range sim.Data {
		siMap[si.Id] = &apb.CharonData_Static_Item{
			Colloq:        si.Colloq,
			ConsumeOnFull: si.ConsumeOnFull,
			Consumed:      si.Consumed,
			Depth:         si.Depth,
			Description: &apb.CharonData_Static_TextPair{
				Raw:       si.Description,
				Sanitized: si.SanitizedDescription,
			},
			Effect: si.Effect,
			From:   si.From,
			Gold: &apb.CharonData_Static_Item_Gold{
				Base:        si.Gold.Base,
				Purchasable: si.Gold.Purchasable,
				Sell:        si.Gold.Sell,
				Total:       si.Gold.Total,
			},
			Group:            si.Group,
			HideFromAll:      si.HideFromAll,
			Id:               si.Id,
			Image:            parseImage(si.Image),
			InStore:          si.InStore,
			Into:             si.Into,
			Maps:             si.Maps,
			Plaintext:        si.Plaintext,
			RequiredChampion: si.RequiredChampion,
			SpecialRecipe:    si.SpecialRecipe,
			Stacks:           si.Stacks,
			Stats: &apb.CharonData_Static_Item_Stats{
				FlatArmor:             si.Stats.FlatArmor,
				FlatAttackSpeed:       si.Stats.FlatAttackSpeed,
				FlatBlock:             si.Stats.FlatBlock,
				FlatCritChance:        si.Stats.FlatCritChance,
				FlatCritDamage:        si.Stats.FlatCritDamage,
				FlatEnergyRegen:       si.Stats.FlatEnergyRegen,
				FlatEnergyPool:        si.Stats.FlatEnergyPool,
				Flat_XPBonus:          si.Stats.FlatEXPBonus,
				Flat_HPPool:           si.Stats.FlatHPPool,
				Flat_HPRegen:          si.Stats.FlatHPRegen,
				Flat_MPPool:           si.Stats.FlatMPPool,
				Flat_MPRegen:          si.Stats.FlatMPRegen,
				FlatMagicDamage:       si.Stats.FlatMagicDamage,
				FlatMovementSpeed:     si.Stats.FlatMovementSpeed,
				FlatPhysicalDamage:    si.Stats.FlatPhysicalDamage,
				FlatSpellBlock:        si.Stats.FlatSpellBlock,
				PercentArmor:          si.Stats.PercentArmor,
				PercentAttackSpeed:    si.Stats.PercentAttackSpeed,
				PercentBlock:          si.Stats.PercentBlock,
				PercentCritChance:     si.Stats.PercentCritChance,
				PercentCritDamage:     si.Stats.PercentCritDamage,
				PercentDodge:          si.Stats.PercentCritDamage,
				Percent_XPBonus:       si.Stats.PercentEXPBonus,
				Percent_HPPool:        si.Stats.PercentHPPool,
				Percent_HPRegen:       si.Stats.PercentHPRegen,
				Percent_MPPool:        si.Stats.PercentMPPool,
				Percent_MPRegen:       si.Stats.PercentMPRegen,
				PercentMagicDamage:    si.Stats.PercentMagicDamage,
				PercentMovementSpeed:  si.Stats.PercentMovementSpeed,
				PercentPhysicalDamage: si.Stats.PercentPhysicalDamage,
				PercentSpellBlock:     si.Stats.PercentSpellBlock,
				PercentSpellVamp:      si.Stats.PercentSpellVamp,
				PercentLifeSteal:      si.Stats.PercentLifeSteal,
			},
			Tags: si.Tags,
		}
	}
	return siMap
}
