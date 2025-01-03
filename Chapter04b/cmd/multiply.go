/*
Copyright © 2023 Marian Montagnino <mmontagnino@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter04b/storage"
	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply number",
	Short: "Multiply value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("only accepts a single argument")
			return
		}
		if len(args) == 0 {
			fmt.Println("command requires input value")
			return
		}
		floatVal, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("unable to parse input[%s]: %v", args[0], err)
			return
		}
		value = storage.GetValue()
		value *= floatVal
		err = storage.SetValue(value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%f\n", value)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}
