//package question
package main

import (
	"database/sql"
	"errors"
	"reflect"

	"example.com/account"
	"example.com/manage"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Question struct {
	QuestionID   uuid.UUID       `json:"questionID"`
	Auther       account.Account `json:"author"`
	QuestionName string          `json:"questionName"`
	QuestionTag  []string        `json:"questionTag"`
	QuestionType string          `json:"questionType"`
	CreateTime   string          `json:"createTime"`
	UpdateTime   string          `json:"updateTime"`
	QuestionBody string          `json:"questionBody"`
	Values       []string        `json:"values"`
	Answers      []string        `json:"answers"`
}

type QuestionParam struct {
	UserID       string
	QuestionName string
	QuestionTag  []string
	QuestionType string
	QuestionBody string
	Values       []string
	Answers      []string
}

func (q *Question) ID() string {
	return q.QuestionID.String()
}

func NewQuestion(param QuestionParam) (Question, error) {
	q := Question{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return q, err
	}

	q.QuestionID = uuid
	q.Auther = account.Account{} // DBからもってくる
	q.QuestionTag = param.QuestionTag
	q.QuestionBody = param.QuestionBody
	q.Values = param.Values
	q.Answers = param.Answers

	return q, nil
}

func (q *Question) Equals(question *Question) bool {
	if reflect.TypeOf(*q) != reflect.TypeOf(*question) {
		return false
	}

	if q.QuestionID != question.QuestionID {
		return false
	}

	if !reflect.DeepEqual(q.Auther, question.Auther) {
		return false
	}

	if len(q.QuestionTag) != len(question.QuestionTag) {
		return false
	}

	for i, v := range q.QuestionTag {
		if v != question.QuestionTag[i] {
			return false
		}
	}

	if len(q.CreateTime) != len(question.CreateTime) {
		return false
	}

	if len(q.UpdateTime) != len(question.UpdateTime) {
		return false
	}

	if q.QuestionBody != question.QuestionBody {
		return false
	}

	for i, v := range q.Values {
		if v != question.Values[i] {
			return false
		}
	}

	for i, v := range q.Answers {
		if v != question.Answers[i] {
			return false
		}
	}

	return true
}

func GetQuestion(questionID string) (*Question, error) {
	var user Question

	// データベースへのアクセス
	db, err := manage.NewDBConnection()

	// エラー処理
	if err != nil {
		return nil, err
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// クエリ
	rows1, err := db.Query("SELECT * FROM question WHERE question_id = ?", questionID)

	defer rows1.Close()

	// エラー処理
	if err != nil {
		return nil, err
	}

	var uuid_tmp1, uuid_tmp2 string

	// ここからデータベースでの処理
	for rows1.Next() {
		if err := rows1.Scan(&uuid_tmp1, &user.QuestionName, &uuid_tmp2, &user.QuestionType, &user.CreateTime, &user.UpdateTime, &user.QuestionBody); err != nil {
			return nil, err
		}
		user.QuestionID, _ = uuid.Parse(uuid_tmp1)
		user.Auther.UserID, _ = uuid.Parse(uuid_tmp2)
	}

	rows2, err := db.Query("SELECT * FROM user WHERE user_id = ?", uuid_tmp2)

	if err != nil {
		return nil, err
	}

	var name_tmp string

	for rows2.Next() {
		if err := rows2.Scan(&name_tmp, &user.Auther.UserName); err != nil {
			return nil, err
		}
	}

	rows3, err := db.Query("SELECT * FROM question_tag_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	var questionid_tmp, array_tmp string

	var tag_int int

	for rows3.Next() {
		if err := rows3.Scan(&questionid_tmp, &tag_int); err != nil {
			return nil, err
		}

		rows4, err := db.Query("SELECT * FROM question_tag WHERE question_tag_id = ?", tag_int)
		if err != nil {
			return nil, err
		}

		for rows4.Next() {
			if err := rows4.Scan(&tag_int, &array_tmp); err != nil {
				return nil, err
			}
			user.QuestionTag = append(user.QuestionTag, array_tmp)
		}
	}

	rows5, err := db.Query("SELECT * FROM question_value_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	var int_tmp int

	for rows5.Next() {
		if err := rows5.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return nil, err
		}
		user.Values = append(user.Values, array_tmp)
	}

	rows6, err := db.Query("SELECT * FROM question_answer_map WHERE question_id = ?", uuid_tmp1)

	if err != nil {
		return nil, err
	}

	for rows6.Next() {
		if err := rows6.Scan(&questionid_tmp, &int_tmp, &array_tmp); err != nil {
			return nil, err
		}
		user.Answers = append(user.Answers, array_tmp)
	}

	// uuidのnil
	var u uuid.NullUUID

	if user.QuestionID == u.UUID || user.Auther.UserID == u.UUID || len(user.Auther.UserName) == 0 || len(user.QuestionTag) == 0 || len(user.QuestionType) == 0 || len(user.CreateTime) == 0 || len(user.UpdateTime) == 0 || len(user.QuestionBody) == 0 || len(user.Values) == 0 || len(user.Answers) == 0 {
		var isEmpty error

		// エラーになる場合は，戻り値としてerrorをtrueとすることで，uuidによる00初期化を防止．．あしあし
		isEmpty = errors.New("true")
		return nil, isEmpty
	}
	// user を question.Questionとして定義している.
	return &user, nil

	//return question.Question{}, nil
}

/*
引数: QuestionParam
type QuestionParam struct {
	UserID       string		必須
	QuestionName string     必須
	QuestionTag  []string	0個以上
	QuestionType string		必須
	QuestionBody string		QuestionTypeがanaumeの場合"以下の空欄を埋めなさい"で固定
	Values       []string	必須
	Answers      []string	必須
}

戻り値: string, error
データの挿入が成功した場合はQuestionIDとnilを返す
データの挿入に失敗した場合は空文字とerrorを返す
*/

// 渡されたQuestionParam構造体をデータベースに保存する関数
func SetQuestion(q *QuestionParam) (string, error) {
	var question Question
	// データベースへのアクセス
	db, err := manage.NewDBConnection()

	// エラー処理
	if err != nil {
		return string(""), nil
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// questionテーブルに挿入している．
	// ここで，挿入しているのは
	// question_id , question_name , user_id , question_type ,  question_body　である

	name_row, err := db.Query("SELECT user_name FROM user WHERE user_id = ?", q.UserID)

	if err != nil {
		return string(""), nil
	}

	for name_row.Next() {
		if err := name_row.Scan(&question.Auther.UserName); err != nil {
			return string(""), nil
		}
	}

	// 最初のinsert文

	question.QuestionID = uuid.New()

	stmtInsert1, err := db.Prepare("INSERT INTO question (question_id,question_name,user_id,question_type,question_body) VALUES(?,?,?,?,?)")

	if err != nil {
		return string(""), nil
	}

	_, err = stmtInsert1.Exec(question.QuestionID, q.QuestionName, q.UserID, q.QuestionType, q.QuestionBody)

	if err != nil {
		return string(""), nil
	}

	var subject_nums []int

	cnt := 0

	for i := 0; i < len(q.QuestionTag); i++ {

		var subject_name string

		subject_name = q.QuestionTag[i]

		subject_row, err := db.Query("SELECT question_tag_id FROM question_tag WHERE question_tag_name = ?", subject_name)

		if err != nil {
			return string(""), nil
		}

		for subject_row.Next() {

			var tmp_question_tag_id int

			if err := subject_row.Scan(&tmp_question_tag_id); err != nil {
				return string(""), nil
			}

			cnt++

			subject_nums = append(subject_nums, tmp_question_tag_id)

		}

		if cnt == i {
			cnt++

			count_row, err := db.Query("SELECT count(question_tag_id) FROM question_tag")

			var count_tmp int

			for count_row.Next() {
				if err := count_row.Scan(&count_tmp); err != nil {
					return string(""), nil
				}

			}

			count_tmp++

			stmtInsert2, err := db.Prepare("INSERT IGNORE INTO question_tag (question_tag_id,question_tag_name) VALUES(?,?)")

			if err != nil {
				return string(""), nil
			}

			_, err = stmtInsert2.Exec(count_tmp, subject_name)

			subject_nums = append(subject_nums, count_tmp)

		}
	}

	// question_tag_mapに挿入する処理
	for i := 0; i < len(subject_nums); i++ {

		var row_num int

		row_num = subject_nums[i]

		stmtInsert3, err := db.Prepare("INSERT IGNORE INTO question_tag_map (question_id,question_tag_id) VALUES(?,?)")

		if err != nil {
			return string(""), nil
		}

		_, err = stmtInsert3.Exec(question.QuestionID, row_num)

		if err != nil {
			return string(""), nil
		}
	}

	// question_value_mapに挿入する処理

	for i := 0; i < len(q.Values); i++ {

		var question_value string

		question_value = q.Values[i]

		stmtInsert4, err := db.Prepare("INSERT IGNORE INTO question_value_map (question_id,value_order,question_value) VALUES(?,?,?)")

		if err != nil {
			return string(""), nil
		}

		_, err = stmtInsert4.Exec(question.QuestionID, i, question_value)

		if err != nil {
			return string(""), nil
		}
	}

	// question_answer_mapに挿入する処理

	for i := 0; i < len(q.Answers); i++ {
		var answer_tmp string

		answer_tmp = q.Answers[i]

		stmtInsert5, err := db.Prepare("INSERT IGNORE INTO question_answer_map (question_id,answer_order,question_answer) VALUES(?,?,?)")

		if err != nil {
			return string(""), nil
		}

		_, err = stmtInsert5.Exec(question.QuestionID, i, answer_tmp)

		if err != nil {
			return string(""), nil
		}
	}

	// 成功したら
	return question.QuestionID.String(), nil

	//return string(""), nil
}

/*
引数: Question
type Question struct {
	QuestionID   uuid.UUID       `json:"questionID"`	必須
	Auther       account.Account `json:"auther"`		必須
	QuestionTag  []string        `json:"questionTag"`	0個以上
	QuestionType string          `json:"questionType"`	必須
	CreateTime   string          `json:"createTime"`	更新しない
	UpdateTime   string          `json:"updateTime"`	勝手に更新されるはず
	QuestionBody string          `json:"questionBody"`	QuestionTypeがanaumeの場合"以下の空欄を埋めなさい"で固定
	Values       []string        `json:"values"`		必須: QuestionTypeで適切かどうか判別してください
	Answers      []string        `json:"answers"`		必須: QuestionTypeで適切かどうか判別してください
}

戻り値: string, error
データの更新が成功した場合はQuestionIDとnilを返す
データの更新に失敗した場合は空文字とerrorを返す
*/

func UpdateQuestion(q *Question) (string, error) {

	db, err := sql.Open("mysql", "root@/MYSQL_DATABASE")

	// エラー処理
	if err != nil {
		return string(""), nil
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// update
	// stmtUpdate, err := db.Prepare("UPDATE users SET name=? WHERE id=?")

	return string(""), nil
}

/*
引数: string
QuestionID

戻り値: string, error
データの削除が成功した場合はnilを返す
データの削除に失敗した場合はerrorを返す
*/
func DeleteQuestion(id string) error {

	db, err := manage.NewDBConnection()

	// エラー処理
	if err != nil {
		return err
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	// 消すデータベース
	// question_answer_map , question_tag_map , question_value_map , question

	stmtDelete1, err := db.Prepare("DELETE FROM question_answer_map WHERE question_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDelete1.Exec(id)
	if err != nil {
		return err
	}

	stmtDelete2, err := db.Prepare("DELETE FROM question_tag_map WHERE question_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDelete2.Exec(id)
	if err != nil {
		return err
	}

	stmtDelete3, err := db.Prepare("DELETE FROM question_value_map WHERE question_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDelete3.Exec(id)
	if err != nil {
		return err
	}

	stmtDelete4, err := db.Prepare("DELETE FROM question WHERE question_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDelete4.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

/*
引数: string
QuestionID

戻り値: Account, error
データの取得が成功した場合はAccountとnilを返す
データの取得に失敗した場合はnilとerrorを返す
*/
func GetAuthor(id string) (*account.Account, error) {
	var author account.Account

	author.UserID, _ = uuid.Parse(id)

	db, err := manage.NewDBConnection()

	// エラー処理
	if err != nil {
		return nil, err
	}

	// returnされた後に実行され，DBとの接続を切る
	defer db.Close()

	rows, err := db.Query("SELECT user_name FROM user WHERE user_id = ?", id)

	defer rows.Close()

	// エラー処理
	if err != nil {
		return nil, err
	}

	// ここからデータベースでの処理
	for rows.Next() {
		if err := rows.Scan(&author.UserName); err != nil {
			return nil, err
		}
	}

	// author を account.Account{}としている

	return &author, nil
	//return &account.Account{}, nil
}

func main() {

	/*
	   // GetAuthorのテスト
	   	st := "116aa423-d551-2b81-2091-132e022d40c5"

	   	ans, err := GetAuthor(st)

	   	if err != nil {
	   		fmt.Println("error")
	   	}

	   	fmt.Println(ans)
	*/
	/*
		// DeleteQuestionのテスト
		st := "dec9318b-83db-4f3c-bc14-62037c0ff722"

		ans := DeleteQuestion(st)

		fmt.Println(ans)
	*/
	/*
		// setQuestionのテスト
		var test QuestionParam

		test.UserID = "XYZ"
		test.QuestionName = "ひぐらし"
		test.QuestionTag = append(test.QuestionTag, "国語")
		test.QuestionTag = append(test.QuestionTag, "アニメ")
		test.QuestionType = "4taku"
		test.QuestionBody = "にぱー"
		test.Values = append(test.Values, "orz")
		test.Values = append(test.Values, "aiueo")
		test.Answers = append(test.Answers, "面白い")
		test.Answers = append(test.Answers, "神アニメ")

		ans, err := SetQuestion(&test)

		if err != nil {
			fmt.Println("error")
		}

		fmt.Println(ans)
	*/

	/*
		// GetQuestionのテスト
		var err error
		st := "1a4ee1ec-1073-f9fb-8281-f3041d15a9d2"
		tt, err := GetQuestion(st)
		if err != nil {
			fmt.Println("hello")
		}
		fmt.Println(tt)
	*/
}
