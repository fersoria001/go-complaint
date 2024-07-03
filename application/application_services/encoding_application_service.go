package application_services

import (
	"net/url"
	"regexp"
	"sync"
	"unicode/utf16"
)

var encodingApplicationService *EncodingApplicationService
var onceEncodingApplicationService sync.Once

func EncodingApplicationServiceInstance() *EncodingApplicationService {
	onceEncodingApplicationService.Do(func() {
		encodingApplicationService = NewEncodingApplicationService()
	})
	return encodingApplicationService
}

type EncodingApplicationService struct {
}

func NewEncodingApplicationService() *EncodingApplicationService {
	return &EncodingApplicationService{}
}

func (e *EncodingApplicationService) SafeUtf16Encode(input string) string {
	regex := regexp.MustCompile(`/%([0-9A-F]{2})/g`)
	scape := url.QueryEscape(input)
	replaced := regex.ReplaceAllStringFunc(scape, func(match string) string {
		return string(utf16.Decode([]uint16{uint16(match[1])}))
	})
	return replaced
}
