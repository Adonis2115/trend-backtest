package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

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
		"fromDate":        "2023-06-16",
		"toDate":          "2023-06-20",
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
	err = json.Unmarshal((body), &response)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		os.Exit(1)
	}

	t := convertEpochToTime(int(response.StartTime[0]))

	fmt.Println("Response:", response)
	fmt.Println("Open:", response.Open[0])
	fmt.Println("High:", response.High[0])
	fmt.Println("Low:", response.Low[0])
	fmt.Println("Close:", response.Close[0])
	fmt.Println("Volume:", response.Volume[0])
	fmt.Println("Start Time:", response.StartTime[0])
	fmt.Println("Time:", t)
}

type DhanResponse struct {
	Open      []float64 `json:"open"`
	High      []float64 `json:"high"`
	Low       []float64 `json:"low"`
	Close     []float64 `json:"close"`
	Volume    []float64 `json:"volume"`
	StartTime []float64 `json:"start_Time"`
}

func convertEpochToTime(epoch int) time.Time {
	// Get the time zone offset in minutes
	_, offset := time.Now().Zone()
	offsetMinutes := int(offset) / 60

	// Calculate the IST offset
	istOffset := int(330)

	// Adjust the epoch timestamp
	n := epoch - (istOffset+offsetMinutes)*60

	// Create a base time using January 1, 1980, 05:30:00 IST
	baseTime := time.Date(1980, 1, 1, 5, 30, 0, 0, time.FixedZone("IST", istOffset*60))

	// Create the final time by setting seconds and adjusting the date
	finalTime := baseTime.Add(time.Duration(n) * time.Second)

	finalTime = finalTime.AddDate(0, 0, 1) //add 1 day to it

	return finalTime
}
