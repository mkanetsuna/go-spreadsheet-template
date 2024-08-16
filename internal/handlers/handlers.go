package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	spreadsheetID     string
	serviceAccountKey string
)

func SetSpreadsheetID(id string) {
	spreadsheetID = id
	log.Printf("Spreadsheet ID set: %s", id)
}

func SetServiceAccountKey(key string) {
	serviceAccountKey = key
	log.Printf("Service Account Key set (length: %d)", len(key))
}

func ScheduledTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Scheduled task executed at %s", time.Now().Format(time.RFC3339))
}

func Webhook1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Webhook 1 received")
}

func Webhook2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Webhook 2 received")
}

func DevHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	log.Println("DevHandler started")

	log.Printf("Using Spreadsheet ID: %s", spreadsheetID)

	readData, err := readFromSpreadsheet(ctx)
	if err != nil {
		log.Printf("Error reading from spreadsheet: %v", err)
		http.Error(w, "Failed to read from spreadsheet", http.StatusInternalServerError)
		return
	}

	writeData := [][]interface{}{
		{"New", "Data"},
		{"To", "Write"},
	}
	err = writeToSpreadsheet(ctx, writeData)
	if err != nil {
		log.Printf("Error writing to spreadsheet: %v", err)
		http.Error(w, "Failed to write to spreadsheet", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"read_data":  readData,
		"write_data": writeData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("DevHandler completed successfully")
}

func readFromSpreadsheet(ctx context.Context) ([][]interface{}, error) {
	log.Println("Starting readFromSpreadsheet")
	srv, err := getSheetService(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}
	log.Println("Sheets service created successfully")

	// スプレッドシートの情報を取得
	spreadsheet, err := srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		log.Printf("Error retrieving spreadsheet info: %v", err)
		return nil, fmt.Errorf("unable to retrieve spreadsheet info: %v", err)
	}

	// 最初のシートの名前を取得
	if len(spreadsheet.Sheets) == 0 {
		return nil, fmt.Errorf("no sheets found in the spreadsheet")
	}
	sheetName := spreadsheet.Sheets[0].Properties.Title

	readRange := fmt.Sprintf("%s!A1:B", sheetName)
	log.Printf("Attempting to read range: %s", readRange)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Printf("Error retrieving data: %v", err)
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	log.Printf("Data retrieved successfully. Rows: %d", len(resp.Values))
	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	return resp.Values, nil
}

func writeToSpreadsheet(ctx context.Context, values [][]interface{}) error {
	log.Println("Starting writeToSpreadsheet")
	srv, err := getSheetService(ctx)
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}
	log.Println("Sheets service created successfully")

	writeRange := "Sheet1!A1" // 書き込む範囲
	log.Printf("Attempting to write to range: %s", writeRange)
	valueRange := &sheets.ValueRange{
		Values: values,
	}
	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Printf("Error writing data: %v", err)
		return fmt.Errorf("unable to write data to sheet: %v", err)
	}

	log.Println("Data written successfully")
	return nil
}

func getSheetService(ctx context.Context) (*sheets.Service, error) {
	log.Println("Creating Sheets service")
	if serviceAccountKey == "" {
		return nil, fmt.Errorf("service account key is empty")
	}
	srv, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(serviceAccountKey)))
	if err != nil {
		log.Printf("Error creating Sheets service: %v", err)
		return nil, err
	}
	log.Println("Sheets service created successfully")
	return srv, nil
}