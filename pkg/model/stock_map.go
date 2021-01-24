package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

type StockMap struct {
	m            map[string]Stock
	StockAddChan chan Stock
	sync.RWMutex
}

func NewStockMap() *StockMap {
	return &StockMap{
		m:            map[string]Stock{},
		StockAddChan: make(chan Stock, 100),
	}
}

func (s *StockMap) Add(code string, stock Stock) {
	s.Lock()
	defer s.Unlock()
	s.m[code] = stock
	s.StockAddChan <- stock
}

func (s *StockMap) Get(code string) (Stock, bool) {
	s.RLock()
	defer s.RUnlock()
	stock, ok := s.m[code]
	return stock, ok
}

func (s *StockMap) Remove(code string) {
	s.Lock()
	s.Unlock()
	delete(s.m, code)
}

func (s *StockMap) Has(code string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[code]
	return ok
}

func (s *StockMap) Len() int {
	return len(s.List())
}

func (s *StockMap) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]Stock{}
}

func (s *StockMap) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *StockMap) List() []Stock {
	s.RLock()
	defer s.RUnlock()
	list := []Stock{}
	for _, item := range s.m {
		list = append(list, item)
	}
	return list
}

func (sm *StockMap) String() string {
	b, err := json.Marshal(sm)
	if err != nil {
		return fmt.Sprint("Error Marshal")
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprint("Error Parse")
	}
	return out.String()
}
