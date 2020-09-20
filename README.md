# Golang SimConnect binding for FS2020 :airplane:

_This library is in developement and I have write little documentation. Please considerate this library as a testing version. It is possible non backwards compatible in the future version. But you can enjoy :partying_face: and good fly !_

For more information on how to use this library, please read [example_test.go](https://github.com/micmonay/simconnect/blob/master/example_test.go).

With this library you can in simulator:
- Read SimVar. (ex: Altitude, Longitude, Latitude, AP master status, Fuel, Engine...)
- Write SimVar. All the SimVar have not a possibility to be written show SimVar.Settable
- Send SimEvent for change Throttle or other
- Receive system event (ex: When the aircraft crash). _Not all implemented_ 
- Show text in the screen on the simulator

## A simple example of how to use this library
```go
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	sim "github.com/micmonay/simconnect"
)

func main() {
	sc, err := sim.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	sc.SetLoggerLevel(sim.LogInfo) // It is better if you the set before connect
	c, err := sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	<-c // Wait connection confirmation
	cSimVar, err := sc.ConnectToSimVar(
		sim.SimVarPlaneAltitude(),
		sim.SimVarPlaneLatitude(sim.UnitDegrees), // You can force the units
		sim.SimVarPlaneLongitude(),
		sim.SimVarIndicatedAltitude(),
		sim.SimVarGeneralEngRpm(1),
		sim.SimVarAutopilotMaster(),
	)
	if err != nil {
		panic(err)
	}
	cSimStatus := sc.ConnectSysEventSim()
	//wait sim start
	for {
		if <-cSimStatus {
			break
		}
	}
	crashed := sc.ConnectSysEventCrashed()
	for {
		select {
		case result := <-cSimVar:
			for _, simVar := range result {
				var f float64
				var err error
				if strings.Contains(string(simVar.Unit), "String") {
					log.Printf("%s : %#v\n", simVar.Name, simVar.GetString())
				} else if simVar.Unit == "SIMCONNECT_DATA_LATLONALT" {
					data, _ := simVar.GetDataLatLonAlt()
					log.Printf("%s : %#v\n", simVar.Name, data)
				} else if simVar.Unit == "SIMCONNECT_DATA_XYZ" {
					data, _ := simVar.GetDataXYZ()
					log.Printf("%s : %#v\n", simVar.Name, data)
				} else if simVar.Unit == "SIMCONNECT_DATA_WAYPOINT" {
					data, _ := simVar.GetDataWaypoint()
					log.Printf("%s : %#v\n", simVar.Name, data)
				} else {
					f, err = simVar.GetFloat64()
					log.Println(simVar.Name, fmt.Sprintf("%f", f))
				}
				if err != nil {
					log.Println("return error :", err)
				}
			}

		case <-crashed:
			log.Println("Your are crashed !!")
			<-sc.Close() // Wait close confirmation
			return       // This example close after crash in the sim
		}

	}

}

```
