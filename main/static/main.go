package static

import "embed"

//go:embed www
var Box embed.FS

func init() {
	Box.Open("d")
}
