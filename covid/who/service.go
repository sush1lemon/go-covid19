package who

import (
	"encoding/json"
	"github.com/jezerdave/go-covid19/covid/util"
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"net/http"
	"strings"
	"time"
)

//Service interface
type Service interface {
	History() (*[]HistoryData, error)
	Find(string) (*HistoryData, error)
}

//Service struct
type service struct {
	h         http.Client
	countries jsons.CountryList
}

//History Get Historical covid19 data from WHO
func (s service) History() (*[]HistoryData, error) {

	var response []HistoryData

	r, err := s.h.Get(api)
	if err != nil {
		return nil, err
	}

	var raw rawData
	err = json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		return nil, err
	}

	var lastAbv string
	var hs HistoryData

	for i, v := range raw.Rows {
		cAbv := v[1].(string)

		var source *string
		if v[2] != nil {
			sHold := v[2].(string)
			source = &sHold
		}

		if cAbv != lastAbv && i != 0 {
			hs.Country = lastAbv
			response = append(response, hs)
			hs = HistoryData{}
		}

		data := Data{
			Date:               v[0].(float64),
			Region:             source,
			DeathDailyIncrease: v[3].(float64),
			TotalDeaths:        v[4].(float64),
			CaseDailyIncrease:  v[5].(float64),
			TotalCases:         v[6].(float64),
		}

		hs.Data = append(hs.Data, data)
		lastAbv = cAbv
	}

	var res []HistoryData

	for _, v := range response {
		for _, x := range s.countries.Countries {
			if strings.EqualFold(x.Cca2, v.Country) || strings.EqualFold(x.Cca3, v.Country) {
				v.Country = x.Name.Common
				v.CountryInfo = x
				res = append(res, v)
			}
		}
	}

	return &res, nil
}

//Find given country
func (s service) Find(param string) (*HistoryData, error) {

	list, err := s.History()
	if err != nil {
		return nil, err
	}

	for _, v := range *list {
		if strings.EqualFold(v.CountryInfo.Name.Common, param) || strings.EqualFold(v.CountryInfo.Name.Official, param) ||
			strings.EqualFold(v.CountryInfo.Cca2, param) || strings.EqualFold(v.CountryInfo.Cca3, param) ||
			strings.EqualFold(v.CountryInfo.Ccn3, param) {
			return &v, nil
		}
	}

	return nil, nil
}

//Latest get latest / last data from array
func (h *HistoryData) Latest() Data {
	return h.Data[len(h.Data)-1]
}

//Simplify simplify data to reduce size
func (h *HistoryData) Simplify() *SimpleHistoryData {
	var response SimpleHistoryData
	var data SimpleData

	data.Cases = make(map[string]float64)
	data.Deaths = make(map[string]float64)

	for _, v := range h.Data {
		date := int64(v.Date) / int64(time.Microsecond)
		t := time.Unix(date, 0)
		data.Deaths[t.Format(util.DATEFORMAT)] = v.TotalDeaths
		data.Cases[t.Format(util.DATEFORMAT)] = v.TotalCases
	}

	response.Data = data
	response.CountryInfo = h.CountryInfo
	response.Country = h.Country

	return &response
}

// NewService create new service
func NewService(h http.Client, list jsons.CountryList) Service {
	return service{h: h, countries: list}
}
