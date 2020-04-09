package worldometer

const countriesApi = "https://www.worldometers.info/coronavirus/"
const statesApi = "https://www.worldometers.info/coronavirus/country/us/"

//CountryStats
type CountryStats struct {
	Country             string      `json:"country"`
	CountryInfo         interface{} `json:"country_info"`
	TotalCases          int         `json:"total_cases"`
	NewCases            int         `json:"new_cases"`
	TotalDeaths         int         `json:"total_deaths"`
	NewDeaths           int         `json:"new_deaths"`
	TotalRecovered      int         `json:"total_recovered"`
	ActiveCases         int         `json:"active_cases"`
	SeriousCritical     int         `json:"serious_critical"`
	CasesPerOneMillion  float64     `json:"cases_per_one_million"`
	DeathsPerOneMillion float64     `json:"deaths_per_one_million"`
	TotalTests          int         `json:"total_tests"`
	TestsPerOneMillion  float64     `json:"tests_per_one_million"`
}

//StatesStats
type StatesStats struct {
	State               string      `json:"state"`
	StateInfo           interface{} `json:"state_info"`
	TotalCases          int         `json:"total_cases"`
	NewCases            int         `json:"new_cases"`
	TotalDeaths         int         `json:"total_deaths"`
	NewDeaths           int         `json:"new_deaths"`
	TotalRecovered      int         `json:"total_recovered"`
	ActiveCases         int         `json:"active_cases"`
	SeriousCritical     int         `json:"serious_critical"`
	CasesPerOneMillion  float64     `json:"cases_per_one_million"`
	DeathsPerOneMillion float64     `json:"deaths_per_one_million"`
	TotalTests          int         `json:"total_tests"`
	TestsPerOneMillion  float64     `json:"tests_per_one_million"`
}
