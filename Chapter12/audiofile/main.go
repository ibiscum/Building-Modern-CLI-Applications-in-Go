package main

import (
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter12/audiofile/cmd"
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter12/audiofile/utils"
)

func main() {
	cmd.Configure()
	utils.InitCLILogger()
	cmd.Execute()
}
