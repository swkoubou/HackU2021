package Question

type User struct {
	UserID   uuid   `json:"name"`
	UserName string `json:"userName"`
}
type Question struct {
	QuestionTag  []string `json:questionTag`
	QuestionID   uuid     `json:"questionID"`
	Auther       user     `json:"auther`
	CreateTime   time     `json:"createTime"`
	UpdateTime   time     `json:"updateTime"`
	QuestionBody []string `json:"questionBody"`
	Values       []string `json:"values"`
	Answers      []string `json:"answers"`
}

func newQuestion() {
	q := Question{QuestionID: 11}
}
