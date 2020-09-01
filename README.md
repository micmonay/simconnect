# Golang SimConnect binding

_This library is in developement and I have write little documentation. Please considerate this library as a testing version._

## Example how to use my library
```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/micmonay/simconnect"
)

func main() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	cSimVar := sc.ConnectStructToSimObject(
		simconnect.SimVarPlaneAltitude(),
		simconnect.SimVarPlaneLatitude(),
		simconnect.SimVarPlaneLongitude(),
		simconnect.SimVarIndicatedAltitude(),
		simconnect.SimVarAutopilotAltitudeLockVar(),
	)
	sc.SetDelay(1 * time.Second)
	crashed := sc.ConnectSysEventCrashed()
	for {
		select {
		case result := <-cSimVar:
			for _, simVar := range result {
				var f float64
				var err error
				if simVar.Units == "Radians" {
					f, err = simVar.GetDegrees()
				} else {
					f, err = simVar.GetFloat64()
				}
				if err != nil {
					log.Println(err)
				}
				log.Println(simVar.Name, fmt.Sprintf("%f", f))
			}

		case <-crashed:
            log.Println("Your are crashed !!")
            return
		}

	}

}

```