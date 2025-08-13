package enums

type ContextSyntax string

const (
	XML    ContextSyntax = "application/rdf+xml"
	JSONLD ContextSyntax = "application/ld+json"
	Turtle ContextSyntax = "text/turtle"
)
