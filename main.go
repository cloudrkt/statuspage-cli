package main

import (
	"log"

	"github.com/cloudrkt/statuspage-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd.Execute()

	err := doc.GenMarkdownTree(cmd.RootCmd, "/tmp")
	if err != nil {
		log.Fatal(err)
	}
}
