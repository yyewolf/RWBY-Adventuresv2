package files

import "os"

func ReadFile(ID string) ([]byte, error) {
	return os.ReadFile("/upload/" + ID)
}
