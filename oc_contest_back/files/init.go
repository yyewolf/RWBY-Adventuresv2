package files

import "os"

func init() {
	// Make sure the upload folder exists
	if _, err := os.Stat("/upload"); os.IsNotExist(err) {
		err := os.Mkdir("/upload", 0755)
		if err != nil {
			panic(err)
		}
	}
}
