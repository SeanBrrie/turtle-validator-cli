package types

import "github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"

type ValidationPayload struct {
	ContentToValidate string
	ContentSyntax     enums.ContextSyntax
	ValidationType    enums.ValidationType
}
