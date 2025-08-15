package enums

import (
	"fmt"
	"strings"
)

type ValidationType string

const (
	V3Full1 ValidationType = "v3.Full1"
	V200    ValidationType = "v2.0.0"
)

// Parse ValidationType from string input
func GetValidationType(input string) (ValidationType, error) {
	switch strings.ToLower(input) {
	case "v3full1":
		return V3Full1, nil
	case "v200":
		return V200, nil
	default:
		return "", fmt.Errorf("invalid validation type: %s", input)
	}
}
