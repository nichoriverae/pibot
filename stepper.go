package main

import (
	"time"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	servo := gpio.NewServoDriver(r, os.Args[1])

	work := func() {
		servo.Center()
		time.Sleep(1000 * time.Millisecond)
		servo.Max()
		time.Sleep(1000 * time.Millisecond)
		servo.Min()
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{r},
		[]gobot.Device{servo},
		work,
	)

	robot.Start()
}
