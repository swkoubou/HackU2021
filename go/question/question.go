package question

import (
	"reflect"

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

type QuestionParam struct {
	UserID       string
	QuestionTag  []string
	QuestionType string
	QuestionBody string
	Values       []string
	Answers      []string
}

func (q *Question) ID() string {
	return q.QuestionID.String()
}

func NewQuestion(param QuestionParam) (Question, error) {
	q := Question{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return q, err
	}

	q.QuestionID = uuid
	q.Auther = account.Account{} // DBからもってくる
	q.QuestionTag = param.QuestionTag
	q.QuestionBody = param.QuestionBody
	q.Values = param.Values
	q.Answers = param.Answers

	return q, nil
}

func (q *Question) Equals(question *Question) bool {
	if reflect.TypeOf(*q) != reflect.TypeOf(*question) {
		return false
	}

	if q.QuestionID != question.QuestionID {
		return false
	}

	if !reflect.DeepEqual(q.Auther, question.Auther) {
		return false
	}

	if len(q.QuestionTag) != len(q.QuestionTag) {
		return false
	}

	for i, v := range q.QuestionTag {
		if v != question.QuestionTag[i] {
			return false
		}
	}

	if len(q.CreateTime) != len(question.CreateTime) {
		return false
	}

	if len(q.UpdateTime) != len(question.UpdateTime) {
		return false
	}

	if q.QuestionBody != question.QuestionBody {
		return false
	}

	for i, v := range q.Values {
		if v != question.Values[i] {
			return false
		}
	}

	for i, v := range q.Answers {
		if v != question.Answers[i] {
			return false
		}
	}

	return true
}
