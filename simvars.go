package simconnect

import (
	"bytes"
	"encoding/binary"
	"math"
	"strconv"
	"strings"
	"unsafe"

	"github.com/sirupsen/logrus"
)

// SimVar is usued for all SimVar describtion
type SimVar struct {
	Name     string
	Units    string
	Settable bool
	Index    int
	data     []byte
}

func (s *SimVar) getUnitsForDataDefinition() string {
	if strings.Contains(s.Units, "String") ||
		strings.Contains(s.Units, "string") ||
		s.Units == "SIMCONNECT_DATA_LATLONALT" ||
		s.Units == "SIMCONNECT_DATA_XYZ" ||
		s.Units == "SIMCONNECT_DATA_WAYPOINT" {
		return ""
	}
	return s.Units
}
func (s *SimVar) getNameForDataDefinition() string {
	if strings.Contains(s.Name, ":index") {
		return strings.Replace(s.Name, ":index", ":"+strconv.Itoa(s.Index), 1)
	}
	return s.Name
}
func (s *SimVar) GetData() []byte {
	return s.data
}

func (s *SimVar) GetDatumType() uint32 {
	switch s.Units {
	case "Bool":
		return SIMCONNECT_DATATYPE_INT32
	case "String8":
		return SIMCONNECT_DATATYPE_STRING8
	case "String64":
		return SIMCONNECT_DATATYPE_STRING64
	case "String":
		return SIMCONNECT_DATATYPE_STRING256
	case "SIMCONNECT_DATA_LATLONALT":
		return SIMCONNECT_DATATYPE_LATLONALT
	case "SIMCONNECT_DATA_XYZ":
		return SIMCONNECT_DATATYPE_XYZ
	case "SIMCONNECT_DATA_WAYPOINT":
		return SIMCONNECT_DATATYPE_WAYPOINT
	default:
		return SIMCONNECT_DATATYPE_FLOAT64
	}
}

func (s *SimVar) GetSize() int {
	switch s.GetDatumType() {
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
	case SIMCONNECT_DATATYPE_LATLONALT:
		return int(unsafe.Sizeof(SIMCONNECT_DATA_LATLONALT{}))
	case SIMCONNECT_DATATYPE_XYZ:
		return int(unsafe.Sizeof(SIMCONNECT_DATA_XYZ{}))
	case SIMCONNECT_DATATYPE_WAYPOINT:
		return int(unsafe.Sizeof(SIMCONNECT_DATA_WAYPOINT{}))
	}
	logrus.Warnln("Not found size for the type : ", s.GetDatumType())
	return 0
}

func (s *SimVar) GetFloat64() (float64, error) {
	var f float64
	err := binary.Read(bytes.NewReader(s.data), binary.LittleEndian, &f)
	if err != nil {
		return 0, err
	}
	return f, nil
}

//GetInt lost precision
func (s *SimVar) GetInt() (int, error) {
	f, err := s.GetFloat64()
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func (s *SimVar) GetDegrees() (float64, error) {
	f, err := s.GetFloat64()
	if err != nil {
		return 0, err
	}
	return f * 180 / math.Pi, nil
}

func (s *SimVar) GetDataXYZ() (*SIMCONNECT_DATA_XYZ, error) {
	var data SIMCONNECT_DATA_XYZ
	err := binary.Read(bytes.NewReader(s.data), binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *SimVar) GetDataLatLonAlt() (*SIMCONNECT_DATA_LATLONALT, error) {
	var data SIMCONNECT_DATA_LATLONALT
	err := binary.Read(bytes.NewReader(s.data), binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *SimVar) GetDataWaypoint() (*SIMCONNECT_DATA_WAYPOINT, error) {
	var data SIMCONNECT_DATA_WAYPOINT
	err := binary.Read(bytes.NewReader(s.data), binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *SimVar) SetFloat64(f float64) {
	s.data = make([]byte, 8)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &f)
	copy(s.data, buf.Bytes())
}

func (s *SimVar) GetString() string {
	return convStrToGoString(s.data)
}

// SimVarAutopilotPitchHold Simvar
// args contain optional index
func SimVarAutopilotPitchHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT PITCH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarStructAmbientWind Simvar
// args contain optional index
func SimVarStructAmbientWind(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT AMBIENT WIND",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarLaunchbarPosition Simvar
// args contain optional index
func SimVarLaunchbarPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LAUNCHBAR POSITION",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarNumberOfCatapults Simvar
// args contain optional index
func SimVarNumberOfCatapults(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NUMBER OF CATAPULTS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHoldbackBarInstalled Simvar
// args contain optional index
func SimVarHoldbackBarInstalled(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HOLDBACK BAR INSTALLED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarBlastShieldPosition Simvar
// args contain optional index
func SimVarBlastShieldPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BLAST SHIELD POSITION:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRecipEngDetonating Simvar
// args contain optional index
func SimVarRecipEngDetonating(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG DETONATING:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRecipEngCylinderHealth Simvar
// args contain optional index
func SimVarRecipEngCylinderHealth(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG CYLINDER HEALTH:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRecipEngNumCylinders Simvar
// args contain optional index
func SimVarRecipEngNumCylinders(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NUM CYLINDERS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipEngNumCylindersFailed Simvar
// args contain optional index
func SimVarRecipEngNumCylindersFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NUM CYLINDERS FAILED",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipEngAntidetonationTankValve Simvar
// args contain optional index
func SimVarRecipEngAntidetonationTankValve(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK VALVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankQuantity Simvar
// args contain optional index
func SimVarRecipEngAntidetonationTankQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK QUANTITY:index",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankMaxQuantity Simvar
// args contain optional index
func SimVarRecipEngAntidetonationTankMaxQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK MAX QUANTITY:index",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarRecipEngNitrousTankValve Simvar
// args contain optional index
func SimVarRecipEngNitrousTankValve(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK VALVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankQuantity Simvar
// args contain optional index
func SimVarRecipEngNitrousTankQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK QUANTITY:index",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankMaxQuantity Simvar
// args contain optional index
func SimVarRecipEngNitrousTankMaxQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK MAX QUANTITY:index",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarPayloadStationObject Simvar
// args contain optional index
func SimVarPayloadStationObject(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION OBJECT:index",
		Units:    "String",
		Settable: true,
	}
}

// SimVarPayloadStationNumSimobjects Simvar
// args contain optional index
func SimVarPayloadStationNumSimobjects(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION NUM SIMOBJECTS:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarSlingObjectAttached Simvar
// args contain optional index
func SimVarSlingObjectAttached(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING OBJECT ATTACHED:index",
		Units:    "Bool/String",
		Settable: false,
	}
}

// SimVarSlingCableBroken Simvar
// args contain optional index
func SimVarSlingCableBroken(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING CABLE BROKEN:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSlingCableExtendedLength Simvar
// args contain optional index
func SimVarSlingCableExtendedLength(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING CABLE EXTENDED LENGTH:index",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarSlingActivePayloadStation Simvar
// args contain optional index
func SimVarSlingActivePayloadStation(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING ACTIVE PAYLOAD STATION:index",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarSlingHoistPercentDeployed Simvar
// args contain optional index
func SimVarSlingHoistPercentDeployed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING HOIST PERCENT DEPLOYED:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarSlingHookInPickupMode Simvar
// args contain optional index
func SimVarSlingHookInPickupMode(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SLING HOOK IN PICKUP MODE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAttachedToSling Simvar
// args contain optional index
func SimVarIsAttachedToSling(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS ATTACHED TO SLING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAlternateStaticSourceOpen Simvar
// args contain optional index
func SimVarAlternateStaticSourceOpen(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ALTERNATE STATIC SOURCE OPEN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAileronTrimPct Simvar
// args contain optional index
func SimVarAileronTrimPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON TRIM PCT",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: true,
	}
}

// SimVarRudderTrimPct Simvar
// args contain optional index
func SimVarRudderTrimPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER TRIM PCT",
		Units:    "Percent over 100",
		Settable: true,
	}
}

// SimVarLightOnStates Simvar
// args contain optional index
func SimVarLightOnStates(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT ON STATES",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarLightStates Simvar
// args contain optional index
func SimVarLightStates(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT STATES",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarLandingLightPbh Simvar
// args contain optional index
func SimVarLandingLightPbh(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LANDING LIGHT PBH",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarLightTaxiOn Simvar
// args contain optional index
func SimVarLightTaxiOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT TAXI ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightStrobeOn Simvar
// args contain optional index
func SimVarLightStrobeOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT STROBE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightPanelOn Simvar
// args contain optional index
func SimVarLightPanelOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT PANEL ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightRecognitionOn Simvar
// args contain optional index
func SimVarLightRecognitionOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT RECOGNITION ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightWingOn Simvar
// args contain optional index
func SimVarLightWingOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT WING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLogoOn Simvar
// args contain optional index
func SimVarLightLogoOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT LOGO ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightCabinOn Simvar
// args contain optional index
func SimVarLightCabinOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT CABIN ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightHeadOn Simvar
// args contain optional index
func SimVarLightHeadOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT HEAD ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBrakeOn Simvar
// args contain optional index
func SimVarLightBrakeOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT BRAKE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightNavOn Simvar
// args contain optional index
func SimVarLightNavOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT NAV ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBeaconOn Simvar
// args contain optional index
func SimVarLightBeaconOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT BEACON ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLandingOn Simvar
// args contain optional index
func SimVarLightLandingOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT LANDING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAiDesiredSpeed Simvar
// args contain optional index
func SimVarAiDesiredSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI DESIRED SPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiWaypointList Actually not supported
// args contain optional index
func SimVarAiWaypointList(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI WAYPOINT LIST",
		Units:    "SIMCONNECT_DATA_WAYPOINT",
		Settable: true,
	}
}

// SimVarAiCurrentWaypoint Simvar
// args contain optional index
func SimVarAiCurrentWaypoint(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI CURRENT WAYPOINT",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarAiDesiredHeading Simvar
// args contain optional index
func SimVarAiDesiredHeading(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI DESIRED HEADING",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarAiGroundturntime Simvar
// args contain optional index
func SimVarAiGroundturntime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDTURNTIME",
		Units:    "Seconds",
		Settable: true,
	}
}

// SimVarAiGroundcruisespeed Simvar
// args contain optional index
func SimVarAiGroundcruisespeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDCRUISESPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiGroundturnspeed Simvar
// args contain optional index
func SimVarAiGroundturnspeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDTURNSPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiTrafficIsifr Simvar
// args contain optional index
func SimVarAiTrafficIsifr(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ISIFR",
		Units:    "Boolean",
		Settable: false,
	}
}

// SimVarAiTrafficState Simvar
// args contain optional index
func SimVarAiTrafficState(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC STATE",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficCurrentAirport Simvar
// args contain optional index
func SimVarAiTrafficCurrentAirport(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC CURRENT AIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficAssignedRunway Simvar
// args contain optional index
func SimVarAiTrafficAssignedRunway(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ASSIGNED RUNWAY",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficAssignedParking Simvar
// args contain optional index
func SimVarAiTrafficAssignedParking(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ASSIGNED PARKING",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficFromairport Simvar
// args contain optional index
func SimVarAiTrafficFromairport(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC FROMAIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficToairport Simvar
// args contain optional index
func SimVarAiTrafficToairport(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC TOAIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficEtd Simvar
// args contain optional index
func SimVarAiTrafficEtd(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ETD",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarAiTrafficEta Simvar
// args contain optional index
func SimVarAiTrafficEta(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarDroppableObjectsType Simvar
// args contain optional index
func SimVarDroppableObjectsType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS TYPE:index",
		Units:    "String",
		Settable: true,
	}
}

// SimVarDroppableObjectsCount Simvar
// args contain optional index
func SimVarDroppableObjectsCount(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS COUNT:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarWingFlexPct Simvar
// args contain optional index
func SimVarWingFlexPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WING FLEX PCT:index",
		Units:    "Percent over 100",
		Settable: true,
	}
}

// SimVarApplyHeatToSystems Simvar
// args contain optional index
func SimVarApplyHeatToSystems(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APPLY HEAT TO SYSTEMS",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAdfLatlonalt Simvar
// args contain optional index
func SimVarAdfLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavVorLatlonalt Simvar
// args contain optional index
func SimVarNavVorLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV VOR LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavGsLatlonalt Simvar
// args contain optional index
func SimVarNavGsLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GS LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavDmeLatlonalt Simvar
// args contain optional index
func SimVarNavDmeLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV DME LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarInnerMarkerLatlonalt Simvar
// args contain optional index
func SimVarInnerMarkerLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INNER MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarMiddleMarkerLatlonalt Simvar
// args contain optional index
func SimVarMiddleMarkerLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MIDDLE MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarOuterMarkerLatlonalt Simvar
// args contain optional index
func SimVarOuterMarkerLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "OUTER MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructLatlonalt Simvar
// args contain optional index
func SimVarStructLatlonalt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructLatlonaltpbh Simvar
// args contain optional index
func SimVarStructLatlonaltpbh(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT LATLONALTPBH",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructSurfaceRelativeVelocity Simvar
// args contain optional index
func SimVarStructSurfaceRelativeVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT SURFACE RELATIVE VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldvelocity Simvar
// args contain optional index
func SimVarStructWorldvelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLDVELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldRotationVelocity Simvar
// args contain optional index
func SimVarStructWorldRotationVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLD ROTATION VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructBodyVelocity Simvar
// args contain optional index
func SimVarStructBodyVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT BODY VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructBodyRotationVelocity Simvar
// args contain optional index
func SimVarStructBodyRotationVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT BODY ROTATION VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldAcceleration Simvar
// args contain optional index
func SimVarStructWorldAcceleration(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLD ACCELERATION",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEnginePosition Simvar
// args contain optional index
func SimVarStructEnginePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT ENGINE POSITION:index",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEyepointDynamicAngle Simvar
// args contain optional index
func SimVarStructEyepointDynamicAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT EYEPOINT DYNAMIC ANGLE",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEyepointDynamicOffset Simvar
// args contain optional index
func SimVarStructEyepointDynamicOffset(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCT EYEPOINT DYNAMIC OFFSET",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarEyepointPosition Simvar
// args contain optional index
func SimVarEyepointPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EYEPOINT POSITION",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarFlyByWireElacSwitch Simvar
// args contain optional index
func SimVarFlyByWireElacSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE ELAC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireFacSwitch Simvar
// args contain optional index
func SimVarFlyByWireFacSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE FAC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireSecSwitch Simvar
// args contain optional index
func SimVarFlyByWireSecSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE SEC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireElacFailed Simvar
// args contain optional index
func SimVarFlyByWireElacFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE ELAC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireFacFailed Simvar
// args contain optional index
func SimVarFlyByWireFacFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE FAC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireSecFailed Simvar
// args contain optional index
func SimVarFlyByWireSecFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE SEC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNumberOfEngines Simvar
// args contain optional index
func SimVarNumberOfEngines(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NUMBER OF ENGINES",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngineControlSelect Simvar
// args contain optional index
func SimVarEngineControlSelect(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENGINE CONTROL SELECT",
		Units:    "Mask",
		Settable: true,
	}
}

// SimVarThrottleLowerLimit Simvar
// args contain optional index
func SimVarThrottleLowerLimit(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "THROTTLE LOWER LIMIT",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngineType Simvar
// args contain optional index
func SimVarEngineType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENGINE TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarMasterIgnitionSwitch Simvar
// args contain optional index
func SimVarMasterIgnitionSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MASTER IGNITION SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngCombustion Simvar
// args contain optional index
func SimVarGeneralEngCombustion(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG COMBUSTION:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGeneralEngMasterAlternator Simvar
// args contain optional index
func SimVarGeneralEngMasterAlternator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MASTER ALTERNATOR:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpSwitch Simvar
// args contain optional index
func SimVarGeneralEngFuelPumpSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PUMP SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpOn Simvar
// args contain optional index
func SimVarGeneralEngFuelPumpOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PUMP ON:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngRpm Simvar
// args contain optional index
func SimVarGeneralEngRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngPctMaxRpm Simvar
// args contain optional index
func SimVarGeneralEngPctMaxRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG PCT MAX RPM:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngMaxReachedRpm Simvar
// args contain optional index
func SimVarGeneralEngMaxReachedRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MAX REACHED RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngThrottleLeverPosition Simvar
// args contain optional index
func SimVarGeneralEngThrottleLeverPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG THROTTLE LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngMixtureLeverPosition Simvar
// args contain optional index
func SimVarGeneralEngMixtureLeverPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MIXTURE LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngPropellerLeverPosition Simvar
// args contain optional index
func SimVarGeneralEngPropellerLeverPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG PROPELLER LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngStarter Simvar
// args contain optional index
func SimVarGeneralEngStarter(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG STARTER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngExhaustGasTemperature Simvar
// args contain optional index
func SimVarGeneralEngExhaustGasTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG EXHAUST GAS TEMPERATURE:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarGeneralEngOilPressure Simvar
// args contain optional index
func SimVarGeneralEngOilPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarGeneralEngOilLeakedPercent Simvar
// args contain optional index
func SimVarGeneralEngOilLeakedPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL LEAKED PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngCombustionSoundPercent Simvar
// args contain optional index
func SimVarGeneralEngCombustionSoundPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG COMBUSTION SOUND PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngDamagePercent Simvar
// args contain optional index
func SimVarGeneralEngDamagePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG DAMAGE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngOilTemperature Simvar
// args contain optional index
func SimVarGeneralEngOilTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL TEMPERATURE:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarGeneralEngFailed Simvar
// args contain optional index
func SimVarGeneralEngFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FAILED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorSwitch Simvar
// args contain optional index
func SimVarGeneralEngGeneratorSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG GENERATOR SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorActive Simvar
// args contain optional index
func SimVarGeneralEngGeneratorActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG GENERATOR ACTIVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGeneralEngAntiIcePosition Simvar
// args contain optional index
func SimVarGeneralEngAntiIcePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG ANTI ICE POSITION:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelValve Simvar
// args contain optional index
func SimVarGeneralEngFuelValve(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL VALVE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPressure Simvar
// args contain optional index
func SimVarGeneralEngFuelPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarGeneralEngElapsedTime Simvar
// args contain optional index
func SimVarGeneralEngElapsedTime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG ELAPSED TIME:index",
		Units:    "Hours",
		Settable: false,
	}
}

// SimVarRecipEngCowlFlapPosition Simvar
// args contain optional index
func SimVarRecipEngCowlFlapPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG COWL FLAP POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngPrimer Simvar
// args contain optional index
func SimVarRecipEngPrimer(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG PRIMER:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngManifoldPressure Simvar
// args contain optional index
func SimVarRecipEngManifoldPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG MANIFOLD PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarRecipEngAlternateAirPosition Simvar
// args contain optional index
func SimVarRecipEngAlternateAirPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ALTERNATE AIR POSITION:index",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRecipEngCoolantReservoirPercent Simvar
// args contain optional index
func SimVarRecipEngCoolantReservoirPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG COOLANT RESERVOIR PERCENT:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngLeftMagneto Simvar
// args contain optional index
func SimVarRecipEngLeftMagneto(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG LEFT MAGNETO:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngRightMagneto Simvar
// args contain optional index
func SimVarRecipEngRightMagneto(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG RIGHT MAGNETO:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngBrakePower Simvar
// args contain optional index
func SimVarRecipEngBrakePower(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG BRAKE POWER:index",
		Units:    "ft lb per second",
		Settable: true,
	}
}

// SimVarRecipEngStarterTorque Simvar
// args contain optional index
func SimVarRecipEngStarterTorque(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG STARTER TORQUE:index",
		Units:    "Foot pound",
		Settable: true,
	}
}

// SimVarRecipEngTurbochargerFailed Simvar
// args contain optional index
func SimVarRecipEngTurbochargerFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG TURBOCHARGER FAILED:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostActive Simvar
// args contain optional index
func SimVarRecipEngEmergencyBoostActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG EMERGENCY BOOST ACTIVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostElapsedTime Simvar
// args contain optional index
func SimVarRecipEngEmergencyBoostElapsedTime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG EMERGENCY BOOST ELAPSED TIME:index",
		Units:    "Hours",
		Settable: true,
	}
}

// SimVarRecipEngWastegatePosition Simvar
// args contain optional index
func SimVarRecipEngWastegatePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG WASTEGATE POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngTurbineInletTemperature Simvar
// args contain optional index
func SimVarRecipEngTurbineInletTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG TURBINE INLET TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngCylinderHeadTemperature Simvar
// args contain optional index
func SimVarRecipEngCylinderHeadTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG CYLINDER HEAD TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngRadiatorTemperature Simvar
// args contain optional index
func SimVarRecipEngRadiatorTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG RADIATOR TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngFuelAvailable Simvar
// args contain optional index
func SimVarRecipEngFuelAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL AVAILABLE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngFuelFlow Simvar
// args contain optional index
func SimVarRecipEngFuelFlow(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL FLOW:index",
		Units:    "Pounds per hour",
		Settable: true,
	}
}

// SimVarRecipEngFuelTankSelector Simvar
// args contain optional index
func SimVarRecipEngFuelTankSelector(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarRecipEngFuelTanksUsed Simvar
// args contain optional index
func SimVarRecipEngFuelTanksUsed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL TANKS USED:index",
		Units:    "Mask",
		Settable: true,
	}
}

// SimVarRecipEngFuelNumberTanksUsed Simvar
// args contain optional index
func SimVarRecipEngFuelNumberTanksUsed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL NUMBER TANKS USED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipCarburetorTemperature Simvar
// args contain optional index
func SimVarRecipCarburetorTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP CARBURETOR TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipMixtureRatio Simvar
// args contain optional index
func SimVarRecipMixtureRatio(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RECIP MIXTURE RATIO:index",
		Units:    "Ratio",
		Settable: true,
	}
}

// SimVarTurbEngN1 Simvar
func SimVarTurbEngN1(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG N1:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngN2 Simvar
func SimVarTurbEngN2(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG N2:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN1 Simvar
func SimVarTurbEngCorrectedN1(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED N1:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN2 Simvar
func SimVarTurbEngCorrectedN2(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED N2:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedFf Simvar
// args contain optional index
func SimVarTurbEngCorrectedFf(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED FF:index",
		Units:    "Pounds per hour",
		Settable: true,
	}
}

// SimVarTurbEngMaxTorquePercent Simvar
// args contain optional index
func SimVarTurbEngMaxTorquePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG MAX TORQUE PERCENT:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngPressureRatio Simvar
// args contain optional index
func SimVarTurbEngPressureRatio(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG PRESSURE RATIO:index",
		Units:    "Ratio",
		Settable: true,
	}
}

// SimVarTurbEngItt Simvar
// args contain optional index
func SimVarTurbEngItt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG ITT:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarTurbEngAfterburner Simvar
// args contain optional index
func SimVarTurbEngAfterburner(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG AFTERBURNER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngJetThrust Simvar
// args contain optional index
func SimVarTurbEngJetThrust(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG JET THRUST:index",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarTurbEngBleedAir Simvar
// args contain optional index
func SimVarTurbEngBleedAir(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG BLEED AIR:index",
		Units:    "Psi",
		Settable: false,
	}
}

// SimVarTurbEngTankSelector Simvar
// args contain optional index
func SimVarTurbEngTankSelector(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarTurbEngTanksUsed Simvar
// args contain optional index
func SimVarTurbEngTanksUsed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG TANKS USED:index",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarTurbEngNumTanksUsed Simvar
// args contain optional index
func SimVarTurbEngNumTanksUsed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG NUM TANKS USED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTurbEngFuelFlowPph Simvar
// args contain optional index
func SimVarTurbEngFuelFlowPph(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG FUEL FLOW PPH:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarTurbEngFuelAvailable Simvar
// args contain optional index
func SimVarTurbEngFuelAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG FUEL AVAILABLE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngReverseNozzlePercent Simvar
// args contain optional index
func SimVarTurbEngReverseNozzlePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG REVERSE NOZZLE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarTurbEngVibration Simvar
// args contain optional index
func SimVarTurbEngVibration(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG VIBRATION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngFailed Simvar
// args contain optional index
func SimVarEngFailed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG FAILED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngRpmAnimationPercent Simvar
// args contain optional index
func SimVarEngRpmAnimationPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG RPM ANIMATION PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngOnFire Simvar
// args contain optional index
func SimVarEngOnFire(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG ON FIRE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarEngFuelFlowBugPosition Simvar
// args contain optional index
func SimVarEngFuelFlowBugPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL FLOW BUG POSITION:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarPropRpm Simvar
// args contain optional index
func SimVarPropRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP RPM:index",
		Units:    "Rpm",
		Settable: true,
	}
}

// SimVarPropMaxRpmPercent Simvar
// args contain optional index
func SimVarPropMaxRpmPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP MAX RPM PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarPropThrust Simvar
// args contain optional index
func SimVarPropThrust(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP THRUST:index",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarPropBeta Simvar
// args contain optional index
func SimVarPropBeta(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP BETA:index",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropFeatheringInhibit Simvar
// args contain optional index
func SimVarPropFeatheringInhibit(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHERING INHIBIT:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropFeathered Simvar
// args contain optional index
func SimVarPropFeathered(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHERED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropSyncDeltaLever Simvar
// args contain optional index
func SimVarPropSyncDeltaLever(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP SYNC DELTA LEVER:index",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarPropAutoFeatherArmed Simvar
// args contain optional index
func SimVarPropAutoFeatherArmed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP AUTO FEATHER ARMED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropFeatherSwitch Simvar
// args contain optional index
func SimVarPropFeatherSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHER SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPanelAutoFeatherSwitch Simvar
// args contain optional index
func SimVarPanelAutoFeatherSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PANEL AUTO FEATHER SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropSyncActive Simvar
// args contain optional index
func SimVarPropSyncActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP SYNC ACTIVE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropDeiceSwitch Simvar
// args contain optional index
func SimVarPropDeiceSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP DEICE SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngCombustion Simvar
// args contain optional index
func SimVarEngCombustion(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG COMBUSTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngN1Rpm Simvar
func SimVarEngN1Rpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG N1 RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarEngN2Rpm Simvar
func SimVarEngN2Rpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG N2 RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarEngFuelFlowPph Simvar
// args contain optional index
func SimVarEngFuelFlowPph(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL FLOW PPH:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarEngTorque Simvar
// args contain optional index
func SimVarEngTorque(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG TORQUE:index",
		Units:    "Foot pounds",
		Settable: false,
	}
}

// SimVarEngAntiIce Simvar
// args contain optional index
func SimVarEngAntiIce(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG ANTI ICE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngPressureRatio Simvar
// args contain optional index
func SimVarEngPressureRatio(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG PRESSURE RATIO:index",
		Units:    "Ratio (0-16384)",
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperature Simvar
// args contain optional index
func SimVarEngExhaustGasTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG EXHAUST GAS TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperatureGes Simvar
// args contain optional index
func SimVarEngExhaustGasTemperatureGes(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG EXHAUST GAS TEMPERATURE GES:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngCylinderHeadTemperature Simvar
// args contain optional index
func SimVarEngCylinderHeadTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG CYLINDER HEAD TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngOilTemperature Simvar
// args contain optional index
func SimVarEngOilTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG OIL TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngOilPressure Simvar
// args contain optional index
func SimVarEngOilPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG OIL PRESSURE:index",
		Units:    "pound-force per square inch",
		Settable: false,
	}
}

// SimVarEngOilQuantity Simvar
// args contain optional index
func SimVarEngOilQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG OIL QUANTITY:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngHydraulicPressure Simvar
// args contain optional index
func SimVarEngHydraulicPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG HYDRAULIC PRESSURE:index",
		Units:    "pound-force per square inch",
		Settable: false,
	}
}

// SimVarEngHydraulicQuantity Simvar
// args contain optional index
func SimVarEngHydraulicQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG HYDRAULIC QUANTITY:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngManifoldPressure Simvar
// args contain optional index
func SimVarEngManifoldPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG MANIFOLD PRESSURE:index",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarEngVibration Simvar
// args contain optional index
func SimVarEngVibration(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG VIBRATION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngRpmScaler Simvar
// args contain optional index
func SimVarEngRpmScaler(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG RPM SCALER:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngTurbineTemperature Simvar
// args contain optional index
func SimVarEngTurbineTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG TURBINE TEMPERATURE:index",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarEngTorquePercent Simvar
// args contain optional index
func SimVarEngTorquePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG TORQUE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngFuelPressure Simvar
// args contain optional index
func SimVarEngFuelPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL PRESSURE:index",
		Units:    "PSI",
		Settable: false,
	}
}

// SimVarEngElectricalLoad Simvar
// args contain optional index
func SimVarEngElectricalLoad(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG ELECTRICAL LOAD:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngTransmissionPressure Simvar
// args contain optional index
func SimVarEngTransmissionPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG TRANSMISSION PRESSURE:index",
		Units:    "PSI",
		Settable: false,
	}
}

// SimVarEngTransmissionTemperature Simvar
// args contain optional index
func SimVarEngTransmissionTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG TRANSMISSION TEMPERATURE:index",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarEngRotorRpm Simvar
// args contain optional index
func SimVarEngRotorRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG ROTOR RPM:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngMaxRpm Simvar
// args contain optional index
func SimVarEngMaxRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENG MAX RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngStarterActive Simvar
// args contain optional index
func SimVarGeneralEngStarterActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG STARTER ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelUsedSinceStart Simvar
// args contain optional index
func SimVarGeneralEngFuelUsedSinceStart(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL USED SINCE START",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarTurbEngPrimaryNozzlePercent Simvar
// args contain optional index
func SimVarTurbEngPrimaryNozzlePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG PRIMARY NOZZLE PERCENT:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarTurbEngIgnitionSwitch Simvar
// args contain optional index
func SimVarTurbEngIgnitionSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG IGNITION SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngMasterStarterSwitch Simvar
// args contain optional index
func SimVarTurbEngMasterStarterSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURB ENG MASTER STARTER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFuelTankCenterLevel Simvar
// args contain optional index
func SimVarFuelTankCenterLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenter2Level Simvar
func SimVarFuelTankCenter2Level(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenter3Level Simvar
func SimVarFuelTankCenter3Level(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftMainLevel Simvar
// args contain optional index
func SimVarFuelTankLeftMainLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxLevel Simvar
// args contain optional index
func SimVarFuelTankLeftAuxLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftTipLevel Simvar
// args contain optional index
func SimVarFuelTankLeftTipLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightMainLevel Simvar
// args contain optional index
func SimVarFuelTankRightMainLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightAuxLevel Simvar
// args contain optional index
func SimVarFuelTankRightAuxLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightTipLevel Simvar
// args contain optional index
func SimVarFuelTankRightTipLevel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankExternal1Level Simvar
func SimVarFuelTankExternal1Level(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankExternal2Level Simvar
func SimVarFuelTankExternal2Level(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenterCapacity Simvar
// args contain optional index
func SimVarFuelTankCenterCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenter2Capacity Simvar
func SimVarFuelTankCenter2Capacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenter3Capacity Simvar
func SimVarFuelTankCenter3Capacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftMainCapacity Simvar
// args contain optional index
func SimVarFuelTankLeftMainCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftAuxCapacity Simvar
// args contain optional index
func SimVarFuelTankLeftAuxCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftTipCapacity Simvar
// args contain optional index
func SimVarFuelTankLeftTipCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightMainCapacity Simvar
// args contain optional index
func SimVarFuelTankRightMainCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightAuxCapacity Simvar
// args contain optional index
func SimVarFuelTankRightAuxCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightTipCapacity Simvar
// args contain optional index
func SimVarFuelTankRightTipCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankExternal1Capacity Simvar
func SimVarFuelTankExternal1Capacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankExternal2Capacity Simvar
func SimVarFuelTankExternal2Capacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelLeftCapacity Simvar
// args contain optional index
func SimVarFuelLeftCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL LEFT CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelRightCapacity Simvar
// args contain optional index
func SimVarFuelRightCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL RIGHT CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenterQuantity Simvar
// args contain optional index
func SimVarFuelTankCenterQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankCenter2Quantity Simvar
func SimVarFuelTankCenter2Quantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankCenter3Quantity Simvar
func SimVarFuelTankCenter3Quantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftMainQuantity Simvar
// args contain optional index
func SimVarFuelTankLeftMainQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxQuantity Simvar
// args contain optional index
func SimVarFuelTankLeftAuxQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftTipQuantity Simvar
// args contain optional index
func SimVarFuelTankLeftTipQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightMainQuantity Simvar
// args contain optional index
func SimVarFuelTankRightMainQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightAuxQuantity Simvar
// args contain optional index
func SimVarFuelTankRightAuxQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightTipQuantity Simvar
// args contain optional index
func SimVarFuelTankRightTipQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankExternal1Quantity Simvar
func SimVarFuelTankExternal1Quantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankExternal2Quantity Simvar
func SimVarFuelTankExternal2Quantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelLeftQuantity Simvar
// args contain optional index
func SimVarFuelLeftQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL LEFT QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelRightQuantity Simvar
// args contain optional index
func SimVarFuelRightQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL RIGHT QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTotalQuantity Simvar
// args contain optional index
func SimVarFuelTotalQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelWeightPerGallon Simvar
// args contain optional index
func SimVarFuelWeightPerGallon(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL WEIGHT PER GALLON",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarFuelTankSelector Simvar
// args contain optional index
func SimVarFuelTankSelector(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarFuelCrossFeed Simvar
// args contain optional index
func SimVarFuelCrossFeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL CROSS FEED",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarFuelTotalCapacity Simvar
// args contain optional index
func SimVarFuelTotalCapacity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelSelectedQuantityPercent Simvar
// args contain optional index
func SimVarFuelSelectedQuantityPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED QUANTITY PERCENT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFuelSelectedQuantity Simvar
// args contain optional index
func SimVarFuelSelectedQuantity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTotalQuantityWeight Simvar
// args contain optional index
func SimVarFuelTotalQuantityWeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL QUANTITY WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarNumFuelSelectors Simvar
// args contain optional index
func SimVarNumFuelSelectors(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NUM FUEL SELECTORS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUnlimitedFuel Simvar
// args contain optional index
func SimVarUnlimitedFuel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "UNLIMITED FUEL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEstimatedFuelFlow Simvar
// args contain optional index
func SimVarEstimatedFuelFlow(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ESTIMATED FUEL FLOW",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarLightStrobe Simvar
// args contain optional index
func SimVarLightStrobe(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT STROBE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightPanel Simvar
// args contain optional index
func SimVarLightPanel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT PANEL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLanding Simvar
// args contain optional index
func SimVarLightLanding(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT LANDING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightTaxi Simvar
// args contain optional index
func SimVarLightTaxi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT TAXI",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBeacon Simvar
// args contain optional index
func SimVarLightBeacon(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT BEACON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightNav Simvar
// args contain optional index
func SimVarLightNav(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT NAV",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLogo Simvar
// args contain optional index
func SimVarLightLogo(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT LOGO",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightWing Simvar
// args contain optional index
func SimVarLightWing(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT WING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightRecognition Simvar
// args contain optional index
func SimVarLightRecognition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT RECOGNITION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightCabin Simvar
// args contain optional index
func SimVarLightCabin(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LIGHT CABIN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGroundVelocity Simvar
// args contain optional index
func SimVarGroundVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GROUND VELOCITY",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarTotalWorldVelocity Simvar
// args contain optional index
func SimVarTotalWorldVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WORLD VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarVelocityBodyZ Simvar
// args contain optional index
func SimVarVelocityBodyZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityBodyX Simvar
// args contain optional index
func SimVarVelocityBodyX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityBodyY Simvar
// args contain optional index
func SimVarVelocityBodyY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldZ Simvar
// args contain optional index
func SimVarVelocityWorldZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldX Simvar
// args contain optional index
func SimVarVelocityWorldX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldY Simvar
// args contain optional index
func SimVarVelocityWorldY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarAccelerationWorldX Simvar
// args contain optional index
func SimVarAccelerationWorldX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD X",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationWorldY Simvar
// args contain optional index
func SimVarAccelerationWorldY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD Y",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationWorldZ Simvar
// args contain optional index
func SimVarAccelerationWorldZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD Z",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyX Simvar
// args contain optional index
func SimVarAccelerationBodyX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY X",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyY Simvar
// args contain optional index
func SimVarAccelerationBodyY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY Y",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyZ Simvar
// args contain optional index
func SimVarAccelerationBodyZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY Z",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyX Simvar
// args contain optional index
func SimVarRotationVelocityBodyX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyY Simvar
// args contain optional index
func SimVarRotationVelocityBodyY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyZ Simvar
// args contain optional index
func SimVarRotationVelocityBodyZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRelativeWindVelocityBodyX Simvar
// args contain optional index
func SimVarRelativeWindVelocityBodyX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyY Simvar
// args contain optional index
func SimVarRelativeWindVelocityBodyY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyZ Simvar
// args contain optional index
func SimVarRelativeWindVelocityBodyZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPlaneAltAboveGround Simvar
// args contain optional index
func SimVarPlaneAltAboveGround(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE ALT ABOVE GROUND",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarPlaneLatitude Simvar
// args contain optional index
func SimVarPlaneLatitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE LATITUDE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneLongitude Simvar
// args contain optional index
func SimVarPlaneLongitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE LONGITUDE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneAltitude Simvar
// args contain optional index
func SimVarPlaneAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE ALTITUDE",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarPlanePitchDegrees Simvar
// args contain optional index
func SimVarPlanePitchDegrees(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE PITCH DEGREES",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneBankDegrees Simvar
// args contain optional index
func SimVarPlaneBankDegrees(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE BANK DEGREES",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesTrue Simvar
// args contain optional index
func SimVarPlaneHeadingDegreesTrue(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES TRUE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesMagnetic Simvar
// args contain optional index
func SimVarPlaneHeadingDegreesMagnetic(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES MAGNETIC",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarMagvar Simvar
// args contain optional index
func SimVarMagvar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MAGVAR",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGroundAltitude Simvar
// args contain optional index
func SimVarGroundAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GROUND ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarSurfaceType Simvar
// args contain optional index
func SimVarSurfaceType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SURFACE TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarSimOnGround Simvar
// args contain optional index
func SimVarSimOnGround(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SIM ON GROUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIncidenceAlpha Simvar
// args contain optional index
func SimVarIncidenceAlpha(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INCIDENCE ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarIncidenceBeta Simvar
// args contain optional index
func SimVarIncidenceBeta(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INCIDENCE BETA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAirspeedTrue Simvar
// args contain optional index
func SimVarAirspeedTrue(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED TRUE",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAirspeedIndicated Simvar
// args contain optional index
func SimVarAirspeedIndicated(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED INDICATED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAirspeedTrueCalibrate Simvar
// args contain optional index
func SimVarAirspeedTrueCalibrate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED TRUE CALIBRATE",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarAirspeedBarberPole Simvar
// args contain optional index
func SimVarAirspeedBarberPole(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED BARBER POLE",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAirspeedMach Simvar
// args contain optional index
func SimVarAirspeedMach(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED MACH",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarVerticalSpeed Simvar
// args contain optional index
func SimVarVerticalSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VERTICAL SPEED",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarMachMaxOperate Simvar
// args contain optional index
func SimVarMachMaxOperate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MACH MAX OPERATE",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarStallWarning Simvar
// args contain optional index
func SimVarStallWarning(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STALL WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarOverspeedWarning Simvar
// args contain optional index
func SimVarOverspeedWarning(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "OVERSPEED WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarBarberPoleMach Simvar
// args contain optional index
func SimVarBarberPoleMach(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BARBER POLE MACH",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarIndicatedAltitude Simvar
// args contain optional index
func SimVarIndicatedAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INDICATED ALTITUDE",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarKohlsmanSettingMb Simvar
// args contain optional index
func SimVarKohlsmanSettingMb(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "KOHLSMAN SETTING MB",
		Units:    "Millibars",
		Settable: true,
	}
}

// SimVarKohlsmanSettingHg Simvar
// args contain optional index
func SimVarKohlsmanSettingHg(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "KOHLSMAN SETTING HG",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarAttitudeIndicatorPitchDegrees Simvar
// args contain optional index
func SimVarAttitudeIndicatorPitchDegrees(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE INDICATOR PITCH DEGREES",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAttitudeIndicatorBankDegrees Simvar
// args contain optional index
func SimVarAttitudeIndicatorBankDegrees(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE INDICATOR BANK DEGREES",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAttitudeBarsPosition Simvar
// args contain optional index
func SimVarAttitudeBarsPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE BARS POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAttitudeCage Simvar
// args contain optional index
func SimVarAttitudeCage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE CAGE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarWiskeyCompassIndicationDegrees Simvar
// args contain optional index
func SimVarWiskeyCompassIndicationDegrees(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WISKEY COMPASS INDICATION DEGREES",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesGyro Simvar
// args contain optional index
func SimVarPlaneHeadingDegreesGyro(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES GYRO",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarHeadingIndicator Simvar
// args contain optional index
func SimVarHeadingIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HEADING INDICATOR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGyroDriftError Simvar
// args contain optional index
func SimVarGyroDriftError(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GYRO DRIFT ERROR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDeltaHeadingRate Simvar
// args contain optional index
func SimVarDeltaHeadingRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DELTA HEADING RATE",
		Units:    "Radians per second",
		Settable: true,
	}
}

// SimVarTurnCoordinatorBall Simvar
// args contain optional index
func SimVarTurnCoordinatorBall(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURN COORDINATOR BALL",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarAngleOfAttackIndicator Simvar
// args contain optional index
func SimVarAngleOfAttackIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ANGLE OF ATTACK INDICATOR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRadioHeight Simvar
// args contain optional index
func SimVarRadioHeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RADIO HEIGHT",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPartialPanelAdf Simvar
// args contain optional index
func SimVarPartialPanelAdf(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ADF",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAirspeed Simvar
// args contain optional index
func SimVarPartialPanelAirspeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL AIRSPEED",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAltimeter Simvar
// args contain optional index
func SimVarPartialPanelAltimeter(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ALTIMETER",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAttitude Simvar
// args contain optional index
func SimVarPartialPanelAttitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ATTITUDE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelComm Simvar
// args contain optional index
func SimVarPartialPanelComm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL COMM",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelCompass Simvar
// args contain optional index
func SimVarPartialPanelCompass(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL COMPASS",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelElectrical Simvar
// args contain optional index
func SimVarPartialPanelElectrical(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ELECTRICAL",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAvionics Simvar
// args contain optional index
func SimVarPartialPanelAvionics(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL AVIONICS",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelEngine Simvar
// args contain optional index
func SimVarPartialPanelEngine(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ENGINE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelFuelIndicator Simvar
// args contain optional index
func SimVarPartialPanelFuelIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL FUEL INDICATOR",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelHeading Simvar
// args contain optional index
func SimVarPartialPanelHeading(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL HEADING",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelVerticalVelocity Simvar
// args contain optional index
func SimVarPartialPanelVerticalVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL VERTICAL VELOCITY",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelTransponder Simvar
// args contain optional index
func SimVarPartialPanelTransponder(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL TRANSPONDER",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelNav Simvar
// args contain optional index
func SimVarPartialPanelNav(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL NAV",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelPitot Simvar
// args contain optional index
func SimVarPartialPanelPitot(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL PITOT",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelTurnCoordinator Simvar
// args contain optional index
func SimVarPartialPanelTurnCoordinator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL TURN COORDINATOR",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelVacuum Simvar
// args contain optional index
func SimVarPartialPanelVacuum(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL VACUUM",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarMaxGForce Simvar
// args contain optional index
func SimVarMaxGForce(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MAX G FORCE",
		Units:    "Gforce",
		Settable: false,
	}
}

// SimVarMinGForce Simvar
// args contain optional index
func SimVarMinGForce(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MIN G FORCE",
		Units:    "Gforce",
		Settable: false,
	}
}

// SimVarSuctionPressure Simvar
// args contain optional index
func SimVarSuctionPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SUCTION PRESSURE",
		Units:    "inHg",
		Settable: true,
	}
}

// SimVarAvionicsMasterSwitch Simvar
// args contain optional index
func SimVarAvionicsMasterSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AVIONICS MASTER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavSound Simvar
// args contain optional index
func SimVarNavSound(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV SOUND:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarDmeSound Simvar
// args contain optional index
func SimVarDmeSound(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DME SOUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfSound Simvar
// args contain optional index
func SimVarAdfSound(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF SOUND:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarMarkerSound Simvar
// args contain optional index
func SimVarMarkerSound(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MARKER SOUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComTransmit Simvar
// args contain optional index
func SimVarComTransmit(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM TRANSMIT:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComRecieveAll Simvar
// args contain optional index
func SimVarComRecieveAll(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM RECIEVE ALL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComActiveFrequency Simvar
// args contain optional index
func SimVarComActiveFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM ACTIVE FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarComStandbyFrequency Simvar
// args contain optional index
func SimVarComStandbyFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM STANDBY FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarComStatus Simvar
// args contain optional index
func SimVarComStatus(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM STATUS:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarNavAvailable Simvar
// args contain optional index
func SimVarNavAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV AVAILABLE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavActiveFrequency Simvar
// args contain optional index
func SimVarNavActiveFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV ACTIVE FREQUENCY:index",
		Units:    "MHz",
		Settable: false,
	}
}

// SimVarNavStandbyFrequency Simvar
// args contain optional index
func SimVarNavStandbyFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV STANDBY FREQUENCY:index",
		Units:    "MHz",
		Settable: false,
	}
}

// SimVarNavSignal Simvar
// args contain optional index
func SimVarNavSignal(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV SIGNAL:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavHasNav Simvar
// args contain optional index
func SimVarNavHasNav(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV HAS NAV:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasLocalizer Simvar
// args contain optional index
func SimVarNavHasLocalizer(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV HAS LOCALIZER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasDme Simvar
// args contain optional index
func SimVarNavHasDme(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV HAS DME:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasGlideSlope Simvar
// args contain optional index
func SimVarNavHasGlideSlope(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV HAS GLIDE SLOPE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavBackCourseFlags Simvar
// args contain optional index
func SimVarNavBackCourseFlags(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV BACK COURSE FLAGS:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavMagvar Simvar
// args contain optional index
func SimVarNavMagvar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV MAGVAR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavRadial Simvar
// args contain optional index
func SimVarNavRadial(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV RADIAL:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavRadialError Simvar
// args contain optional index
func SimVarNavRadialError(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV RADIAL ERROR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavLocalizer Simvar
// args contain optional index
func SimVarNavLocalizer(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV LOCALIZER:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavGlideSlopeError Simvar
// args contain optional index
func SimVarNavGlideSlopeError(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GLIDE SLOPE ERROR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavCdi Simvar
// args contain optional index
func SimVarNavCdi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV CDI:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavGsi Simvar
// args contain optional index
func SimVarNavGsi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GSI:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavTofrom Simvar
// args contain optional index
func SimVarNavTofrom(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV TOFROM:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarNavGsFlag Simvar
// args contain optional index
func SimVarNavGsFlag(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GS FLAG:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavObs Simvar
// args contain optional index
func SimVarNavObs(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV OBS:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavDme Simvar
// args contain optional index
func SimVarNavDme(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV DME:index",
		Units:    "Nautical miles",
		Settable: false,
	}
}

// SimVarNavDmespeed Simvar
// args contain optional index
func SimVarNavDmespeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV DMESPEED:index",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAdfActiveFrequency Simvar
// args contain optional index
func SimVarAdfActiveFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF ACTIVE FREQUENCY:index",
		Units:    "Frequency ADF BCD32",
		Settable: false,
	}
}

// SimVarAdfStandbyFrequency Simvar
// args contain optional index
func SimVarAdfStandbyFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF STANDBY FREQUENCY:index",
		Units:    "Hz",
		Settable: false,
	}
}

// SimVarAdfRadial Simvar
// args contain optional index
func SimVarAdfRadial(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF RADIAL:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAdfSignal Simvar
// args contain optional index
func SimVarAdfSignal(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF SIGNAL:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTransponderCode Simvar
// args contain optional index
func SimVarTransponderCode(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRANSPONDER CODE:index",
		Units:    "BCO16",
		Settable: false,
	}
}

// SimVarMarkerBeaconState Simvar
// args contain optional index
func SimVarMarkerBeaconState(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MARKER BEACON STATE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarInnerMarker Simvar
// args contain optional index
func SimVarInnerMarker(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INNER MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarMiddleMarker Simvar
// args contain optional index
func SimVarMiddleMarker(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MIDDLE MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarOuterMarker Simvar
// args contain optional index
func SimVarOuterMarker(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "OUTER MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarNavRawGlideSlope Simvar
// args contain optional index
func SimVarNavRawGlideSlope(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV RAW GLIDE SLOPE:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAdfCard Simvar
// args contain optional index
func SimVarAdfCard(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF CARD",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarHsiCdiNeedle Simvar
// args contain optional index
func SimVarHsiCdiNeedle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI CDI NEEDLE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHsiGsiNeedle Simvar
// args contain optional index
func SimVarHsiGsiNeedle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI GSI NEEDLE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHsiCdiNeedleValid Simvar
// args contain optional index
func SimVarHsiCdiNeedleValid(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI CDI NEEDLE VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiGsiNeedleValid Simvar
// args contain optional index
func SimVarHsiGsiNeedleValid(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI GSI NEEDLE VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiTfFlags Simvar
// args contain optional index
func SimVarHsiTfFlags(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI TF FLAGS",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarHsiBearingValid Simvar
// args contain optional index
func SimVarHsiBearingValid(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI BEARING VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiBearing Simvar
// args contain optional index
func SimVarHsiBearing(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI BEARING",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarHsiHasLocalizer Simvar
// args contain optional index
func SimVarHsiHasLocalizer(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI HAS LOCALIZER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiSpeed Simvar
// args contain optional index
func SimVarHsiSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI SPEED",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarHsiDistance Simvar
// args contain optional index
func SimVarHsiDistance(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI DISTANCE",
		Units:    "Nautical miles",
		Settable: false,
	}
}

// SimVarGpsPositionLat Simvar
// args contain optional index
func SimVarGpsPositionLat(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsPositionLon Simvar
// args contain optional index
func SimVarGpsPositionLon(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsPositionAlt Simvar
// args contain optional index
func SimVarGpsPositionAlt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsMagvar Simvar
// args contain optional index
func SimVarGpsMagvar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS MAGVAR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsIsActiveFlightPlan Simvar
// args contain optional index
func SimVarGpsIsActiveFlightPlan(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE FLIGHT PLAN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsActiveWayPoint Simvar
// args contain optional index
func SimVarGpsIsActiveWayPoint(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE WAY POINT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsArrived Simvar
// args contain optional index
func SimVarGpsIsArrived(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS ARRIVED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsDirecttoFlightplan Simvar
// args contain optional index
func SimVarGpsIsDirecttoFlightplan(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS DIRECTTO FLIGHTPLAN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsGroundSpeed Simvar
// args contain optional index
func SimVarGpsGroundSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND SPEED",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarGpsGroundTrueHeading Simvar
// args contain optional index
func SimVarGpsGroundTrueHeading(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND TRUE HEADING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsGroundMagneticTrack Simvar
// args contain optional index
func SimVarGpsGroundMagneticTrack(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND MAGNETIC TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsGroundTrueTrack Simvar
// args contain optional index
func SimVarGpsGroundTrueTrack(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND TRUE TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpDistance Simvar
// args contain optional index
func SimVarGpsWpDistance(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP DISTANCE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpBearing Simvar
// args contain optional index
func SimVarGpsWpBearing(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP BEARING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpTrueBearing Simvar
// args contain optional index
func SimVarGpsWpTrueBearing(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRUE BEARING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpCrossTrk Simvar
// args contain optional index
func SimVarGpsWpCrossTrk(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP CROSS TRK",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpDesiredTrack Simvar
// args contain optional index
func SimVarGpsWpDesiredTrack(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP DESIRED TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpTrueReqHdg Simvar
// args contain optional index
func SimVarGpsWpTrueReqHdg(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRUE REQ HDG",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpVerticalSpeed Simvar
// args contain optional index
func SimVarGpsWpVerticalSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP VERTICAL SPEED",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarGpsWpTrackAngleError Simvar
// args contain optional index
func SimVarGpsWpTrackAngleError(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRACK ANGLE ERROR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsEte Simvar
// args contain optional index
func SimVarGpsEte(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS ETE",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsEta Simvar
// args contain optional index
func SimVarGpsEta(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsWpNextLat Simvar
// args contain optional index
func SimVarGpsWpNextLat(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpNextLon Simvar
// args contain optional index
func SimVarGpsWpNextLon(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpNextAlt Simvar
// args contain optional index
func SimVarGpsWpNextAlt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpPrevValid Simvar
// args contain optional index
func SimVarGpsWpPrevValid(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsWpPrevLat Simvar
// args contain optional index
func SimVarGpsWpPrevLat(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpPrevLon Simvar
// args contain optional index
func SimVarGpsWpPrevLon(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpPrevAlt Simvar
// args contain optional index
func SimVarGpsWpPrevAlt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpEte Simvar
// args contain optional index
func SimVarGpsWpEte(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP ETE",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsWpEta Simvar
// args contain optional index
func SimVarGpsWpEta(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsCourseToSteer Simvar
// args contain optional index
func SimVarGpsCourseToSteer(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS COURSE TO STEER",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpIndex Simvar
// args contain optional index
func SimVarGpsFlightPlanWpIndex(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS FLIGHT PLAN WP INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpCount Simvar
// args contain optional index
func SimVarGpsFlightPlanWpCount(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS FLIGHT PLAN WP COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsIsActiveWpLocked Simvar
// args contain optional index
func SimVarGpsIsActiveWpLocked(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE WP LOCKED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsApproachLoaded Simvar
// args contain optional index
func SimVarGpsIsApproachLoaded(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS APPROACH LOADED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsApproachActive Simvar
// args contain optional index
func SimVarGpsIsApproachActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS IS APPROACH ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachMode Simvar
// args contain optional index
func SimVarGpsApproachMode(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH MODE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachWpType Simvar
// args contain optional index
func SimVarGpsApproachWpType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachIsWpRunway Simvar
// args contain optional index
func SimVarGpsApproachIsWpRunway(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS WP RUNWAY",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachSegmentType Simvar
// args contain optional index
func SimVarGpsApproachSegmentType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH SEGMENT TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachApproachIndex Simvar
// args contain optional index
func SimVarGpsApproachApproachIndex(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachApproachType Simvar
// args contain optional index
func SimVarGpsApproachApproachType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachTransitionIndex Simvar
// args contain optional index
func SimVarGpsApproachTransitionIndex(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TRANSITION INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachIsFinal Simvar
// args contain optional index
func SimVarGpsApproachIsFinal(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS FINAL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachIsMissed Simvar
// args contain optional index
func SimVarGpsApproachIsMissed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS MISSED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachTimezoneDeviation Simvar
// args contain optional index
func SimVarGpsApproachTimezoneDeviation(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TIMEZONE DEVIATION",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsApproachWpIndex Simvar
// args contain optional index
func SimVarGpsApproachWpIndex(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachWpCount Simvar
// args contain optional index
func SimVarGpsApproachWpCount(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsDrivesNav1 Simvar
func SimVarGpsDrivesNav1(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS DRIVES NAV1",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComReceiveAll Simvar
// args contain optional index
func SimVarComReceiveAll(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM RECEIVE ALL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComAvailable Simvar
// args contain optional index
func SimVarComAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComTest Simvar
// args contain optional index
func SimVarComTest(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "COM TEST:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTransponderAvailable Simvar
// args contain optional index
func SimVarTransponderAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRANSPONDER AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfAvailable Simvar
// args contain optional index
func SimVarAdfAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfFrequency Simvar
// args contain optional index
func SimVarAdfFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarAdfExtFrequency Simvar
// args contain optional index
func SimVarAdfExtFrequency(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF EXT FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarAdfIdent Simvar
// args contain optional index
func SimVarAdfIdent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF IDENT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAdfName Simvar
// args contain optional index
func SimVarAdfName(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ADF NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavIdent Simvar
// args contain optional index
func SimVarNavIdent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV IDENT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavName Simvar
// args contain optional index
func SimVarNavName(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavCodes Simvar
// args contain optional index
func SimVarNavCodes(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV CODES:index",
		Units:    "Flags",
		Settable: false,
	}
}

// SimVarNavGlideSlope Simvar
// args contain optional index
func SimVarNavGlideSlope(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GLIDE SLOPE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavRelativeBearingToStation Simvar
// args contain optional index
func SimVarNavRelativeBearingToStation(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV RELATIVE BEARING TO STATION:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarSelectedDme Simvar
// args contain optional index
func SimVarSelectedDme(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SELECTED DME",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsWpNextId Simvar
// args contain optional index
func SimVarGpsWpNextId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsWpPrevId Simvar
// args contain optional index
func SimVarGpsWpPrevId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsTargetDistance Simvar
// args contain optional index
func SimVarGpsTargetDistance(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS TARGET DISTANCE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsTargetAltitude Simvar
// args contain optional index
func SimVarGpsTargetAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS TARGET ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarYokeYPosition Simvar
// args contain optional index
func SimVarYokeYPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YOKE Y POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarYokeXPosition Simvar
// args contain optional index
func SimVarYokeXPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YOKE X POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRudderPedalPosition Simvar
// args contain optional index
func SimVarRudderPedalPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER PEDAL POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRudderPosition Simvar
// args contain optional index
func SimVarRudderPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarElevatorPosition Simvar
// args contain optional index
func SimVarElevatorPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarAileronPosition Simvar
// args contain optional index
func SimVarAileronPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarElevatorTrimPosition Simvar
// args contain optional index
func SimVarElevatorTrimPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM POSITION",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarElevatorTrimIndicator Simvar
// args contain optional index
func SimVarElevatorTrimIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarElevatorTrimPct Simvar
// args contain optional index
func SimVarElevatorTrimPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarBrakeLeftPosition Simvar
// args contain optional index
func SimVarBrakeLeftPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE LEFT POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeRightPosition Simvar
// args contain optional index
func SimVarBrakeRightPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE RIGHT POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeIndicator Simvar
// args contain optional index
func SimVarBrakeIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarBrakeParkingPosition Simvar
// args contain optional index
func SimVarBrakeParkingPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE PARKING POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeParkingIndicator Simvar
// args contain optional index
func SimVarBrakeParkingIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE PARKING INDICATOR",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilersArmed Simvar
// args contain optional index
func SimVarSpoilersArmed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SPOILERS ARMED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilersHandlePosition Simvar
// args contain optional index
func SimVarSpoilersHandlePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SPOILERS HANDLE POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarSpoilersLeftPosition Simvar
// args contain optional index
func SimVarSpoilersLeftPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SPOILERS LEFT POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarSpoilersRightPosition Simvar
// args contain optional index
func SimVarSpoilersRightPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SPOILERS RIGHT POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFlapsHandlePercent Simvar
// args contain optional index
func SimVarFlapsHandlePercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAPS HANDLE PERCENT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFlapsHandleIndex Simvar
// args contain optional index
func SimVarFlapsHandleIndex(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAPS HANDLE INDEX",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarFlapsNumHandlePositions Simvar
// args contain optional index
func SimVarFlapsNumHandlePositions(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAPS NUM HANDLE POSITIONS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsLeftPercent Simvar
// args contain optional index
func SimVarTrailingEdgeFlapsLeftPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsRightPercent Simvar
// args contain optional index
func SimVarTrailingEdgeFlapsRightPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsLeftAngle Simvar
// args contain optional index
func SimVarTrailingEdgeFlapsLeftAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS LEFT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsRightAngle Simvar
// args contain optional index
func SimVarTrailingEdgeFlapsRightAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS RIGHT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsLeftPercent Simvar
// args contain optional index
func SimVarLeadingEdgeFlapsLeftPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsRightPercent Simvar
// args contain optional index
func SimVarLeadingEdgeFlapsRightPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsLeftAngle Simvar
// args contain optional index
func SimVarLeadingEdgeFlapsLeftAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS LEFT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsRightAngle Simvar
// args contain optional index
func SimVarLeadingEdgeFlapsRightAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS RIGHT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarIsGearRetractable Simvar
// args contain optional index
func SimVarIsGearRetractable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS GEAR RETRACTABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearSkis Simvar
// args contain optional index
func SimVarIsGearSkis(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS GEAR SKIS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearFloats Simvar
// args contain optional index
func SimVarIsGearFloats(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS GEAR FLOATS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearSkids Simvar
// args contain optional index
func SimVarIsGearSkids(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS GEAR SKIDS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearWheels Simvar
// args contain optional index
func SimVarIsGearWheels(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS GEAR WHEELS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearHandlePosition Simvar
// args contain optional index
func SimVarGearHandlePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR HANDLE POSITION",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGearHydraulicPressure Simvar
// args contain optional index
func SimVarGearHydraulicPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR HYDRAULIC PRESSURE",
		Units:    "psf",
		Settable: false,
	}
}

// SimVarTailwheelLockOn Simvar
// args contain optional index
func SimVarTailwheelLockOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TAILWHEEL LOCK ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearCenterPosition Simvar
// args contain optional index
func SimVarGearCenterPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearLeftPosition Simvar
// args contain optional index
func SimVarGearLeftPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearRightPosition Simvar
// args contain optional index
func SimVarGearRightPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearTailPosition Simvar
// args contain optional index
func SimVarGearTailPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR TAIL POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxPosition Simvar
// args contain optional index
func SimVarGearAuxPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearPosition Simvar
// args contain optional index
func SimVarGearPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR POSITION:index",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarGearAnimationPosition Simvar
// args contain optional index
func SimVarGearAnimationPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR ANIMATION POSITION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGearTotalPctExtended Simvar
// args contain optional index
func SimVarGearTotalPctExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR TOTAL PCT EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarAutoBrakeSwitchCb Simvar
// args contain optional index
func SimVarAutoBrakeSwitchCb(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTO BRAKE SWITCH CB",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarWaterRudderHandlePosition Simvar
// args contain optional index
func SimVarWaterRudderHandlePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER RUDDER HANDLE POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarElevatorDeflection Simvar
// args contain optional index
func SimVarElevatorDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarElevatorDeflectionPct Simvar
// args contain optional index
func SimVarElevatorDeflectionPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderExtended Simvar
// args contain optional index
func SimVarWaterLeftRudderExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarWaterRightRudderExtended Simvar
// args contain optional index
func SimVarWaterRightRudderExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarGearCenterSteerAngle Simvar
// args contain optional index
func SimVarGearCenterSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearLeftSteerAngle Simvar
// args contain optional index
func SimVarGearLeftSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearRightSteerAngle Simvar
// args contain optional index
func SimVarGearRightSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxSteerAngle Simvar
// args contain optional index
func SimVarGearAuxSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearSteerAngle Simvar
// args contain optional index
func SimVarGearSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR STEER ANGLE:index",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAngle Simvar
// args contain optional index
func SimVarWaterLeftRudderSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAngle Simvar
// args contain optional index
func SimVarWaterRightRudderSteerAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearCenterSteerAnglePct Simvar
// args contain optional index
func SimVarGearCenterSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearLeftSteerAnglePct Simvar
// args contain optional index
func SimVarGearLeftSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearRightSteerAnglePct Simvar
// args contain optional index
func SimVarGearRightSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxSteerAnglePct Simvar
// args contain optional index
func SimVarGearAuxSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearSteerAnglePct Simvar
// args contain optional index
func SimVarGearSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR STEER ANGLE PCT:index",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAnglePct Simvar
// args contain optional index
func SimVarWaterLeftRudderSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAnglePct Simvar
// args contain optional index
func SimVarWaterRightRudderSteerAnglePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronLeftDeflection Simvar
// args contain optional index
func SimVarAileronLeftDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON LEFT DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronLeftDeflectionPct Simvar
// args contain optional index
func SimVarAileronLeftDeflectionPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON LEFT DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronRightDeflection Simvar
// args contain optional index
func SimVarAileronRightDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON RIGHT DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronRightDeflectionPct Simvar
// args contain optional index
func SimVarAileronRightDeflectionPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON RIGHT DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronAverageDeflection Simvar
// args contain optional index
func SimVarAileronAverageDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON AVERAGE DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronTrim Simvar
// args contain optional index
func SimVarAileronTrim(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AILERON TRIM",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRudderDeflection Simvar
// args contain optional index
func SimVarRudderDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRudderDeflectionPct Simvar
// args contain optional index
func SimVarRudderDeflectionPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRudderTrim Simvar
// args contain optional index
func SimVarRudderTrim(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER TRIM",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarFlapsAvailable Simvar
// args contain optional index
func SimVarFlapsAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAPS AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearDamageBySpeed Simvar
// args contain optional index
func SimVarGearDamageBySpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR DAMAGE BY SPEED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearSpeedExceeded Simvar
// args contain optional index
func SimVarGearSpeedExceeded(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR SPEED EXCEEDED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlapDamageBySpeed Simvar
// args contain optional index
func SimVarFlapDamageBySpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAP DAMAGE BY SPEED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlapSpeedExceeded Simvar
// args contain optional index
func SimVarFlapSpeedExceeded(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FLAP SPEED EXCEEDED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCenterWheelRpm Simvar
// args contain optional index
func SimVarCenterWheelRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CENTER WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarLeftWheelRpm Simvar
// args contain optional index
func SimVarLeftWheelRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEFT WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarRightWheelRpm Simvar
// args contain optional index
func SimVarRightWheelRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RIGHT WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarAutopilotAvailable Simvar
// args contain optional index
func SimVarAutopilotAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMaster Simvar
// args contain optional index
func SimVarAutopilotMaster(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MASTER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotNavSelected Simvar
// args contain optional index
func SimVarAutopilotNavSelected(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT NAV SELECTED",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotWingLeveler Simvar
// args contain optional index
func SimVarAutopilotWingLeveler(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT WING LEVELER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotHeadingLock Simvar
// args contain optional index
func SimVarAutopilotHeadingLock(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT HEADING LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotHeadingLockDir Simvar
// args contain optional index
func SimVarAutopilotHeadingLockDir(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT HEADING LOCK DIR",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLock Simvar
// args contain optional index
func SimVarAutopilotAltitudeLock(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ALTITUDE LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLockVar Simvar
// args contain optional index
func SimVarAutopilotAltitudeLockVar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ALTITUDE LOCK VAR",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarAutopilotAttitudeHold Simvar
// args contain optional index
func SimVarAutopilotAttitudeHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ATTITUDE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotGlideslopeHold Simvar
// args contain optional index
func SimVarAutopilotGlideslopeHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT GLIDESLOPE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotPitchHoldRef Simvar
// args contain optional index
func SimVarAutopilotPitchHoldRef(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT PITCH HOLD REF",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotApproachHold Simvar
// args contain optional index
func SimVarAutopilotApproachHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT APPROACH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotBackcourseHold Simvar
// args contain optional index
func SimVarAutopilotBackcourseHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT BACKCOURSE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotVerticalHoldVar Simvar
// args contain optional index
func SimVarAutopilotVerticalHoldVar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT VERTICAL HOLD VAR",
		Units:    "Feet/minute",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorActive Simvar
// args contain optional index
func SimVarAutopilotFlightDirectorActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorPitch Simvar
// args contain optional index
func SimVarAutopilotFlightDirectorPitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR PITCH",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorBank Simvar
// args contain optional index
func SimVarAutopilotFlightDirectorBank(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR BANK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHold Simvar
// args contain optional index
func SimVarAutopilotAirspeedHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AIRSPEED HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHoldVar Simvar
// args contain optional index
func SimVarAutopilotAirspeedHoldVar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AIRSPEED HOLD VAR",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAutopilotMachHold Simvar
// args contain optional index
func SimVarAutopilotMachHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MACH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMachHoldVar Simvar
// args contain optional index
func SimVarAutopilotMachHoldVar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MACH HOLD VAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotYawDamper Simvar
// args contain optional index
func SimVarAutopilotYawDamper(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT YAW DAMPER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotRpmHoldVar Simvar
// args contain optional index
func SimVarAutopilotRpmHoldVar(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT RPM HOLD VAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotThrottleArm Simvar
// args contain optional index
func SimVarAutopilotThrottleArm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT THROTTLE ARM",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotTakeoffPowerActive Simvar
// args contain optional index
func SimVarAutopilotTakeoffPowerActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT TAKEOFF POWER ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutothrottleActive Simvar
// args contain optional index
func SimVarAutothrottleActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOTHROTTLE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotNav1Lock Simvar
func SimVarAutopilotNav1Lock(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT NAV1 LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotVerticalHold Simvar
// args contain optional index
func SimVarAutopilotVerticalHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT VERTICAL HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotRpmHold Simvar
// args contain optional index
func SimVarAutopilotRpmHold(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT RPM HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMaxBank Simvar
// args contain optional index
func SimVarAutopilotMaxBank(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MAX BANK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarWheelRpm Simvar
// args contain optional index
func SimVarWheelRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarAuxWheelRpm Simvar
// args contain optional index
func SimVarAuxWheelRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUX WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarWheelRotationAngle Simvar
// args contain optional index
func SimVarWheelRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCenterWheelRotationAngle Simvar
// args contain optional index
func SimVarCenterWheelRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CENTER WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeftWheelRotationAngle Simvar
// args contain optional index
func SimVarLeftWheelRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LEFT WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRightWheelRotationAngle Simvar
// args contain optional index
func SimVarRightWheelRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RIGHT WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAuxWheelRotationAngle Simvar
// args contain optional index
func SimVarAuxWheelRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUX WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGearEmergencyHandlePosition Simvar
// args contain optional index
func SimVarGearEmergencyHandlePosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR EMERGENCY HANDLE POSITION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearWarning Simvar
// args contain optional index
func SimVarGearWarning(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GEAR WARNING",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarAntiskidBrakesActive Simvar
// args contain optional index
func SimVarAntiskidBrakesActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ANTISKID BRAKES ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRetractFloatSwitch Simvar
// args contain optional index
func SimVarRetractFloatSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RETRACT FLOAT SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRetractLeftFloatExtended Simvar
// args contain optional index
func SimVarRetractLeftFloatExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RETRACT LEFT FLOAT EXTENDED",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarRetractRightFloatExtended Simvar
// args contain optional index
func SimVarRetractRightFloatExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RETRACT RIGHT FLOAT EXTENDED",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarSteerInputControl Simvar
// args contain optional index
func SimVarSteerInputControl(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STEER INPUT CONTROL",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarAmbientDensity Simvar
// args contain optional index
func SimVarAmbientDensity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT DENSITY",
		Units:    "Slugs per cubic feet",
		Settable: false,
	}
}

// SimVarAmbientTemperature Simvar
// args contain optional index
func SimVarAmbientTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT TEMPERATURE",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarAmbientPressure Simvar
// args contain optional index
func SimVarAmbientPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT PRESSURE",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarAmbientWindVelocity Simvar
// args contain optional index
func SimVarAmbientWindVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND VELOCITY",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAmbientWindDirection Simvar
// args contain optional index
func SimVarAmbientWindDirection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND DIRECTION",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAmbientWindX Simvar
// args contain optional index
func SimVarAmbientWindX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND X",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientWindY Simvar
// args contain optional index
func SimVarAmbientWindY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND Y",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientWindZ Simvar
// args contain optional index
func SimVarAmbientWindZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND Z",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientPrecipState Simvar
// args contain optional index
func SimVarAmbientPrecipState(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT PRECIP STATE",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarAircraftWindX Simvar
// args contain optional index
func SimVarAircraftWindX(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND X",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAircraftWindY Simvar
// args contain optional index
func SimVarAircraftWindY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND Y",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAircraftWindZ Simvar
// args contain optional index
func SimVarAircraftWindZ(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND Z",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarBarometerPressure Simvar
// args contain optional index
func SimVarBarometerPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BAROMETER PRESSURE",
		Units:    "Millibars",
		Settable: false,
	}
}

// SimVarSeaLevelPressure Simvar
// args contain optional index
func SimVarSeaLevelPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SEA LEVEL PRESSURE",
		Units:    "Millibars",
		Settable: false,
	}
}

// SimVarTotalAirTemperature Simvar
// args contain optional index
func SimVarTotalAirTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL AIR TEMPERATURE",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarWindshieldRainEffectAvailable Simvar
// args contain optional index
func SimVarWindshieldRainEffectAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WINDSHIELD RAIN EFFECT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAmbientInCloud Simvar
// args contain optional index
func SimVarAmbientInCloud(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT IN CLOUD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAmbientVisibility Simvar
// args contain optional index
func SimVarAmbientVisibility(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AMBIENT VISIBILITY",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarStandardAtmTemperature Simvar
// args contain optional index
func SimVarStandardAtmTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STANDARD ATM TEMPERATURE",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarRotorBrakeHandlePos Simvar
// args contain optional index
func SimVarRotorBrakeHandlePos(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR BRAKE HANDLE POS",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRotorBrakeActive Simvar
// args contain optional index
func SimVarRotorBrakeActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR BRAKE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorClutchSwitchPos Simvar
// args contain optional index
func SimVarRotorClutchSwitchPos(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR CLUTCH SWITCH POS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorClutchActive Simvar
// args contain optional index
func SimVarRotorClutchActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR CLUTCH ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorTemperature Simvar
// args contain optional index
func SimVarRotorTemperature(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR TEMPERATURE",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarRotorChipDetected Simvar
// args contain optional index
func SimVarRotorChipDetected(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR CHIP DETECTED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorGovSwitchPos Simvar
// args contain optional index
func SimVarRotorGovSwitchPos(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR GOV SWITCH POS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorGovActive Simvar
// args contain optional index
func SimVarRotorGovActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR GOV ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorLateralTrimPct Simvar
// args contain optional index
func SimVarRotorLateralTrimPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR LATERAL TRIM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRotorRpmPct Simvar
// args contain optional index
func SimVarRotorRpmPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR RPM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarSmokeEnable Simvar
// args contain optional index
func SimVarSmokeEnable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SMOKE ENABLE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarSmokesystemAvailable Simvar
// args contain optional index
func SimVarSmokesystemAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SMOKESYSTEM AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPitotHeat Simvar
// args contain optional index
func SimVarPitotHeat(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PITOT HEAT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFoldingWingLeftPercent Simvar
// args contain optional index
func SimVarFoldingWingLeftPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FOLDING WING LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFoldingWingRightPercent Simvar
// args contain optional index
func SimVarFoldingWingRightPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FOLDING WING RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarCanopyOpen Simvar
// args contain optional index
func SimVarCanopyOpen(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CANOPY OPEN",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTailhookPosition Simvar
// args contain optional index
func SimVarTailhookPosition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TAILHOOK POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarExitOpen Simvar
// args contain optional index
func SimVarExitOpen(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EXIT OPEN:index",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarStallHornAvailable Simvar
// args contain optional index
func SimVarStallHornAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STALL HORN AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngineMixureAvailable Simvar
// args contain optional index
func SimVarEngineMixureAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ENGINE MIXURE AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCarbHeatAvailable Simvar
// args contain optional index
func SimVarCarbHeatAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CARB HEAT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilerAvailable Simvar
// args contain optional index
func SimVarSpoilerAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SPOILER AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsTailDragger Simvar
// args contain optional index
func SimVarIsTailDragger(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS TAIL DRAGGER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarStrobesAvailable Simvar
// args contain optional index
func SimVarStrobesAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STROBES AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarToeBrakesAvailable Simvar
// args contain optional index
func SimVarToeBrakesAvailable(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOE BRAKES AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPushbackState Simvar
// args contain optional index
func SimVarPushbackState(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK STATE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarElectricalMasterBattery Simvar
// args contain optional index
func SimVarElectricalMasterBattery(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MASTER BATTERY",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarElectricalTotalLoadAmps Simvar
// args contain optional index
func SimVarElectricalTotalLoadAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL TOTAL LOAD AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryLoad Simvar
// args contain optional index
func SimVarElectricalBatteryLoad(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY LOAD",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryVoltage Simvar
// args contain optional index
func SimVarElectricalBatteryVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalMainBusVoltage Simvar
// args contain optional index
func SimVarElectricalMainBusVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MAIN BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalMainBusAmps Simvar
// args contain optional index
func SimVarElectricalMainBusAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MAIN BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusVoltage Simvar
// args contain optional index
func SimVarElectricalAvionicsBusVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL AVIONICS BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusAmps Simvar
// args contain optional index
func SimVarElectricalAvionicsBusAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL AVIONICS BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusVoltage Simvar
// args contain optional index
func SimVarElectricalHotBatteryBusVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL HOT BATTERY BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusAmps Simvar
// args contain optional index
func SimVarElectricalHotBatteryBusAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL HOT BATTERY BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryBusVoltage Simvar
// args contain optional index
func SimVarElectricalBatteryBusVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalBatteryBusAmps Simvar
// args contain optional index
func SimVarElectricalBatteryBusAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalGenaltBusVoltage Simvar
// args contain optional index
func SimVarElectricalGenaltBusVoltage(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL GENALT BUS VOLTAGE:index",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalGenaltBusAmps Simvar
// args contain optional index
func SimVarElectricalGenaltBusAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL GENALT BUS AMPS:index",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarCircuitGeneralPanelOn Simvar
// args contain optional index
func SimVarCircuitGeneralPanelOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GENERAL PANEL ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitFlapMotorOn Simvar
// args contain optional index
func SimVarCircuitFlapMotorOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT FLAP MOTOR ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitGearMotorOn Simvar
// args contain optional index
func SimVarCircuitGearMotorOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GEAR MOTOR ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutopilotOn Simvar
// args contain optional index
func SimVarCircuitAutopilotOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTOPILOT ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAvionicsOn Simvar
// args contain optional index
func SimVarCircuitAvionicsOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AVIONICS ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitPitotHeatOn Simvar
// args contain optional index
func SimVarCircuitPitotHeatOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT PITOT HEAT ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitPropSyncOn Simvar
// args contain optional index
func SimVarCircuitPropSyncOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT PROP SYNC ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutoFeatherOn Simvar
// args contain optional index
func SimVarCircuitAutoFeatherOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTO FEATHER ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutoBrakesOn Simvar
// args contain optional index
func SimVarCircuitAutoBrakesOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTO BRAKES ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitStandyVacuumOn Simvar
// args contain optional index
func SimVarCircuitStandyVacuumOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT STANDY VACUUM ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitMarkerBeaconOn Simvar
// args contain optional index
func SimVarCircuitMarkerBeaconOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT MARKER BEACON ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitGearWarningOn Simvar
// args contain optional index
func SimVarCircuitGearWarningOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GEAR WARNING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitHydraulicPumpOn Simvar
// args contain optional index
func SimVarCircuitHydraulicPumpOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT HYDRAULIC PUMP ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHydraulicPressure Simvar
// args contain optional index
func SimVarHydraulicPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC PRESSURE:index",
		Units:    "Pound force per square foot",
		Settable: false,
	}
}

// SimVarHydraulicReservoirPercent Simvar
// args contain optional index
func SimVarHydraulicReservoirPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC RESERVOIR PERCENT:index",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarHydraulicSystemIntegrity Simvar
// args contain optional index
func SimVarHydraulicSystemIntegrity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC SYSTEM INTEGRITY",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarStructuralDeiceSwitch Simvar
// args contain optional index
func SimVarStructuralDeiceSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCTURAL DEICE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTotalWeight Simvar
// args contain optional index
func SimVarTotalWeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarMaxGrossWeight Simvar
// args contain optional index
func SimVarMaxGrossWeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MAX GROSS WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarEmptyWeight Simvar
// args contain optional index
func SimVarEmptyWeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarIsUserSim Simvar
// args contain optional index
func SimVarIsUserSim(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS USER SIM",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSimDisabled Simvar
// args contain optional index
func SimVarSimDisabled(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SIM DISABLED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGForce Simvar
// args contain optional index
func SimVarGForce(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "G FORCE",
		Units:    "GForce",
		Settable: true,
	}
}

// SimVarAtcHeavy Simvar
// args contain optional index
func SimVarAtcHeavy(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC HEAVY",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAutoCoordination Simvar
// args contain optional index
func SimVarAutoCoordination(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AUTO COORDINATION",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRealism Simvar
// args contain optional index
func SimVarRealism(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "REALISM",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarTrueAirspeedSelected Simvar
// args contain optional index
func SimVarTrueAirspeedSelected(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TRUE AIRSPEED SELECTED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarDesignSpeedVc Simvar
// args contain optional index
func SimVarDesignSpeedVc(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VC",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarMinDragVelocity Simvar
// args contain optional index
func SimVarMinDragVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MIN DRAG VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarEstimatedCruiseSpeed Simvar
// args contain optional index
func SimVarEstimatedCruiseSpeed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ESTIMATED CRUISE SPEED",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarCgPercent Simvar
// args contain optional index
func SimVarCgPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG PERCENT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgPercentLateral Simvar
// args contain optional index
func SimVarCgPercentLateral(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG PERCENT LATERAL",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarIsSlewActive Simvar
// args contain optional index
func SimVarIsSlewActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS SLEW ACTIVE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarIsSlewAllowed Simvar
// args contain optional index
func SimVarIsSlewAllowed(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS SLEW ALLOWED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAtcSuggestedMinRwyTakeoff Simvar
// args contain optional index
func SimVarAtcSuggestedMinRwyTakeoff(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC SUGGESTED MIN RWY TAKEOFF",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarAtcSuggestedMinRwyLanding Simvar
// args contain optional index
func SimVarAtcSuggestedMinRwyLanding(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC SUGGESTED MIN RWY LANDING",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPayloadStationWeight Simvar
// args contain optional index
func SimVarPayloadStationWeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION WEIGHT:index",
		Units:    "Pounds",
		Settable: true,
	}
}

// SimVarPayloadStationCount Simvar
// args contain optional index
func SimVarPayloadStationCount(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUserInputEnabled Simvar
// args contain optional index
func SimVarUserInputEnabled(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "USER INPUT ENABLED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarTypicalDescentRate Simvar
// args contain optional index
func SimVarTypicalDescentRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TYPICAL DESCENT RATE",
		Units:    "Feet per minute",
		Settable: false,
	}
}

// SimVarVisualModelRadius Simvar
// args contain optional index
func SimVarVisualModelRadius(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VISUAL MODEL RADIUS",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarCategory Simvar
// args contain optional index
func SimVarCategory(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CATEGORY",
		Units:    "String",
		Settable: false,
	}
}

// SimVarSigmaSqrt Simvar
// args contain optional index
func SimVarSigmaSqrt(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SIGMA SQRT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarDynamicPressure Simvar
// args contain optional index
func SimVarDynamicPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DYNAMIC PRESSURE",
		Units:    "Pounds per square foot",
		Settable: false,
	}
}

// SimVarTotalVelocity Simvar
// args contain optional index
func SimVarTotalVelocity(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarAirspeedSelectIndicatedOrTrue Simvar
// args contain optional index
func SimVarAirspeedSelectIndicatedOrTrue(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED SELECT INDICATED OR TRUE",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarVariometerRate Simvar
// args contain optional index
func SimVarVariometerRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VARIOMETER RATE",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarVariometerSwitch Simvar
// args contain optional index
func SimVarVariometerSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "VARIOMETER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarDesignSpeedVs0 Simvar
func SimVarDesignSpeedVs0(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VS0",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarDesignSpeedVs1 Simvar
func SimVarDesignSpeedVs1(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VS1",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPressureAltitude Simvar
// args contain optional index
func SimVarPressureAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURE ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarMagneticCompass Simvar
// args contain optional index
func SimVarMagneticCompass(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MAGNETIC COMPASS",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarTurnIndicatorRate Simvar
// args contain optional index
func SimVarTurnIndicatorRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURN INDICATOR RATE",
		Units:    "Radians per second",
		Settable: false,
	}
}

// SimVarTurnIndicatorSwitch Simvar
// args contain optional index
func SimVarTurnIndicatorSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TURN INDICATOR SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarYokeYIndicator Simvar
// args contain optional index
func SimVarYokeYIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YOKE Y INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarYokeXIndicator Simvar
// args contain optional index
func SimVarYokeXIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YOKE X INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarRudderPedalIndicator Simvar
// args contain optional index
func SimVarRudderPedalIndicator(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RUDDER PEDAL INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarBrakeDependentHydraulicPressure Simvar
// args contain optional index
func SimVarBrakeDependentHydraulicPressure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BRAKE DEPENDENT HYDRAULIC PRESSURE",
		Units:    "foot pounds",
		Settable: false,
	}
}

// SimVarPanelAntiIceSwitch Simvar
// args contain optional index
func SimVarPanelAntiIceSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PANEL ANTI ICE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarWingArea Simvar
// args contain optional index
func SimVarWingArea(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WING AREA",
		Units:    "Square feet",
		Settable: false,
	}
}

// SimVarWingSpan Simvar
// args contain optional index
func SimVarWingSpan(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WING SPAN",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarBetaDot Simvar
// args contain optional index
func SimVarBetaDot(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BETA DOT",
		Units:    "Radians per second",
		Settable: false,
	}
}

// SimVarLinearClAlpha Simvar
// args contain optional index
func SimVarLinearClAlpha(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LINEAR CL ALPHA",
		Units:    "Per radian",
		Settable: false,
	}
}

// SimVarStallAlpha Simvar
// args contain optional index
func SimVarStallAlpha(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STALL ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarZeroLiftAlpha Simvar
// args contain optional index
func SimVarZeroLiftAlpha(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZERO LIFT ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCgAftLimit Simvar
// args contain optional index
func SimVarCgAftLimit(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG AFT LIMIT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgFwdLimit Simvar
// args contain optional index
func SimVarCgFwdLimit(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG FWD LIMIT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgMaxMach Simvar
// args contain optional index
func SimVarCgMaxMach(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG MAX MACH",
		Units:    "Machs",
		Settable: false,
	}
}

// SimVarCgMinMach Simvar
// args contain optional index
func SimVarCgMinMach(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CG MIN MACH",
		Units:    "Machs",
		Settable: false,
	}
}

// SimVarPayloadStationName Simvar
// args contain optional index
func SimVarPayloadStationName(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarElevonDeflection Simvar
// args contain optional index
func SimVarElevonDeflection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELEVON DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarExitType Simvar
// args contain optional index
func SimVarExitType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EXIT TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarExitPosx Simvar
// args contain optional index
func SimVarExitPosx(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EXIT POSX",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarExitPosy Simvar
// args contain optional index
func SimVarExitPosy(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EXIT POSY",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarExitPosz Simvar
// args contain optional index
func SimVarExitPosz(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EXIT POSZ",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarDecisionHeight Simvar
// args contain optional index
func SimVarDecisionHeight(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DECISION HEIGHT",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarDecisionAltitudeMsl Simvar
// args contain optional index
func SimVarDecisionAltitudeMsl(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DECISION ALTITUDE MSL",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarEmptyWeightPitchMoi Simvar
// args contain optional index
func SimVarEmptyWeightPitchMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT PITCH MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightRollMoi Simvar
// args contain optional index
func SimVarEmptyWeightRollMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT ROLL MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightYawMoi Simvar
// args contain optional index
func SimVarEmptyWeightYawMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT YAW MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightCrossCoupledMoi Simvar
// args contain optional index
func SimVarEmptyWeightCrossCoupledMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT CROSS COUPLED MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightPitchMoi Simvar
// args contain optional index
func SimVarTotalWeightPitchMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT PITCH MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightRollMoi Simvar
// args contain optional index
func SimVarTotalWeightRollMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT ROLL MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightYawMoi Simvar
// args contain optional index
func SimVarTotalWeightYawMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT YAW MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightCrossCoupledMoi Simvar
// args contain optional index
func SimVarTotalWeightCrossCoupledMoi(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT CROSS COUPLED MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarWaterBallastValve Simvar
// args contain optional index
func SimVarWaterBallastValve(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "WATER BALLAST VALVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarMaxRatedEngineRpm Simvar
// args contain optional index
func SimVarMaxRatedEngineRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MAX RATED ENGINE RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarFullThrottleThrustToWeightRatio Simvar
// args contain optional index
func SimVarFullThrottleThrustToWeightRatio(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FULL THROTTLE THRUST TO WEIGHT RATIO",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarPropAutoCruiseActive Simvar
// args contain optional index
func SimVarPropAutoCruiseActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP AUTO CRUISE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropRotationAngle Simvar
// args contain optional index
func SimVarPropRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMax Simvar
// args contain optional index
func SimVarPropBetaMax(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MAX",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMin Simvar
// args contain optional index
func SimVarPropBetaMin(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MIN",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMinReverse Simvar
// args contain optional index
func SimVarPropBetaMinReverse(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MIN REVERSE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarFuelSelectedTransferMode Simvar
// args contain optional index
func SimVarFuelSelectedTransferMode(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED TRANSFER MODE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarDroppableObjectsUiName Simvar
// args contain optional index
func SimVarDroppableObjectsUiName(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS UI NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarManualFuelPumpHandle Simvar
// args contain optional index
func SimVarManualFuelPumpHandle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MANUAL FUEL PUMP HANDLE",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarBleedAirSourceControl Simvar
// args contain optional index
func SimVarBleedAirSourceControl(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "BLEED AIR SOURCE CONTROL",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarElectricalOldChargingAmps Simvar
// args contain optional index
func SimVarElectricalOldChargingAmps(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL OLD CHARGING AMPS",
		Units:    "Amps",
		Settable: false,
	}
}

// SimVarHydraulicSwitch Simvar
// args contain optional index
func SimVarHydraulicSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarConcordeVisorNoseHandle Simvar
// args contain optional index
func SimVarConcordeVisorNoseHandle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CONCORDE VISOR NOSE HANDLE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarConcordeVisorPositionPercent Simvar
// args contain optional index
func SimVarConcordeVisorPositionPercent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CONCORDE VISOR POSITION PERCENT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarConcordeNoseAngle Simvar
// args contain optional index
func SimVarConcordeNoseAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CONCORDE NOSE ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRealismCrashWithOthers Simvar
// args contain optional index
func SimVarRealismCrashWithOthers(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "REALISM CRASH WITH OTHERS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRealismCrashDetection Simvar
// args contain optional index
func SimVarRealismCrashDetection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "REALISM CRASH DETECTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarManualInstrumentLights Simvar
// args contain optional index
func SimVarManualInstrumentLights(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "MANUAL INSTRUMENT LIGHTS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPitotIcePct Simvar
// args contain optional index
func SimVarPitotIcePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PITOT ICE PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorY Simvar
// args contain optional index
func SimVarSemibodyLoadfactorY(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SEMIBODY LOADFACTOR Y",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorYdot Simvar
// args contain optional index
func SimVarSemibodyLoadfactorYdot(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SEMIBODY LOADFACTOR YDOT",
		Units:    "Per second",
		Settable: false,
	}
}

// SimVarRadInsSwitch Simvar
// args contain optional index
func SimVarRadInsSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "RAD INS SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSimulatedRadius Simvar
// args contain optional index
func SimVarSimulatedRadius(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SIMULATED RADIUS",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarStructuralIcePct Simvar
// args contain optional index
func SimVarStructuralIcePct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STRUCTURAL ICE PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarArtificialGroundElevation Simvar
// args contain optional index
func SimVarArtificialGroundElevation(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ARTIFICIAL GROUND ELEVATION",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarSurfaceInfoValid Simvar
// args contain optional index
func SimVarSurfaceInfoValid(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SURFACE INFO VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSurfaceCondition Simvar
// args contain optional index
func SimVarSurfaceCondition(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SURFACE CONDITION",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPushbackAngle Simvar
// args contain optional index
func SimVarPushbackAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPushbackContactx Simvar
// args contain optional index
func SimVarPushbackContactx(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTX",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackContacty Simvar
// args contain optional index
func SimVarPushbackContacty(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTY",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackContactz Simvar
// args contain optional index
func SimVarPushbackContactz(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTZ",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackWait Simvar
// args contain optional index
func SimVarPushbackWait(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK WAIT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarYawStringAngle Simvar
// args contain optional index
func SimVarYawStringAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YAW STRING ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarYawStringPctExtended Simvar
// args contain optional index
func SimVarYawStringPctExtended(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "YAW STRING PCT EXTENDED",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarInductorCompassPercentDeviation Simvar
// args contain optional index
func SimVarInductorCompassPercentDeviation(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INDUCTOR COMPASS PERCENT DEVIATION",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarInductorCompassHeadingRef Simvar
// args contain optional index
func SimVarInductorCompassHeadingRef(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "INDUCTOR COMPASS HEADING REF",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAnemometerPctRpm Simvar
// args contain optional index
func SimVarAnemometerPctRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ANEMOMETER PCT RPM",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRotorRotationAngle Simvar
// args contain optional index
func SimVarRotorRotationAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ROTOR ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskPitchAngle Simvar
// args contain optional index
func SimVarDiskPitchAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DISK PITCH ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskBankAngle Simvar
// args contain optional index
func SimVarDiskBankAngle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DISK BANK ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskPitchPct Simvar
// args contain optional index
func SimVarDiskPitchPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DISK PITCH PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarDiskBankPct Simvar
// args contain optional index
func SimVarDiskBankPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DISK BANK PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarDiskConingPct Simvar
// args contain optional index
func SimVarDiskConingPct(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "DISK CONING PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarNavVorLlaf64 Simvar
func SimVarNavVorLlaf64(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV VOR LLAF64",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavGsLlaf64 Simvar
func SimVarNavGsLlaf64(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "NAV GS LLAF64",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStaticCgToGround Simvar
// args contain optional index
func SimVarStaticCgToGround(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STATIC CG TO GROUND",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarStaticPitch Simvar
// args contain optional index
func SimVarStaticPitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "STATIC PITCH",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCrashSequence Simvar
// args contain optional index
func SimVarCrashSequence(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CRASH SEQUENCE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarCrashFlag Simvar
// args contain optional index
func SimVarCrashFlag(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CRASH FLAG",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarTowReleaseHandle Simvar
// args contain optional index
func SimVarTowReleaseHandle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOW RELEASE HANDLE",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarTowConnection Simvar
// args contain optional index
func SimVarTowConnection(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TOW CONNECTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuPctRpm Simvar
// args contain optional index
func SimVarApuPctRpm(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU PCT RPM",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarApuPctStarter Simvar
// args contain optional index
func SimVarApuPctStarter(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU PCT STARTER",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarApuVolts Simvar
// args contain optional index
func SimVarApuVolts(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU VOLTS",
		Units:    "Volts",
		Settable: false,
	}
}

// SimVarApuGeneratorSwitch Simvar
// args contain optional index
func SimVarApuGeneratorSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU GENERATOR SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuGeneratorActive Simvar
// args contain optional index
func SimVarApuGeneratorActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU GENERATOR ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuOnFireDetected Simvar
// args contain optional index
func SimVarApuOnFireDetected(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "APU ON FIRE DETECTED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitude Simvar
// args contain optional index
func SimVarPressurizationCabinAltitude(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeGoal Simvar
// args contain optional index
func SimVarPressurizationCabinAltitudeGoal(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE GOAL",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeRate Simvar
// args contain optional index
func SimVarPressurizationCabinAltitudeRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE RATE",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPressurizationPressureDifferential Simvar
// args contain optional index
func SimVarPressurizationPressureDifferential(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION PRESSURE DIFFERENTIAL",
		Units:    "foot pounds",
		Settable: false,
	}
}

// SimVarPressurizationDumpSwitch Simvar
// args contain optional index
func SimVarPressurizationDumpSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION DUMP SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFireBottleSwitch Simvar
// args contain optional index
func SimVarFireBottleSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FIRE BOTTLE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFireBottleDischarged Simvar
// args contain optional index
func SimVarFireBottleDischarged(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "FIRE BOTTLE DISCHARGED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCabinNoSmokingAlertSwitch Simvar
// args contain optional index
func SimVarCabinNoSmokingAlertSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CABIN NO SMOKING ALERT SWITCH",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarCabinSeatbeltsAlertSwitch Simvar
// args contain optional index
func SimVarCabinSeatbeltsAlertSwitch(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "CABIN SEATBELTS ALERT SWITCH",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGpwsWarning Simvar
// args contain optional index
func SimVarGpwsWarning(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPWS WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpwsSystemActive Simvar
// args contain optional index
func SimVarGpwsSystemActive(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPWS SYSTEM ACTIVE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarIsLatitudeLongitudeFreezeOn Simvar
// args contain optional index
func SimVarIsLatitudeLongitudeFreezeOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS LATITUDE LONGITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAltitudeFreezeOn Simvar
// args contain optional index
func SimVarIsAltitudeFreezeOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS ALTITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAttitudeFreezeOn Simvar
// args contain optional index
func SimVarIsAttitudeFreezeOn(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "IS ATTITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAtcType Simvar
// args contain optional index
func SimVarAtcType(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC TYPE",
		Units:    "String64",
		Settable: false,
	}
}

// SimVarAtcModel Simvar
// args contain optional index
func SimVarAtcModel(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC MODEL",
		Units:    "String64",
		Settable: false,
	}
}

// SimVarAtcId Simvar
// args contain optional index
func SimVarAtcId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC ID",
		Units:    "String64",
		Settable: true,
	}
}

// SimVarAtcAirline Simvar
// args contain optional index
func SimVarAtcAirline(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC AIRLINE",
		Units:    "String64",
		Settable: true,
	}
}

// SimVarAtcFlightNumber Simvar
// args contain optional index
func SimVarAtcFlightNumber(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ATC FLIGHT NUMBER",
		Units:    "String8",
		Settable: true,
	}
}

// SimVarTitle Actually not supported and crash FS2020
// args contain optional index
/*func SimVarTitle(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TITLE",
		Units:    "Variable length string",
		Settable: false,
	}
}*/

// SimVarHsiStationIdent Simvar
// args contain optional index
func SimVarHsiStationIdent(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "HSI STATION IDENT",
		Units:    "String8",
		Settable: false,
	}
}

// SimVarGpsApproachAirportId Simvar
// args contain optional index
func SimVarGpsApproachAirportId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH AIRPORT ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsApproachApproachId Simvar
// args contain optional index
func SimVarGpsApproachApproachId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsApproachTransitionId Simvar
// args contain optional index
func SimVarGpsApproachTransitionId(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TRANSITION ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAbsoluteTime Simvar
// args contain optional index
func SimVarAbsoluteTime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ABSOLUTE TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarZuluTime Simvar
// args contain optional index
func SimVarZuluTime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarZuluDayOfWeek Simvar
// args contain optional index
func SimVarZuluDayOfWeek(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF WEEK",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluDayOfMonth Simvar
// args contain optional index
func SimVarZuluDayOfMonth(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF MONTH",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluMonthOfYear Simvar
// args contain optional index
func SimVarZuluMonthOfYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU MONTH OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluDayOfYear Simvar
// args contain optional index
func SimVarZuluDayOfYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluYear Simvar
// args contain optional index
func SimVarZuluYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "ZULU YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalTime Simvar
// args contain optional index
func SimVarLocalTime(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarLocalDayOfWeek Simvar
// args contain optional index
func SimVarLocalDayOfWeek(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF WEEK",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalDayOfMonth Simvar
// args contain optional index
func SimVarLocalDayOfMonth(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF MONTH",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalMonthOfYear Simvar
// args contain optional index
func SimVarLocalMonthOfYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL MONTH OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalDayOfYear Simvar
// args contain optional index
func SimVarLocalDayOfYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalYear Simvar
// args contain optional index
func SimVarLocalYear(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "LOCAL YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTimeZoneOffset Simvar
// args contain optional index
func SimVarTimeZoneOffset(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TIME ZONE OFFSET",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarTimeOfDay Simvar
// args contain optional index
func SimVarTimeOfDay(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "TIME OF DAY",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarSimulationRate Simvar
// args contain optional index
func SimVarSimulationRate(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "SIMULATION RATE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUnitsOfMeasure Simvar
// args contain optional index
func SimVarUnitsOfMeasure(args ...int) SimVar {
	index := 0
	if len(args) > 0 {
		index = args[0]
	}
	return SimVar{
		Index:    index,
		Name:     "UNITS OF MEASURE",
		Units:    "Enum",
		Settable: false,
	}
}
