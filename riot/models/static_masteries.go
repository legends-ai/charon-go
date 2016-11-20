package models

type StaticMasteryMap struct {
	Data map[string]struct {
		Description          []string    `json:"description"`
		Id                   uint32      `json:"id"`
		Image                StaticImage `json:"image"`
		MasteryTree          string      `json:"masteryTree"`
		Name                 string      `json:"name"`
		Prereq               string      `json:"prereq"`
		Ranks                uint32      `json:"ranks"`
		SanitizedDescription []string    `json:"sanitizedDescription"`
	} `json:"data"`
	Tree struct {
		Cunning  []MasteryTreeRow `json:"cunning"`
		Ferocity []MasteryTreeRow `json:"ferocity"`
		Resolve  []MasteryTreeRow `json:"resolve"`
	} `json:"tree"`
}

type MasteryTreeRow struct {
	MasteryTreeItems []struct {
		MasteryId uint32 `json:"masteryId"`
		Prereq    string `json:"prereq"`
	} `json:"masteryTreeItems"`
}
