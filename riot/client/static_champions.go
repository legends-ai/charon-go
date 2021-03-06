package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func (r *API) StaticChampions(locale apb.Locale, version string) (*models.StaticChampionMap, error) {
	res, err := r.fetchWithParams(
		"static-champion",
		fmt.Sprintf("%s/static-data/%s/v1.2/champion", r.apiLol, r.Region),
		url.Values{
			"champData": []string{"all"},
			"locale":    []string{locale.String()},
			"version":   []string{version},
		},
		false,
	)

	if err != nil {
		return nil, err
	}

	var sc models.StaticChampionMap
	if err = json.Unmarshal(res, &sc); err != nil {
		return nil, fmt.Errorf("could not unmarshal static champion response: %v", err)
	}

	return &sc, nil
}
