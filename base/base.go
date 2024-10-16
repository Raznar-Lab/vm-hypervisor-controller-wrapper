package base

import "net/http"

type BaseService struct {
	URL         string
	TokenID     string
	TokenSecret string
	Client      *http.Client
}