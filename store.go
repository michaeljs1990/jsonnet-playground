package main

import (
	"errors"
	"sync"
)

var (
	DatastoreNoSuchID = errors.New("The ID you request to lookup doesn't exist in our backend store")
)

type PersistJsonnet interface {
	Get(id string) (string, error)
	Store(id string, code string) error
}

// InMemory is for testing or a quick deploy
type InMemory struct {
	sync.RWMutex
	store map[string]string
}

func (i InMemory) Get(id string) (string, error) {
	i.RLock()
	defer i.RUnlock()
	val, ok := i.store[id]
	if ok {
		return val, nil
	}
	return val, DatastoreNoSuchID
}

func (i InMemory) Store(id string, code string) error {
	i.Lock()
	i.store[id] = code
	i.Unlock()
	return nil
}
