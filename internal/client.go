package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cryptwire/gogocryptohopper/pkg"
)

// Client implements a generic HTTP client for simple interaction
type Client struct {
	// apiEndpoint holds the API endpoint to interact with
	apiEndpoint string
	// headers contains additional HTTP headers to send
	headers map[string]string
	// httpClient holds the http client
	httpClient http.Client
}

// NewClient creates a new instance of the generic Client
func NewClient(apiEndpoint string, headers map[string]string) (*Client, error) {
	if apiEndpoint == "" {
		return nil, errors.New("apiEndpoint must be provided")
	}
	return &Client{
		apiEndpoint: apiEndpoint,
		headers:     headers,
	}, nil
}

// Post the given API path with the given payload object
// Returns the body of the response and possible error
func (client *Client) Post(path string, payloadObject interface{}) ([]byte, error) {
	payload, err := json.Marshal(payloadObject)
	if err != nil {
		return nil, err
	}

	callEndpoint := fmt.Sprintf("%s/%s", client.apiEndpoint, path)
	req, err := http.NewRequest("POST", callEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("Unable to create API request '%s': %s", path, err)
	}

	// Workaround, Cryptohopper added additional security to the primary domain
	// to avoid bots from accessing it. Here we fake to be Firefox to work
	// around it
	if strings.Contains(path, "oauth") {
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/81.0")
	}

	req.Header.Set("Content-Type", "application/json")
	for header, value := range client.headers {
		req.Header.Set(header, value)
	}

	client.httpClient.Timeout = 1 * time.Minute
	response, err := client.httpClient.Do(req)
	if err != nil && response == nil {
		return nil, fmt.Errorf("Unable to call API '%s': %s", path, err)
	}
	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read API response: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK &&
		response.StatusCode != http.StatusCreated {
		var apiError pkg.APIError
		err = json.Unmarshal(responseContent, &apiError)
		if err != nil {

			// Try parsing it as an auth error
			var authError pkg.AuthError
			err = json.Unmarshal(responseContent, &authError)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("%s: %s", authError.Error, authError.ErrorDescription)
		}

		return nil, fmt.Errorf("%d: %s", apiError.Error, apiError.Message)
	}

	// Attempt to parse this as a ratelimit
	var genericError pkg.GenericError
	err = json.Unmarshal(responseContent, &genericError)
	if err == nil {
		if genericError.Status == 429 {
			return nil, errors.New("Ratelimit")
		}
	}

	return responseContent, nil
}

// Get the given API path
// Returns the body of the response and possible error
func (client *Client) Get(path string) ([]byte, error) {

	callEndpoint := fmt.Sprintf("%s/%s", client.apiEndpoint, path)
	req, err := http.NewRequest("GET", callEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to create API request '%s': %s", path, err)
	}

	req.Header.Set("Content-Type", "application/json")
	for header, value := range client.headers {
		req.Header.Set(header, value)
	}
	client.httpClient.Timeout = 10 * time.Minute
	response, err := client.httpClient.Do(req)
	if err != nil && response == nil {
		return nil, fmt.Errorf("Unable to call API '%s': %s", path, err)
	}
	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read API response: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK &&
		response.StatusCode != http.StatusCreated {
		var apiError pkg.APIError
		err = json.Unmarshal(responseContent, &apiError)
		if err != nil {

			// Try parsing it as an auth error
			var authError pkg.AuthError
			err = json.Unmarshal(responseContent, &authError)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("%s: %s", authError.Error, authError.ErrorDescription)
		}

		return nil, fmt.Errorf("%d: %s", apiError.Error, apiError.Message)
	}

	// Attempt to parse this as a ratelimit
	var genericError pkg.GenericError
	err = json.Unmarshal(responseContent, &genericError)
	if err == nil {
		if genericError.Status == 429 {
			return nil, errors.New("Ratelimit")
		}
	}

	return responseContent, nil
}

// createHTTPClient creates a new HTTP client
func createHTTPClient(requestTimeoutSeconds int) *http.Client {
	client := &http.Client{
		Timeout: time.Duration(requestTimeoutSeconds) * time.Second,
	}
	return client
}
