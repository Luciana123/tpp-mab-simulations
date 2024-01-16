package jackpot

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPClient struct {
	BaseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
	}
}

func (c *HTTPClient) Get(path string) (string, error) {
	url := c.BaseURL + path

	// Make a GET request to the service
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %v", err)
	}
	defer response.Body.Close()

	// Check if the request was successful (status code 200)
	if response.StatusCode == http.StatusOK {
		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body: %v", err)
		}
		return string(body), nil
	}

	return "", fmt.Errorf("request failed with status code: %v", response.StatusCode)
}

func (c *HTTPClient) Post(path string, payload []byte) (string, error) {
	url := c.BaseURL + path

	// Make a POST request to the service with the payload
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("error making POST request: %v", err)
	}
	defer response.Body.Close()

	// Check if the request was successful (status code 200)
	if response.StatusCode == http.StatusOK {
		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", fmt.Errorf("error reading response body: %v", err)
		}
		return string(body), nil
	}

	return "", fmt.Errorf("request failed with status code: %v", response.StatusCode)
}
