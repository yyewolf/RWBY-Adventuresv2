package main

import (
	"fmt"
	"os"
	"os/signal"
	"rwby-adventures/models"
	"rwby-adventures/oc_contest_back/routes"
	"syscall"
)

func main() {
	routes.Serve()

	for i := 0; i < 10; i++ {
		test := models.Submission{
			SubmissionID: fmt.Sprintf("%d", i),
			DiscordID:    "1",
			Name:         "This is a long title to test line break",
			Color:        "Red",
			ShortDesc:    "This is a short description but it is still a bit long.",
			LongDesc:     "This is a long description.",
			Author:       "@Someone#0000",
			Files: []*models.SubmissionFile{
				{
					FileID:       1,
					SubmissionID: fmt.Sprintf("%d", i),
					Name:         "A really cool icon",
					URI:          "files/get/42d9c39a-ffda-4bf2-8c3c-2993cb7c428e_xero.png",
					Path:         "./upload/42d9c39a-ffda-4bf2-8c3c-2993cb7c428e_xero.png",
				},
			},
		}
		test.Save()
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
