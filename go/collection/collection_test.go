package collection_test

import (
	"testing"

	"example.com/account"
	"example.com/collection"
	"example.com/question"
	"github.com/google/uuid"
)

func newUUID(s string) uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return uuid
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
			Author:       users[0],
			QuestionTag:  []string{"国語", "算数", "理科", "社会"},
			QuestionType: "4taku",
			CreateTime:   "1000-01-01 00:00:00",
			UpdateTime:   "1000-01-01 00:00:00",
			QuestionBody: "hogeってhoge?",
			Values:       []string{"hoge", "huga", "piyo", "foo"},
			Answers:      []string{"hoge"},
		},
		{
			QuestionID:   newUUID("c9d09c9e-e445-d1cf-5cb8-03d4507ad1f5"),
			Author:       users[1],
			QuestionTag:  []string{"Go", "Dart", "C", "JavaScript"},
			QuestionType: "anaume",
			CreateTime:   "1000-01-01 00:00:00",
			UpdateTime:   "1000-01-01 00:00:00",
			QuestionBody: "以下の空欄を埋めなさい",
			Values:       []string{"[]って[]?"},
			Answers:      []string{"hoge", "hoge"},
		},
	}

	collections = []collection.Collection{
		{
			CollectionID:          newUUID("81b989be-5f8a-a979-ed6d-7a4613ff08e3"),
			CollectionName:        "piyo問題集",
			CollectionDescription: "テスト用のコレクション",
			Author:                users[2],
			Questions:             questions,
			CreateTime:            "1000-01-01 00:00:00",
			UpdateTime:            "1000-01-01 00:00:00",
		},
		{
			CollectionID:          newUUID("2a2fd701-bb7e-6719-1f7b-ec67ffa5a269"),
			CollectionName:        "foo問題集",
			CollectionDescription: "ンョシクレコの用トステ",
			Author:                users[3],
			Questions: []question.Question{
				questions[1],
				questions[0],
			},
			CreateTime: "1000-01-01 00:00:00",
			UpdateTime: "1000-01-01 00:00:00",
		},
	}
)

type collectionCase struct {
	col collection.Collection
	err bool
}
type test struct {
	arg  string
	want collectionCase
}

func TestGetCollection(t *testing.T) {
	tests := []test{
		{
			// 正しい
			arg: collections[0].Author.ID(),
			want: collectionCase{
				col: collections[0],
				err: false,
			},
		},
		{
			// 正しい
			arg: collections[1].Author.ID(),
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
		col, err := collection.GetCollection(v.arg)
		if !v.want.col.Equals(col) {
			t.Error("case: ", i, " collection unamatch")
		}

		if (err != nil) != v.want.err {
			t.Error("case: ", i, " err status not valid")
		}
	}
}
