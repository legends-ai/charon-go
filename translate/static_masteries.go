package translate

import (
	"strings"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func StaticMasteries(
	smm *models.StaticMasteryMap,
) (map[uint32]*apb.CharonData_Static_Mastery, *apb.CharonData_Static_MasteryTrees) {
	return staticMasteryMap(smm), staticMasteryTrees(smm)
}

func staticMasteryMap(smm *models.StaticMasteryMap) map[uint32]*apb.CharonData_Static_Mastery {
	smMap := map[uint32]*apb.CharonData_Static_Mastery{}
	for _, sm := range smm.Data {
		smMap[sm.Id] = &apb.CharonData_Static_Mastery{
			Description: parseDescription(sm.Description, sm.SanitizedDescription),
			Id:          sm.Id,
			Image:       parseImage(sm.Image),
			MasteryTree: apb.MasteryTree(apb.MasteryTree_value[strings.ToUpper(sm.MasteryTree+"_TREE")]),
			Name:        sm.Name,
			Prereq:      sm.Prereq,
			Ranks:       sm.Ranks,
		}
	}
	return smMap
}

func parseDescription(raw []string, sanitized []string) []*apb.VulgateData_League_TextPair {
	tp := []*apb.VulgateData_League_TextPair{}
	for i, r := range raw {
		tp = append(tp, &apb.VulgateData_League_TextPair{
			Raw:       r,
			Sanitized: sanitized[i],
		})
	}
	return tp
}

func staticMasteryTrees(smt *models.StaticMasteryMap) *apb.CharonData_Static_MasteryTrees {
	return &apb.CharonData_Static_MasteryTrees{
		Cunning:  staticMasteryTree(smt.Tree.Cunning),
		Ferocity: staticMasteryTree(smt.Tree.Ferocity),
		Resolve:  staticMasteryTree(smt.Tree.Resolve),
	}
}

func staticMasteryTree(smtrs []models.MasteryTreeRow) []*apb.CharonData_Static_MasteryTrees_Row {
	tree := []*apb.CharonData_Static_MasteryTrees_Row{}
	for _, row := range smtrs {
		tree = append(tree, &apb.CharonData_Static_MasteryTrees_Row{
			Entries: staticMasteryTreeRow(row),
		})
	}
	return tree
}

func staticMasteryTreeRow(smtr models.MasteryTreeRow) []*apb.CharonData_Static_MasteryTrees_Row_Entry {
	row := []*apb.CharonData_Static_MasteryTrees_Row_Entry{}
	for _, item := range smtr.MasteryTreeItems {
		row = append(row, &apb.CharonData_Static_MasteryTrees_Row_Entry{
			Id:     item.MasteryId,
			Prereq: item.Prereq,
		})
	}
	return row
}
