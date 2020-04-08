package worldometer

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jezerdave/go-covid19/covid/util"
	"net/http"
)

// Service struct
type service struct {
	h http.Client
}

// Service
type Service interface {
	GetCountriesData() (*[]CountryStats, error)
	GetStatesData() (*[]StatesStats, error)
}

// NewService
func NewService(h http.Client) Service {
	return service{h: h}
}

// Fetch data from api
func (s service) GetCountriesData() (*[]CountryStats, error) {

	res, err := s.h.Get(countriesApi)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
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
				break
			case 1:
				holder.TotalCases = formatToInt(sTD.Text())
				break
			case 2:
				holder.NewCases = formatToInt(sTD.Text())
				break
			case 3:
				holder.TotalDeaths = formatToInt(sTD.Text())
				break
			case 4:
				holder.NewDeaths = formatToInt(sTD.Text())
				break
			case 5:
				holder.TotalRecovered = formatToInt(sTD.Text())
				break
			case 6:
				holder.ActiveCases = formatToInt(sTD.Text())
				break
			case 7:
				holder.SeriousCritical = formatToInt(sTD.Text())
				break
			case 8:
				holder.CasesPerOneMillion = formatToFloat(sTD.Text())
				break
			case 9:
				holder.DeathsPerOneMillion = formatToFloat(sTD.Text())
				break
			case 10:
				holder.TotalTests = formatToInt(sTD.Text())
				break
			case 11:
				holder.TestsPerOneMillion = formatToFloat(sTD.Text())
				break
			}
		})
		if holder.Country != "" {
			data = append(data, holder)
		}
	})

	return &data, nil
}

// GetStatesData
func (s service) GetStatesData() (*[]StatesStats, error) {
	res, err := s.h.Get(statesApi)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
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
				break
			case 1:
				holder.TotalCases = formatToInt(sTD.Text())
				break
			case 2:
				holder.NewCases = formatToInt(sTD.Text())
				break
			case 3:
				holder.TotalDeaths = formatToInt(sTD.Text())
				break
			case 4:
				holder.NewDeaths = formatToInt(sTD.Text())
				break
			case 5:
				holder.TotalRecovered = formatToInt(sTD.Text())
				break
			case 6:
				holder.ActiveCases = formatToInt(sTD.Text())
				break
			case 7:
				holder.SeriousCritical = formatToInt(sTD.Text())
				break
			case 8:
				holder.CasesPerOneMillion = formatToFloat(sTD.Text())
				break
			case 9:
				holder.DeathsPerOneMillion = formatToFloat(sTD.Text())
				break
			case 10:
				holder.TotalTests = formatToInt(sTD.Text())
				break
			case 11:
				holder.TestsPerOneMillion = formatToFloat(sTD.Text())
				break
			}
		})
		if holder.State != "" && !util.InArray(holder.State, exception) {
			data = append(data, holder)
		}
	})

	return &data, nil
}
