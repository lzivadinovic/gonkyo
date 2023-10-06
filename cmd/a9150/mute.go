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
	Short: "Mutes output",
	Long:  `Mutes output. This command is not toggle, you will need to use unmute!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mute called")
		wrapper.SendCommand(A9150Predefined["mute"])
	},
}

func init() {
}
