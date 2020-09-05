package simconnect

import (
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
)

// EasySimConnect use for best easy use SimConnect in golang
type EasySimConnect struct {
	sc         *SimConnect
	delay      time.Duration
	listSimVar [][]SimVar
	listChan   []chan []SimVar
	indexEvent int
	listEvent  map[int]func(interface{})
}

// NewEasySimConnect create instance of EasySimConnect
func NewEasySimConnect() (*EasySimConnect, error) {
	sc, err := NewSimConnect()
	if err != nil {
		return nil, err
	}
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	return &EasySimConnect{
		sc,
		100 * time.Millisecond,
		make([][]SimVar, 0),
		make([]chan []SimVar, 0),
		0,
		make(map[int]func(interface{})),
	}, nil
}

//SetDelay select delay update
func (esc *EasySimConnect) SetDelay(t time.Duration) {
	esc.delay = t
}

// Connect to sim and run dispatch or return error
func (esc *EasySimConnect) Connect(appName string) error {
	err := esc.sc.Open(appName)
	if err != nil {
		return err
	}
	go esc.runDispatch()
	return nil
}

func (esc *EasySimConnect) runDispatch() {
	defer esc.sc.Close()
	for {
		time.Sleep(esc.delay)
		var ppdata unsafe.Pointer
		var pcbData uint32
		err := esc.sc.GetNextDispatch(&ppdata, &pcbData)
		//créer un buffer en copy les data ppdata avec longueur pcbdata et utiliser le buffer pour la suite
		if err != nil {
			continue
		}
		buf, err := convCBytesToGoBytes(ppdata, int(pcbData))
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		recvInfo := *(*SIMCONNECT_RECV)(ppdata)
		switch recvInfo.dwID {
		case SIMCONNECT_RECV_ID_OPEN:
			recv := *(*SIMCONNECT_RECV_OPEN)(ppdata)
			logrus.Infoln("Connected to", convStrToGoString(recv.szApplicationName[:]))
		case SIMCONNECT_RECV_ID_EVENT:
			recv := *(*SIMCONNECT_RECV_EVENT)(ppdata)
			cb, found := esc.listEvent[int(recv.uEventID)]
			if !found {
				logrus.Infof("Ignored event : %#v\n", recv)
				continue
			}
			cb(recv)
		case SIMCONNECT_RECV_ID_EVENT_FILENAME:
			recv := *(*SIMCONNECT_RECV_EVENT_FILENAME)(ppdata)
			esc.listEvent[int(recv.uEventID)](recv)
		case SIMCONNECT_RECV_ID_EXCEPTION:
			recv := *(*SIMCONNECT_RECV_EXCEPTION)(ppdata)
			logrus.Infoln("SimConnect Exception : ", getTextException(recv.dwException), recv.dwSendID)
		case SIMCONNECT_RECV_ID_SIMOBJECT_DATA, SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
			recv := *(*SIMCONNECT_RECV_SIMOBJECT_DATA)(ppdata)
			if len(esc.listSimVar) < int(recv.dwDefineID) {
				logrus.Warnf("ListSimVar not found: %#v\n %#v\n %d>=%d", recv, esc.listSimVar, len(esc.listSimVar), int(recv.dwDefineID))
				continue
			}
			listSimVar := esc.listSimVar[recv.dwDefineID]
			if len(listSimVar) != int(recv.dwDefineCount) {
				logrus.Warnf("ListSimVar size not equal %#v ?= %#v\n", int(recv.dwDefineCount), len(listSimVar))
				continue
			}
			position := int(unsafe.Offsetof(recv.dwData))
			returnSimVar := make([]SimVar, len(listSimVar))
			for i, simVar := range listSimVar {
				size := simVar.GetSize()
				if position+size > int(pcbData) {
					logrus.Errorln("slice bounds out of range")
					break
				}
				simVar.data = buf[position : position+size]
				returnSimVar[i] = simVar
				position = position + size
			}
			select {
			case esc.listChan[recv.dwDefineID] <- returnSimVar:
			case <-time.After(esc.delay):
			}
			go func() {
				time.Sleep(esc.delay)
				esc.sc.RequestDataOnSimObjectType(uint32(0), recv.dwDefineID, uint32(0), uint32(0))
			}()

		default:
			logrus.Infof("%#v\n", recvInfo)
		}
	}
}

// ConnectStructToSimObject this function return a chan. This chan return update with SimVar in order of argument
func (esc *EasySimConnect) ConnectStructToSimObject(listSimVar ...SimVar) chan []SimVar {
	defineID := uint32(len(esc.listSimVar))
	addedSimVar := make([]SimVar, 0)
	for i, simVar := range listSimVar {
		err := esc.sc.AddToDataDefinition(defineID, simVar.getNameForDataDefinition(), simVar.getUnitsForDataDefinition(), simVar.GetDatumType(), 0, uint32(i))
		if err != nil {
			logrus.Infoln("Error add SimVar (", simVar.Name, ") error :", err)
			continue
		}
		addedSimVar = append(addedSimVar, simVar)
	}
	esc.listSimVar = append(esc.listSimVar, addedSimVar)
	chanSimVar := make(chan []SimVar)
	esc.listChan = append(esc.listChan, chanSimVar)
	esc.sc.RequestDataOnSimObjectType(uint32(0), defineID, uint32(0), uint32(0))
	return chanSimVar
}

//SetSimObject use for set SimVar in FS
func (esc *EasySimConnect) SetSimObject(simVar SimVar) {
	defineID := uint32(1 << 30)
	err := esc.sc.AddToDataDefinition(defineID, simVar.Name, simVar.Units, simVar.GetDatumType(), 0, 0)
	if err != nil {
		logrus.Infoln("Error set SimVar (", simVar.Name, ") in AddToDataDefinition error :", err)
		return
	}
	//esc.listSimVar = append(esc.listSimVar, []*SimVar{&simVar})
	err = esc.sc.SetDataOnSimObject(defineID, SIMCONNECT_OBJECT_ID_USER, 0, 0, 8, simVar.data)
	if err != nil {
		logrus.Infoln("Error set SimVar (", simVar.Name, ") in SetDataOnSimObject error :", err)
		return
	}
	err = esc.sc.ClearDataDefinition(uint32(defineID))
	if err != nil {
		logrus.Infoln("Error set SimVar (", simVar.Name, ") in ClearDataDefinition error :", err)
		return
	}
}
func (esc *EasySimConnect) connectSysEvent(name string, cb func(interface{})) {
	esc.listEvent[esc.indexEvent] = cb
	err := esc.sc.SubscribeToSystemEvent(uint32(esc.indexEvent), name)
	esc.indexEvent++
	if err != nil {
		logrus.Infoln("Error connect to Event ", name, " in ConnectSysEventCrashed error :", err)
	}
}

//ConnectSysEventCrashed Request a notification if the user aircraft crashes.
func (esc *EasySimConnect) ConnectSysEventCrashed() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SimEventCrashed, func(data interface{}) {
		c <- true
	})
	return c
}

//ConnectSysEventCrashReset Request a notification when the crash cut-scene has completed.
func (esc *EasySimConnect) ConnectSysEventCrashReset() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SimEventCrashReset, func(data interface{}) {
		c <- true
	})
	return c
}

//ConnectSysEventPause Request notifications when the flight is paused or unpaused, and also immediately returns the current pause state (1 = paused or 0 = unpaused). The state is returned in the dwData parameter.
func (esc *EasySimConnect) ConnectSysEventPause() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SimEventPause, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT)
		c <- event.dwData > 0
	})
	return c
}

//ConnectSysEventPaused Request a notification when the flight is paused.
func (esc *EasySimConnect) ConnectSysEventPaused() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SimEventPaused, func(data interface{}) {
		c <- true
	})
	return c
}

//ConnectSysEventFlightPlanDeactivated Request a notification when the active flight plan is de-activated.
func (esc *EasySimConnect) ConnectSysEventFlightPlanDeactivated() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SimEventFlightPlanDeactivated, func(data interface{}) {
		c <- true
	})
	return c
}

//ConnectSysEventAircraftLoaded Request a notification when the aircraft flight dynamics file is changed. These files have a .AIR extension. The filename is returned in a string.
func (esc *EasySimConnect) ConnectSysEventAircraftLoaded() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SimEventAircraftLoaded, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

//ConnectSysEventFlightLoaded 	Request a notification when a flight is loaded. Note that when a flight is ended, a default flight is typically loaded, so these events will occur when flights and missions are started and finished. The filename of the flight loaded is returned in a string
func (esc *EasySimConnect) ConnectSysEventFlightLoaded() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SimEventFlightLoaded, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

//ConnectSysEventFlightSaved 	Request a notification when a flight is saved correctly. The filename of the flight saved is returned in a string
func (esc *EasySimConnect) ConnectSysEventFlightSaved() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SimEventFlightSaved, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

//ConnectSysEventFlightPlanActivated Request a notification when a new flight plan is activated. The filename of the activated flight plan is returned in a string.
func (esc *EasySimConnect) ConnectSysEventFlightPlanActivated() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SimEventFlightPlanActivated, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

//ShowText is used to display a text menu, or scrolling or static text, on the screen.
//
//Time is in second and return chan with event
func (esc *EasySimConnect) ShowText(str string, time float32, color PrintColor) (<-chan int, error) {
	buf := convGoStringtoBytes(str)
	cReturn := make(chan int)
	esc.listEvent[esc.indexEvent] = func(data interface{}) {
		cReturn <- int(data.(SIMCONNECT_RECV_EVENT).dwData)
	}
	esc.indexEvent++
	return cReturn, esc.sc.Text(uint32(color), 60, 0, buf)
}
