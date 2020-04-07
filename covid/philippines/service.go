package philippines

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type service struct {
	h http.Client
}

//Service
type Service interface {
	GetStats() (*StatsAttributes, error)
	GetHospitalPUI() (*[]HsPUIsAttributes, error)
}

//GetStats
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

	response := StatsAttributes(baseResponse.Features[0].Attributes)
	return &response, nil
}

//GetHospitalPUI
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

//NewService
func NewService(h http.Client) Service {
	return &service{h: h}
}