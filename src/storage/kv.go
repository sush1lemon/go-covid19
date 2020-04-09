package storage

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
)

// storage
type storage struct {
	kV *redis.Client
}

//Find
func (s storage) Find(query string) (*[]string, error) {
	keys, err := s.kV.Keys(query).Result()
	if err != nil {
		return nil, err
	}

	var result []string

	for _, v := range keys {
		data, err := s.kV.Get(v).Result()
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return &result, nil
}

//New
func (s storage) New(key string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.kV.Set(key, string(b), 0).Err()
}

//GetAll
func (s storage) GetAll(param string) (*[]string, error) {
	keys, err := s.kV.Keys(param).Result()
	if err != nil {
		return nil, err
	}

	var result []string

	for _, v := range keys {
		data, err := s.kV.Get(v).Result()
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return &result, nil
}

//Get
func (s storage) Get(param string) (*string, error) {
	data, err := s.kV.Get(param).Result()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//Storage
type Storage interface {
	Find(key string) (*[]string, error)
	GetAll(param string) (*[]string, error)
	New(key string, data interface{}) error
	Get(param string) (*string, error)
}

//NewStorage
func NewStorage(kv *redis.Client) Storage {
	return &storage{kV: kv}
}
