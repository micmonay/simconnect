package simconnect

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
)

// EasySimConnectLogLevel is a type of Log level
type EasySimConnectLogLevel int

// Log Level
const (
	LogNo EasySimConnectLogLevel = iota
	LogError
	LogWarn
	LogInfo
)

// EasySimConnect for easy use of SimConnect in golang
// Please show example_test.go for use case
type EasySimConnect struct {
	sc           *SimConnect
	delay        time.Duration
	listSimVar   [][]SimVar
	listChan     []chan []SimVar
	indexEvent   uint32
	listEvent    map[uint32]func(interface{})
	listSimEvent map[KeySimEvent]SimEvent
	logLevel     EasySimConnectLogLevel
	cOpen        chan bool
	alive        bool
	cException   chan *SIMCONNECT_RECV_EXCEPTION
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
		make(map[uint32]func(interface{})),
		make(map[KeySimEvent]SimEvent),
		LogNo,
		make(chan bool, 1),
		true,
		make(chan *SIMCONNECT_RECV_EXCEPTION),
	}, nil
}

// SetLoggerLevel you can set log level in EasySimConnect
func (esc *EasySimConnect) SetLoggerLevel(level EasySimConnectLogLevel) {
	esc.logLevel = level
}

// Close Finishing EasySimConnect, All object created with this EasySimConnect's instance is perished after call this function
func (esc *EasySimConnect) Close() <-chan bool {
	esc.alive = false
	return esc.cOpen
}

// IsAlive return true if connected
func (esc *EasySimConnect) IsAlive() bool {
	return esc.alive
}

// SetDelay Select delay update SimVar and
func (esc *EasySimConnect) SetDelay(t time.Duration) {
	esc.delay = t
}

// Connect to sim and run dispatch or return error
func (esc *EasySimConnect) Connect(appName string) (<-chan bool, error) {
	err, _ := esc.sc.Open(appName)
	if err != nil {
		return nil, err
	}
	go esc.runDispatch()
	return esc.cOpen, nil
}

func (esc *EasySimConnect) logf(level EasySimConnectLogLevel, format string, args ...interface{}) {
	if level > esc.logLevel {
		return
	}
	if level == LogInfo {
		logrus.Infof(format, args...)
	}
	if level == LogWarn {
		logrus.Warnf(format, args...)
	}
	if level == LogError {
		logrus.Errorf(format, args...)
	}
}

func (esc *EasySimConnect) runDispatch() {
	for esc.alive {
		var ppdata unsafe.Pointer
		var pcbData uint32
		err, _ := esc.sc.GetNextDispatch(&ppdata, &pcbData)
		//créer un buffer en copy les data ppdata avec longueur pcbdata et utiliser le buffer pour la suite
		if err != nil {
			time.Sleep(esc.delay / 2)
			continue
		}
		buf, err := convCBytesToGoBytes(ppdata, int(pcbData))
		if err != nil {
			esc.logf(LogError, "%v#", err)
			continue
		}
		recvInfo := *(*SIMCONNECT_RECV)(ppdata)
		switch recvInfo.dwID {
		case SIMCONNECT_RECV_ID_OPEN:
			recv := *(*SIMCONNECT_RECV_OPEN)(ppdata)
			esc.logf(LogInfo, "Connected to %s", convStrToGoString(recv.szApplicationName[:]))
			esc.cOpen <- true
		case SIMCONNECT_RECV_ID_EVENT:
			recv := *(*SIMCONNECT_RECV_EVENT)(ppdata)
			cb, found := esc.listEvent[recv.uEventID]
			if !found {
				esc.logf(LogInfo, "Ignored event : %#v\n", recv)
				continue
			}
			cb(recv)
		case SIMCONNECT_RECV_ID_QUIT:
			esc.sc.Close()
			esc.cOpen <- false
			return
		case SIMCONNECT_RECV_ID_EVENT_FILENAME:
			recv := *(*SIMCONNECT_RECV_EVENT_FILENAME)(ppdata)
			esc.listEvent[recv.uEventID](recv)
		case SIMCONNECT_RECV_ID_EXCEPTION:
			recv := (*SIMCONNECT_RECV_EXCEPTION)(ppdata)
			select {
			case esc.cException <- recv:
			case <-time.After(100 * time.Millisecond):
			}
			esc.logf(LogInfo, "SimConnect Exception : %s %#v\n", getTextException(recv.dwException), *recv)
		case SIMCONNECT_RECV_ID_SIMOBJECT_DATA, SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
			recv := *(*SIMCONNECT_RECV_SIMOBJECT_DATA)(ppdata)
			if len(esc.listSimVar) < int(recv.dwDefineID) {
				esc.logf(LogWarn, "ListSimVar not found: %#v\n %#v\n %d>=%d", recv, esc.listSimVar, len(esc.listSimVar), int(recv.dwDefineID))
				continue
			}
			listSimVar := esc.listSimVar[recv.dwDefineID]
			if len(listSimVar) != int(recv.dwDefineCount) {
				esc.logf(LogWarn, "ListSimVar size not equal %#v ?= %#v\n", int(recv.dwDefineCount), len(listSimVar))
				continue
			}
			position := int(unsafe.Offsetof(recv.dwData))
			returnSimVar := make([]SimVar, len(listSimVar))
			for i, simVar := range listSimVar {
				size := simVar.GetSize()
				if position+size > int(pcbData) {
					esc.logf(LogError, "slice bounds out of range")
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
			esc.logf(LogInfo, "%#v\n", recvInfo)
		}
	}
	esc.sc.Close()
	esc.cOpen <- false
}

// ConnectToSimVar return a chan. This chan return an array when updating they SimVars in order of argument of this function
func (esc *EasySimConnect) ConnectToSimVar(listSimVar ...SimVar) (<-chan []SimVar, error) {
	defineID := uint32(len(esc.listSimVar))
	addedSimVar := make([]SimVar, 0)
	for i, simVar := range listSimVar {
		err, id := esc.sc.AddToDataDefinition(defineID, simVar.getNameForDataDefinition(), simVar.getUnitForDataDefinition(), simVar.GetDatumType(), 0, uint32(i))
		if err != nil {
			esc.logf(LogInfo, "Error add SimVar ( %s ) in AddToDataDefinition error : %#v", simVar.Name, err)
			return nil, fmt.Errorf(
				"Error add SimVar ( %s ) in AddToDataDefinition error : %#v",
				simVar.Name,
				err,
			)
		}
		var exception *SIMCONNECT_RECV_EXCEPTION
		select {
		case exception = <-esc.cException:
		case <-time.After(100 * time.Millisecond):
		}
		if exception != nil && exception.dwSendID == id {
			return nil, fmt.Errorf(
				"Error add SimVar ( %s ) in AddToDataDefinition : %s. Please control name ( %s ) and unit ( %s )",
				simVar.Name,
				getTextException(exception.dwException),
				simVar.Name,
				simVar.Unit,
			)
		}
		addedSimVar = append(addedSimVar, simVar)
	}
	esc.listSimVar = append(esc.listSimVar, addedSimVar)
	chanSimVar := make(chan []SimVar)
	esc.listChan = append(esc.listChan, chanSimVar)
	esc.sc.RequestDataOnSimObjectType(uint32(0), defineID, uint32(0), uint32(0))
	return chanSimVar, nil
}

// ConnectToSimVarObject return a chan. This chan return an array when updating they SimVars in order of argument of this function
//
// Deprecated: Use ConnectToSimVar instead.
func (esc *EasySimConnect) ConnectToSimVarObject(listSimVar ...SimVar) <-chan []SimVar {
	c, err := esc.ConnectToSimVar(listSimVar...)
	if err != nil {
		esc.logf(LogError, err.Error())
		return make(<-chan []SimVar)
	}
	return c
}

// ConnectInterfaceToSimVar return a chan. This chan return interface when updating
func (esc *EasySimConnect) ConnectInterfaceToSimVar(iFace interface{}) (<-chan interface{}, error) {
	simVars, err := SimVarGenerator(iFace)
	if err != nil {
		return nil, err
	}
	csimVars, err := esc.ConnectToSimVar(simVars...)
	if err != nil {
		return nil, err
	}
	cInterface := make(chan interface{})
	go func() {
		for {
			cInterface <- SimVarAssignInterface(iFace, <-csimVars)
		}
	}()
	return cInterface, nil
}

// SetSimObject edit the SimVar in the simulator
func (esc *EasySimConnect) SetSimObject(simVar SimVar) {
	defineID := uint32(1 << 30)
	err, _ := esc.sc.AddToDataDefinition(defineID, simVar.Name, simVar.getUnitForDataDefinition(), simVar.GetDatumType(), 0, 0)
	if err != nil {
		esc.logf(LogInfo, "Error add SimVar ( %s ) in AddToDataDefinition error : %#v", simVar.Name, err)
		return
	}
	//esc.listSimVar = append(esc.listSimVar, []*SimVar{&simVar})
	err, _ = esc.sc.SetDataOnSimObject(defineID, SIMCONNECT_OBJECT_ID_USER, 0, 0, uint32(len(simVar.data)), simVar.data)
	if err != nil {
		esc.logf(LogInfo, "Error add SimVar ( %s ) in SetDataOnSimObject error : %#v", simVar.Name, err)
		return
	}
	err, _ = esc.sc.ClearDataDefinition(uint32(defineID))
	if err != nil {
		esc.logf(LogInfo, "Error add SimVar ( %s ) in ClearDataDefinition error : %#v", simVar.Name, err)
		return
	}
}
func (esc *EasySimConnect) connectSysEvent(name SystemEvent, cb func(interface{})) {
	esc.indexEvent++
	esc.listEvent[esc.indexEvent] = cb
	err, _ := esc.sc.SubscribeToSystemEvent(uint32(esc.indexEvent), name)
	if err != nil {
		esc.logf(LogInfo, "Error connect to Event %s in ConnectSysEventCrashed error : %#v", name, err)
	}
}

// ConnectSysEventCrashed Request a notification if the user aircraft crashes.
func (esc *EasySimConnect) ConnectSysEventCrashed() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventCrashed, func(data interface{}) {
		c <- true
	})
	return c
}

// ConnectSysEventCrashReset Request a notification when the crash cut-scene has completed.
func (esc *EasySimConnect) ConnectSysEventCrashReset() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventCrashReset, func(data interface{}) {
		c <- true
	})
	return c
}

// ConnectSysEventPause Request notifications when the flight is paused or unpaused, and also immediately returns the current pause state (1 = paused or 0 = unpaused). The state is returned in the dwData parameter.
func (esc *EasySimConnect) ConnectSysEventPause() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventPause, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT)
		c <- event.dwData > 0
	})
	return c
}

// ConnectSysEventPaused Request a notification when the flight is paused.
func (esc *EasySimConnect) ConnectSysEventPaused() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventPaused, func(data interface{}) {
		c <- true
	})
	return c
}

// ConnectSysEventSim Request a notification when Sim start and stop.
func (esc *EasySimConnect) ConnectSysEventSim() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventSim, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT)
		c <- event.dwData > 0
	})
	return c
}

// ConnectSysEventFlightPlanDeactivated Request a notification when the active flight plan is de-activated.
func (esc *EasySimConnect) ConnectSysEventFlightPlanDeactivated() <-chan bool {
	c := make(chan bool)
	esc.connectSysEvent(SystemEventFlightPlanDeactivated, func(data interface{}) {
		c <- true
	})
	return c
}

// ConnectSysEventAircraftLoaded Request a notification when the aircraft flight dynamics file is changed. These files have a .AIR extension. The filename is returned in a string.
func (esc *EasySimConnect) ConnectSysEventAircraftLoaded() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SystemEventAircraftLoaded, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

// ConnectSysEventFlightLoaded 	Request a notification when a flight is loaded. Note that when a flight is ended, a default flight is typically loaded, so these events will occur when flights and missions are started and finished. The filename of the flight loaded is returned in a string
func (esc *EasySimConnect) ConnectSysEventFlightLoaded() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SystemEventFlightLoaded, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

// ConnectSysEventFlightSaved 	Request a notification when a flight is saved correctly. The filename of the flight saved is returned in a string
func (esc *EasySimConnect) ConnectSysEventFlightSaved() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SystemEventFlightSaved, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

// ConnectSysEventFlightPlanActivated Request a notification when a new flight plan is activated. The filename of the activated flight plan is returned in a string.
func (esc *EasySimConnect) ConnectSysEventFlightPlanActivated() <-chan string {
	c := make(chan string)
	esc.connectSysEvent(SystemEventFlightPlanActivated, func(data interface{}) {
		event := data.(SIMCONNECT_RECV_EVENT_FILENAME)
		c <- convStrToGoString(event.szFileName[:])
	})
	return c
}

// ShowText display a text on the screen in the simulator.
//
// ime is in second and return chan a confirmation for the simulator
func (esc *EasySimConnect) ShowText(str string, time float32, color PrintColor) (<-chan int, error) {
	cReturn := make(chan int)
	esc.indexEvent++
	esc.listEvent[esc.indexEvent] = func(data interface{}) {
		cReturn <- int(data.(SIMCONNECT_RECV_EVENT).dwData)
	}
	err, _ := esc.sc.Text(uint32(color), time, esc.indexEvent, str)
	return cReturn, err
}
func (esc *EasySimConnect) runSimEvent(simEvent SimEvent) {
	esc.sc.TransmitClientEvent(SIMCONNECT_OBJECT_ID_USER, simEvent.eventID, simEvent.Value, SIMCONNECT_GROUP_PRIORITY_HIGHEST, SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY)
}

// NewSimEvent return new instance of SimEvent and you can run SimEvent.Run()
func (esc *EasySimConnect) NewSimEvent(simEventStr KeySimEvent) SimEvent {
	instance, found := esc.listSimEvent[simEventStr]
	if found {
		return instance
	}

	esc.indexEvent++
	c := make(chan int32)
	simEvent := SimEvent{
		simEventStr,
		0,
		esc.runSimEvent,
		c,
		esc.indexEvent,
	}
	esc.listEvent[esc.indexEvent] = func(data interface{}) {
		recv := data.(SIMCONNECT_RECV_EVENT)
		c <- int32(recv.dwData)
	}
	esc.sc.MapClientEventToSimEvent(esc.indexEvent, string(simEventStr))
	esc.sc.AddClientEventToNotificationGroup(0, esc.indexEvent, false)
	esc.sc.SetNotificationGroupPriority(0, SIMCONNECT_GROUP_PRIORITY_HIGHEST)
	esc.listSimEvent[simEventStr] = simEvent
	return simEvent
}

// SimEvent Use for generate action in the simulator
type SimEvent struct {
	Mapping KeySimEvent
	Value   int
	run     func(simEvent SimEvent)
	cb      <-chan int32
	eventID uint32
}

// Run return chan bool when receive the event is finish
func (s SimEvent) Run() <-chan int32 {
	s.run(s)
	return s.cb
}

// RunWithValue return chan bool when receive the event is finish
func (s SimEvent) RunWithValue(value int) <-chan int32 {
	s.Value = value
	return s.Run()
}
