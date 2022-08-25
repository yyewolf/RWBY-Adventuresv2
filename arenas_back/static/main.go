package static

import (
	"embed"
	"io/fs"
)

//go:embed www
var webFS embed.FS
var WebFS, _ = fs.Sub(webFS, "www")

//go:embed www/css www/js www/fonts
var assets embed.FS
var Assets, _ = fs.Sub(assets, "www")

func init() {
	//
}
