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

	// Prepare parameters for creating a group
	// You can add additional parameters if needed
	params := map[string]string{
		"groupName": "My Group", // Optional group name
	}

	// Create a group
	result, err := client.Messages.CreateGroup(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Group created successfully!")
	fmt.Println("Group ID:", result.GroupId)
	fmt.Println("Account ID:", result.AccountId)
	fmt.Println("Status:", result.Status)
	fmt.Println("Date Created:", result.DateCreated)

	// Store the group ID for use in other examples
	fmt.Println("\nUse this Group ID in other examples:")
	fmt.Println(result.GroupId)
}
