package main

import (
	"fmt"

	"github.com/cryptwire/gogocryptohopper"
)

func main() {
	fmt.Println("Example: Fetch marketplace items")

	ch, err := gogocryptohopper.New(
		"https://www.cryptohopper.com",
		"https://api.cryptohopper.com/v1",
		accessToken)
	if err != nil {
		panic(err)
	}

	// NOTE: The strategy ID returned by the baseconfig call doesn't match
	// any strategies in the marketplace.
	// Will contact cryptohopper to clarify
	// fetchStrategy("stratid", ch)
	fetchUserMarketStrategy("strat", ch)

}

func fetchStrategy(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Marketplace.GetStrategy(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func fetchUserMarketStrategy(id string, ch *gogocryptohopper.GoGoCryptohopper) {
	result, err := ch.Marketplace.GetUserMarketStrategy(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Name)
	// fmt.Println(result)
}
