package main

//package src
// 外部ファイルとして使えるようにする場合はsrc

import (
	"database/sql"
	"fmt"

	//"time"
	//"github.com/jinzhu/gorm"

	//"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// 動作確認用データベースで使用する構造体
/*
type DBTableInfos struct {
	id   uuid.UUID `db:userId json:id`
	name string    `db:userName json:name`
	//answer string    `db:answer json:answer`
}
*/

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

// 問題IDから問題の構造体を返す関数
func getQuestion(id uuid.UUID) Question {

	// データベースアクセス
	db, err := sql.Open("mysql", "root@/hacku")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// uuidを文字列へ変換
	searchId := id.String()

	// クエリ
	rows, err := db.Query("SELECT * FROM test WHERE QuestionId = ?", searchId)

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

// 問題をデータベースに登録する関数
func setQuestion(question Question) bool {
	db, err := sql.Open("mysql", "root@/hacku")
	if err != nil {
		panic(err.Error())
		return false
	}
	defer db.Close()

	// uuidを文字列へ変換している
	searchUserQuestion := question.QuestionId.String()
	userAuthorId := question.Author.String()

	stmtInsert, err := db.Prepare("INSERT IGNORE INTO test (QuestionId,Author,QuestionTag,QuestionType,CreateTime,UpdateTime,QuestionBody,Value,Answer) VALUES(?,?,?,?,?,?,?,?,?)")

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

func SaveData(Question_id interface{}) {
	//getQuestion(id)
	/*
		switch judge := Questionid.(type) {
			case uuid.UUID:
				tmp := getQuestion(judge)
				fmt.Println(tmp.id, tmp.name)
			case DBTableInfos:
				isAllRegist := setQuestion(judge)
				if isAllRegist {
					fmt.Println("OK")
				} else {
					fmt.Println("NG")
				}
		}
	*/
	fmt.Println("sql終了")
}

func main() {
	//SaveData(123)

	//u, _ := uuid.FromString("efcc59e4-9220-47df-b6a4-14f81313abc0")
	u, _ := uuid.FromString("3cc807ab-8e31-3071-aee4-f8f03781cb91")
	u2, _ := uuid.FromString("efcc59e4-9220-47df-b6a4-14f81313abc0")

	// テスト用構造体
	st := Question{u, u2, "{Go,Dart,C,Javascript}", "anaume", "1000-01-01 00:00:00.000000", "9999-12-31 23:59:59.999999", "以下の空欄を埋めなさい", "{[]って[]?}", "{hoge, hoge}"}

	var Question_ID interface{}

	Question_ID = st

	switch judge := Question_ID.(type) {
	case uuid.UUID:
		tmp := getQuestion(judge)

		// 構造体の確認
		fmt.Println(tmp)
	case Question:
		isAllRegist := setQuestion(judge)
		if isAllRegist {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}
