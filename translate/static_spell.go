package translate

import (
	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func StaticSummonerSpells(ssm *models.StaticSpellMap) map[uint32]*apb.CharonData_Static_Spell {
	ssMap := map[uint32]*apb.CharonData_Static_Spell{}
	for _, spell := range ssm.Data {
		ssMap[spell.Id] = parseStaticSpell(spell)
	}
	return ssMap
}

func parseStaticSpellList(spells []models.StaticSpell) []*apb.CharonData_Static_Spell {
	out := []*apb.CharonData_Static_Spell{}
	for _, spell := range spells {
		out = append(out, parseStaticSpell(spell))
	}
	return out
}

func parseStaticSpell(spell models.StaticSpell) *apb.CharonData_Static_Spell {
	return &apb.CharonData_Static_Spell{
		AltImages:    parseSpellAltImages(spell.AltImages),
		Cooldown:     spell.Cooldown,
		CooldownBurn: spell.CooldownBurn,
		Cost:         spell.Cost,
		CostBurn:     spell.CostBurn,
		CostType:     spell.CostType,
		Description: &apb.CharonData_Static_TextPair{
			Raw:       spell.Description,
			Sanitized: spell.SanitizedDescription,
		},
		Effect:     parseSpellEffect(spell.Effect),
		EffectBurn: spell.EffectBurn,
		Id:         spell.Id,
		Image: &apb.CharonData_Static_Image{
			Full:   spell.Image.Full,
			Group:  spell.Image.Group,
			H:      spell.Image.H,
			Sprite: spell.Image.Sprite,
			W:      spell.Image.W,
			X:      spell.Image.X,
			Y:      spell.Image.Y,
		},
		Key: spell.Key,
		LevelTip: &apb.CharonData_Static_Spell_LevelTip{
			Effect: spell.LevelTip.Effect,
			Label:  spell.LevelTip.Label,
		},
		MaxRank:       spell.MaxRank,
		Modes:         spell.Modes,
		Name:          spell.Name,
		Range:         parseRange(spell.Range),
		RangeBurn:     spell.RangeBurn,
		Resource:      spell.Resource,
		SummonerLevel: spell.SummonerLevel,
		Tooltip: &apb.CharonData_Static_TextPair{
			Raw:       spell.Tooltip,
			Sanitized: spell.SanitizedTooltip,
		},
		Vars: parseSpellVars(spell.Vars),
	}
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
