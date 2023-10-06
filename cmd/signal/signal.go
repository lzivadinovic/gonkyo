/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package signal

import (
	"fmt"
	"github.com/spf13/cobra"
)

// SignalCmd represents the signal command
var SignalCmd = &cobra.Command{
	Use:   "signal",
	Short: "Send custom message (hex/uint) to Onkyo-RI interface",
	Long: `You can use this command to send custom message to your onkyo-RI interface.
You can use either hex or uint encoded values for your command.
Program will then send blablabla`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signal called")
		//wrapper.SendCommand(230)
	},
}

func init() {
	//cmd.rootCmd.AddCommand(SignalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
