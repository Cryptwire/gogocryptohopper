package pkg

// Strategy defines the structure of a marketplace strategy
type Strategy struct {
	Type                string `json:"type"`
	ID                  string `json:"id"`
	MarketID            string `json:"market_id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Image               string `json:"image"`
	HiddenConfiguration string `json:"hidden_configuration"`
	CreatedDate         string `json:"created_date"`
	Link                string `json:"link"`
}

// GetStrategyResponse holds an array of HopperResponses
type GetStrategyResponse struct {
	// Data contains the response information
	Data Strategy `json:"data"`
}

// GetUserMarketStrategyResponse holds the information for a user
// market strategy
type GetUserMarketStrategyResponse struct {
	Data Strategy `json:"data"`
}
