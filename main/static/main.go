package static

import "embed"

//go:embed www
var Box embed.FS

//go:embed database
var CharBox embed.FS

func init() {
	Box.Open("d")
}
