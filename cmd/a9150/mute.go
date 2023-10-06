/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package a9150

import (
	"fmt"
	"github.com/lzivadinovic/gonkyo/wrapper"
	"github.com/spf13/cobra"
)

// muteCmd represents the mute command
var muteCmd = &cobra.Command{
	Use:   "mute",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mute called")
		wrapper.SendCommand(A9150Predefined["mute"])
	},
}

func init() {
}
