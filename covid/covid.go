package covid

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const api = "https://www.worldometers.info/coronavirus/"

// Stats from api
type Stats struct {
	Country             string  `json:"country"`
	TotalCases          int     `json:"total_cases"`
	NewCases            int     `json:"new_cases"`
	TotalDeaths         int     `json:"total_deaths"`
	NewDeaths           int     `json:"new_deaths"`
	TotalRecovered      int     `json:"total_recovered"`
	ActiveCases         int     `json:"active_cases"`
	SeriousCritical     int     `json:"serious_critical"`
	CasesPerOneMillion  float64 `json:"cases_per_one_million""`
	DeathsPerOneMillion float64 `json:"deaths_per_one_million""`
	TotalTests          int     `json:"total_tests"`
	TestsPerOneMillion  float64 `json:"tests_per_one_million""`
}

// Fetch data from api
func FetchData() (*[]Stats, error) {

	res, err := http.Get(api)
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

	var data []Stats

	table := doc.Find("#main_table_countries_today")
	table.Find("tr").Each(func(i int, sTR *goquery.Selection) {
		var holder Stats
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
