package philippines

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

//service
type service struct {
	h http.Client
}

//Service interface
type Service interface {
	GetStats() (*StatsAttributes, error)
	GetHospitalPUI() (*[]HsPUIsAttributes, error)
}

//GetStats get ph data
func (s service) GetStats() (*StatsAttributes, error) {

	param := phParams{
		F:              "json",
		Where:          "1=1",
		ReturnGeometry: "false",
		SpatialRel:     "esriSpatialRelIntersects",
		OutFields:      "*",
		CacheHint:      "false",
	}

	params, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s%s", baseApi, currentStats, params.Encode())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var baseResponse statsBase
	err = json.NewDecoder(res.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	var response StatsAttributes

	if len(baseResponse.Features) > 0 {
		response = baseResponse.Features[0].Attributes
	}

	return &response, nil
}

//GetHospitalPUI get hospital with puis
func (s service) GetHospitalPUI() (*[]HsPUIsAttributes, error) {
	param := phParams{
		F:              "json",
		Where:          "1=1",
		ReturnGeometry: "false",
		SpatialRel:     "esriSpatialRelIntersects",
		OutFields:      "*",
		CacheHint:      "false",
	}

	params, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s%s", baseApi, hospitalPUI, params.Encode())
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var baseResponse hsPUIBase
	err = json.NewDecoder(res.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	var response []HsPUIsAttributes

	for _, v := range baseResponse.Features {
		response = append(response, v.Attributes)
	}

	return &response, nil
}

//NewService create new service
func NewService(h http.Client) Service {
	return &service{h: h}
}
