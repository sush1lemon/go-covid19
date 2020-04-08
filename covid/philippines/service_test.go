package philippines

import (
	"net/http"
	"testing"
)

// Test service
func TestService(t *testing.T) {
	c := http.Client{
		Timeout:       30,
	}

	srv := NewService(c)
	srv.GetHospitalPUI()
}
