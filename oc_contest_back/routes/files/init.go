package files

import (
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	// Make sure the upload folder exists
	if _, err := os.Stat("./upload"); os.IsNotExist(err) {
		err := os.Mkdir("./upload", 0755)
		if err != nil {
			panic(err)
		}
	}
}

func ServeFiles(path *gin.RouterGroup) {
	subpath := path.Group("/files")

	subpath.GET("/get/:file", Get)
}
