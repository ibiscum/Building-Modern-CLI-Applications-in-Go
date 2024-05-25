package main

import (
	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter10a/dashboard"
)

func main() {
	// Guiding users with prompts
	//survey.UserExperience()

	// Building a useful dashboard - Learn about termdash
	err := dashboard.Basic()
	if err != nil {
		panic(err)
	}

	err = dashboard.BinaryTreeWithStyle()
	if err != nil {
		panic(err)
	}
}
