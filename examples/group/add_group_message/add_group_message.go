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

	// Replace with your actual group ID
	// You can get a group ID by running the create_group example
	groupId := "GROUP_ID_HERE"

	// Prepare messages to add to the group
	messages := []map[string]interface{}{
		{
			"to":   "01000000001", // Recipient phone number
			"from": "01000000000", // Sender phone number
			"text": "This is the first message in the group",
			"type": "SMS",
		},
		{
			"to":   "01000000002", // Recipient phone number
			"from": "01000000000", // Sender phone number
			"text": "This is the second message in the group",
			"type": "SMS",
		},
	}

	// Add messages to the group
	result, err := client.Messages.AddGroupMessage(groupId, messages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Messages added to group successfully!")
	fmt.Println("Error Count:", result.ErrorCount)
	fmt.Println("Number of messages added:", len(result.ResultList))

	// Print details of each added message
	for i, message := range result.ResultList {
		fmt.Printf("\nMessage %d:\n", i+1)
		fmt.Println("  Message ID:", message.MessageId)
		fmt.Println("  To:", message.To)
		fmt.Println("  From:", message.From)
		fmt.Println("  Type:", message.Type)
		fmt.Println("  Status Code:", message.StatusCode)
	}
}
