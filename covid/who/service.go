package who

import (
	"encoding/json"
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"io/ioutil"
	"net/http"
	"strings"
)

// Service interface
type Service interface {
	History() (*[]HistoryData, error)
}

// Service struct
type service struct {
	h     http.Client
	cList jsons.CountryList
}

// Get Historical covid19 data from WHO
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
		for _, x := range s.cList.Countries {
			if strings.EqualFold(x.Cca2, v.Country) || strings.EqualFold(x.Cca3, v.Country) {
				v.Country = x.Name.Common
				v.CountryInfo = x
				res = append(res, v)
			}
		}
	}

	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	return &res, nil
}

// NewService
func NewService(h http.Client, list jsons.CountryList) Service {
	return service{h: h, cList: list}
}
