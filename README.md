# Go-Covid19
**GO-COVID-19** is just another API 
######(but written in golang)

## Installation
#### Package
```
go get github.com/jezerdave/go-covid19/covid
```

#### Import
```
import "github.com/jezerdave/go-covid19/covid"
```
#### Sample
```
client := covid.NewClient()
data, err := client.Worldometer.GetCountriesData()
```

#### Available Packages
[Worldometer](https://github.com/jezerdave/go-covid19/blob/master/covid/worldometer/service.go#L17) \
[Philippines](https://github.com/jezerdave/go-covid19/blob/master/covid/philippines/service.go#L15)

## REST API
https://go-covid19.herokuapp.com/
#### DOCS **(WIP)**


##Data Source
https://www.worldometers.info/coronavirus/
https://www.worldometers.info/coronavirus/country/us/
https://ncovtracker.doh.gov.ph/