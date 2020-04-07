package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jezerdave/go-covid19/covid"
	"github.com/jezerdave/go-covid19/src/http/rest"
	"github.com/jezerdave/go-covid19/src/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	port      = "1919"
	redisHost = "127.0.0.1"
	redisPass = ""
	redisPort = "6379"
)

func main() {

	countries, states, err := getJsons()
	if err != nil {
		log.Fatal(err)
	}

	kV := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPass,
		DB:       0,
	})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	srv := covid.NewClient()
	repo := storage.NewStorage(kV)
	routes := rest.Handler{
		PH:          srv.Philippines,
		WOM:         srv.Worldometer,
		R:           repo,
		CountryList: *countries,
		States:      *states,
	}

	v := e.Group("api/v1")


	cG := v.Group("/countries")
	cG.GET("", routes.GetCountriesData)
	cG.GET("/", routes.GetCountriesData)
	cG.GET("/:country", routes.FindCountry)

	sG := v.Group("/states")
	sG.GET("", routes.GetStatesData)
	sG.GET("/", routes.GetStatesData)
	sG.GET("/:param", routes.FindStates)

	phDoh := v.Group("/doh/ph")
	phDoh.GET("", routes.GetPHStats)
	phDoh.GET("/", routes.GetPHStats)
	phDoh.GET("/hospital-pui", routes.GetPHHospitalPUIs)

	v.GET("/update", routes.UpdateData)


	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	kV.Close()
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func getJsons() (*rest.CountryList, *rest.States, error) {

	var countryList rest.CountryList
	var stateList rest.States

	clByte, err := ioutil.ReadFile("./src/storage/countries.json")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(clByte, &countryList)
	if err != nil {
		return nil, nil, err
	}

	stByte, err := ioutil.ReadFile("./src/storage/states.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(stByte, &stateList)
	if err != nil {
		return nil, nil, err
	}

	return &countryList, &stateList, nil
}

func getConfig() {
	port = os.Getenv("APP_PORT")
	redisHost = os.Getenv("REDIS_HOST")
	redisPass = os.Getenv("REDIS_PASS")
	redisPort = os.Getenv("REDIS_PORT")
}
