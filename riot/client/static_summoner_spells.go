package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func (r *API) StaticSummonerSpells(locale apb.Locale, version string) (*models.StaticSpellMap, error) {
	res, err := r.fetchWithParams(
		"static-summoner-spells",
		fmt.Sprintf("%s/static-data/%s/v1.2/summoner-spell", r.apiLol, r.Region),
		url.Values{
			"dataById":  []string{"true"},
			"spellData": []string{"all"},
			"locale":    []string{locale.String()},
			"version":   []string{version},
		},
		false,
	)

	if err != nil {
		return nil, err
	}

	var ss models.StaticSpellMap
	if err = json.Unmarshal(res, &ss); err != nil {
		return nil, fmt.Errorf("could not unmarshal static summoner spells response: %v", err)
	}

	return &ss, nil
}
