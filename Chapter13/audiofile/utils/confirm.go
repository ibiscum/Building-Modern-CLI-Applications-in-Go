package utils

import (
	"log"

	survey "github.com/AlecAivazis/survey/v2"
)

func Confirm(confirmationText string) bool {
	confirmed := false
	prompt := &survey.Confirm{
		Message: confirmationText,
	}
	err := survey.AskOne(prompt, &confirmed)
	if err != nil {
		log.Fatal(err)
	}
	return confirmed
}
