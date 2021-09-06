package handle

import (
	"log"
	"net/http"

	"example.com/account"
	"example.com/collection"
	"example.com/question"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type Handle struct {
	app  *firebase.App
	auth *auth.Client
}

func NewHandle(app *firebase.App, auth *auth.Client) *Handle {
	return &Handle{app: app, auth: auth}

}

func newUUID() uuid.UUID {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return uuid
}

var accounts []account.Account = []account.Account{
	{
		UserID:   newUUID(),
		UserName: "hoge",
	},
	{
		UserID:   newUUID(),
		UserName: "huga",
	},
	{
		UserID:   newUUID(),
		UserName: "piyo",
	},
	{
		UserID:   newUUID(),
		UserName: "foo",
	},
}

var questions []question.Question = []question.Question{
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionTag:  []string{"国語", "算数", "理科", "社会"},
		QuestionType: "4taku",
		CreateTime:   "9999-12-31 23:59:59.999999",
		UpdateTime:   "",
		QuestionBody: "hogeってhoge?",
		Values:       []string{"hoge", "huga", "piyo", "foo"},
		Answers:      []string{"hoge"},
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionTag:  []string{"Go", "Dart", "C", "JavaScript"},
		QuestionType: "anaume",
		CreateTime:   "1000-01-01 00:00:00.000000",
		UpdateTime:   "9999-12-31 23:59:59.999999",
		QuestionBody: "以下の空欄を埋めなさい",
		Values:       []string{"[]って[]?"},
		Answers:      []string{"hoge", "hoge"},
	},
}

var collections []collection.Collection = []collection.Collection{
	{
		CollectionID:           newUUID(),
		CollectionName:         "foo",
		CollectionDescripition: "テスト用のコレクション",
		Auther:                 accounts[2],
		Questions:              questions,
		CreateTime:             "9999-12-31 23:59:59.999999",
		UpdateTime:             "",
	},
	{
		CollectionID:           newUUID(),
		CollectionName:         "bar",
		CollectionDescripition: "ンョシクレコの用トステ",
		Auther:                 accounts[2],
		Questions:              questions,
		CreateTime:             "1000-01-01 00:00:00.000000",
		UpdateTime:             "9999-12-31 23:59:59.999999",
	},
}

func (h *Handle) GetQuestionsHandler(c *gin.Context) {
	questionID := c.Param("questionID")

	if _, err := uuid.Parse(questionID); err != nil {
		log.Println("uuid.Parse(): ", err)
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	//DBからもらう
	for _, v := range questions {
		if v.QuestionID.String() == questionID {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, "Not Found")
}

func (h *Handle) GetAllQuestionsHandler(c *gin.Context) {
	//DBからもらう

	c.JSON(http.StatusOK, questions)
}

func (h *Handle) GetCollectionHandler(c *gin.Context) {
	collectionID := c.Param("collectionID")

	if _, err := uuid.Parse(collectionID); err != nil {
		log.Println("uuid.Parse(): ", err)
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	// DBからもらう
	for _, v := range collections {
		if v.CollectionID.String() == collectionID {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, "Not Found")
}

func (h *Handle) GetAllCollectionsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, collections)
}
