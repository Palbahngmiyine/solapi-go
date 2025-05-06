package messages

import (
	"fmt"

	"github.com/solapi/solapi-go/pkg/solapi/fetcher"
	"github.com/solapi/solapi-go/pkg/solapi/types"
)

// Messages struct
type Messages struct {
	Config types.Config
}

// GetMessages gets the list of messages
func (r *Messages) GetMessages(params map[string]string) (types.MessageList, error) {
	result := types.MessageList{}
	err := fetcher.Request("GET", "messages/v4/list", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Send sends a simple message
func (r *Messages) Send(params map[string]interface{}) (types.MessageStruct, error) {
	// Create a fetcher instance to access its properties

	// Get the messages array
	messagesArray, ok := params["messages"].([]map[string]interface{})
	if !ok {
		// If messages key doesn't exist or is not an array of maps, wrap the entire params in a messages key
		// as a single element in the array
		messagesArray = []map[string]interface{}{params}
		params = map[string]interface{}{
			"messages": messagesArray,
		}
	}

	// Create agent info
	// agent := map[string]string{"sdkVersion": fetcher.Fetcher.Request(), "osPlatform": request.OsPlatform}
	// TODO: Will fix
	//if request.AppId != "" {
	//	agent["appId"] = request.AppId
	//}
	// params["agent"] = agent

	result := types.MessageStruct{}
	err := fetcher.Request("POST", "messages/v4/send-many/detail", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateGroup creeate message group
func (r *Messages) CreateGroup(params map[string]string) (types.Group, error) {
	// Create a fetcher instance to access its properties
	//params["sdkVersion"] = request.SdkVersion
	//params["osPlatform"] = request.OsPlatform
	// TODO: Will fix
	//if request.AppId != "" {
	//	params["appId"] = request.AppId
	//}
	result := types.Group{}
	err := fetcher.Request("POST", "messages/v4/groups", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// AddGroupMessage adds a group message
func (r *Messages) AddGroupMessage(groupId string, params interface{}) (types.AddGroupMessageList, error) {
	result := types.AddGroupMessageList{}
	url := fmt.Sprintf("messages/v4/groups/%s/messages", groupId)
	err := fetcher.Request("PUT", url, params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SendGroup send a group
func (r *Messages) SendGroup(groupId string) (types.Group, error) {
	result := types.Group{}
	url := fmt.Sprintf("messages/v4/groups/%s/send", groupId)
	params := make(map[string]string)
	err := fetcher.Request("POST", url, params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// DeleteGroup delete a group
func (r *Messages) DeleteGroup(groupId string) (types.Group, error) {
	result := types.Group{}
	url := fmt.Sprintf("messages/v4/groups/%s", groupId)
	params := make(map[string]string)
	err := fetcher.Request("DELETE", url, params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroupList gets the list of groups
func (r *Messages) GetGroupList(params map[string]string) (types.GroupList, error) {
	result := types.GroupList{}
	err := fetcher.Request("GET", "messages/v4/groups", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroup get a group
func (r *Messages) GetGroup(groupId string) (types.Group, error) {
	result := types.Group{}
	params := map[string]string{}
	url := fmt.Sprintf("messages/v4/groups/%s", groupId)
	err := fetcher.Request("GET", url, params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroupMessageList returns a list of group messages
func (r *Messages) GetGroupMessageList(groupId string, params map[string]string) (types.MessageList, error) {
	result := types.MessageList{}
	url := fmt.Sprintf("messages/v4/groups/%s/messages", groupId)
	err := fetcher.Request("GET", url, params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}
