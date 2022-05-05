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
