package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mkanetsuna/jooto-go-webapp-template/internal/handlers"
	"github.com/mkanetsuna/jooto-go-webapp-template/pkg/health"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 環境変数の設定
	handlers.SetSpreadsheetID(os.Getenv("SPREADSHEET_ID"))
	handlers.SetServiceAccountKey(os.Getenv("GCP_SA_KEY"))	

	http.HandleFunc("/scheduled-task", handlers.ScheduledTaskHandler)
	http.HandleFunc("/webhook1", handlers.Webhook1Handler)
	http.HandleFunc("/webhook2", handlers.Webhook2Handler)
	http.HandleFunc("/health", health.HealthCheckHandler)
	http.HandleFunc("/", handlers.DevHandler)

	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}