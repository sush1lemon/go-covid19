package rest

import (
	"encoding/json"
	"fmt"
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/who"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"github.com/jezerdave/go-covid19/src/storage"
	"github.com/jezerdave/go-covid19/src/updating"
	"github.com/labstack/echo/v4"
	"strings"
)

// Handler
type Handler struct {
	R     storage.Storage
	UpSrv updating.Service
}

// GetCountriesData get all countries
// @Tags Worldometer
// @Summary Get all countries data
// @Description get current covid data
// @Produce  json
// @Success 200 {array} worldometer.CountryStats
// @Router /countries [get]
func (h Handler) GetCountriesData(c echo.Context) error {

	raw, err := h.R.GetAll("country:*")
	if err != nil {
		return err
	}

	var response []worldometer.CountryStats
	for _, v := range *raw {
		var holder worldometer.CountryStats
		err := json.Unmarshal([]byte(v), &holder)
		if err == nil {
			response = append(response, holder)
		}
	}

	return c.JSON(200, &response)
}

//GetStatesData godoc
// @Tags Worldometer
// @Summary Get all US-States data
// @Description get current US-States covid data
// @Produce  json
// @Success 200 {array} worldometer.StatesStats
// @Router /states [get]
func (h Handler) GetStatesData(c echo.Context) error {
	raw, err := h.R.GetAll("us-state*")
	if err != nil {
		return err
	}

	var response []worldometer.StatesStats
	for _, v := range *raw {
		var holder worldometer.StatesStats
		err := json.Unmarshal([]byte(v), &holder)
		if err == nil {
			response = append(response, holder)
		}
	}

	return c.JSON(200, &response)
}

//FindCountry godoc
// @Tags Worldometer
// @Summary Find Country
// @Description find covid19 related data by country
// @Produce  json
// @Param country path string true "country name / code (philippines/ph/608/phl)"
// @Success 200 {object} worldometer.CountryStats
// @Failure 404 {object} BasicError
// @Router /countries/{country} [get]
func (h Handler) FindCountry(c echo.Context) error {
	var basicErr BasicError
	parameter := strings.ToLower(c.Param("country"))

	var response worldometer.CountryStats

	raw, err := h.R.Find(fmt.Sprintf("country:*:%s:*", parameter))
	if err != nil {
		basicErr.Message = "Cannot find given country"
		return c.JSON(404, basicErr)
	}

	for _, v := range *raw {
		err := json.Unmarshal([]byte(v), &response)
		if err != nil {
			return err
		}
	}

	if response.Country == "" {
		basicErr.Message = "Cannot find given country"
		return c.JSON(404, basicErr)
	}

	return c.JSON(200, response)
}

//FindStates godoc
// @Tags Worldometer
// @Summary Find State
// @Description find covid19 related data by state
// @Produce  json
// @Param state path string true "State Name / State Abbreviation"
// @Success 200 {object} worldometer.StatesStats
// @Failure 404 {object} BasicError
// @Router /states/{state} [get]
func (h Handler) FindStates(c echo.Context) error {
	var basicErr BasicError
	basicErr.Message = "Cannot find given state"
	parameter := strings.ToLower(c.Param("state"))

	var response worldometer.StatesStats

	raw, err := h.R.Find(fmt.Sprintf("*us-state:*:%s:*", parameter))
	if err != nil {
		return c.JSON(404, basicErr)
	}

	for _, v := range *raw {
		json.Unmarshal([]byte(v), &response)
	}

	if response.State == "" {
		return c.JSON(404, basicErr)
	}

	return c.JSON(200, response)
}

//GetPHStats godoc
// @Tags Unavailable
// @Summary GET DOH PHILIPPINES DATA
// @Description get doh philippines official data from https://ncovtracker.doh.gov.ph/
// @Produce  json
// @Success 200 {array} philippines.StatsAttributes
// @Router /doh/ph [get]
func (h Handler) GetPHStats(c echo.Context) error {
	var basicErr BasicError
	basicErr.Message = "Cannot find DOH Philippines latest data"

	key := "doh-philippines-latest"
	raw, err := h.R.Get(key)
	if err != nil {
		return err
	}

	if raw == nil {
		return c.JSON(404, basicErr)
	}

	var response philippines.StatsAttributes
	err = json.Unmarshal([]byte(*raw), &response)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

//GetPHHospitalPUIs godoc
// @Tags Unavailable
// @Summary GET DOH PHILIPPINES DATA (HOSPITAL PUIs)
// @Description get doh philippines official data (HOSPITAL PUIs) from https://ncovtracker.doh.gov.ph/
// @Produce  json
// @Success 200 {array} philippines.HsPUIsAttributes
// @Router /doh/ph/hospital-pui [get]
func (h Handler) GetPHHospitalPUIs(c echo.Context) error {
	var basicErr BasicError
	basicErr.Message = "Cannot find DOH Philippines Hospital PUIs latest data"

	key := "doh-philippines-hospitalpui-latest"
	raw, err := h.R.Get(key)
	if err != nil {
		return err
	}

	if raw == nil {
		return c.JSON(404, basicErr)
	}

	var response []philippines.HsPUIsAttributes
	err = json.Unmarshal([]byte(*raw), &response)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}

//GetHistories
// @Tags World Health Organization (WHO)
// @Summary Get all countries historical data
// @Description get histories
// @Produce  json
// @Success 200 {array} who.HistoryData
// @Router /histories [get]
func (h Handler) GetHistories(c echo.Context) error {
	raw, err := h.R.GetAll("history:*")
	if err != nil {
		return err
	}

	var response []who.SimpleHistoryData
	for _, v := range *raw {
		var holder who.HistoryData
		err := json.Unmarshal([]byte(v), &holder)
		if err == nil{
			simple := holder.Simplify()
			response = append(response, *simple)
		}
	}

	return c.JSON(200, &response)
}

//FindCountryHistories
// @Tags World Health Organization (WHO)
// @Summary Find Country Histories
// @Description find covid19 related historical data by country
// @Produce  json
// @Param country path string true "country name / code (philippines/ph/608/phl)"
// @Success 200 {object} who.HistoryData
// @Failure 404 {object} BasicError
// @Router /histories/{country} [get]
func (h Handler) FindCountryHistories(c echo.Context) error {
	var basicErr BasicError
	parameter := strings.ToLower(c.Param("country"))

	var response who.HistoryData
	var simple *who.SimpleHistoryData

	raw, err := h.R.Find(fmt.Sprintf("history:*:%s:*", parameter))
	if err != nil {
		basicErr.Message = "Cannot find history of the given country"
		return c.JSON(404, basicErr)
	}

	for _, v := range *raw {
		err := json.Unmarshal([]byte(v), &response)
		if err != nil {
			return err
		}
		simple = response.Simplify()
	}

	if response.Country == "" {
		basicErr.Message = "Cannot find history of the given country"
		return c.JSON(404, basicErr)
	}

	return c.JSON(200, &simple)
}

//UpdateData - update
func (h Handler) UpdateData(c echo.Context) error {

	response, err := h.UpSrv.UpdateData()
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}
