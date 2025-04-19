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

	// Prepare parameters for file upload
	// Replace "path/to/your/file.jpg" with the actual file path
	params := map[string]string{
		"file": "path/to/your/file.jpg",
		"name": "example_file",
	}

	// Upload the file
	result, err := client.Storage.UploadFile(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("File uploaded successfully!")
	fmt.Println("File ID:", result.FileId)
	fmt.Println("File Name:", result.Name)
	fmt.Println("File URL:", result.Url)
}