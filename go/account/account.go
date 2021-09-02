package account

import (
	"github.com/google/uuid"
)

type Account struct {
	UserID   uuid.UUID `json:"userID"`
	UserName string    `json:"userName"`
}

func (a *Account) ID() string {
	return a.UserID.String()
}

func NewAccount(userName string) (Account, error) {
	a := Account{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return a, err
	}

	a.UserID = uuid
	a.UserName = userName

	return a, nil
}
