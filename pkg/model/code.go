package model

import "sync"

type CodeSet struct {
	m map[string]bool
	sync.RWMutex
}

func NewCodeSet() *CodeSet {
	return &CodeSet{
		m: map[string]bool{},
	}
}

func (s *CodeSet) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *CodeSet) AddSet(items []string) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *CodeSet) Remove(item string) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *CodeSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *CodeSet) Len() int {
	return len(s.List())
}

func (s *CodeSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *CodeSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *CodeSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
