/*
Copyright © 2023 Lazar Zivadinovic <git@lzivadinovic.com>
*/
package signal

import (
	"fmt"
	"github.com/lzivadinovic/gonkyo/wrapper"
	"github.com/spf13/cobra"
	"strconv"
)

// SignalCmd represents the signal command
var SignalCmd = &cobra.Command{
	Use:   "signal [<uInt|hex>]",
	Short: "Send custom message (hex/uInt) to Onkyo-RI interface",
	Long: `You can use this command to send custom message to your onkyo-RI interface.
Command can be hex or uInt, eg sending "0xD5" or "213" will yield equivalent result.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get arg and pars it as uint16
		arg, err := strconv.ParseUint(args[0], 0, 0)
		comm := uint16(arg)

		if err != nil {
			fmt.Println("Something went wrong while parsing command!")
		}

		for i := 0; i < repeatTimes; i++ {
			wrapper.SendCommand(comm)
		}

	},
}

var repeatTimes int

func init() {
	SignalCmd.Flags().IntVarP(&repeatTimes, "times", "t", 1, "times to repeat the command")
}
