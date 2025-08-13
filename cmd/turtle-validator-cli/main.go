package turtle_validator_cli

import (
	"github.com/SeanBrrie/turtle-validator-cli/internal/clients"
	"github.com/SeanBrrie/turtle-validator-cli/internal/services"
	"log"
)

func main() {
	client := clients.NewItbEuropaClient()

	itbEuropaServices, err := services.NewItbEuropaServices(client)
	if err != nil {
		log.Fatalf("failed to create ITB Europa service: %v", err)
	}

}
