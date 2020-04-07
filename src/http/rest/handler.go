package rest

import (
	"encoding/json"
	"fmt"
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"github.com/jezerdave/go-covid19/src/storage"
	"github.com/labstack/echo/v4"
	"strings"
	"sync"
)

type Handler struct {
	PH  philippines.Service
	WOM worldometer.Service
	R   storage.Storage
	CountryList
	States
}

// GetCountriesData godoc
// @Summary Get all countries data
// @Description get current covid data
// @Produce  json
// @Success 200 {array} worldometer.CountryStats
// @Router /countries [get]
func (h Handler) GetCountriesData(c echo.Context) error {

	raw, err := h.R.GetAll("countries:*")
	if err != nil {
		return err
	}

	var response []worldometer.CountryStats
	for _, v := range *raw {
		var holder worldometer.CountryStats
		json.Unmarshal([]byte(v), &holder)
		response = append(response, holder)
	}

	return c.JSON(200, &response)
}

//GetStatesData
// @Summary Get all US-States data
// @Description get current US-States covid data
// @Produce  json
// @Success 200 {array} worldometer.StatesStats
// @Router /states [get]
func (h Handler) GetStatesData(c echo.Context) error {
	raw, err := h.R.GetAll("us-states:*")
	if err != nil {
		return err
	}

	var response []worldometer.StatesStats
	for _, v := range *raw {
		var holder worldometer.StatesStats
		json.Unmarshal([]byte(v), &holder)
		response = append(response, holder)
	}

	return c.JSON(200, &response)
}

//FindCountry
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

	raw, err := h.R.Find(fmt.Sprintf("countries:*:%s:*", parameter))
	if err != nil {
		basicErr.Message = "Cannot find given country"
		return c.JSON(404, basicErr)
	}

	for _, v := range *raw {
		json.Unmarshal([]byte(v), &response)
	}

	if response.CountryInfo == nil {
		basicErr.Message = "Cannot find given country"
		return c.JSON(404, basicErr)
	}

	return c.JSON(200, response)
}

//FindStates
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

	raw, err := h.R.Find(fmt.Sprintf("*us-states:*:%s:*", parameter))
	if err != nil {
		return  c.JSON(404, basicErr)
	}

	for _, v := range *raw {
		json.Unmarshal([]byte(v), &response)
	}

	if response.StateInfo == nil {
		return  c.JSON(404, basicErr)
	}

	return c.JSON(200, response)
}

//GetPHStats
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

//GetPHHospitalPUIs
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

//UpdateData
func (h Handler) UpdateData(c echo.Context) error {

	var countries *[]worldometer.CountryStats
	var states *[]worldometer.StatesStats
	var phStats *philippines.StatsAttributes
	var phhspui *[]philippines.HsPUIsAttributes

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		countries, _ = h.WOM.GetCountriesData()
	}()
	go func() {
		defer wg.Done()
		states, _ = h.WOM.GetStatesData()
	}()
	go func() {
		defer wg.Done()
		phStats, _ = h.PH.GetStats()
	}()
	go func() {
		defer wg.Done()
		phhspui, _ = h.PH.GetHospitalPUI()
	}()

	wg.Wait()

	pSK := "doh-philippines-latest"
	err := h.R.New(pSK, phStats)
	if err != nil {
		return err
	}

	pHK := "doh-philippines-hospitalpui-latest"
	err = h.R.New(pHK, phhspui)
	if err != nil {
		return err
	}

	for _, v := range *countries {
		for _, c := range h.CountryList.Countries {
			if v.Country == c.Name.Official || v.Country == c.Name.Common {
				key := fmt.Sprintf("countries::%s:%s:%s:%s:", strings.ToLower(c.Name.Common),
					strings.ToLower(c.Ccn3), strings.ToLower(c.Cca2), strings.ToLower(c.Cca3))
				v.CountryInfo = c
				h.R.New(key, v)
			}
		}

	}

	for _, v := range *states {
		for i, s := range h.States.States {
			if v.State == s {
				key := fmt.Sprintf("us-states::%s:%s:", strings.ToLower(v.State), strings.ToLower(i))
				info := make(map[string]interface{})
				info["name"] = s
				info["abbreviation"] = i
				v.StateInfo = info
				h.R.New(key, v)
			}
		}
	}

	response := make(map[string]interface{})

	response["worldometer_countries"] = countries
	response["worldometer_states"] = states
	response["doh_ph_stats"] = phStats
	response["doh_ph_hospital_pui"] = phhspui

	return c.JSON(200, response)
}
