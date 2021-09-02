package manage

import (
	"errors"
	"reflect"
	"testing"

	"example.com/account"
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
