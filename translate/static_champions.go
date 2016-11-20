package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func StaticChampions(scm *models.StaticChampionMap) map[uint32]*apb.CharonData_Static_Champion {
	scMap := map[uint32]*apb.CharonData_Static_Champion{}
	for _, sc := range scm.Data {
		scMap[sc.Id] = &apb.CharonData_Static_Champion{
			AllyTips:  sc.AllyTips,
			Blurb:     sc.Blurb,
			EnemyTips: sc.EnemyTips,
			Id:        sc.Id,
			Image:     parseImage(sc.Image),
			Info: &apb.CharonData_Static_Champion_Info{
				Attack:     sc.Info.Attack,
				Defense:    sc.Info.Defense,
				Difficulty: sc.Info.Difficulty,
				Magic:      sc.Info.Magic,
			},
			Key:     sc.Key,
			Lore:    sc.Lore,
			Name:    sc.Name,
			Partype: sc.Partype,
			Passive: &apb.CharonData_Static_Champion_Passive{
				Description: &apb.VulgateData_League_TextPair{
					Raw:       sc.Passive.Description,
					Sanitized: sc.Passive.SanitizedDescription,
				},
				Image: &apb.VulgateData_League_Image{
					Full:   sc.Passive.Image.Full,
					Group:  sc.Passive.Image.Group,
					H:      sc.Passive.Image.H,
					Sprite: sc.Passive.Image.Sprite,
					W:      sc.Passive.Image.W,
					X:      sc.Passive.Image.X,
					Y:      sc.Passive.Image.Y,
				},
				Name: sc.Passive.Name,
			},
			Recommended: parseRecommended(sc.Recommended),
			Skins:       parseSkins(sc.Skins),
			Spells:      parseSpells(sc.Spells),
			Stats: &apb.CharonData_Static_Champion_Stats{
				Armor:                sc.Stats.Armor,
				ArmorPerLevel:        sc.Stats.ArmorPerLevel,
				AttackDamage:         sc.Stats.AttackDamage,
				AttackDamagePerLevel: sc.Stats.AttackDamagePerLevel,
				AttackRange:          sc.Stats.AttackRange,
				AttackSpeedOffset:    sc.Stats.AttackSpeedOffset,
				Crit:                 sc.Stats.Crit,
				Hp:                   sc.Stats.HP,
				HpPerLevel:           sc.Stats.HPPerLevel,
				HpRegen:              sc.Stats.HPRegen,
				HpRegenPerLevel:      sc.Stats.HPRegenPerLevel,
				MovementSpeed:        sc.Stats.MovementSpeed,
				Mp:                   sc.Stats.MP,
				MpPerLevel:           sc.Stats.MPPerLevel,
				MpRegenPerLevel:      sc.Stats.MPRegenPerLevel,
				SpellBlock:           sc.Stats.SpellBlock,
				SpellBlockPerLevel:   sc.Stats.SpellBlockPerLevel,
			},
			Tags:  sc.Tags,
			Title: sc.Title,
		}
	}

	return scMap
}

func parseRecommended(
	rs []models.ChampionRecommended,
) []*apb.CharonData_Static_Champion_Recommended {
	out := []*apb.CharonData_Static_Champion_Recommended{}
	for _, r := range rs {
		out = append(out, &apb.CharonData_Static_Champion_Recommended{
			Blocks:   parseRecommendedBlocks(r.Blocks),
			Champion: r.Champion,
			Map:      r.Map,
			Mode:     r.Mode,
			Priority: r.Priority,
			Title:    r.Title,
			Type:     r.Type,
		})
	}
	return out
}

func parseRecommendedBlocks(
	rbs []models.ChampionBlocks,
) []*apb.CharonData_Static_Champion_Recommended_Block {
	out := []*apb.CharonData_Static_Champion_Recommended_Block{}
	for _, b := range rbs {
		out = append(out, &apb.CharonData_Static_Champion_Recommended_Block{
			Items:   parseRecommendedBlocksItems(b.Items),
			RecMath: b.RecMath,
			Type:    b.Type,
		})
	}
	return out
}

func parseRecommendedBlocksItems(
	rbis []models.ChampionBlockItem,
) []*apb.CharonData_Static_Champion_Recommended_Block_Item {
	out := []*apb.CharonData_Static_Champion_Recommended_Block_Item{}
	for _, i := range rbis {
		out = append(out, &apb.CharonData_Static_Champion_Recommended_Block_Item{
			Count: i.Count,
			Id:    i.Id,
		})
	}
	return out
}

func parseSkins(
	skins []models.ChampionSkin,
) []*apb.CharonData_Static_Champion_Skin {
	out := []*apb.CharonData_Static_Champion_Skin{}
	for _, s := range skins {
		out = append(out, &apb.CharonData_Static_Champion_Skin{
			Id:   s.Id,
			Name: s.Name,
			Num:  s.Num,
		})
	}
	return out
}

func parseSpells(
	spells []models.ChampionSpell,
) []*apb.CharonData_Static_Champion_Spell {
	out := []*apb.CharonData_Static_Champion_Spell{}
	for _, s := range spells {
		out = append(out, &apb.CharonData_Static_Champion_Spell{
			AltImages:    parseSpellAltImages(s.AltImages),
			Cooldown:     s.Cooldown,
			CooldownBurn: s.CooldownBurn,
			Cost:         s.Cost,
			CostBurn:     s.CostBurn,
			CostType:     s.CostType,
			Description: &apb.VulgateData_League_TextPair{
				Raw:       s.Description,
				Sanitized: s.SanitizedDescription,
			},
			Effect:     parseSpellEffect(s.Effect),
			EffectBurn: s.EffectBurn,
			Image: &apb.VulgateData_League_Image{
				Full:   s.Image.Full,
				Group:  s.Image.Group,
				H:      s.Image.H,
				Sprite: s.Image.Sprite,
				W:      s.Image.W,
				X:      s.Image.X,
				Y:      s.Image.Y,
			},
			Key: s.Key,
			LevelTip: &apb.CharonData_Static_Champion_Spell_LevelTip{
				Effect: s.LevelTip.Effect,
				Label:  s.LevelTip.Label,
			},
			MaxRank:   s.MaxRank,
			Name:      s.Name,
			Range:     parseRange(s.Range),
			RangeBurn: s.RangeBurn,
			Resource:  s.Resource,
			Tooltip: &apb.VulgateData_League_TextPair{
				Raw:       s.Tooltip,
				Sanitized: s.SanitizedTooltip,
			},
			Vars: parseSpellVars(s.Vars),
		})
	}
	return out
}

func parseSpellAltImages(
	images []models.StaticImage,
) []*apb.VulgateData_League_Image {
	out := []*apb.VulgateData_League_Image{}
	for _, image := range images {
		out = append(out, &apb.VulgateData_League_Image{
			Full:   image.Full,
			Group:  image.Group,
			H:      image.H,
			Sprite: image.Sprite,
			W:      image.W,
			X:      image.X,
			Y:      image.Y,
		})
	}
	return out
}

func parseSpellEffect(
	effects [][]float64,
) []*apb.CharonData_Static_Champion_Spell_Effect {
	out := []*apb.CharonData_Static_Champion_Spell_Effect{}
	for _, e := range effects {
		out = append(out, &apb.CharonData_Static_Champion_Spell_Effect{List: e})
	}
	return out
}

func parseRange(rg interface{}) *apb.CharonData_Static_Champion_Spell_Range {
	switch rg.(type) {
	case string:
		return &apb.CharonData_Static_Champion_Spell_Range{
			Value: &apb.CharonData_Static_Champion_Spell_Range_Self{Self: true},
		}
	case []interface{}:
		return &apb.CharonData_Static_Champion_Spell_Range{
			Value: &apb.CharonData_Static_Champion_Spell_Range_Range{
				Range: &apb.CharonData_Static_Champion_Spell_Range_Ranges{
					Range: parseRangeArray(rg.([]interface{})),
				},
			},
		}
	}

	return &apb.CharonData_Static_Champion_Spell_Range{}
}

func parseRangeArray(a []interface{}) []float64 {
	out := []float64{}
	for _, v := range a {
		out = append(out, v.(float64))
	}
	return out
}

func parseSpellVars(
	vars []models.ChampionSpellVars,
) []*apb.CharonData_Static_Champion_Spell_Var {
	out := []*apb.CharonData_Static_Champion_Spell_Var{}
	for _, v := range vars {
		out = append(out, &apb.CharonData_Static_Champion_Spell_Var{
			Coeff:     v.Coeff,
			Dyn:       v.Dyn,
			Key:       v.Key,
			Link:      v.Link,
			RanksWith: v.RanksWith,
		})
	}
	return out
}
