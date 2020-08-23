package simconnect

import (
	"bytes"
	"encoding/binary"
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

func (esc *EasySimConnect) RunDispatch() {
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
			if len(esc.listSimVar) > int(recv.dwDefineID) {
				logrus.Warnf("ListSimVar not found: %#v\n", recv)
				continue
			}
			listSimVar := esc.listSimVar[recv.dwDefineID]
			if len(listSimVar) > int(recv.dwDefineCount) {
				logrus.Warnf("ListSimVar size not equal %#v ?= %#v\n", recv, listSimVar)
				continue
			}
			for i, simVar := range listSimVar {
				var buf = make([]byte, 8*i)
				buf = *(*[]byte)(unsafe.Pointer(&recv.dwData))
				var f float64
				binary.Read(bytes.NewReader(buf[:]), binary.LittleEndian, &f)
				println(f)
				bytes := buf[i*8 : i+1*8]
				simVar.data = &bytes
				// TODO: créer les autres type de variable
			}

		default:
			logrus.Infof("%#v\n", recvInfo)
		}
	}
}
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
