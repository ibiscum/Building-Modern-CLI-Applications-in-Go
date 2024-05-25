package command

import (
	"fmt"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter03/audiofile/internal/interfaces"
)

type Parser struct {
	commands []interfaces.Command
}

func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string) error {
	if len(args) < 1 {
		help()
		return nil
	}

	subcommand := args[0]
	for _, cmd := range p.commands {
		if cmd.Name() == subcommand {
			err := cmd.ParseFlags(args[1:])
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
