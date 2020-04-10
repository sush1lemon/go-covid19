# Go-Covid19 [![Go Report Card](https://goreportcard.com/badge/github.com/jezerdave/go-covid19)](https://goreportcard.com/report/github.com/jezerdave/go-covid19)
**go-covid19** is a Go Package / Rest API for getting covid19 related data from different source


## REST API
#### LINK
https://go-covid19.sideprojects.fun/api/v1/
#### Documentation
https://go-covid19.sideprojects.fun/api/v1/docs/index.html


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

#### Available Methods
[Worldometer](https://github.com/jezerdave/go-covid19/blob/master/covid/worldometer/service.go#L17) \
[Philippines](https://github.com/jezerdave/go-covid19/blob/master/covid/philippines/service.go#L15) \
[WHO](https://github.com/jezerdave/go-covid19/blob/master/covid/who/service.go#12)

## Run on your machine
#### Clone repo
    git clone https://github.com/jezerdave/go-covid19
#### Run
    go run .
### Run with ENV
    REDIS_HOST=<value> REDIS_PASS=<value> REDIS_PORT=<value> go run .

## Data Source
https://www.worldometers.info/coronavirus/ \
https://www.worldometers.info/coronavirus/country/us/ \
https://ncovtracker.doh.gov.ph/ \
https://who.sprinklr.com/

## API Status
**Worldometer** - _available_ \
**WHO** - _available_ \
**DOH Philippines** - _[not available / maintenance](https://ncovtracker.doh.gov.ph/)_ 
