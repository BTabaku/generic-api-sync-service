package scheduler

import (
	"generic-api-sync-service/internal/config"
	"log"
	"time"
)

// StartSync starts the synchronization process based on the provided configuration
func StartSync(cfg *config.Config) {
	log.Println("Starting synchronization process...")
	// Add your scheduler code here
	ticker := time.NewTicker(time.Duration(cfg.Sync.Interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running sync task...")
		// Add your sync task code here
	}
}
