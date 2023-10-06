package main

import (
	"github.com/lzivadinovic/gonkyo/cmd"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
)

func main() {
	// Setup gpio handler
	if err := rpio.Open(); err != nil {
		log.Panic("Could not open /dev/gpiomem or /dev/mem; Try running with sudo or adding your user to `gpio` group!")
		os.Exit(1)
	}

	//// unmap gpio memory when done
	defer rpio.Close()

	// Setup cobra cli
	cmd.Execute()

}
