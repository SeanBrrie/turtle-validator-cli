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

	var contentFilePath string
	fmt.Print("File path to .ttl file: ")
	fmt.Scan(&contentFilePath)

	content, err := getFileContent(contentFilePath)
	if err != nil {
		log.Fatalf("Failed to read content file: %v", err)
	}

	var contextSyntax enums.ContextSyntax
	fmt.Print("Context syntax (e.g., XML, JSONLD, Turtle): ")
	fmt.Scan(&contextSyntax)

	var validationType enums.ValidationType
	fmt.Print("Validation type (e.g., V3Full1, V200): ")
	fmt.Scan(&validationType)

	return domain, content, contextSyntax, validationType
}

func getFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("empty file content: %s", filePath)
	}

	return string(data), nil
}
