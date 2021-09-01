package src

import (
	"database/sql"
	"fmt"

	//"time"
	//"github.com/jinzhu/gorm"
        "example.com/question"
        "example.com/account"
	//"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// 問題をデータベースに登録する関数
func Savedata(question Question) bool {
	db, err := sql.Open("mysql", "root@/hacku_db")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()

	// uuidを文字列へ変換している
	searchUserQuestion := question.QuestionId.String()
	userAuthorId := question.Author.String()

	stmtInsert, err := db.Prepare("INSERT IGNORE INTO question (QuestionId,Author,QuestionTag,QuestionType,CreateTime,UpdateTime,QuestionBody,Value,Answers) VALUES(?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
		return false
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(searchUserQuestion, userAuthorId, question.QuestionTag, question.QuestionType, question.CreateTime, question.UpdateTime, question.QuestionBody, question.QuestionBody, question.Answers)
	if err != nil {
		panic(err.Error())
		return false
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
		return false
	}
	fmt.Println(lastInsertID)

	return true
}
