package worldometer

import (
	"math"
	"strconv"
	"strings"
)

// Convert Api Country format to country
func getCountry(t string) string {
	t = strings.Trim(t, " ")
	t = strings.TrimSuffix(t, "\n")
	t = strings.Replace(t, "\n", "", -1)
	switch t {
	case "S. Korea":
		t = "South Korea"
	case "USA":
		t = "United States"
	case "UK":
		t = "United Kingdom"
	case "UAE":
		t = "United Arab Emirates"
	case "DRC":
		t = "Democratic Republic of the Congo"
	case "Congo":
		t = "Republic of the Congo"
	case "Faeroe Islands":
		t = "Faroe Islands"
	case "CAR":
		t = "Central African Republic"
	case "Macao":
		t = "Macau"
	case "Turks and Caicos":
		t = "Turks and Caicos Islands"
	case "Cabo Verde":
		t = "Cape Verde"
	case "St. Barth":
		t = "Saint Barthélemy"
	case "St. Vincent Grenadines":
		t = "Saint Vincent and the Grenadines"
	case "Sao Tome and Principe":
		t = "São Tomé and Príncipe"
	case "Saint Pierre Miquelon":
		t = "Saint Pierre and Miquelon"
	case "Total":
		t = ""
	case "World":
		t = "all"
	}
	return t
}

// Format string to int
func formatToInt(value string) int {
	value = strings.TrimSpace(value)
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, "+", "", -1)
	var def int
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return parsed
}

// Format string to float64
func formatToFloat(value string) float64 {
	value = strings.TrimSpace(value)
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, "+", "", -1)
	var def float64
	parsed, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return def
	}
	return math.Round(parsed*100) / 100
}
