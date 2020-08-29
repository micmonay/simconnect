package simconnect_test

import (
	"log"
	"simconnect"
)

// ExampleGetSimVar this example show how to get SimVar with EasySimConnect
func ExampleGetSimVar() {
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
	newalt := simconnect.SimVarPlaneAltitude()
	newalt.SetFloat64(10000.0)
	sc.SetSimObject(newalt)
	for {
		result := <-cSimVar
		for _, simVar := range result {
			f, err := simVar.GetFloat64()
			if err != nil {
				panic(err)
			}
			log.Printf("%#v", f)
		}

	}

}
