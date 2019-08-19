package main

import (
   "fmt"

   "gobot.io/x/gobot"
   "gobot.io/x/gobot/drivers/gpio"
   "gobot.io/x/gobot/platforms/raspi"
)

func main() {
   r := raspi.NewAdaptor()

   sensor := gpio.NewPIRMotionDriver(r, "7")
   led := gpio.NewLedDriver(r, "12")

   work := func() {
      sensor.On(gpio.MotionDetected, func(data interface{}) {
         fmt.Println(gpio.MotionDetected)
         led.Toggle()
	fmt.Println(led.State())
	   })
	}

   robot := gobot.NewRobot("motionBot",
      []gobot.Connection{r},
      []gobot.Device{sensor, led},
      work,
   )

   robot.Start()
}
