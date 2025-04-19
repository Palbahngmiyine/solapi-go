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

	// Prepare parameters for getting group list
	// You can add query parameters here if needed
	params := map[string]string{
		"limit": "10", // Number of groups to retrieve
	}

	// Get the group list
	result, err := client.Messages.GetGroupList(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Groups retrieved successfully!")
	fmt.Println("Start Key:", result.StartKey)
	fmt.Println("Next Key:", result.NextKey)
	fmt.Println("Limit:", result.Limit)
	fmt.Println("Number of groups:", len(result.GroupList))

	// Print details of each group
	for groupId, group := range result.GroupList {
		fmt.Printf("\nGroup ID: %s\n", groupId)
		fmt.Println("  Status:", group.Status)
		fmt.Println("  Account ID:", group.AccountId)
		fmt.Println("  Date Created:", group.DateCreated)
		fmt.Println("  Count Total:", group.Count.Total)
		fmt.Println("  Count Sent Success:", group.Count.SentSuccess)
		fmt.Println("  Count Sent Failed:", group.Count.SentFailed)
	}
}