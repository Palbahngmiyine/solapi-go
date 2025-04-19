package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/solapi/solapi-go/pkg/solapi/types"
)

// createMessagesInstance creates a Messages instance with the standard config
func createMessagesInstance(serverURL string) *Messages {
	return &Messages{
		Config: map[string]string{
			"APIKey":    "test_api_key",
			"APISecret": "test_api_secret",
			"Protocol":  "http",
			"Domain":    serverURL[7:], // Remove "http://" prefix
			"Prefix":    "",
		},
	}
}

// mockServer creates a test HTTP server that returns the specified response
func mockServer(t *testing.T, expectedMethod, expectedPath string, statusCode int, response interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check method
		if r.Method != expectedMethod {
			t.Errorf("Expected method %s, got %s", expectedMethod, r.Method)
		}

		// Check path
		if r.URL.Path != "/"+expectedPath {
			t.Errorf("Expected path /%s, got %s", expectedPath, r.URL.Path)
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		// Write response
		if response != nil {
			if err := json.NewEncoder(w).Encode(response); err != nil {
				t.Errorf("Failed to encode response: %v", err)
			}
		}
	}))
}

// Test SendSimpleMessage
func TestSendSimpleMessage(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		params       map[string]interface{}
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name: "Successful message send",
			params: map[string]interface{}{
				"to":   "01000000000",
				"from": "01000000000",
				"text": "Test message",
				"type": "SMS",
			},
			mockResponse: types.SimpleMessage{
				GroupId:       "G4V20180307105937H3PFASXMNJG2JID",
				MessageId:     "M4V20180307105937H3PTASXMNJG2JID",
				AccountId:     "12345678901234",
				StatusMessage: "정상 접수(이통사로 접수 예정) ",
				StatusCode:    "2000",
				To:            "01000000000",
				From:          "01000000000",
				Type:          "SMS",
				Country:       "82",
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name: "API error",
			params: map[string]interface{}{
				"to":   "01000000000",
				"from": "01000000000",
				"text": "Test message",
				"type": "SMS",
			},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			server := mockServer(t, "POST", "messages/v4/send-many/detail", tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.SendSimpleMessage(tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test GetMessageList
func TestGetMessageList(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		params       map[string]string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name: "Successful get message list",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.MessageList{
				StartKey: "",
				NextKey:  "",
				Limit:    10,
				MessageList: map[string]types.Message{
					"M4V20180307105937H3PTASXMNJG2JID": {
						MessageId:  "M4V20180307105937H3PTASXMNJG2JID",
						GroupId:    "G4V20180307105937H3PFASXMNJG2JID",
						StatusCode: "2000",
						To:         "01000000000",
						From:       "01000000000",
						Type:       "SMS",
						Text:       "Test message",
					},
				},
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name: "API error",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			server := mockServer(t, "GET", "messages/v4/list", tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.GetMessageList(tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test CreateGroup
func TestCreateGroup(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		params       map[string]string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:   "Successful create group",
			params: map[string]string{},
			mockResponse: types.Group{
				GroupId:   "G4V20180307105937H3PFASXMNJG2JID",
				AccountId: "12345678901234",
				Status:    "PENDING",
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:   "API error",
			params: map[string]string{},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			server := mockServer(t, "POST", "messages/v4/groups", tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.CreateGroup(tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test AddGroupMessage
func TestAddGroupMessage(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		groupId      string
		params       interface{}
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:    "Successful add group message",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			params: []map[string]interface{}{
				{
					"to":   "01000000000",
					"from": "01000000000",
					"text": "Test message",
					"type": "SMS",
				},
			},
			mockResponse: types.AddGroupMessageList{
				ErrorCount: 0,
				ResultList: []types.AddGroupMessage{
					{
						MessageId:     "M4V20180307105937H3PTASXMNJG2JID",
						To:            "01000000000",
						From:          "01000000000",
						Type:          "SMS",
						StatusCode:    "2000",
						StatusMessage: "정상 접수(이통사로 접수 예정) ",
						AccountId:     "12345678901234",
					},
				},
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:    "API error",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			params:  []map[string]interface{}{},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			path := fmt.Sprintf("messages/v4/groups/%s/messages", tc.groupId)
			server := mockServer(t, "PUT", path, tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.AddGroupMessage(tc.groupId, tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test SendGroup
func TestSendGroup(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		groupId      string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:    "Successful send group",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.Group{
				GroupId:   "G4V20180307105937H3PFASXMNJG2JID",
				AccountId: "12345678901234",
				Status:    "SENDING",
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:    "API error",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			path := fmt.Sprintf("messages/v4/groups/%s/send", tc.groupId)
			server := mockServer(t, "POST", path, tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.SendGroup(tc.groupId)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test DeleteGroup
func TestDeleteGroup(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		groupId      string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:    "Successful delete group",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.Group{
				GroupId:   "G4V20180307105937H3PFASXMNJG2JID",
				AccountId: "12345678901234",
				Status:    "DELETED",
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:    "API error",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			path := fmt.Sprintf("messages/v4/groups/%s", tc.groupId)
			server := mockServer(t, "DELETE", path, tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.DeleteGroup(tc.groupId)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test GetGroupList
func TestGetGroupList(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		params       map[string]string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name: "Successful get group list",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.GroupList{
				StartKey: "",
				NextKey:  "",
				Limit:    10,
				GroupList: map[string]types.Group{
					"G4V20180307105937H3PFASXMNJG2JID": {
						GroupId:   "G4V20180307105937H3PFASXMNJG2JID",
						AccountId: "12345678901234",
						Status:    "PENDING",
					},
				},
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name: "API error",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			server := mockServer(t, "GET", "messages/v4/groups", tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.GetGroupList(tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test GetGroup
func TestGetGroup(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		groupId      string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:    "Successful get group",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.Group{
				GroupId:   "G4V20180307105937H3PFASXMNJG2JID",
				AccountId: "12345678901234",
				Status:    "PENDING",
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:    "API error",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			path := fmt.Sprintf("messages/v4/groups/%s", tc.groupId)
			server := mockServer(t, "GET", path, tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.GetGroup(tc.groupId)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}

// Test GetGroupMessageList
func TestGetGroupMessageList(t *testing.T) {
	// Test cases
	testCases := []struct {
		name         string
		groupId      string
		params       map[string]string
		mockResponse interface{}
		statusCode   int
		expectError  bool
	}{
		{
			name:    "Successful get group message list",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.MessageList{
				StartKey: "",
				NextKey:  "",
				Limit:    10,
				MessageList: map[string]types.Message{
					"M4V20180307105937H3PTASXMNJG2JID": {
						MessageId:  "M4V20180307105937H3PTASXMNJG2JID",
						GroupId:    "G4V20180307105937H3PFASXMNJG2JID",
						StatusCode: "2000",
						To:         "01000000000",
						From:       "01000000000",
						Type:       "SMS",
						Text:       "Test message",
					},
				},
			},
			statusCode:  200,
			expectError: false,
		},
		{
			name:    "API error",
			groupId: "G4V20180307105937H3PFASXMNJG2JID",
			params: map[string]string{
				"limit": "10",
			},
			mockResponse: types.CustomError{
				ErrorCode:    "API_ERROR",
				ErrorMessage: "Bad Request",
			},
			statusCode:  400,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			path := fmt.Sprintf("messages/v4/groups/%s/messages", tc.groupId)
			server := mockServer(t, "GET", path, tc.statusCode, tc.mockResponse)
			defer server.Close()

			// Create a Messages instance with the mock server URL
			messages := createMessagesInstance(server.URL)

			// Call the function
			result, err := messages.GetGroupMessageList(tc.groupId, tc.params)

			// Check error
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got: %v", tc.expectError, err != nil)
				if err != nil {
					t.Errorf("Error: %v", err)
				}
			}

			// If no error and status code is 200, check result
			if err == nil && tc.statusCode == 200 {
				if !reflect.DeepEqual(result, tc.mockResponse) {
					t.Errorf("Expected result %v, got %v", tc.mockResponse, result)
				}
			}
		})
	}
}
