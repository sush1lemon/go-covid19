package who

import (
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"net/http"
	"testing"
	"time"
)

func Test_service_History(t *testing.T) {
	list, _ := jsons.JsonCountries()
	c := http.Client{Timeout: time.Minute * 1}
	srv := NewService(c, *list)
	srv.History()
}