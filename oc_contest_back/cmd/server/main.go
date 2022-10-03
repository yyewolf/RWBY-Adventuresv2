package main

import (
	"os"
	"os/signal"
	"rwby-adventures/models"
	"rwby-adventures/oc_contest_back/routes"
	"syscall"
)

func main() {
	routes.Serve()

	test := models.Submission{
		SubmissionID: "1",
		DiscordID:    "1",
		Name:         "This is a long title to test line break",
		Color:        "Red",
		ShortDesc:    "This is a short description but it is still a bit long.",
		LongDesc:     "This is a long description.",
		Author:       "@Someone#0000",
		Votes:        50,
		Files: []*models.SubmissionFile{
			{
				FileID:       "1",
				SubmissionID: "1",
				Name:         "Image name.png",
				URI:          "image file.png",
				Path:         "image file.png",
			},
		},
	}
	test.Save()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
