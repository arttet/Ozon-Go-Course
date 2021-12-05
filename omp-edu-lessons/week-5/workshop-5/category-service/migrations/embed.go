package migrations

import "embed"

//go:embed *.sql
var EmbedFS embed.FS
