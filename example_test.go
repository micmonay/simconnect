package simconnect_test

import (
	"log"
	"time"

	"github.com/micmonay/simconnect"
)

// ExampleGetSimVar this example show how to get SimVar with EasySimConnect
func Example_getSimVar() {
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
	for i := 0; i < 1; i++ {
		result := <-cSimVar
		for _, simVar := range result {
			f, err := simVar.GetFloat64()
			if err != nil {
				panic(err)
			}
			log.Printf("%#v\n", f)
		}

	}
	// Output:

}

func Example_getSimVarWithIndex() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	cSimVar := sc.ConnectStructToSimObject(
		simconnect.SimVarGeneralEngRpm(1),
		simconnect.SimVarTransponderCode(1),
	)
	for i := 0; i < 1; i++ {
		result := <-cSimVar
		for _, simVar := range result {

			if simVar.Name == simconnect.SimVarTransponderCode().Name {
				i, err := simVar.GetInt()
				if err != nil {
					panic(err)
				}
				log.Printf("%s : %x\n", simVar.Name, i)
			} else {
				f, err := simVar.GetFloat64()
				if err != nil {
					panic(err)
				}
				log.Printf("%s : %f\n", simVar.Name, f)
			}
		}

	}
	// Output:
}

//
func Example_setSimVar() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	newalt := simconnect.SimVarPlaneAltitude()
	newalt.SetFloat64(6000.0)
	sc.SetSimObject(newalt)
	time.Sleep(1000 * time.Millisecond)
	// NOEXEC Output:
}

func Example_getLatLonAlt() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	cSimVar := sc.ConnectStructToSimObject(
		simconnect.SimVarStructLatlonalt(),
	)
	for i := 0; i < 1; i++ {
		result := <-cSimVar
		for _, simVar := range result {
			latlonalt, err := simVar.GetDataLatLonAlt()
			if err != nil {
				panic(err)
			}
			log.Printf("%s : %#v\n", simVar.Name, latlonalt)
		}

	}
	// Output:
}

func Example_getXYZ() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	cSimVar := sc.ConnectStructToSimObject(
		simconnect.SimVarEyepointPosition(),
	)
	for i := 0; i < 1; i++ {
		result := <-cSimVar
		for _, simVar := range result {
			xyz, err := simVar.GetDataXYZ()
			if err != nil {
				panic(err)
			}
			log.Printf("%s : %#v\n", simVar.Name, xyz)
		}

	}
	// Output:
}

func Example_getString() {
	sc, err := simconnect.NewEasySimConnect()
	if err != nil {
		panic(err)
	}
	err = sc.Connect("MyApp")
	if err != nil {
		panic(err)
	}
	cSimVar := sc.ConnectStructToSimObject(
		simconnect.SimVarAtcAirline(),
		simconnect.SimVarCategory(),
	)
	for i := 0; i < 1; i++ {
		result := <-cSimVar
		for _, simVar := range result {
			str := simVar.GetString()
			log.Printf("%s : %#v\n", simVar.Name, str)
		}

	}
	// Output:
}
