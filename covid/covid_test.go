package covid

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jezerdave/go-covid19/covid/util"
	"io/ioutil"
	"strings"
	"testing"
)

type Countries struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Name Name     `json:"name"`
	TLD  []string `json:"tld"`
	Cca2 string   `json:"cca2"`
	Ccn3 string   `json:"ccn3"`
	Cca3 string   `json:"cca3"`
	Cioc string   `json:"cioc"`
}

type Name struct {
	Common   string            `json:"common"`
	Official string            `json:"official"`
}

func TestScrape(t *testing.T) {

	srv := NewClient()

	data, err := srv.Worldometer.GetCountriesData()
	if err != nil {
		t.Fatal(err)
	}

	byteValue, err := ioutil.ReadFile("./countries.json")
	if err != nil {
		fmt.Print(err)
	}

	kV := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	var result Countries
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("[TOTAL] ",len(*data))
	var test []string
	for _, v := range *data {
		for _, c := range result.Countries {
			if v.Country == c.Name.Official || v.Country == c.Name.Common {
				key := fmt.Sprintf("countries::%s:%s:%s:%s", strings.ToLower(c.Name.Common),
					strings.ToLower(c.Ccn3), strings.ToLower(c.Cca2), strings.ToLower(c.Cca3))
				b, err := json.Marshal(v)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(key)
				err = kV.Set(key, string(b), 0).Err()
				if err != nil {
					panic(err)
				}
				test = append(test, v.Country)
			}
		}

	}

	for _, v := range *data {
		if !util.InArray(v.Country, test) {
			fmt.Println(v.Country)
		}
	}

	fmt.Println(len(test))

}

func TestPH(t *testing.T)  {
	srv := NewClient()

	_, err := srv.Philippines.GetStats()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStates(t *testing.T) {
	srv := NewClient()

	data, err := srv.Worldometer.GetStatesData()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range *data {
		fmt.Println(v)
	}

}

