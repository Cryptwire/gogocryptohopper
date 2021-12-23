// Package gogocryptohopper contains the client implementation
package gogocryptohopper

import (
	"encoding/json"
	"errors"

	"github.com/cryptwire/gogocryptohopper/internal"
	"github.com/cryptwire/gogocryptohopper/internal/hopper"
	"github.com/cryptwire/gogocryptohopper/internal/marketplace"
	"github.com/cryptwire/gogocryptohopper/pkg"
)

// GoGoCryptohopper implements the JSON API for Cryptohopper
type GoGoCryptohopper struct {
	// Hopper implements hopper related APIs
	Hopper *hopper.Hopper
	// Marketplace implements marketplace related APIs
	Marketplace *marketplace.Marketplace
	// client is the HTTP client used to communicate with the API
	client *internal.Client
}

// New creates and returns new instance of the Cryptohopper API client
func New(
	authenticationEndpoint string,
	apiEndpoint string,
	accessToken string,
) (*GoGoCryptohopper, error) {

	headers := map[string]string{}
	client, err := internal.NewClient(authenticationEndpoint, headers)
	if err != nil {
		return nil, err
	}

	hopperClient, err := hopper.New(apiEndpoint, accessToken)
	if err != nil {
		return nil, err
	}

	marketplaceClient, err := marketplace.New(apiEndpoint, accessToken)
	if err != nil {
		return nil, err
	}

	return &GoGoCryptohopper{
		Hopper:      hopperClient,
		Marketplace: marketplaceClient,
		client:      client,
	}, nil
}

// ExchangeGrantToken exchanges a given grant code to an access token
// Cryptohopper documentation: https://www.cryptohopper.com/api-documentation/authentication
func (ch *GoGoCryptohopper) ExchangeGrantToken(
	request pkg.TokenExchangeRequest,
) (pkg.TokenExchangeResponse, error) {

	var tokenResponse pkg.TokenExchangeResponse
	response, err := ch.client.Post("oauth2/token", request)
	if err != nil {
		return tokenResponse, err
	}
	if response == nil {
		return tokenResponse, errors.New("Empty body received, expected TokenExchangeResponse")
	}

	err = json.Unmarshal(response, &tokenResponse)
	if err != nil {
		return tokenResponse, err
	}
	return tokenResponse, nil
}
