package simconnect

// Dcumentation on http://www.prepar3d.com/SDKv3/LearningCenter/utilities/variables/simulation_variables.html
import (
	"bytes"
	"encoding/binary"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

type SimVarUnit string

const (
	UnitBool                    SimVarUnit = "Bool"
	UnitFeetpersecond           SimVarUnit = "Feetpersecond"
	UnitPercentover100          SimVarUnit = "Percentover100"
	UnitNumber                  SimVarUnit = "Number"
	UnitGallons                 SimVarUnit = "Gallons"
	UnitString                  SimVarUnit = "String"
	UnitBoolString              SimVarUnit = "Bool/String"
	UnitFeet                    SimVarUnit = "Feet"
	UnitSimconnectDataXyz       SimVarUnit = "SimconnectDataXyz"
	UnitMask                    SimVarUnit = "Mask"
	UnitKnots                   SimVarUnit = "Knots"
	UnitSimconnectDataWaypoint  SimVarUnit = "SimconnectDataWaypoint"
	UnitDegrees                 SimVarUnit = "Degrees"
	UnitSeconds                 SimVarUnit = "Seconds"
	UnitBoolean                 SimVarUnit = "Boolean"
	UnitSimconnectDataLatlonalt SimVarUnit = "SimconnectDataLatlonalt"
	UnitPercent                 SimVarUnit = "Percent"
	UnitEnum                    SimVarUnit = "Enum"
	UnitRpm                     SimVarUnit = "Rpm"
	UnitRankine                 SimVarUnit = "Rankine"
	UnitPsi                     SimVarUnit = "Psi"
	UnitHours                   SimVarUnit = "Hours"
	UnitPosition                SimVarUnit = "Position"
	Unitftlbpersecond           SimVarUnit = "ftlbpersecond"
	UnitFootpound               SimVarUnit = "Footpound"
	UnitCelsius                 SimVarUnit = "Celsius"
	UnitPoundsperhour           SimVarUnit = "Poundsperhour"
	UnitRatio                   SimVarUnit = "Ratio"
	UnitPounds                  SimVarUnit = "Pounds"
	UnitRadians                 SimVarUnit = "Radians"
	UnitFootpounds              SimVarUnit = "Footpounds"
	UnitpoundForcepersquareinch SimVarUnit = "pound-forcepersquareinch"
	UnitinHg                    SimVarUnit = "inHg"
	UnitPSI                     SimVarUnit = "PSI"
	UnitFeetpersecondsquared    SimVarUnit = "Feetpersecondsquared"
	UnitMeters                  SimVarUnit = "Meters"
	UnitMach                    SimVarUnit = "Mach"
	UnitMillibars               SimVarUnit = "Millibars"
	UnitRadianspersecond        SimVarUnit = "Radianspersecond"
	UnitGforce                  SimVarUnit = "Gforce"
	UnitFrequencyBCD16          SimVarUnit = "FrequencyBCD16"
	UnitMHz                     SimVarUnit = "MHz"
	UnitNauticalmiles           SimVarUnit = "Nauticalmiles"
	UnitFrequencyADFBCD32       SimVarUnit = "FrequencyADFBCD32"
	UnitHz                      SimVarUnit = "Hz"
	UnitBCO16                   SimVarUnit = "BCO16"
	UnitMeterspersecond         SimVarUnit = "Meterspersecond"
	UnitFlags                   SimVarUnit = "Flags"
	Unitpsf                     SimVarUnit = "psf"
	UnitPercentage              SimVarUnit = "Percentage"
	UnitFeetPMinute             SimVarUnit = "Feet/minute"
	UnitSlugspercubicfeet       SimVarUnit = "Slugspercubicfeet"
	UnitAmperes                 SimVarUnit = "Amperes"
	UnitVolts                   SimVarUnit = "Volts"
	UnitPoundforcepersquarefoot SimVarUnit = "Poundforcepersquarefoot"
	UnitGForce                  SimVarUnit = "GForce"
	UnitFeetperminute           SimVarUnit = "Feetperminute"
	UnitPoundspersquarefoot     SimVarUnit = "Poundspersquarefoot"
	Unitfootpounds              SimVarUnit = "footpounds"
	UnitSquarefeet              SimVarUnit = "Squarefeet"
	UnitPerradian               SimVarUnit = "Perradian"
	UnitMachs                   SimVarUnit = "Machs"
	Unitslugfeetsquared         SimVarUnit = "slugfeetsquared"
	UnitAmps                    SimVarUnit = "Amps"
	UnitPersecond               SimVarUnit = "Persecond"
	UnitString64                SimVarUnit = "String64"
	UnitString8                 SimVarUnit = "String8"
	UnitVariablelengthstring    SimVarUnit = "Variablelengthstring"
)

func readArgs(args []interface{}, defIndex int, defUnit SimVarUnit) (int, SimVarUnit) {
	for _, arg := range args {
		if s, ok := arg.(SimVarUnit); ok {
			defUnit = s
		}
		if i, ok := arg.(int); ok {
			defIndex = i
		}
	}
	return defIndex, defUnit
}

// SimVar is usued for all SimVar describtion
type SimVar struct {
	Name     string
	Unit     SimVarUnit
	Settable bool
	Index    int
	data     []byte
}

func (s *SimVar) getUnitForDataDefinition() string {
	if strings.Contains(string(s.Unit), "String") ||
		strings.Contains(string(s.Unit), "string") ||
		s.Unit == "SIMCONNECT_DATA_LATLONALT" ||
		s.Unit == "SIMCONNECT_DATA_XYZ" ||
		s.Unit == "SIMCONNECT_DATA_WAYPOINT" {
		return ""
	}
	return string(s.Unit)
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
	switch s.Unit {
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
	return 8
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
// args contain optional index and/or unit
func SimVarAutopilotPitchHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT PITCH HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructAmbientWind Simvar
// args contain optional index and/or unit
func SimVarStructAmbientWind(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "STRUCT AMBIENT WIND",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLaunchbarPosition Simvar
// args contain optional index and/or unit
func SimVarLaunchbarPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "LAUNCHBAR POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNumberOfCatapults Simvar
// args contain optional index and/or unit
func SimVarNumberOfCatapults(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NUMBER OF CATAPULTS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHoldbackBarInstalled Simvar
// args contain optional index and/or unit
func SimVarHoldbackBarInstalled(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HOLDBACK BAR INSTALLED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBlastShieldPosition Simvar
// args contain optional index and/or unit
func SimVarBlastShieldPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "BLAST SHIELD POSITION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngDetonating Simvar
// args contain optional index and/or unit
func SimVarRecipEngDetonating(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG DETONATING:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngCylinderHealth Simvar
// args contain optional index and/or unit
func SimVarRecipEngCylinderHealth(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG CYLINDER HEALTH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngNumCylinders Simvar
// args contain optional index and/or unit
func SimVarRecipEngNumCylinders(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NUM CYLINDERS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngNumCylindersFailed Simvar
// args contain optional index and/or unit
func SimVarRecipEngNumCylindersFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NUM CYLINDERS FAILED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngAntidetonationTankValve Simvar
// args contain optional index and/or unit
func SimVarRecipEngAntidetonationTankValve(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK VALVE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankQuantity Simvar
// args contain optional index and/or unit
func SimVarRecipEngAntidetonationTankQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK QUANTITY:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankMaxQuantity Simvar
// args contain optional index and/or unit
func SimVarRecipEngAntidetonationTankMaxQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ANTIDETONATION TANK MAX QUANTITY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngNitrousTankValve Simvar
// args contain optional index and/or unit
func SimVarRecipEngNitrousTankValve(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK VALVE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankQuantity Simvar
// args contain optional index and/or unit
func SimVarRecipEngNitrousTankQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK QUANTITY:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankMaxQuantity Simvar
// args contain optional index and/or unit
func SimVarRecipEngNitrousTankMaxQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG NITROUS TANK MAX QUANTITY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPayloadStationObject Simvar
// args contain optional index and/or unit
func SimVarPayloadStationObject(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION OBJECT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPayloadStationNumSimobjects Simvar
// args contain optional index and/or unit
func SimVarPayloadStationNumSimobjects(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION NUM SIMOBJECTS:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSlingObjectAttached Simvar
// args contain optional index and/or unit
func SimVarSlingObjectAttached(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool/String")
	return SimVar{
		Index:    index,
		Name:     "SLING OBJECT ATTACHED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSlingCableBroken Simvar
// args contain optional index and/or unit
func SimVarSlingCableBroken(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SLING CABLE BROKEN:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSlingCableExtendedLength Simvar
// args contain optional index and/or unit
func SimVarSlingCableExtendedLength(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "SLING CABLE EXTENDED LENGTH:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarSlingActivePayloadStation Simvar
// args contain optional index and/or unit
func SimVarSlingActivePayloadStation(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "SLING ACTIVE PAYLOAD STATION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarSlingHoistPercentDeployed Simvar
// args contain optional index and/or unit
func SimVarSlingHoistPercentDeployed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "SLING HOIST PERCENT DEPLOYED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSlingHookInPickupMode Simvar
// args contain optional index and/or unit
func SimVarSlingHookInPickupMode(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SLING HOOK IN PICKUP MODE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsAttachedToSling Simvar
// args contain optional index and/or unit
func SimVarIsAttachedToSling(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS ATTACHED TO SLING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAlternateStaticSourceOpen Simvar
// args contain optional index and/or unit
func SimVarAlternateStaticSourceOpen(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ALTERNATE STATIC SOURCE OPEN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronTrimPct Simvar
// args contain optional index and/or unit
func SimVarAileronTrimPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "AILERON TRIM PCT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRudderTrimPct Simvar
// args contain optional index and/or unit
func SimVarRudderTrimPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "RUDDER TRIM PCT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarLightOnStates Simvar
// args contain optional index and/or unit
func SimVarLightOnStates(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "LIGHT ON STATES",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightStates Simvar
// args contain optional index and/or unit
func SimVarLightStates(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "LIGHT STATES",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLandingLightPbh Simvar
// args contain optional index and/or unit
func SimVarLandingLightPbh(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "LANDING LIGHT PBH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightTaxiOn Simvar
// args contain optional index and/or unit
func SimVarLightTaxiOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT TAXI ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightStrobeOn Simvar
// args contain optional index and/or unit
func SimVarLightStrobeOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT STROBE ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightPanelOn Simvar
// args contain optional index and/or unit
func SimVarLightPanelOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT PANEL ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightRecognitionOn Simvar
// args contain optional index and/or unit
func SimVarLightRecognitionOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT RECOGNITION ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightWingOn Simvar
// args contain optional index and/or unit
func SimVarLightWingOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT WING ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightLogoOn Simvar
// args contain optional index and/or unit
func SimVarLightLogoOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT LOGO ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightCabinOn Simvar
// args contain optional index and/or unit
func SimVarLightCabinOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT CABIN ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightHeadOn Simvar
// args contain optional index and/or unit
func SimVarLightHeadOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT HEAD ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightBrakeOn Simvar
// args contain optional index and/or unit
func SimVarLightBrakeOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT BRAKE ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightNavOn Simvar
// args contain optional index and/or unit
func SimVarLightNavOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT NAV ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightBeaconOn Simvar
// args contain optional index and/or unit
func SimVarLightBeaconOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT BEACON ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightLandingOn Simvar
// args contain optional index and/or unit
func SimVarLightLandingOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT LANDING ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiDesiredSpeed Simvar
// args contain optional index and/or unit
func SimVarAiDesiredSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AI DESIRED SPEED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiWaypointList Actually not supported
// args contain optional index and/or unit
func SimVarAiWaypointList(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_WAYPOINT")
	return SimVar{
		Index:    index,
		Name:     "AI WAYPOINT LIST",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiCurrentWaypoint Simvar
// args contain optional index and/or unit
func SimVarAiCurrentWaypoint(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "AI CURRENT WAYPOINT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiDesiredHeading Simvar
// args contain optional index and/or unit
func SimVarAiDesiredHeading(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "AI DESIRED HEADING",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiGroundturntime Simvar
// args contain optional index and/or unit
func SimVarAiGroundturntime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDTURNTIME",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiGroundcruisespeed Simvar
// args contain optional index and/or unit
func SimVarAiGroundcruisespeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDCRUISESPEED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiGroundturnspeed Simvar
// args contain optional index and/or unit
func SimVarAiGroundturnspeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AI GROUNDTURNSPEED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAiTrafficIsifr Simvar
// args contain optional index and/or unit
func SimVarAiTrafficIsifr(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Boolean")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ISIFR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficState Simvar
// args contain optional index and/or unit
func SimVarAiTrafficState(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC STATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficCurrentAirport Simvar
// args contain optional index and/or unit
func SimVarAiTrafficCurrentAirport(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC CURRENT AIRPORT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficAssignedRunway Simvar
// args contain optional index and/or unit
func SimVarAiTrafficAssignedRunway(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ASSIGNED RUNWAY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficAssignedParking Simvar
// args contain optional index and/or unit
func SimVarAiTrafficAssignedParking(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ASSIGNED PARKING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficFromairport Simvar
// args contain optional index and/or unit
func SimVarAiTrafficFromairport(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC FROMAIRPORT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficToairport Simvar
// args contain optional index and/or unit
func SimVarAiTrafficToairport(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC TOAIRPORT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficEtd Simvar
// args contain optional index and/or unit
func SimVarAiTrafficEtd(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ETD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAiTrafficEta Simvar
// args contain optional index and/or unit
func SimVarAiTrafficEta(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "AI TRAFFIC ETA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDroppableObjectsType Simvar
// args contain optional index and/or unit
func SimVarDroppableObjectsType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS TYPE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarDroppableObjectsCount Simvar
// args contain optional index and/or unit
func SimVarDroppableObjectsCount(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS COUNT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWingFlexPct Simvar
// args contain optional index and/or unit
func SimVarWingFlexPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WING FLEX PCT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarApplyHeatToSystems Simvar
// args contain optional index and/or unit
func SimVarApplyHeatToSystems(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "APPLY HEAT TO SYSTEMS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAdfLatlonalt Simvar
// args contain optional index and/or unit
func SimVarAdfLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "ADF LATLONALT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavVorLatlonalt Simvar
// args contain optional index and/or unit
func SimVarNavVorLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "NAV VOR LATLONALT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGsLatlonalt Simvar
// args contain optional index and/or unit
func SimVarNavGsLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "NAV GS LATLONALT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavDmeLatlonalt Simvar
// args contain optional index and/or unit
func SimVarNavDmeLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "NAV DME LATLONALT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarInnerMarkerLatlonalt Simvar
// args contain optional index and/or unit
func SimVarInnerMarkerLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "INNER MARKER LATLONALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMiddleMarkerLatlonalt Simvar
// args contain optional index and/or unit
func SimVarMiddleMarkerLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "MIDDLE MARKER LATLONALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarOuterMarkerLatlonalt Simvar
// args contain optional index and/or unit
func SimVarOuterMarkerLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "OUTER MARKER LATLONALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructLatlonalt Simvar
// args contain optional index and/or unit
func SimVarStructLatlonalt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "STRUCT LATLONALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructLatlonaltpbh Simvar
// args contain optional index and/or unit
func SimVarStructLatlonaltpbh(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "STRUCT LATLONALTPBH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructSurfaceRelativeVelocity Simvar
// args contain optional index and/or unit
func SimVarStructSurfaceRelativeVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT SURFACE RELATIVE VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructWorldvelocity Simvar
// args contain optional index and/or unit
func SimVarStructWorldvelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLDVELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructWorldRotationVelocity Simvar
// args contain optional index and/or unit
func SimVarStructWorldRotationVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLD ROTATION VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructBodyVelocity Simvar
// args contain optional index and/or unit
func SimVarStructBodyVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT BODY VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructBodyRotationVelocity Simvar
// args contain optional index and/or unit
func SimVarStructBodyRotationVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT BODY ROTATION VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructWorldAcceleration Simvar
// args contain optional index and/or unit
func SimVarStructWorldAcceleration(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT WORLD ACCELERATION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructEnginePosition Simvar
// args contain optional index and/or unit
func SimVarStructEnginePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT ENGINE POSITION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructEyepointDynamicAngle Simvar
// args contain optional index and/or unit
func SimVarStructEyepointDynamicAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT EYEPOINT DYNAMIC ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructEyepointDynamicOffset Simvar
// args contain optional index and/or unit
func SimVarStructEyepointDynamicOffset(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "STRUCT EYEPOINT DYNAMIC OFFSET",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEyepointPosition Simvar
// args contain optional index and/or unit
func SimVarEyepointPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_XYZ")
	return SimVar{
		Index:    index,
		Name:     "EYEPOINT POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireElacSwitch Simvar
// args contain optional index and/or unit
func SimVarFlyByWireElacSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE ELAC SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireFacSwitch Simvar
// args contain optional index and/or unit
func SimVarFlyByWireFacSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE FAC SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireSecSwitch Simvar
// args contain optional index and/or unit
func SimVarFlyByWireSecSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE SEC SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireElacFailed Simvar
// args contain optional index and/or unit
func SimVarFlyByWireElacFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE ELAC FAILED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireFacFailed Simvar
// args contain optional index and/or unit
func SimVarFlyByWireFacFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE FAC FAILED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlyByWireSecFailed Simvar
// args contain optional index and/or unit
func SimVarFlyByWireSecFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLY BY WIRE SEC FAILED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNumberOfEngines Simvar
// args contain optional index and/or unit
func SimVarNumberOfEngines(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NUMBER OF ENGINES",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngineControlSelect Simvar
// args contain optional index and/or unit
func SimVarEngineControlSelect(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "ENGINE CONTROL SELECT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarThrottleLowerLimit Simvar
// args contain optional index and/or unit
func SimVarThrottleLowerLimit(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "THROTTLE LOWER LIMIT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngineType Simvar
// args contain optional index and/or unit
func SimVarEngineType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "ENGINE TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMasterIgnitionSwitch Simvar
// args contain optional index and/or unit
func SimVarMasterIgnitionSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "MASTER IGNITION SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngCombustion Simvar
// args contain optional index and/or unit
func SimVarGeneralEngCombustion(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG COMBUSTION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngMasterAlternator Simvar
// args contain optional index and/or unit
func SimVarGeneralEngMasterAlternator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MASTER ALTERNATOR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpSwitch Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFuelPumpSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PUMP SWITCH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpOn Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFuelPumpOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PUMP ON:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngRpm Simvar
// args contain optional index and/or unit
func SimVarGeneralEngRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngPctMaxRpm Simvar
// args contain optional index and/or unit
func SimVarGeneralEngPctMaxRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG PCT MAX RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngMaxReachedRpm Simvar
// args contain optional index and/or unit
func SimVarGeneralEngMaxReachedRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MAX REACHED RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngThrottleLeverPosition Simvar
// args contain optional index and/or unit
func SimVarGeneralEngThrottleLeverPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG THROTTLE LEVER POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngMixtureLeverPosition Simvar
// args contain optional index and/or unit
func SimVarGeneralEngMixtureLeverPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG MIXTURE LEVER POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngPropellerLeverPosition Simvar
// args contain optional index and/or unit
func SimVarGeneralEngPropellerLeverPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG PROPELLER LEVER POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngStarter Simvar
// args contain optional index and/or unit
func SimVarGeneralEngStarter(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG STARTER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngExhaustGasTemperature Simvar
// args contain optional index and/or unit
func SimVarGeneralEngExhaustGasTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG EXHAUST GAS TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngOilPressure Simvar
// args contain optional index and/or unit
func SimVarGeneralEngOilPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Psi")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL PRESSURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngOilLeakedPercent Simvar
// args contain optional index and/or unit
func SimVarGeneralEngOilLeakedPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL LEAKED PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngCombustionSoundPercent Simvar
// args contain optional index and/or unit
func SimVarGeneralEngCombustionSoundPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG COMBUSTION SOUND PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngDamagePercent Simvar
// args contain optional index and/or unit
func SimVarGeneralEngDamagePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG DAMAGE PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngOilTemperature Simvar
// args contain optional index and/or unit
func SimVarGeneralEngOilTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG OIL TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngFailed Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FAILED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorSwitch Simvar
// args contain optional index and/or unit
func SimVarGeneralEngGeneratorSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG GENERATOR SWITCH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorActive Simvar
// args contain optional index and/or unit
func SimVarGeneralEngGeneratorActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG GENERATOR ACTIVE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngAntiIcePosition Simvar
// args contain optional index and/or unit
func SimVarGeneralEngAntiIcePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG ANTI ICE POSITION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngFuelValve Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFuelValve(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL VALVE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngFuelPressure Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFuelPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Psi")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL PRESSURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGeneralEngElapsedTime Simvar
// args contain optional index and/or unit
func SimVarGeneralEngElapsedTime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Hours")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG ELAPSED TIME:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngCowlFlapPosition Simvar
// args contain optional index and/or unit
func SimVarRecipEngCowlFlapPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG COWL FLAP POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngPrimer Simvar
// args contain optional index and/or unit
func SimVarRecipEngPrimer(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG PRIMER:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngManifoldPressure Simvar
// args contain optional index and/or unit
func SimVarRecipEngManifoldPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Psi")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG MANIFOLD PRESSURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngAlternateAirPosition Simvar
// args contain optional index and/or unit
func SimVarRecipEngAlternateAirPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG ALTERNATE AIR POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngCoolantReservoirPercent Simvar
// args contain optional index and/or unit
func SimVarRecipEngCoolantReservoirPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG COOLANT RESERVOIR PERCENT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngLeftMagneto Simvar
// args contain optional index and/or unit
func SimVarRecipEngLeftMagneto(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG LEFT MAGNETO:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngRightMagneto Simvar
// args contain optional index and/or unit
func SimVarRecipEngRightMagneto(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG RIGHT MAGNETO:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngBrakePower Simvar
// args contain optional index and/or unit
func SimVarRecipEngBrakePower(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "ft lb per second")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG BRAKE POWER:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngStarterTorque Simvar
// args contain optional index and/or unit
func SimVarRecipEngStarterTorque(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Foot pound")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG STARTER TORQUE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngTurbochargerFailed Simvar
// args contain optional index and/or unit
func SimVarRecipEngTurbochargerFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG TURBOCHARGER FAILED:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostActive Simvar
// args contain optional index and/or unit
func SimVarRecipEngEmergencyBoostActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG EMERGENCY BOOST ACTIVE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostElapsedTime Simvar
// args contain optional index and/or unit
func SimVarRecipEngEmergencyBoostElapsedTime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Hours")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG EMERGENCY BOOST ELAPSED TIME:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngWastegatePosition Simvar
// args contain optional index and/or unit
func SimVarRecipEngWastegatePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG WASTEGATE POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngTurbineInletTemperature Simvar
// args contain optional index and/or unit
func SimVarRecipEngTurbineInletTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG TURBINE INLET TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngCylinderHeadTemperature Simvar
// args contain optional index and/or unit
func SimVarRecipEngCylinderHeadTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG CYLINDER HEAD TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngRadiatorTemperature Simvar
// args contain optional index and/or unit
func SimVarRecipEngRadiatorTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG RADIATOR TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngFuelAvailable Simvar
// args contain optional index and/or unit
func SimVarRecipEngFuelAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL AVAILABLE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngFuelFlow Simvar
// args contain optional index and/or unit
func SimVarRecipEngFuelFlow(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL FLOW:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngFuelTankSelector Simvar
// args contain optional index and/or unit
func SimVarRecipEngFuelTankSelector(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL TANK SELECTOR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipEngFuelTanksUsed Simvar
// args contain optional index and/or unit
func SimVarRecipEngFuelTanksUsed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL TANKS USED:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipEngFuelNumberTanksUsed Simvar
// args contain optional index and/or unit
func SimVarRecipEngFuelNumberTanksUsed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "RECIP ENG FUEL NUMBER TANKS USED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRecipCarburetorTemperature Simvar
// args contain optional index and/or unit
func SimVarRecipCarburetorTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "RECIP CARBURETOR TEMPERATURE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRecipMixtureRatio Simvar
// args contain optional index and/or unit
func SimVarRecipMixtureRatio(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Ratio")
	return SimVar{
		Index:    index,
		Name:     "RECIP MIXTURE RATIO:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngN1 Simvar
func SimVarTurbEngN1(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG N1:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngN2 Simvar
func SimVarTurbEngN2(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG N2:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN1 Simvar
func SimVarTurbEngCorrectedN1(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED N1:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN2 Simvar
func SimVarTurbEngCorrectedN2(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED N2:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngCorrectedFf Simvar
// args contain optional index and/or unit
func SimVarTurbEngCorrectedFf(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG CORRECTED FF:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngMaxTorquePercent Simvar
// args contain optional index and/or unit
func SimVarTurbEngMaxTorquePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG MAX TORQUE PERCENT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngPressureRatio Simvar
// args contain optional index and/or unit
func SimVarTurbEngPressureRatio(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Ratio")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG PRESSURE RATIO:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngItt Simvar
// args contain optional index and/or unit
func SimVarTurbEngItt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG ITT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurbEngAfterburner Simvar
// args contain optional index and/or unit
func SimVarTurbEngAfterburner(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG AFTERBURNER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngJetThrust Simvar
// args contain optional index and/or unit
func SimVarTurbEngJetThrust(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG JET THRUST:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngBleedAir Simvar
// args contain optional index and/or unit
func SimVarTurbEngBleedAir(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Psi")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG BLEED AIR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngTankSelector Simvar
// args contain optional index and/or unit
func SimVarTurbEngTankSelector(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG TANK SELECTOR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngTanksUsed Simvar
// args contain optional index and/or unit
func SimVarTurbEngTanksUsed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG TANKS USED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngNumTanksUsed Simvar
// args contain optional index and/or unit
func SimVarTurbEngNumTanksUsed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG NUM TANKS USED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngFuelFlowPph Simvar
// args contain optional index and/or unit
func SimVarTurbEngFuelFlowPph(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG FUEL FLOW PPH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngFuelAvailable Simvar
// args contain optional index and/or unit
func SimVarTurbEngFuelAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG FUEL AVAILABLE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngReverseNozzlePercent Simvar
// args contain optional index and/or unit
func SimVarTurbEngReverseNozzlePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG REVERSE NOZZLE PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngVibration Simvar
// args contain optional index and/or unit
func SimVarTurbEngVibration(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG VIBRATION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngFailed Simvar
// args contain optional index and/or unit
func SimVarEngFailed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ENG FAILED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngRpmAnimationPercent Simvar
// args contain optional index and/or unit
func SimVarEngRpmAnimationPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "ENG RPM ANIMATION PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngOnFire Simvar
// args contain optional index and/or unit
func SimVarEngOnFire(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ENG ON FIRE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarEngFuelFlowBugPosition Simvar
// args contain optional index and/or unit
func SimVarEngFuelFlowBugPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL FLOW BUG POSITION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropRpm Simvar
// args contain optional index and/or unit
func SimVarPropRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "PROP RPM:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPropMaxRpmPercent Simvar
// args contain optional index and/or unit
func SimVarPropMaxRpmPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "PROP MAX RPM PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropThrust Simvar
// args contain optional index and/or unit
func SimVarPropThrust(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "PROP THRUST:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropBeta Simvar
// args contain optional index and/or unit
func SimVarPropBeta(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PROP BETA:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropFeatheringInhibit Simvar
// args contain optional index and/or unit
func SimVarPropFeatheringInhibit(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHERING INHIBIT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropFeathered Simvar
// args contain optional index and/or unit
func SimVarPropFeathered(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHERED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropSyncDeltaLever Simvar
// args contain optional index and/or unit
func SimVarPropSyncDeltaLever(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "PROP SYNC DELTA LEVER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropAutoFeatherArmed Simvar
// args contain optional index and/or unit
func SimVarPropAutoFeatherArmed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP AUTO FEATHER ARMED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropFeatherSwitch Simvar
// args contain optional index and/or unit
func SimVarPropFeatherSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP FEATHER SWITCH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPanelAutoFeatherSwitch Simvar
// args contain optional index and/or unit
func SimVarPanelAutoFeatherSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PANEL AUTO FEATHER SWITCH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropSyncActive Simvar
// args contain optional index and/or unit
func SimVarPropSyncActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP SYNC ACTIVE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropDeiceSwitch Simvar
// args contain optional index and/or unit
func SimVarPropDeiceSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP DEICE SWITCH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngCombustion Simvar
// args contain optional index and/or unit
func SimVarEngCombustion(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ENG COMBUSTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngN1Rpm Simvar
func SimVarEngN1Rpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "ENG N1 RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngN2Rpm Simvar
func SimVarEngN2Rpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "ENG N2 RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngFuelFlowPph Simvar
// args contain optional index and/or unit
func SimVarEngFuelFlowPph(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL FLOW PPH:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngTorque Simvar
// args contain optional index and/or unit
func SimVarEngTorque(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Foot pounds")
	return SimVar{
		Index:    index,
		Name:     "ENG TORQUE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngAntiIce Simvar
// args contain optional index and/or unit
func SimVarEngAntiIce(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ENG ANTI ICE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngPressureRatio Simvar
// args contain optional index and/or unit
func SimVarEngPressureRatio(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Ratio (0-16384)")
	return SimVar{
		Index:    index,
		Name:     "ENG PRESSURE RATIO:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperature Simvar
// args contain optional index and/or unit
func SimVarEngExhaustGasTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "ENG EXHAUST GAS TEMPERATURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperatureGes Simvar
// args contain optional index and/or unit
func SimVarEngExhaustGasTemperatureGes(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ENG EXHAUST GAS TEMPERATURE GES:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngCylinderHeadTemperature Simvar
// args contain optional index and/or unit
func SimVarEngCylinderHeadTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "ENG CYLINDER HEAD TEMPERATURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngOilTemperature Simvar
// args contain optional index and/or unit
func SimVarEngOilTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "ENG OIL TEMPERATURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngOilPressure Simvar
// args contain optional index and/or unit
func SimVarEngOilPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "pound-force per square inch")
	return SimVar{
		Index:    index,
		Name:     "ENG OIL PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngOilQuantity Simvar
// args contain optional index and/or unit
func SimVarEngOilQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ENG OIL QUANTITY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngHydraulicPressure Simvar
// args contain optional index and/or unit
func SimVarEngHydraulicPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "pound-force per square inch")
	return SimVar{
		Index:    index,
		Name:     "ENG HYDRAULIC PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngHydraulicQuantity Simvar
// args contain optional index and/or unit
func SimVarEngHydraulicQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ENG HYDRAULIC QUANTITY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngManifoldPressure Simvar
// args contain optional index and/or unit
func SimVarEngManifoldPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "inHg")
	return SimVar{
		Index:    index,
		Name:     "ENG MANIFOLD PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngVibration Simvar
// args contain optional index and/or unit
func SimVarEngVibration(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ENG VIBRATION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngRpmScaler Simvar
// args contain optional index and/or unit
func SimVarEngRpmScaler(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ENG RPM SCALER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngTurbineTemperature Simvar
// args contain optional index and/or unit
func SimVarEngTurbineTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "ENG TURBINE TEMPERATURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngTorquePercent Simvar
// args contain optional index and/or unit
func SimVarEngTorquePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "ENG TORQUE PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngFuelPressure Simvar
// args contain optional index and/or unit
func SimVarEngFuelPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "PSI")
	return SimVar{
		Index:    index,
		Name:     "ENG FUEL PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngElectricalLoad Simvar
// args contain optional index and/or unit
func SimVarEngElectricalLoad(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "ENG ELECTRICAL LOAD:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngTransmissionPressure Simvar
// args contain optional index and/or unit
func SimVarEngTransmissionPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "PSI")
	return SimVar{
		Index:    index,
		Name:     "ENG TRANSMISSION PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngTransmissionTemperature Simvar
// args contain optional index and/or unit
func SimVarEngTransmissionTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "ENG TRANSMISSION TEMPERATURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngRotorRpm Simvar
// args contain optional index and/or unit
func SimVarEngRotorRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "ENG ROTOR RPM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngMaxRpm Simvar
// args contain optional index and/or unit
func SimVarEngMaxRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "ENG MAX RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngStarterActive Simvar
// args contain optional index and/or unit
func SimVarGeneralEngStarterActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG STARTER ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGeneralEngFuelUsedSinceStart Simvar
// args contain optional index and/or unit
func SimVarGeneralEngFuelUsedSinceStart(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "GENERAL ENG FUEL USED SINCE START",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngPrimaryNozzlePercent Simvar
// args contain optional index and/or unit
func SimVarTurbEngPrimaryNozzlePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG PRIMARY NOZZLE PERCENT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngIgnitionSwitch Simvar
// args contain optional index and/or unit
func SimVarTurbEngIgnitionSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG IGNITION SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurbEngMasterStarterSwitch Simvar
// args contain optional index and/or unit
func SimVarTurbEngMasterStarterSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TURB ENG MASTER STARTER SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankCenterLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankCenterLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankCenter2Level Simvar
func SimVarFuelTankCenter2Level(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankCenter3Level Simvar
func SimVarFuelTankCenter3Level(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftMainLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftMainLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftAuxLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftTipLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftTipLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightMainLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightMainLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightAuxLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightAuxLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightTipLevel Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightTipLevel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankExternal1Level Simvar
func SimVarFuelTankExternal1Level(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankExternal2Level Simvar
func SimVarFuelTankExternal2Level(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 LEVEL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankCenterCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankCenterCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankCenter2Capacity Simvar
func SimVarFuelTankCenter2Capacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankCenter3Capacity Simvar
func SimVarFuelTankCenter3Capacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankLeftMainCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftMainCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankLeftAuxCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftAuxCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankLeftTipCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftTipCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankRightMainCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightMainCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankRightAuxCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightAuxCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankRightTipCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightTipCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankExternal1Capacity Simvar
func SimVarFuelTankExternal1Capacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankExternal2Capacity Simvar
func SimVarFuelTankExternal2Capacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelLeftCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelLeftCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL LEFT CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelRightCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelRightCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL RIGHT CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankCenterQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankCenterQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankCenter2Quantity Simvar
func SimVarFuelTankCenter2Quantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER2 QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankCenter3Quantity Simvar
func SimVarFuelTankCenter3Quantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK CENTER3 QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftMainQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftMainQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT MAIN QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftAuxQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT AUX QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankLeftTipQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankLeftTipQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK LEFT TIP QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightMainQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightMainQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT MAIN QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightAuxQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightAuxQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT AUX QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankRightTipQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTankRightTipQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK RIGHT TIP QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankExternal1Quantity Simvar
func SimVarFuelTankExternal1Quantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL1 QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelTankExternal2Quantity Simvar
func SimVarFuelTankExternal2Quantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK EXTERNAL2 QUANTITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFuelLeftQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelLeftQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL LEFT QUANTITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelRightQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelRightQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL RIGHT QUANTITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTotalQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelTotalQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL QUANTITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelWeightPerGallon Simvar
// args contain optional index and/or unit
func SimVarFuelWeightPerGallon(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "FUEL WEIGHT PER GALLON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTankSelector Simvar
// args contain optional index and/or unit
func SimVarFuelTankSelector(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "FUEL TANK SELECTOR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelCrossFeed Simvar
// args contain optional index and/or unit
func SimVarFuelCrossFeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "FUEL CROSS FEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTotalCapacity Simvar
// args contain optional index and/or unit
func SimVarFuelTotalCapacity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL CAPACITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelSelectedQuantityPercent Simvar
// args contain optional index and/or unit
func SimVarFuelSelectedQuantityPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED QUANTITY PERCENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelSelectedQuantity Simvar
// args contain optional index and/or unit
func SimVarFuelSelectedQuantity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gallons")
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED QUANTITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelTotalQuantityWeight Simvar
// args contain optional index and/or unit
func SimVarFuelTotalQuantityWeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "FUEL TOTAL QUANTITY WEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNumFuelSelectors Simvar
// args contain optional index and/or unit
func SimVarNumFuelSelectors(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NUM FUEL SELECTORS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarUnlimitedFuel Simvar
// args contain optional index and/or unit
func SimVarUnlimitedFuel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "UNLIMITED FUEL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEstimatedFuelFlow Simvar
// args contain optional index and/or unit
func SimVarEstimatedFuelFlow(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per hour")
	return SimVar{
		Index:    index,
		Name:     "ESTIMATED FUEL FLOW",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightStrobe Simvar
// args contain optional index and/or unit
func SimVarLightStrobe(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT STROBE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightPanel Simvar
// args contain optional index and/or unit
func SimVarLightPanel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT PANEL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightLanding Simvar
// args contain optional index and/or unit
func SimVarLightLanding(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT LANDING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightTaxi Simvar
// args contain optional index and/or unit
func SimVarLightTaxi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT TAXI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightBeacon Simvar
// args contain optional index and/or unit
func SimVarLightBeacon(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT BEACON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightNav Simvar
// args contain optional index and/or unit
func SimVarLightNav(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT NAV",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightLogo Simvar
// args contain optional index and/or unit
func SimVarLightLogo(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT LOGO",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightWing Simvar
// args contain optional index and/or unit
func SimVarLightWing(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT WING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightRecognition Simvar
// args contain optional index and/or unit
func SimVarLightRecognition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT RECOGNITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLightCabin Simvar
// args contain optional index and/or unit
func SimVarLightCabin(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "LIGHT CABIN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGroundVelocity Simvar
// args contain optional index and/or unit
func SimVarGroundVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "GROUND VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWorldVelocity Simvar
// args contain optional index and/or unit
func SimVarTotalWorldVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WORLD VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarVelocityBodyZ Simvar
// args contain optional index and/or unit
func SimVarVelocityBodyZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY Z",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarVelocityBodyX Simvar
// args contain optional index and/or unit
func SimVarVelocityBodyX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY X",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarVelocityBodyY Simvar
// args contain optional index and/or unit
func SimVarVelocityBodyY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY BODY Y",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarVelocityWorldZ Simvar
// args contain optional index and/or unit
func SimVarVelocityWorldZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD Z",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarVelocityWorldX Simvar
// args contain optional index and/or unit
func SimVarVelocityWorldX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD X",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarVelocityWorldY Simvar
// args contain optional index and/or unit
func SimVarVelocityWorldY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VELOCITY WORLD Y",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationWorldX Simvar
// args contain optional index and/or unit
func SimVarAccelerationWorldX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD X",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationWorldY Simvar
// args contain optional index and/or unit
func SimVarAccelerationWorldY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD Y",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationWorldZ Simvar
// args contain optional index and/or unit
func SimVarAccelerationWorldZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION WORLD Z",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationBodyX Simvar
// args contain optional index and/or unit
func SimVarAccelerationBodyX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY X",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationBodyY Simvar
// args contain optional index and/or unit
func SimVarAccelerationBodyY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY Y",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAccelerationBodyZ Simvar
// args contain optional index and/or unit
func SimVarAccelerationBodyZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second squared")
	return SimVar{
		Index:    index,
		Name:     "ACCELERATION BODY Z",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRotationVelocityBodyX Simvar
// args contain optional index and/or unit
func SimVarRotationVelocityBodyX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY X",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRotationVelocityBodyY Simvar
// args contain optional index and/or unit
func SimVarRotationVelocityBodyY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY Y",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRotationVelocityBodyZ Simvar
// args contain optional index and/or unit
func SimVarRotationVelocityBodyZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "ROTATION VELOCITY BODY Z",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRelativeWindVelocityBodyX Simvar
// args contain optional index and/or unit
func SimVarRelativeWindVelocityBodyX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY X",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyY Simvar
// args contain optional index and/or unit
func SimVarRelativeWindVelocityBodyY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY Y",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyZ Simvar
// args contain optional index and/or unit
func SimVarRelativeWindVelocityBodyZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "RELATIVE WIND VELOCITY BODY Z",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPlaneAltAboveGround Simvar
// args contain optional index and/or unit
func SimVarPlaneAltAboveGround(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PLANE ALT ABOVE GROUND",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneLatitude Simvar
// args contain optional index and/or unit
func SimVarPlaneLatitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE LATITUDE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneLongitude Simvar
// args contain optional index and/or unit
func SimVarPlaneLongitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE LONGITUDE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneAltitude Simvar
// args contain optional index and/or unit
func SimVarPlaneAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PLANE ALTITUDE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlanePitchDegrees Simvar
// args contain optional index and/or unit
func SimVarPlanePitchDegrees(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE PITCH DEGREES",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneBankDegrees Simvar
// args contain optional index and/or unit
func SimVarPlaneBankDegrees(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE BANK DEGREES",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesTrue Simvar
// args contain optional index and/or unit
func SimVarPlaneHeadingDegreesTrue(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES TRUE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesMagnetic Simvar
// args contain optional index and/or unit
func SimVarPlaneHeadingDegreesMagnetic(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES MAGNETIC",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarMagvar Simvar
// args contain optional index and/or unit
func SimVarMagvar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "MAGVAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGroundAltitude Simvar
// args contain optional index and/or unit
func SimVarGroundAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GROUND ALTITUDE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSurfaceType Simvar
// args contain optional index and/or unit
func SimVarSurfaceType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "SURFACE TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSimOnGround Simvar
// args contain optional index and/or unit
func SimVarSimOnGround(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SIM ON GROUND",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIncidenceAlpha Simvar
// args contain optional index and/or unit
func SimVarIncidenceAlpha(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "INCIDENCE ALPHA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIncidenceBeta Simvar
// args contain optional index and/or unit
func SimVarIncidenceBeta(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "INCIDENCE BETA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAirspeedTrue Simvar
// args contain optional index and/or unit
func SimVarAirspeedTrue(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED TRUE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAirspeedIndicated Simvar
// args contain optional index and/or unit
func SimVarAirspeedIndicated(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED INDICATED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAirspeedTrueCalibrate Simvar
// args contain optional index and/or unit
func SimVarAirspeedTrueCalibrate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED TRUE CALIBRATE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAirspeedBarberPole Simvar
// args contain optional index and/or unit
func SimVarAirspeedBarberPole(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED BARBER POLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAirspeedMach Simvar
// args contain optional index and/or unit
func SimVarAirspeedMach(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mach")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED MACH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarVerticalSpeed Simvar
// args contain optional index and/or unit
func SimVarVerticalSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VERTICAL SPEED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarMachMaxOperate Simvar
// args contain optional index and/or unit
func SimVarMachMaxOperate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mach")
	return SimVar{
		Index:    index,
		Name:     "MACH MAX OPERATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStallWarning Simvar
// args contain optional index and/or unit
func SimVarStallWarning(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "STALL WARNING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarOverspeedWarning Simvar
// args contain optional index and/or unit
func SimVarOverspeedWarning(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "OVERSPEED WARNING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBarberPoleMach Simvar
// args contain optional index and/or unit
func SimVarBarberPoleMach(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mach")
	return SimVar{
		Index:    index,
		Name:     "BARBER POLE MACH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIndicatedAltitude Simvar
// args contain optional index and/or unit
func SimVarIndicatedAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "INDICATED ALTITUDE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarKohlsmanSettingMb Simvar
// args contain optional index and/or unit
func SimVarKohlsmanSettingMb(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Millibars")
	return SimVar{
		Index:    index,
		Name:     "KOHLSMAN SETTING MB",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarKohlsmanSettingHg Simvar
// args contain optional index and/or unit
func SimVarKohlsmanSettingHg(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "inHg")
	return SimVar{
		Index:    index,
		Name:     "KOHLSMAN SETTING HG",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAttitudeIndicatorPitchDegrees Simvar
// args contain optional index and/or unit
func SimVarAttitudeIndicatorPitchDegrees(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE INDICATOR PITCH DEGREES",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAttitudeIndicatorBankDegrees Simvar
// args contain optional index and/or unit
func SimVarAttitudeIndicatorBankDegrees(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE INDICATOR BANK DEGREES",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAttitudeBarsPosition Simvar
// args contain optional index and/or unit
func SimVarAttitudeBarsPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE BARS POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAttitudeCage Simvar
// args contain optional index and/or unit
func SimVarAttitudeCage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ATTITUDE CAGE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWiskeyCompassIndicationDegrees Simvar
// args contain optional index and/or unit
func SimVarWiskeyCompassIndicationDegrees(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "WISKEY COMPASS INDICATION DEGREES",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesGyro Simvar
// args contain optional index and/or unit
func SimVarPlaneHeadingDegreesGyro(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PLANE HEADING DEGREES GYRO",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarHeadingIndicator Simvar
// args contain optional index and/or unit
func SimVarHeadingIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "HEADING INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGyroDriftError Simvar
// args contain optional index and/or unit
func SimVarGyroDriftError(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GYRO DRIFT ERROR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDeltaHeadingRate Simvar
// args contain optional index and/or unit
func SimVarDeltaHeadingRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians per second")
	return SimVar{
		Index:    index,
		Name:     "DELTA HEADING RATE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTurnCoordinatorBall Simvar
// args contain optional index and/or unit
func SimVarTurnCoordinatorBall(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "TURN COORDINATOR BALL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAngleOfAttackIndicator Simvar
// args contain optional index and/or unit
func SimVarAngleOfAttackIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ANGLE OF ATTACK INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRadioHeight Simvar
// args contain optional index and/or unit
func SimVarRadioHeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "RADIO HEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPartialPanelAdf Simvar
// args contain optional index and/or unit
func SimVarPartialPanelAdf(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ADF",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelAirspeed Simvar
// args contain optional index and/or unit
func SimVarPartialPanelAirspeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL AIRSPEED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelAltimeter Simvar
// args contain optional index and/or unit
func SimVarPartialPanelAltimeter(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ALTIMETER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelAttitude Simvar
// args contain optional index and/or unit
func SimVarPartialPanelAttitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ATTITUDE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelComm Simvar
// args contain optional index and/or unit
func SimVarPartialPanelComm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL COMM",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelCompass Simvar
// args contain optional index and/or unit
func SimVarPartialPanelCompass(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL COMPASS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelElectrical Simvar
// args contain optional index and/or unit
func SimVarPartialPanelElectrical(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ELECTRICAL",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelAvionics Simvar
// args contain optional index and/or unit
func SimVarPartialPanelAvionics(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL AVIONICS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPartialPanelEngine Simvar
// args contain optional index and/or unit
func SimVarPartialPanelEngine(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL ENGINE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelFuelIndicator Simvar
// args contain optional index and/or unit
func SimVarPartialPanelFuelIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL FUEL INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPartialPanelHeading Simvar
// args contain optional index and/or unit
func SimVarPartialPanelHeading(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL HEADING",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelVerticalVelocity Simvar
// args contain optional index and/or unit
func SimVarPartialPanelVerticalVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL VERTICAL VELOCITY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelTransponder Simvar
// args contain optional index and/or unit
func SimVarPartialPanelTransponder(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL TRANSPONDER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelNav Simvar
// args contain optional index and/or unit
func SimVarPartialPanelNav(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL NAV",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelPitot Simvar
// args contain optional index and/or unit
func SimVarPartialPanelPitot(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL PITOT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPartialPanelTurnCoordinator Simvar
// args contain optional index and/or unit
func SimVarPartialPanelTurnCoordinator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL TURN COORDINATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPartialPanelVacuum Simvar
// args contain optional index and/or unit
func SimVarPartialPanelVacuum(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PARTIAL PANEL VACUUM",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarMaxGForce Simvar
// args contain optional index and/or unit
func SimVarMaxGForce(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gforce")
	return SimVar{
		Index:    index,
		Name:     "MAX G FORCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMinGForce Simvar
// args contain optional index and/or unit
func SimVarMinGForce(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Gforce")
	return SimVar{
		Index:    index,
		Name:     "MIN G FORCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSuctionPressure Simvar
// args contain optional index and/or unit
func SimVarSuctionPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "inHg")
	return SimVar{
		Index:    index,
		Name:     "SUCTION PRESSURE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAvionicsMasterSwitch Simvar
// args contain optional index and/or unit
func SimVarAvionicsMasterSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AVIONICS MASTER SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavSound Simvar
// args contain optional index and/or unit
func SimVarNavSound(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV SOUND:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDmeSound Simvar
// args contain optional index and/or unit
func SimVarDmeSound(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "DME SOUND",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfSound Simvar
// args contain optional index and/or unit
func SimVarAdfSound(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ADF SOUND:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMarkerSound Simvar
// args contain optional index and/or unit
func SimVarMarkerSound(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "MARKER SOUND",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComTransmit Simvar
// args contain optional index and/or unit
func SimVarComTransmit(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "COM TRANSMIT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComRecieveAll Simvar
// args contain optional index and/or unit
func SimVarComRecieveAll(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "COM RECIEVE ALL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComActiveFrequency Simvar
// args contain optional index and/or unit
func SimVarComActiveFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Frequency BCD16")
	return SimVar{
		Index:    index,
		Name:     "COM ACTIVE FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComStandbyFrequency Simvar
// args contain optional index and/or unit
func SimVarComStandbyFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Frequency BCD16")
	return SimVar{
		Index:    index,
		Name:     "COM STANDBY FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComStatus Simvar
// args contain optional index and/or unit
func SimVarComStatus(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "COM STATUS:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavAvailable Simvar
// args contain optional index and/or unit
func SimVarNavAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV AVAILABLE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavActiveFrequency Simvar
// args contain optional index and/or unit
func SimVarNavActiveFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "MHz")
	return SimVar{
		Index:    index,
		Name:     "NAV ACTIVE FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavStandbyFrequency Simvar
// args contain optional index and/or unit
func SimVarNavStandbyFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "MHz")
	return SimVar{
		Index:    index,
		Name:     "NAV STANDBY FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavSignal Simvar
// args contain optional index and/or unit
func SimVarNavSignal(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NAV SIGNAL:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavHasNav Simvar
// args contain optional index and/or unit
func SimVarNavHasNav(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV HAS NAV:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavHasLocalizer Simvar
// args contain optional index and/or unit
func SimVarNavHasLocalizer(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV HAS LOCALIZER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavHasDme Simvar
// args contain optional index and/or unit
func SimVarNavHasDme(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV HAS DME:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavHasGlideSlope Simvar
// args contain optional index and/or unit
func SimVarNavHasGlideSlope(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV HAS GLIDE SLOPE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavBackCourseFlags Simvar
// args contain optional index and/or unit
func SimVarNavBackCourseFlags(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV BACK COURSE FLAGS:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavMagvar Simvar
// args contain optional index and/or unit
func SimVarNavMagvar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV MAGVAR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavRadial Simvar
// args contain optional index and/or unit
func SimVarNavRadial(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV RADIAL:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavRadialError Simvar
// args contain optional index and/or unit
func SimVarNavRadialError(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV RADIAL ERROR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavLocalizer Simvar
// args contain optional index and/or unit
func SimVarNavLocalizer(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV LOCALIZER:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGlideSlopeError Simvar
// args contain optional index and/or unit
func SimVarNavGlideSlopeError(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV GLIDE SLOPE ERROR:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavCdi Simvar
// args contain optional index and/or unit
func SimVarNavCdi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NAV CDI:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGsi Simvar
// args contain optional index and/or unit
func SimVarNavGsi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NAV GSI:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavTofrom Simvar
// args contain optional index and/or unit
func SimVarNavTofrom(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "NAV TOFROM:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGsFlag Simvar
// args contain optional index and/or unit
func SimVarNavGsFlag(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "NAV GS FLAG:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavObs Simvar
// args contain optional index and/or unit
func SimVarNavObs(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV OBS:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavDme Simvar
// args contain optional index and/or unit
func SimVarNavDme(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Nautical miles")
	return SimVar{
		Index:    index,
		Name:     "NAV DME:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavDmespeed Simvar
// args contain optional index and/or unit
func SimVarNavDmespeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "NAV DMESPEED:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfActiveFrequency Simvar
// args contain optional index and/or unit
func SimVarAdfActiveFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Frequency ADF BCD32")
	return SimVar{
		Index:    index,
		Name:     "ADF ACTIVE FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfStandbyFrequency Simvar
// args contain optional index and/or unit
func SimVarAdfStandbyFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Hz")
	return SimVar{
		Index:    index,
		Name:     "ADF STANDBY FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfRadial Simvar
// args contain optional index and/or unit
func SimVarAdfRadial(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "ADF RADIAL:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfSignal Simvar
// args contain optional index and/or unit
func SimVarAdfSignal(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ADF SIGNAL:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTransponderCode Simvar
// args contain optional index and/or unit
func SimVarTransponderCode(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "BCO16")
	return SimVar{
		Index:    index,
		Name:     "TRANSPONDER CODE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMarkerBeaconState Simvar
// args contain optional index and/or unit
func SimVarMarkerBeaconState(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "MARKER BEACON STATE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarInnerMarker Simvar
// args contain optional index and/or unit
func SimVarInnerMarker(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "INNER MARKER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarMiddleMarker Simvar
// args contain optional index and/or unit
func SimVarMiddleMarker(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "MIDDLE MARKER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarOuterMarker Simvar
// args contain optional index and/or unit
func SimVarOuterMarker(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "OUTER MARKER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarNavRawGlideSlope Simvar
// args contain optional index and/or unit
func SimVarNavRawGlideSlope(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV RAW GLIDE SLOPE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfCard Simvar
// args contain optional index and/or unit
func SimVarAdfCard(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "ADF CARD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiCdiNeedle Simvar
// args contain optional index and/or unit
func SimVarHsiCdiNeedle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "HSI CDI NEEDLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiGsiNeedle Simvar
// args contain optional index and/or unit
func SimVarHsiGsiNeedle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "HSI GSI NEEDLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiCdiNeedleValid Simvar
// args contain optional index and/or unit
func SimVarHsiCdiNeedleValid(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HSI CDI NEEDLE VALID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiGsiNeedleValid Simvar
// args contain optional index and/or unit
func SimVarHsiGsiNeedleValid(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HSI GSI NEEDLE VALID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiTfFlags Simvar
// args contain optional index and/or unit
func SimVarHsiTfFlags(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "HSI TF FLAGS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiBearingValid Simvar
// args contain optional index and/or unit
func SimVarHsiBearingValid(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HSI BEARING VALID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiBearing Simvar
// args contain optional index and/or unit
func SimVarHsiBearing(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "HSI BEARING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiHasLocalizer Simvar
// args contain optional index and/or unit
func SimVarHsiHasLocalizer(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HSI HAS LOCALIZER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiSpeed Simvar
// args contain optional index and/or unit
func SimVarHsiSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "HSI SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiDistance Simvar
// args contain optional index and/or unit
func SimVarHsiDistance(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Nautical miles")
	return SimVar{
		Index:    index,
		Name:     "HSI DISTANCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsPositionLat Simvar
// args contain optional index and/or unit
func SimVarGpsPositionLat(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION LAT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsPositionLon Simvar
// args contain optional index and/or unit
func SimVarGpsPositionLon(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION LON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsPositionAlt Simvar
// args contain optional index and/or unit
func SimVarGpsPositionAlt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS POSITION ALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsMagvar Simvar
// args contain optional index and/or unit
func SimVarGpsMagvar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS MAGVAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsActiveFlightPlan Simvar
// args contain optional index and/or unit
func SimVarGpsIsActiveFlightPlan(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE FLIGHT PLAN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsActiveWayPoint Simvar
// args contain optional index and/or unit
func SimVarGpsIsActiveWayPoint(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE WAY POINT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsArrived Simvar
// args contain optional index and/or unit
func SimVarGpsIsArrived(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS ARRIVED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsDirecttoFlightplan Simvar
// args contain optional index and/or unit
func SimVarGpsIsDirecttoFlightplan(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS DIRECTTO FLIGHTPLAN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsGroundSpeed Simvar
// args contain optional index and/or unit
func SimVarGpsGroundSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters per second")
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsGroundTrueHeading Simvar
// args contain optional index and/or unit
func SimVarGpsGroundTrueHeading(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND TRUE HEADING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsGroundMagneticTrack Simvar
// args contain optional index and/or unit
func SimVarGpsGroundMagneticTrack(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND MAGNETIC TRACK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsGroundTrueTrack Simvar
// args contain optional index and/or unit
func SimVarGpsGroundTrueTrack(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS GROUND TRUE TRACK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpDistance Simvar
// args contain optional index and/or unit
func SimVarGpsWpDistance(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS WP DISTANCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpBearing Simvar
// args contain optional index and/or unit
func SimVarGpsWpBearing(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS WP BEARING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpTrueBearing Simvar
// args contain optional index and/or unit
func SimVarGpsWpTrueBearing(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRUE BEARING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpCrossTrk Simvar
// args contain optional index and/or unit
func SimVarGpsWpCrossTrk(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS WP CROSS TRK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpDesiredTrack Simvar
// args contain optional index and/or unit
func SimVarGpsWpDesiredTrack(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS WP DESIRED TRACK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpTrueReqHdg Simvar
// args contain optional index and/or unit
func SimVarGpsWpTrueReqHdg(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRUE REQ HDG",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpVerticalSpeed Simvar
// args contain optional index and/or unit
func SimVarGpsWpVerticalSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters per second")
	return SimVar{
		Index:    index,
		Name:     "GPS WP VERTICAL SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpTrackAngleError Simvar
// args contain optional index and/or unit
func SimVarGpsWpTrackAngleError(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS WP TRACK ANGLE ERROR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsEte Simvar
// args contain optional index and/or unit
func SimVarGpsEte(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "GPS ETE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsEta Simvar
// args contain optional index and/or unit
func SimVarGpsEta(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "GPS ETA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpNextLat Simvar
// args contain optional index and/or unit
func SimVarGpsWpNextLat(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT LAT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpNextLon Simvar
// args contain optional index and/or unit
func SimVarGpsWpNextLon(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT LON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpNextAlt Simvar
// args contain optional index and/or unit
func SimVarGpsWpNextAlt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT ALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpPrevValid Simvar
// args contain optional index and/or unit
func SimVarGpsWpPrevValid(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV VALID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpPrevLat Simvar
// args contain optional index and/or unit
func SimVarGpsWpPrevLat(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV LAT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpPrevLon Simvar
// args contain optional index and/or unit
func SimVarGpsWpPrevLon(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV LON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpPrevAlt Simvar
// args contain optional index and/or unit
func SimVarGpsWpPrevAlt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV ALT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpEte Simvar
// args contain optional index and/or unit
func SimVarGpsWpEte(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "GPS WP ETE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpEta Simvar
// args contain optional index and/or unit
func SimVarGpsWpEta(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "GPS WP ETA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsCourseToSteer Simvar
// args contain optional index and/or unit
func SimVarGpsCourseToSteer(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "GPS COURSE TO STEER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpIndex Simvar
// args contain optional index and/or unit
func SimVarGpsFlightPlanWpIndex(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS FLIGHT PLAN WP INDEX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpCount Simvar
// args contain optional index and/or unit
func SimVarGpsFlightPlanWpCount(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS FLIGHT PLAN WP COUNT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsActiveWpLocked Simvar
// args contain optional index and/or unit
func SimVarGpsIsActiveWpLocked(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS ACTIVE WP LOCKED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsApproachLoaded Simvar
// args contain optional index and/or unit
func SimVarGpsIsApproachLoaded(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS APPROACH LOADED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsIsApproachActive Simvar
// args contain optional index and/or unit
func SimVarGpsIsApproachActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS IS APPROACH ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachMode Simvar
// args contain optional index and/or unit
func SimVarGpsApproachMode(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH MODE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachWpType Simvar
// args contain optional index and/or unit
func SimVarGpsApproachWpType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachIsWpRunway Simvar
// args contain optional index and/or unit
func SimVarGpsApproachIsWpRunway(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS WP RUNWAY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachSegmentType Simvar
// args contain optional index and/or unit
func SimVarGpsApproachSegmentType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH SEGMENT TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachApproachIndex Simvar
// args contain optional index and/or unit
func SimVarGpsApproachApproachIndex(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH INDEX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachApproachType Simvar
// args contain optional index and/or unit
func SimVarGpsApproachApproachType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachTransitionIndex Simvar
// args contain optional index and/or unit
func SimVarGpsApproachTransitionIndex(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TRANSITION INDEX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachIsFinal Simvar
// args contain optional index and/or unit
func SimVarGpsApproachIsFinal(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS FINAL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachIsMissed Simvar
// args contain optional index and/or unit
func SimVarGpsApproachIsMissed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH IS MISSED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachTimezoneDeviation Simvar
// args contain optional index and/or unit
func SimVarGpsApproachTimezoneDeviation(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TIMEZONE DEVIATION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachWpIndex Simvar
// args contain optional index and/or unit
func SimVarGpsApproachWpIndex(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP INDEX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachWpCount Simvar
// args contain optional index and/or unit
func SimVarGpsApproachWpCount(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH WP COUNT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsDrivesNav1 Simvar
func SimVarGpsDrivesNav1(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPS DRIVES NAV1",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComReceiveAll Simvar
// args contain optional index and/or unit
func SimVarComReceiveAll(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "COM RECEIVE ALL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComAvailable Simvar
// args contain optional index and/or unit
func SimVarComAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "COM AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarComTest Simvar
// args contain optional index and/or unit
func SimVarComTest(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "COM TEST:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTransponderAvailable Simvar
// args contain optional index and/or unit
func SimVarTransponderAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TRANSPONDER AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfAvailable Simvar
// args contain optional index and/or unit
func SimVarAdfAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ADF AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfFrequency Simvar
// args contain optional index and/or unit
func SimVarAdfFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Frequency BCD16")
	return SimVar{
		Index:    index,
		Name:     "ADF FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfExtFrequency Simvar
// args contain optional index and/or unit
func SimVarAdfExtFrequency(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Frequency BCD16")
	return SimVar{
		Index:    index,
		Name:     "ADF EXT FREQUENCY:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfIdent Simvar
// args contain optional index and/or unit
func SimVarAdfIdent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "ADF IDENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAdfName Simvar
// args contain optional index and/or unit
func SimVarAdfName(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "ADF NAME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavIdent Simvar
// args contain optional index and/or unit
func SimVarNavIdent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "NAV IDENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavName Simvar
// args contain optional index and/or unit
func SimVarNavName(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "NAV NAME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavCodes Simvar
// args contain optional index and/or unit
func SimVarNavCodes(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Flags")
	return SimVar{
		Index:    index,
		Name:     "NAV CODES:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGlideSlope Simvar
// args contain optional index and/or unit
func SimVarNavGlideSlope(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "NAV GLIDE SLOPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavRelativeBearingToStation Simvar
// args contain optional index and/or unit
func SimVarNavRelativeBearingToStation(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "NAV RELATIVE BEARING TO STATION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSelectedDme Simvar
// args contain optional index and/or unit
func SimVarSelectedDme(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "SELECTED DME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpNextId Simvar
// args contain optional index and/or unit
func SimVarGpsWpNextId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "GPS WP NEXT ID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsWpPrevId Simvar
// args contain optional index and/or unit
func SimVarGpsWpPrevId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "GPS WP PREV ID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsTargetDistance Simvar
// args contain optional index and/or unit
func SimVarGpsTargetDistance(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS TARGET DISTANCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsTargetAltitude Simvar
// args contain optional index and/or unit
func SimVarGpsTargetAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "GPS TARGET ALTITUDE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarYokeYPosition Simvar
// args contain optional index and/or unit
func SimVarYokeYPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "YOKE Y POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarYokeXPosition Simvar
// args contain optional index and/or unit
func SimVarYokeXPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "YOKE X POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRudderPedalPosition Simvar
// args contain optional index and/or unit
func SimVarRudderPedalPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "RUDDER PEDAL POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRudderPosition Simvar
// args contain optional index and/or unit
func SimVarRudderPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "RUDDER POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElevatorPosition Simvar
// args contain optional index and/or unit
func SimVarElevatorPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAileronPosition Simvar
// args contain optional index and/or unit
func SimVarAileronPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "AILERON POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElevatorTrimPosition Simvar
// args contain optional index and/or unit
func SimVarElevatorTrimPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElevatorTrimIndicator Simvar
// args contain optional index and/or unit
func SimVarElevatorTrimIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarElevatorTrimPct Simvar
// args contain optional index and/or unit
func SimVarElevatorTrimPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR TRIM PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBrakeLeftPosition Simvar
// args contain optional index and/or unit
func SimVarBrakeLeftPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "BRAKE LEFT POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarBrakeRightPosition Simvar
// args contain optional index and/or unit
func SimVarBrakeRightPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "BRAKE RIGHT POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarBrakeIndicator Simvar
// args contain optional index and/or unit
func SimVarBrakeIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "BRAKE INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBrakeParkingPosition Simvar
// args contain optional index and/or unit
func SimVarBrakeParkingPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "BRAKE PARKING POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarBrakeParkingIndicator Simvar
// args contain optional index and/or unit
func SimVarBrakeParkingIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "BRAKE PARKING INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSpoilersArmed Simvar
// args contain optional index and/or unit
func SimVarSpoilersArmed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SPOILERS ARMED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSpoilersHandlePosition Simvar
// args contain optional index and/or unit
func SimVarSpoilersHandlePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "SPOILERS HANDLE POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarSpoilersLeftPosition Simvar
// args contain optional index and/or unit
func SimVarSpoilersLeftPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "SPOILERS LEFT POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSpoilersRightPosition Simvar
// args contain optional index and/or unit
func SimVarSpoilersRightPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "SPOILERS RIGHT POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlapsHandlePercent Simvar
// args contain optional index and/or unit
func SimVarFlapsHandlePercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FLAPS HANDLE PERCENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlapsHandleIndex Simvar
// args contain optional index and/or unit
func SimVarFlapsHandleIndex(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "FLAPS HANDLE INDEX",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFlapsNumHandlePositions Simvar
// args contain optional index and/or unit
func SimVarFlapsNumHandlePositions(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "FLAPS NUM HANDLE POSITIONS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsLeftPercent Simvar
// args contain optional index and/or unit
func SimVarTrailingEdgeFlapsLeftPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS LEFT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsRightPercent Simvar
// args contain optional index and/or unit
func SimVarTrailingEdgeFlapsRightPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS RIGHT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsLeftAngle Simvar
// args contain optional index and/or unit
func SimVarTrailingEdgeFlapsLeftAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS LEFT ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsRightAngle Simvar
// args contain optional index and/or unit
func SimVarTrailingEdgeFlapsRightAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "TRAILING EDGE FLAPS RIGHT ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsLeftPercent Simvar
// args contain optional index and/or unit
func SimVarLeadingEdgeFlapsLeftPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS LEFT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsRightPercent Simvar
// args contain optional index and/or unit
func SimVarLeadingEdgeFlapsRightPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS RIGHT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsLeftAngle Simvar
// args contain optional index and/or unit
func SimVarLeadingEdgeFlapsLeftAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS LEFT ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsRightAngle Simvar
// args contain optional index and/or unit
func SimVarLeadingEdgeFlapsRightAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "LEADING EDGE FLAPS RIGHT ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsGearRetractable Simvar
// args contain optional index and/or unit
func SimVarIsGearRetractable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS GEAR RETRACTABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsGearSkis Simvar
// args contain optional index and/or unit
func SimVarIsGearSkis(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS GEAR SKIS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsGearFloats Simvar
// args contain optional index and/or unit
func SimVarIsGearFloats(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS GEAR FLOATS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsGearSkids Simvar
// args contain optional index and/or unit
func SimVarIsGearSkids(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS GEAR SKIDS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsGearWheels Simvar
// args contain optional index and/or unit
func SimVarIsGearWheels(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS GEAR WHEELS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearHandlePosition Simvar
// args contain optional index and/or unit
func SimVarGearHandlePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GEAR HANDLE POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGearHydraulicPressure Simvar
// args contain optional index and/or unit
func SimVarGearHydraulicPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "psf")
	return SimVar{
		Index:    index,
		Name:     "GEAR HYDRAULIC PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTailwheelLockOn Simvar
// args contain optional index and/or unit
func SimVarTailwheelLockOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TAILWHEEL LOCK ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearCenterPosition Simvar
// args contain optional index and/or unit
func SimVarGearCenterPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGearLeftPosition Simvar
// args contain optional index and/or unit
func SimVarGearLeftPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGearRightPosition Simvar
// args contain optional index and/or unit
func SimVarGearRightPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGearTailPosition Simvar
// args contain optional index and/or unit
func SimVarGearTailPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR TAIL POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearAuxPosition Simvar
// args contain optional index and/or unit
func SimVarGearAuxPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearPosition Simvar
// args contain optional index and/or unit
func SimVarGearPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GEAR POSITION:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGearAnimationPosition Simvar
// args contain optional index and/or unit
func SimVarGearAnimationPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "GEAR ANIMATION POSITION:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearTotalPctExtended Simvar
// args contain optional index and/or unit
func SimVarGearTotalPctExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percentage")
	return SimVar{
		Index:    index,
		Name:     "GEAR TOTAL PCT EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutoBrakeSwitchCb Simvar
// args contain optional index and/or unit
func SimVarAutoBrakeSwitchCb(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "AUTO BRAKE SWITCH CB",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterRudderHandlePosition Simvar
// args contain optional index and/or unit
func SimVarWaterRudderHandlePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WATER RUDDER HANDLE POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElevatorDeflection Simvar
// args contain optional index and/or unit
func SimVarElevatorDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarElevatorDeflectionPct Simvar
// args contain optional index and/or unit
func SimVarElevatorDeflectionPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ELEVATOR DEFLECTION PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterLeftRudderExtended Simvar
// args contain optional index and/or unit
func SimVarWaterLeftRudderExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percentage")
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterRightRudderExtended Simvar
// args contain optional index and/or unit
func SimVarWaterRightRudderExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percentage")
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearCenterSteerAngle Simvar
// args contain optional index and/or unit
func SimVarGearCenterSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearLeftSteerAngle Simvar
// args contain optional index and/or unit
func SimVarGearLeftSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearRightSteerAngle Simvar
// args contain optional index and/or unit
func SimVarGearRightSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearAuxSteerAngle Simvar
// args contain optional index and/or unit
func SimVarGearAuxSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearSteerAngle Simvar
// args contain optional index and/or unit
func SimVarGearSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR STEER ANGLE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAngle Simvar
// args contain optional index and/or unit
func SimVarWaterLeftRudderSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAngle Simvar
// args contain optional index and/or unit
func SimVarWaterRightRudderSteerAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER STEER ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearCenterSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarGearCenterSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR CENTER STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearLeftSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarGearLeftSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR LEFT STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearRightSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarGearRightSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR RIGHT STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearAuxSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarGearAuxSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR AUX STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarGearSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "GEAR STEER ANGLE PCT:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarWaterLeftRudderSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WATER LEFT RUDDER STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAnglePct Simvar
// args contain optional index and/or unit
func SimVarWaterRightRudderSteerAnglePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "WATER RIGHT RUDDER STEER ANGLE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronLeftDeflection Simvar
// args contain optional index and/or unit
func SimVarAileronLeftDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AILERON LEFT DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronLeftDeflectionPct Simvar
// args contain optional index and/or unit
func SimVarAileronLeftDeflectionPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "AILERON LEFT DEFLECTION PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronRightDeflection Simvar
// args contain optional index and/or unit
func SimVarAileronRightDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AILERON RIGHT DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronRightDeflectionPct Simvar
// args contain optional index and/or unit
func SimVarAileronRightDeflectionPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "AILERON RIGHT DEFLECTION PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronAverageDeflection Simvar
// args contain optional index and/or unit
func SimVarAileronAverageDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AILERON AVERAGE DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAileronTrim Simvar
// args contain optional index and/or unit
func SimVarAileronTrim(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AILERON TRIM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRudderDeflection Simvar
// args contain optional index and/or unit
func SimVarRudderDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "RUDDER DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRudderDeflectionPct Simvar
// args contain optional index and/or unit
func SimVarRudderDeflectionPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "RUDDER DEFLECTION PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRudderTrim Simvar
// args contain optional index and/or unit
func SimVarRudderTrim(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "RUDDER TRIM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlapsAvailable Simvar
// args contain optional index and/or unit
func SimVarFlapsAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLAPS AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearDamageBySpeed Simvar
// args contain optional index and/or unit
func SimVarGearDamageBySpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GEAR DAMAGE BY SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearSpeedExceeded Simvar
// args contain optional index and/or unit
func SimVarGearSpeedExceeded(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GEAR SPEED EXCEEDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlapDamageBySpeed Simvar
// args contain optional index and/or unit
func SimVarFlapDamageBySpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLAP DAMAGE BY SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFlapSpeedExceeded Simvar
// args contain optional index and/or unit
func SimVarFlapSpeedExceeded(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FLAP SPEED EXCEEDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCenterWheelRpm Simvar
// args contain optional index and/or unit
func SimVarCenterWheelRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "CENTER WHEEL RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLeftWheelRpm Simvar
// args contain optional index and/or unit
func SimVarLeftWheelRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "LEFT WHEEL RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRightWheelRpm Simvar
// args contain optional index and/or unit
func SimVarRightWheelRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "RIGHT WHEEL RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAvailable Simvar
// args contain optional index and/or unit
func SimVarAutopilotAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotMaster Simvar
// args contain optional index and/or unit
func SimVarAutopilotMaster(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MASTER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotNavSelected Simvar
// args contain optional index and/or unit
func SimVarAutopilotNavSelected(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT NAV SELECTED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotWingLeveler Simvar
// args contain optional index and/or unit
func SimVarAutopilotWingLeveler(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT WING LEVELER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotHeadingLock Simvar
// args contain optional index and/or unit
func SimVarAutopilotHeadingLock(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT HEADING LOCK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotHeadingLockDir Simvar
// args contain optional index and/or unit
func SimVarAutopilotHeadingLockDir(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT HEADING LOCK DIR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLock Simvar
// args contain optional index and/or unit
func SimVarAutopilotAltitudeLock(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ALTITUDE LOCK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLockVar Simvar
// args contain optional index and/or unit
func SimVarAutopilotAltitudeLockVar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ALTITUDE LOCK VAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAttitudeHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotAttitudeHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT ATTITUDE HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotGlideslopeHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotGlideslopeHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT GLIDESLOPE HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotPitchHoldRef Simvar
// args contain optional index and/or unit
func SimVarAutopilotPitchHoldRef(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT PITCH HOLD REF",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotApproachHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotApproachHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT APPROACH HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotBackcourseHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotBackcourseHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT BACKCOURSE HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotVerticalHoldVar Simvar
// args contain optional index and/or unit
func SimVarAutopilotVerticalHoldVar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet/minute")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT VERTICAL HOLD VAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorActive Simvar
// args contain optional index and/or unit
func SimVarAutopilotFlightDirectorActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorPitch Simvar
// args contain optional index and/or unit
func SimVarAutopilotFlightDirectorPitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR PITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorBank Simvar
// args contain optional index and/or unit
func SimVarAutopilotFlightDirectorBank(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT FLIGHT DIRECTOR BANK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotAirspeedHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AIRSPEED HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHoldVar Simvar
// args contain optional index and/or unit
func SimVarAutopilotAirspeedHoldVar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT AIRSPEED HOLD VAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotMachHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotMachHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MACH HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotMachHoldVar Simvar
// args contain optional index and/or unit
func SimVarAutopilotMachHoldVar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MACH HOLD VAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotYawDamper Simvar
// args contain optional index and/or unit
func SimVarAutopilotYawDamper(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT YAW DAMPER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotRpmHoldVar Simvar
// args contain optional index and/or unit
func SimVarAutopilotRpmHoldVar(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT RPM HOLD VAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotThrottleArm Simvar
// args contain optional index and/or unit
func SimVarAutopilotThrottleArm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT THROTTLE ARM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotTakeoffPowerActive Simvar
// args contain optional index and/or unit
func SimVarAutopilotTakeoffPowerActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT TAKEOFF POWER ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutothrottleActive Simvar
// args contain optional index and/or unit
func SimVarAutothrottleActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOTHROTTLE ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotNav1Lock Simvar
func SimVarAutopilotNav1Lock(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT NAV1 LOCK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotVerticalHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotVerticalHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT VERTICAL HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotRpmHold Simvar
// args contain optional index and/or unit
func SimVarAutopilotRpmHold(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT RPM HOLD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAutopilotMaxBank Simvar
// args contain optional index and/or unit
func SimVarAutopilotMaxBank(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AUTOPILOT MAX BANK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWheelRpm Simvar
// args contain optional index and/or unit
func SimVarWheelRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "WHEEL RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAuxWheelRpm Simvar
// args contain optional index and/or unit
func SimVarAuxWheelRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "AUX WHEEL RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWheelRotationAngle Simvar
// args contain optional index and/or unit
func SimVarWheelRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "WHEEL ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCenterWheelRotationAngle Simvar
// args contain optional index and/or unit
func SimVarCenterWheelRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "CENTER WHEEL ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLeftWheelRotationAngle Simvar
// args contain optional index and/or unit
func SimVarLeftWheelRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "LEFT WHEEL ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRightWheelRotationAngle Simvar
// args contain optional index and/or unit
func SimVarRightWheelRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "RIGHT WHEEL ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAuxWheelRotationAngle Simvar
// args contain optional index and/or unit
func SimVarAuxWheelRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "AUX WHEEL ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearEmergencyHandlePosition Simvar
// args contain optional index and/or unit
func SimVarGearEmergencyHandlePosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GEAR EMERGENCY HANDLE POSITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGearWarning Simvar
// args contain optional index and/or unit
func SimVarGearWarning(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "GEAR WARNING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAntiskidBrakesActive Simvar
// args contain optional index and/or unit
func SimVarAntiskidBrakesActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ANTISKID BRAKES ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRetractFloatSwitch Simvar
// args contain optional index and/or unit
func SimVarRetractFloatSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RETRACT FLOAT SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRetractLeftFloatExtended Simvar
// args contain optional index and/or unit
func SimVarRetractLeftFloatExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "RETRACT LEFT FLOAT EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRetractRightFloatExtended Simvar
// args contain optional index and/or unit
func SimVarRetractRightFloatExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent")
	return SimVar{
		Index:    index,
		Name:     "RETRACT RIGHT FLOAT EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSteerInputControl Simvar
// args contain optional index and/or unit
func SimVarSteerInputControl(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "STEER INPUT CONTROL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientDensity Simvar
// args contain optional index and/or unit
func SimVarAmbientDensity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Slugs per cubic feet")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT DENSITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientTemperature Simvar
// args contain optional index and/or unit
func SimVarAmbientTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT TEMPERATURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientPressure Simvar
// args contain optional index and/or unit
func SimVarAmbientPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "inHg")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientWindVelocity Simvar
// args contain optional index and/or unit
func SimVarAmbientWindVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientWindDirection Simvar
// args contain optional index and/or unit
func SimVarAmbientWindDirection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND DIRECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientWindX Simvar
// args contain optional index and/or unit
func SimVarAmbientWindX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters per second")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND X",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientWindY Simvar
// args contain optional index and/or unit
func SimVarAmbientWindY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters per second")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND Y",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientWindZ Simvar
// args contain optional index and/or unit
func SimVarAmbientWindZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters per second")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT WIND Z",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientPrecipState Simvar
// args contain optional index and/or unit
func SimVarAmbientPrecipState(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Mask")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT PRECIP STATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAircraftWindX Simvar
// args contain optional index and/or unit
func SimVarAircraftWindX(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND X",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAircraftWindY Simvar
// args contain optional index and/or unit
func SimVarAircraftWindY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND Y",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAircraftWindZ Simvar
// args contain optional index and/or unit
func SimVarAircraftWindZ(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRCRAFT WIND Z",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBarometerPressure Simvar
// args contain optional index and/or unit
func SimVarBarometerPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Millibars")
	return SimVar{
		Index:    index,
		Name:     "BAROMETER PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSeaLevelPressure Simvar
// args contain optional index and/or unit
func SimVarSeaLevelPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Millibars")
	return SimVar{
		Index:    index,
		Name:     "SEA LEVEL PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalAirTemperature Simvar
// args contain optional index and/or unit
func SimVarTotalAirTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Celsius")
	return SimVar{
		Index:    index,
		Name:     "TOTAL AIR TEMPERATURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWindshieldRainEffectAvailable Simvar
// args contain optional index and/or unit
func SimVarWindshieldRainEffectAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "WINDSHIELD RAIN EFFECT AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientInCloud Simvar
// args contain optional index and/or unit
func SimVarAmbientInCloud(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT IN CLOUD",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAmbientVisibility Simvar
// args contain optional index and/or unit
func SimVarAmbientVisibility(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "AMBIENT VISIBILITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStandardAtmTemperature Simvar
// args contain optional index and/or unit
func SimVarStandardAtmTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "STANDARD ATM TEMPERATURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorBrakeHandlePos Simvar
// args contain optional index and/or unit
func SimVarRotorBrakeHandlePos(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ROTOR BRAKE HANDLE POS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorBrakeActive Simvar
// args contain optional index and/or unit
func SimVarRotorBrakeActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR BRAKE ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorClutchSwitchPos Simvar
// args contain optional index and/or unit
func SimVarRotorClutchSwitchPos(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR CLUTCH SWITCH POS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorClutchActive Simvar
// args contain optional index and/or unit
func SimVarRotorClutchActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR CLUTCH ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorTemperature Simvar
// args contain optional index and/or unit
func SimVarRotorTemperature(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rankine")
	return SimVar{
		Index:    index,
		Name:     "ROTOR TEMPERATURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorChipDetected Simvar
// args contain optional index and/or unit
func SimVarRotorChipDetected(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR CHIP DETECTED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorGovSwitchPos Simvar
// args contain optional index and/or unit
func SimVarRotorGovSwitchPos(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR GOV SWITCH POS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorGovActive Simvar
// args contain optional index and/or unit
func SimVarRotorGovActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ROTOR GOV ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorLateralTrimPct Simvar
// args contain optional index and/or unit
func SimVarRotorLateralTrimPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ROTOR LATERAL TRIM PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorRpmPct Simvar
// args contain optional index and/or unit
func SimVarRotorRpmPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ROTOR RPM PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSmokeEnable Simvar
// args contain optional index and/or unit
func SimVarSmokeEnable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SMOKE ENABLE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarSmokesystemAvailable Simvar
// args contain optional index and/or unit
func SimVarSmokesystemAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SMOKESYSTEM AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPitotHeat Simvar
// args contain optional index and/or unit
func SimVarPitotHeat(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PITOT HEAT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFoldingWingLeftPercent Simvar
// args contain optional index and/or unit
func SimVarFoldingWingLeftPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FOLDING WING LEFT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarFoldingWingRightPercent Simvar
// args contain optional index and/or unit
func SimVarFoldingWingRightPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "FOLDING WING RIGHT PERCENT",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarCanopyOpen Simvar
// args contain optional index and/or unit
func SimVarCanopyOpen(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CANOPY OPEN",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTailhookPosition Simvar
// args contain optional index and/or unit
func SimVarTailhookPosition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "TAILHOOK POSITION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarExitOpen Simvar
// args contain optional index and/or unit
func SimVarExitOpen(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "EXIT OPEN:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarStallHornAvailable Simvar
// args contain optional index and/or unit
func SimVarStallHornAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "STALL HORN AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEngineMixureAvailable Simvar
// args contain optional index and/or unit
func SimVarEngineMixureAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ENGINE MIXURE AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCarbHeatAvailable Simvar
// args contain optional index and/or unit
func SimVarCarbHeatAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CARB HEAT AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSpoilerAvailable Simvar
// args contain optional index and/or unit
func SimVarSpoilerAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SPOILER AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsTailDragger Simvar
// args contain optional index and/or unit
func SimVarIsTailDragger(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS TAIL DRAGGER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStrobesAvailable Simvar
// args contain optional index and/or unit
func SimVarStrobesAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "STROBES AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarToeBrakesAvailable Simvar
// args contain optional index and/or unit
func SimVarToeBrakesAvailable(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TOE BRAKES AVAILABLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackState Simvar
// args contain optional index and/or unit
func SimVarPushbackState(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK STATE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalMasterBattery Simvar
// args contain optional index and/or unit
func SimVarElectricalMasterBattery(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MASTER BATTERY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalTotalLoadAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalTotalLoadAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL TOTAL LOAD AMPS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalBatteryLoad Simvar
// args contain optional index and/or unit
func SimVarElectricalBatteryLoad(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY LOAD",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalBatteryVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalBatteryVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY VOLTAGE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalMainBusVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalMainBusVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MAIN BUS VOLTAGE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalMainBusAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalMainBusAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL MAIN BUS AMPS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalAvionicsBusVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL AVIONICS BUS VOLTAGE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalAvionicsBusAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL AVIONICS BUS AMPS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalHotBatteryBusVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL HOT BATTERY BUS VOLTAGE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalHotBatteryBusAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL HOT BATTERY BUS AMPS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalBatteryBusVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalBatteryBusVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY BUS VOLTAGE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalBatteryBusAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalBatteryBusAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL BATTERY BUS AMPS",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalGenaltBusVoltage Simvar
// args contain optional index and/or unit
func SimVarElectricalGenaltBusVoltage(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL GENALT BUS VOLTAGE:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarElectricalGenaltBusAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalGenaltBusAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amperes")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL GENALT BUS AMPS:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarCircuitGeneralPanelOn Simvar
// args contain optional index and/or unit
func SimVarCircuitGeneralPanelOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GENERAL PANEL ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitFlapMotorOn Simvar
// args contain optional index and/or unit
func SimVarCircuitFlapMotorOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT FLAP MOTOR ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitGearMotorOn Simvar
// args contain optional index and/or unit
func SimVarCircuitGearMotorOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GEAR MOTOR ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitAutopilotOn Simvar
// args contain optional index and/or unit
func SimVarCircuitAutopilotOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTOPILOT ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitAvionicsOn Simvar
// args contain optional index and/or unit
func SimVarCircuitAvionicsOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AVIONICS ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitPitotHeatOn Simvar
// args contain optional index and/or unit
func SimVarCircuitPitotHeatOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT PITOT HEAT ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitPropSyncOn Simvar
// args contain optional index and/or unit
func SimVarCircuitPropSyncOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT PROP SYNC ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitAutoFeatherOn Simvar
// args contain optional index and/or unit
func SimVarCircuitAutoFeatherOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTO FEATHER ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitAutoBrakesOn Simvar
// args contain optional index and/or unit
func SimVarCircuitAutoBrakesOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT AUTO BRAKES ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitStandyVacuumOn Simvar
// args contain optional index and/or unit
func SimVarCircuitStandyVacuumOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT STANDY VACUUM ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitMarkerBeaconOn Simvar
// args contain optional index and/or unit
func SimVarCircuitMarkerBeaconOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT MARKER BEACON ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitGearWarningOn Simvar
// args contain optional index and/or unit
func SimVarCircuitGearWarningOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT GEAR WARNING ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCircuitHydraulicPumpOn Simvar
// args contain optional index and/or unit
func SimVarCircuitHydraulicPumpOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CIRCUIT HYDRAULIC PUMP ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHydraulicPressure Simvar
// args contain optional index and/or unit
func SimVarHydraulicPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pound force per square foot")
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC PRESSURE:index",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHydraulicReservoirPercent Simvar
// args contain optional index and/or unit
func SimVarHydraulicReservoirPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC RESERVOIR PERCENT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarHydraulicSystemIntegrity Simvar
// args contain optional index and/or unit
func SimVarHydraulicSystemIntegrity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC SYSTEM INTEGRITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructuralDeiceSwitch Simvar
// args contain optional index and/or unit
func SimVarStructuralDeiceSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "STRUCTURAL DEICE SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWeight Simvar
// args contain optional index and/or unit
func SimVarTotalWeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMaxGrossWeight Simvar
// args contain optional index and/or unit
func SimVarMaxGrossWeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "MAX GROSS WEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEmptyWeight Simvar
// args contain optional index and/or unit
func SimVarEmptyWeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsUserSim Simvar
// args contain optional index and/or unit
func SimVarIsUserSim(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS USER SIM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSimDisabled Simvar
// args contain optional index and/or unit
func SimVarSimDisabled(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SIM DISABLED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGForce Simvar
// args contain optional index and/or unit
func SimVarGForce(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "GForce")
	return SimVar{
		Index:    index,
		Name:     "G FORCE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAtcHeavy Simvar
// args contain optional index and/or unit
func SimVarAtcHeavy(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "ATC HEAVY",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAutoCoordination Simvar
// args contain optional index and/or unit
func SimVarAutoCoordination(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "AUTO COORDINATION",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarRealism Simvar
// args contain optional index and/or unit
func SimVarRealism(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "REALISM",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTrueAirspeedSelected Simvar
// args contain optional index and/or unit
func SimVarTrueAirspeedSelected(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TRUE AIRSPEED SELECTED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarDesignSpeedVc Simvar
// args contain optional index and/or unit
func SimVarDesignSpeedVc(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VC",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMinDragVelocity Simvar
// args contain optional index and/or unit
func SimVarMinDragVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "MIN DRAG VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEstimatedCruiseSpeed Simvar
// args contain optional index and/or unit
func SimVarEstimatedCruiseSpeed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "ESTIMATED CRUISE SPEED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgPercent Simvar
// args contain optional index and/or unit
func SimVarCgPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CG PERCENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgPercentLateral Simvar
// args contain optional index and/or unit
func SimVarCgPercentLateral(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CG PERCENT LATERAL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsSlewActive Simvar
// args contain optional index and/or unit
func SimVarIsSlewActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS SLEW ACTIVE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarIsSlewAllowed Simvar
// args contain optional index and/or unit
func SimVarIsSlewAllowed(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS SLEW ALLOWED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAtcSuggestedMinRwyTakeoff Simvar
// args contain optional index and/or unit
func SimVarAtcSuggestedMinRwyTakeoff(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "ATC SUGGESTED MIN RWY TAKEOFF",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAtcSuggestedMinRwyLanding Simvar
// args contain optional index and/or unit
func SimVarAtcSuggestedMinRwyLanding(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "ATC SUGGESTED MIN RWY LANDING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPayloadStationWeight Simvar
// args contain optional index and/or unit
func SimVarPayloadStationWeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds")
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION WEIGHT:index",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarPayloadStationCount Simvar
// args contain optional index and/or unit
func SimVarPayloadStationCount(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION COUNT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarUserInputEnabled Simvar
// args contain optional index and/or unit
func SimVarUserInputEnabled(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "USER INPUT ENABLED",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTypicalDescentRate Simvar
// args contain optional index and/or unit
func SimVarTypicalDescentRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per minute")
	return SimVar{
		Index:    index,
		Name:     "TYPICAL DESCENT RATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarVisualModelRadius Simvar
// args contain optional index and/or unit
func SimVarVisualModelRadius(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "VISUAL MODEL RADIUS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCategory Simvar
// args contain optional index and/or unit
func SimVarCategory(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "CATEGORY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSigmaSqrt Simvar
// args contain optional index and/or unit
func SimVarSigmaSqrt(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "SIGMA SQRT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDynamicPressure Simvar
// args contain optional index and/or unit
func SimVarDynamicPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Pounds per square foot")
	return SimVar{
		Index:    index,
		Name:     "DYNAMIC PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalVelocity Simvar
// args contain optional index and/or unit
func SimVarTotalVelocity(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "TOTAL VELOCITY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAirspeedSelectIndicatedOrTrue Simvar
// args contain optional index and/or unit
func SimVarAirspeedSelectIndicatedOrTrue(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Knots")
	return SimVar{
		Index:    index,
		Name:     "AIRSPEED SELECT INDICATED OR TRUE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarVariometerRate Simvar
// args contain optional index and/or unit
func SimVarVariometerRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "VARIOMETER RATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarVariometerSwitch Simvar
// args contain optional index and/or unit
func SimVarVariometerSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "VARIOMETER SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDesignSpeedVs0 Simvar
func SimVarDesignSpeedVs0(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VS0",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDesignSpeedVs1 Simvar
func SimVarDesignSpeedVs1(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "DESIGN SPEED VS1",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressureAltitude Simvar
// args contain optional index and/or unit
func SimVarPressureAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Meters")
	return SimVar{
		Index:    index,
		Name:     "PRESSURE ALTITUDE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMagneticCompass Simvar
// args contain optional index and/or unit
func SimVarMagneticCompass(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Degrees")
	return SimVar{
		Index:    index,
		Name:     "MAGNETIC COMPASS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurnIndicatorRate Simvar
// args contain optional index and/or unit
func SimVarTurnIndicatorRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians per second")
	return SimVar{
		Index:    index,
		Name:     "TURN INDICATOR RATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTurnIndicatorSwitch Simvar
// args contain optional index and/or unit
func SimVarTurnIndicatorSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TURN INDICATOR SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarYokeYIndicator Simvar
// args contain optional index and/or unit
func SimVarYokeYIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "YOKE Y INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarYokeXIndicator Simvar
// args contain optional index and/or unit
func SimVarYokeXIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "YOKE X INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRudderPedalIndicator Simvar
// args contain optional index and/or unit
func SimVarRudderPedalIndicator(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Position")
	return SimVar{
		Index:    index,
		Name:     "RUDDER PEDAL INDICATOR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBrakeDependentHydraulicPressure Simvar
// args contain optional index and/or unit
func SimVarBrakeDependentHydraulicPressure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "foot pounds")
	return SimVar{
		Index:    index,
		Name:     "BRAKE DEPENDENT HYDRAULIC PRESSURE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPanelAntiIceSwitch Simvar
// args contain optional index and/or unit
func SimVarPanelAntiIceSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PANEL ANTI ICE SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWingArea Simvar
// args contain optional index and/or unit
func SimVarWingArea(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Square feet")
	return SimVar{
		Index:    index,
		Name:     "WING AREA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWingSpan Simvar
// args contain optional index and/or unit
func SimVarWingSpan(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "WING SPAN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBetaDot Simvar
// args contain optional index and/or unit
func SimVarBetaDot(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians per second")
	return SimVar{
		Index:    index,
		Name:     "BETA DOT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLinearClAlpha Simvar
// args contain optional index and/or unit
func SimVarLinearClAlpha(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Per radian")
	return SimVar{
		Index:    index,
		Name:     "LINEAR CL ALPHA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStallAlpha Simvar
// args contain optional index and/or unit
func SimVarStallAlpha(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "STALL ALPHA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZeroLiftAlpha Simvar
// args contain optional index and/or unit
func SimVarZeroLiftAlpha(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ZERO LIFT ALPHA",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgAftLimit Simvar
// args contain optional index and/or unit
func SimVarCgAftLimit(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CG AFT LIMIT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgFwdLimit Simvar
// args contain optional index and/or unit
func SimVarCgFwdLimit(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CG FWD LIMIT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgMaxMach Simvar
// args contain optional index and/or unit
func SimVarCgMaxMach(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Machs")
	return SimVar{
		Index:    index,
		Name:     "CG MAX MACH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCgMinMach Simvar
// args contain optional index and/or unit
func SimVarCgMinMach(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Machs")
	return SimVar{
		Index:    index,
		Name:     "CG MIN MACH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPayloadStationName Simvar
// args contain optional index and/or unit
func SimVarPayloadStationName(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "PAYLOAD STATION NAME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarElevonDeflection Simvar
// args contain optional index and/or unit
func SimVarElevonDeflection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ELEVON DEFLECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarExitType Simvar
// args contain optional index and/or unit
func SimVarExitType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "EXIT TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarExitPosx Simvar
// args contain optional index and/or unit
func SimVarExitPosx(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "EXIT POSX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarExitPosy Simvar
// args contain optional index and/or unit
func SimVarExitPosy(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "EXIT POSY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarExitPosz Simvar
// args contain optional index and/or unit
func SimVarExitPosz(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "EXIT POSZ",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDecisionHeight Simvar
// args contain optional index and/or unit
func SimVarDecisionHeight(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "DECISION HEIGHT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDecisionAltitudeMsl Simvar
// args contain optional index and/or unit
func SimVarDecisionAltitudeMsl(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "DECISION ALTITUDE MSL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEmptyWeightPitchMoi Simvar
// args contain optional index and/or unit
func SimVarEmptyWeightPitchMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT PITCH MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEmptyWeightRollMoi Simvar
// args contain optional index and/or unit
func SimVarEmptyWeightRollMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT ROLL MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEmptyWeightYawMoi Simvar
// args contain optional index and/or unit
func SimVarEmptyWeightYawMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT YAW MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarEmptyWeightCrossCoupledMoi Simvar
// args contain optional index and/or unit
func SimVarEmptyWeightCrossCoupledMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "EMPTY WEIGHT CROSS COUPLED MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWeightPitchMoi Simvar
// args contain optional index and/or unit
func SimVarTotalWeightPitchMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT PITCH MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWeightRollMoi Simvar
// args contain optional index and/or unit
func SimVarTotalWeightRollMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT ROLL MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWeightYawMoi Simvar
// args contain optional index and/or unit
func SimVarTotalWeightYawMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT YAW MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTotalWeightCrossCoupledMoi Simvar
// args contain optional index and/or unit
func SimVarTotalWeightCrossCoupledMoi(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "slug feet squared")
	return SimVar{
		Index:    index,
		Name:     "TOTAL WEIGHT CROSS COUPLED MOI",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarWaterBallastValve Simvar
// args contain optional index and/or unit
func SimVarWaterBallastValve(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "WATER BALLAST VALVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarMaxRatedEngineRpm Simvar
// args contain optional index and/or unit
func SimVarMaxRatedEngineRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Rpm")
	return SimVar{
		Index:    index,
		Name:     "MAX RATED ENGINE RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFullThrottleThrustToWeightRatio Simvar
// args contain optional index and/or unit
func SimVarFullThrottleThrustToWeightRatio(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "FULL THROTTLE THRUST TO WEIGHT RATIO",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropAutoCruiseActive Simvar
// args contain optional index and/or unit
func SimVarPropAutoCruiseActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PROP AUTO CRUISE ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropRotationAngle Simvar
// args contain optional index and/or unit
func SimVarPropRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PROP ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropBetaMax Simvar
// args contain optional index and/or unit
func SimVarPropBetaMax(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MAX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropBetaMin Simvar
// args contain optional index and/or unit
func SimVarPropBetaMin(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MIN",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPropBetaMinReverse Simvar
// args contain optional index and/or unit
func SimVarPropBetaMinReverse(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PROP BETA MIN REVERSE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFuelSelectedTransferMode Simvar
// args contain optional index and/or unit
func SimVarFuelSelectedTransferMode(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "FUEL SELECTED TRANSFER MODE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDroppableObjectsUiName Simvar
// args contain optional index and/or unit
func SimVarDroppableObjectsUiName(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "DROPPABLE OBJECTS UI NAME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarManualFuelPumpHandle Simvar
// args contain optional index and/or unit
func SimVarManualFuelPumpHandle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "MANUAL FUEL PUMP HANDLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarBleedAirSourceControl Simvar
// args contain optional index and/or unit
func SimVarBleedAirSourceControl(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "BLEED AIR SOURCE CONTROL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarElectricalOldChargingAmps Simvar
// args contain optional index and/or unit
func SimVarElectricalOldChargingAmps(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Amps")
	return SimVar{
		Index:    index,
		Name:     "ELECTRICAL OLD CHARGING AMPS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHydraulicSwitch Simvar
// args contain optional index and/or unit
func SimVarHydraulicSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "HYDRAULIC SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarConcordeVisorNoseHandle Simvar
// args contain optional index and/or unit
func SimVarConcordeVisorNoseHandle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "CONCORDE VISOR NOSE HANDLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarConcordeVisorPositionPercent Simvar
// args contain optional index and/or unit
func SimVarConcordeVisorPositionPercent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "CONCORDE VISOR POSITION PERCENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarConcordeNoseAngle Simvar
// args contain optional index and/or unit
func SimVarConcordeNoseAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "CONCORDE NOSE ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRealismCrashWithOthers Simvar
// args contain optional index and/or unit
func SimVarRealismCrashWithOthers(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "REALISM CRASH WITH OTHERS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRealismCrashDetection Simvar
// args contain optional index and/or unit
func SimVarRealismCrashDetection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "REALISM CRASH DETECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarManualInstrumentLights Simvar
// args contain optional index and/or unit
func SimVarManualInstrumentLights(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "MANUAL INSTRUMENT LIGHTS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPitotIcePct Simvar
// args contain optional index and/or unit
func SimVarPitotIcePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "PITOT ICE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorY Simvar
// args contain optional index and/or unit
func SimVarSemibodyLoadfactorY(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "SEMIBODY LOADFACTOR Y",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorYdot Simvar
// args contain optional index and/or unit
func SimVarSemibodyLoadfactorYdot(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Per second")
	return SimVar{
		Index:    index,
		Name:     "SEMIBODY LOADFACTOR YDOT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRadInsSwitch Simvar
// args contain optional index and/or unit
func SimVarRadInsSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "RAD INS SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSimulatedRadius Simvar
// args contain optional index and/or unit
func SimVarSimulatedRadius(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "SIMULATED RADIUS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStructuralIcePct Simvar
// args contain optional index and/or unit
func SimVarStructuralIcePct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "STRUCTURAL ICE PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarArtificialGroundElevation Simvar
// args contain optional index and/or unit
func SimVarArtificialGroundElevation(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "ARTIFICIAL GROUND ELEVATION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSurfaceInfoValid Simvar
// args contain optional index and/or unit
func SimVarSurfaceInfoValid(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "SURFACE INFO VALID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSurfaceCondition Simvar
// args contain optional index and/or unit
func SimVarSurfaceCondition(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "SURFACE CONDITION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackAngle Simvar
// args contain optional index and/or unit
func SimVarPushbackAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackContactx Simvar
// args contain optional index and/or unit
func SimVarPushbackContactx(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTX",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackContacty Simvar
// args contain optional index and/or unit
func SimVarPushbackContacty(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackContactz Simvar
// args contain optional index and/or unit
func SimVarPushbackContactz(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK CONTACTZ",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPushbackWait Simvar
// args contain optional index and/or unit
func SimVarPushbackWait(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PUSHBACK WAIT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarYawStringAngle Simvar
// args contain optional index and/or unit
func SimVarYawStringAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "YAW STRING ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarYawStringPctExtended Simvar
// args contain optional index and/or unit
func SimVarYawStringPctExtended(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "YAW STRING PCT EXTENDED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarInductorCompassPercentDeviation Simvar
// args contain optional index and/or unit
func SimVarInductorCompassPercentDeviation(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "INDUCTOR COMPASS PERCENT DEVIATION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarInductorCompassHeadingRef Simvar
// args contain optional index and/or unit
func SimVarInductorCompassHeadingRef(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "INDUCTOR COMPASS HEADING REF",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAnemometerPctRpm Simvar
// args contain optional index and/or unit
func SimVarAnemometerPctRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "ANEMOMETER PCT RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarRotorRotationAngle Simvar
// args contain optional index and/or unit
func SimVarRotorRotationAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "ROTOR ROTATION ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDiskPitchAngle Simvar
// args contain optional index and/or unit
func SimVarDiskPitchAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "DISK PITCH ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDiskBankAngle Simvar
// args contain optional index and/or unit
func SimVarDiskBankAngle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "DISK BANK ANGLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDiskPitchPct Simvar
// args contain optional index and/or unit
func SimVarDiskPitchPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "DISK PITCH PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDiskBankPct Simvar
// args contain optional index and/or unit
func SimVarDiskBankPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "DISK BANK PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarDiskConingPct Simvar
// args contain optional index and/or unit
func SimVarDiskConingPct(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "DISK CONING PCT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavVorLlaf64 Simvar
func SimVarNavVorLlaf64(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "NAV VOR LLAF64",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarNavGsLlaf64 Simvar
func SimVarNavGsLlaf64(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "SIMCONNECT_DATA_LATLONALT")
	return SimVar{
		Index:    index,
		Name:     "NAV GS LLAF64",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStaticCgToGround Simvar
// args contain optional index and/or unit
func SimVarStaticCgToGround(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "STATIC CG TO GROUND",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarStaticPitch Simvar
// args contain optional index and/or unit
func SimVarStaticPitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Radians")
	return SimVar{
		Index:    index,
		Name:     "STATIC PITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCrashSequence Simvar
// args contain optional index and/or unit
func SimVarCrashSequence(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "CRASH SEQUENCE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCrashFlag Simvar
// args contain optional index and/or unit
func SimVarCrashFlag(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "CRASH FLAG",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTowReleaseHandle Simvar
// args contain optional index and/or unit
func SimVarTowReleaseHandle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "TOW RELEASE HANDLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTowConnection Simvar
// args contain optional index and/or unit
func SimVarTowConnection(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "TOW CONNECTION",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuPctRpm Simvar
// args contain optional index and/or unit
func SimVarApuPctRpm(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "APU PCT RPM",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuPctStarter Simvar
// args contain optional index and/or unit
func SimVarApuPctStarter(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Percent over 100")
	return SimVar{
		Index:    index,
		Name:     "APU PCT STARTER",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuVolts Simvar
// args contain optional index and/or unit
func SimVarApuVolts(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Volts")
	return SimVar{
		Index:    index,
		Name:     "APU VOLTS",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuGeneratorSwitch Simvar
// args contain optional index and/or unit
func SimVarApuGeneratorSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "APU GENERATOR SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuGeneratorActive Simvar
// args contain optional index and/or unit
func SimVarApuGeneratorActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "APU GENERATOR ACTIVE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarApuOnFireDetected Simvar
// args contain optional index and/or unit
func SimVarApuOnFireDetected(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "APU ON FIRE DETECTED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitude Simvar
// args contain optional index and/or unit
func SimVarPressurizationCabinAltitude(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeGoal Simvar
// args contain optional index and/or unit
func SimVarPressurizationCabinAltitudeGoal(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet")
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE GOAL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeRate Simvar
// args contain optional index and/or unit
func SimVarPressurizationCabinAltitudeRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Feet per second")
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION CABIN ALTITUDE RATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressurizationPressureDifferential Simvar
// args contain optional index and/or unit
func SimVarPressurizationPressureDifferential(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "foot pounds")
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION PRESSURE DIFFERENTIAL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarPressurizationDumpSwitch Simvar
// args contain optional index and/or unit
func SimVarPressurizationDumpSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "PRESSURIZATION DUMP SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFireBottleSwitch Simvar
// args contain optional index and/or unit
func SimVarFireBottleSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FIRE BOTTLE SWITCH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarFireBottleDischarged Simvar
// args contain optional index and/or unit
func SimVarFireBottleDischarged(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "FIRE BOTTLE DISCHARGED",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarCabinNoSmokingAlertSwitch Simvar
// args contain optional index and/or unit
func SimVarCabinNoSmokingAlertSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CABIN NO SMOKING ALERT SWITCH",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarCabinSeatbeltsAlertSwitch Simvar
// args contain optional index and/or unit
func SimVarCabinSeatbeltsAlertSwitch(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "CABIN SEATBELTS ALERT SWITCH",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarGpwsWarning Simvar
// args contain optional index and/or unit
func SimVarGpwsWarning(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPWS WARNING",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpwsSystemActive Simvar
// args contain optional index and/or unit
func SimVarGpwsSystemActive(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "GPWS SYSTEM ACTIVE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarIsLatitudeLongitudeFreezeOn Simvar
// args contain optional index and/or unit
func SimVarIsLatitudeLongitudeFreezeOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS LATITUDE LONGITUDE FREEZE ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsAltitudeFreezeOn Simvar
// args contain optional index and/or unit
func SimVarIsAltitudeFreezeOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS ALTITUDE FREEZE ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarIsAttitudeFreezeOn Simvar
// args contain optional index and/or unit
func SimVarIsAttitudeFreezeOn(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Bool")
	return SimVar{
		Index:    index,
		Name:     "IS ATTITUDE FREEZE ON",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAtcType Simvar
// args contain optional index and/or unit
func SimVarAtcType(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String64")
	return SimVar{
		Index:    index,
		Name:     "ATC TYPE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAtcModel Simvar
// args contain optional index and/or unit
func SimVarAtcModel(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String64")
	return SimVar{
		Index:    index,
		Name:     "ATC MODEL",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAtcId Simvar
// args contain optional index and/or unit
func SimVarAtcId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String64")
	return SimVar{
		Index:    index,
		Name:     "ATC ID",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAtcAirline Simvar
// args contain optional index and/or unit
func SimVarAtcAirline(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String64")
	return SimVar{
		Index:    index,
		Name:     "ATC AIRLINE",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarAtcFlightNumber Simvar
// args contain optional index and/or unit
func SimVarAtcFlightNumber(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String8")
	return SimVar{
		Index:    index,
		Name:     "ATC FLIGHT NUMBER",
		Unit:     unit,
		Settable: true,
	}
}

// SimVarTitle Simvar
// args contain optional index and/or unit
func SimVarTitle(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "TITLE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarHsiStationIdent Simvar
// args contain optional index and/or unit
func SimVarHsiStationIdent(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String8")
	return SimVar{
		Index:    index,
		Name:     "HSI STATION IDENT",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachAirportId Simvar
// args contain optional index and/or unit
func SimVarGpsApproachAirportId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH AIRPORT ID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachApproachId Simvar
// args contain optional index and/or unit
func SimVarGpsApproachApproachId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH APPROACH ID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarGpsApproachTransitionId Simvar
// args contain optional index and/or unit
func SimVarGpsApproachTransitionId(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "String")
	return SimVar{
		Index:    index,
		Name:     "GPS APPROACH TRANSITION ID",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarAbsoluteTime Simvar
// args contain optional index and/or unit
func SimVarAbsoluteTime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "ABSOLUTE TIME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluTime Simvar
// args contain optional index and/or unit
func SimVarZuluTime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "ZULU TIME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluDayOfWeek Simvar
// args contain optional index and/or unit
func SimVarZuluDayOfWeek(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF WEEK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluDayOfMonth Simvar
// args contain optional index and/or unit
func SimVarZuluDayOfMonth(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF MONTH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluMonthOfYear Simvar
// args contain optional index and/or unit
func SimVarZuluMonthOfYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ZULU MONTH OF YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluDayOfYear Simvar
// args contain optional index and/or unit
func SimVarZuluDayOfYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ZULU DAY OF YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarZuluYear Simvar
// args contain optional index and/or unit
func SimVarZuluYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "ZULU YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalTime Simvar
// args contain optional index and/or unit
func SimVarLocalTime(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "LOCAL TIME",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalDayOfWeek Simvar
// args contain optional index and/or unit
func SimVarLocalDayOfWeek(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF WEEK",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalDayOfMonth Simvar
// args contain optional index and/or unit
func SimVarLocalDayOfMonth(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF MONTH",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalMonthOfYear Simvar
// args contain optional index and/or unit
func SimVarLocalMonthOfYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "LOCAL MONTH OF YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalDayOfYear Simvar
// args contain optional index and/or unit
func SimVarLocalDayOfYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "LOCAL DAY OF YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarLocalYear Simvar
// args contain optional index and/or unit
func SimVarLocalYear(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "LOCAL YEAR",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTimeZoneOffset Simvar
// args contain optional index and/or unit
func SimVarTimeZoneOffset(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Seconds")
	return SimVar{
		Index:    index,
		Name:     "TIME ZONE OFFSET",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarTimeOfDay Simvar
// args contain optional index and/or unit
func SimVarTimeOfDay(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "TIME OF DAY",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarSimulationRate Simvar
// args contain optional index and/or unit
func SimVarSimulationRate(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Number")
	return SimVar{
		Index:    index,
		Name:     "SIMULATION RATE",
		Unit:     unit,
		Settable: false,
	}
}

// SimVarUnitOfMeasure Simvar
// args contain optional index and/or unit
func SimVarUnitOfMeasure(args ...interface{}) SimVar {
	index, unit := readArgs(args, 0, "Enum")
	return SimVar{
		Index:    index,
		Name:     "UNITS OF MEASURE",
		Unit:     unit,
		Settable: false,
	}
}
