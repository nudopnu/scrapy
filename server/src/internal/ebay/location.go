package ebay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nudopnu/scraper/internal"
)

type LocationData struct {
	SearchOptions struct {
		Q                 string `json:"q"`
		Depth             int    `json:"depth"`
		IncludeParentPath bool   `json:"includeParentPath"`
	} `json:"searchOptions"`
	Locations struct {
		Value struct {
			Location []struct {
				ID string `json:"id"`
			} `json:"location"`
		} `json:"value"`
	} `json:"{http://www.ebayclassifiedsgroup.com/schema/location/v1}locations"`
}

func GetLocationId(postalCode string) (string, error) {
	url := fmt.Sprintf("https://api.kleinanzeigen.de/api/locations.json?depth=1&q=%s", postalCode)
	headers := map[string]string{
		"Accept-Encoding":      "gzip",
		"Authorization":        "Basic YW5kcm9pZDpUYVI2MHBFdHRZ",
		"Connection":           "Keep-Alive",
		"Host":                 "api.kleinanzeigen.de",
		"User-Agent":           "Kleinanzeigen/100.20.0 (Android 9; Asus ASUS_Z01QD)",
		"X-EBAYK-APP":          "38f30879-61bc-4589-bb91-ec1aeb066a8d1728011895290",
		"X-EBAYK-GROUPS":       "BAND-7832-Category-Alerts_B|BAND-8364_B|BAND-8483_composeSlider_A|BLN-19260-cis-login_B|BLN-24652_category_alert_B|backend_ab_bln13364_A|backend_ab_bln418_B|backend_ab_bln_abc_B|backend_ab_bln_abc2_A",
		"X-EBAYK-USERID-TOKEN": "",
		"X-ECG-IN":             "id,localized-name,longitude,latitude,radius,regions",
		"X-ECG-USER-AGENT":     "ebayk-android-app-100.20.0",
		"X-ECG-USER-VERSION":   "100.20.0",
	}
	byteValue, err := internal.Fetch(url, headers)
	if err != nil {
		return "", err
	}

	var locationData LocationData
	err = json.Unmarshal(byteValue, &locationData)
	if err != nil {
		return "", fmt.Errorf("error parsing response data: %w", err)
	}

	// Extract and print the location ID
	if len(locationData.Locations.Value.Location) > 0 {
		locationID := locationData.Locations.Value.Location[0].ID
		return locationID, nil
	}

	return "", errors.New("no location id found")
}
