package types

import "github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"

type ValidationPayload struct {
	ContentToValidate string               `json:"contentToValidate"`
	ContentSyntax     enums.ContextSyntax  `json:"contentSyntax"`
	ValidationType    enums.ValidationType `json:"validationType"`
}
