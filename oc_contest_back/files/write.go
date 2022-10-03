package files

import "os"

func WriteID(ID string, data []byte) error {
	return os.WriteFile("./upload/"+ID, data, 0644)
}
