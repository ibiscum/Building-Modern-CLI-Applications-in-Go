//go:build linux

package cmd

import (
	"log"
	"os/exec"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter13/audiofile/utils"
	"github.com/pterm/pterm"
)

func play(audiofilePath string, verbose bool) (int, error) {
	cmd := exec.Command("aplay", audiofilePath)
	if err := cmd.Start(); err != nil {
		return 0, utils.Error("\n  starting aplay command: %v", err, verbose)
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsaTTY() && verbose {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return 0, utils.Error("\n  running aplay command: %v", err, verbose)
	}
	if utils.IsaTTY() && verbose {
		err = spinnerInfo.Stop()
		if err != nil {
			log.Fatal(err)
		}
	}
	return cmd.Process.Pid, nil
}
