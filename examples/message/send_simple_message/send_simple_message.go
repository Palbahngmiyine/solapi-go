package main

import (
	"fmt"
	"os"

	"github.com/solapi/solapi-go/pkg/solapi"
)

func main() {
	// Get API key and API secret from environment variables
	apiKey := os.Getenv("SOLAPI_API_KEY")
	apiSecret := os.Getenv("SOLAPI_API_SECRET")

	// Create a new client
	client := solapi.MessageService(apiKey, apiSecret)

	// Prepare parameters for sending a simple message
	// Replace the phone numbers with actual values
	params := map[string]interface{}{
		"to":   "01000000000", // Recipient phone number
		"from": "01000000000", // Sender phone number
		"text": "This is a test message from solapi-go SDK",
		"type": "SMS", // SMS, LMS, MMS, ATA, CTA, CTI
	}

	// Send the message
	result, err := client.Messages.SendSimpleMessage(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Message sent successfully!")
	fmt.Println("Group ID:", result.GroupId)
	fmt.Println("Message ID:", result.MessageId)
	fmt.Println("Status:", result.StatusCode)
	fmt.Println("To:", result.To)
	fmt.Println("From:", result.From)
	fmt.Println("Type:", result.Type)
}
