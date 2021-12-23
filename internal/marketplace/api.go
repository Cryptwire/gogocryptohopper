package marketplace

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cryptwire/gogocryptohopper/internal"
	"github.com/cryptwire/gogocryptohopper/pkg"
)

// Marketplace implements marketplace related API calls
type Marketplace struct {
	// AccessToken is the authenticated API access token for a user
	accessToken string
	// client is the HTTP client used to communicate with the API
	client *internal.Client
}

// New creates a new Hopper instance
func New(apiEndpoint string, accessToken string) (*Marketplace, error) {
	headers := map[string]string{
		"access-token": accessToken,
	}
	client, err := internal.NewClient(apiEndpoint, headers)
	if err != nil {
		return nil, err
	}
	marketplace := Marketplace{
		accessToken: accessToken,
		client:      client,
	}
	return &marketplace, nil
}

// GetStrategy returns the strategy information for a given ID
func (h *Marketplace) GetStrategy(id string) (pkg.Strategy, error) {

	response, err := h.client.Get(fmt.Sprintf("market/strategies/%s", id))
	if err != nil {
		return pkg.Strategy{}, err
	}
	if response == nil {
		return pkg.Strategy{}, errors.New("Empty body received, expected GetAllHoppersResponse")
	}

	var apiResponse pkg.GetStrategyResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return pkg.Strategy{}, err
	}
	return apiResponse.Data, nil
}

// GetUserMarketStrategy returns the strategy information for a given
// user strategy ID from teh Markerplace
func (h *Marketplace) GetUserMarketStrategy(id string) (pkg.Strategy, error) {

	response, err := h.client.Get(fmt.Sprintf("strategy/market/%s", id))
	if err != nil {
		return pkg.Strategy{}, err
	}
	if response == nil {
		return pkg.Strategy{}, errors.New("Empty body received, expected GetAllHoppersResponse")
	}

	var apiResponse pkg.GetUserMarketStrategyResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return pkg.Strategy{}, err
	}
	return apiResponse.Data, nil
}
