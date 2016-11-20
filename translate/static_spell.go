package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func parseStaticSpells(
	spells []models.StaticSpell,
) []*apb.CharonData_Static_Spell {
	out := []*apb.CharonData_Static_Spell{}
	for _, s := range spells {
		out = append(out, &apb.CharonData_Static_Spell{
			AltImages:    parseSpellAltImages(s.AltImages),
			Cooldown:     s.Cooldown,
			CooldownBurn: s.CooldownBurn,
			Cost:         s.Cost,
			CostBurn:     s.CostBurn,
			CostType:     s.CostType,
			Description: &apb.CharonData_Static_TextPair{
				Raw:       s.Description,
				Sanitized: s.SanitizedDescription,
			},
			Effect:     parseSpellEffect(s.Effect),
			EffectBurn: s.EffectBurn,
			Id:         s.Id,
			Image: &apb.CharonData_Static_Image{
				Full:   s.Image.Full,
				Group:  s.Image.Group,
				H:      s.Image.H,
				Sprite: s.Image.Sprite,
				W:      s.Image.W,
				X:      s.Image.X,
				Y:      s.Image.Y,
			},
			Key: s.Key,
			LevelTip: &apb.CharonData_Static_Spell_LevelTip{
				Effect: s.LevelTip.Effect,
				Label:  s.LevelTip.Label,
			},
			MaxRank:       s.MaxRank,
			Modes:         s.Modes,
			Name:          s.Name,
			Range:         parseRange(s.Range),
			RangeBurn:     s.RangeBurn,
			Resource:      s.Resource,
			SummonerLevel: s.SummonerLevel,
			Tooltip: &apb.CharonData_Static_TextPair{
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
) []*apb.CharonData_Static_Image {
	out := []*apb.CharonData_Static_Image{}
	for _, image := range images {
		out = append(out, &apb.CharonData_Static_Image{
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
) []*apb.CharonData_Static_Spell_Effect {
	out := []*apb.CharonData_Static_Spell_Effect{}
	for _, e := range effects {
		out = append(out, &apb.CharonData_Static_Spell_Effect{List: e})
	}
	return out
}

func parseSpellVars(
	vars []models.StaticSpellVars,
) []*apb.CharonData_Static_Spell_Var {
	out := []*apb.CharonData_Static_Spell_Var{}
	for _, v := range vars {
		out = append(out, &apb.CharonData_Static_Spell_Var{
			Coeff:     v.Coeff,
			Dyn:       v.Dyn,
			Key:       v.Key,
			Link:      v.Link,
			RanksWith: v.RanksWith,
		})
	}
	return out
}
