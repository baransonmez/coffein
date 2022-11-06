package outgoing

import (
	"context"
)

type UserAdapter struct {
}

func NewUserAdapter() (UserAdapter, error) {

	return UserAdapter{}, nil
}

func (u UserAdapter) GetMembershipType(ctx context.Context, userId string) (string, error) {

	return "membershipType", nil
}
