package main

import (
	"flag"
	"log"

	"github.com/mkanetsuna/jooto-go-webapp-template/internal/scheduler"
)

func main() {
	projectID := flag.String("project", "", "GCP Project ID")
	location := flag.String("location", "", "GCP Region")
	serviceURL := flag.String("service-url", "", "Cloud Run Service URL")
	flag.Parse()

	if *projectID == "" || *location == "" || *serviceURL == "" {
		log.Fatal("All flags must be provided")
	}

	if err := scheduler.SetupScheduler(*projectID, *location, *serviceURL); err != nil {
		log.Fatalf("Failed to setup scheduler: %v", err)
	}

	log.Println("Scheduler setup completed successfully")
}