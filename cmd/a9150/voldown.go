/*
Copyright © 2023 Lazar Zivadinovic <git@lzivadinovic.com>
*/
package a9150

import (
	"fmt"
	"github.com/spf13/cobra"
)

// voldownCmd represents the voldown command
var voldownCmd = &cobra.Command{
	Use:   "voldown",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("voldown called")
	},
}

func init() {
	voldownCmd.Flags().IntVarP(&repeatTimes, "times", "t", 5, "times to repeat the command")
}
