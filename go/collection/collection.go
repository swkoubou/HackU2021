package collection

import (
	"errors"
	"reflect"

	"example.com/account"
	"example.com/manage"
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
func (c *Collection) IsEmpty(collection *Collection) bool { //Collectionに空の要素があったらTrueを返す
	var id_null uuid.NullUUID
	return c.CollectionID == id_null.UUID || len(c.CollectionName) == 0 ||
		len(c.CollectionDescripition) == 0 || c.Author.UserID == id_null.UUID ||
		len(c.Author.UserName) == 0 || len(c.Questions) == 0 ||
		len(c.UpdateTime) == 0 || len(c.CreateTime) == 0

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
	var user Collection
	//データベース開く
	db, err := manage.NewDBConnection()
	if err != nil {
		return nil, err
	}
	//クエリ
	defer db.Close()
	rows1, err := db.Query("SELECT * FROM question_collection WHERE collection_id = ?", collectionID)
	//defer rows1.Close()
	if err != nil {
		return nil, err
	}
	var collection_id_tmp, collection_user_id_tmp string
	for rows1.Next() {
		//question_collectionテーブルのデータを構造体に代入
		if err := rows1.Scan(&collection_id_tmp, &user.CollectionName, &user.CollectionDescripition,
			&collection_user_id_tmp, &user.CreateTime, &user.UpdateTime); err != nil {
			return nil, err
		}
		user.CollectionID, _ = uuid.Parse(collection_id_tmp)
		user.Author.UserID, _ = uuid.Parse(collection_user_id_tmp)
	}
	rows2, err := db.Query("SELECT  user_id FROM user WHERE user_id = ?", collection_user_id_tmp)
	if err != nil {
		return nil, err
	}
	for rows2.Next() {
		if err := rows2.Scan(&user.Author.UserName); err != nil {
			return nil, err
		}
	}
	//ここからQuestions取得
	rows3, err := db.Query("SELECT question_id FROM question_collection_map WHERE collection_id = ?", collection_user_id_tmp)
	if err != nil {
		return nil, err
	}
	var idtemp string
	for rows3.Next() {
		if err := rows3.Scan(&idtemp); err != nil {
			return nil, err
		}
		question_tmp, qerr := question.GetQuestion(idtemp)
		if qerr != nil {
			return nil, err
		}
		user.Questions = append(user.Questions, *question_tmp)
		if user.IsEmpty(&user) { // エラーになる場合は，戻り値としてerrorをtrueとすることで，uuidによる00初期化を防止．．あしあし
			var isEmpty error
			isEmpty = errors.New("true")
			return nil, isEmpty
		}
	}

	return &user, nil
}
