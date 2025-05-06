package cash

import (
	"github.com/solapi/solapi-go/pkg/solapi/fetcher"
	"github.com/solapi/solapi-go/pkg/solapi/types"
)

// Cash struct
type Cash struct {
	Config types.Config
}

// GetBalance get balance information
func (r *Cash) GetBalance() (types.Balance, error) {
	result := types.Balance{}
	params := map[string]string{}
	err := fetcher.Request("GET", "cash/v1/balance", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}
