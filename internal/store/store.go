package store

import (
	"Server/internal/stores"
	"sync"
)

type store struct {
	mu sync.RWMutex
	m  map[string]any
}

func NewStore() stores.Store {
	return &store{}
}

func (s *store) Init() {
	s.m = make(map[string]any)
}

func (s *store) Get(key string) any {
	return 2
}

func (s *store) Set(key string, value any) {

}

func (s *store) SetWithTTL(key string, value any, ttl int) {

}

func (s *store) Delete(key string) {

}
