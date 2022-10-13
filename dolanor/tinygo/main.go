package main

import (
	"fmt"
	"time"

	"github.com/tuxago/go/dolanor/tinygo/hw/led"
)

func main() {
	blink()
}

func blink() {
	ledCtrl := make(chan bool, 10)

	l := led.New()

	go func() {
		for {
			select {
			case <-ledCtrl:
				l.Toggle()
			}
		}
	}()
	var ledState bool

	for {
		time.Sleep(time.Millisecond * 100)

		ledState = !ledState
		fmt.Println(ledState)
		//toggleLED(led, ledState)

		ledCtrl <- ledState

		// println("[STM32] blink")
	}
}

//func toggleLED(led *machine.Pin, state *bool) {
//	*state = !*state
//	if *state {
//		led.Low()
//	} else {
//		led.High()
//	}
//}
