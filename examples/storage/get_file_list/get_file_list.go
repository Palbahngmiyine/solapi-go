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

	// Prepare parameters for getting file list
	// You can add query parameters here if needed
	params := map[string]string{}

	// Get the file list
	result, err := client.Storage.GetFileList(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Files retrieved successfully!")
	fmt.Println("Number of files:", len(result.FileList))

	// Print details of each file
	for i, file := range result.FileList {
		fmt.Printf("\nFile %d:\n", i+1)
		fmt.Println("  File ID:", file.FileId)
		fmt.Println("  File Name:", file.Name)
		fmt.Println("  Original Name:", file.OriginalName)
		fmt.Println("  URL:", file.Url)
		fmt.Println("  Date Created:", file.DateCreated)
	}
}