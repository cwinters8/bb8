package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/bb8"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor("BB-3A94")
	bb8 := bb8.NewDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := getRandColor()
			g := getRandColor()
			b := getRandColor()
			bb8.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("bb", []gobot.Connection{bleAdaptor}, []gobot.Device{bb8}, work)
	robot.Start()
}

func getRandColor() uint8 {
	return uint8(gobot.Rand(255))
}
