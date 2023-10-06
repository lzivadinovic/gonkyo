/*
Copyright © 2023 Lazar Zivadinovic <git@lzivadinovic.com>
*/
package a9150

import (
	"fmt"

	"github.com/spf13/cobra"
)

// d1Cmd represents the d1 command
var d1Cmd = &cobra.Command{
	Use:   "d1",
	Short: "Switch to D1 source",
	Long:  `Switches input to Digital 1 (coax 1).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d1 called")
	},
}

func init() {
}
