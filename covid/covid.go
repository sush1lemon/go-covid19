package covid

import (
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"net/http"
	"time"
)

//Client
type Client struct {
	Philippines philippines.Service
	Worldometer worldometer.Service
}

//NewClient
func NewClient() *Client  {
	h := http.Client{
		Timeout:       60 * time.Second,
	}
	return &Client{
		Philippines: philippines.NewService(h),
		Worldometer: worldometer.NewService(h),
	}
}