package collection

import (
	"reflect"

	"example.com/account"
	"example.com/question"
	"github.com/google/uuid"
)

type Collection struct {
	CollectionID           uuid.UUID           `json:"collectionID"`
	CollectionName         string              `json:"collectionName"`
	CollectionDescripition string              `json:"collectionDescription"`
	Author                 account.Account     `json:"author"`
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
	c.Author = account.Account{}        // DBからもってくる
	c.Questions = []question.Question{} // DBからもってくる

	return c, nil
}

func (c *Collection) Equals(collection *Collection) bool {
	if reflect.TypeOf(*c) != reflect.TypeOf(*collection) {
		return false
	}

	if c.CollectionID != collection.CollectionID {
		return false
	}

	if c.CollectionName != collection.CollectionName {
		return false
	}

	if c.CollectionDescripition != collection.CollectionDescripition {
		return false
	}

	if !reflect.DeepEqual(c.Author, collection.Author) {
		return false
	}

	for i, v := range c.Questions {
		if !v.Equals(&collection.Questions[i]) {
			return false
		}
	}

	if len(c.CreateTime) != len(collection.CreateTime) {
		return false
	}

	if len(c.UpdateTime) != len(collection.UpdateTime) {
		return false
	}

	return true
}

func GetCollection(collectionID string) (*Collection, error) {
	return &Collection{}, nil
}
