package manage

import (
	"database/sql"
	"os"

	"example.com/collection"
	"example.com/question"
	_ "github.com/go-sql-driver/mysql"
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
	return question.Question{}, nil
}

func GetCollection(collectionID string) (collection.Collection, error) {
	return collection.Collection{}, nil
}
