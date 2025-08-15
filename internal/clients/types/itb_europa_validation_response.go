package types

import "encoding/xml"

type ValidationResponse struct {
	XMLName          xml.Name         `xml:"RDF"`
	ValidationReport ValidationReport `xml:"ValidationReport"`
}

type ValidationReport struct {
	Conforms bool `xml:"conforms"`
}
