package question

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"name"`
	UserName string    `json:"userName"`
}
type Question struct {
	QuestionTag  []string  `json:"questionTag"`
	QuestionID   uuid.UUID `json:"questionID"`
	Auther       User      `json:"auther"`
	CreateTime   string    `json:"createTime"`
	UpdateTime   string    `json:"updateTime"`
	QuestionBody []string  `json:"questionBody"`
	Values       []string  `json:"values"`
	Answers      []string  `json:"answers"`
}
