package question_test

import (
	"testing"

	"example.com/account"
	"example.com/question"
	"github.com/google/uuid"
)

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
)

func newUUID(s string) uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return uuid
}

func TestGetQuestion(t *testing.T) {
	type questionCase struct {
		question question.Question
		err      bool
	}
	type testCase struct {
		arg  string
		want questionCase
	}

	cases := []testCase{
		{
			arg: questions[0].ID(),
			want: questionCase{
				question: questions[0],
				err:      false,
			},
		},
		{
			arg: questions[1].ID(),
			want: questionCase{
				question: questions[1],
				err:      false,
			},
		},
		{
			// 存在しないQuestion
			arg: "a56ada75-9118-afaa-4bb3-b2b655b9b4b7",
			want: questionCase{
				question: question.Question{},
				err:      true,
			},
		},
		{
			// 正しくないQuestionID
			arg: "hoge",
			want: questionCase{
				question: question.Question{},
				err:      true,
			},
		},
		{
			// 空文字
			arg: "",
			want: questionCase{
				question: question.Question{},
				err:      true,
			},
		},
	}

	for i, v := range cases {
		q, err := question.GetQuestion(v.arg)
		if (err != nil) != v.want.err {
			t.Error("case: ", i, err)
		}

		if ok := v.want.question.Equals(q); !ok {
			t.Error("case: ", i, err)
		}
	}
}

func TestSetQuestion(t *testing.T) {
	type questionCase struct {
		question question.Question
		err      bool
	}
	type testCase struct {
		arg  question.QuestionParam
		want questionCase
	}

	cases := []testCase{
		{
			arg: question.QuestionParam{
				UserID:       users[2].UserID.String(),
				QuestionTag:  []string{"あか", "さた", "な"},
				QuestionType: "4taku",
				QuestionBody: "てすともんだいてすともんだい",
				Values:       []string{"hogehogehoge", "hugahugahuga", "piyopiyopiyo", "foofoofoo"},
				Answers:      []string{"hogehogehoge"},
			},
			want: questionCase{
				question: question.Question{
					QuestionID:   uuid.Nil,
					Author:       users[2],
					QuestionTag:  []string{"あか", "さた", "な"},
					QuestionType: "4taku",
					CreateTime:   "1000-01-01 00:00:00",
					UpdateTime:   "1000-01-01 00:00:00",
					QuestionBody: "てすともんだいてすともんだい",
					Values:       []string{"hogehogehoge", "hugahugahuga", "piyopiyopiyo", "foofoofoo"},
					Answers:      []string{"hogehogehoge"},
				},
				err: false,
			},
		},
		{
			arg: question.QuestionParam{
				UserID:       users[2].UserID.String(),
				QuestionTag:  []string{},
				QuestionType: "anaume",
				QuestionBody: "以下の空欄を埋めなさい",
				Values:       []string{"piyo[]piyo[]piyo[]"},
				Answers:      []string{"poe", "poe", "poe"},
			},
			want: questionCase{
				question: question.Question{
					QuestionID:   uuid.Nil,
					Author:       users[2],
					QuestionTag:  []string{},
					QuestionType: "anaume",
					CreateTime:   "1000-01-01 00:00:00",
					UpdateTime:   "1000-01-01 00:00:00",
					QuestionBody: "以下の空欄を埋めなさい",
					Values:       []string{"piyo[]piyo[]piyo[]"},
					Answers:      []string{"poe", "poe", "poe"},
				},
				err: false,
			},
		},
		{
			arg: question.QuestionParam{
				UserID:       users[2].ID(),
				QuestionTag:  []string{},
				QuestionType: "",
				QuestionBody: "以下の空欄を埋めなさい",
				Values:       []string{"piyo[]piyo[]piyo[]"},
				Answers:      []string{"poe", "poe", "poe"},
			},
			want: questionCase{
				question: question.Question{},
				err:      true,
			},
		},
		{
			arg: question.QuestionParam{
				UserID:       users[2].ID(),
				QuestionTag:  []string{},
				QuestionType: "anaume",
				QuestionBody: "hogehoge",
				Values:       []string{"piyo[]piyo[]piyo[]"},
				Answers:      []string{"poe", "poe", "poe"},
			},
			want: questionCase{
				question: question.Question{},
				err:      true,
			},
		},
	}

	for i, v := range cases {
		id, err := question.SetQuestion(&v.arg)
		if (err != nil) != v.want.err {
			t.Error("case: ", i, err)
		}
		v.want.question.QuestionID = newUUID(id)

		q, _ := question.GetQuestion(id)

		if !v.want.err {
			if !v.want.question.Equals(q) {
				t.Error("case: ", i, "question unmatch")
			}
		} else {
			if q != nil {
				t.Error("case: ", i, "question unmatch")
			}
		}
	}
}
