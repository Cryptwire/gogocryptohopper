package pkg

// AuthError defines the error format for API responses used in authentication
type AuthError struct {
	// Error contains the error code or string
	Error string `json:"error"`
	// ErrorDescription contains the description of the error
	ErrorDescription string `json:"error_description"`
}
