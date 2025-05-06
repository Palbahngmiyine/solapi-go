package types

import "time"

type RequestDefaultParameter struct {
	Messages        RequestMessages `json:"messages"`
	AllowDuplicates *bool           `json:"allowDuplicates"`
	ShowMessageList *bool           `json:"showMessageList"`
	ScheduledDate   *time.Time      `json:"scheduledDate"`
}

type RequestMessages struct {
	To             string               `json:"to"`
	From           *string              `json:"from"`
	Text           *string              `json:"text"`
	FileId         *string              `json:"fileId"`
	Country        *string              `json:"country"`
	AutoTypeDetect *bool                `json:"autoTypeDetect"`
	Subject        *string              `json:"subject"`
	Type           *string              `json:"type"`
	KakaoOptions   *RequestKakaoOptions `json:"kakaoOptions"`
	RcsOptions     *RequestRcsOptions   `json:"rcsOptions"`
}

type RequestKakaoOptions struct {
	PfId       string             `json:"pfId"`
	TemplateId *string            `json:"templateId"`
	Variables  *map[string]string `json:"variables"`
	ImageId    *string            `json:"imageId"`
	Buttons    *map[string]string `json:"buttons"`
}

type RequestRcsOptions struct {
	brandId    string
	templateId string
}
