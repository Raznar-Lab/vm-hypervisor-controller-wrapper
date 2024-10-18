package base_response

import (
	"fmt"
	"strings"
)

type BaseErrorResponse map[string]string


func (bErr BaseErrorResponse) Source() string {
	return bErr["source"]
}

func (bErr BaseErrorResponse) GetAllString() (sources string) {
	var rawSources []string
	for key, value := range bErr {
		rawSources = append(rawSources, fmt.Sprintf("%s=%s", key, value))
	}

	sources = fmt.Sprintf("[%s]", strings.Join(rawSources, ","))
	return
}

func (bErr BaseErrorResponse) Get(key string) string {
	for errKey, value := range bErr {
		if errKey == key {
			return value
		}
	}

	return ""
}