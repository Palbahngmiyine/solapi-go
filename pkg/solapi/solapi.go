package solapi

import (
	"github.com/solapi/solapi-go/pkg/solapi/cash"
	"github.com/solapi/solapi-go/pkg/solapi/messages"
	"github.com/solapi/solapi-go/pkg/solapi/storage"
	"github.com/solapi/solapi-go/pkg/solapi/types"
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
	config := types.Config{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}

	// Assign config to each service
	client.Messages.Config = config
	client.Storage.Config = config
	client.Cash.Config = config

	return &client
}
