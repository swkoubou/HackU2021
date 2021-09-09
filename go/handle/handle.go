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
		QuestionTag:  []string{"音楽", "音楽記号", "コード理論"},
		QuestionType: "4taku",
		CreateTime:   "9999-12-31 23:59:59.999999",
		UpdateTime:   "",
		QuestionBody: "C,G,Am,Em,F,C,F,Gのコード進行は?",
		Values:       []string{"カノン進行", "J-POP王道進行", "小室進行", "イチロクニーゴー"},
		Answers:      []string{"カノン進行"},
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionTag:  []string{"音楽", "音楽用語"},
		QuestionType: "anaume",
		CreateTime:   "9999-12-31 23:59:59.999999",
		UpdateTime:   "",
		QuestionBody: "Poco Allegro con affettoの意味は?",
		Values:       []string{"やや[]く[]こめて"},
		Answers:      []string{"はや", "愛情"},
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionTag:  []string{"音楽", "音楽用語"},
		QuestionType: "anaume",
		CreateTime:   "9999-12-31 23:59:59.999999",
		UpdateTime:   "",
		QuestionBody: "più animato con passioneの意味は?",
		Values:       []string{"[]的にやや[]と[]く"},
		Answers:      []string{"情熱", "生き生きと", "はや"},
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionTag:  []string{"プログラミング言語", "C"},
		QuestionType: "4taku",
		CreateTime:   "9999-12-31 23:59:59.999999",
		UpdateTime:   "",
		QuestionBody: "intは何型?",
		Values:       []string{"整数型", "文字型", "実数型", "構造体"},
		Answers:      []string{"整数型"},
	},
}

var collections []collection.Collection = []collection.Collection{
	{
		CollectionID:           newUUID(),
		CollectionName:         "foo",
		CollectionDescripition: "テスト用のコレクション",
		Author:                 accounts[2],
		Questions:              questions,
		CreateTime:             "9999-12-31 23:59:59.999999",
		UpdateTime:             "",
	},
	{
		CollectionID:           newUUID(),
		CollectionName:         "bar",
		CollectionDescripition: "ンョシクレコの用トステ",
		Author:                 accounts[2],
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
