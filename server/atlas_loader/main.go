package main

import (
	"fmt"
	"io"
	"os"
	"usdrive/db"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	customTypes := `CREATE TYPE file_status AS ENUM ('pending_upload', 'active', 'error');`

	stmts, err := gormschema.New("postgres").Load(
		&db.User{},
		&db.File{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, customTypes)
	io.WriteString(os.Stdout, stmts)
}
