package submissions

import (
	"rwby-adventures/config"
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
	var submissions []*models.Submission

	config.Database.
		Model(&models.Submission{}).
		Joins("JOIN \"submission_votes\" ON \"submission_votes\".\"submission_id\" = \"submissions\".\"id\"").
		Group("\"submissions\".\"id\", \"submissions\".\"discord_id\"").
		Order("COUNT(\"submission_votes\".\"submission_id\") DESC").
		Limit(10).
		Find(&submissions)

	for _, s := range submissions {
		s.Votes = models.GetSubmissionVotes(s.SubmissionID)
		s.Files = models.GetSubmissionFiles(s.SubmissionID)
		s.Icon = s.Files[0]
	}

	c.JSON(200, gin.H{
		"submissions": submissions,
	})
}
