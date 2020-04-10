package worldometer

import "github.com/jezerdave/go-covid19/covid/util/jsons"

const countriesApi = "https://www.worldometers.info/coronavirus/"
const statesApi = "https://www.worldometers.info/coronavirus/country/us/"

//CountryStats struct
type CountryStats struct {
	Country             string        `json:"country"`
	CountryInfo         jsons.Country `json:"country_info"`
	TotalCases          int           `json:"total_cases"`
	NewCases            int           `json:"new_cases"`
	TotalDeaths         int           `json:"total_deaths"`
	NewDeaths           int           `json:"new_deaths"`
	TotalRecovered      int           `json:"total_recovered"`
	ActiveCases         int           `json:"active_cases"`
	SeriousCritical     int           `json:"serious_critical"`
	CasesPerOneMillion  float64       `json:"cases_per_one_million"`
	DeathsPerOneMillion float64       `json:"deaths_per_one_million"`
	TotalTests          int           `json:"total_tests"`
	TestsPerOneMillion  float64       `json:"tests_per_one_million"`
}

//StatesStats struct
type StatesStats struct {
	State               string    `json:"state"`
	StateInfo           StateInfo `json:"state_info"`
	TotalCases          int       `json:"total_cases"`
	NewCases            int       `json:"new_cases"`
	TotalDeaths         int       `json:"total_deaths"`
	NewDeaths           int       `json:"new_deaths"`
	TotalRecovered      int       `json:"total_recovered"`
	ActiveCases         int       `json:"active_cases"`
	SeriousCritical     int       `json:"serious_critical"`
	CasesPerOneMillion  float64   `json:"cases_per_one_million"`
	DeathsPerOneMillion float64   `json:"deaths_per_one_million"`
	TotalTests          int       `json:"total_tests"`
	TestsPerOneMillion  float64   `json:"tests_per_one_million"`
}

//StateInfo struct
type StateInfo struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}
