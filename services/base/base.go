package base

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
)

type BaseService struct {
	URL         string
	TokenID     string
	TokenSecret string
	Client      *http.Client
}

func (s BaseService) NewHttpRequest(method string, endpoint string, body []byte) (r *http.Request, err error) {
	r, err = http.NewRequest(method, s.URL+path.Clean(fmt.Sprintf("/%s", endpoint)), bytes.NewBuffer(body))
	if err != nil {
		return
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s=%s", s.TokenID, s.TokenSecret))

	return
}
