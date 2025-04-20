# solapi-go

[Site](https://www.solapi.com/) |
[Docs](https://docs.solapi.com/) |
[Examples](https://github.com/solapi/solapi-go/tree/master/examples) |

문자 메시지 발송 및 조회 관련 기능들을 쉽게 사용하실 수 있도록 만들어진 SDK 입니다.

## Example

```go
package main

import (
	"fmt"

	"github.com/solapi/solapi-go/pkg/solapi"
)

func main() {
	// API Key와 API Secret Key는 테스트 시 활성화 된 키로 입력해주세요.
	apiKey := os.Getenv("SOLAPI_API_KEY")
	apiSecret := os.Getenv("SOLAPI_API_SECRET")

	client := solapi.MessageService(apiKey, apiSecret)

	// Message Data
	// 관련 파라미터들은 https://docs.solapi.com에서 확인 가능합니다.
	message := make(map[string]interface{})
	message["to"] = "01000000000"
	message["from"] = "029302266"
	message["text"] = "Test Message"
	message["type"] = "SMS"

	params := map[string]interface{}{
		"to":   "01000000000", // Recipient phone number
		"from": "01000000000", // Sender phone number
		"text": "This is a test message from solapi-go SDK",		
	}

	// Call API Resource
	result, err := client.Messages.Send(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
```

[examples folder](https://github.com/solapi/solapi-go/tree/master/examples)에서 자세한 예제파일들을 확인하세요.

## Installation

```
go get github.com/solapi/solapi-go/pkg/solapi
```

## Project Structure

이 프로젝트는 Go 표준 레이아웃을 따릅니다:

- `pkg/solapi`: 메인 패키지 및 하위 패키지들
  - `apirequest`: API 요청 관련 기능
  - `cash`: 잔액 조회 관련 기능
  - `messages`: 메시지 발송 및 조회 관련 기능
  - `storage`: 파일 업로드 및 조회 관련 기능
  - `types`: 데이터 타입 정의
- `examples`: 예제 코드
