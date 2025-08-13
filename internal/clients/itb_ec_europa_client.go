package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/types"
	"io"
	"net/http"
)

type ItbEuropaClient struct {
	client  *http.Client
	baseURL string
}

func NewItbEuropaClient() *ItbEuropaClient {
	client := &http.Client{}

	return &ItbEuropaClient{
		client:  client,
		baseURL: "https://www.itb.ec.europa.eu/shacl",
	}
}

// ShaclValidator sends Turtle content to the ITB SHACL API and returns validation results.
func (c *ItbEuropaClient) ShaclValidator(domain string, payload types.ValidationPayload) (*types.ValidationResponse, error) {
	url := fmt.Sprintf("%s/%s/api/validate", c.baseURL, domain)

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create a new POST request with JSON body
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	var validationResp types.ValidationResponse
	if err := json.Unmarshal(body, &validationResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &validationResp, nil
}
