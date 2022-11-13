package persistence

import (
	"context"
	"errors"
	"github.com/baransonmez/coffein/internal/user/infra/outgoing"
	"sync"
)

type InMem struct {
	store map[string]*outgoing.User
	m     sync.Mutex
}

func NewInMem() *InMem {
	var emptyMap = map[string]*outgoing.User{}
	return &InMem{
		store: emptyMap,
	}
}

func (i *InMem) Create(_ context.Context, user *outgoing.User) error {
	i.m.Lock()
	defer i.m.Unlock()
	i.store[user.ID] = user
	return nil
}

func (i *InMem) Get(id string) (*outgoing.User, error) {
	if i.store[id] == nil {
		return nil, errors.New("not found")
	}
	return i.store[id], nil
}
