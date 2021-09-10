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
		QuestionName: "コード進行1",
		QuestionTag:  []string{"音楽", "音楽記号", "コード理論"},
		QuestionType: "4taku",
		CreateTime:   "2021-08-31 15:30:57.499391",
		UpdateTime:   "2021-08-31 18:33:24.299391",
		QuestionBody: "C,G,Am,Em,F,C,F,Gのコード進行は?",
		Values:       []string{"カノン進行", "J-POP王道進行", "小室進行", "イチロクニーゴー"},
		Answers:      []string{"カノン進行"},
		Color:        "#A2F1FF",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionName: "音楽用語1",
		QuestionTag:  []string{"音楽", "音楽用語"},
		QuestionType: "anaume",
		CreateTime:   "2021-08-31 15:36:27.135791",
		UpdateTime:   "2021-08-31 19:32:24.359291",
		QuestionBody: "Poco Allegro con affettoの意味は?",
		Values:       []string{"やや[]く[]こめて"},
		Answers:      []string{"はや", "愛情"},
		Color:        "#FFB4C5",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionName: "音楽用語2",
		QuestionTag:  []string{"音楽", "音楽用語"},
		QuestionType: "anaume",
		CreateTime:   "2021-09-01 02:15:30.321691",
		UpdateTime:   "2021-09-01 04:18:21.321691",
		QuestionBody: "più animato con passioneの意味は?",
		Values:       []string{"[]的にやや[]と[]く"},
		Answers:      []string{"情熱", "生き生きと", "はや"},
		Color:        "#9AF7B6",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[0],
		QuestionTag:  []string{"プログラミング言語", "C"},
		QuestionName: "C言語1",
		QuestionType: "4taku",
		CreateTime:   "2021-09-01 05:32:47.401391",
		UpdateTime:   "2021-09-01 10:24:25.401391",
		QuestionBody: "intは何型?",
		Values:       []string{"整数型", "文字型", "実数型", "構造体"},
		Answers:      []string{"整数型"},
		Color:        "#FFE796",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "ミラーリングについて",
		QuestionTag:  []string{"資格", "基本情報", "IT"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-01 15:20:12.821491",
		UpdateTime:   "2021-09-01 22:34:51.821491",
		QuestionBody: "ミラーリングを用いることで，信頼性を高め障害発生時にデータ復元を行う方式はどれか？",
		Values:       []string{"RAID1", "RAID2", "RAID3", "RAID4"},
		Answers:      []string{"RAID1"},
		Color:        "#D4D9F2",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "計算誤差について",
		QuestionTag:  []string{"資格", "基本情報", "IT"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-01 15:25:32.135491",
		UpdateTime:   "2021-09-01 22:57:12.135491",
		QuestionBody: "値がほぼ等しい浮動小数点同士の計算により，有効桁数が減ってしまうのはどれか？",
		Values:       []string{"桁落ち", "丸め誤差", "情報落ち", "桁あふれ誤差"},
		Answers:      []string{"桁落ち"},
		Color:        "#A2F1FF",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "記憶方式について",
		QuestionTag:  []string{"資格", "基本情報", "IT"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-01 18:49:32.632182",
		UpdateTime:   "2021-09-01 20:32:47.632182",
		QuestionBody: "仮想記憶空間と実記憶空間を，固定長の領域に区切り対応させる管理方式はどれか？",
		Values:       []string{"動的再配置", "メモリインタリーブ", "ページング方式", "ブロック化"},
		Answers:      []string{"ページング方式"},
		Color:        "#FFB4C5",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "攻撃手法について",
		QuestionTag:  []string{"資格", "基本情報", "IT"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-01 23:20:18.311491",
		UpdateTime:   "2021-09-02 22:43:32.311491",
		QuestionBody: "パスワードを解読する手法として，総当たりで調べるのはどれか？",
		Values:       []string{"線形解読法", "ブルートフォース攻撃", "差分解読法", "関連鍵攻撃"},
		Answers:      []string{"ブルートフォース攻撃"},
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英文法",
		QuestionTag:  []string{"英語", "英文法", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-02 22:19:32.231891",
		UpdateTime:   "2021-09-02 23:15:21.231891",
		QuestionBody: "The International Peace Conference ___ as planned of Saturday.",
		Values:       []string{"concluded", "held", "achieved", "invited"},
		Answers:      []string{"concluded"},
		Color:        "#9AF7B6",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英文法",
		QuestionTag:  []string{"英語", "英文法", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-03 19:13:32.102491",
		UpdateTime:   "2021-09-03 21:45:33.102491",
		QuestionBody: "The new guidelines regarding jury duty should be ___ to strictly.",
		Values:       []string{"respected", "adhered", "willing", "set"},
		Answers:      []string{"adhered"},
		Color:        "#FFE796",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英文法",
		QuestionTag:  []string{"英語", "英文法", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-04 12:43:42.502491",
		UpdateTime:   "2021-09-04 19:21:11.502491",
		QuestionBody: "Our new facility ___ right across from the city's new convention center.",
		Values:       []string{"is located", "will locate", "is to locate", "had located"},
		Answers:      []string{"is located"},
		Color:        "#D4D9F2",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英文法",
		QuestionTag:  []string{"英語", "英文法", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-06 18:13:29.319491",
		UpdateTime:   "2021-09-06 23:43:11.319491",
		QuestionBody: "A sudden drop in sales would ___ cuts in production and employment.",
		Values:       []string{"necessitate", "necessary", "necessity", "necessarily"},
		Answers:      []string{"necessitate"},
		Color:        "#A2F1FF",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英単語",
		QuestionTag:  []string{"英語", "英単語", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-07 20:12:12.812391",
		UpdateTime:   "2021-09-08 18:15:29.319491",
		QuestionBody: "protractor の意味は？",
		Values:       []string{"コンパス", "分度器", "消しゴム", "黒板消し"},
		Answers:      []string{"分度器"},
		Color:        "#FFB4C5",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英単語",
		QuestionTag:  []string{"英語", "英単語", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-08 03:41:22.352390",
		UpdateTime:   "2021-09-08 14:22:45.352390",
		QuestionBody: "apprentice の意味は？",
		Values:       []string{"人類", "謝る", "初心者", "食欲"},
		Answers:      []string{"初心者"},
		Color:        "#9AF7B6",
	},
	{
		QuestionID:   newUUID(),
		Author:       accounts[1],
		QuestionName: "英単語",
		QuestionTag:  []string{"英語", "英単語", "TOEIC"},
		QuestionType: "4taku",
		CreateTime:   "2021-09-09 23:22:42.292391",
		UpdateTime:   "2021-09-10 11:34:45.292391",
		QuestionBody: "photosynthesis の意味は？",
		Values:       []string{"写真", "定説", "映像", "光合成"},
		Answers:      []string{"光合成"},
		Color:        "#FFE796",
	},
}

var collections []collection.Collection = []collection.Collection{
	{
		CollectionID:          newUUID(),
		CollectionName:        "music",
		CollectionDescription: "音楽関連の問題集",
		Author:                accounts[0],
		Questions:             questions[0:2],
		CreateTime:            "2021-09-02 22:21:31.231891",
		UpdateTime:            "2021-09-02 23:32:11.133891",
	},
	{
		CollectionID:          newUUID(),
		CollectionName:        "基本情報",
		CollectionDescription: "基本情報技術者試験勉強用の問題集",
		Author:                accounts[1],
		Questions:             questions[4:7],
		CreateTime:            "2021-09-02 22:29:12.431891",
		UpdateTime:            "2021-09-02 22:35:32.231891",
	},
	{
		CollectionID:          newUUID(),
		CollectionName:        "英文法",
		CollectionDescription: "英文法に関する問題集",
		Author:                accounts[2],
		Questions:             questions[8:11],
		CreateTime:            "2021-09-03 12:21:42.831890",
		UpdateTime:            "2021-09-04 22:19:36.291891",
	},
	{
		CollectionID:          newUUID(),
		CollectionName:        "英単語",
		CollectionDescription: "英単語に関する問題集",
		Author:                accounts[3],
		Questions:             questions[12:14],
		CreateTime:            "2021-09-02 18:38:12.831891",
		UpdateTime:            "2021-09-05 16:19:49.539891",
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
