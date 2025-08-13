package services

import (
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
)

type ItbEuropaServices struct {
	ItbEuropaClient *clients.ItbEuropaClient
}

func NewItbEuropaServices(c *clients.ItbEuropaClient) (ItbEuropaServices, error) {
	return ItbEuropaServices{
		ItbEuropaClient: c,
	}, nil
}
