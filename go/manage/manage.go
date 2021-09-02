package manage

import (
	"database/sql"
	"errors"
	"os"

	"example.com/question"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type dsn struct {
	user     string
	pass     string
	protocol string
	address  string
	dbName   string
}

func NewDBConnection() (*sql.DB, error) {
	dsn := &dsn{
		user:     "root",
		pass:     os.Getenv("MYSQL_ROOT_PASSWORD"),
		protocol: "tcp",
		address:  "db:3306",
		dbName:   os.Getenv("MYSQL_DATABASE"),
	}

	return sql.Open("mysql", dsn.user+":"+dsn.pass+"@"+dsn.protocol+"("+dsn.address+")/"+dsn.dbName)
}

func GetQuestion(questionID string) (question.Question, error) {

	var user question.Question

	// データベースへのアクセス
	db, err := NewDBConnection()

	// エラー処理
	if err != nil {
		return user, err
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// クエリ
	rows1, err := db.Query("SELECT * FROM question WHERE question_id = ?", questionID)

	defer rows1.Close()

	// エラー処理
	if err != nil {
		return user, err
	}

	var uuid_tmp1, uuid_tmp2 string

	var tmp string

	// ここからデータベースでの処理
	for rows1.Next() {
		if err := rows1.Scan(&uuid_tmp1, &tmp, &uuid_tmp2, &user.QuestionType, &user.CreateTime, &user.UpdateTime, &user.QuestionBody); err != nil {
			return user, err
		}
		user.QuestionID, _ = uuid.Parse(uuid_tmp1)
		user.Auther.UserID, _ = uuid.Parse(uuid_tmp2)
	}

	rows2, err := db.Query("SELECT * FROM user WHERE user_id = ?", uuid_tmp2)

	if err != nil {
		return user, err
	}

	var name_tmp string

	for rows2.Next() {
		if err := rows2.Scan(&name_tmp, &user.Auther.UserName); err != nil {
			return user, err
		}
	}

	rows3, err := db.Query("SELECT * FROM question_tag_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return user, err
	}

	var questionid_tmp, array_tmp string

	var tag_int int

	for rows3.Next() {
		if err := rows3.Scan(&questionid_tmp, &tag_int); err != nil {
			return user, err
		}

		rows4, err := db.Query("SELECT * FROM question_tag WHERE question_tag_id = ?", tag_int)
		if err != nil {
			return user, err
		}

		for rows4.Next() {
			if err := rows4.Scan(&tag_int, &array_tmp); err != nil {
				return user, err
			}
			user.QuestionTag = append(user.QuestionTag, array_tmp)
		}
	}

	rows5, err := db.Query("SELECT * FROM question_value_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return user, err
	}

	var int_tmp int

	for rows5.Next() {
		if err := rows5.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return user, err
		}
		user.Values = append(user.Values, array_tmp)
	}

	rows6, err := db.Query("SELECT * FROM question_answer_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return user, err
	}

	for rows6.Next() {
		if err := rows6.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return user, err
		}
		user.Answers = append(user.Answers, array_tmp)
	}

	// uuidのnil
	var u uuid.NullUUID

	if user.QuestionID == u.UUID || user.Auther.UserID == u.UUID || len(user.Auther.UserName) == 0 || len(user.QuestionTag) == 0 || len(user.QuestionType) == 0 || len(user.CreateTime) == 0 || len(user.UpdateTime) == 0 || len(user.QuestionBody) == 0 || len(user.Values) == 0 || len(user.Answers) == 0 {
		user2 := question.Question{}
		var isEmpty error

		// エラーになる場合は，戻り値としてerrorをtrueとすることで，uuidによる00初期化を防止．．あしあし
		isEmpty = errors.New("true")
		return user2, isEmpty
	}

	// user を question.Questionとして定義している.
	return user, nil

	//return question.Question{}, nil
}
