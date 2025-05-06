package storage

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"github.com/solapi/solapi-go/pkg/solapi/fetcher"
	"github.com/solapi/solapi-go/pkg/solapi/types"
)

var (
	errFailToReadFile = errors.New("FailToReadFile")
	errFileNotFound   = errors.New("FileNotFound")
)

// Storage struct
type Storage struct {
	Config types.Config
}

// UploadFile upload a file
func (r *Storage) UploadFile(params map[string]string) (types.File, error) {
	result := types.File{}

	// 파일이 없다면 에러
	if _, ok := params["file"]; !ok {
		return result, errFileNotFound
	}

	// Open file
	f, err := os.Open(params["file"])
	if err != nil {
		return result, errFileNotFound
	}

	// Read entire contents into byte slice.
	reader := bufio.NewReader(f)
	content, err := io.ReadAll(reader)
	if err != nil {
		return result, errFailToReadFile
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to params.
	params["file"] = encoded

	err = fetcher.Request("POST", "storage/v1/files", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetFiles gets the list of files
func (r *Storage) GetFiles(params map[string]string) (types.FileList, error) {
	result := types.FileList{}
	err := fetcher.Request("GET", "storage/v1/files", params, &result, r.Config.ApiKey, r.Config.ApiSecret)
	if err != nil {
		return result, err
	}

	return result, nil
}
