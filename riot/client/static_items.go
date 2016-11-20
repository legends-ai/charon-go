package riot

import (
	"encoding/json"
	"fmt"
	"net/url"

	apb "github.com/asunaio/charon/gen-go/asuna"
	"github.com/asunaio/charon/riot/models"
)

func (r *API) StaticItems(locale apb.Locale, version string) (*models.StaticItemMap, error) {
	res, err := r.fetchWithParams(
		"static-items",
		fmt.Sprintf("%s/static-data/%s/v1.2/item", r.apiLol, r.Region),
		url.Values{
			"itemListData": []string{"all"},
			"locale":       []string{locale.String()},
			"version":      []string{version},
		},
		false,
	)

	if err != nil {
		return nil, err
	}

	var si models.StaticItemMap
	if err = json.Unmarshal(res, &si); err != nil {
		return nil, fmt.Errorf("could not unmarshal static items response: %v", err)
	}

	return &si, nil
}
