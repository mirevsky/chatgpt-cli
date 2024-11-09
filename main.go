package main

import (
	"github.com/mirevsky/chatgpt-cli/cmd"
	"os"
)

func main() {
	err := cmd.NewRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
