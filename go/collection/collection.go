package collection

import (
	"example.com/account"
	"example.com/question"
	"github.com/google/uuid"
)

type Collection struct {
	CollectionID           uuid.UUID           `json:"collectionID"`
	CollectionName         string              `json:"collectionName"`
	CollectionDescripition string              `json:"collectionDescription"`
	Auther                 account.Account     `json:"auther"`
	Questions              []question.Question `json:"questions"`
	CreateTime             string              `json:"createtime"`
	UpdateTime             string              `json:"updatetime"`
}

type CollectionParam struct {
	CollectionName         string
	CollectionDescripition string
	UserID                 string
	Questions              []string
}

func (c *Collection) ID() string {
	return c.CollectionID.String()
}

func NewCollection(param CollectionParam) (Collection, error) {
	c := Collection{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return c, err
	}

	c.CollectionID = uuid
	c.CollectionName = param.CollectionName
	c.CollectionDescripition = param.CollectionDescripition
	c.Auther = account.Account{}        // DBからもってくる
	c.Questions = []question.Question{} // DBからもってくる

	return c, nil
}
