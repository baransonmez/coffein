package persistence

import (
	"context"
	"errors"
	"github.com/baransonmez/coffein/internal/recipe/infra/outgoing"
	"sync"
)

type InMem struct {
	store map[string]*outgoing.Recipe
	m     sync.Mutex
}

func NewInMem() *InMem {
	var emptyMap = map[string]*outgoing.Recipe{}
	return &InMem{
		store: emptyMap,
	}
}

func (i *InMem) Create(_ context.Context, recipe outgoing.Recipe) error {
	i.m.Lock()
	defer i.m.Unlock()
	i.store[recipe.ID] = &recipe
	return nil
}

func (i *InMem) Get(id string) (*outgoing.Recipe, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return i.store[id], nil
}
