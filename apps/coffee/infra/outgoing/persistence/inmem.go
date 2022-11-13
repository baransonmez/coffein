package persistence

import (
	"context"
	"errors"
	"github.com/baransonmez/coffein/internal/coffee/infra/outgoing"
	"sync"
)

type InMem struct {
	store map[string]*outgoing.Bean
	m     sync.Mutex
}

func NewInMem() *InMem {
	var emptyMap = map[string]*outgoing.Bean{}
	return &InMem{
		store: emptyMap,
	}
}

func (i *InMem) Create(_ context.Context, recipe *outgoing.Bean) error {
	i.m.Lock()
	defer i.m.Unlock()
	i.store[recipe.ID] = recipe
	return nil
}

func (i *InMem) Get(id string) (*outgoing.Bean, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return i.store[id], nil
}
