package PDP

type StockJSON struct {
	Brand      string `json:"brand"`
	Categories []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"categories"`
	Description  string    `json:"description"`
	DropShip     bool      `json:"dropShip"`
	FreeShipping bool      `json:"freeShipping"`
	GiftCosts    []float64 `json:"giftCosts"`
	Images       []struct {
		Code       string `json:"code"`
		Variations []struct {
			AltText string `json:"altText"`
			Format  string `json:"format"`
			URL     string `json:"url"`
		} `json:"variations"`
	} `json:"images"`
	IsNewProduct  bool   `json:"isNewProduct"`
	IsSaleProduct bool   `json:"isSaleProduct"`
	ModelNumber   string `json:"modelNumber"`
	Name          string `json:"name"`
	SellableUnits []struct {
		Attributes []struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"attributes"`
		BarCode         string `json:"barCode"`
		Code            string `json:"code"`
		IsBackOrderable bool   `json:"isBackOrderable"`
		IsPreOrder      bool   `json:"isPreOrder"`
		IsRecaptchaOn   bool   `json:"isRecaptchaOn"`
		Price           struct {
			CurrencyIso            string  `json:"currencyIso"`
			FormattedOriginalPrice string  `json:"formattedOriginalPrice"`
			FormattedValue         string  `json:"formattedValue"`
			OriginalPrice          float64 `json:"originalPrice"`
			Value                  float64 `json:"value"`
		} `json:"price"`
		SingleStoreInventory  bool   `json:"singleStoreInventory"`
		SizeAvailableInStores bool   `json:"sizeAvailableInStores"`
		StockLevelStatus      string `json:"stockLevelStatus"`
	} `json:"sellableUnits"`
	SizeChartGridMap []struct {
		Label string   `json:"label"`
		Sizes []string `json:"sizes"`
	} `json:"sizeChartGridMap"`
	SizeChartImage    string `json:"sizeChartImage"`
	SizeChartTipTx    string `json:"sizeChartTipTx"`
	VariantAttributes []struct {
		Code                           string   `json:"code"`
		DisplayCountDownTimer          bool     `json:"displayCountDownTimer"`
		EligiblePaymentTypesForProduct []string `json:"eligiblePaymentTypesForProduct"`
		FitVariant                     string   `json:"fitVariant"`
		FreeShipping                   bool     `json:"freeShipping"`
		FreeShippingMessage            string   `json:"freeShippingMessage"`
		IsSelected                     bool     `json:"isSelected"`
		LaunchProduct                  bool     `json:"launchProduct"`
		MapEnable                      bool     `json:"mapEnable"`
		Price                          struct {
			CurrencyIso            string  `json:"currencyIso"`
			FormattedOriginalPrice string  `json:"formattedOriginalPrice"`
			FormattedValue         string  `json:"formattedValue"`
			OriginalPrice          float64 `json:"originalPrice"`
			Value                  float64 `json:"value"`
		} `json:"price"`
		RecaptchaOn               bool   `json:"recaptchaOn"`
		Riskified                 bool   `json:"riskified"`
		ShipToAndFromStore        bool   `json:"shipToAndFromStore"`
		ShippingRestrictionExists bool   `json:"shippingRestrictionExists"`
		Sku                       string `json:"sku"`
		SkuExclusions             bool   `json:"skuExclusions"`
		StockLevelStatus          string `json:"stockLevelStatus"`
		WebOnlyLaunch             bool   `json:"webOnlyLaunch"`
		Width                     string `json:"width"`
	} `json:"variantAttributes"`
}


type Stock struct {
	StockLevel string `json:"StockLevel"`
	ID         string `json:"ID"`
	Size       string `json:"Size"`
}