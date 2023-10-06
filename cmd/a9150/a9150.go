/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package a9150

import (
	"github.com/spf13/cobra"
)

// Define known commands for onkyo a9150
var (
	A9150Predefined = map[string]uint16{
		"volu":   0x2,
		"vold":   0x3,
		"pwr":    0x4,
		"mute":   0x5,
		"d1":     0x20,
		"unmute": 0xD8,
		"d2":     0xE0,
		"d3":     0x170,
		"srcu":   0xD5,
		"srcd":   0xD6,
	}
)

// A9150Cmd represents the a9150 command
var A9150Cmd = &cobra.Command{
	Use:   "a9150",
	Short: "a9150 is pallet for onkyo stereo amp A-9150 predefined commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubcommands() {
	A9150Cmd.AddCommand(pwrCmd)

	A9150Cmd.AddCommand(muteCmd)
	A9150Cmd.AddCommand(unmuteCmd)

	A9150Cmd.AddCommand(volupCmd)
	A9150Cmd.AddCommand(voldownCmd)

	A9150Cmd.AddCommand(srcupCmd)
	A9150Cmd.AddCommand(srcdownCmd)

	A9150Cmd.AddCommand(d1Cmd)

}

func init() {
	addSubcommands()
}
