package types

type ValidationResponse struct {
	Valid   bool     `xml:"RDF"`
	Reports []string `xml:"ValidationReport"`
}
