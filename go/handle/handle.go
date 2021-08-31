package handle

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type param struct {
	QuestionID string `json:"questionID"`
}

func GetQuestionsHandler(c *gin.Context) {
	questionID := c.Param("questionID")

	if _, err := uuid.Parse(questionID); err != nil {
		log.Println("uuid.Parse(): ", err)
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	c.JSON(http.StatusOK, &param{
		QuestionID: questionID,
	})
}

func GetAllQuestionsHandler(c *gin.Context) {
	log.Println("hoge")
	c.JSON(http.StatusOK, &param{
		QuestionID: "all!",
	})
}

func PostQuestionHandler(c *gin.Context) {

}
