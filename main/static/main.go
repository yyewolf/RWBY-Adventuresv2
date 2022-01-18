package static

import (
	"embed"
	"io/fs"
)

//go:embed www
var webFS embed.FS
var WebFS, _ = fs.Sub(webFS, "www")

//go:embed www/assets
var assets embed.FS
var Assets, _ = fs.Sub(assets, "www/assets")

//go:embed database
var DatabaseFS embed.FS

func init() {
	//
}
