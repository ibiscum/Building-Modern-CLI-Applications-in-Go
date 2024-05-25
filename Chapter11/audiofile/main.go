package main

import (
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter11/audiofile/cmd"
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter11/audiofile/utils"
)

func main() {
	cmd.Configure()
	utils.InitCLILogger()
	cmd.Execute()
}
