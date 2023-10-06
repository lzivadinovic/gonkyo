/*
Copyright © 2023 Lazar Zivadinovic <git@lzivadinovic.com>
*/
package a9150

import (
	"fmt"
	"github.com/spf13/cobra"
)

// pwrCmd represents the pwr command
var pwrCmd = &cobra.Command{
	Use:   "pwr",
	Short: "Power on/off",
	Long: `Send power on or off command to you amp.
This command is toggle, it does not care about your current state!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pwr called")
	},
}

func init() {
}
