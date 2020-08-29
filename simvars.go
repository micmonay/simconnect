package simconnect

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/sirupsen/logrus"
)

// SimVar is usued for all SimVar describtion
type SimVar struct {
	Name     string
	Units    string
	Settable bool
	data     []byte
}

func (s *SimVar) GetData() []byte {
	return s.data
}
func (s *SimVar) GetDatumType() uint32 {
	switch s.Units {
	case "Bool":
		return SIMCONNECT_DATATYPE_INT32
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

func (s *SimVar) GetDegrees() (float64, error) {
	f, err := s.GetFloat64()
	if err != nil {
		return 0, err
	}
	return f * 180 / math.Pi, nil
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
func SimVarAutopilotPitchHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT PITCH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarStructAmbientWind Simvar
func SimVarStructAmbientWind() SimVar {
	return SimVar{
		Name:     "STRUCT AMBIENT WIND",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarLaunchbarPosition Simvar
func SimVarLaunchbarPosition() SimVar {
	return SimVar{
		Name:     "LAUNCHBAR POSITION",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarNumberOfCatapults Simvar
func SimVarNumberOfCatapults() SimVar {
	return SimVar{
		Name:     "NUMBER OF CATAPULTS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHoldbackBarInstalled Simvar
func SimVarHoldbackBarInstalled() SimVar {
	return SimVar{
		Name:     "HOLDBACK BAR INSTALLED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarBlastShieldPosition Simvar
func SimVarBlastShieldPosition() SimVar {
	return SimVar{
		Name:     "BLAST SHIELD POSITION:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRecipEngDetonating Simvar
func SimVarRecipEngDetonating() SimVar {
	return SimVar{
		Name:     "RECIP ENG DETONATING:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRecipEngCylinderHealth Simvar
func SimVarRecipEngCylinderHealth() SimVar {
	return SimVar{
		Name:     "RECIP ENG CYLINDER HEALTH:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRecipEngNumCylinders Simvar
func SimVarRecipEngNumCylinders() SimVar {
	return SimVar{
		Name:     "RECIP ENG NUM CYLINDERS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipEngNumCylindersFailed Simvar
func SimVarRecipEngNumCylindersFailed() SimVar {
	return SimVar{
		Name:     "RECIP ENG NUM CYLINDERS FAILED",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipEngAntidetonationTankValve Simvar
func SimVarRecipEngAntidetonationTankValve() SimVar {
	return SimVar{
		Name:     "RECIP ENG ANTIDETONATION TANK VALVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankQuantity Simvar
func SimVarRecipEngAntidetonationTankQuantity() SimVar {
	return SimVar{
		Name:     "RECIP ENG ANTIDETONATION TANK QUANTITY:index",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarRecipEngAntidetonationTankMaxQuantity Simvar
func SimVarRecipEngAntidetonationTankMaxQuantity() SimVar {
	return SimVar{
		Name:     "RECIP ENG ANTIDETONATION TANK MAX QUANTITY:index",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarRecipEngNitrousTankValve Simvar
func SimVarRecipEngNitrousTankValve() SimVar {
	return SimVar{
		Name:     "RECIP ENG NITROUS TANK VALVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankQuantity Simvar
func SimVarRecipEngNitrousTankQuantity() SimVar {
	return SimVar{
		Name:     "RECIP ENG NITROUS TANK QUANTITY:index",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarRecipEngNitrousTankMaxQuantity Simvar
func SimVarRecipEngNitrousTankMaxQuantity() SimVar {
	return SimVar{
		Name:     "RECIP ENG NITROUS TANK MAX QUANTITY:index",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarPayloadStationObject Simvar
func SimVarPayloadStationObject() SimVar {
	return SimVar{
		Name:     "PAYLOAD STATION OBJECT:index",
		Units:    "String",
		Settable: true,
	}
}

// SimVarPayloadStationNumSimobjects Simvar
func SimVarPayloadStationNumSimobjects() SimVar {
	return SimVar{
		Name:     "PAYLOAD STATION NUM SIMOBJECTS:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarSlingObjectAttached Simvar
func SimVarSlingObjectAttached() SimVar {
	return SimVar{
		Name:     "SLING OBJECT ATTACHED:index",
		Units:    "Bool/String",
		Settable: false,
	}
}

// SimVarSlingCableBroken Simvar
func SimVarSlingCableBroken() SimVar {
	return SimVar{
		Name:     "SLING CABLE BROKEN:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSlingCableExtendedLength Simvar
func SimVarSlingCableExtendedLength() SimVar {
	return SimVar{
		Name:     "SLING CABLE EXTENDED LENGTH:index",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarSlingActivePayloadStation Simvar
func SimVarSlingActivePayloadStation() SimVar {
	return SimVar{
		Name:     "SLING ACTIVE PAYLOAD STATION:index",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarSlingHoistPercentDeployed Simvar
func SimVarSlingHoistPercentDeployed() SimVar {
	return SimVar{
		Name:     "SLING HOIST PERCENT DEPLOYED:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarSlingHookInPickupMode Simvar
func SimVarSlingHookInPickupMode() SimVar {
	return SimVar{
		Name:     "SLING HOOK IN PICKUP MODE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAttachedToSling Simvar
func SimVarIsAttachedToSling() SimVar {
	return SimVar{
		Name:     "IS ATTACHED TO SLING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAlternateStaticSourceOpen Simvar
func SimVarAlternateStaticSourceOpen() SimVar {
	return SimVar{
		Name:     "ALTERNATE STATIC SOURCE OPEN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAileronTrimPct Simvar
func SimVarAileronTrimPct() SimVar {
	return SimVar{
		Name:     "AILERON TRIM PCT",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: true,
	}
}

// SimVarRudderTrimPct Simvar
func SimVarRudderTrimPct() SimVar {
	return SimVar{
		Name:     "RUDDER TRIM PCT",
		Units:    "Percent over 100",
		Settable: true,
	}
}

// SimVarLightOnStates Simvar
func SimVarLightOnStates() SimVar {
	return SimVar{
		Name:     "LIGHT ON STATES",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarLightStates Simvar
func SimVarLightStates() SimVar {
	return SimVar{
		Name:     "LIGHT STATES",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarLandingLightPbh Simvar
func SimVarLandingLightPbh() SimVar {
	return SimVar{
		Name:     "LANDING LIGHT PBH",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarLightTaxiOn Simvar
func SimVarLightTaxiOn() SimVar {
	return SimVar{
		Name:     "LIGHT TAXI ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightStrobeOn Simvar
func SimVarLightStrobeOn() SimVar {
	return SimVar{
		Name:     "LIGHT STROBE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightPanelOn Simvar
func SimVarLightPanelOn() SimVar {
	return SimVar{
		Name:     "LIGHT PANEL ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightRecognitionOn Simvar
func SimVarLightRecognitionOn() SimVar {
	return SimVar{
		Name:     "LIGHT RECOGNITION ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightWingOn Simvar
func SimVarLightWingOn() SimVar {
	return SimVar{
		Name:     "LIGHT WING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLogoOn Simvar
func SimVarLightLogoOn() SimVar {
	return SimVar{
		Name:     "LIGHT LOGO ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightCabinOn Simvar
func SimVarLightCabinOn() SimVar {
	return SimVar{
		Name:     "LIGHT CABIN ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightHeadOn Simvar
func SimVarLightHeadOn() SimVar {
	return SimVar{
		Name:     "LIGHT HEAD ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBrakeOn Simvar
func SimVarLightBrakeOn() SimVar {
	return SimVar{
		Name:     "LIGHT BRAKE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightNavOn Simvar
func SimVarLightNavOn() SimVar {
	return SimVar{
		Name:     "LIGHT NAV ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBeaconOn Simvar
func SimVarLightBeaconOn() SimVar {
	return SimVar{
		Name:     "LIGHT BEACON ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLandingOn Simvar
func SimVarLightLandingOn() SimVar {
	return SimVar{
		Name:     "LIGHT LANDING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAiDesiredSpeed Simvar
func SimVarAiDesiredSpeed() SimVar {
	return SimVar{
		Name:     "AI DESIRED SPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiWaypointList Simvar
func SimVarAiWaypointList() SimVar {
	return SimVar{
		Name:     "AI WAYPOINT LIST",
		Units:    "SIMCONNECT_DATA_WAYPOINT",
		Settable: true,
	}
}

// SimVarAiCurrentWaypoint Simvar
func SimVarAiCurrentWaypoint() SimVar {
	return SimVar{
		Name:     "AI CURRENT WAYPOINT",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarAiDesiredHeading Simvar
func SimVarAiDesiredHeading() SimVar {
	return SimVar{
		Name:     "AI DESIRED HEADING",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarAiGroundturntime Simvar
func SimVarAiGroundturntime() SimVar {
	return SimVar{
		Name:     "AI GROUNDTURNTIME",
		Units:    "Seconds",
		Settable: true,
	}
}

// SimVarAiGroundcruisespeed Simvar
func SimVarAiGroundcruisespeed() SimVar {
	return SimVar{
		Name:     "AI GROUNDCRUISESPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiGroundturnspeed Simvar
func SimVarAiGroundturnspeed() SimVar {
	return SimVar{
		Name:     "AI GROUNDTURNSPEED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAiTrafficIsifr Simvar
func SimVarAiTrafficIsifr() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC ISIFR",
		Units:    "Boolean",
		Settable: false,
	}
}

// SimVarAiTrafficState Simvar
func SimVarAiTrafficState() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC STATE",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficCurrentAirport Simvar
func SimVarAiTrafficCurrentAirport() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC CURRENT AIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficAssignedRunway Simvar
func SimVarAiTrafficAssignedRunway() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC ASSIGNED RUNWAY",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficAssignedParking Simvar
func SimVarAiTrafficAssignedParking() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC ASSIGNED PARKING",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficFromairport Simvar
func SimVarAiTrafficFromairport() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC FROMAIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficToairport Simvar
func SimVarAiTrafficToairport() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC TOAIRPORT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAiTrafficEtd Simvar
func SimVarAiTrafficEtd() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC ETD",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarAiTrafficEta Simvar
func SimVarAiTrafficEta() SimVar {
	return SimVar{
		Name:     "AI TRAFFIC ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarDroppableObjectsType Simvar
func SimVarDroppableObjectsType() SimVar {
	return SimVar{
		Name:     "DROPPABLE OBJECTS TYPE:index",
		Units:    "String",
		Settable: true,
	}
}

// SimVarDroppableObjectsCount Simvar
func SimVarDroppableObjectsCount() SimVar {
	return SimVar{
		Name:     "DROPPABLE OBJECTS COUNT:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarWingFlexPct Simvar
func SimVarWingFlexPct() SimVar {
	return SimVar{
		Name:     "WING FLEX PCT:index",
		Units:    "Percent over 100",
		Settable: true,
	}
}

// SimVarApplyHeatToSystems Simvar
func SimVarApplyHeatToSystems() SimVar {
	return SimVar{
		Name:     "APPLY HEAT TO SYSTEMS",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAdfLatlonalt Simvar
func SimVarAdfLatlonalt() SimVar {
	return SimVar{
		Name:     "ADF LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavVorLatlonalt Simvar
func SimVarNavVorLatlonalt() SimVar {
	return SimVar{
		Name:     "NAV VOR LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavGsLatlonalt Simvar
func SimVarNavGsLatlonalt() SimVar {
	return SimVar{
		Name:     "NAV GS LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavDmeLatlonalt Simvar
func SimVarNavDmeLatlonalt() SimVar {
	return SimVar{
		Name:     "NAV DME LATLONALT:index",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarInnerMarkerLatlonalt Simvar
func SimVarInnerMarkerLatlonalt() SimVar {
	return SimVar{
		Name:     "INNER MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarMiddleMarkerLatlonalt Simvar
func SimVarMiddleMarkerLatlonalt() SimVar {
	return SimVar{
		Name:     "MIDDLE MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarOuterMarkerLatlonalt Simvar
func SimVarOuterMarkerLatlonalt() SimVar {
	return SimVar{
		Name:     "OUTER MARKER LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructLatlonalt Simvar
func SimVarStructLatlonalt() SimVar {
	return SimVar{
		Name:     "STRUCT LATLONALT",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructLatlonaltpbh Simvar
func SimVarStructLatlonaltpbh() SimVar {
	return SimVar{
		Name:     "STRUCT LATLONALTPBH",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStructSurfaceRelativeVelocity Simvar
func SimVarStructSurfaceRelativeVelocity() SimVar {
	return SimVar{
		Name:     "STRUCT SURFACE RELATIVE VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldvelocity Simvar
func SimVarStructWorldvelocity() SimVar {
	return SimVar{
		Name:     "STRUCT WORLDVELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldRotationVelocity Simvar
func SimVarStructWorldRotationVelocity() SimVar {
	return SimVar{
		Name:     "STRUCT WORLD ROTATION VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructBodyVelocity Simvar
func SimVarStructBodyVelocity() SimVar {
	return SimVar{
		Name:     "STRUCT BODY VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructBodyRotationVelocity Simvar
func SimVarStructBodyRotationVelocity() SimVar {
	return SimVar{
		Name:     "STRUCT BODY ROTATION VELOCITY",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructWorldAcceleration Simvar
func SimVarStructWorldAcceleration() SimVar {
	return SimVar{
		Name:     "STRUCT WORLD ACCELERATION",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEnginePosition Simvar
func SimVarStructEnginePosition() SimVar {
	return SimVar{
		Name:     "STRUCT ENGINE POSITION:index",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEyepointDynamicAngle Simvar
func SimVarStructEyepointDynamicAngle() SimVar {
	return SimVar{
		Name:     "STRUCT EYEPOINT DYNAMIC ANGLE",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarStructEyepointDynamicOffset Simvar
func SimVarStructEyepointDynamicOffset() SimVar {
	return SimVar{
		Name:     "STRUCT EYEPOINT DYNAMIC OFFSET",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarEyepointPosition Simvar
func SimVarEyepointPosition() SimVar {
	return SimVar{
		Name:     "EYEPOINT POSITION",
		Units:    "SIMCONNECT_DATA_XYZ",
		Settable: false,
	}
}

// SimVarFlyByWireElacSwitch Simvar
func SimVarFlyByWireElacSwitch() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE ELAC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireFacSwitch Simvar
func SimVarFlyByWireFacSwitch() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE FAC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireSecSwitch Simvar
func SimVarFlyByWireSecSwitch() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE SEC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireElacFailed Simvar
func SimVarFlyByWireElacFailed() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE ELAC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireFacFailed Simvar
func SimVarFlyByWireFacFailed() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE FAC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlyByWireSecFailed Simvar
func SimVarFlyByWireSecFailed() SimVar {
	return SimVar{
		Name:     "FLY BY WIRE SEC FAILED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNumberOfEngines Simvar
func SimVarNumberOfEngines() SimVar {
	return SimVar{
		Name:     "NUMBER OF ENGINES",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngineControlSelect Simvar
func SimVarEngineControlSelect() SimVar {
	return SimVar{
		Name:     "ENGINE CONTROL SELECT",
		Units:    "Mask",
		Settable: true,
	}
}

// SimVarThrottleLowerLimit Simvar
func SimVarThrottleLowerLimit() SimVar {
	return SimVar{
		Name:     "THROTTLE LOWER LIMIT",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngineType Simvar
func SimVarEngineType() SimVar {
	return SimVar{
		Name:     "ENGINE TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarMasterIgnitionSwitch Simvar
func SimVarMasterIgnitionSwitch() SimVar {
	return SimVar{
		Name:     "MASTER IGNITION SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngCombustion Simvar
func SimVarGeneralEngCombustion() SimVar {
	return SimVar{
		Name:     "GENERAL ENG COMBUSTION:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGeneralEngMasterAlternator Simvar
func SimVarGeneralEngMasterAlternator() SimVar {
	return SimVar{
		Name:     "GENERAL ENG MASTER ALTERNATOR:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpSwitch Simvar
func SimVarGeneralEngFuelPumpSwitch() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FUEL PUMP SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPumpOn Simvar
func SimVarGeneralEngFuelPumpOn() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FUEL PUMP ON:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngRpm Simvar
func SimVarGeneralEngRpm() SimVar {
	return SimVar{
		Name:     "GENERAL ENG RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngPctMaxRpm Simvar
func SimVarGeneralEngPctMaxRpm() SimVar {
	return SimVar{
		Name:     "GENERAL ENG PCT MAX RPM:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngMaxReachedRpm Simvar
func SimVarGeneralEngMaxReachedRpm() SimVar {
	return SimVar{
		Name:     "GENERAL ENG MAX REACHED RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngThrottleLeverPosition Simvar
func SimVarGeneralEngThrottleLeverPosition() SimVar {
	return SimVar{
		Name:     "GENERAL ENG THROTTLE LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngMixtureLeverPosition Simvar
func SimVarGeneralEngMixtureLeverPosition() SimVar {
	return SimVar{
		Name:     "GENERAL ENG MIXTURE LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngPropellerLeverPosition Simvar
func SimVarGeneralEngPropellerLeverPosition() SimVar {
	return SimVar{
		Name:     "GENERAL ENG PROPELLER LEVER POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarGeneralEngStarter Simvar
func SimVarGeneralEngStarter() SimVar {
	return SimVar{
		Name:     "GENERAL ENG STARTER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngExhaustGasTemperature Simvar
func SimVarGeneralEngExhaustGasTemperature() SimVar {
	return SimVar{
		Name:     "GENERAL ENG EXHAUST GAS TEMPERATURE:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarGeneralEngOilPressure Simvar
func SimVarGeneralEngOilPressure() SimVar {
	return SimVar{
		Name:     "GENERAL ENG OIL PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarGeneralEngOilLeakedPercent Simvar
func SimVarGeneralEngOilLeakedPercent() SimVar {
	return SimVar{
		Name:     "GENERAL ENG OIL LEAKED PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngCombustionSoundPercent Simvar
func SimVarGeneralEngCombustionSoundPercent() SimVar {
	return SimVar{
		Name:     "GENERAL ENG COMBUSTION SOUND PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngDamagePercent Simvar
func SimVarGeneralEngDamagePercent() SimVar {
	return SimVar{
		Name:     "GENERAL ENG DAMAGE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarGeneralEngOilTemperature Simvar
func SimVarGeneralEngOilTemperature() SimVar {
	return SimVar{
		Name:     "GENERAL ENG OIL TEMPERATURE:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarGeneralEngFailed Simvar
func SimVarGeneralEngFailed() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FAILED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorSwitch Simvar
func SimVarGeneralEngGeneratorSwitch() SimVar {
	return SimVar{
		Name:     "GENERAL ENG GENERATOR SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngGeneratorActive Simvar
func SimVarGeneralEngGeneratorActive() SimVar {
	return SimVar{
		Name:     "GENERAL ENG GENERATOR ACTIVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGeneralEngAntiIcePosition Simvar
func SimVarGeneralEngAntiIcePosition() SimVar {
	return SimVar{
		Name:     "GENERAL ENG ANTI ICE POSITION:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelValve Simvar
func SimVarGeneralEngFuelValve() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FUEL VALVE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelPressure Simvar
func SimVarGeneralEngFuelPressure() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FUEL PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarGeneralEngElapsedTime Simvar
func SimVarGeneralEngElapsedTime() SimVar {
	return SimVar{
		Name:     "GENERAL ENG ELAPSED TIME:index",
		Units:    "Hours",
		Settable: false,
	}
}

// SimVarRecipEngCowlFlapPosition Simvar
func SimVarRecipEngCowlFlapPosition() SimVar {
	return SimVar{
		Name:     "RECIP ENG COWL FLAP POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngPrimer Simvar
func SimVarRecipEngPrimer() SimVar {
	return SimVar{
		Name:     "RECIP ENG PRIMER:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngManifoldPressure Simvar
func SimVarRecipEngManifoldPressure() SimVar {
	return SimVar{
		Name:     "RECIP ENG MANIFOLD PRESSURE:index",
		Units:    "Psi",
		Settable: true,
	}
}

// SimVarRecipEngAlternateAirPosition Simvar
func SimVarRecipEngAlternateAirPosition() SimVar {
	return SimVar{
		Name:     "RECIP ENG ALTERNATE AIR POSITION:index",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRecipEngCoolantReservoirPercent Simvar
func SimVarRecipEngCoolantReservoirPercent() SimVar {
	return SimVar{
		Name:     "RECIP ENG COOLANT RESERVOIR PERCENT:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngLeftMagneto Simvar
func SimVarRecipEngLeftMagneto() SimVar {
	return SimVar{
		Name:     "RECIP ENG LEFT MAGNETO:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngRightMagneto Simvar
func SimVarRecipEngRightMagneto() SimVar {
	return SimVar{
		Name:     "RECIP ENG RIGHT MAGNETO:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngBrakePower Simvar
func SimVarRecipEngBrakePower() SimVar {
	return SimVar{
		Name:     "RECIP ENG BRAKE POWER:index",
		Units:    "ft lb per second",
		Settable: true,
	}
}

// SimVarRecipEngStarterTorque Simvar
func SimVarRecipEngStarterTorque() SimVar {
	return SimVar{
		Name:     "RECIP ENG STARTER TORQUE:index",
		Units:    "Foot pound",
		Settable: true,
	}
}

// SimVarRecipEngTurbochargerFailed Simvar
func SimVarRecipEngTurbochargerFailed() SimVar {
	return SimVar{
		Name:     "RECIP ENG TURBOCHARGER FAILED:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostActive Simvar
func SimVarRecipEngEmergencyBoostActive() SimVar {
	return SimVar{
		Name:     "RECIP ENG EMERGENCY BOOST ACTIVE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngEmergencyBoostElapsedTime Simvar
func SimVarRecipEngEmergencyBoostElapsedTime() SimVar {
	return SimVar{
		Name:     "RECIP ENG EMERGENCY BOOST ELAPSED TIME:index",
		Units:    "Hours",
		Settable: true,
	}
}

// SimVarRecipEngWastegatePosition Simvar
func SimVarRecipEngWastegatePosition() SimVar {
	return SimVar{
		Name:     "RECIP ENG WASTEGATE POSITION:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarRecipEngTurbineInletTemperature Simvar
func SimVarRecipEngTurbineInletTemperature() SimVar {
	return SimVar{
		Name:     "RECIP ENG TURBINE INLET TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngCylinderHeadTemperature Simvar
func SimVarRecipEngCylinderHeadTemperature() SimVar {
	return SimVar{
		Name:     "RECIP ENG CYLINDER HEAD TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngRadiatorTemperature Simvar
func SimVarRecipEngRadiatorTemperature() SimVar {
	return SimVar{
		Name:     "RECIP ENG RADIATOR TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipEngFuelAvailable Simvar
func SimVarRecipEngFuelAvailable() SimVar {
	return SimVar{
		Name:     "RECIP ENG FUEL AVAILABLE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRecipEngFuelFlow Simvar
func SimVarRecipEngFuelFlow() SimVar {
	return SimVar{
		Name:     "RECIP ENG FUEL FLOW:index",
		Units:    "Pounds per hour",
		Settable: true,
	}
}

// SimVarRecipEngFuelTankSelector Simvar
func SimVarRecipEngFuelTankSelector() SimVar {
	return SimVar{
		Name:     "RECIP ENG FUEL TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarRecipEngFuelTanksUsed Simvar
func SimVarRecipEngFuelTanksUsed() SimVar {
	return SimVar{
		Name:     "RECIP ENG FUEL TANKS USED:index",
		Units:    "Mask",
		Settable: true,
	}
}

// SimVarRecipEngFuelNumberTanksUsed Simvar
func SimVarRecipEngFuelNumberTanksUsed() SimVar {
	return SimVar{
		Name:     "RECIP ENG FUEL NUMBER TANKS USED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarRecipCarburetorTemperature Simvar
func SimVarRecipCarburetorTemperature() SimVar {
	return SimVar{
		Name:     "RECIP CARBURETOR TEMPERATURE:index",
		Units:    "Celsius",
		Settable: true,
	}
}

// SimVarRecipMixtureRatio Simvar
func SimVarRecipMixtureRatio() SimVar {
	return SimVar{
		Name:     "RECIP MIXTURE RATIO:index",
		Units:    "Ratio",
		Settable: true,
	}
}

// SimVarTurbEngN1 Simvar
func SimVarTurbEngN1() SimVar {
	return SimVar{
		Name:     "TURB ENG N1:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngN2 Simvar
func SimVarTurbEngN2() SimVar {
	return SimVar{
		Name:     "TURB ENG N2:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN1 Simvar
func SimVarTurbEngCorrectedN1() SimVar {
	return SimVar{
		Name:     "TURB ENG CORRECTED N1:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedN2 Simvar
func SimVarTurbEngCorrectedN2() SimVar {
	return SimVar{
		Name:     "TURB ENG CORRECTED N2:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngCorrectedFf Simvar
func SimVarTurbEngCorrectedFf() SimVar {
	return SimVar{
		Name:     "TURB ENG CORRECTED FF:index",
		Units:    "Pounds per hour",
		Settable: true,
	}
}

// SimVarTurbEngMaxTorquePercent Simvar
func SimVarTurbEngMaxTorquePercent() SimVar {
	return SimVar{
		Name:     "TURB ENG MAX TORQUE PERCENT:index",
		Units:    "Percent",
		Settable: true,
	}
}

// SimVarTurbEngPressureRatio Simvar
func SimVarTurbEngPressureRatio() SimVar {
	return SimVar{
		Name:     "TURB ENG PRESSURE RATIO:index",
		Units:    "Ratio",
		Settable: true,
	}
}

// SimVarTurbEngItt Simvar
func SimVarTurbEngItt() SimVar {
	return SimVar{
		Name:     "TURB ENG ITT:index",
		Units:    "Rankine",
		Settable: true,
	}
}

// SimVarTurbEngAfterburner Simvar
func SimVarTurbEngAfterburner() SimVar {
	return SimVar{
		Name:     "TURB ENG AFTERBURNER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngJetThrust Simvar
func SimVarTurbEngJetThrust() SimVar {
	return SimVar{
		Name:     "TURB ENG JET THRUST:index",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarTurbEngBleedAir Simvar
func SimVarTurbEngBleedAir() SimVar {
	return SimVar{
		Name:     "TURB ENG BLEED AIR:index",
		Units:    "Psi",
		Settable: false,
	}
}

// SimVarTurbEngTankSelector Simvar
func SimVarTurbEngTankSelector() SimVar {
	return SimVar{
		Name:     "TURB ENG TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarTurbEngTanksUsed Simvar
func SimVarTurbEngTanksUsed() SimVar {
	return SimVar{
		Name:     "TURB ENG TANKS USED:index",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarTurbEngNumTanksUsed Simvar
func SimVarTurbEngNumTanksUsed() SimVar {
	return SimVar{
		Name:     "TURB ENG NUM TANKS USED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTurbEngFuelFlowPph Simvar
func SimVarTurbEngFuelFlowPph() SimVar {
	return SimVar{
		Name:     "TURB ENG FUEL FLOW PPH:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarTurbEngFuelAvailable Simvar
func SimVarTurbEngFuelAvailable() SimVar {
	return SimVar{
		Name:     "TURB ENG FUEL AVAILABLE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngReverseNozzlePercent Simvar
func SimVarTurbEngReverseNozzlePercent() SimVar {
	return SimVar{
		Name:     "TURB ENG REVERSE NOZZLE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarTurbEngVibration Simvar
func SimVarTurbEngVibration() SimVar {
	return SimVar{
		Name:     "TURB ENG VIBRATION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngFailed Simvar
func SimVarEngFailed() SimVar {
	return SimVar{
		Name:     "ENG FAILED:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngRpmAnimationPercent Simvar
func SimVarEngRpmAnimationPercent() SimVar {
	return SimVar{
		Name:     "ENG RPM ANIMATION PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngOnFire Simvar
func SimVarEngOnFire() SimVar {
	return SimVar{
		Name:     "ENG ON FIRE:index",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarEngFuelFlowBugPosition Simvar
func SimVarEngFuelFlowBugPosition() SimVar {
	return SimVar{
		Name:     "ENG FUEL FLOW BUG POSITION:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarPropRpm Simvar
func SimVarPropRpm() SimVar {
	return SimVar{
		Name:     "PROP RPM:index",
		Units:    "Rpm",
		Settable: true,
	}
}

// SimVarPropMaxRpmPercent Simvar
func SimVarPropMaxRpmPercent() SimVar {
	return SimVar{
		Name:     "PROP MAX RPM PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarPropThrust Simvar
func SimVarPropThrust() SimVar {
	return SimVar{
		Name:     "PROP THRUST:index",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarPropBeta Simvar
func SimVarPropBeta() SimVar {
	return SimVar{
		Name:     "PROP BETA:index",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropFeatheringInhibit Simvar
func SimVarPropFeatheringInhibit() SimVar {
	return SimVar{
		Name:     "PROP FEATHERING INHIBIT:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropFeathered Simvar
func SimVarPropFeathered() SimVar {
	return SimVar{
		Name:     "PROP FEATHERED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropSyncDeltaLever Simvar
func SimVarPropSyncDeltaLever() SimVar {
	return SimVar{
		Name:     "PROP SYNC DELTA LEVER:index",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarPropAutoFeatherArmed Simvar
func SimVarPropAutoFeatherArmed() SimVar {
	return SimVar{
		Name:     "PROP AUTO FEATHER ARMED:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropFeatherSwitch Simvar
func SimVarPropFeatherSwitch() SimVar {
	return SimVar{
		Name:     "PROP FEATHER SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPanelAutoFeatherSwitch Simvar
func SimVarPanelAutoFeatherSwitch() SimVar {
	return SimVar{
		Name:     "PANEL AUTO FEATHER SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropSyncActive Simvar
func SimVarPropSyncActive() SimVar {
	return SimVar{
		Name:     "PROP SYNC ACTIVE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropDeiceSwitch Simvar
func SimVarPropDeiceSwitch() SimVar {
	return SimVar{
		Name:     "PROP DEICE SWITCH:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngCombustion Simvar
func SimVarEngCombustion() SimVar {
	return SimVar{
		Name:     "ENG COMBUSTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngN1Rpm Simvar
func SimVarEngN1Rpm() SimVar {
	return SimVar{
		Name:     "ENG N1 RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarEngN2Rpm Simvar
func SimVarEngN2Rpm() SimVar {
	return SimVar{
		Name:     "ENG N2 RPM:index",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarEngFuelFlowPph Simvar
func SimVarEngFuelFlowPph() SimVar {
	return SimVar{
		Name:     "ENG FUEL FLOW PPH:index",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarEngTorque Simvar
func SimVarEngTorque() SimVar {
	return SimVar{
		Name:     "ENG TORQUE:index",
		Units:    "Foot pounds",
		Settable: false,
	}
}

// SimVarEngAntiIce Simvar
func SimVarEngAntiIce() SimVar {
	return SimVar{
		Name:     "ENG ANTI ICE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngPressureRatio Simvar
func SimVarEngPressureRatio() SimVar {
	return SimVar{
		Name:     "ENG PRESSURE RATIO:index",
		Units:    "Ratio (0-16384)",
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperature Simvar
func SimVarEngExhaustGasTemperature() SimVar {
	return SimVar{
		Name:     "ENG EXHAUST GAS TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngExhaustGasTemperatureGes Simvar
func SimVarEngExhaustGasTemperatureGes() SimVar {
	return SimVar{
		Name:     "ENG EXHAUST GAS TEMPERATURE GES:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngCylinderHeadTemperature Simvar
func SimVarEngCylinderHeadTemperature() SimVar {
	return SimVar{
		Name:     "ENG CYLINDER HEAD TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngOilTemperature Simvar
func SimVarEngOilTemperature() SimVar {
	return SimVar{
		Name:     "ENG OIL TEMPERATURE:index",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarEngOilPressure Simvar
func SimVarEngOilPressure() SimVar {
	return SimVar{
		Name:     "ENG OIL PRESSURE:index",
		Units:    "pound-force per square inch",
		Settable: false,
	}
}

// SimVarEngOilQuantity Simvar
func SimVarEngOilQuantity() SimVar {
	return SimVar{
		Name:     "ENG OIL QUANTITY:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngHydraulicPressure Simvar
func SimVarEngHydraulicPressure() SimVar {
	return SimVar{
		Name:     "ENG HYDRAULIC PRESSURE:index",
		Units:    "pound-force per square inch",
		Settable: false,
	}
}

// SimVarEngHydraulicQuantity Simvar
func SimVarEngHydraulicQuantity() SimVar {
	return SimVar{
		Name:     "ENG HYDRAULIC QUANTITY:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarEngManifoldPressure Simvar
func SimVarEngManifoldPressure() SimVar {
	return SimVar{
		Name:     "ENG MANIFOLD PRESSURE:index",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarEngVibration Simvar
func SimVarEngVibration() SimVar {
	return SimVar{
		Name:     "ENG VIBRATION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngRpmScaler Simvar
func SimVarEngRpmScaler() SimVar {
	return SimVar{
		Name:     "ENG RPM SCALER:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarEngTurbineTemperature Simvar
func SimVarEngTurbineTemperature() SimVar {
	return SimVar{
		Name:     "ENG TURBINE TEMPERATURE:index",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarEngTorquePercent Simvar
func SimVarEngTorquePercent() SimVar {
	return SimVar{
		Name:     "ENG TORQUE PERCENT:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngFuelPressure Simvar
func SimVarEngFuelPressure() SimVar {
	return SimVar{
		Name:     "ENG FUEL PRESSURE:index",
		Units:    "PSI",
		Settable: false,
	}
}

// SimVarEngElectricalLoad Simvar
func SimVarEngElectricalLoad() SimVar {
	return SimVar{
		Name:     "ENG ELECTRICAL LOAD:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngTransmissionPressure Simvar
func SimVarEngTransmissionPressure() SimVar {
	return SimVar{
		Name:     "ENG TRANSMISSION PRESSURE:index",
		Units:    "PSI",
		Settable: false,
	}
}

// SimVarEngTransmissionTemperature Simvar
func SimVarEngTransmissionTemperature() SimVar {
	return SimVar{
		Name:     "ENG TRANSMISSION TEMPERATURE:index",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarEngRotorRpm Simvar
func SimVarEngRotorRpm() SimVar {
	return SimVar{
		Name:     "ENG ROTOR RPM:index",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarEngMaxRpm Simvar
func SimVarEngMaxRpm() SimVar {
	return SimVar{
		Name:     "ENG MAX RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarGeneralEngStarterActive Simvar
func SimVarGeneralEngStarterActive() SimVar {
	return SimVar{
		Name:     "GENERAL ENG STARTER ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGeneralEngFuelUsedSinceStart Simvar
func SimVarGeneralEngFuelUsedSinceStart() SimVar {
	return SimVar{
		Name:     "GENERAL ENG FUEL USED SINCE START",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarTurbEngPrimaryNozzlePercent Simvar
func SimVarTurbEngPrimaryNozzlePercent() SimVar {
	return SimVar{
		Name:     "TURB ENG PRIMARY NOZZLE PERCENT:index",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarTurbEngIgnitionSwitch Simvar
func SimVarTurbEngIgnitionSwitch() SimVar {
	return SimVar{
		Name:     "TURB ENG IGNITION SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTurbEngMasterStarterSwitch Simvar
func SimVarTurbEngMasterStarterSwitch() SimVar {
	return SimVar{
		Name:     "TURB ENG MASTER STARTER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFuelTankCenterLevel Simvar
func SimVarFuelTankCenterLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenter2Level Simvar
func SimVarFuelTankCenter2Level() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER2 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenter3Level Simvar
func SimVarFuelTankCenter3Level() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER3 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftMainLevel Simvar
func SimVarFuelTankLeftMainLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT MAIN LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxLevel Simvar
func SimVarFuelTankLeftAuxLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT AUX LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankLeftTipLevel Simvar
func SimVarFuelTankLeftTipLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT TIP LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightMainLevel Simvar
func SimVarFuelTankRightMainLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT MAIN LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightAuxLevel Simvar
func SimVarFuelTankRightAuxLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT AUX LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankRightTipLevel Simvar
func SimVarFuelTankRightTipLevel() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT TIP LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankExternal1Level Simvar
func SimVarFuelTankExternal1Level() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL1 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankExternal2Level Simvar
func SimVarFuelTankExternal2Level() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL2 LEVEL",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFuelTankCenterCapacity Simvar
func SimVarFuelTankCenterCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenter2Capacity Simvar
func SimVarFuelTankCenter2Capacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER2 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenter3Capacity Simvar
func SimVarFuelTankCenter3Capacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER3 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftMainCapacity Simvar
func SimVarFuelTankLeftMainCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT MAIN CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftAuxCapacity Simvar
func SimVarFuelTankLeftAuxCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT AUX CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankLeftTipCapacity Simvar
func SimVarFuelTankLeftTipCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT TIP CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightMainCapacity Simvar
func SimVarFuelTankRightMainCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT MAIN CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightAuxCapacity Simvar
func SimVarFuelTankRightAuxCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT AUX CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankRightTipCapacity Simvar
func SimVarFuelTankRightTipCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT TIP CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankExternal1Capacity Simvar
func SimVarFuelTankExternal1Capacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL1 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankExternal2Capacity Simvar
func SimVarFuelTankExternal2Capacity() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL2 CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelLeftCapacity Simvar
func SimVarFuelLeftCapacity() SimVar {
	return SimVar{
		Name:     "FUEL LEFT CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelRightCapacity Simvar
func SimVarFuelRightCapacity() SimVar {
	return SimVar{
		Name:     "FUEL RIGHT CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTankCenterQuantity Simvar
func SimVarFuelTankCenterQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankCenter2Quantity Simvar
func SimVarFuelTankCenter2Quantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER2 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankCenter3Quantity Simvar
func SimVarFuelTankCenter3Quantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK CENTER3 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftMainQuantity Simvar
func SimVarFuelTankLeftMainQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT MAIN QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftAuxQuantity Simvar
func SimVarFuelTankLeftAuxQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT AUX QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankLeftTipQuantity Simvar
func SimVarFuelTankLeftTipQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK LEFT TIP QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightMainQuantity Simvar
func SimVarFuelTankRightMainQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT MAIN QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightAuxQuantity Simvar
func SimVarFuelTankRightAuxQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT AUX QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankRightTipQuantity Simvar
func SimVarFuelTankRightTipQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK RIGHT TIP QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankExternal1Quantity Simvar
func SimVarFuelTankExternal1Quantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL1 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelTankExternal2Quantity Simvar
func SimVarFuelTankExternal2Quantity() SimVar {
	return SimVar{
		Name:     "FUEL TANK EXTERNAL2 QUANTITY",
		Units:    "Gallons",
		Settable: true,
	}
}

// SimVarFuelLeftQuantity Simvar
func SimVarFuelLeftQuantity() SimVar {
	return SimVar{
		Name:     "FUEL LEFT QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelRightQuantity Simvar
func SimVarFuelRightQuantity() SimVar {
	return SimVar{
		Name:     "FUEL RIGHT QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTotalQuantity Simvar
func SimVarFuelTotalQuantity() SimVar {
	return SimVar{
		Name:     "FUEL TOTAL QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelWeightPerGallon Simvar
func SimVarFuelWeightPerGallon() SimVar {
	return SimVar{
		Name:     "FUEL WEIGHT PER GALLON",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarFuelTankSelector Simvar
func SimVarFuelTankSelector() SimVar {
	return SimVar{
		Name:     "FUEL TANK SELECTOR:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarFuelCrossFeed Simvar
func SimVarFuelCrossFeed() SimVar {
	return SimVar{
		Name:     "FUEL CROSS FEED",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarFuelTotalCapacity Simvar
func SimVarFuelTotalCapacity() SimVar {
	return SimVar{
		Name:     "FUEL TOTAL CAPACITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelSelectedQuantityPercent Simvar
func SimVarFuelSelectedQuantityPercent() SimVar {
	return SimVar{
		Name:     "FUEL SELECTED QUANTITY PERCENT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFuelSelectedQuantity Simvar
func SimVarFuelSelectedQuantity() SimVar {
	return SimVar{
		Name:     "FUEL SELECTED QUANTITY",
		Units:    "Gallons",
		Settable: false,
	}
}

// SimVarFuelTotalQuantityWeight Simvar
func SimVarFuelTotalQuantityWeight() SimVar {
	return SimVar{
		Name:     "FUEL TOTAL QUANTITY WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarNumFuelSelectors Simvar
func SimVarNumFuelSelectors() SimVar {
	return SimVar{
		Name:     "NUM FUEL SELECTORS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUnlimitedFuel Simvar
func SimVarUnlimitedFuel() SimVar {
	return SimVar{
		Name:     "UNLIMITED FUEL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEstimatedFuelFlow Simvar
func SimVarEstimatedFuelFlow() SimVar {
	return SimVar{
		Name:     "ESTIMATED FUEL FLOW",
		Units:    "Pounds per hour",
		Settable: false,
	}
}

// SimVarLightStrobe Simvar
func SimVarLightStrobe() SimVar {
	return SimVar{
		Name:     "LIGHT STROBE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightPanel Simvar
func SimVarLightPanel() SimVar {
	return SimVar{
		Name:     "LIGHT PANEL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLanding Simvar
func SimVarLightLanding() SimVar {
	return SimVar{
		Name:     "LIGHT LANDING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightTaxi Simvar
func SimVarLightTaxi() SimVar {
	return SimVar{
		Name:     "LIGHT TAXI",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightBeacon Simvar
func SimVarLightBeacon() SimVar {
	return SimVar{
		Name:     "LIGHT BEACON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightNav Simvar
func SimVarLightNav() SimVar {
	return SimVar{
		Name:     "LIGHT NAV",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightLogo Simvar
func SimVarLightLogo() SimVar {
	return SimVar{
		Name:     "LIGHT LOGO",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightWing Simvar
func SimVarLightWing() SimVar {
	return SimVar{
		Name:     "LIGHT WING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightRecognition Simvar
func SimVarLightRecognition() SimVar {
	return SimVar{
		Name:     "LIGHT RECOGNITION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarLightCabin Simvar
func SimVarLightCabin() SimVar {
	return SimVar{
		Name:     "LIGHT CABIN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGroundVelocity Simvar
func SimVarGroundVelocity() SimVar {
	return SimVar{
		Name:     "GROUND VELOCITY",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarTotalWorldVelocity Simvar
func SimVarTotalWorldVelocity() SimVar {
	return SimVar{
		Name:     "TOTAL WORLD VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarVelocityBodyZ Simvar
func SimVarVelocityBodyZ() SimVar {
	return SimVar{
		Name:     "VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityBodyX Simvar
func SimVarVelocityBodyX() SimVar {
	return SimVar{
		Name:     "VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityBodyY Simvar
func SimVarVelocityBodyY() SimVar {
	return SimVar{
		Name:     "VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldZ Simvar
func SimVarVelocityWorldZ() SimVar {
	return SimVar{
		Name:     "VELOCITY WORLD Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldX Simvar
func SimVarVelocityWorldX() SimVar {
	return SimVar{
		Name:     "VELOCITY WORLD X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarVelocityWorldY Simvar
func SimVarVelocityWorldY() SimVar {
	return SimVar{
		Name:     "VELOCITY WORLD Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarAccelerationWorldX Simvar
func SimVarAccelerationWorldX() SimVar {
	return SimVar{
		Name:     "ACCELERATION WORLD X",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationWorldY Simvar
func SimVarAccelerationWorldY() SimVar {
	return SimVar{
		Name:     "ACCELERATION WORLD Y",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationWorldZ Simvar
func SimVarAccelerationWorldZ() SimVar {
	return SimVar{
		Name:     "ACCELERATION WORLD Z",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyX Simvar
func SimVarAccelerationBodyX() SimVar {
	return SimVar{
		Name:     "ACCELERATION BODY X",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyY Simvar
func SimVarAccelerationBodyY() SimVar {
	return SimVar{
		Name:     "ACCELERATION BODY Y",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarAccelerationBodyZ Simvar
func SimVarAccelerationBodyZ() SimVar {
	return SimVar{
		Name:     "ACCELERATION BODY Z",
		Units:    "Feet per second squared",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyX Simvar
func SimVarRotationVelocityBodyX() SimVar {
	return SimVar{
		Name:     "ROTATION VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyY Simvar
func SimVarRotationVelocityBodyY() SimVar {
	return SimVar{
		Name:     "ROTATION VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRotationVelocityBodyZ Simvar
func SimVarRotationVelocityBodyZ() SimVar {
	return SimVar{
		Name:     "ROTATION VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarRelativeWindVelocityBodyX Simvar
func SimVarRelativeWindVelocityBodyX() SimVar {
	return SimVar{
		Name:     "RELATIVE WIND VELOCITY BODY X",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyY Simvar
func SimVarRelativeWindVelocityBodyY() SimVar {
	return SimVar{
		Name:     "RELATIVE WIND VELOCITY BODY Y",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarRelativeWindVelocityBodyZ Simvar
func SimVarRelativeWindVelocityBodyZ() SimVar {
	return SimVar{
		Name:     "RELATIVE WIND VELOCITY BODY Z",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPlaneAltAboveGround Simvar
func SimVarPlaneAltAboveGround() SimVar {
	return SimVar{
		Name:     "PLANE ALT ABOVE GROUND",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarPlaneLatitude Simvar
func SimVarPlaneLatitude() SimVar {
	return SimVar{
		Name:     "PLANE LATITUDE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneLongitude Simvar
func SimVarPlaneLongitude() SimVar {
	return SimVar{
		Name:     "PLANE LONGITUDE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneAltitude Simvar
func SimVarPlaneAltitude() SimVar {
	return SimVar{
		Name:     "PLANE ALTITUDE",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarPlanePitchDegrees Simvar
func SimVarPlanePitchDegrees() SimVar {
	return SimVar{
		Name:     "PLANE PITCH DEGREES",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneBankDegrees Simvar
func SimVarPlaneBankDegrees() SimVar {
	return SimVar{
		Name:     "PLANE BANK DEGREES",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesTrue Simvar
func SimVarPlaneHeadingDegreesTrue() SimVar {
	return SimVar{
		Name:     "PLANE HEADING DEGREES TRUE",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesMagnetic Simvar
func SimVarPlaneHeadingDegreesMagnetic() SimVar {
	return SimVar{
		Name:     "PLANE HEADING DEGREES MAGNETIC",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarMagvar Simvar
func SimVarMagvar() SimVar {
	return SimVar{
		Name:     "MAGVAR",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGroundAltitude Simvar
func SimVarGroundAltitude() SimVar {
	return SimVar{
		Name:     "GROUND ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarSurfaceType Simvar
func SimVarSurfaceType() SimVar {
	return SimVar{
		Name:     "SURFACE TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarSimOnGround Simvar
func SimVarSimOnGround() SimVar {
	return SimVar{
		Name:     "SIM ON GROUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIncidenceAlpha Simvar
func SimVarIncidenceAlpha() SimVar {
	return SimVar{
		Name:     "INCIDENCE ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarIncidenceBeta Simvar
func SimVarIncidenceBeta() SimVar {
	return SimVar{
		Name:     "INCIDENCE BETA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAirspeedTrue Simvar
func SimVarAirspeedTrue() SimVar {
	return SimVar{
		Name:     "AIRSPEED TRUE",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAirspeedIndicated Simvar
func SimVarAirspeedIndicated() SimVar {
	return SimVar{
		Name:     "AIRSPEED INDICATED",
		Units:    "Knots",
		Settable: true,
	}
}

// SimVarAirspeedTrueCalibrate Simvar
func SimVarAirspeedTrueCalibrate() SimVar {
	return SimVar{
		Name:     "AIRSPEED TRUE CALIBRATE",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarAirspeedBarberPole Simvar
func SimVarAirspeedBarberPole() SimVar {
	return SimVar{
		Name:     "AIRSPEED BARBER POLE",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAirspeedMach Simvar
func SimVarAirspeedMach() SimVar {
	return SimVar{
		Name:     "AIRSPEED MACH",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarVerticalSpeed Simvar
func SimVarVerticalSpeed() SimVar {
	return SimVar{
		Name:     "VERTICAL SPEED",
		Units:    "Feet per second",
		Settable: true,
	}
}

// SimVarMachMaxOperate Simvar
func SimVarMachMaxOperate() SimVar {
	return SimVar{
		Name:     "MACH MAX OPERATE",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarStallWarning Simvar
func SimVarStallWarning() SimVar {
	return SimVar{
		Name:     "STALL WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarOverspeedWarning Simvar
func SimVarOverspeedWarning() SimVar {
	return SimVar{
		Name:     "OVERSPEED WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarBarberPoleMach Simvar
func SimVarBarberPoleMach() SimVar {
	return SimVar{
		Name:     "BARBER POLE MACH",
		Units:    "Mach",
		Settable: false,
	}
}

// SimVarIndicatedAltitude Simvar
func SimVarIndicatedAltitude() SimVar {
	return SimVar{
		Name:     "INDICATED ALTITUDE",
		Units:    "Feet",
		Settable: true,
	}
}

// SimVarKohlsmanSettingMb Simvar
func SimVarKohlsmanSettingMb() SimVar {
	return SimVar{
		Name:     "KOHLSMAN SETTING MB",
		Units:    "Millibars",
		Settable: true,
	}
}

// SimVarKohlsmanSettingHg Simvar
func SimVarKohlsmanSettingHg() SimVar {
	return SimVar{
		Name:     "KOHLSMAN SETTING HG",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarAttitudeIndicatorPitchDegrees Simvar
func SimVarAttitudeIndicatorPitchDegrees() SimVar {
	return SimVar{
		Name:     "ATTITUDE INDICATOR PITCH DEGREES",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAttitudeIndicatorBankDegrees Simvar
func SimVarAttitudeIndicatorBankDegrees() SimVar {
	return SimVar{
		Name:     "ATTITUDE INDICATOR BANK DEGREES",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAttitudeBarsPosition Simvar
func SimVarAttitudeBarsPosition() SimVar {
	return SimVar{
		Name:     "ATTITUDE BARS POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAttitudeCage Simvar
func SimVarAttitudeCage() SimVar {
	return SimVar{
		Name:     "ATTITUDE CAGE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarWiskeyCompassIndicationDegrees Simvar
func SimVarWiskeyCompassIndicationDegrees() SimVar {
	return SimVar{
		Name:     "WISKEY COMPASS INDICATION DEGREES",
		Units:    "Degrees",
		Settable: true,
	}
}

// SimVarPlaneHeadingDegreesGyro Simvar
func SimVarPlaneHeadingDegreesGyro() SimVar {
	return SimVar{
		Name:     "PLANE HEADING DEGREES GYRO",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarHeadingIndicator Simvar
func SimVarHeadingIndicator() SimVar {
	return SimVar{
		Name:     "HEADING INDICATOR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGyroDriftError Simvar
func SimVarGyroDriftError() SimVar {
	return SimVar{
		Name:     "GYRO DRIFT ERROR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDeltaHeadingRate Simvar
func SimVarDeltaHeadingRate() SimVar {
	return SimVar{
		Name:     "DELTA HEADING RATE",
		Units:    "Radians per second",
		Settable: true,
	}
}

// SimVarTurnCoordinatorBall Simvar
func SimVarTurnCoordinatorBall() SimVar {
	return SimVar{
		Name:     "TURN COORDINATOR BALL",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarAngleOfAttackIndicator Simvar
func SimVarAngleOfAttackIndicator() SimVar {
	return SimVar{
		Name:     "ANGLE OF ATTACK INDICATOR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRadioHeight Simvar
func SimVarRadioHeight() SimVar {
	return SimVar{
		Name:     "RADIO HEIGHT",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPartialPanelAdf Simvar
func SimVarPartialPanelAdf() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL ADF",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAirspeed Simvar
func SimVarPartialPanelAirspeed() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL AIRSPEED",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAltimeter Simvar
func SimVarPartialPanelAltimeter() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL ALTIMETER",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAttitude Simvar
func SimVarPartialPanelAttitude() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL ATTITUDE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelComm Simvar
func SimVarPartialPanelComm() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL COMM",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelCompass Simvar
func SimVarPartialPanelCompass() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL COMPASS",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelElectrical Simvar
func SimVarPartialPanelElectrical() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL ELECTRICAL",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelAvionics Simvar
func SimVarPartialPanelAvionics() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL AVIONICS",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelEngine Simvar
func SimVarPartialPanelEngine() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL ENGINE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelFuelIndicator Simvar
func SimVarPartialPanelFuelIndicator() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL FUEL INDICATOR",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelHeading Simvar
func SimVarPartialPanelHeading() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL HEADING",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelVerticalVelocity Simvar
func SimVarPartialPanelVerticalVelocity() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL VERTICAL VELOCITY",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelTransponder Simvar
func SimVarPartialPanelTransponder() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL TRANSPONDER",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelNav Simvar
func SimVarPartialPanelNav() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL NAV",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelPitot Simvar
func SimVarPartialPanelPitot() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL PITOT",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarPartialPanelTurnCoordinator Simvar
func SimVarPartialPanelTurnCoordinator() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL TURN COORDINATOR",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPartialPanelVacuum Simvar
func SimVarPartialPanelVacuum() SimVar {
	return SimVar{
		Name:     "PARTIAL PANEL VACUUM",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarMaxGForce Simvar
func SimVarMaxGForce() SimVar {
	return SimVar{
		Name:     "MAX G FORCE",
		Units:    "Gforce",
		Settable: false,
	}
}

// SimVarMinGForce Simvar
func SimVarMinGForce() SimVar {
	return SimVar{
		Name:     "MIN G FORCE",
		Units:    "Gforce",
		Settable: false,
	}
}

// SimVarSuctionPressure Simvar
func SimVarSuctionPressure() SimVar {
	return SimVar{
		Name:     "SUCTION PRESSURE",
		Units:    "inHg",
		Settable: true,
	}
}

// SimVarAvionicsMasterSwitch Simvar
func SimVarAvionicsMasterSwitch() SimVar {
	return SimVar{
		Name:     "AVIONICS MASTER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavSound Simvar
func SimVarNavSound() SimVar {
	return SimVar{
		Name:     "NAV SOUND:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarDmeSound Simvar
func SimVarDmeSound() SimVar {
	return SimVar{
		Name:     "DME SOUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfSound Simvar
func SimVarAdfSound() SimVar {
	return SimVar{
		Name:     "ADF SOUND:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarMarkerSound Simvar
func SimVarMarkerSound() SimVar {
	return SimVar{
		Name:     "MARKER SOUND",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComTransmit Simvar
func SimVarComTransmit() SimVar {
	return SimVar{
		Name:     "COM TRANSMIT:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComRecieveAll Simvar
func SimVarComRecieveAll() SimVar {
	return SimVar{
		Name:     "COM RECIEVE ALL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComActiveFrequency Simvar
func SimVarComActiveFrequency() SimVar {
	return SimVar{
		Name:     "COM ACTIVE FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarComStandbyFrequency Simvar
func SimVarComStandbyFrequency() SimVar {
	return SimVar{
		Name:     "COM STANDBY FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarComStatus Simvar
func SimVarComStatus() SimVar {
	return SimVar{
		Name:     "COM STATUS:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarNavAvailable Simvar
func SimVarNavAvailable() SimVar {
	return SimVar{
		Name:     "NAV AVAILABLE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavActiveFrequency Simvar
func SimVarNavActiveFrequency() SimVar {
	return SimVar{
		Name:     "NAV ACTIVE FREQUENCY:index",
		Units:    "MHz",
		Settable: false,
	}
}

// SimVarNavStandbyFrequency Simvar
func SimVarNavStandbyFrequency() SimVar {
	return SimVar{
		Name:     "NAV STANDBY FREQUENCY:index",
		Units:    "MHz",
		Settable: false,
	}
}

// SimVarNavSignal Simvar
func SimVarNavSignal() SimVar {
	return SimVar{
		Name:     "NAV SIGNAL:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavHasNav Simvar
func SimVarNavHasNav() SimVar {
	return SimVar{
		Name:     "NAV HAS NAV:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasLocalizer Simvar
func SimVarNavHasLocalizer() SimVar {
	return SimVar{
		Name:     "NAV HAS LOCALIZER:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasDme Simvar
func SimVarNavHasDme() SimVar {
	return SimVar{
		Name:     "NAV HAS DME:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavHasGlideSlope Simvar
func SimVarNavHasGlideSlope() SimVar {
	return SimVar{
		Name:     "NAV HAS GLIDE SLOPE:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavBackCourseFlags Simvar
func SimVarNavBackCourseFlags() SimVar {
	return SimVar{
		Name:     "NAV BACK COURSE FLAGS:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavMagvar Simvar
func SimVarNavMagvar() SimVar {
	return SimVar{
		Name:     "NAV MAGVAR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavRadial Simvar
func SimVarNavRadial() SimVar {
	return SimVar{
		Name:     "NAV RADIAL:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavRadialError Simvar
func SimVarNavRadialError() SimVar {
	return SimVar{
		Name:     "NAV RADIAL ERROR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavLocalizer Simvar
func SimVarNavLocalizer() SimVar {
	return SimVar{
		Name:     "NAV LOCALIZER:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavGlideSlopeError Simvar
func SimVarNavGlideSlopeError() SimVar {
	return SimVar{
		Name:     "NAV GLIDE SLOPE ERROR:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavCdi Simvar
func SimVarNavCdi() SimVar {
	return SimVar{
		Name:     "NAV CDI:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavGsi Simvar
func SimVarNavGsi() SimVar {
	return SimVar{
		Name:     "NAV GSI:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavTofrom Simvar
func SimVarNavTofrom() SimVar {
	return SimVar{
		Name:     "NAV TOFROM:index",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarNavGsFlag Simvar
func SimVarNavGsFlag() SimVar {
	return SimVar{
		Name:     "NAV GS FLAG:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarNavObs Simvar
func SimVarNavObs() SimVar {
	return SimVar{
		Name:     "NAV OBS:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarNavDme Simvar
func SimVarNavDme() SimVar {
	return SimVar{
		Name:     "NAV DME:index",
		Units:    "Nautical miles",
		Settable: false,
	}
}

// SimVarNavDmespeed Simvar
func SimVarNavDmespeed() SimVar {
	return SimVar{
		Name:     "NAV DMESPEED:index",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAdfActiveFrequency Simvar
func SimVarAdfActiveFrequency() SimVar {
	return SimVar{
		Name:     "ADF ACTIVE FREQUENCY:index",
		Units:    "Frequency ADF BCD32",
		Settable: false,
	}
}

// SimVarAdfStandbyFrequency Simvar
func SimVarAdfStandbyFrequency() SimVar {
	return SimVar{
		Name:     "ADF STANDBY FREQUENCY:index",
		Units:    "Hz",
		Settable: false,
	}
}

// SimVarAdfRadial Simvar
func SimVarAdfRadial() SimVar {
	return SimVar{
		Name:     "ADF RADIAL:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAdfSignal Simvar
func SimVarAdfSignal() SimVar {
	return SimVar{
		Name:     "ADF SIGNAL:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTransponderCode Simvar
func SimVarTransponderCode() SimVar {
	return SimVar{
		Name:     "TRANSPONDER CODE:index",
		Units:    "BCO16",
		Settable: false,
	}
}

// SimVarMarkerBeaconState Simvar
func SimVarMarkerBeaconState() SimVar {
	return SimVar{
		Name:     "MARKER BEACON STATE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarInnerMarker Simvar
func SimVarInnerMarker() SimVar {
	return SimVar{
		Name:     "INNER MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarMiddleMarker Simvar
func SimVarMiddleMarker() SimVar {
	return SimVar{
		Name:     "MIDDLE MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarOuterMarker Simvar
func SimVarOuterMarker() SimVar {
	return SimVar{
		Name:     "OUTER MARKER",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarNavRawGlideSlope Simvar
func SimVarNavRawGlideSlope() SimVar {
	return SimVar{
		Name:     "NAV RAW GLIDE SLOPE:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAdfCard Simvar
func SimVarAdfCard() SimVar {
	return SimVar{
		Name:     "ADF CARD",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarHsiCdiNeedle Simvar
func SimVarHsiCdiNeedle() SimVar {
	return SimVar{
		Name:     "HSI CDI NEEDLE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHsiGsiNeedle Simvar
func SimVarHsiGsiNeedle() SimVar {
	return SimVar{
		Name:     "HSI GSI NEEDLE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarHsiCdiNeedleValid Simvar
func SimVarHsiCdiNeedleValid() SimVar {
	return SimVar{
		Name:     "HSI CDI NEEDLE VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiGsiNeedleValid Simvar
func SimVarHsiGsiNeedleValid() SimVar {
	return SimVar{
		Name:     "HSI GSI NEEDLE VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiTfFlags Simvar
func SimVarHsiTfFlags() SimVar {
	return SimVar{
		Name:     "HSI TF FLAGS",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarHsiBearingValid Simvar
func SimVarHsiBearingValid() SimVar {
	return SimVar{
		Name:     "HSI BEARING VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiBearing Simvar
func SimVarHsiBearing() SimVar {
	return SimVar{
		Name:     "HSI BEARING",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarHsiHasLocalizer Simvar
func SimVarHsiHasLocalizer() SimVar {
	return SimVar{
		Name:     "HSI HAS LOCALIZER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHsiSpeed Simvar
func SimVarHsiSpeed() SimVar {
	return SimVar{
		Name:     "HSI SPEED",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarHsiDistance Simvar
func SimVarHsiDistance() SimVar {
	return SimVar{
		Name:     "HSI DISTANCE",
		Units:    "Nautical miles",
		Settable: false,
	}
}

// SimVarGpsPositionLat Simvar
func SimVarGpsPositionLat() SimVar {
	return SimVar{
		Name:     "GPS POSITION LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsPositionLon Simvar
func SimVarGpsPositionLon() SimVar {
	return SimVar{
		Name:     "GPS POSITION LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsPositionAlt Simvar
func SimVarGpsPositionAlt() SimVar {
	return SimVar{
		Name:     "GPS POSITION ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsMagvar Simvar
func SimVarGpsMagvar() SimVar {
	return SimVar{
		Name:     "GPS MAGVAR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsIsActiveFlightPlan Simvar
func SimVarGpsIsActiveFlightPlan() SimVar {
	return SimVar{
		Name:     "GPS IS ACTIVE FLIGHT PLAN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsActiveWayPoint Simvar
func SimVarGpsIsActiveWayPoint() SimVar {
	return SimVar{
		Name:     "GPS IS ACTIVE WAY POINT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsArrived Simvar
func SimVarGpsIsArrived() SimVar {
	return SimVar{
		Name:     "GPS IS ARRIVED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsDirecttoFlightplan Simvar
func SimVarGpsIsDirecttoFlightplan() SimVar {
	return SimVar{
		Name:     "GPS IS DIRECTTO FLIGHTPLAN",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsGroundSpeed Simvar
func SimVarGpsGroundSpeed() SimVar {
	return SimVar{
		Name:     "GPS GROUND SPEED",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarGpsGroundTrueHeading Simvar
func SimVarGpsGroundTrueHeading() SimVar {
	return SimVar{
		Name:     "GPS GROUND TRUE HEADING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsGroundMagneticTrack Simvar
func SimVarGpsGroundMagneticTrack() SimVar {
	return SimVar{
		Name:     "GPS GROUND MAGNETIC TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsGroundTrueTrack Simvar
func SimVarGpsGroundTrueTrack() SimVar {
	return SimVar{
		Name:     "GPS GROUND TRUE TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpDistance Simvar
func SimVarGpsWpDistance() SimVar {
	return SimVar{
		Name:     "GPS WP DISTANCE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpBearing Simvar
func SimVarGpsWpBearing() SimVar {
	return SimVar{
		Name:     "GPS WP BEARING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpTrueBearing Simvar
func SimVarGpsWpTrueBearing() SimVar {
	return SimVar{
		Name:     "GPS WP TRUE BEARING",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpCrossTrk Simvar
func SimVarGpsWpCrossTrk() SimVar {
	return SimVar{
		Name:     "GPS WP CROSS TRK",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpDesiredTrack Simvar
func SimVarGpsWpDesiredTrack() SimVar {
	return SimVar{
		Name:     "GPS WP DESIRED TRACK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpTrueReqHdg Simvar
func SimVarGpsWpTrueReqHdg() SimVar {
	return SimVar{
		Name:     "GPS WP TRUE REQ HDG",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsWpVerticalSpeed Simvar
func SimVarGpsWpVerticalSpeed() SimVar {
	return SimVar{
		Name:     "GPS WP VERTICAL SPEED",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarGpsWpTrackAngleError Simvar
func SimVarGpsWpTrackAngleError() SimVar {
	return SimVar{
		Name:     "GPS WP TRACK ANGLE ERROR",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsEte Simvar
func SimVarGpsEte() SimVar {
	return SimVar{
		Name:     "GPS ETE",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsEta Simvar
func SimVarGpsEta() SimVar {
	return SimVar{
		Name:     "GPS ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsWpNextLat Simvar
func SimVarGpsWpNextLat() SimVar {
	return SimVar{
		Name:     "GPS WP NEXT LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpNextLon Simvar
func SimVarGpsWpNextLon() SimVar {
	return SimVar{
		Name:     "GPS WP NEXT LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpNextAlt Simvar
func SimVarGpsWpNextAlt() SimVar {
	return SimVar{
		Name:     "GPS WP NEXT ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpPrevValid Simvar
func SimVarGpsWpPrevValid() SimVar {
	return SimVar{
		Name:     "GPS WP PREV VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsWpPrevLat Simvar
func SimVarGpsWpPrevLat() SimVar {
	return SimVar{
		Name:     "GPS WP PREV LAT",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpPrevLon Simvar
func SimVarGpsWpPrevLon() SimVar {
	return SimVar{
		Name:     "GPS WP PREV LON",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarGpsWpPrevAlt Simvar
func SimVarGpsWpPrevAlt() SimVar {
	return SimVar{
		Name:     "GPS WP PREV ALT",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsWpEte Simvar
func SimVarGpsWpEte() SimVar {
	return SimVar{
		Name:     "GPS WP ETE",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsWpEta Simvar
func SimVarGpsWpEta() SimVar {
	return SimVar{
		Name:     "GPS WP ETA",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsCourseToSteer Simvar
func SimVarGpsCourseToSteer() SimVar {
	return SimVar{
		Name:     "GPS COURSE TO STEER",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpIndex Simvar
func SimVarGpsFlightPlanWpIndex() SimVar {
	return SimVar{
		Name:     "GPS FLIGHT PLAN WP INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsFlightPlanWpCount Simvar
func SimVarGpsFlightPlanWpCount() SimVar {
	return SimVar{
		Name:     "GPS FLIGHT PLAN WP COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsIsActiveWpLocked Simvar
func SimVarGpsIsActiveWpLocked() SimVar {
	return SimVar{
		Name:     "GPS IS ACTIVE WP LOCKED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsApproachLoaded Simvar
func SimVarGpsIsApproachLoaded() SimVar {
	return SimVar{
		Name:     "GPS IS APPROACH LOADED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsIsApproachActive Simvar
func SimVarGpsIsApproachActive() SimVar {
	return SimVar{
		Name:     "GPS IS APPROACH ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachMode Simvar
func SimVarGpsApproachMode() SimVar {
	return SimVar{
		Name:     "GPS APPROACH MODE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachWpType Simvar
func SimVarGpsApproachWpType() SimVar {
	return SimVar{
		Name:     "GPS APPROACH WP TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachIsWpRunway Simvar
func SimVarGpsApproachIsWpRunway() SimVar {
	return SimVar{
		Name:     "GPS APPROACH IS WP RUNWAY",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachSegmentType Simvar
func SimVarGpsApproachSegmentType() SimVar {
	return SimVar{
		Name:     "GPS APPROACH SEGMENT TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachApproachIndex Simvar
func SimVarGpsApproachApproachIndex() SimVar {
	return SimVar{
		Name:     "GPS APPROACH APPROACH INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachApproachType Simvar
func SimVarGpsApproachApproachType() SimVar {
	return SimVar{
		Name:     "GPS APPROACH APPROACH TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarGpsApproachTransitionIndex Simvar
func SimVarGpsApproachTransitionIndex() SimVar {
	return SimVar{
		Name:     "GPS APPROACH TRANSITION INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachIsFinal Simvar
func SimVarGpsApproachIsFinal() SimVar {
	return SimVar{
		Name:     "GPS APPROACH IS FINAL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachIsMissed Simvar
func SimVarGpsApproachIsMissed() SimVar {
	return SimVar{
		Name:     "GPS APPROACH IS MISSED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpsApproachTimezoneDeviation Simvar
func SimVarGpsApproachTimezoneDeviation() SimVar {
	return SimVar{
		Name:     "GPS APPROACH TIMEZONE DEVIATION",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarGpsApproachWpIndex Simvar
func SimVarGpsApproachWpIndex() SimVar {
	return SimVar{
		Name:     "GPS APPROACH WP INDEX",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsApproachWpCount Simvar
func SimVarGpsApproachWpCount() SimVar {
	return SimVar{
		Name:     "GPS APPROACH WP COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsDrivesNav1 Simvar
func SimVarGpsDrivesNav1() SimVar {
	return SimVar{
		Name:     "GPS DRIVES NAV1",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComReceiveAll Simvar
func SimVarComReceiveAll() SimVar {
	return SimVar{
		Name:     "COM RECEIVE ALL",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComAvailable Simvar
func SimVarComAvailable() SimVar {
	return SimVar{
		Name:     "COM AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarComTest Simvar
func SimVarComTest() SimVar {
	return SimVar{
		Name:     "COM TEST:index",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTransponderAvailable Simvar
func SimVarTransponderAvailable() SimVar {
	return SimVar{
		Name:     "TRANSPONDER AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfAvailable Simvar
func SimVarAdfAvailable() SimVar {
	return SimVar{
		Name:     "ADF AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAdfFrequency Simvar
func SimVarAdfFrequency() SimVar {
	return SimVar{
		Name:     "ADF FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarAdfExtFrequency Simvar
func SimVarAdfExtFrequency() SimVar {
	return SimVar{
		Name:     "ADF EXT FREQUENCY:index",
		Units:    "Frequency BCD16",
		Settable: false,
	}
}

// SimVarAdfIdent Simvar
func SimVarAdfIdent() SimVar {
	return SimVar{
		Name:     "ADF IDENT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAdfName Simvar
func SimVarAdfName() SimVar {
	return SimVar{
		Name:     "ADF NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavIdent Simvar
func SimVarNavIdent() SimVar {
	return SimVar{
		Name:     "NAV IDENT",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavName Simvar
func SimVarNavName() SimVar {
	return SimVar{
		Name:     "NAV NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarNavCodes Simvar
func SimVarNavCodes() SimVar {
	return SimVar{
		Name:     "NAV CODES:index",
		Units:    "Flags",
		Settable: false,
	}
}

// SimVarNavGlideSlope Simvar
func SimVarNavGlideSlope() SimVar {
	return SimVar{
		Name:     "NAV GLIDE SLOPE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarNavRelativeBearingToStation Simvar
func SimVarNavRelativeBearingToStation() SimVar {
	return SimVar{
		Name:     "NAV RELATIVE BEARING TO STATION:index",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarSelectedDme Simvar
func SimVarSelectedDme() SimVar {
	return SimVar{
		Name:     "SELECTED DME",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGpsWpNextId Simvar
func SimVarGpsWpNextId() SimVar {
	return SimVar{
		Name:     "GPS WP NEXT ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsWpPrevId Simvar
func SimVarGpsWpPrevId() SimVar {
	return SimVar{
		Name:     "GPS WP PREV ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsTargetDistance Simvar
func SimVarGpsTargetDistance() SimVar {
	return SimVar{
		Name:     "GPS TARGET DISTANCE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarGpsTargetAltitude Simvar
func SimVarGpsTargetAltitude() SimVar {
	return SimVar{
		Name:     "GPS TARGET ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarYokeYPosition Simvar
func SimVarYokeYPosition() SimVar {
	return SimVar{
		Name:     "YOKE Y POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarYokeXPosition Simvar
func SimVarYokeXPosition() SimVar {
	return SimVar{
		Name:     "YOKE X POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRudderPedalPosition Simvar
func SimVarRudderPedalPosition() SimVar {
	return SimVar{
		Name:     "RUDDER PEDAL POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarRudderPosition Simvar
func SimVarRudderPosition() SimVar {
	return SimVar{
		Name:     "RUDDER POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarElevatorPosition Simvar
func SimVarElevatorPosition() SimVar {
	return SimVar{
		Name:     "ELEVATOR POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarAileronPosition Simvar
func SimVarAileronPosition() SimVar {
	return SimVar{
		Name:     "AILERON POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarElevatorTrimPosition Simvar
func SimVarElevatorTrimPosition() SimVar {
	return SimVar{
		Name:     "ELEVATOR TRIM POSITION",
		Units:    "Radians",
		Settable: true,
	}
}

// SimVarElevatorTrimIndicator Simvar
func SimVarElevatorTrimIndicator() SimVar {
	return SimVar{
		Name:     "ELEVATOR TRIM INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarElevatorTrimPct Simvar
func SimVarElevatorTrimPct() SimVar {
	return SimVar{
		Name:     "ELEVATOR TRIM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarBrakeLeftPosition Simvar
func SimVarBrakeLeftPosition() SimVar {
	return SimVar{
		Name:     "BRAKE LEFT POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeRightPosition Simvar
func SimVarBrakeRightPosition() SimVar {
	return SimVar{
		Name:     "BRAKE RIGHT POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeIndicator Simvar
func SimVarBrakeIndicator() SimVar {
	return SimVar{
		Name:     "BRAKE INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarBrakeParkingPosition Simvar
func SimVarBrakeParkingPosition() SimVar {
	return SimVar{
		Name:     "BRAKE PARKING POSITION",
		Units:    "Position",
		Settable: true,
	}
}

// SimVarBrakeParkingIndicator Simvar
func SimVarBrakeParkingIndicator() SimVar {
	return SimVar{
		Name:     "BRAKE PARKING INDICATOR",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilersArmed Simvar
func SimVarSpoilersArmed() SimVar {
	return SimVar{
		Name:     "SPOILERS ARMED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilersHandlePosition Simvar
func SimVarSpoilersHandlePosition() SimVar {
	return SimVar{
		Name:     "SPOILERS HANDLE POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarSpoilersLeftPosition Simvar
func SimVarSpoilersLeftPosition() SimVar {
	return SimVar{
		Name:     "SPOILERS LEFT POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarSpoilersRightPosition Simvar
func SimVarSpoilersRightPosition() SimVar {
	return SimVar{
		Name:     "SPOILERS RIGHT POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFlapsHandlePercent Simvar
func SimVarFlapsHandlePercent() SimVar {
	return SimVar{
		Name:     "FLAPS HANDLE PERCENT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarFlapsHandleIndex Simvar
func SimVarFlapsHandleIndex() SimVar {
	return SimVar{
		Name:     "FLAPS HANDLE INDEX",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarFlapsNumHandlePositions Simvar
func SimVarFlapsNumHandlePositions() SimVar {
	return SimVar{
		Name:     "FLAPS NUM HANDLE POSITIONS",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsLeftPercent Simvar
func SimVarTrailingEdgeFlapsLeftPercent() SimVar {
	return SimVar{
		Name:     "TRAILING EDGE FLAPS LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsRightPercent Simvar
func SimVarTrailingEdgeFlapsRightPercent() SimVar {
	return SimVar{
		Name:     "TRAILING EDGE FLAPS RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTrailingEdgeFlapsLeftAngle Simvar
func SimVarTrailingEdgeFlapsLeftAngle() SimVar {
	return SimVar{
		Name:     "TRAILING EDGE FLAPS LEFT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarTrailingEdgeFlapsRightAngle Simvar
func SimVarTrailingEdgeFlapsRightAngle() SimVar {
	return SimVar{
		Name:     "TRAILING EDGE FLAPS RIGHT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsLeftPercent Simvar
func SimVarLeadingEdgeFlapsLeftPercent() SimVar {
	return SimVar{
		Name:     "LEADING EDGE FLAPS LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsRightPercent Simvar
func SimVarLeadingEdgeFlapsRightPercent() SimVar {
	return SimVar{
		Name:     "LEADING EDGE FLAPS RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarLeadingEdgeFlapsLeftAngle Simvar
func SimVarLeadingEdgeFlapsLeftAngle() SimVar {
	return SimVar{
		Name:     "LEADING EDGE FLAPS LEFT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeadingEdgeFlapsRightAngle Simvar
func SimVarLeadingEdgeFlapsRightAngle() SimVar {
	return SimVar{
		Name:     "LEADING EDGE FLAPS RIGHT ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarIsGearRetractable Simvar
func SimVarIsGearRetractable() SimVar {
	return SimVar{
		Name:     "IS GEAR RETRACTABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearSkis Simvar
func SimVarIsGearSkis() SimVar {
	return SimVar{
		Name:     "IS GEAR SKIS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearFloats Simvar
func SimVarIsGearFloats() SimVar {
	return SimVar{
		Name:     "IS GEAR FLOATS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearSkids Simvar
func SimVarIsGearSkids() SimVar {
	return SimVar{
		Name:     "IS GEAR SKIDS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsGearWheels Simvar
func SimVarIsGearWheels() SimVar {
	return SimVar{
		Name:     "IS GEAR WHEELS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearHandlePosition Simvar
func SimVarGearHandlePosition() SimVar {
	return SimVar{
		Name:     "GEAR HANDLE POSITION",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGearHydraulicPressure Simvar
func SimVarGearHydraulicPressure() SimVar {
	return SimVar{
		Name:     "GEAR HYDRAULIC PRESSURE",
		Units:    "psf",
		Settable: false,
	}
}

// SimVarTailwheelLockOn Simvar
func SimVarTailwheelLockOn() SimVar {
	return SimVar{
		Name:     "TAILWHEEL LOCK ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearCenterPosition Simvar
func SimVarGearCenterPosition() SimVar {
	return SimVar{
		Name:     "GEAR CENTER POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearLeftPosition Simvar
func SimVarGearLeftPosition() SimVar {
	return SimVar{
		Name:     "GEAR LEFT POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearRightPosition Simvar
func SimVarGearRightPosition() SimVar {
	return SimVar{
		Name:     "GEAR RIGHT POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarGearTailPosition Simvar
func SimVarGearTailPosition() SimVar {
	return SimVar{
		Name:     "GEAR TAIL POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxPosition Simvar
func SimVarGearAuxPosition() SimVar {
	return SimVar{
		Name:     "GEAR AUX POSITION",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearPosition Simvar
func SimVarGearPosition() SimVar {
	return SimVar{
		Name:     "GEAR POSITION:index",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarGearAnimationPosition Simvar
func SimVarGearAnimationPosition() SimVar {
	return SimVar{
		Name:     "GEAR ANIMATION POSITION:index",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarGearTotalPctExtended Simvar
func SimVarGearTotalPctExtended() SimVar {
	return SimVar{
		Name:     "GEAR TOTAL PCT EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarAutoBrakeSwitchCb Simvar
func SimVarAutoBrakeSwitchCb() SimVar {
	return SimVar{
		Name:     "AUTO BRAKE SWITCH CB",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarWaterRudderHandlePosition Simvar
func SimVarWaterRudderHandlePosition() SimVar {
	return SimVar{
		Name:     "WATER RUDDER HANDLE POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarElevatorDeflection Simvar
func SimVarElevatorDeflection() SimVar {
	return SimVar{
		Name:     "ELEVATOR DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarElevatorDeflectionPct Simvar
func SimVarElevatorDeflectionPct() SimVar {
	return SimVar{
		Name:     "ELEVATOR DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderExtended Simvar
func SimVarWaterLeftRudderExtended() SimVar {
	return SimVar{
		Name:     "WATER LEFT RUDDER EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarWaterRightRudderExtended Simvar
func SimVarWaterRightRudderExtended() SimVar {
	return SimVar{
		Name:     "WATER RIGHT RUDDER EXTENDED",
		Units:    "Percentage",
		Settable: false,
	}
}

// SimVarGearCenterSteerAngle Simvar
func SimVarGearCenterSteerAngle() SimVar {
	return SimVar{
		Name:     "GEAR CENTER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearLeftSteerAngle Simvar
func SimVarGearLeftSteerAngle() SimVar {
	return SimVar{
		Name:     "GEAR LEFT STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearRightSteerAngle Simvar
func SimVarGearRightSteerAngle() SimVar {
	return SimVar{
		Name:     "GEAR RIGHT STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxSteerAngle Simvar
func SimVarGearAuxSteerAngle() SimVar {
	return SimVar{
		Name:     "GEAR AUX STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearSteerAngle Simvar
func SimVarGearSteerAngle() SimVar {
	return SimVar{
		Name:     "GEAR STEER ANGLE:index",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAngle Simvar
func SimVarWaterLeftRudderSteerAngle() SimVar {
	return SimVar{
		Name:     "WATER LEFT RUDDER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAngle Simvar
func SimVarWaterRightRudderSteerAngle() SimVar {
	return SimVar{
		Name:     "WATER RIGHT RUDDER STEER ANGLE",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearCenterSteerAnglePct Simvar
func SimVarGearCenterSteerAnglePct() SimVar {
	return SimVar{
		Name:     "GEAR CENTER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearLeftSteerAnglePct Simvar
func SimVarGearLeftSteerAnglePct() SimVar {
	return SimVar{
		Name:     "GEAR LEFT STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearRightSteerAnglePct Simvar
func SimVarGearRightSteerAnglePct() SimVar {
	return SimVar{
		Name:     "GEAR RIGHT STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearAuxSteerAnglePct Simvar
func SimVarGearAuxSteerAnglePct() SimVar {
	return SimVar{
		Name:     "GEAR AUX STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarGearSteerAnglePct Simvar
func SimVarGearSteerAnglePct() SimVar {
	return SimVar{
		Name:     "GEAR STEER ANGLE PCT:index",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterLeftRudderSteerAnglePct Simvar
func SimVarWaterLeftRudderSteerAnglePct() SimVar {
	return SimVar{
		Name:     "WATER LEFT RUDDER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarWaterRightRudderSteerAnglePct Simvar
func SimVarWaterRightRudderSteerAnglePct() SimVar {
	return SimVar{
		Name:     "WATER RIGHT RUDDER STEER ANGLE PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronLeftDeflection Simvar
func SimVarAileronLeftDeflection() SimVar {
	return SimVar{
		Name:     "AILERON LEFT DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronLeftDeflectionPct Simvar
func SimVarAileronLeftDeflectionPct() SimVar {
	return SimVar{
		Name:     "AILERON LEFT DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronRightDeflection Simvar
func SimVarAileronRightDeflection() SimVar {
	return SimVar{
		Name:     "AILERON RIGHT DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronRightDeflectionPct Simvar
func SimVarAileronRightDeflectionPct() SimVar {
	return SimVar{
		Name:     "AILERON RIGHT DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarAileronAverageDeflection Simvar
func SimVarAileronAverageDeflection() SimVar {
	return SimVar{
		Name:     "AILERON AVERAGE DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAileronTrim Simvar
func SimVarAileronTrim() SimVar {
	return SimVar{
		Name:     "AILERON TRIM",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRudderDeflection Simvar
func SimVarRudderDeflection() SimVar {
	return SimVar{
		Name:     "RUDDER DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRudderDeflectionPct Simvar
func SimVarRudderDeflectionPct() SimVar {
	return SimVar{
		Name:     "RUDDER DEFLECTION PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRudderTrim Simvar
func SimVarRudderTrim() SimVar {
	return SimVar{
		Name:     "RUDDER TRIM",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarFlapsAvailable Simvar
func SimVarFlapsAvailable() SimVar {
	return SimVar{
		Name:     "FLAPS AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearDamageBySpeed Simvar
func SimVarGearDamageBySpeed() SimVar {
	return SimVar{
		Name:     "GEAR DAMAGE BY SPEED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearSpeedExceeded Simvar
func SimVarGearSpeedExceeded() SimVar {
	return SimVar{
		Name:     "GEAR SPEED EXCEEDED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlapDamageBySpeed Simvar
func SimVarFlapDamageBySpeed() SimVar {
	return SimVar{
		Name:     "FLAP DAMAGE BY SPEED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFlapSpeedExceeded Simvar
func SimVarFlapSpeedExceeded() SimVar {
	return SimVar{
		Name:     "FLAP SPEED EXCEEDED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCenterWheelRpm Simvar
func SimVarCenterWheelRpm() SimVar {
	return SimVar{
		Name:     "CENTER WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarLeftWheelRpm Simvar
func SimVarLeftWheelRpm() SimVar {
	return SimVar{
		Name:     "LEFT WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarRightWheelRpm Simvar
func SimVarRightWheelRpm() SimVar {
	return SimVar{
		Name:     "RIGHT WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarAutopilotAvailable Simvar
func SimVarAutopilotAvailable() SimVar {
	return SimVar{
		Name:     "AUTOPILOT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMaster Simvar
func SimVarAutopilotMaster() SimVar {
	return SimVar{
		Name:     "AUTOPILOT MASTER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotNavSelected Simvar
func SimVarAutopilotNavSelected() SimVar {
	return SimVar{
		Name:     "AUTOPILOT NAV SELECTED",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotWingLeveler Simvar
func SimVarAutopilotWingLeveler() SimVar {
	return SimVar{
		Name:     "AUTOPILOT WING LEVELER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotHeadingLock Simvar
func SimVarAutopilotHeadingLock() SimVar {
	return SimVar{
		Name:     "AUTOPILOT HEADING LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotHeadingLockDir Simvar
func SimVarAutopilotHeadingLockDir() SimVar {
	return SimVar{
		Name:     "AUTOPILOT HEADING LOCK DIR",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLock Simvar
func SimVarAutopilotAltitudeLock() SimVar {
	return SimVar{
		Name:     "AUTOPILOT ALTITUDE LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotAltitudeLockVar Simvar
func SimVarAutopilotAltitudeLockVar() SimVar {
	return SimVar{
		Name:     "AUTOPILOT ALTITUDE LOCK VAR",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarAutopilotAttitudeHold Simvar
func SimVarAutopilotAttitudeHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT ATTITUDE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotGlideslopeHold Simvar
func SimVarAutopilotGlideslopeHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT GLIDESLOPE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotPitchHoldRef Simvar
func SimVarAutopilotPitchHoldRef() SimVar {
	return SimVar{
		Name:     "AUTOPILOT PITCH HOLD REF",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotApproachHold Simvar
func SimVarAutopilotApproachHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT APPROACH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotBackcourseHold Simvar
func SimVarAutopilotBackcourseHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT BACKCOURSE HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotVerticalHoldVar Simvar
func SimVarAutopilotVerticalHoldVar() SimVar {
	return SimVar{
		Name:     "AUTOPILOT VERTICAL HOLD VAR",
		Units:    "Feet/minute",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorActive Simvar
func SimVarAutopilotFlightDirectorActive() SimVar {
	return SimVar{
		Name:     "AUTOPILOT FLIGHT DIRECTOR ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorPitch Simvar
func SimVarAutopilotFlightDirectorPitch() SimVar {
	return SimVar{
		Name:     "AUTOPILOT FLIGHT DIRECTOR PITCH",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotFlightDirectorBank Simvar
func SimVarAutopilotFlightDirectorBank() SimVar {
	return SimVar{
		Name:     "AUTOPILOT FLIGHT DIRECTOR BANK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHold Simvar
func SimVarAutopilotAirspeedHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT AIRSPEED HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotAirspeedHoldVar Simvar
func SimVarAutopilotAirspeedHoldVar() SimVar {
	return SimVar{
		Name:     "AUTOPILOT AIRSPEED HOLD VAR",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAutopilotMachHold Simvar
func SimVarAutopilotMachHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT MACH HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMachHoldVar Simvar
func SimVarAutopilotMachHoldVar() SimVar {
	return SimVar{
		Name:     "AUTOPILOT MACH HOLD VAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotYawDamper Simvar
func SimVarAutopilotYawDamper() SimVar {
	return SimVar{
		Name:     "AUTOPILOT YAW DAMPER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotRpmHoldVar Simvar
func SimVarAutopilotRpmHoldVar() SimVar {
	return SimVar{
		Name:     "AUTOPILOT RPM HOLD VAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarAutopilotThrottleArm Simvar
func SimVarAutopilotThrottleArm() SimVar {
	return SimVar{
		Name:     "AUTOPILOT THROTTLE ARM",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotTakeoffPowerActive Simvar
func SimVarAutopilotTakeoffPowerActive() SimVar {
	return SimVar{
		Name:     "AUTOPILOT TAKEOFF POWER ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutothrottleActive Simvar
func SimVarAutothrottleActive() SimVar {
	return SimVar{
		Name:     "AUTOTHROTTLE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotNav1Lock Simvar
func SimVarAutopilotNav1Lock() SimVar {
	return SimVar{
		Name:     "AUTOPILOT NAV1 LOCK",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotVerticalHold Simvar
func SimVarAutopilotVerticalHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT VERTICAL HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotRpmHold Simvar
func SimVarAutopilotRpmHold() SimVar {
	return SimVar{
		Name:     "AUTOPILOT RPM HOLD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAutopilotMaxBank Simvar
func SimVarAutopilotMaxBank() SimVar {
	return SimVar{
		Name:     "AUTOPILOT MAX BANK",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarWheelRpm Simvar
func SimVarWheelRpm() SimVar {
	return SimVar{
		Name:     "WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarAuxWheelRpm Simvar
func SimVarAuxWheelRpm() SimVar {
	return SimVar{
		Name:     "AUX WHEEL RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarWheelRotationAngle Simvar
func SimVarWheelRotationAngle() SimVar {
	return SimVar{
		Name:     "WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCenterWheelRotationAngle Simvar
func SimVarCenterWheelRotationAngle() SimVar {
	return SimVar{
		Name:     "CENTER WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarLeftWheelRotationAngle Simvar
func SimVarLeftWheelRotationAngle() SimVar {
	return SimVar{
		Name:     "LEFT WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRightWheelRotationAngle Simvar
func SimVarRightWheelRotationAngle() SimVar {
	return SimVar{
		Name:     "RIGHT WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAuxWheelRotationAngle Simvar
func SimVarAuxWheelRotationAngle() SimVar {
	return SimVar{
		Name:     "AUX WHEEL ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarGearEmergencyHandlePosition Simvar
func SimVarGearEmergencyHandlePosition() SimVar {
	return SimVar{
		Name:     "GEAR EMERGENCY HANDLE POSITION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGearWarning Simvar
func SimVarGearWarning() SimVar {
	return SimVar{
		Name:     "GEAR WARNING",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarAntiskidBrakesActive Simvar
func SimVarAntiskidBrakesActive() SimVar {
	return SimVar{
		Name:     "ANTISKID BRAKES ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRetractFloatSwitch Simvar
func SimVarRetractFloatSwitch() SimVar {
	return SimVar{
		Name:     "RETRACT FLOAT SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRetractLeftFloatExtended Simvar
func SimVarRetractLeftFloatExtended() SimVar {
	return SimVar{
		Name:     "RETRACT LEFT FLOAT EXTENDED",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarRetractRightFloatExtended Simvar
func SimVarRetractRightFloatExtended() SimVar {
	return SimVar{
		Name:     "RETRACT RIGHT FLOAT EXTENDED",
		Units:    "Percent",
		Settable: false,
	}
}

// SimVarSteerInputControl Simvar
func SimVarSteerInputControl() SimVar {
	return SimVar{
		Name:     "STEER INPUT CONTROL",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarAmbientDensity Simvar
func SimVarAmbientDensity() SimVar {
	return SimVar{
		Name:     "AMBIENT DENSITY",
		Units:    "Slugs per cubic feet",
		Settable: false,
	}
}

// SimVarAmbientTemperature Simvar
func SimVarAmbientTemperature() SimVar {
	return SimVar{
		Name:     "AMBIENT TEMPERATURE",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarAmbientPressure Simvar
func SimVarAmbientPressure() SimVar {
	return SimVar{
		Name:     "AMBIENT PRESSURE",
		Units:    "inHg",
		Settable: false,
	}
}

// SimVarAmbientWindVelocity Simvar
func SimVarAmbientWindVelocity() SimVar {
	return SimVar{
		Name:     "AMBIENT WIND VELOCITY",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAmbientWindDirection Simvar
func SimVarAmbientWindDirection() SimVar {
	return SimVar{
		Name:     "AMBIENT WIND DIRECTION",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarAmbientWindX Simvar
func SimVarAmbientWindX() SimVar {
	return SimVar{
		Name:     "AMBIENT WIND X",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientWindY Simvar
func SimVarAmbientWindY() SimVar {
	return SimVar{
		Name:     "AMBIENT WIND Y",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientWindZ Simvar
func SimVarAmbientWindZ() SimVar {
	return SimVar{
		Name:     "AMBIENT WIND Z",
		Units:    "Meters per second",
		Settable: false,
	}
}

// SimVarAmbientPrecipState Simvar
func SimVarAmbientPrecipState() SimVar {
	return SimVar{
		Name:     "AMBIENT PRECIP STATE",
		Units:    "Mask",
		Settable: false,
	}
}

// SimVarAircraftWindX Simvar
func SimVarAircraftWindX() SimVar {
	return SimVar{
		Name:     "AIRCRAFT WIND X",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAircraftWindY Simvar
func SimVarAircraftWindY() SimVar {
	return SimVar{
		Name:     "AIRCRAFT WIND Y",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarAircraftWindZ Simvar
func SimVarAircraftWindZ() SimVar {
	return SimVar{
		Name:     "AIRCRAFT WIND Z",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarBarometerPressure Simvar
func SimVarBarometerPressure() SimVar {
	return SimVar{
		Name:     "BAROMETER PRESSURE",
		Units:    "Millibars",
		Settable: false,
	}
}

// SimVarSeaLevelPressure Simvar
func SimVarSeaLevelPressure() SimVar {
	return SimVar{
		Name:     "SEA LEVEL PRESSURE",
		Units:    "Millibars",
		Settable: false,
	}
}

// SimVarTotalAirTemperature Simvar
func SimVarTotalAirTemperature() SimVar {
	return SimVar{
		Name:     "TOTAL AIR TEMPERATURE",
		Units:    "Celsius",
		Settable: false,
	}
}

// SimVarWindshieldRainEffectAvailable Simvar
func SimVarWindshieldRainEffectAvailable() SimVar {
	return SimVar{
		Name:     "WINDSHIELD RAIN EFFECT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAmbientInCloud Simvar
func SimVarAmbientInCloud() SimVar {
	return SimVar{
		Name:     "AMBIENT IN CLOUD",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAmbientVisibility Simvar
func SimVarAmbientVisibility() SimVar {
	return SimVar{
		Name:     "AMBIENT VISIBILITY",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarStandardAtmTemperature Simvar
func SimVarStandardAtmTemperature() SimVar {
	return SimVar{
		Name:     "STANDARD ATM TEMPERATURE",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarRotorBrakeHandlePos Simvar
func SimVarRotorBrakeHandlePos() SimVar {
	return SimVar{
		Name:     "ROTOR BRAKE HANDLE POS",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRotorBrakeActive Simvar
func SimVarRotorBrakeActive() SimVar {
	return SimVar{
		Name:     "ROTOR BRAKE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorClutchSwitchPos Simvar
func SimVarRotorClutchSwitchPos() SimVar {
	return SimVar{
		Name:     "ROTOR CLUTCH SWITCH POS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorClutchActive Simvar
func SimVarRotorClutchActive() SimVar {
	return SimVar{
		Name:     "ROTOR CLUTCH ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorTemperature Simvar
func SimVarRotorTemperature() SimVar {
	return SimVar{
		Name:     "ROTOR TEMPERATURE",
		Units:    "Rankine",
		Settable: false,
	}
}

// SimVarRotorChipDetected Simvar
func SimVarRotorChipDetected() SimVar {
	return SimVar{
		Name:     "ROTOR CHIP DETECTED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorGovSwitchPos Simvar
func SimVarRotorGovSwitchPos() SimVar {
	return SimVar{
		Name:     "ROTOR GOV SWITCH POS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorGovActive Simvar
func SimVarRotorGovActive() SimVar {
	return SimVar{
		Name:     "ROTOR GOV ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRotorLateralTrimPct Simvar
func SimVarRotorLateralTrimPct() SimVar {
	return SimVar{
		Name:     "ROTOR LATERAL TRIM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarRotorRpmPct Simvar
func SimVarRotorRpmPct() SimVar {
	return SimVar{
		Name:     "ROTOR RPM PCT",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarSmokeEnable Simvar
func SimVarSmokeEnable() SimVar {
	return SimVar{
		Name:     "SMOKE ENABLE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarSmokesystemAvailable Simvar
func SimVarSmokesystemAvailable() SimVar {
	return SimVar{
		Name:     "SMOKESYSTEM AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPitotHeat Simvar
func SimVarPitotHeat() SimVar {
	return SimVar{
		Name:     "PITOT HEAT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFoldingWingLeftPercent Simvar
func SimVarFoldingWingLeftPercent() SimVar {
	return SimVar{
		Name:     "FOLDING WING LEFT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarFoldingWingRightPercent Simvar
func SimVarFoldingWingRightPercent() SimVar {
	return SimVar{
		Name:     "FOLDING WING RIGHT PERCENT",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarCanopyOpen Simvar
func SimVarCanopyOpen() SimVar {
	return SimVar{
		Name:     "CANOPY OPEN",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarTailhookPosition Simvar
func SimVarTailhookPosition() SimVar {
	return SimVar{
		Name:     "TAILHOOK POSITION",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarExitOpen Simvar
func SimVarExitOpen() SimVar {
	return SimVar{
		Name:     "EXIT OPEN:index",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarStallHornAvailable Simvar
func SimVarStallHornAvailable() SimVar {
	return SimVar{
		Name:     "STALL HORN AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarEngineMixureAvailable Simvar
func SimVarEngineMixureAvailable() SimVar {
	return SimVar{
		Name:     "ENGINE MIXURE AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCarbHeatAvailable Simvar
func SimVarCarbHeatAvailable() SimVar {
	return SimVar{
		Name:     "CARB HEAT AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSpoilerAvailable Simvar
func SimVarSpoilerAvailable() SimVar {
	return SimVar{
		Name:     "SPOILER AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsTailDragger Simvar
func SimVarIsTailDragger() SimVar {
	return SimVar{
		Name:     "IS TAIL DRAGGER",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarStrobesAvailable Simvar
func SimVarStrobesAvailable() SimVar {
	return SimVar{
		Name:     "STROBES AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarToeBrakesAvailable Simvar
func SimVarToeBrakesAvailable() SimVar {
	return SimVar{
		Name:     "TOE BRAKES AVAILABLE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPushbackState Simvar
func SimVarPushbackState() SimVar {
	return SimVar{
		Name:     "PUSHBACK STATE",
		Units:    "Enum",
		Settable: true,
	}
}

// SimVarElectricalMasterBattery Simvar
func SimVarElectricalMasterBattery() SimVar {
	return SimVar{
		Name:     "ELECTRICAL MASTER BATTERY",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarElectricalTotalLoadAmps Simvar
func SimVarElectricalTotalLoadAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL TOTAL LOAD AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryLoad Simvar
func SimVarElectricalBatteryLoad() SimVar {
	return SimVar{
		Name:     "ELECTRICAL BATTERY LOAD",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryVoltage Simvar
func SimVarElectricalBatteryVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL BATTERY VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalMainBusVoltage Simvar
func SimVarElectricalMainBusVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL MAIN BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalMainBusAmps Simvar
func SimVarElectricalMainBusAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL MAIN BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusVoltage Simvar
func SimVarElectricalAvionicsBusVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL AVIONICS BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalAvionicsBusAmps Simvar
func SimVarElectricalAvionicsBusAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL AVIONICS BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusVoltage Simvar
func SimVarElectricalHotBatteryBusVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL HOT BATTERY BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalHotBatteryBusAmps Simvar
func SimVarElectricalHotBatteryBusAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL HOT BATTERY BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalBatteryBusVoltage Simvar
func SimVarElectricalBatteryBusVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL BATTERY BUS VOLTAGE",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalBatteryBusAmps Simvar
func SimVarElectricalBatteryBusAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL BATTERY BUS AMPS",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarElectricalGenaltBusVoltage Simvar
func SimVarElectricalGenaltBusVoltage() SimVar {
	return SimVar{
		Name:     "ELECTRICAL GENALT BUS VOLTAGE:index",
		Units:    "Volts",
		Settable: true,
	}
}

// SimVarElectricalGenaltBusAmps Simvar
func SimVarElectricalGenaltBusAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL GENALT BUS AMPS:index",
		Units:    "Amperes",
		Settable: true,
	}
}

// SimVarCircuitGeneralPanelOn Simvar
func SimVarCircuitGeneralPanelOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT GENERAL PANEL ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitFlapMotorOn Simvar
func SimVarCircuitFlapMotorOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT FLAP MOTOR ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitGearMotorOn Simvar
func SimVarCircuitGearMotorOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT GEAR MOTOR ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutopilotOn Simvar
func SimVarCircuitAutopilotOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT AUTOPILOT ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAvionicsOn Simvar
func SimVarCircuitAvionicsOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT AVIONICS ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitPitotHeatOn Simvar
func SimVarCircuitPitotHeatOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT PITOT HEAT ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitPropSyncOn Simvar
func SimVarCircuitPropSyncOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT PROP SYNC ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutoFeatherOn Simvar
func SimVarCircuitAutoFeatherOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT AUTO FEATHER ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitAutoBrakesOn Simvar
func SimVarCircuitAutoBrakesOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT AUTO BRAKES ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitStandyVacuumOn Simvar
func SimVarCircuitStandyVacuumOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT STANDY VACUUM ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitMarkerBeaconOn Simvar
func SimVarCircuitMarkerBeaconOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT MARKER BEACON ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitGearWarningOn Simvar
func SimVarCircuitGearWarningOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT GEAR WARNING ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCircuitHydraulicPumpOn Simvar
func SimVarCircuitHydraulicPumpOn() SimVar {
	return SimVar{
		Name:     "CIRCUIT HYDRAULIC PUMP ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarHydraulicPressure Simvar
func SimVarHydraulicPressure() SimVar {
	return SimVar{
		Name:     "HYDRAULIC PRESSURE:index",
		Units:    "Pound force per square foot",
		Settable: false,
	}
}

// SimVarHydraulicReservoirPercent Simvar
func SimVarHydraulicReservoirPercent() SimVar {
	return SimVar{
		Name:     "HYDRAULIC RESERVOIR PERCENT:index",
		Units:    "Percent Over 100",
		Settable: true,
	}
}

// SimVarHydraulicSystemIntegrity Simvar
func SimVarHydraulicSystemIntegrity() SimVar {
	return SimVar{
		Name:     "HYDRAULIC SYSTEM INTEGRITY",
		Units:    "Percent Over 100",
		Settable: false,
	}
}

// SimVarStructuralDeiceSwitch Simvar
func SimVarStructuralDeiceSwitch() SimVar {
	return SimVar{
		Name:     "STRUCTURAL DEICE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarTotalWeight Simvar
func SimVarTotalWeight() SimVar {
	return SimVar{
		Name:     "TOTAL WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarMaxGrossWeight Simvar
func SimVarMaxGrossWeight() SimVar {
	return SimVar{
		Name:     "MAX GROSS WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarEmptyWeight Simvar
func SimVarEmptyWeight() SimVar {
	return SimVar{
		Name:     "EMPTY WEIGHT",
		Units:    "Pounds",
		Settable: false,
	}
}

// SimVarIsUserSim Simvar
func SimVarIsUserSim() SimVar {
	return SimVar{
		Name:     "IS USER SIM",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSimDisabled Simvar
func SimVarSimDisabled() SimVar {
	return SimVar{
		Name:     "SIM DISABLED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGForce Simvar
func SimVarGForce() SimVar {
	return SimVar{
		Name:     "G FORCE",
		Units:    "GForce",
		Settable: true,
	}
}

// SimVarAtcHeavy Simvar
func SimVarAtcHeavy() SimVar {
	return SimVar{
		Name:     "ATC HEAVY",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAutoCoordination Simvar
func SimVarAutoCoordination() SimVar {
	return SimVar{
		Name:     "AUTO COORDINATION",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarRealism Simvar
func SimVarRealism() SimVar {
	return SimVar{
		Name:     "REALISM",
		Units:    "Number",
		Settable: true,
	}
}

// SimVarTrueAirspeedSelected Simvar
func SimVarTrueAirspeedSelected() SimVar {
	return SimVar{
		Name:     "TRUE AIRSPEED SELECTED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarDesignSpeedVc Simvar
func SimVarDesignSpeedVc() SimVar {
	return SimVar{
		Name:     "DESIGN SPEED VC",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarMinDragVelocity Simvar
func SimVarMinDragVelocity() SimVar {
	return SimVar{
		Name:     "MIN DRAG VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarEstimatedCruiseSpeed Simvar
func SimVarEstimatedCruiseSpeed() SimVar {
	return SimVar{
		Name:     "ESTIMATED CRUISE SPEED",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarCgPercent Simvar
func SimVarCgPercent() SimVar {
	return SimVar{
		Name:     "CG PERCENT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgPercentLateral Simvar
func SimVarCgPercentLateral() SimVar {
	return SimVar{
		Name:     "CG PERCENT LATERAL",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarIsSlewActive Simvar
func SimVarIsSlewActive() SimVar {
	return SimVar{
		Name:     "IS SLEW ACTIVE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarIsSlewAllowed Simvar
func SimVarIsSlewAllowed() SimVar {
	return SimVar{
		Name:     "IS SLEW ALLOWED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarAtcSuggestedMinRwyTakeoff Simvar
func SimVarAtcSuggestedMinRwyTakeoff() SimVar {
	return SimVar{
		Name:     "ATC SUGGESTED MIN RWY TAKEOFF",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarAtcSuggestedMinRwyLanding Simvar
func SimVarAtcSuggestedMinRwyLanding() SimVar {
	return SimVar{
		Name:     "ATC SUGGESTED MIN RWY LANDING",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPayloadStationWeight Simvar
func SimVarPayloadStationWeight() SimVar {
	return SimVar{
		Name:     "PAYLOAD STATION WEIGHT:index",
		Units:    "Pounds",
		Settable: true,
	}
}

// SimVarPayloadStationCount Simvar
func SimVarPayloadStationCount() SimVar {
	return SimVar{
		Name:     "PAYLOAD STATION COUNT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUserInputEnabled Simvar
func SimVarUserInputEnabled() SimVar {
	return SimVar{
		Name:     "USER INPUT ENABLED",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarTypicalDescentRate Simvar
func SimVarTypicalDescentRate() SimVar {
	return SimVar{
		Name:     "TYPICAL DESCENT RATE",
		Units:    "Feet per minute",
		Settable: false,
	}
}

// SimVarVisualModelRadius Simvar
func SimVarVisualModelRadius() SimVar {
	return SimVar{
		Name:     "VISUAL MODEL RADIUS",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarCategory Simvar
func SimVarCategory() SimVar {
	return SimVar{
		Name:     "CATEGORY",
		Units:    "String",
		Settable: false,
	}
}

// SimVarSigmaSqrt Simvar
func SimVarSigmaSqrt() SimVar {
	return SimVar{
		Name:     "SIGMA SQRT",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarDynamicPressure Simvar
func SimVarDynamicPressure() SimVar {
	return SimVar{
		Name:     "DYNAMIC PRESSURE",
		Units:    "Pounds per square foot",
		Settable: false,
	}
}

// SimVarTotalVelocity Simvar
func SimVarTotalVelocity() SimVar {
	return SimVar{
		Name:     "TOTAL VELOCITY",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarAirspeedSelectIndicatedOrTrue Simvar
func SimVarAirspeedSelectIndicatedOrTrue() SimVar {
	return SimVar{
		Name:     "AIRSPEED SELECT INDICATED OR TRUE",
		Units:    "Knots",
		Settable: false,
	}
}

// SimVarVariometerRate Simvar
func SimVarVariometerRate() SimVar {
	return SimVar{
		Name:     "VARIOMETER RATE",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarVariometerSwitch Simvar
func SimVarVariometerSwitch() SimVar {
	return SimVar{
		Name:     "VARIOMETER SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarDesignSpeedVs0 Simvar
func SimVarDesignSpeedVs0() SimVar {
	return SimVar{
		Name:     "DESIGN SPEED VS0",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarDesignSpeedVs1 Simvar
func SimVarDesignSpeedVs1() SimVar {
	return SimVar{
		Name:     "DESIGN SPEED VS1",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPressureAltitude Simvar
func SimVarPressureAltitude() SimVar {
	return SimVar{
		Name:     "PRESSURE ALTITUDE",
		Units:    "Meters",
		Settable: false,
	}
}

// SimVarMagneticCompass Simvar
func SimVarMagneticCompass() SimVar {
	return SimVar{
		Name:     "MAGNETIC COMPASS",
		Units:    "Degrees",
		Settable: false,
	}
}

// SimVarTurnIndicatorRate Simvar
func SimVarTurnIndicatorRate() SimVar {
	return SimVar{
		Name:     "TURN INDICATOR RATE",
		Units:    "Radians per second",
		Settable: false,
	}
}

// SimVarTurnIndicatorSwitch Simvar
func SimVarTurnIndicatorSwitch() SimVar {
	return SimVar{
		Name:     "TURN INDICATOR SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarYokeYIndicator Simvar
func SimVarYokeYIndicator() SimVar {
	return SimVar{
		Name:     "YOKE Y INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarYokeXIndicator Simvar
func SimVarYokeXIndicator() SimVar {
	return SimVar{
		Name:     "YOKE X INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarRudderPedalIndicator Simvar
func SimVarRudderPedalIndicator() SimVar {
	return SimVar{
		Name:     "RUDDER PEDAL INDICATOR",
		Units:    "Position",
		Settable: false,
	}
}

// SimVarBrakeDependentHydraulicPressure Simvar
func SimVarBrakeDependentHydraulicPressure() SimVar {
	return SimVar{
		Name:     "BRAKE DEPENDENT HYDRAULIC PRESSURE",
		Units:    "foot pounds",
		Settable: false,
	}
}

// SimVarPanelAntiIceSwitch Simvar
func SimVarPanelAntiIceSwitch() SimVar {
	return SimVar{
		Name:     "PANEL ANTI ICE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarWingArea Simvar
func SimVarWingArea() SimVar {
	return SimVar{
		Name:     "WING AREA",
		Units:    "Square feet",
		Settable: false,
	}
}

// SimVarWingSpan Simvar
func SimVarWingSpan() SimVar {
	return SimVar{
		Name:     "WING SPAN",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarBetaDot Simvar
func SimVarBetaDot() SimVar {
	return SimVar{
		Name:     "BETA DOT",
		Units:    "Radians per second",
		Settable: false,
	}
}

// SimVarLinearClAlpha Simvar
func SimVarLinearClAlpha() SimVar {
	return SimVar{
		Name:     "LINEAR CL ALPHA",
		Units:    "Per radian",
		Settable: false,
	}
}

// SimVarStallAlpha Simvar
func SimVarStallAlpha() SimVar {
	return SimVar{
		Name:     "STALL ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarZeroLiftAlpha Simvar
func SimVarZeroLiftAlpha() SimVar {
	return SimVar{
		Name:     "ZERO LIFT ALPHA",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCgAftLimit Simvar
func SimVarCgAftLimit() SimVar {
	return SimVar{
		Name:     "CG AFT LIMIT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgFwdLimit Simvar
func SimVarCgFwdLimit() SimVar {
	return SimVar{
		Name:     "CG FWD LIMIT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarCgMaxMach Simvar
func SimVarCgMaxMach() SimVar {
	return SimVar{
		Name:     "CG MAX MACH",
		Units:    "Machs",
		Settable: false,
	}
}

// SimVarCgMinMach Simvar
func SimVarCgMinMach() SimVar {
	return SimVar{
		Name:     "CG MIN MACH",
		Units:    "Machs",
		Settable: false,
	}
}

// SimVarPayloadStationName Simvar
func SimVarPayloadStationName() SimVar {
	return SimVar{
		Name:     "PAYLOAD STATION NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarElevonDeflection Simvar
func SimVarElevonDeflection() SimVar {
	return SimVar{
		Name:     "ELEVON DEFLECTION",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarExitType Simvar
func SimVarExitType() SimVar {
	return SimVar{
		Name:     "EXIT TYPE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarExitPosx Simvar
func SimVarExitPosx() SimVar {
	return SimVar{
		Name:     "EXIT POSX",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarExitPosy Simvar
func SimVarExitPosy() SimVar {
	return SimVar{
		Name:     "EXIT POSY",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarExitPosz Simvar
func SimVarExitPosz() SimVar {
	return SimVar{
		Name:     "EXIT POSZ",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarDecisionHeight Simvar
func SimVarDecisionHeight() SimVar {
	return SimVar{
		Name:     "DECISION HEIGHT",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarDecisionAltitudeMsl Simvar
func SimVarDecisionAltitudeMsl() SimVar {
	return SimVar{
		Name:     "DECISION ALTITUDE MSL",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarEmptyWeightPitchMoi Simvar
func SimVarEmptyWeightPitchMoi() SimVar {
	return SimVar{
		Name:     "EMPTY WEIGHT PITCH MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightRollMoi Simvar
func SimVarEmptyWeightRollMoi() SimVar {
	return SimVar{
		Name:     "EMPTY WEIGHT ROLL MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightYawMoi Simvar
func SimVarEmptyWeightYawMoi() SimVar {
	return SimVar{
		Name:     "EMPTY WEIGHT YAW MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarEmptyWeightCrossCoupledMoi Simvar
func SimVarEmptyWeightCrossCoupledMoi() SimVar {
	return SimVar{
		Name:     "EMPTY WEIGHT CROSS COUPLED MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightPitchMoi Simvar
func SimVarTotalWeightPitchMoi() SimVar {
	return SimVar{
		Name:     "TOTAL WEIGHT PITCH MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightRollMoi Simvar
func SimVarTotalWeightRollMoi() SimVar {
	return SimVar{
		Name:     "TOTAL WEIGHT ROLL MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightYawMoi Simvar
func SimVarTotalWeightYawMoi() SimVar {
	return SimVar{
		Name:     "TOTAL WEIGHT YAW MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarTotalWeightCrossCoupledMoi Simvar
func SimVarTotalWeightCrossCoupledMoi() SimVar {
	return SimVar{
		Name:     "TOTAL WEIGHT CROSS COUPLED MOI",
		Units:    "slug feet squared",
		Settable: false,
	}
}

// SimVarWaterBallastValve Simvar
func SimVarWaterBallastValve() SimVar {
	return SimVar{
		Name:     "WATER BALLAST VALVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarMaxRatedEngineRpm Simvar
func SimVarMaxRatedEngineRpm() SimVar {
	return SimVar{
		Name:     "MAX RATED ENGINE RPM",
		Units:    "Rpm",
		Settable: false,
	}
}

// SimVarFullThrottleThrustToWeightRatio Simvar
func SimVarFullThrottleThrustToWeightRatio() SimVar {
	return SimVar{
		Name:     "FULL THROTTLE THRUST TO WEIGHT RATIO",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarPropAutoCruiseActive Simvar
func SimVarPropAutoCruiseActive() SimVar {
	return SimVar{
		Name:     "PROP AUTO CRUISE ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPropRotationAngle Simvar
func SimVarPropRotationAngle() SimVar {
	return SimVar{
		Name:     "PROP ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMax Simvar
func SimVarPropBetaMax() SimVar {
	return SimVar{
		Name:     "PROP BETA MAX",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMin Simvar
func SimVarPropBetaMin() SimVar {
	return SimVar{
		Name:     "PROP BETA MIN",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPropBetaMinReverse Simvar
func SimVarPropBetaMinReverse() SimVar {
	return SimVar{
		Name:     "PROP BETA MIN REVERSE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarFuelSelectedTransferMode Simvar
func SimVarFuelSelectedTransferMode() SimVar {
	return SimVar{
		Name:     "FUEL SELECTED TRANSFER MODE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarDroppableObjectsUiName Simvar
func SimVarDroppableObjectsUiName() SimVar {
	return SimVar{
		Name:     "DROPPABLE OBJECTS UI NAME",
		Units:    "String",
		Settable: false,
	}
}

// SimVarManualFuelPumpHandle Simvar
func SimVarManualFuelPumpHandle() SimVar {
	return SimVar{
		Name:     "MANUAL FUEL PUMP HANDLE",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarBleedAirSourceControl Simvar
func SimVarBleedAirSourceControl() SimVar {
	return SimVar{
		Name:     "BLEED AIR SOURCE CONTROL",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarElectricalOldChargingAmps Simvar
func SimVarElectricalOldChargingAmps() SimVar {
	return SimVar{
		Name:     "ELECTRICAL OLD CHARGING AMPS",
		Units:    "Amps",
		Settable: false,
	}
}

// SimVarHydraulicSwitch Simvar
func SimVarHydraulicSwitch() SimVar {
	return SimVar{
		Name:     "HYDRAULIC SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarConcordeVisorNoseHandle Simvar
func SimVarConcordeVisorNoseHandle() SimVar {
	return SimVar{
		Name:     "CONCORDE VISOR NOSE HANDLE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarConcordeVisorPositionPercent Simvar
func SimVarConcordeVisorPositionPercent() SimVar {
	return SimVar{
		Name:     "CONCORDE VISOR POSITION PERCENT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarConcordeNoseAngle Simvar
func SimVarConcordeNoseAngle() SimVar {
	return SimVar{
		Name:     "CONCORDE NOSE ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarRealismCrashWithOthers Simvar
func SimVarRealismCrashWithOthers() SimVar {
	return SimVar{
		Name:     "REALISM CRASH WITH OTHERS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarRealismCrashDetection Simvar
func SimVarRealismCrashDetection() SimVar {
	return SimVar{
		Name:     "REALISM CRASH DETECTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarManualInstrumentLights Simvar
func SimVarManualInstrumentLights() SimVar {
	return SimVar{
		Name:     "MANUAL INSTRUMENT LIGHTS",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPitotIcePct Simvar
func SimVarPitotIcePct() SimVar {
	return SimVar{
		Name:     "PITOT ICE PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorY Simvar
func SimVarSemibodyLoadfactorY() SimVar {
	return SimVar{
		Name:     "SEMIBODY LOADFACTOR Y",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarSemibodyLoadfactorYdot Simvar
func SimVarSemibodyLoadfactorYdot() SimVar {
	return SimVar{
		Name:     "SEMIBODY LOADFACTOR YDOT",
		Units:    "Per second",
		Settable: false,
	}
}

// SimVarRadInsSwitch Simvar
func SimVarRadInsSwitch() SimVar {
	return SimVar{
		Name:     "RAD INS SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSimulatedRadius Simvar
func SimVarSimulatedRadius() SimVar {
	return SimVar{
		Name:     "SIMULATED RADIUS",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarStructuralIcePct Simvar
func SimVarStructuralIcePct() SimVar {
	return SimVar{
		Name:     "STRUCTURAL ICE PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarArtificialGroundElevation Simvar
func SimVarArtificialGroundElevation() SimVar {
	return SimVar{
		Name:     "ARTIFICIAL GROUND ELEVATION",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarSurfaceInfoValid Simvar
func SimVarSurfaceInfoValid() SimVar {
	return SimVar{
		Name:     "SURFACE INFO VALID",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarSurfaceCondition Simvar
func SimVarSurfaceCondition() SimVar {
	return SimVar{
		Name:     "SURFACE CONDITION",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarPushbackAngle Simvar
func SimVarPushbackAngle() SimVar {
	return SimVar{
		Name:     "PUSHBACK ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarPushbackContactx Simvar
func SimVarPushbackContactx() SimVar {
	return SimVar{
		Name:     "PUSHBACK CONTACTX",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackContacty Simvar
func SimVarPushbackContacty() SimVar {
	return SimVar{
		Name:     "PUSHBACK CONTACTY",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackContactz Simvar
func SimVarPushbackContactz() SimVar {
	return SimVar{
		Name:     "PUSHBACK CONTACTZ",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPushbackWait Simvar
func SimVarPushbackWait() SimVar {
	return SimVar{
		Name:     "PUSHBACK WAIT",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarYawStringAngle Simvar
func SimVarYawStringAngle() SimVar {
	return SimVar{
		Name:     "YAW STRING ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarYawStringPctExtended Simvar
func SimVarYawStringPctExtended() SimVar {
	return SimVar{
		Name:     "YAW STRING PCT EXTENDED",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarInductorCompassPercentDeviation Simvar
func SimVarInductorCompassPercentDeviation() SimVar {
	return SimVar{
		Name:     "INDUCTOR COMPASS PERCENT DEVIATION",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarInductorCompassHeadingRef Simvar
func SimVarInductorCompassHeadingRef() SimVar {
	return SimVar{
		Name:     "INDUCTOR COMPASS HEADING REF",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarAnemometerPctRpm Simvar
func SimVarAnemometerPctRpm() SimVar {
	return SimVar{
		Name:     "ANEMOMETER PCT RPM",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarRotorRotationAngle Simvar
func SimVarRotorRotationAngle() SimVar {
	return SimVar{
		Name:     "ROTOR ROTATION ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskPitchAngle Simvar
func SimVarDiskPitchAngle() SimVar {
	return SimVar{
		Name:     "DISK PITCH ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskBankAngle Simvar
func SimVarDiskBankAngle() SimVar {
	return SimVar{
		Name:     "DISK BANK ANGLE",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarDiskPitchPct Simvar
func SimVarDiskPitchPct() SimVar {
	return SimVar{
		Name:     "DISK PITCH PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarDiskBankPct Simvar
func SimVarDiskBankPct() SimVar {
	return SimVar{
		Name:     "DISK BANK PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarDiskConingPct Simvar
func SimVarDiskConingPct() SimVar {
	return SimVar{
		Name:     "DISK CONING PCT",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarNavVorLlaf64 Simvar
func SimVarNavVorLlaf64() SimVar {
	return SimVar{
		Name:     "NAV VOR LLAF64",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarNavGsLlaf64 Simvar
func SimVarNavGsLlaf64() SimVar {
	return SimVar{
		Name:     "NAV GS LLAF64",
		Units:    "SIMCONNECT_DATA_LATLONALT",
		Settable: false,
	}
}

// SimVarStaticCgToGround Simvar
func SimVarStaticCgToGround() SimVar {
	return SimVar{
		Name:     "STATIC CG TO GROUND",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarStaticPitch Simvar
func SimVarStaticPitch() SimVar {
	return SimVar{
		Name:     "STATIC PITCH",
		Units:    "Radians",
		Settable: false,
	}
}

// SimVarCrashSequence Simvar
func SimVarCrashSequence() SimVar {
	return SimVar{
		Name:     "CRASH SEQUENCE",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarCrashFlag Simvar
func SimVarCrashFlag() SimVar {
	return SimVar{
		Name:     "CRASH FLAG",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarTowReleaseHandle Simvar
func SimVarTowReleaseHandle() SimVar {
	return SimVar{
		Name:     "TOW RELEASE HANDLE",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarTowConnection Simvar
func SimVarTowConnection() SimVar {
	return SimVar{
		Name:     "TOW CONNECTION",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuPctRpm Simvar
func SimVarApuPctRpm() SimVar {
	return SimVar{
		Name:     "APU PCT RPM",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarApuPctStarter Simvar
func SimVarApuPctStarter() SimVar {
	return SimVar{
		Name:     "APU PCT STARTER",
		Units:    "Percent over 100",
		Settable: false,
	}
}

// SimVarApuVolts Simvar
func SimVarApuVolts() SimVar {
	return SimVar{
		Name:     "APU VOLTS",
		Units:    "Volts",
		Settable: false,
	}
}

// SimVarApuGeneratorSwitch Simvar
func SimVarApuGeneratorSwitch() SimVar {
	return SimVar{
		Name:     "APU GENERATOR SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuGeneratorActive Simvar
func SimVarApuGeneratorActive() SimVar {
	return SimVar{
		Name:     "APU GENERATOR ACTIVE",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarApuOnFireDetected Simvar
func SimVarApuOnFireDetected() SimVar {
	return SimVar{
		Name:     "APU ON FIRE DETECTED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitude Simvar
func SimVarPressurizationCabinAltitude() SimVar {
	return SimVar{
		Name:     "PRESSURIZATION CABIN ALTITUDE",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeGoal Simvar
func SimVarPressurizationCabinAltitudeGoal() SimVar {
	return SimVar{
		Name:     "PRESSURIZATION CABIN ALTITUDE GOAL",
		Units:    "Feet",
		Settable: false,
	}
}

// SimVarPressurizationCabinAltitudeRate Simvar
func SimVarPressurizationCabinAltitudeRate() SimVar {
	return SimVar{
		Name:     "PRESSURIZATION CABIN ALTITUDE RATE",
		Units:    "Feet per second",
		Settable: false,
	}
}

// SimVarPressurizationPressureDifferential Simvar
func SimVarPressurizationPressureDifferential() SimVar {
	return SimVar{
		Name:     "PRESSURIZATION PRESSURE DIFFERENTIAL",
		Units:    "foot pounds",
		Settable: false,
	}
}

// SimVarPressurizationDumpSwitch Simvar
func SimVarPressurizationDumpSwitch() SimVar {
	return SimVar{
		Name:     "PRESSURIZATION DUMP SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFireBottleSwitch Simvar
func SimVarFireBottleSwitch() SimVar {
	return SimVar{
		Name:     "FIRE BOTTLE SWITCH",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarFireBottleDischarged Simvar
func SimVarFireBottleDischarged() SimVar {
	return SimVar{
		Name:     "FIRE BOTTLE DISCHARGED",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarCabinNoSmokingAlertSwitch Simvar
func SimVarCabinNoSmokingAlertSwitch() SimVar {
	return SimVar{
		Name:     "CABIN NO SMOKING ALERT SWITCH",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarCabinSeatbeltsAlertSwitch Simvar
func SimVarCabinSeatbeltsAlertSwitch() SimVar {
	return SimVar{
		Name:     "CABIN SEATBELTS ALERT SWITCH",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarGpwsWarning Simvar
func SimVarGpwsWarning() SimVar {
	return SimVar{
		Name:     "GPWS WARNING",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarGpwsSystemActive Simvar
func SimVarGpwsSystemActive() SimVar {
	return SimVar{
		Name:     "GPWS SYSTEM ACTIVE",
		Units:    "Bool",
		Settable: true,
	}
}

// SimVarIsLatitudeLongitudeFreezeOn Simvar
func SimVarIsLatitudeLongitudeFreezeOn() SimVar {
	return SimVar{
		Name:     "IS LATITUDE LONGITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAltitudeFreezeOn Simvar
func SimVarIsAltitudeFreezeOn() SimVar {
	return SimVar{
		Name:     "IS ALTITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarIsAttitudeFreezeOn Simvar
func SimVarIsAttitudeFreezeOn() SimVar {
	return SimVar{
		Name:     "IS ATTITUDE FREEZE ON",
		Units:    "Bool",
		Settable: false,
	}
}

// SimVarAtcType Simvar
func SimVarAtcType() SimVar {
	return SimVar{
		Name:     "ATC TYPE",
		Units:    "String64",
		Settable: false,
	}
}

// SimVarAtcModel Simvar
func SimVarAtcModel() SimVar {
	return SimVar{
		Name:     "ATC MODEL",
		Units:    "String64",
		Settable: false,
	}
}

// SimVarAtcId Simvar
func SimVarAtcId() SimVar {
	return SimVar{
		Name:     "ATC ID",
		Units:    "String64",
		Settable: true,
	}
}

// SimVarAtcAirline Simvar
func SimVarAtcAirline() SimVar {
	return SimVar{
		Name:     "ATC AIRLINE",
		Units:    "String64",
		Settable: true,
	}
}

// SimVarAtcFlightNumber Simvar
func SimVarAtcFlightNumber() SimVar {
	return SimVar{
		Name:     "ATC FLIGHT NUMBER",
		Units:    "String8",
		Settable: true,
	}
}

// SimVarTitle Simvar
func SimVarTitle() SimVar {
	return SimVar{
		Name:     "TITLE",
		Units:    "Variable length string",
		Settable: false,
	}
}

// SimVarHsiStationIdent Simvar
func SimVarHsiStationIdent() SimVar {
	return SimVar{
		Name:     "HSI STATION IDENT",
		Units:    "String8",
		Settable: false,
	}
}

// SimVarGpsApproachAirportId Simvar
func SimVarGpsApproachAirportId() SimVar {
	return SimVar{
		Name:     "GPS APPROACH AIRPORT ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsApproachApproachId Simvar
func SimVarGpsApproachApproachId() SimVar {
	return SimVar{
		Name:     "GPS APPROACH APPROACH ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarGpsApproachTransitionId Simvar
func SimVarGpsApproachTransitionId() SimVar {
	return SimVar{
		Name:     "GPS APPROACH TRANSITION ID",
		Units:    "String",
		Settable: false,
	}
}

// SimVarAbsoluteTime Simvar
func SimVarAbsoluteTime() SimVar {
	return SimVar{
		Name:     "ABSOLUTE TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarZuluTime Simvar
func SimVarZuluTime() SimVar {
	return SimVar{
		Name:     "ZULU TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarZuluDayOfWeek Simvar
func SimVarZuluDayOfWeek() SimVar {
	return SimVar{
		Name:     "ZULU DAY OF WEEK",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluDayOfMonth Simvar
func SimVarZuluDayOfMonth() SimVar {
	return SimVar{
		Name:     "ZULU DAY OF MONTH",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluMonthOfYear Simvar
func SimVarZuluMonthOfYear() SimVar {
	return SimVar{
		Name:     "ZULU MONTH OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluDayOfYear Simvar
func SimVarZuluDayOfYear() SimVar {
	return SimVar{
		Name:     "ZULU DAY OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarZuluYear Simvar
func SimVarZuluYear() SimVar {
	return SimVar{
		Name:     "ZULU YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalTime Simvar
func SimVarLocalTime() SimVar {
	return SimVar{
		Name:     "LOCAL TIME",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarLocalDayOfWeek Simvar
func SimVarLocalDayOfWeek() SimVar {
	return SimVar{
		Name:     "LOCAL DAY OF WEEK",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalDayOfMonth Simvar
func SimVarLocalDayOfMonth() SimVar {
	return SimVar{
		Name:     "LOCAL DAY OF MONTH",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalMonthOfYear Simvar
func SimVarLocalMonthOfYear() SimVar {
	return SimVar{
		Name:     "LOCAL MONTH OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalDayOfYear Simvar
func SimVarLocalDayOfYear() SimVar {
	return SimVar{
		Name:     "LOCAL DAY OF YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarLocalYear Simvar
func SimVarLocalYear() SimVar {
	return SimVar{
		Name:     "LOCAL YEAR",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarTimeZoneOffset Simvar
func SimVarTimeZoneOffset() SimVar {
	return SimVar{
		Name:     "TIME ZONE OFFSET",
		Units:    "Seconds",
		Settable: false,
	}
}

// SimVarTimeOfDay Simvar
func SimVarTimeOfDay() SimVar {
	return SimVar{
		Name:     "TIME OF DAY",
		Units:    "Enum",
		Settable: false,
	}
}

// SimVarSimulationRate Simvar
func SimVarSimulationRate() SimVar {
	return SimVar{
		Name:     "SIMULATION RATE",
		Units:    "Number",
		Settable: false,
	}
}

// SimVarUnitsOfMeasure Simvar
func SimVarUnitsOfMeasure() SimVar {
	return SimVar{
		Name:     "UNITS OF MEASURE",
		Units:    "Enum",
		Settable: false,
	}
}
