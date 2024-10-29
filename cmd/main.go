package main

import (
	"generic-api-sync-service/internal/config"
	"generic-api-sync-service/internal/scheduler"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Start the scheduler
	scheduler.StartSync(cfg)
}
