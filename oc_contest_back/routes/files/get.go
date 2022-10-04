package files

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// Get the file
	filename := c.Param("file")

	data, err := os.ReadFile("./upload/" + filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return
	}

	// We reply image for images type and octet stream for other types
	if strings.HasSuffix(filename, "png") {
		c.Writer.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(filename, "jpg") {
		c.Writer.Header().Set("Content-Type", "image/jpg")
	} else if strings.HasSuffix(filename, "jpeg") {
		c.Writer.Header().Set("Content-Type", "image/jpeg")
	} else {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
	}
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	c.Writer.Write(data)
}
