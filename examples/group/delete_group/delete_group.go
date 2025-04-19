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

	// Delete the group
	result, err := client.Messages.DeleteGroup(groupId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Group deleted successfully!")
	fmt.Println("Group ID:", result.GroupId)
	fmt.Println("Status:", result.Status)
	fmt.Println("Account ID:", result.AccountId)
	fmt.Println("Date Created:", result.DateCreated)
	fmt.Println("Date Updated:", result.DateUpdated)
}
