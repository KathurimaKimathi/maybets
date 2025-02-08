package assets

import "embed"

//go:embed migrations/*.sql
var DBMigrations embed.FS
