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

	// Prepare parameters for getting group message list
	// You can add query parameters here if needed
	params := map[string]string{
		"limit": "10", // Number of messages to retrieve
	}

	// Get the group message list
	result, err := client.Messages.GetGroupMessageList(groupId, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Group messages retrieved successfully!")
	fmt.Println("Start Key:", result.StartKey)
	fmt.Println("Next Key:", result.NextKey)
	fmt.Println("Limit:", result.Limit)
	fmt.Println("Number of messages:", len(result.MessageList))

	// Print details of each message
	for messageId, message := range result.MessageList {
		fmt.Printf("\nMessage ID: %s\n", messageId)
		fmt.Println("  From:", message.From)
		fmt.Println("  To:", message.To)
		fmt.Println("  Type:", message.Type)
		fmt.Println("  Status:", message.Status)
		fmt.Println("  Text:", message.Text)
		fmt.Println("  Date Created:", message.DateCreated)
	}
}