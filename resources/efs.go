package resources

import "embed"

//go:embed "public" "views"
var EmbeddedFiles embed.FS
