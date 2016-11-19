package riot

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (r *API) StaticVersions() ([]string, error) {
	res, err := r.fetchWithParams(
		"static-version",
		fmt.Sprintf("%s/static-data/%s/v1.2/versions", r.apiLol, r.Region),
		url.Values{},
		false,
	)

	if err != nil {
		return nil, err
	}

	var sv []string
	if err = json.Unmarshal(res, &sv); err != nil {
		return nil, fmt.Errorf("could not unmarshal static versions response: %v", err)
	}

	return sv, nil
}
