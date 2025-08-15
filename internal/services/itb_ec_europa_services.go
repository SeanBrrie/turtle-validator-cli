package services

import (
	"errors"
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
	if domain == "" {
		return false, errors.New("domain is required")
	}

	if content == "" {
		return false, errors.New("content is required")
	}

	payload := types.ValidationPayload{
		ContentToValidate: content,
		ContentSyntax:     cs,
		ValidationType:    vt,
	}

	valid, err := s.ItbEuropaClient.ShaclValidator(domain, payload)
	if err != nil {
		return valid, err
	}

	return valid, nil
}
