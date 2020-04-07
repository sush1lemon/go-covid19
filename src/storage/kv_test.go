package storage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"net/http"
	"testing"
)

func Test_storage_Find(t *testing.T) {
	kV := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	db := NewStorage(kV)
	fmt.Println(db.Find("ph"))

	r, err := http.Get("https://services9.arcgis.com/N9p5hsImWXAccRNI/arcgis/rest/services/")
	if err != nil {
		t.Fatal(err)
	}

	if r.StatusCode != 200 {
		t.Fatal(fmt.Sprintf("status code error: %d %s", r.StatusCode, r.Status))
	}

}