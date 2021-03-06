// Copyright 2020 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/embeddedgo/kendryte/devboard/maixbit/board/leds"
)

func blink(led leds.LED, delay time.Duration) {
	//runtime.LockOSThread()
	for {
		led.SetOn()
		time.Sleep(delay)
		led.SetOff()
		time.Sleep(delay)
	}
}

func main() {
	go blink(leds.Red, time.Second*2)
	go blink(leds.Green, time.Second/3)
	blink(leds.Blue, time.Second/11)
}
