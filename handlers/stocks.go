package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Stocks(c *fiber.Ctx) error {
	DhanHistoricalData()
	return c.SendString("Stocks Fetched")
}

func DhanHistoricalData() {
	// Prepare the request data
	token := os.Getenv("TOKEN")
	// bearerToken := fmt.Sprintf("Bearer %s", token)
	data := map[string]interface{}{
		"symbol":          "TCS",
		"exchangeSegment": "NSE_EQ",
		"instrument":      "EQUITY",
		"expiryCode":      0,
		"fromDate":        "2021-01-01",
		"toDate":          "2023-06-19",
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.dhan.co/charts/historical", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access-token", token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	// Convert body to string and print
	// Parse response into struct
	var response DhanResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		os.Exit(1)
	}
	fmt.Println("Open prices:", response)
	fmt.Println("Close prices:", response.Close)
}

type DhanResponse struct {
	Open      []float64 `json:"open"`
	High      []float64 `json:"high"`
	Low       []float64 `json:"low"`
	Close     []float64 `json:"close"`
	Volume    []int64   `json:"volume"`
	StartTime []int64   `json:"start_Time"`
}
