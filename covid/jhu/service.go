package jhu

import "net/http"

type service struct {
	h http.Client
}

//Service interface
type Service interface {
}

//NewService create new service
func NewService(h http.Client) Service {
	return service{h: h}
}
