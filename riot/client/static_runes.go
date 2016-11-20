package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func (r *API) StaticRunes(locale apb.Locale, version string) (*models.StaticRuneMap, error) {
	res, err := r.fetchWithParams(
		"static-runes",
		fmt.Sprintf("%s/static-data/%s/v1.2/rune", r.apiLol, r.Region),
		url.Values{
			"runeListData": []string{"all"},
			"locale":       []string{locale.String()},
			"version":      []string{version},
		},
		false,
	)

	if err != nil {
		return nil, err
	}

	var sr models.StaticRuneMap
	if err = json.Unmarshal(res, &sr); err != nil {
		return nil, fmt.Errorf("could not unmarshal static runes response: %v", err)
	}

	return &sr, nil
}
