package worldometer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jezerdave/go-covid19/covid/util"
	"net/http"
)

// Service struct
type service struct {
	h http.Client
}

//Service
type Service interface {
	GetCountriesData() (*[]CountryStats, error)
	GetStatesData() (*[]StatesStats, error)
}

//NewService
func NewService(h http.Client) Service {
	return service{h: h}
}

//GetCountriesData
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
			data = append(data, holder)
		}
	})

	return &data, nil
}

//GetStatesData
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
			data = append(data, holder)
		}
	})

	return &data, nil
}
