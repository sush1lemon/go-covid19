package storage

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

type storage struct {
	kV *redis.Client
}

func (s storage) Find(query string) (*[]string, error) {
	keys, err := s.kV.Keys(fmt.Sprintf("countries:*%s:*", query)).Result()
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

func (s storage) New(key string, data string) error {
	return s.kV.Set(key, data, 0).Err()
}

func (s storage) GetAll() (*[]string, error) {
	keys, err := s.kV.Keys("countries:*").Result()
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

type Storage interface {
	Find(key string) (*[]string, error)
	GetAll() (*[]string, error)
}

func NewStorage(kv *redis.Client) Storage  {
	return &storage{kV:kv}
}