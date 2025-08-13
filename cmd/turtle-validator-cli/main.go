package main

import (
	"fmt"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"
	"github.com/SeanBrrie/turtle-validator-cli/internal/services"
)

// todo:
// 1. Validate through api request
// 2. Create functions in cli
// 3.

const test = `@prefix dcat: <http://www.w3.org/ns/dcat#> .
@prefix dct: <http://purl.org/dc/terms/> .
@prefix ex: <http://example.com/ns#> .

ex:ValidExample a dcat:Dataset ;
dct:title "Example Valid Dataset" ;
dct:description "This is an example of a dataset that should pass the ITB validation." .
`

func main() {
	client := clients.NewItbEuropaClient()

	itbEuropaServices, err := services.NewItbEuropaServices(client)
	if err != nil {
		fmt.Print(err)
	}

	valid, err := itbEuropaServices.ValidateContent("dcat-ap", test, enums.Turtle, enums.V3Full1)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("isvalid: ", valid)
}
