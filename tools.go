package simconnect

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/sirupsen/logrus"
)

func convStrToGoString(buf []byte) string {
	var index int
	var value byte
	for index, value = range buf {
		if value == 0x00 {
			break
		}
	}
	return string(buf[:index])

}

func convGoStringtoBytes(str string) []byte {
	str = str + "\x00"
	return []byte(str)
}

func convCBytesToGoBytes(ptr unsafe.Pointer, size int) ([]byte, error) {
	if size > 1<<30 {
		return nil, errors.New("Dispatch return to big size array data")
	}
	buf := make([]byte, size)
	copy(buf, (*[1 << 30]byte)(ptr)[:size:size])
	return buf, nil
}

func getByVarName(name string, listSimVar []SimVar) *SimVar {
	if strings.Contains(name, ":") {
		name = strings.Split(name, ":")[0]
	}
	for _, simVar := range listSimVar {
		if name == simVar.Name {
			return &simVar
		}
	}
	return nil
}

func getUnitForType(t string) SimVarUnit {
	switch t {
	case "string":
		return UnitString
	case "*SIMCONNECT_DATA_XYZ":
		return UnitSimconnectDataXyz
	case "*SIMCONNECT_DATA_LATLONALT":
		return UnitSimconnectDataLatlonalt
	case "*SIMCONNECT_DATA_WAYPOINT":
		return UnitSimconnectDataWaypoint
	default:
		return ""
	}
}

func SimVarGenerator(iFace interface{}) ([]SimVar, error) {
	rt := reflect.TypeOf(iFace)
	if rt.Kind() != reflect.Struct {
		return nil, errors.New("Interface error : " + rt.Name())
	}
	simVars := []SimVar{}
	var err error
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tag := f.Tag.Get("sim")
		if tag == "" {
			continue
		}
		index := 0
		if strings.Contains(tag, ":") {
			v := strings.Split(tag, ":")
			tag = v[0]
			index, err = strconv.Atoi(v[1])
			if err != nil {
				return nil, err
			}
		}
		unit := SimVarUnit(f.Tag.Get("simUnit"))
		if unit == "" {
			unit = getUnitForType(f.Type.Name())
		}
		simVar := SimVar{
			Name:  tag,
			Unit:  unit,
			Index: index,
		}
		simVars = append(simVars, simVar)
	}
	return simVars, nil
}

func InterfaceAssignSimVar(listSimVar []SimVar, iFace interface{}) {
	rt := reflect.ValueOf(iFace)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		switch field.Type().Name() {
		case "float64":
			listSimVar[i].SetFloat64(field.Float())
		}
	}
}
func SimVarAssignInterface(iFace interface{}, listSimVar []SimVar) interface{} {
	rt := reflect.TypeOf(iFace)
	if rt.Kind() != reflect.Struct {
		logrus.Warn("Interface error in SimVarAssignInterface:", rt.Name())
		return nil
	}
	reflectElem := reflect.New(rt).Elem()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tag := f.Tag.Get("sim")
		if tag == "" {
			continue
		}
		logWarm := func(i int64, err error) {
			logrus.Warn("#"+strconv.FormatInt(i, 10), "ignored field in AssignInterface", f.Name, " tag:", tag, "error:", err)
		}
		simVar := getByVarName(tag, listSimVar)
		if simVar == nil {
			logWarm(8, errors.New("SimVar not found"))
			continue
		}
		reflectValue := reflectElem.FieldByName(f.Name)

		if reflectValue.Kind() == 0 {
			logWarm(1, nil)
			continue
		}
		typeValue := reflectValue.Type().String()
		switch typeValue {
		case "string":
			reflectValue.SetString(simVar.GetString())
		case "float64":
			f, err := simVar.GetFloat64()
			if err != nil {
				logWarm(2, err)
				continue
			}
			reflectValue.SetFloat(f)
		case "bool":
			b, err := simVar.GetBool()
			if err != nil {
				logWarm(3, err)
				continue
			}
			reflectValue.SetBool(b)
		case "int":
			i, err := simVar.GetInt()
			if err != nil {
				logWarm(4, err)
				continue
			}
			reflectValue.SetInt(int64(i))
		case "*SIMCONNECT_DATA_XYZ":
			data, err := simVar.GetDataXYZ()
			if err != nil {
				logWarm(5, err)
				continue
			}
			reflectValue.Set(reflect.ValueOf(data))
		case "*SIMCONNECT_DATA_LATLONALT":
			data, err := simVar.GetDataLatLonAlt()
			if err != nil {
				logWarm(6, err)
				continue
			}
			reflectValue.Set(reflect.ValueOf(data))
		case "*SIMCONNECT_DATA_WAYPOINT":
			data, err := simVar.GetDataWaypoint()
			if err != nil {
				logWarm(7, err)
				continue
			}
			reflectValue.Set(reflect.ValueOf(data))
		default:
			logrus.Infoln("Type :", reflectValue.Type(), "?")
		}
	}
	return reflectElem.Interface()
}
