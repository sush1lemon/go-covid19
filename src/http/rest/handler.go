package rest

import (
	"encoding/json"
	"fmt"
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"github.com/jezerdave/go-covid19/src/storage"
	"github.com/labstack/echo"
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

//GetCountriesData
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
func (h Handler) FindStates(c echo.Context) error {
	var basicErr BasicError
	basicErr.Message = "Cannot find given state"
	parameter := strings.ToLower(c.Param("param"))

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

	fmt.Println(states)

	for _, v := range *states {
		for i, s := range h.States.States {
			fmt.Println(v.State)
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
