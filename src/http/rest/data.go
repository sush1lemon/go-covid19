package rest

type CountryList struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Name CountryName     `json:"name"`
	TLD  []string `json:"tld"`
	Cca2 string   `json:"cca2"`
	Ccn3 string   `json:"ccn3"`
	Cca3 string   `json:"cca3"`
	Cioc string   `json:"cioc"`
}

type CountryName struct {
	Common   string            `json:"common"`
	Official string            `json:"official"`
}

type States struct {
	States map[string]string `json:"states"`
}

type BasicError struct {
	Message string `json:"message"`
}