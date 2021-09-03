package manage

import (
	"errors"
	"reflect"
	"testing"

	"example.com/account"
	"example.com/collection"
	"example.com/question"
	"github.com/google/uuid"
)

type testCase struct {
	arg string
	ret
}

type ret struct {
	question question.Question
	err      bool
}

func matchStrArr(want, got []string) bool {
	if len(want) != len(got) {
		return false
	}

	for i, v := range want {
		if v != got[i] {
			return false
		}
	}
	return true
}

func equal(want, got question.Question) error {
	if reflect.TypeOf(want) != reflect.TypeOf(got) {
		return errors.New("type is not valid")
	}

	if want.QuestionID != got.QuestionID {
		return errors.New("QuestionID is not valid")
	}

	if want.Auther != got.Auther {
		return errors.New("Auther is not valid")
	}

	if !matchStrArr(want.QuestionTag, got.QuestionTag) {
		return errors.New("QuestionTag is not valid")
	}

	// 時刻は決め撃ちでは無い為、長さだけ検証
	if len(got.CreateTime) != len(want.CreateTime) {
		return errors.New("CreateTime is not valid")
	}

	if len(got.UpdateTime) != len(want.UpdateTime) {
		return errors.New("UpdateTime is not valid")
	}

	if want.QuestionBody != got.QuestionBody {
		return errors.New("QuestionBody is not valid")
	}

	if !matchStrArr(want.Values, got.Values) {
		return errors.New("Values is not valid")
	}

	if !matchStrArr(want.Answers, got.Answers) {
		return errors.New("Answers is not valid")
	}

	return nil
}

var (
	users = []account.Account{
		{
			UserID:   newUUID("c74b366b-5289-d4f5-1d80-66fedaafe97b"),
			UserName: "hoge",
		},
		{
			UserID:   newUUID("7125174d-6b38-d3ed-b808-ae0a4416f589"),
			UserName: "huga",
		},
		{
			UserID:   newUUID("116aa423-d551-2b81-2091-132e022d40c5"),
			UserName: "piyo",
		},
		{
			UserID:   newUUID("30df398b-63f3-351d-2cd2-78ce9df47faf"),
			UserName: "foo",
		},
	}

	questions = []question.Question{
		{
			QuestionID:   newUUID("1a4ee1ec-1073-f9fb-8281-f3041d15a9d2"),
			Auther:       users[0],
			QuestionTag:  []string{"国語", "算数", "理科", "社会"},
			QuestionType: "4taku",
			CreateTime:   "1000-01-01 00:00:00.000000",
			UpdateTime:   "1000-01-01 00:00:00.000000",
			QuestionBody: "hogeってhoge?",
			Values:       []string{"hoge", "huga", "piyo", "foo"},
			Answers:      []string{"hoge"},
		},
		{
			QuestionID:   newUUID("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5"),
			Auther:       users[1],
			QuestionTag:  []string{"Go", "Dart", "C", "JavaScript"},
			QuestionType: "anaume",
			CreateTime:   "1000-01-01 00:00:00.000000",
			UpdateTime:   "1000-01-01 00:00:00.000000",
			QuestionBody: "以下の空欄を埋めなさい",
			Values:       []string{"[]って[]?"},
			Answers:      []string{"hoge", "hoge"},
		},
	}

	collections = []collection.Collection{
		{
			CollectionID:           newUUID("81b989be-5f8a-a979-ed6d-7a4613ff08e3"),
			CollectionName:         "piyo問題集",
			CollectionDescripition: "テスト用のコレクション",
			Auther:                 users[2],
			Questions:              questions,
			CreateTime:             "1000-01-01 00:00:00.000000",
			UpdateTime:             "1000-01-01 00:00:00.000000",
		},
		{
			CollectionID:           newUUID("2a2fd701-bb7e-6719-1f7b-ec67ffa5a269"),
			CollectionName:         "foo問題集",
			CollectionDescripition: "ンョシクレコの用トステ",
			Auther:                 users[3],
			Questions: []question.Question{
				questions[1],
				questions[0],
			},
			CreateTime: "1000-01-01 00:00:00.000000",
			UpdateTime: "1000-01-01 00:00:00.000000",
		},
	}
)

func newUUID(s string) uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return uuid
}

func TestGetQuestion(t *testing.T) {
	var ids []uuid.UUID
	for _, v := range []string{"1a4ee1ec-1073-f9fb-8281-f3041d15a9d2", "c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5", "c74b366b-5289-d4f5-1d80-66fedaafe97b", "7125174d-6b38-d3ed-b808-ae0a4416f589"} {
		uuid, err := uuid.Parse(v)
		if err != nil {
			t.Fail()
		}
		ids = append(ids, uuid)
	}

	cases := []testCase{
		{
			arg: "1a4ee1ec-1073-f9fb-8281-f3041d15a9d2",
			ret: ret{
				question: question.Question{
					QuestionID: ids[0],
					Auther: account.Account{
						UserID:   ids[2],
						UserName: "hoge",
					},
					QuestionTag:  []string{"国語", "算数", "理科", "社会"},
					QuestionType: "4taku",
					CreateTime:   "2021-09-01 22:42:00",
					UpdateTime:   "2021-09-01 22:42:00",
					QuestionBody: "hogeってhoge?",
					Values:       []string{"hoge", "huga", "piyo", "foo"},
					Answers:      []string{"hoge"},
				},
				err: false,
			},
		},
		{
			arg: "c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5",
			ret: ret{
				question: question.Question{
					QuestionID: ids[1],
					Auther: account.Account{
						UserID:   ids[3],
						UserName: "huga",
					},
					QuestionTag:  []string{"Go", "Dart", "C", "JavaScript"},
					QuestionType: "anaume",
					CreateTime:   "2021-09-01 22:42:00",
					UpdateTime:   "2021-09-01 22:42:00",
					QuestionBody: "以下の空欄を埋めなさい",
					Values:       []string{"[]って[]?"},
					Answers:      []string{"hoge", "hoge"},
				},
				err: false,
			},
		},
		{
			// 存在しないQuestion
			arg: "a56ada75-9118-afaa-4bb3-b2b655b9b4b7",
			ret: ret{
				question: question.Question{},
				err:      true,
			},
		},
		{
			// 正しくないQuestionID
			arg: "hoge",
			ret: ret{
				question: question.Question{},
				err:      true,
			},
		},
		{
			// 空文字
			arg: "",
			ret: ret{
				question: question.Question{},
				err:      true,
			},
		},
	}

	for i, v := range cases {
		q, err := GetQuestion(v.arg)
		if (err != nil) != v.err {
			t.Error("case: ", i, err)
		}

		if err := equal(v.question, q); err != nil {
			t.Error("case: ", i, err)
		}
	}

}

type test struct {
	arg  string
	want collectionCase
}

type collectionCase struct {
	col collection.Collection
	err bool
}

func TestGetCollection(t *testing.T) {
	tests := []test{
		{
			// 正しい
			arg: collections[0].Auther.ID(),
			want: collectionCase{
				col: collections[0],
				err: false,
			},
		},
		{
			// 正しい
			arg: collections[1].Auther.ID(),
			want: collectionCase{
				col: collections[1],
				err: false,
			},
		},
		{
			// 存在しない
			arg: "3bb406d7-0ad9-dc9f-bec6-6c26e17726bf",
			want: collectionCase{
				col: collection.Collection{},
				err: true,
			},
		},
		{
			// IDが正しくない
			arg: "hoge",
			want: collectionCase{
				col: collection.Collection{},
				err: true,
			},
		},
		{
			// 空文字
			arg: "",
			want: collectionCase{
				col: collection.Collection{},
				err: true,
			},
		},
	}

	for i, v := range tests {
		col, err := GetCollection(v.arg)
		if !v.want.col.Equals(&col) {
			t.Error("case: ", i, " collection unamatch")
		}

		if (err != nil) != v.want.err {
			t.Error("case: ", i, " err status not valid")
		}
	}
}
