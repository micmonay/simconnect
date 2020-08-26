package simconnect

import (
	"errors"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
)

// EasySimConnect use for best easy use SimConnect in golang
type EasySimConnect struct {
	sc         *SimConnect
	listSimVar [][]SimVar
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
		make([][]SimVar, 0),
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
	go esc.runDispatch()
	return nil
}

func convertToGoBytes(ptr unsafe.Pointer, size int) ([]byte, error) {
	if size > 1<<30 {
		return nil, errors.New("Dispatch return to big size array data")
	}
	buf := make([]byte, size)
	copy(buf, (*[1 << 30]byte)(ptr)[:size:size])
	return buf, nil
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
		buf, err := convertToGoBytes(ppdata, int(pcbData))
		if err != nil {
			logrus.Errorln(err)
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
			logrus.Infoln("SimConnect Exception : ", getTextException(recv.dwException), recv.dwSendID)
		case SIMCONNECT_RECV_ID_SIMOBJECT_DATA, SIMCONNECT_RECV_ID_SIMOBJECT_DATA_BYTYPE:
			recv := *(*SIMCONNECT_RECV_SIMOBJECT_DATA)(ppdata)
			if len(esc.listSimVar) < int(recv.dwDefineID) {
				logrus.Warnf("ListSimVar not found: %#v\n %#v\n %d>=%d", recv, esc.listSimVar, len(esc.listSimVar), int(recv.dwDefineID))
				continue
			}
			listSimVar := esc.listSimVar[recv.dwDefineID]
			if len(listSimVar) != int(recv.dwDefineCount) {
				logrus.Warnf("ListSimVar size not equal %#v ?= %#v\n", recv, listSimVar)
				continue
			}
			position := int(unsafe.Offsetof(recv.dwData))
			returnSimVar := make([]SimVar, len(listSimVar))
			for i, simVar := range listSimVar {
				size := simVar.GetSize()
				if position > int(pcbData) {
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

			esc.sc.RequestDataOnSimObjectType(uint32(0), recv.dwDefineID, uint32(0), uint32(0))

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
		err := esc.sc.AddToDataDefinition(defineID, simVar.Name, simVar.Units, simVar.GetDatumType(), 0, uint32(i))
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
