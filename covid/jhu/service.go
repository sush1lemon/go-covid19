package jhu

import "net/http"

type service struct {
	h http.Client
}

//Service
type Service interface {

}

//NewService
func NewService(h http.Client) Service {
	return service{h:h}
}