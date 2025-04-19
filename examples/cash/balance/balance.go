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

	// Get balance information
	result, err := client.Cash.Balance()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println("Balance information retrieved successfully!")
	fmt.Println("Account ID:", result.AccountId)
	fmt.Println("Balance:", result.Balance)
	fmt.Println("Point:", result.Point)

	// Print low balance alert information
	fmt.Println("\nLow Balance Alert Information:")
	fmt.Println("  Enabled:", result.LowBalanceAlert.Enabled)
	fmt.Println("  Current Balance:", result.LowBalanceAlert.CurrentBalance)
	fmt.Println("  Notification Balance:", result.LowBalanceAlert.NotificationBalance)

	// Print auto recharge information
	fmt.Println("\nAuto Recharge Information:")
	fmt.Println("  Auto Recharge:", result.AutoRecharge)
	fmt.Println("  Recharge To:", result.RechargeTo)
	fmt.Println("  Recharge Try Count:", result.RechargeTryCount)
}
