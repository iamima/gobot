package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

var Master *gobot.Master = gobot.GobotMaster()

func TurnBlue(params map[string]interface{}) bool {
	sphero := Master.FindRobotDevice(params["robotname"].(string), "sphero")
	gobot.Call(sphero.Driver, "SetRGB", uint8(0), uint8(0), uint8(255))
	return true
}

func main() {
	gobot.Api(Master)

	spheros := map[string]string{
		"Sphero-BPO": "/dev/rfcomm0",
	}

	for name, port := range spheros {
		spheroAdaptor := new(gobotSphero.SpheroAdaptor)
		spheroAdaptor.Name = "sphero"
		spheroAdaptor.Port = port

		sphero := gobotSphero.NewSphero(spheroAdaptor)
		sphero.Name = "sphero"
		sphero.Interval = "0.5s"

		work := func() {
			sphero.SetRGB(uint8(255), uint8(0), uint8(0))
		}

		Master.Robots = append(Master.Robots, &gobot.Robot{
			Name:        name,
			Connections: []gobot.Connection{spheroAdaptor},
			Devices:     []gobot.Device{sphero},
			Work:        work,
			Commands:    map[string]interface{}{"TurnBlue": TurnBlue},
		})
	}

	Master.Start()
}
