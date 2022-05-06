# Price Errors
## About This Project
This project is designed with the goal of providing a cutting edge toolkit for developers and deal enthusiasts alike in order to assist in automating the deal-finding process as well as catching the increasingly rare "price errors" that we all seem to miss out on. Most price errors are often caught, replicated, and shared within artificially exclusive, paid deal groups. The vast majority of these groups are a waste of time and my hope is that this project can provide a sufficient alternative to the few groups that do seem to garner regular success.

The two main sites that I aim to provide consistent updates for are Amazon and BestBuy.

## Usage
##### BestBuy
This toolkit utilizes a few key BestBuy endpoints and exposes them to an easy to use api.

A few of these are demonstrated below.
```
deals, err := bestbuy.FetchLatestDeals("Zephyrus laptop", 10)
if err != nil {
	panic(err)
}

for _, deal := range deals {
	fmt.Println(fmt.Sprintf("%s - %d percent off - $%f", deal.Name, deal.Discount, deal.Price))

	openBoxOffers, err := bestbuy.FetchOpenBoxOffers(deal.Sku)
	if err != nil {
		panic(err)
	}

	log.Println(openBoxOffers.Excellent.Price)
}
```
