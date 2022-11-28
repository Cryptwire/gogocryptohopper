package hopper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cryptwire/gogocryptohopper/internal"
	"github.com/cryptwire/gogocryptohopper/pkg"
)

// Hopper implements hopper related API calls
type Hopper struct {
	// AccessToken is the authenticated API access token for a user
	accessToken string
	// client is the HTTP client used to communicate with the API
	client *internal.Client
}

// New creates a new Hopper instance
func New(apiEndpoint string, accessToken string) (*Hopper, error) {
	headers := map[string]string{
		"access-token": accessToken,
		"deviceId":     "0",
	}
	client, err := internal.NewClient(apiEndpoint, headers)
	if err != nil {
		return nil, err
	}
	hopper := Hopper{
		accessToken: accessToken,
		client:      client,
	}
	return &hopper, nil
}

// GetAll returns all the hoppers for the user
func (h *Hopper) GetAll() ([]pkg.Hopper, error) {

	response, err := h.client.Get("hopper")
	if err != nil {
		return []pkg.Hopper{}, err
	}
	if response == nil {
		return []pkg.Hopper{}, errors.New("Empty body received, expected GetAllHoppersResponse")
	}

	var apiResponse pkg.GetAllHoppersResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return []pkg.Hopper{}, err
	}
	return apiResponse.Data.Hoppers, nil
}

// Get returns the hopper with given ID
func (h *Hopper) Get(id string) (pkg.Hopper, error) {
	response, err := h.client.Get(fmt.Sprintf("hopper/%s", id))
	if err != nil {
		return pkg.Hopper{}, err
	}
	if response == nil {
		return pkg.Hopper{}, errors.New("Empty body received, expected GetAllHoppersResponse")
	}

	var apiResponse pkg.GetHopperResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return pkg.Hopper{}, err
	}
	return apiResponse.Data.Hopper, nil
}

// GetBaseconfig retrieves the current baseconfig for a hopper
func (h *Hopper) GetBaseconfig(id string) (pkg.HopperBaseconfig, error) {
	response, err := h.client.Get(fmt.Sprintf("hopper/%s/config", id))
	if err != nil {
		return pkg.HopperBaseconfig{}, err
	}
	if response == nil {
		return pkg.HopperBaseconfig{}, errors.New("Empty body received, expected GetHopperConfigResponse")
	}

	var apiResponse pkg.GetHopperConfigResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return pkg.HopperBaseconfig{}, err
	}
	return apiResponse.Data, nil
}

func (h *Hopper) GetOpenPositions(id string) ([]pkg.Position, error) {
	var positions []pkg.Position
	response, err := h.client.Get(fmt.Sprintf("hopper/%s/position", id))
	if err != nil {
		return positions, err
	}
	if response == nil {
		return positions, errors.New("Empty body received, expected GetHopperOpenPositionsResponse")
	}

	var apiResponse pkg.GetHopperOpenPositionsResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return positions, err
	}
	return apiResponse.Data, nil
}

func (h *Hopper) GetReservedPositions(id string) ([]pkg.Position, error) {
	var positions []pkg.Position
	response, err := h.client.Get(fmt.Sprintf("hopper/%s/reserved", id))
	if err != nil {
		return positions, err
	}
	if response == nil {
		return positions, errors.New("Empty body received, expected GetReservedPositions response")
	}

	var apiResponse pkg.GetHopperReservedPositionsResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return positions, err
	}
	return apiResponse.Data, nil
}

func (h *Hopper) GetShortPositions(id string) ([]pkg.Position, error) {
	var positions []pkg.Position
	response, err := h.client.Get(fmt.Sprintf("hopper/%s/short", id))
	if err != nil {
		return positions, err
	}
	if response == nil {
		return positions, errors.New("Empty body received, expected GetShortPositions response")
	}

	var apiResponse pkg.GetHopperShortPositionsResponse

	// Cryptohopper, for some reason, returns the string "stop_loss 0.000000"
	// as the first line of the response instead of the JSON object
	// So we cut out everything before the first curly brace {
	response = response[strings.Index(string(response), "{"):]

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return positions, err
	}
	return apiResponse.Data, nil
}

// GetTradeHistory returns the most recent trades for the given hopper
func (h *Hopper) GetTradeHistory(id string) ([]pkg.Trade, error) {
	var trades []pkg.Trade
	response, err := h.client.Get(fmt.Sprintf("hopper/%s/trade", id))
	if err != nil {
		return trades, err
	}
	if response == nil {
		return trades, errors.New("Empty body received, expected GetHopperTradesResponse")
	}

	var apiResponse pkg.GetHopperTradesResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return trades, err
	}
	return apiResponse.Data.Trades, nil
}

// GetTradeHistoryWithTimeframes returns all the trades within a specific
// timeframe
func (h *Hopper) GetTradeHistoryWithTimeframes(id string, dateFrom time.Time, dateTo time.Time, offset int) (pkg.TradeList, error) {
	var data pkg.TradeList

	dateFromPlain := dateFrom.Format("2006-01-02")
	dateToPlain := dateTo.Format("2006-01-02")

	_ = dateToPlain

	apiUrl := fmt.Sprintf("hopper/%s/trade?dateFrom=%s&offset=%d", id, dateFromPlain, offset)
	response, err := h.client.Get(apiUrl)
	if err != nil {
		return data, err
	}
	if response == nil {
		return data, errors.New("Empty body received, expected GetHopperTradesResponse")
	}

	var apiResponse pkg.GetHopperTradesResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return data, err
	}
	return apiResponse.Data, nil
}

// GetStats returns the stats for a hopper
func (h *Hopper) GetStats(id string) (pkg.HopperStats, error) {
	var data pkg.HopperStats

	apiUrl := fmt.Sprintf("hopper/%s/stats", id)
	response, err := h.client.Get(apiUrl)
	if err != nil {
		return data, err
	}
	if response == nil {
		return data, errors.New("Empty body received, expected GetHopperStatsResponse")
	}

	var apiResponse pkg.GetHopperStatsResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return data, err
	}
	return apiResponse.Data, nil

}

// GetAssets returns the assets for a hopper
func (h *Hopper) GetAssets(id string) (map[string]string, error) {
	var data map[string]string

	apiUrl := fmt.Sprintf("hopper/%s/assets", id)
	response, err := h.client.Get(apiUrl)
	if err != nil {
		return data, err
	}
	if response == nil {
		return data, errors.New("Empty body received, expected GetHopperAssetsResponse")
	}

	var apiResponse pkg.GetHopperAssetsResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return data, err
	}
	return apiResponse.Data, nil

}

// GetLogs returns the logs for a hopper
func (h *Hopper) GetLogs(id string) ([]pkg.HopperLog, error) {
	var data []pkg.HopperLog

	apiUrl := fmt.Sprintf("hopper/%s/output", id)
	response, err := h.client.Get(apiUrl)
	if err != nil {
		return data, err
	}
	if response == nil {
		return data, errors.New("Empty body received, expected GetHopperLogsResponse")
	}

	var apiResponse pkg.GetHopperLogsResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return data, err
	}
	return apiResponse.Data, nil

}

// PanicSell initiates a panic sell on the provided Hopper
// where ID is the Cryptohopper Hopper ID
// This call requires manage and trade permissions
func (h *Hopper) PanicSell(id string) error {

	apiUrl := fmt.Sprintf("hopper/%s/panicbutton", id)
	response, err := h.client.Get(apiUrl)
	if err != nil {
		return err
	}
	if response == nil {
		return errors.New("Empty body received, expected GetHopperPanicSellResponse")
	}

	var apiResponse pkg.GetHopperPanicSellResponse

	err = json.Unmarshal(response, &apiResponse)
	if err != nil {
		return err
	}

	if strings.Contains(strings.ToLower(apiResponse.Data), "panic mode") {
		return nil
	}

	return errors.New("Panic mode refused")

}
