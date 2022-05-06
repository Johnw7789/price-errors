# Price Errors
## About This Project
This project is designed with the goal of providing a cutting edge toolkit for developers and deal enthusiasts alike in order to assist in automating the deal-finding process as well as catching the increasingly rare "price errors" that we all seem to miss out on. Most price errors are often caught, replicated, and shared within artificially exclusive, paid deal groups. The vast majority of these groups are a waste of time and my hope is that this project can provide a sufficient alternative to the few groups that do seem to garner regular success.

The two main sites that I aim to provide consistent updates for are Amazon and BestBuy.

## Usage
##### BestBuy
This toolkit utilizes a few key BestBuy endpoints and exposes them to an easy to use api.

A few of these are demonstrated below.



```Getting deals using a search query```

Note: the more simplistic and vague the search query, the more items will be evaluated. This is because BestBuy's search recommondation api favors less complex queries.
```
deals, err := bestbuy.FetchLatestDeals("Zephyrus laptop", 10)
if err != nil {
	panic(err)
}

for _, deal := range deals {
	fmt.Println(fmt.Sprintf("%s - %d percent off - $%f", deal.Name, deal.Discount, deal.Price))
	
	// Name     string
	// Price    float64
	// Brand    string
	// Link     string
	// Discount int
	// CartLink string
}
```

```Getting Open-Box offers for a specific sku```

Note: if the sku is not available under a particular open-box condition, then the price and savings will be 0, so always check to see if it is available.
```
openBoxOffers, err := bestbuy.FetchOpenBoxOffers("123456")
if err != nil {
	panic(err)
}

// There are 4 possible open-box conditions:
// - ExcellentCertified
// - Excellent
// - Satisfactory
// - Fair

log.Println(openBoxOffers.Excellent.Avail)   // bool
log.Println(openBoxOffers.Excellent.Price)   // float64
log.Println(openBoxOffers.Excellent.Savings) // float64
```

```Getting finance options for a sku```

```
financeOptions, err := bestbuy.GetFinanceOptions("4901811")
if err != nil {
	panic(err)
}

for _, financeOption := range financeOptions {
	fmt.Println(financeOption.Name)           // string
	fmt.Println(financeOption.OfferId)        // string
	fmt.Println(financeOption.MonthlyPayment) // float64
	fmt.Println(financeOption.TotalCost)      // float64
}
```

```Getting inventory status for a zipcode```

```
isAvailable, err := bestbuy.GetInventoryByZipcode("sku", "zipcode")
if err != nil {
	panic(err)
}

if isAvailable {
	log.Println("In Stock!")
}
```
