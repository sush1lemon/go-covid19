package updating

import (
	"fmt"
	"github.com/jezerdave/go-covid19/covid"
	"github.com/jezerdave/go-covid19/covid/philippines"
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"github.com/jezerdave/go-covid19/covid/who"
	"github.com/jezerdave/go-covid19/covid/worldometer"
	"github.com/jezerdave/go-covid19/src/storage"
	"strings"
	"sync"
)

// service
type service struct {
	repo      storage.Storage
	cl        covid.Client
	countries jsons.CountryList
	states    jsons.States
}

//UpdateData
func (s service) UpdateData() (interface{}, error) {

	var countries *[]worldometer.CountryStats
	var states *[]worldometer.StatesStats
	var phStats *philippines.StatsAttributes
	var phhspui *[]philippines.HsPUIsAttributes
	var history *[]who.HistoryData

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		countries, _ = s.cl.Worldometer.GetCountriesData()
	}()
	go func() {
		defer wg.Done()
		states, _ = s.cl.Worldometer.GetStatesData()
	}()
	go func() {
		defer wg.Done()
		phStats, _ = s.cl.Philippines.GetStats()
	}()
	go func() {
		defer wg.Done()
		phhspui, _ = s.cl.Philippines.GetHospitalPUI()
	}()
	go func() {
		defer wg.Done()
		history, _ = s.cl.WHO.History()
	}()

	wg.Wait()

	pSK := "doh-philippines-latest"
	err := s.repo.New(pSK, phStats)
	if err != nil {
		return nil, err
	}

	pHK := "doh-philippines-hospitalpui-latest"
	err = s.repo.New(pHK, phhspui)
	if err != nil {
		return nil, err
	}

	for _, v := range *countries {
		for _, c := range s.countries.Countries {
			if v.Country == c.Name.Official || v.Country == c.Name.Common {
				key := fmt.Sprintf("country::%s:%s:%s:%s:", strings.ToLower(c.Name.Common),
					strings.ToLower(c.Ccn3), strings.ToLower(c.Cca2), strings.ToLower(c.Cca3))
				v.CountryInfo = c
				go s.repo.New(key, v)
			}
		}

	}

	for _, v := range *states {
		for i, st := range s.states.States {
			if v.State == st {
				key := fmt.Sprintf("us-state::%s:%s:", strings.ToLower(v.State), strings.ToLower(i))
				info := make(map[string]interface{})
				info["name"] = st
				info["abbreviation"] = i
				v.StateInfo = info
				go s.repo.New(key, v)
			}
		}
	}

	for _, v := range *history {
		key := fmt.Sprintf("history::%s:%s:%s:%s:", strings.ToLower(v.CountryInfo.Name.Common),
			strings.ToLower(v.CountryInfo.Ccn3), strings.ToLower(v.CountryInfo.Cca2), strings.ToLower(v.CountryInfo.Cca3))
		go s.repo.New(key, v)
	}

	response := make(map[string]interface{})

	response["worldometer_countries"] = countries
	response["worldometer_states"] = states
	response["doh_ph_stats"] = phStats
	response["doh_ph_hospital_pui"] = phhspui
	response["who_history"] = history

	return response, nil
}

//Service
type Service interface {
	UpdateData() (interface{}, error)
}

//NewService
func NewService(repo storage.Storage, api covid.Client, list jsons.CountryList, states jsons.States) Service {
	return service{
		repo:      repo,
		cl:        api,
		countries: list,
		states:    states,
	}
}
