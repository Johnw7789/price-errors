package bestbuy

type OpenBoxResponse struct {
	CurrentAsOfDate     string  `json:"currentAsOfDate"`
	CurrentPrice        float64 `json:"currentPrice"`
	CustomerPrice       float64 `json:"customerPrice"`
	DotComDisplayStatus string  `json:"dotComDisplayStatus"`
	FinanceOption       struct {
		DefaultPlan                bool    `json:"defaultPlan"`
		FinanceCode                int64   `json:"financeCode"`
		FinanceCodeDescLong        string  `json:"financeCodeDescLong"`
		FinanceCodeName            string  `json:"financeCodeName"`
		FinanceTerm                int64   `json:"financeTerm"`
		MonthlyPayment             float64 `json:"monthlyPayment"`
		MonthlyPaymentIncludingTax float64 `json:"monthlyPaymentIncludingTax"`
		OfferID                    string  `json:"offerId"`
		PlanType                   string  `json:"planType"`
		Priority                   int64   `json:"priority"`
		Rank                       int64   `json:"rank"`
		Rate                       int64   `json:"rate"`
		TermsAndConditions         string  `json:"termsAndConditions"`
		TotalCost                  float64 `json:"totalCost"`
		TotalCostIncludingTax      float64 `json:"totalCostIncludingTax"`
	} `json:"financeOption"`
	GiftSkus []struct {
		IsRequiredWithOffer bool    `json:"isRequiredWithOffer"`
		OfferID             string  `json:"offerId"`
		Quantity            int64   `json:"quantity"`
		Savings             float64 `json:"savings"`
		SkuID               string  `json:"skuId"`
	} `json:"giftSkus"`
	GspUnitPrice          float64 `json:"gspUnitPrice"`
	HasTotalTechWarranty  bool    `json:"hasTotalTechWarranty"`
	InstantSavings        int64   `json:"instantSavings"`
	IsMAP                 bool    `json:"isMAP"`
	IsPriceMatchGuarantee bool    `json:"isPriceMatchGuarantee"`
	LocationID            string  `json:"locationId"`
	OfferQualifiers       []struct {
		ComOfferType              string `json:"comOfferType"`
		ComRuleType               string `json:"comRuleType"`
		ExcludeFromBundleBreakage bool   `json:"excludeFromBundleBreakage"`
		ID                        int64  `json:"id"`
		InstanceID                int64  `json:"instanceId"`
		OfferDiscountType         string `json:"offerDiscountType"`
		OfferID                   string `json:"offerId"`
		OfferName                 string `json:"offerName"`
		OfferRevocableOnReturns   bool   `json:"offerRevocableOnReturns"`
		OfferVersion              int64  `json:"offerVersion"`
	} `json:"offerQualifiers"`
	OpenBoxCondition    int64   `json:"openBoxCondition"`
	OpenBoxPrice        float64 `json:"openBoxPrice"`
	OpenBoxSavings      int64   `json:"openBoxSavings"`
	PreferredBadging    string  `json:"preferredBadging"`
	PriceEventType      string  `json:"priceEventType"`
	RegularPrice        float64 `json:"regularPrice"`
	ShowTotalSavings    bool    `json:"showTotalSavings"`
	SkuID               string  `json:"skuId"`
	TotalGiftSavings    float64 `json:"totalGiftSavings"`
	TotalSavings        int64   `json:"totalSavings"`
	TotalSavingsPercent int64   `json:"totalSavingsPercent"`
}

type ItemDiscountResponse struct {
	ShowBuyNowPayLater              bool          `json:"showBuyNowPayLater"`
	PaymentOptions                  []interface{} `json:"paymentOptions"`
	ShowCustomerPrice               bool          `json:"showCustomerPrice"`
	CustomerPrice                   float64       `json:"customerPrice"`
	ShowFavorFree                   bool          `json:"showFavorFree"`
	ShowTotalTechInfo               bool          `json:"showTotalTechInfo"`
	ShowLivelyPricingInfo           bool          `json:"showLivelyPricingInfo"`
	ShowPrefix                      bool          `json:"showPrefix"`
	ShowPerMonth                    bool          `json:"showPerMonth"`
	ShowPriceDisclaimer             bool          `json:"showPriceDisclaimer"`
	ShowSaleMessage                 bool          `json:"showSaleMessage"`
	AbbreviatedRegularPrice         string        `json:"abbreviatedRegularPrice"`
	Description                     string        `json:"description"`
	IsAlone                         bool          `json:"isAlone"`
	ShowSavingsRegularPrice         bool          `json:"showSavingsRegularPrice"`
	MessagePrefix                   string        `json:"messagePrefix"`
	MessageSuffix                   string        `json:"messageSuffix"`
	MessagePrefixWithSavingsPercent string        `json:"messagePrefixWithSavingsPercent"`
	Total                           int           `json:"total"`
	ShowFinanceOption               bool          `json:"showFinanceOption"`
	TotalSavingsPercent             int           `json:"totalSavingsPercent"`
	ShowPercentOff                  bool          `json:"showPercentOff"`
	DisableSavingsPercent           bool          `json:"disableSavingsPercent"`
	HasMobileContracts              bool          `json:"hasMobileContracts"`
	DisableGreenSavingsBrick        bool          `json:"disableGreenSavingsBrick"`
	ShowContextualPriceMessaging    bool          `json:"showContextualPriceMessaging"`
	DealExpirationTimeStamp         interface{}   `json:"dealExpirationTimeStamp"`
	DisableExpirationTimer          bool          `json:"disableExpirationTimer"`
	ExpirationDateThreshold         int           `json:"expirationDateThreshold"`
	CountDownThreshold              int           `json:"countDownThreshold"`
	ShowBundleAndSave               bool          `json:"showBundleAndSave"`
	Href                            string        `json:"href"`
	DataTrack                       string        `json:"dataTrack"`
	ShowMyBByPrice                  bool          `json:"showMyBByPrice"`
	GwpEmbedClientSideOnly          bool          `json:"gwpEmbedClientSideOnly"`
	Layout                          string        `json:"layout"`
	SkuID                           string        `json:"skuId"`
	MmtPricing                      bool          `json:"mmtPricing"`
	ShowLink                        bool          `json:"showLink"`
	ContextParam                    string        `json:"contextParam"`
	ShowPriceBadge                  bool          `json:"showPriceBadge"`
	ShowPreviousPrice               bool          `json:"showPreviousPrice"`
	ShowPaidMemberBadge             bool          `json:"showPaidMemberBadge"`
	ShowPMGBadge                    bool          `json:"showPMGBadge"`
	PreferredBadgingStr             string        `json:"preferredBadgingStr"`
	ShowMyBBYOffer                  bool          `json:"showMyBBYOffer"`
	ShowHolidayMessageBadge         bool          `json:"showHolidayMessageBadge"`
	OverridePaidMemberBadgeHTML     string        `json:"overridePaidMemberBadgeHTML"`
	ShowOffers                      bool          `json:"showOffers"`
	IsMobileContractSelected        bool          `json:"isMobileContractSelected"`
	ShowSuco                        bool          `json:"showSuco"`
	ShowSmallPrice                  bool          `json:"showSmallPrice"`
	ShowTotalCostClarity            bool          `json:"showTotalCostClarity"`
	ShowBlackFridayHolidayMessage   bool          `json:"showBlackFridayHolidayMessage"`
	ShowLowestPriceOfTheSeason      bool          `json:"showLowestPriceOfTheSeason"`
	PriceChangeTotalSavingsAmount   int           `json:"priceChangeTotalSavingsAmount"`
	PriceChangeTotalSavingsPercent  int           `json:"priceChangeTotalSavingsPercent"`
}

type QueryRecResponse struct {
	SuggestionResponse struct {
		Count             int `json:"count"`
		QueryResponseTime int `json:"queryResponseTime"`
		Track             struct {
			Location struct {
				Applied   string `json:"applied"`
				Requested string `json:"requested"`
			} `json:"location"`
			Variant string `json:"variant"`
		} `json:"track"`
		SpellCheck struct {
			CorrectlySpelled bool   `json:"correctlySpelled"`
			CorrectedQuery   string `json:"correctedQuery"`
			OriginalQuery    string `json:"originalQuery"`
		} `json:"spellCheck"`
		Suggestions []struct {
			Products []struct {
				SkuID  string `json:"skuId"`
				Source string `json:"source"`
				Type   string `json:"type"`
			} `json:"products"`
			Category []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"category,omitempty"`
			Term string `json:"term"`
		} `json:"suggestions"`
	} `json:"suggestionResponse"`
}

type DealInfoResponse []struct {
	Sku struct {
		Attributes struct {
			LowPriceGuaranteedProduct bool `json:"lowPriceGuaranteedProduct"`
		} `json:"attributes"`
		Brand struct {
			Brand string `json:"brand"`
		} `json:"brand"`
		ButtonState struct {
			Purchasable bool   `json:"purchasable"`
			ButtonState string `json:"buttonState"`
			DisplayText string `json:"displayText"`
			SkuID       string `json:"skuId"`
		} `json:"buttonState"`
		Condition string `json:"condition"`
		Names     struct {
			Short string `json:"short"`
		} `json:"names"`
		Price struct {
			CurrentPrice       float64 `json:"currentPrice"`
			PricingType        string  `json:"pricingType"`
			SmartPricerEnabled bool    `json:"smartPricerEnabled"`
			PriceDomain        struct {
				SkuID               string  `json:"skuId"`
				PriceEventType      string  `json:"priceEventType"`
				RegularPrice        float64 `json:"regularPrice"`
				CurrentPrice        float64 `json:"currentPrice"`
				CustomerPrice       float64 `json:"customerPrice"`
				TotalSavings        float64 `json:"totalSavings"`
				TotalSavingsPercent float64 `json:"totalSavingsPercent"`
				IsMAP               bool    `json:"isMAP"`
				OfferQualifiers     []struct {
					OfferID      string `json:"offerId"`
					OfferName    string `json:"offerName"`
					OfferVersion int    `json:"offerVersion"`
					ID           int    `json:"id"`
				} `json:"offerQualifiers"`
				GiftSkus []struct {
					SkuID    string `json:"skuId"`
					Quantity int    `json:"quantity"`
					OfferID  string `json:"offerId"`
				} `json:"giftSkus"`
				CurrentAsOfDate string `json:"currentAsOfDate"`
			} `json:"priceDomain"`
		} `json:"price"`
		ProductType string `json:"productType"`
		Properties  struct {
			ServicePlansType      string `json:"servicePlansType"`
			ServicePlanDetailsURL string `json:"servicePlanDetailsURL"`
		} `json:"properties"`
		GspAppleCare []struct {
			SkuID string `json:"skuId"`
			Name  string `json:"name"`
			Price struct {
				CurrentPrice float64 `json:"currentPrice"`
			} `json:"price"`
			Type           string `json:"type"`
			ProtectionType string `json:"protectionType"`
			Term           string `json:"term"`
			PaymentType    string `json:"paymentType"`
		} `json:"gspAppleCare"`
		SkuID string `json:"skuId"`
		URL   string `json:"url"`
		Class struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
		} `json:"class"`
		Department struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
		} `json:"department"`
		Subclass struct {
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
		} `json:"subclass"`
		InkSubscriptions []interface{} `json:"inkSubscriptions"`
	} `json:"sku"`
}
