package simconnect

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
)

// EasySimConnect use for best easy use SimConnect in golang
type EasySimConnect struct {
	sc         *SimConnect
	listSimVar [][]*SimVar
	listChan   []chan []SimVar
	delay      time.Duration
}

// NewEasySimConnect create instance of EasySimConnect
func NewEasySimConnect() (*EasySimConnect, error) {
	sc, err := NewSimConnect()
	if err != nil {
		return nil, err
	}
	return &EasySimConnect{
		sc,
		make([][]*SimVar, 0),
		make([]chan []SimVar, 0),
		100 * time.Millisecond,
	}, nil
}

// Connect to sim and run dispatch or return error
func (esc *EasySimConnect) Connect(appName string) error {
	err := esc.sc.Open(appName)
	if err != nil {
		return err
	}
	go esc.RunDispatch()
	return nil
}

func getMemoryByte(startPos uintptr, size uint) []byte {
	buf := make([]byte, size)
	for i := uint(0); i < size; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(startPos + uintptr(i)))
	}
	return buf
}

func getSize(t uint32) int {
	switch t {
	case SIMCONNECT_DATATYPE_FLOAT64, SIMCONNECT_DATATYPE_INT64, SIMCONNECT_DATATYPE_STRING8:
		return 8
	case SIMCONNECT_DATATYPE_FLOAT32, SIMCONNECT_DATATYPE_INT32:
		return 4
	case SIMCONNECT_DATATYPE_STRING32:
		return 32
	case SIMCONNECT_DATATYPE_STRING64:
		return 64
	case SIMCONNECT_DATATYPE_STRING128:
		return 128
	case SIMCONNECT_DATATYPE_STRING256:
		return 256
	case SIMCONNECT_DATATYPE_STRING260:
		return 260
	}
	logrus.Warnln("Not found size for the type : ", t)
	return 0
}

func (esc *EasySimConnect) RunDispatch() {
	defer esc.sc.Close()
	for {
		time.Sleep(esc.delay)
		var ppdata unsafe.Pointer
		var pcbData uint32
		err := esc.sc.GetNextDispatch(&ppdata, &pcbData)
		if err != nil {
			continue
		}
		recvInfo := *(*SIMCONNECT_RECV)(ppdata)
		switch recvInfo.dwID {
		case SIMCONNECT_RECV_ID_EVENT:
			//err = sc.RequestDataOnSimObjectType(0, 0, 0, 0)
			/*if err != nil {
				println(err)
			}*/
		case SIMCONNECT_RECV_ID_EXCEPTION:
			recv := *(*SIMCONNECT_RECV_EXCEPTION)(ppdata)
			fmt.Printf("%#v\n", recv)
		case SIMCONNECT_RECV_ID_SIMOBJECT_DATA, SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
			recv := *(*SIMCONNECT_RECV_SIMOBJECT_DATA)(ppdata)
			if len(esc.listSimVar) < int(recv.dwDefineID) {
				logrus.Warnf("ListSimVar not found: %#v\n %#v\n %d>=%d", recv, esc.listSimVar, len(esc.listSimVar), int(recv.dwDefineID))
				continue
			}
			listSimVar := esc.listSimVar[recv.dwDefineID]
			if len(listSimVar) < int(recv.dwDefineCount) {
				logrus.Warnf("ListSimVar size not equal %#v ?= %#v\n", recv, listSimVar)
				continue
			}
			buf := (*[1 << 30]byte)(ppdata)[:recv.dwSize:recv.dwSize]
			position := int(unsafe.Offsetof(recv.dwData))
			returnSimVar := make([]SimVar, len(listSimVar))
			for i, simVar := range listSimVar {
				size := getSize(simVar.GetDatumType())
				buf := buf[position : position+size]
				position = position + size
				simVar.data = &buf
				returnSimVar[i] = *simVar
			}
			esc.listChan[recv.dwDefineID] <- returnSimVar
			esc.sc.RequestDataOnSimObjectType(uint32(0), recv.dwDefineID, uint32(0), uint32(0))

		default:
			logrus.Infof("%#v\n", recvInfo)
		}
	}
}

// ConnConnectStructToSimObject
func (esc *EasySimConnect) ConnectStructToSimObject(listSimVar ...SimVar) chan []SimVar {
	defineID := uint32(len(esc.listSimVar))
	addedSimVar := make([]*SimVar, 0)
	for i, simVar := range listSimVar {
		err := esc.sc.AddToDataDefinition(defineID, simVar.Name, simVar.Units, simVar.GetDatumType(), 0, uint32(i))
		if err != nil {
			logrus.Infoln("Error add SimVar (", simVar.Name, ") error :", err)
			continue
		}
		addedSimVar = append(addedSimVar, &simVar)
	}
	esc.listSimVar = append(esc.listSimVar, addedSimVar)
	chanSimVar := make(chan []SimVar)
	esc.listChan = append(esc.listChan, chanSimVar)
	esc.sc.RequestDataOnSimObjectType(uint32(0), defineID, uint32(0), uint32(0))
	return chanSimVar
}
