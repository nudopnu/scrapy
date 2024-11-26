package ebay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/nudopnu/scraper/internal"
)

type GetAdsRequest struct {
	Keyword    string
	PageNumber int
	LocationId string
	Distance   int
}

type ParsedAd struct {
	Id          string
	Title       string
	Price       string
	Description string
	Location    string
	ZipCode     string
	CategoryId  string
	Date        string
	Pictures    []map[string]string
	Link        string
}

func GetAdsRaw(getAdsRequest GetAdsRequest) ([]byte, error) {
	url := fmt.Sprintf("https://api.kleinanzeigen.de/api/ads.json?_in=id,title,description,displayoptions,start-date-time,category.id,category.localized_name,ad-address.state,ad-address.zip-code,price,pictures,link,features-active,search-distance,negotiation-enabled,attributes,medias,medias.media,medias.media.title,medias.media.media-link,buy-now,placeholder-image-present,store-id,store-title&q=%s&page=%d&sortType=DATE_DESCENDING&size=31&locationId=%s&pictureRequired=false&distance=%d&includeTopAds=true&buyNowOnly=false&limitTotalResultCount=true", getAdsRequest.Keyword, getAdsRequest.PageNumber, getAdsRequest.LocationId, getAdsRequest.Distance)
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
		log.Fatal(err)
	}
	return byteValue, nil
}

func GetAdsMock(GetAdsRequest) ([]byte, error) {
	path, err := filepath.Abs("./samples/ads.json")
	if err != nil {
		return nil, fmt.Errorf("error locating file: %w", err)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening mock file: %w", err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading mock file: %w", err)
	}
	return bytes, nil
}

func UnmarshalAds(bytes []byte) (AdsResult, error) {
	var result AdsResult
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return AdsResult{}, fmt.Errorf("error parsing result: %w", err)
	}
	return result, nil
}

func ParseAds(adsResult AdsResult) []ParsedAd {
	adList := adsResult.HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1Ads.Value.Ad
	parsedAds := make([]ParsedAd, 0, len(adList))
	for _, ad := range adList {
		link, err := getLink(ad.Link)
		if err != nil {
			fmt.Println(err)
			continue
		}
		parsedAd := ParsedAd{
			Id:          ad.ID,
			Title:       ad.Title.Value,
			Price:       getPrice(ad.Price),
			Description: ad.Description.Value,
			Location:    ad.AdAddress.State.Value,
			ZipCode:     ad.AdAddress.ZipCode.Value,
			Date:        ad.StartDateTime.Value,
			CategoryId:  ad.Category.ID,
			Pictures:    getPictures(ad.Pictures),
			Link:        link,
		}
		parsedAds = append(parsedAds, parsedAd)
	}
	return parsedAds
}

func getPrice(price Price) string {
	res := ""
	if price.Amount != nil {
		res += fmt.Sprintf("%d%s", price.Amount.Value, price.CurrencyISOCode.Value.LocalizedLabel)
	}
	if price.PriceType.Value == "PLEASE_CONTACT" {
		res += " VB"
	}
	return res
}

func getPictures(pictures Pictures) []map[string]string {
	pics := make([]map[string]string, 0, len(pictures.Picture))
	for _, picture := range pictures.Picture {
		sizes := make(map[string]string)
		pics = append(pics, sizes)
		for _, link := range picture.Link {
			sizes[string(link.Rel)] = link.Href
		}
	}
	return pics
}

func getLink(links []Link) (string, error) {
	for _, link := range links {
		if link.Rel == "self-public-website" {
			return link.Href, nil
		}
	}
	return "", errors.New("no public link available")
}
