package submissions

import (
	"encoding/json"
	"io"
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// Read the JSON Body first
	jsonFile, err := c.FormFile("data")
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid form"})
		return
	}
	jsonData, err := jsonFile.Open()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid form"})
		return
	}
	data, err := io.ReadAll(jsonData)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid form"})
		return
	}
	submission := &models.Submission{}
	err = json.Unmarshal(data, submission)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid form"})
		return
	}
	jsonData.Close()

}
