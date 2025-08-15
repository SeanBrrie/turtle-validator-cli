package main

import (
	"fmt"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients/enums"
	"github.com/SeanBrrie/turtle-validator-cli/internal/services"
	"log"
	"os"
)

func main() {

	client := clients.NewItbEuropaClient()

	itbEuropaServices, err := services.NewItbEuropaServices(client)
	if err != nil {
		log.Fatalf("Failed to create ITB Europa service: %v", err)
	}

	for {
		fmt.Println("\n--- New Validation Request ---")
		fmt.Println("Type 'exit' to quit.")

		domain, content, contextSyntax, validationType := getUserInput()

		valid, err := itbEuropaServices.ValidateContent(domain, content, contextSyntax, validationType)
		if err != nil {
			log.Printf("Validation error: %v", err)
		}

		fmt.Println("Validation result:", valid)
	}
}

func getUserInput() (string, string, enums.ContextSyntax, enums.ValidationType) {
	var domain string
	fmt.Print("Domain: (e.g., dcat-ap, healthri): ")
	fmt.Scan(&domain)

	var contentFileName string
	fmt.Print("Content file name: ")
	fmt.Scan(&contentFileName)

	content, err := getFileContent(contentFileName)
	if err != nil {
		log.Fatalf("Failed to read content file: %v", err)
	}
	var contextSyntaxStr string
	fmt.Print("Context syntax (e.g., XML, JSONLD, Turtle): ")
	fmt.Scan(&contextSyntaxStr)
	contextSyntax, err := enums.GetContextSyntax(contextSyntaxStr)
	if err != nil {
		log.Fatalf("Failed to parse context syntax: %v", err)
	}

	var validationTypeStr string
	fmt.Print("Validation type (e.g., V3Full1, V200): ")
	fmt.Scan(&validationTypeStr)
	validationType, err := enums.GetValidationType(validationTypeStr)
	if err != nil {
		log.Fatalf("Failed to parse context syntax: %v", err)
	}

	return domain, content, contextSyntax, validationType
}

func getFileContent(fileName string) (string, error) {
	data, err := os.ReadFile("data/" + fileName)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("file not found: %s", fileName)
	}

	return string(data), nil
}
