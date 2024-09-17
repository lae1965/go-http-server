package utils

import (
	"errors"
	"strings"
)

func ParseUrlParams(url string, endpoint string) (*string, error) {
	urlLength := len(url)
	endpointLength := len(endpoint)
	remainIndex := strings.LastIndex(url, endpoint)
	if remainIndex == -1 {
		return nil, errors.New("Эндпойнт " + endpoint + " не задан")
	}

	remainIndex += endpointLength
	if remainIndex >= urlLength {
		return nil, nil
	}
	if url[remainIndex] == '/' {
		remainIndex++
		if remainIndex >= urlLength {
			return nil, nil
		}
	}
	result := url[remainIndex:]
	return &result, nil
}
