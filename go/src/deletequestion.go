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

func Deletequestion(question uuid.UUID) bool {
	db, err := sql.Open("mysql", "root@/hacku_db")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()
	stmtDelete, err := db.Prepare("DELETE FROM question WHERE QuestionId=?")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer stmtDelete.Close()

	// uuidを文字列に変換
	searchUserQuestion := question.String()

	result, err := stmtDelete.Exec(searchUserQuestion)
	if err != nil {
		panic(err.Error())
		return false
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
		return false
	}
	fmt.Println(rowsAffect)
	return true
}
