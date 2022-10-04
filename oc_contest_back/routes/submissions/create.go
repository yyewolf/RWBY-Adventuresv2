package submissions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/yyewolf/goth"
)

var maxMemory int64 = 8 << 20 // 8 MB

func Create(c *gin.Context) {
	// Check how many submission the user has
	u := c.MustGet("User").(*goth.User)
	s := models.GetUserSubmissions(u.UserID)
	if len(s) >= 15 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You have already submitted 5 characters."})
		return
	}

	// Read the JSON Body first
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(maxMemory); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "There has been an error with the server : C_001"})
			return
		}
	}

	// Check the amount of uploaded files
	f, _ := c.Request.MultipartForm.File["files"]
	if len(f) >= 6 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You cannot upload more than 6 files (including icon)."})
		return
	}
	if len(f) < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "You have to upload at least an icon."})
		return
	}

	jsonFiles, ok := c.Request.MultipartForm.Value["data"]
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"error": "There has been an error with the server : C_002"})
		return
	}
	if len(jsonFiles) != 1 {
		c.AbortWithStatusJSON(400, gin.H{"error": "There has been an error with the server : C_003"})
		return
	}
	data := jsonFiles[0]
	submission := &models.Submission{}
	err := json.Unmarshal([]byte(data), submission)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "There has been an error with the server : C_004"})
		return
	}

	// Check if the submission is valid
	if len(submission.Name) <= 5 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Minimum character name length is 5."})
		return
	}
	if len(submission.Name) >= 30 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Maximum character name length is 30."})
		return
	}
	if len(submission.ShortDesc) <= 10 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Minimum short description length is 10."})
		return
	}
	if len(submission.ShortDesc) >= 150 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Maximum short description length is 150."})
		return
	}
	if len(submission.LongDesc) <= 300 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Minimum long description length is 300."})
		return
	}
	if len(submission.LongDesc) >= 5000 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Maximum long description length is 1000."})
		return
	}

	// Generate submission id
	ID, err := uuid.NewV4()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "There has been an error with the server : C_005"})
		return
	}

	// Read all files
	f, ok = c.Request.MultipartForm.File["files"]
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"error": "There has been an error with the server : C_006"})
		return
	}

	var allFiles []*models.SubmissionFile
	for i, file := range f {
		// Get the path
		newName, err := uuid.NewV4()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "There has been an error with the server : C_007"})
			return
		}

		filename := fmt.Sprintf("%s_%s", newName, file.Filename)
		path := fmt.Sprintf("./upload/%s", filename)

		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "There has been an error with the server : C_008"})
			return
		}

		newFile := &models.SubmissionFile{
			SubmissionID: ID.String(),
			FileID:       i,
			Name:         file.Filename,
			URI:          "files/get/" + filename,
			Path:         path,
		}
		allFiles = append(allFiles, newFile)
	}

	submission.Files = allFiles
	user := c.MustGet("User").(*goth.User)

	submission.SubmissionID = ID.String()
	submission.DiscordID = user.UserID
	submission.Author = fmt.Sprintf("@%s#%s", user.RawData["username"], user.RawData["discriminator"])
	submission.Votes = []*models.SubmissionVote{}

	submission.Save()

	c.JSON(200, gin.H{
		"submission": submission,
	})
}
