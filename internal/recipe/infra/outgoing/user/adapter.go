package user

import (
	"context"
)

type Adapter struct {
}

func NewUserAdapter() (Adapter, error) {

	return Adapter{}, nil
}

func (u Adapter) GetMembershipType(ctx context.Context, userId string) (string, error) {

	return "membershipType", nil
}
