package account

import "github.com/google/uuid"

type Account struct {
	UserID   uuid.UUID `json:"userID"`
	UserName string    `json:"userName"`
}
