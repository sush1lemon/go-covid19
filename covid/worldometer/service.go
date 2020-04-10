package worldometer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jezerdave/go-covid19/covid/util"
	"github.com/jezerdave/go-covid19/covid/util/jsons"
	"net/http"
	"strings"
)

// Service struct
type service struct {
	h         http.Client
	countries jsons.CountryList
	states    jsons.States
}

//Service interface
type Service interface {
	GetCountriesData() (*[]CountryStats, error)
	GetStatesData() (*[]StatesStats, error)
	FindCountry(string) (*CountryStats, error)
	FindState(string) (*StatesStats, error)
}

//NewService create new service
func NewService(h http.Client, countries jsons.CountryList, states jsons.States) Service {
	return service{h, countries, states}
}

//GetCountriesData get all countries data
func (s service) GetCountriesData() (*[]CountryStats, error) {

	res, err := s.h.Get(countriesApi)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var data []CountryStats

	table := doc.Find("#main_table_countries_today")
	table.Find("tr").Each(func(i int, sTR *goquery.Selection) {
		var holder CountryStats
		sTR.Find("td").Each(func(itd int, sTD *goquery.Selection) {
			switch itd {
			case 0:
				holder.Country = getCountry(sTD.Text())
			case 1:
				holder.TotalCases = formatToInt(sTD.Text())
			case 2:
				holder.NewCases = formatToInt(sTD.Text())
			case 3:
				holder.TotalDeaths = formatToInt(sTD.Text())
			case 4:
				holder.NewDeaths = formatToInt(sTD.Text())
			case 5:
				holder.TotalRecovered = formatToInt(sTD.Text())
			case 6:
				holder.ActiveCases = formatToInt(sTD.Text())
			case 7:
				holder.SeriousCritical = formatToInt(sTD.Text())
			case 8:
				holder.CasesPerOneMillion = formatToFloat(sTD.Text())
			case 9:
				holder.DeathsPerOneMillion = formatToFloat(sTD.Text())
			case 10:
				holder.TotalTests = formatToInt(sTD.Text())
			case 11:
				holder.TestsPerOneMillion = formatToFloat(sTD.Text())
			}
		})
		if holder.Country != "" {

			for _, v := range s.countries.Countries {
				if strings.EqualFold(holder.Country, v.Name.Common) || strings.EqualFold(holder.Country, v.Name.Official) ||
					strings.EqualFold(holder.Country, v.Cca3) || strings.EqualFold(holder.Country, v.Cca2) ||
					strings.EqualFold(holder.Country, v.Ccn3) {
					holder.CountryInfo = v
					data = append(data, holder)
				}
			}
		}
	})

	return &data, nil
}

//GetStatesData get all states data
func (s service) GetStatesData() (*[]StatesStats, error) {
	res, err := s.h.Get(statesApi)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var data []StatesStats
	exception := []string{"Wuhan Repatriated", "Diamond Princess Cruise", "Total:", "USA Total"}

	table := doc.Find("#usa_table_countries_today")
	table.Find("tr").Each(func(i int, sTR *goquery.Selection) {
		var holder StatesStats
		sTR.Find("td").Each(func(itd int, sTD *goquery.Selection) {
			switch itd {
			case 0:
				holder.State = getCountry(sTD.Text())
			case 1:
				holder.TotalCases = formatToInt(sTD.Text())
			case 2:
				holder.NewCases = formatToInt(sTD.Text())
			case 3:
				holder.TotalDeaths = formatToInt(sTD.Text())
			case 4:
				holder.NewDeaths = formatToInt(sTD.Text())
			case 5:
				holder.TotalRecovered = formatToInt(sTD.Text())
			case 6:
				holder.ActiveCases = formatToInt(sTD.Text())
			case 7:
				holder.SeriousCritical = formatToInt(sTD.Text())
			case 8:
				holder.CasesPerOneMillion = formatToFloat(sTD.Text())
			case 9:
				holder.DeathsPerOneMillion = formatToFloat(sTD.Text())
			case 10:
				holder.TotalTests = formatToInt(sTD.Text())
			case 11:
				holder.TestsPerOneMillion = formatToFloat(sTD.Text())
			}
		})

		if holder.State != "" && !util.InArray(holder.State, exception) {
			for i, v := range s.states.States {
				if strings.EqualFold(holder.State, v) {
					holder.StateInfo = StateInfo{
						Name:         v,
						Abbreviation: i,
					}
					data = append(data, holder)
				}
			}
		}
	})

	return &data, nil
}

//FindCountry get country by given string
func (s service) FindCountry(param string) (*CountryStats, error) {
	list, err := s.GetCountriesData()
	if err != nil {
		return nil, err
	}

	for _, v := range *list {
		if strings.EqualFold(param, v.CountryInfo.Name.Common) || strings.EqualFold(param, v.CountryInfo.Name.Official) ||
			strings.EqualFold(param, v.CountryInfo.Cca3) || strings.EqualFold(param, v.CountryInfo.Cca2) ||
			strings.EqualFold(param, v.CountryInfo.Ccn3) {
			return &v, nil
		}
	}

	return nil, nil
}

//FindState get state by given string
func (s service) FindState(param string) (*StatesStats, error) {
	list, err := s.GetStatesData()
	if err != nil {
		return nil, err
	}

	for _, v := range *list {
		if strings.EqualFold(v.StateInfo.Name, param) || strings.EqualFold(v.State, param) {
			return &v, nil
		}
	}

	return nil, nil
}
