package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func (r *API) StaticMasteries(locale apb.Locale, version string) (*models.StaticMasteryMap, error) {
	res, err := r.fetchWithParams(
		"static-masteries",
		fmt.Sprintf("%s/static-data/%s/v1.2/mastery", r.apiLol, r.Region),
		url.Values{
			"masteryListData": []string{"all"},
			"locale":          []string{locale.String()},
			"version":         []string{version},
		},
		false,
	)

	if err != nil {
		return nil, err
	}

	var sm models.StaticMasteryMap
	if err = json.Unmarshal(res, &sm); err != nil {
		return nil, fmt.Errorf("could not unmarshal static masteries response: %v", err)
	}

	return &sm, nil
}
