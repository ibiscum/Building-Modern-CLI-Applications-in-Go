package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/chapter03/cmd/cli/command"
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/chapter03/internal/interfaces"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewGetCommand(client),
		command.NewUploadCommand(client),
		command.NewListCommand(client),
	}
	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
