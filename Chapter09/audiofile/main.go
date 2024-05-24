package main

import (
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter09/audiofile/cmd"
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter09/audiofile/utils"
)

func main() {
	cmd.Configure()
	utils.InitCLILogger()
	cmd.Execute()
}
