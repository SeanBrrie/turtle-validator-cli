package enums

import (
	"fmt"
	"strings"
)

type ContextSyntax string

const (
	XML    ContextSyntax = "application/rdf+xml"
	JSONLD ContextSyntax = "application/ld+json"
	Turtle ContextSyntax = "text/turtle"
)

// Parse ContextSyntax from string input
func GetContextSyntax(input string) (ContextSyntax, error) {
	switch strings.ToLower(input) {
	case "xml":
		return XML, nil
	case "jsonld":
		return JSONLD, nil
	case "turtle":
		return Turtle, nil
	default:
		return "", fmt.Errorf("invalid context syntax: %s", input)
	}
}
