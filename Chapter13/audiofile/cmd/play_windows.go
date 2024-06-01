//go:build windows

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter13/audiofile/utils"
)

func play(audiofilePath string, verbose bool) (int, error) {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return 0, utils.Error("\n  starting start command: %v", err, verbose)
	}
	if verbose {
		fmt.Println("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return 0, utils.Error("\n  running start command: %v", err, verbose)
	}
	return cmd.Process.Pid, nil
}
