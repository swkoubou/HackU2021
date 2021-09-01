package src

import (
	"database/sql"

	//"time"
	//"github.com/jinzhu/gorm"
    "example.com/question"
    "example.com/account"
	//"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// 問題IDから問題の構造体を返す関数
func GetQuestion(id uuid.UUID) Question {

	// データベースアクセス
	db, err := sql.Open("mysql", "root@/hacku_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// uuidを文字列へ変換
	searchId := id.String()

	// クエリ
	rows, err := db.Query("SELECT * FROM question WHERE QuestionId = ?", searchId)

	defer rows.Close()

	if err != nil {
		panic(err.Error())
	}
	// 返す構造体を作成
	var check Question

	var tmp1, tmp2 string

	for rows.Next() {
		if err := rows.Scan(&tmp1, &tmp2, &check.QuestionTag, &check.QuestionType, &check.CreateTime, &check.UpdateTime, &check.QuestionBody, &check.Value, &check.Answers); err != nil {
			panic(err.Error())
		}
		// 文字列をuuidに変換
		check.QuestionId, _ = uuid.FromString(tmp1)
		check.Author, _ = uuid.FromString(tmp2)
	}
	return check
}
