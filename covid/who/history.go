package who

import "github.com/jezerdave/go-covid19/covid/util/jsons"

const api = "https://dashboards-dev.sprinklr.com/data/9043/global-covid19-who-gis.json"

//HistoryData
type HistoryData struct {
	Country     string        `json:"country"`
	CountryInfo jsons.Country `json:"country_info"`
	Data        []Data        `json:"data"`
}

//Data
type Data struct {
	Date               float64 `json:"date"`
	Region             *string `json:"region"`
	DeathDailyIncrease float64 `json:"death_daily_increase"`
	TotalDeaths        float64 `json:"total_deaths"`
	CaseDailyIncrease  float64 `json:"case_daily_increase"`
	TotalCases         float64 `json:"total_cases"`
}

type rawData struct {
	Rows [][]interface{} `json:"rows"`
}
