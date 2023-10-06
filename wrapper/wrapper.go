package wrapper

import (
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

var (
	// gpio 25 -> physical pin 37
	pin = rpio.Pin(25)
)

// Creates header signal for onkyo-ri message
// Header signal consists of 3 ms of high pull, and 1 ms of low
func headerSignal() {
	pin.High()
	time.Sleep(time.Millisecond * 3)
	pin.Low()
	time.Sleep(time.Millisecond * 1)
}

// Creates signal for one message bit
// 1 is encoded as 1 ms of high pull, and 2 ms of low pull
// 0 is encoded as 1ms of high pull, and 1 ms of low pull
func bitPull(b uint16) {
	pin.High()
	time.Sleep(time.Millisecond * 1)
	if b == 1 {
		pin.Low()
		time.Sleep(time.Millisecond * 1)
	}
	pin.Low()
	time.Sleep(time.Millisecond * 1)

}

// Creates trailer signal which is sent after message signal
// Trailer signal consist of 1ms of high pull, and at leas 20ms of low pull
func trailerSignal() {
	pin.High()
	time.Sleep(time.Millisecond * 1)
	pin.Low()
	time.Sleep(time.Millisecond * 70)
}

// Combines header, command and trail signal in full command
func SendCommand(command uint16) {
	// set pin to output
	// -- REFACTOR --
	pin.Output()
	// -- REFACTOR --
	headerSignal()
	for i := 0; i < 12; i++ {
		bit := command >> (11 - i) & uint16(0x1)
		bitPull(bit)
	}
	trailerSignal()
}
