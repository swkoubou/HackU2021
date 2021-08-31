package src

import (
	"database/sql"
	"fmt"

	//"time"
	//"github.com/jinzhu/gorm"

	//"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	UserID   uuid.UUID `db:"userId"`
	UserName string    `db:"userName"`
}

type Question struct {
	QuestionId   uuid.UUID `db:"QuestionId"`
	Author       uuid.UUID `db:"Author"`
	QuestionTag  string    `db:"QuestionTag"`
	QuestionType string    `db:QuestionType`
	CreateTime   string    `db:"CreateTime"`
	UpdateTime   string    `db:"UpdateTime"`
	QuestionBody string    `db:"QuestionBody"`
	Value        string    `db:"Value"`
	Answers      string    `db:"Answers"`
}

func Deletequestion(question uuid.UUID) {
	db, err := sql.Open("mysql", "root@/hacku_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtDelete, err := db.Prepare("DELETE FROM question WHERE QuestionId=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	// uuidを文字列に変換
	searchUserQuestion := question.String()

	result, err := stmtDelete.Exec(searchUserQuestion)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)
}
