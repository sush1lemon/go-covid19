package philippines

import (
	"net/http"
	"testing"
)

// TEst
func TestService(t *testing.T) {
	c := http.Client{
		Timeout:       30,
	}

	srv := NewService(c)
	srv.GetHospitalPUI()
}
