package clients

import (
	"net/http"
)

type ItbEuropaClient struct {
	client *http.Client
}

func NewItbEuropaClient() *ItbEuropaClient {
	client := &http.Client{}

	return &ItbEuropaClient{
		client: client,
	}
}
