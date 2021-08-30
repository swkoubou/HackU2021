package main

//package src
// 外部ファイルとして使えるようにする場合はsrc

import (
	"database/sql"
	"fmt"

	//"time"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 動作確認用データベースで使用する構造体

type DBTableInfos struct {
	id     int    `db:id json:id`
	name   string `db:name json:name`
	answer string `db:answer json:answer`
}

/*
type DBTableInfos struct {
	questionId      uint   `db:"id"`
	questionName    string `db:"questionName"`
	userId          int
	questionType    string
	questionTag     []string
	author          string
	createTime      time.Time
	updateTime      time.Time
	questionAnswer  string
	quesitonAnswers []string
	value           [4]string
}
*/

// 問題IDから問題の構造体を返す関数
func getQuestion(id int) DBTableInfos {
	db, err := sql.Open("mysql", "root@/hacku")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM question_test WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	// 返す構造体を作成
	var check DBTableInfos

	for rows.Next() {
		var id int
		var name string
		var answer string

		// 構造体に直接でも良さそう
		if err := rows.Scan(&id, &name, &answer); err != nil {
			panic(err.Error())
		}

		// プログラムが動くかの確認用のprint文
		fmt.Println(id, name, answer)

		// 構造体に入れている
		check.id = id
		check.name = name
		check.answer = answer

	}

	return check
}

// 問題をデータベースに登録する関数
func setQuestion(question DBTableInfos) bool {
	db, err := sql.Open("mysql", "root@/hacku")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO question_test (id,name,answer) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(question.id, question.name, question.answer)
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

func SaveData(id int) {
	getQuestion(id)
	/*
		switch id.(type) {
		case int:
			// idがint型なら構造体型を返す
			getQuestion(id)
		default:
			// 構造体型ならデータベースにinsertする

		}
	*/
	fmt.Println("sql終了")
}

func main() {
	//SaveData(123)
	var QuestionID interface{}

	// テスト用構造体

	st := DBTableInfos{234, "ひぐらし", "北条沙都子"}

	QuestionID = st

	switch judge := QuestionID.(type) {
	case int:
		tmp := getQuestion(judge)
		fmt.Println(tmp)
	case DBTableInfos:
		isAllRegist := setQuestion(judge)
		if isAllRegist {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

// 以下はデータベースにアクセスできるかの確認で使用
/*
func main() {
	db, err := sql.Open("mysql", "root@/hacku")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM question_test")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		var answer string
		if err := rows.Scan(&id, &name, &answer); err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name, answer)
	}
}
*/
