package covid

import (
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"github.com/jezerdave/go-covid19/covid/who"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"net/http"
	"time"
)

//Client struct
type Client struct {
	Philippines philippines.Service
	Worldometer worldometer.Service
	WHO         who.Service
}

//NewClient create new client
func NewClient() *Client {

	list, _ := jsons.JsonCountries()
	states, _ := jsons.JsonStates()

	h := http.Client{
		Timeout: 60 * time.Second,
	}
	return &Client{
		Philippines: philippines.NewService(h),
		Worldometer: worldometer.NewService(h, *list, *states),
		WHO:         who.NewService(h, *list),
	}
}
