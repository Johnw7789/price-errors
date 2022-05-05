package bestbuy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Open Box identifers:
// excellent (certified) - 3 (?)
// excellent - 2
// satisfactory - 1
// fair - 0

type OpenBoxItemInfo struct {
	Avail   bool
	Price   float64
	Savings int64
}

type OpenBoxItem struct {
	ExcellentCertified OpenBoxItemInfo
	Excellent          OpenBoxItemInfo
	Satisfactory       OpenBoxItemInfo
	Fair               OpenBoxItemInfo
}

type DealInfo struct {
	Name     string
	Price    float64
	Brand    string
	Link     string
	CartLink string
}

func getOpenBoxInfo(skuId string, openBoxId string) (OpenBoxItemInfo, error) {
	url := "https://www.bestbuy.com/pricing/v1/price/item?catalog=bby&context=PDP-Buying-Options-Price-Block&includeOpenboxPrice=true&openBoxCondition=" + openBoxId + "&salesChannel=LargeView&skuId=" + skuId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return OpenBoxItemInfo{}, err
	}
	req.Header.Add("authority", "www.bestbuy.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("referer", "https://www.bestbuy.com")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Microsoft Edge\";v=\"101\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32")
	req.Header.Add("x-client-id", "lib-price-browser")

	res, err := client.Do(req)
	if err != nil {
		return OpenBoxItemInfo{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return OpenBoxItemInfo{}, err
	}

	openBoxResp := OpenBoxResponse{}

	json.Unmarshal(body, &openBoxResp)

	var displayStatus = false
	if openBoxResp.DotComDisplayStatus == "active" {
		displayStatus = true
	}

	return OpenBoxItemInfo{Avail: displayStatus, Price: openBoxResp.CurrentPrice, Savings: openBoxResp.TotalSavings}, nil
}

// gets the latest open box offers for an item id
func fetchOpenBoxOffers(skuId string) (OpenBoxItem, error) {
	var openBoxItem OpenBoxItem
	var err error

	openBoxItem.ExcellentCertified, err = getOpenBoxInfo(skuId, "3")
	if err != nil {
		return OpenBoxItem{}, err
	}

	openBoxItem.Excellent, err = getOpenBoxInfo(skuId, "2")
	if err != nil {
		return OpenBoxItem{}, err
	}

	openBoxItem.Satisfactory, err = getOpenBoxInfo(skuId, "1")
	if err != nil {
		return OpenBoxItem{}, err
	}

	openBoxItem.Fair, err = getOpenBoxInfo(skuId, "0")
	if err != nil {
		return OpenBoxItem{}, err
	}

	return openBoxItem, nil
}

func sendToDiscord(deal DealInfo) error {
	return nil
}

func sendToTwitter(deal DealInfo) error {
	return nil
}

func sendToTelegram(deal DealInfo) error {
	return nil
}

// sends deal notifications (twitter, discord, telegram, slickdeals)
func AlertDeal(alertType string) {
	switch alertType {
	case "discord":
		go func() {

		}()

	case "twitter":
		go func() {

		}()

	case "telegram":
		go func() {

		}()
	}

}

// evaluates a deal to see if AlertDeal() should be called
func EvaluateDeal() {

}

// returns the latest deal offers
func GetLatestDealOffers() {

}

// returns the latest deals
func FetchLatestDeals(searchQuery string, minimumDiscountPercent int) {

}
