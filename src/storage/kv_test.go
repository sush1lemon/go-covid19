package storage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"testing"
)

func Test_storage_Find(t *testing.T) {
	kV := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	db := NewStorage(kV)
	fmt.Println(db.Find("oman"))
}