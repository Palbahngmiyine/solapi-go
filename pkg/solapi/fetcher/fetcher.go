package fetcher

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/solapi/solapi-go/pkg/solapi/authenticator"
	"github.com/solapi/solapi-go/pkg/solapi/types"
)

const sdkVersion string = "golang/2.0.0"

var (
	errFailedToConvertJSON   = errors.New("FailedToConvertJSON")
	errFailedToClientRequest = errors.New("FailedToClientRequest")
)

// Fetcher api
type Fetcher struct {
	// Config
	APIKey     string
	APISecret  string
	Protocol   string
	Domain     string
	Prefix     string
	SdkVersion string
	OsPlatform string

	// Authenticator
	Auth *authenticator.Authenticator
}

// NewFetcherInstance creates a new Fetcher instance with the provided API key and API secret
// This is used internally by the Request function and can be used when you need access to the Fetcher properties
func NewFetcherInstance(apiKey, apiSecret string) *Fetcher {
	goos := runtime.GOOS
	goVersion := runtime.Version()
	osPlatform := fmt.Sprintf("%s/%s", goos, goVersion)

	fetcher := &Fetcher{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		Protocol:   "https",
		Domain:     "api.solapi.com",
		Prefix:     "",
		SdkVersion: sdkVersion,
		OsPlatform: osPlatform,
	}

	// Initialize authenticator
	fetcher.Auth = authenticator.NewAuthenticator(apiKey, apiSecret)

	return fetcher
}

// Request creates a new Fetcher instance with the provided API key and API secret
// and performs an HTTP request with the specified method, resource, params, and customStruct
func Request(method string, resource string, params interface{}, customStruct interface{}, apiKey, apiSecret string) error {
	fetcher := NewFetcherInstance(apiKey, apiSecret)

	// If method is GET, handle it differently
	if method == "GET" {
		// Convert params to map[string]string if it's not nil
		paramsMap := make(map[string]string)
		if params != nil {
			// Try to convert params to map[string]string
			if paramsAsMap, ok := params.(map[string]string); ok {
				paramsMap = paramsAsMap
			}
		}
		return fetcher.GET(resource, paramsMap, customStruct)
	}

	// For other methods, use the Request method
	return fetcher.Request(method, resource, params, customStruct)
}

// NewFetcher creates a new Fetcher instance with the provided API key and API secret
// This is kept for backward compatibility
func NewFetcher(apiKey, apiSecret string) *Fetcher {
	goos := runtime.GOOS
	goVersion := runtime.Version()
	osPlatform := fmt.Sprintf("%s/%s", goos, goVersion)

	fetcher := &Fetcher{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		Protocol:   "https",
		Domain:     "api.solapi.com",
		Prefix:     "",
		SdkVersion: sdkVersion,
		OsPlatform: osPlatform,
	}

	// Initialize authenticator
	fetcher.Auth = authenticator.NewAuthenticator(apiKey, apiSecret)

	return fetcher
}


// GET method request
func (f *Fetcher) GET(resource string, params map[string]string, customStruct interface{}) error {
	// Prepare for Http Request
	client := &http.Client{}
	url := fmt.Sprintf("%s://%s/%s%s", f.Protocol, f.Domain, f.Prefix, resource)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set Query Parameters
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Set Headers
	authorization := f.Auth.GetAuthorization()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errFailedToClientRequest
	}
	defer resp.Body.Close()

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		err = json.NewDecoder(resp.Body).Decode(&errorStruct)
		if err != nil {
			return err
		}
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	err = json.NewDecoder(resp.Body).Decode(&customStruct)
	if err != nil {
		return err
	}
	return nil
}

// Request method request
func (f *Fetcher) Request(method string, resource string, params interface{}, customStruct interface{}) error {
	// Convert to json string
	jsonString, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return errFailedToConvertJSON
	}

	// Prepare for Http Request
	client := &http.Client{}
	url := fmt.Sprintf("%s://%s/%s%s", f.Protocol, f.Domain, f.Prefix, resource)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonString))
	if err != nil {
		return err
	}

	// Set Headers
	authorization := f.Auth.GetAuthorization()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errFailedToClientRequest
	}
	defer resp.Body.Close()

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		err = json.NewDecoder(resp.Body).Decode(&errorStruct)
		if err != nil {
			return err
		}
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	err = json.NewDecoder(resp.Body).Decode(&customStruct)
	if err != nil {
		return err
	}

	return nil
}

// POST method request
func (f *Fetcher) POST(resource string, params interface{}, customStruct interface{}) error {
	return f.Request("POST", resource, params, customStruct)
}

// PUT method request
func (f *Fetcher) PUT(resource string, params interface{}, customStruct interface{}) error {
	return f.Request("PUT", resource, params, customStruct)
}

// DELETE method request
func (f *Fetcher) DELETE(resource string, params interface{}, customStruct interface{}) error {
	return f.Request("DELETE", resource, params, customStruct)
}
