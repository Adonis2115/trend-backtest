package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	bearerToken := fmt.Sprintf("Bearer %s", token)
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
	req.Header.Set("Authorization", bearerToken)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response
	fmt.Println("Response status code:", resp.StatusCode)
	// You can read the response body if needed using resp.Body
}
