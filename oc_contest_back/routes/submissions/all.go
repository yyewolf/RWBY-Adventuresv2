package submissions

import (
	"rwby-adventures/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {
	data := c.Param("page")
	page, err := strconv.Atoi(data)
	if err != nil {
		page = 0
	}
	count := models.GetAmountOfSubmissions()

	maxPage := count / 10
	if page > maxPage {
		page = maxPage
	}
	if page < 0 {
		page = 0
	}

	s := models.GetSubmissions(10, page*10)

	c.JSON(200, gin.H{
		"page":        page,
		"max_page":    maxPage,
		"submissions": s,
	})
}
