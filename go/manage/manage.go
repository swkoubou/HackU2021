package manage

import (
	"database/sql"
	"fmt"

	"example.com/account"
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
	/*
		dsn := &dsn{
			user:     "root",
			pass:     os.Getenv("MYSQL_ROOT_PASSWORD"),
			protocol: "tcp",
			address:  "db:3306",
			dbName:   os.Getenv("MYSQL_DATABASE"),
		}
	*/
	//return sql.Open("mysql", dsn.user+":"+dsn.pass+"@"+dsn.protocol+"("+dsn.address+")/"+dsn.dbName)
	return sql.Open("mysql", "root@/MYSQL_DATABASE")
}

func GetQuestion(questionID string) (question.Question, error) {

	// データベースへのアクセス
	db, err := NewDBConnection()

	// エラー処理
	if err != nil {
		panic(err.Error())
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// クエリ
	rows, err := db.Query("SELECT * FROM question WHERE question_id = ?", questionID)

	defer rows.Close()

	// エラー処理
	if err != nil {
		panic(err.Error())
	}

	var uuid_tmp1, uuid_tmp2 string

	var tmp string

	var user question.Question

	var userInfo account.Account

	for rows.Next() {
		if err := rows.Scan(&uuid_tmp1, &tmp, &uuid_tmp2, &user.QuestionType, &user.CreateTime, &user.UpdateTime, &user.QuestionBody); err != nil {
			panic(err.Error())
		}
		user.QuestionID, _ = uuid.Parse(uuid_tmp1)
		userInfo.UserID, _ = uuid.Parse(uuid_tmp2)
	}

	//rows, err := db.Query("SELECT * FROM question WHERE question_id = ?", questionID)

	fmt.Println(user)
	return question.Question{}, nil
}
