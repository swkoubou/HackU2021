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
	c.JSON(http.StatusOK, &param{
		QuestionID: "get all questions",
	})
}

func PostQuestionHandler(c *gin.Context) {

}

func GetCollectionHandler(c *gin.Context) {
	collectionID := c.Param("collectionID")

	if _, err := uuid.Parse(collectionID); err != nil {
		log.Println("uuid.Parse(): ", err)
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	// DBからもらう

	c.JSON(http.StatusOK, &param{
		QuestionID: collectionID,
	})
}

func GetAllCollectionsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, &param{
		QuestionID: "get all collection",
	})
}
