package worldometer

import (
	"math"
	"strconv"
	"strings"
)

//Convert Api Country format to country
func getCountry(t string) string {
	t = strings.Trim(t, " ")
	t = strings.TrimSuffix(t, "\n")
	t = strings.Replace(t, "\n", "", -1)
	switch t {
	case "S. Korea":
		t = "South Korea"
		break
	case "USA":
		t = "United States"
		break
	case "UK":
		t = "United Kingdom"
		break
	case "UAE":
		t = "United Arab Emirates"
		break
	case "DRC":
		t = "Democratic Republic of the Congo"
		break
	case "Congo":
		t = "Republic of the Congo"
		break
	case "Faeroe Islands":
		t = "Faroe Islands"
		break
	case "CAR":
		t = "Central African Republic"
		break
	case "Macao":
		t = "Macau"
		break
	case "Turks and Caicos":
		t = "Turks and Caicos Islands"
		break
	case "Cabo Verde":
		t = "Cape Verde"
		break
	case "St. Barth":
		t = "Saint Barthélemy"
		break
	case "St. Vincent Grenadines":
		t = "Saint Vincent and the Grenadines"
		break
	case "Sao Tome and Principe":
		t = "São Tomé and Príncipe"
		break
	case "Saint Pierre Miquelon":
		t = "Saint Pierre and Miquelon"
		break
	case "Total":
		t = ""
		break
	case "World":
		t = "all"
		break
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
	return math.Round(parsed*100)/100
}