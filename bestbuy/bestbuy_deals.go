package bestbuy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// Open Box identifers:
// excellent (certified) - 3 (?)
// excellent - 2
// satisfactory - 1
// fair - 0

type OpenBoxItemInfo struct {
	Avail   bool
	Price   float64
	Savings float64
}

type OpenBoxItem struct {
	ExcellentCertified OpenBoxItemInfo
	Excellent          OpenBoxItemInfo
	Satisfactory       OpenBoxItemInfo
	Fair               OpenBoxItemInfo
}

type Deal struct {
	Sku      string
	Name     string
	Price    float64
	Brand    string
	Link     string
	Discount int
	CartLink string
}

type FinanceOption struct {
	OfferId        string
	Name           string
	TotalCost      float64
	MonthlyPayment float64
}

func GetFinanceOptions(skuId string) ([]FinanceOption, error) {
	url := "https://www.bestbuy.com/pricing/v1/price/item?allFinanceOffers=true&catalog=bby&context=price-block&includeOpenboxPrice=true&includeTotalTechWarranty=false&lib-price=19.2218.14&paidMemberSkuInCart=false&salesChannel=LargeView&skuId=" + skuId + "&usePriceWithCart=true"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	financeResp := FinanceResponse{}
	json.Unmarshal(body, &financeResp)

	var financeOptions []FinanceOption

	for _, financeOption := range financeResp.FinanceOptions {
		var financeOpt FinanceOption

		financeOpt.OfferId = financeOption.OfferID
		financeOpt.Name = financeOption.FinanceCodeName
		financeOpt.TotalCost = financeOption.TotalCost
		financeOpt.MonthlyPayment = financeOption.MonthlyPayment

		financeOptions = append(financeOptions, financeOpt)
	}
	return financeOptions, nil
}

func GetInventoryByZipcode(skuId string, zipCode string) (bool, error) {
	url := "https://www.bestbuy.com/button-state/api/v4/button-state?skus=" + skuId + "&context=pdp&source=buttonView&storeId=&destinationZipCode=" + zipCode
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return false, err
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
	req.Header.Add("x-client-id", "FRV")

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	if res.StatusCode == 200 {
		inventoryResp := ZipcodeInventoryResponse{}
		json.Unmarshal(body, &inventoryResp)

		if inventoryResp.ButtonStateResponseInfos[0].ButtonState == "ADD_TO_CART" {
			return true, nil
		}
	}

	return false, nil
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

	return OpenBoxItemInfo{Avail: displayStatus, Price: openBoxResp.OpenBoxPrice, Savings: openBoxResp.TotalSavings}, nil
}

// gets the latest open box offers for an item id
func FetchOpenBoxOffers(skuId string) (OpenBoxItem, error) {
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

func sendToDiscord(deal Deal) error {
	return nil
}

func sendToTwitter(deal Deal) error {
	return nil
}

func sendToTelegram(deal Deal) error {
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

func GetDealInfo(client *http.Client, skuId string) (Deal, error) {
	url := "https://www.bestbuy.com/api/3.0/priceBlocks?skus=" + skuId
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return Deal{}, err
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

	resp, err := client.Do(req)
	if err != nil {
		return Deal{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Deal{}, err
	}

	dealInfoResp := DealInfoResponse{}

	json.Unmarshal(body, &dealInfoResp)

	return Deal{Sku: dealInfoResp[0].Sku.SkuID, Name: dealInfoResp[0].Sku.Names.Short, Price: dealInfoResp[0].Sku.Price.CurrentPrice, Brand: dealInfoResp[0].Sku.Brand.Brand, Link: "https://www.bestbuy.com" + dealInfoResp[0].Sku.URL, Discount: 0, CartLink: "https://api.bestbuy.com/click/-/" + skuId + "/cart"}, nil
}

func GetItemDiscount(client *http.Client, skuId string) (int, error) {
	url := "https://www.bestbuy.com/priceview/query/sku?includes=price,holidayPriceDealMessages&context=shop-holiday&layout=small&viewType=price&skuId=" + skuId
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, err
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

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	itemDiscountResp := ItemDiscountResponse{}

	json.Unmarshal(body, &itemDiscountResp)
	return itemDiscountResp.TotalSavingsPercent, nil
}

func skuExists(skuList []string, sku string) bool {
	for _, v := range skuList {
		if v == sku {
			return true
		}
	}
	return false
}

func getQueryRecommendations(client *http.Client, searchQuery string) ([]string, error) {
	searchQuery = url.QueryEscape(searchQuery)

	url := "https://www.bestbuy.com/suggest/v1/fragment/suggest/www?query=" + searchQuery
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
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

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	queryRecResp := QueryRecResponse{}

	json.Unmarshal(body, &queryRecResp)

	var skuList []string

	for _, suggestions := range queryRecResp.SuggestionResponse.Suggestions {
		for _, products := range suggestions.Products {
			if skuExists(skuList, products.SkuID) == false {
				skuList = append(skuList, products.SkuID)
			}
		}
	}

	return skuList, nil
}

// returns the latest deals under a particular search query
func FetchLatestDeals(searchQuery string, minimumDiscountPercent int) ([]Deal, error) {
	client := &http.Client{}
	var wg sync.WaitGroup

	skus, err := getQueryRecommendations(client, searchQuery)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var dealList []Deal

	for _, sku := range skus {

		go func(skuId string) {
			wg.Add(1)

			var itemDiscount int
			var deal Deal

			itemDiscount, err = GetItemDiscount(client, skuId)
			if err != nil {
				log.Println(err)
				return
			}

			if itemDiscount >= minimumDiscountPercent {
				// the request made in GetDealInfo also contains the total savings percent
				// but GetItemDiscount may account for further (hidden?) discounts during holidays
				deal, err = GetDealInfo(client, skuId)
				if err != nil {
					log.Println(err)
					return
				}

				deal.Discount = itemDiscount

				dealList = append(dealList, deal)
			}
			wg.Done()
		}(sku)
	}

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return dealList, nil
}
