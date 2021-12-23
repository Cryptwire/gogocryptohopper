package pkg

// TokenExchangeRequest defines the structure for grant token exchange requests
type TokenExchangeRequest struct {
	// ClientID is the application's client ID
	ClientID string `json:"client_id"`
	// ClientSecret is the application's secret
	ClientSecret string `json:"client_secret"`
	// GrantType is the type of grant being exchanged
	GrantType string `json:"grant_type"`
	// RedirectURI is the redirect URL of the application
	RedirectURI string `json:"redirect_uri"`
	// GrantCode is the grant code returned by authorize
	GrantCode string `json:"code"`
}

// TokenExchangeResponse defines the structure of grant token exchange responses
type TokenExchangeResponse struct {
	// AccessToken to be used in authenticated API calls
	AccessToken string `json:"access_token"`
	// ExpiresIn contains the time until the access code expires, in seconds
	ExpiresIn uint `json:"expires_in"`
	// TokenType determines the type of token received
	TokenType string `json:"token_type"`
	// Scope contains the granted permissions
	Scope string `json:"scope"`
	// RefreshToken contains the token to refresh the access token with
	RefreshToken string `json:"refresh_token"`
}
