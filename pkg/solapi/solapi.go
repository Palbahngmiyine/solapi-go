package solapi

import (
	"github.com/solapi/solapi-go/pkg/solapi/cash"
	"github.com/solapi/solapi-go/pkg/solapi/messages"
	"github.com/solapi/solapi-go/pkg/solapi/storage"
)

// Client struct
type Client struct {
	Messages messages.Messages
	Storage  storage.Storage
	Cash     cash.Cash
}

// MessageService returns a new client with API credentials
func MessageService(apiKey string, apiSecret string) *Client {
	client := Client{}

	// Initialize config map for each service
	config := map[string]string{
		"APIKey":    apiKey,
		"APISecret": apiSecret,
	}

	// Assign config to each service
	client.Messages.Config = config
	client.Storage.Config = config
	client.Cash.Config = config

	return &client
}
