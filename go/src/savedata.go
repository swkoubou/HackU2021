package src

import (
	"database/sql"
	"fmt"

	//"time"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 動作確認用データベースで使用する構造体

type DBTableInfos struct {
	id     uint   `db:id json:id`
	name   string `db:name json:name`
	answer string `db:answer json:answer`
}

/*
type DBTableInfos struct {
	questionId      uint   `db:"id" json:"id"`
	questionName    string `db:"questionName" json:"questionName"`
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

/*
func access_to_db() *gorm.DB {
	DBMS := "mysql"
	USER := "localhost"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}
*/

// 問題IDから問題の構造体を返す関数
func getQuestion(id int) {
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

// 問題をデータベースに登録する関数
/*
func setQuestion() bool {

}
*/

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

/*
func main() {
	SaveData(123)
}
*/

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
