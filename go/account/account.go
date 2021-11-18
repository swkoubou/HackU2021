package account

import (
	"github.com/google/uuid"
)

type Account struct {
	UserID   string `json:"userID" db:"user_id"`
	UserName string `json:"userName" db:"user_name"`
}

func NewAccount(userName string) (Account, error) {
	a := Account{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return a, err
	}

	a.UserID = uuid.String()
	a.UserName = userName

	return a, nil
}
