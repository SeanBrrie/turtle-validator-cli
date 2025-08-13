package services

import (
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/types"
)

type ItbEuropaServices struct {
	ItbEuropaClient *clients.ItbEuropaClient
}

func NewItbEuropaServices(c *clients.ItbEuropaClient) (ItbEuropaServices, error) {
	return ItbEuropaServices{
		ItbEuropaClient: c,
	}, nil
}

// ValidateContent validates Turtle content using the ITB Europa SHACL API
func (s *ItbEuropaServices) ValidateContent(domain string, content string, cs enums.ContextSyntax, vt enums.ValidationType) (bool, error) {
	payload := types.ValidationPayload{
		ContentToValidate: content,
		ContentSyntax:     cs,
		ValidationType:    vt,
	}

	_, err := s.ItbEuropaClient.ShaclValidator(domain, payload)
	if err != nil {
		return false, err
	}

	return true, nil
}
