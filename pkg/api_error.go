package pkg

// APIError defines the error format for API responses
type APIError struct {
	// Error contains the error code
	Error int `json:"error"`
	// Message contains the description of the error
	Message string `json:"message"`
}

type GenericError struct {
	// Status contains the status code
	Status int `json:"status"`
	// Message contains the description of the error
	Message string `json:"message"`
}
