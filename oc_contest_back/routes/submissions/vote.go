package submissions

import (
	"rwby-adventures/config"
	"rwby-adventures/models"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

func Vote(c *gin.Context) {
	id := c.Param("id")
	s := models.GetSubmission(id)
	if s == nil {
		c.JSON(404, gin.H{
			"error": "Submission not found",
		})
		return
	}
	u := c.MustGet("User").(*goth.User)
	if u.UserID == s.DiscordID {
		c.JSON(403, gin.H{
			"error": "You are not allowed to vote for your own submission",
		})
		return
	}

	// Count amount of user votes
	amount := models.GetAmountOfVote(u.UserID)
	if amount >= 5 {
		c.JSON(403, gin.H{
			"error": "You are not allowed to vote more than 5 times",
		})
		return
	}

	// Check if user already voted
	v := &models.SubmissionVote{}
	config.Database.Where("submission_id = ? AND discord_id = ?", s.SubmissionID, u.UserID).First(&v)
	if v.SubmissionID != "" {
		c.JSON(403, gin.H{
			"error": "You are not allowed to vote more than once",
		})
		return
	}

	vote := &models.SubmissionVote{
		SubmissionID: s.SubmissionID,
		DiscordID:    u.UserID,
	}
	vote.Save()

	c.JSON(200, gin.H{
		"status": "Your vote has been counted",
	})
}
