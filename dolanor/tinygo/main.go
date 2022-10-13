package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	blink()
}

type LED struct {
	led machine.Pin
	on  bool
}

func (l *LED) On() {
	l.led.High()
	l.on = true
}

func (l *LED) Off() {
	l.led.Low()
	l.on = false
}

func (l *LED) Toggle() {
	if !l.on {
		l.On()
	} else {
		l.Off()
	}
	//toggleLED(&l.led, &l.on)
}

func blink() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledCtrl := make(chan bool, 10)

	l := LED{
		led: led,
	}

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
