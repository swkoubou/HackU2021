package question

import (
	"example.com/account"
	"github.com/google/uuid"
)

type Question struct {
	QuestionID   uuid.UUID       `json:"questionID"`
	Auther       account.Account `json:"auther"`
	QuestionTag  []string        `json:"questionTag"`
	QuestionType string          `json:"questionType"`
	CreateTime   string          `json:"createTime"`
	UpdateTime   string          `json:"updateTime"`
	QuestionBody string          `json:"questionBody"`
	Values       []string        `json:"values"`
	Answers      []string        `json:"answers"`
}
