package question

import (
	"errors"
	"reflect"

	"example.com/account"
	"example.com/manage"
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

	if len(q.QuestionTag) != len(question.QuestionTag) {
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

func GetQuestion(questionID string) (*Question, error) {

	var user Question

	// データベースへのアクセス
	db, err := manage.NewDBConnection()

	// エラー処理
	if err != nil {
		return nil, err
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// クエリ
	rows1, err := db.Query("SELECT * FROM question WHERE question_id = ?", questionID)

	defer rows1.Close()

	// エラー処理
	if err != nil {
		return nil, err
	}

	var uuid_tmp1, uuid_tmp2 string

	var tmp string

	// ここからデータベースでの処理
	for rows1.Next() {
		if err := rows1.Scan(&uuid_tmp1, &tmp, &uuid_tmp2, &user.QuestionType, &user.CreateTime, &user.UpdateTime, &user.QuestionBody); err != nil {
			return nil, err
		}
		user.QuestionID, _ = uuid.Parse(uuid_tmp1)
		user.Auther.UserID, _ = uuid.Parse(uuid_tmp2)
	}

	rows2, err := db.Query("SELECT * FROM user WHERE user_id = ?", uuid_tmp2)

	if err != nil {
		return nil, err
	}

	var name_tmp string

	for rows2.Next() {
		if err := rows2.Scan(&name_tmp, &user.Auther.UserName); err != nil {
			return nil, err
		}
	}

	rows3, err := db.Query("SELECT * FROM question_tag_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	var questionid_tmp, array_tmp string

	var tag_int int

	for rows3.Next() {
		if err := rows3.Scan(&questionid_tmp, &tag_int); err != nil {
			return nil, err
		}

		rows4, err := db.Query("SELECT * FROM question_tag WHERE question_tag_id = ?", tag_int)
		if err != nil {
			return nil, err
		}

		for rows4.Next() {
			if err := rows4.Scan(&tag_int, &array_tmp); err != nil {
				return nil, err
			}
			user.QuestionTag = append(user.QuestionTag, array_tmp)
		}
	}

	rows5, err := db.Query("SELECT * FROM question_value_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	var int_tmp int

	for rows5.Next() {
		if err := rows5.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return nil, err
		}
		user.Values = append(user.Values, array_tmp)
	}

	rows6, err := db.Query("SELECT * FROM question_answer_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	for rows6.Next() {
		if err := rows6.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return nil, err
		}
		user.Answers = append(user.Answers, array_tmp)
	}

	// uuidのnil
	var u uuid.NullUUID

	if user.QuestionID == u.UUID || user.Auther.UserID == u.UUID || len(user.Auther.UserName) == 0 || len(user.QuestionTag) == 0 || len(user.QuestionType) == 0 || len(user.CreateTime) == 0 || len(user.UpdateTime) == 0 || len(user.QuestionBody) == 0 || len(user.Values) == 0 || len(user.Answers) == 0 {
		var isEmpty error

		// エラーになる場合は，戻り値としてerrorをtrueとすることで，uuidによる00初期化を防止．．あしあし
		isEmpty = errors.New("true")
		return nil, isEmpty
	}

	// user を question.Questionとして定義している.
	return &user, nil

	//return question.Question{}, nil
}

/*
引数: QuestionParam
type QuestionParam struct {
	UserID       string		必須
	QuestionTag  []string	0個以上
	QuestionType string		必須
	QuestionBody string		QuestionTypeがanaumeの場合"以下の空欄を埋めなさい"で固定
	Values       []string	必須
	Answers      []string	必須
}

戻り値: string, error
データの挿入が成功した場合はQuestionIDとnilを返す
データの挿入に失敗した場合は空文字とerrorを返す
*/
func SetQuestion(q *QuestionParam) (string, error) {
	return string(""), nil
}

func UpdateQuestion(q *Question) (string, error) {
	return string(""), nil
}

func DeleteQuestion(id string) error {
	return nil
}
