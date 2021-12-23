package main

import (
	"fmt"

	"github.com/cryptwire/gogocryptohopper"
	"github.com/cryptwire/gogocryptohopper/pkg"
)

func main() {
	fmt.Println("Example: Authentication")

	ch, err := gogocryptohopper.New(
		"https://www.cryptohopper.com",
		"https://api.cryptohopper.com/v1",
		"accessToken")
	if err != nil {
		panic(err)
	}

	tokenExchangeRequest := pkg.TokenExchangeRequest{
		ClientID:     "your client ID",
		ClientSecret: "your client secret",
		GrantType:    "authorization_code", // This must be authorization_code
		RedirectURI:  "https://www.your.com/redirect/url",
		GrantCode:    "your grant code",
	}

	token, err := ch.ExchangeGrantToken(tokenExchangeRequest)
	if err != nil {
		panic(err)
	}

	fmt.Println("AccessToken: ", token.AccessToken)
	fmt.Println("ExpiresIn: ", token.ExpiresIn)
	fmt.Println("TokenType: ", token.TokenType)
	fmt.Println("Scope: ", token.Scope)
	fmt.Println("RefreshToken: ", token.RefreshToken)

}
