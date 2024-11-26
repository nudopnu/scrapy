package ebay

type AdsResult struct {
	SearchOptions                               SearchOptions                               `json:"searchOptions"`
	HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1Ads HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1Ads `json:"{http://www.ebayclassifiedsgroup.com/schema/ad/v1}ads"`
}

type HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1Ads struct {
	Name            string                                           `json:"name"`
	DeclaredType    string                                           `json:"declaredType"`
	Scope           string                                           `json:"scope"`
	Value           HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1AdsValue `json:"value"`
	Nil             bool                                             `json:"nil"`
	GlobalScope     bool                                             `json:"globalScope"`
	TypeSubstituted bool                                             `json:"typeSubstituted"`
}

type HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1AdsValue struct {
	Ad                      []Ad                    `json:"ad"`
	Paging                  Paging                  `json:"paging"`
	PageLayout              string                  `json:"page-layout"`
	AdsSearchOptions        AdsSearchOptions        `json:"ads-search-options"`
	AdsSearchHistograms     AdsSearchHistograms     `json:"ads-search-histograms"`
	AdsSearchResultMetadata AdsSearchResultMetadata `json:"ads-search-result-metadata"`
}

type Ad struct {
	Price                   Price          `json:"price,omitempty"`
	Title                   BuyNowOnly     `json:"title"`
	Description             BuyNowOnly     `json:"description"`
	AdAddress               AdAddress      `json:"ad-address"`
	SearchDistance          SearchDistance `json:"search-distance"`
	StoreID                 BuyNowOnly     `json:"store-id"`
	StoreTitle              BuyNowOnly     `json:"store-title"`
	PlaceholderImagePresent BuyNowOnly     `json:"placeholder-image-present"`
	StartDateTime           BuyNowOnly     `json:"start-date-time"`
	FeaturesActive          FeaturesActive `json:"features-active"`
	Category                Category       `json:"category"`
	BuyNow                  BuyNow         `json:"buy-now"`
	Attributes              Attributes     `json:"attributes,omitempty"`
	Link                    []Link         `json:"link"`
	Pictures                Pictures       `json:"pictures"`
	Displayoptions          Displayoptions `json:"displayoptions"`
	Medias                  Medias         `json:"medias"`
	ID                      string         `json:"id"`
}

type AdAddress struct {
	State   BuyNowOnly `json:"state"`
	ZipCode BuyNowOnly `json:"zip-code"`
}

type BuyNowOnly struct {
	Value string `json:"value,omitempty"`
}

type Attributes struct {
	Attribute []Attribute `json:"attribute"`
}

type Attribute struct {
	Value             []ValueElement `json:"value"`
	Name              Name           `json:"name"`
	Unit              string         `json:"unit"`
	SearchDisplay     string         `json:"search-display"`
	FakeSubCategory   string         `json:"fake-sub-category"`
	SearchMultiSelect string         `json:"search-multi-select"`
	Type              Type           `json:"type"`
	LocalizedLabel    LocalizedLabel `json:"localized-label"`
	LocalizedTag      string         `json:"localized-tag,omitempty"`
}

type ValueElement struct {
	Value          string `json:"value"`
	LocalizedLabel string `json:"localized-label"`
}

type BuyNow struct {
	Selected string `json:"selected"`
}

type Category struct {
	Link     []interface{} `json:"link"`
	Category []interface{} `json:"category"`
	ID       string        `json:"id"`
}

type Displayoptions struct {
	ReducedAdsOnVip                BuyNowOnly `json:"reduced-ads-on-vip"`
	HideAdsOnVip                   BuyNowOnly `json:"hide-ads-on-vip"`
	ShowRatings                    BuyNowOnly `json:"show-ratings"`
	AdFlaggingEnabled              BuyNowOnly `json:"ad-flagging-enabled"`
	HideSimilardAdsOnVip           BuyNowOnly `json:"hide-similard-ads-on-vip"`
	CategoryChangeAllowed          BuyNowOnly `json:"category-change-allowed"`
	OfferAllowed                   BuyNowOnly `json:"offer-allowed"`
	SecurePaymentPossible          BuyNowOnly `json:"secure-payment-possible"`
	AdDescriptionFormattingEnabled BuyNowOnly `json:"ad-description-formatting-enabled"`
	AdDocumentsUploadEnabled       BuyNowOnly `json:"ad-documents-upload-enabled"`
	DisplayMapOnVip                BuyNowOnly `json:"display-map-on-vip"`
	LicencePlateBlurringEnabled    BuyNowOnly `json:"licence-plate-blurring-enabled"`
	ContactPosterFeedback          string     `json:"contact-poster-feedback"`
	LeadGenerationTarget           BuyNowOnly `json:"lead-generation-target"`
	EnableImageTitles              BuyNowOnly `json:"enable-image-titles"`
}

type FeaturesActive struct {
	FeatureActive []interface{} `json:"feature-active"`
}

type Link struct {
	Href string `json:"href"`
	Rel  Rel    `json:"rel"`
}

type Medias struct {
	Media []interface{} `json:"media"`
}

type Pictures struct {
	Picture []Picture `json:"picture"`
}

type Picture struct {
	Link     []Link   `json:"link"`
	Viewport Viewport `json:"viewport,omitempty"`
}

type Viewport struct {
	MarginTop    float64 `json:"marginTop"`
	MarginBottom float64 `json:"marginBottom"`
	MarginRight  float64 `json:"marginRight"`
	MarginLeft   float64 `json:"marginLeft"`
}

type Price struct {
	CurrencyISOCode CurrencyISOCode `json:"currency-iso-code"`
	Amount          *Amount         `json:"amount"`
	PriceType       BuyNowOnly      `json:"price-type"`
}

type Amount struct {
	Value int64 `json:"value,omitempty"`
}

type CurrencyISOCode struct {
	Value ValueElement `json:"value"`
}

type SearchDistance struct {
	Distance        BuyNowOnly `json:"distance"`
	DistanceUnit    BuyNowOnly `json:"distance-unit"`
	DisplayDistance BuyNowOnly `json:"display-distance"`
}

type AdsSearchHistograms struct {
	AdsCategoryHistogram AdsCategoryHistogram `json:"ads-category-histogram"`
}

type AdsCategoryHistogram struct {
	AdsCategory []AdsCategory `json:"ads-category"`
}

type AdsCategory struct {
	Value string `json:"value"`
	ID    string `json:"id"`
}

type AdsSearchOptions struct {
	Q                       BuyNowOnly `json:"q"`
	LocationID              BuyNowOnly `json:"locationId"`
	Page                    BuyNowOnly `json:"page"`
	Size                    BuyNowOnly `json:"size"`
	LimitTotalResultCount   BuyNowOnly `json:"limitTotalResultCount"`
	Distance                BuyNowOnly `json:"distance"`
	DistanceUnit            BuyNowOnly `json:"distanceUnit"`
	SortType                BuyNowOnly `json:"sortType"`
	PictureRequired         BuyNowOnly `json:"pictureRequired"`
	IncludeTopAds           BuyNowOnly `json:"includeTopAds"`
	Histograms              Histograms `json:"histograms"`
	BuyNowOnly              BuyNowOnly `json:"buyNowOnly"`
	LabelsGenerationEnabled BuyNowOnly `json:"labelsGenerationEnabled"`
}

type Histograms struct {
	Value []string `json:"value"`
}

type AdsSearchResultMetadata struct {
	Title              BuyNowOnly `json:"title"`
	Description        BuyNowOnly `json:"description"`
	DominantCategoryID string     `json:"dominant-category-id"`
}

type Paging struct {
	NumFound      string `json:"numFound"`
	ExactNumFound string `json:"exactNumFound"`
	Link          []Link `json:"link"`
}

type SearchOptions struct {
	Q                       string  `json:"q"`
	Page                    int64   `json:"page"`
	Size                    int64   `json:"size"`
	LimitTotalResultCount   bool    `json:"limitTotalResultCount"`
	Extension               Attr    `json:"extension"`
	LocationID              []int64 `json:"locationId"`
	Distance                float64 `json:"distance"`
	Attr                    Attr    `json:"attr"`
	SortType                string  `json:"sortType"`
	PictureRequired         bool    `json:"pictureRequired"`
	IncludeTopAds           bool    `json:"includeTopAds"`
	LastSearchPush          bool    `json:"lastSearchPush"`
	BuyNowOnly              bool    `json:"buyNowOnly"`
	LabelsGenerationEnabled bool    `json:"labelsGenerationEnabled"`
	UserInventorySearch     bool    `json:"userInventorySearch"`
	BizBrandingSearch       bool    `json:"bizBrandingSearch"`
}

type Attr struct {
}

type LocalizedLabel string

const (
	Art     LocalizedLabel = "Art"
	Typ     LocalizedLabel = "Typ"
	Versand LocalizedLabel = "Versand"
	Zustand LocalizedLabel = "Zustand"
)

type Name string

const (
	AutoteileReifenArt     Name = "autoteile_reifen.art"
	AutoteileReifenVersand Name = "autoteile_reifen.versand"
	FahrraederArt          Name = "fahrraeder.art"
	FahrraederCondition    Name = "fahrraeder.condition"
	FahrraederType         Name = "fahrraeder.type"
	FahrraederVersand      Name = "fahrraeder.versand"
)

type Type string

const (
	Enum Type = "ENUM"
)

type Rel string

const (
	CanonicalURL               Rel = "canonicalUrl"
	ExtraLarge                 Rel = "extraLarge"
	Large                      Rel = "large"
	Next                       Rel = "next"
	Self                       Rel = "self"
	SelfPublicWebsite          Rel = "self-public-website"
	SelfPublicWebsiteCanonical Rel = "self-public-website-canonical"
	Teaser                     Rel = "teaser"
	Thumbnail                  Rel = "thumbnail"
	Xxl                        Rel = "XXL"
)
