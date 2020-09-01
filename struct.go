package simconnect

type SIMCONNECT_RECV struct {
	dwSize    uint32 // record size
	dwVersuib uint32 // interface version
	dwID      uint32 // see SIMCONNECT_RECV_ID
}

type SIMCONNECT_RECV_EXCEPTION struct {
	SIMCONNECT_RECV
	dwException uint32 // see SIMCONNECT_EXCEPTION
	dwSendID    uint32 // see SimConnect_GetLastSentPacketID
	dwIndex     uint32 // index of parameter that was source of error
}

type SIMCONNECT_RECV_OPEN struct {
	SIMCONNECT_RECV
	szApplicationName         [256]byte
	dwApplicationVersionMajor uint32
	dwApplicationVersionMinor uint32
	dwApplicationBuildMajor   uint32
	dwApplicationBuildMinor   uint32
	dwSimConnectVersionMajor  uint32
	dwSimConnectVersionMinor  uint32
	dwSimConnectBuildMajor    uint32
	dwSimConnectBuildMinor    uint32
	dwReserved1               uint32
	dwReserved2               uint32
}

type SIMCONNECT_RECV_QUIT struct {
	SIMCONNECT_RECV
}

type SIMCONNECT_RECV_EVENT struct {
	SIMCONNECT_RECV
	uGroupID uint32
	uEventID uint32
	dwData   uint32 // uEventID-dependent context
}

type SIMCONNECT_RECV_EVENT_FILENAME struct {
	SIMCONNECT_RECV_EVENT
	szFileName [260]byte // depend of MAX_PATH
	dwFlags    uint32
}

type SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE struct {
	SIMCONNECT_RECV_EVENT
	eObjType uint32
}

type SIMCONNECT_RECV_EVENT_FRAME struct {
	SIMCONNECT_RECV_EVENT
	fFrameRate float32
	fSimSpeed  float32
}

type SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED struct {
	SIMCONNECT_RECV_EVENT
}

type SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED struct {
	SIMCONNECT_RECV_EVENT
}

type SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED struct {
	SIMCONNECT_RECV_EVENT
}

type GUID struct {
	Data1 uint64
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type SIMCONNECT_DATA_RACE_RESULT struct {
	dwNumberOfRacers uint32    // The total number of racers
	MissionGUID      *GUID     // The name of the mission to execute, NULL if no mission
	szPlayerName     [260]byte // The name of the player (depend of MAX_PATH)
	szSessionType    [260]byte // The type of the multiplayer session: "LAN", "GAMESPY") (depend of MAX_PATH)
	szAircraft       [260]byte // The aircraft type (depend of MAX_PATH)
	szPlayerRole     [260]byte // The player role in the mission (depend of MAX_PATH)
	fTotalTime       float64   // Total time in seconds, 0 means DNF
	fPenaltyTime     float64   // Total penalty time in seconds
	dwIsDisqualified uint32    // non 0 - disqualified, 0 - not disqualified
}

type SIMCONNECT_RECV_EVENT_RACE_END struct {
	SIMCONNECT_RECV_EVENT
	dwRacerNumber uint32 // The index of the racer the results are for
	RacerData     SIMCONNECT_DATA_RACE_RESULT
}

type SIMCONNECT_RECV_EVENT_RACE_LAP struct {
	SIMCONNECT_RECV_EVENT
	dwRacerNumber uint32 // The index of the lap the results are for
	RacerData     SIMCONNECT_DATA_RACE_RESULT
}

type SIMCONNECT_RECV_SIMOBJECT_DATA struct {
	SIMCONNECT_RECV
	dwRequestID   uint32
	dwObjectID    uint32
	dwDefineID    uint32
	dwFlags       uint32 // SIMCONNECT_DATA_REQUEST_FLAG
	dwentrynumber uint32 // if multiple objects returned, this is number <entrynumber> out of <outof>.
	dwoutof       uint32 // note: starts with 1, not 0.
	dwDefineCount uint32 // data count (number of datums, *not* byte count)
	dwData        []byte
}

type SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE struct {
	SIMCONNECT_RECV_SIMOBJECT_DATA
}

type SIMCONNECT_RECV_CLIENT_DATA struct {
	SIMCONNECT_RECV_SIMOBJECT_DATA
}

type SIMCONNECT_RECV_WEATHER_OBSERVATION struct {
	SIMCONNECT_RECV
	dwRequestID uint32
	szMetar     [1]byte // Variable length string whose maximum size is MAX_METAR_LENGTH
}

const (
	SIMCONNECT_CLOUD_STATE_ARRAY_WIDTH = 64
	SIMCONNECT_CLOUD_STATE_ARRAY_SIZE  = SIMCONNECT_CLOUD_STATE_ARRAY_WIDTH * SIMCONNECT_CLOUD_STATE_ARRAY_WIDTH
)

// when dwID == SIMCONNECT_RECV_ID_CLOUD_STATE
type SIMCONNECT_RECV_CLOUD_STATE struct {
	SIMCONNECT_RECV
	dwRequestID uint32
	dwArraySize uint32
	rgbData     [1]byte //variable size
}

// when dwID == SIMCONNECT_RECV_ID_ASSIGNED_OBJECT_ID
type SIMCONNECT_RECV_ASSIGNED_OBJECT_ID struct {
	SIMCONNECT_RECV
	dwRequestID uint32
	dwObjectID  uint32
}

// when dwID == SIMCONNECT_RECV_ID_RESERVED_KEY
type SIMCONNECT_RECV_RESERVED_KEY struct {
	SIMCONNECT_RECV
	szChoiceReserved [30]byte
	szReservedKey    [50]byte
}

// when dwID == SIMCONNECT_RECV_ID_SYSTEM_STATE
type SIMCONNECT_RECV_SYSTEM_STATE struct {
	SIMCONNECT_RECV
	dwRequestID uint32
	dwInteger   uint32
	fFloat      float32
	szString    [260]byte // depend of MAX_PATH
}

type SIMCONNECT_RECV_CUSTOM_ACTION struct {
	SIMCONNECT_RECV_EVENT
	guidInstanceId      GUID
	dwWaitForCompletion uint32
	szPayLoad           [1]byte
}

type SIMCONNECT_RECV_EVENT_WEATHER_MODE struct {
	SIMCONNECT_RECV_EVENT
}

type SIMCONNECT_RECV_FACILITIES_LIST struct {
	SIMCONNECT_RECV
	dwRequestID   uint32
	dwArraySize   uint32
	dwEntryNumber uint32 // when the array of items is too big for one send, which send this is (0..dwOutOf-1)
	dwOutOf       uint32 // total number of transmissions the list is chopped into
}

type SIMCONNECT_DATA_FACILITY_AIRPORT struct {
	Icao      [9]byte // ICAO of the object
	Latitude  float64 // degrees
	Longitude float64 // degrees
	Altitude  float64 // meters
}

type SIMCONNECT_RECV_AIRPORT_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	rgData [1]SIMCONNECT_DATA_FACILITY_AIRPORT
}

type SIMCONNECT_DATA_FACILITY_WAYPOINT struct {
	SIMCONNECT_DATA_FACILITY_AIRPORT
	fMagVar float32 // Magvar in degrees
}

type SIMCONNECT_RECV_WAYPOINT_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	rgData [1]SIMCONNECT_DATA_FACILITY_WAYPOINT
}

type SIMCONNECT_DATA_FACILITY_NDB struct {
	SIMCONNECT_DATA_FACILITY_WAYPOINT
	fFrequency uint32 // frequency in Hz
}

type SIMCONNECT_RECV_NDB_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	rgData [1]SIMCONNECT_DATA_FACILITY_NDB
}

type SIMCONNECT_DATA_FACILITY_VOR struct {
	SIMCONNECT_DATA_FACILITY_NDB
	Flags            uint32  // SIMCONNECT_VOR_FLAGS
	fLocalizer       float32 // Localizer in degrees
	GlideLat         float64 // Glide Slope Location (deg, deg, meters)
	GlideLon         float64
	GlideAlt         float64
	fGlideSlopeAngle float32 // Glide Slope in degrees
}

type SIMCONNECT_RECV_VOR_LIST struct {
	SIMCONNECT_RECV_FACILITIES_LIST
	rgData [1]SIMCONNECT_DATA_FACILITY_VOR
}

type SIMCONNECT_DATA_INITPOSITION struct {
	Latitude  float64 // degrees
	Longitude float64 // degrees
	Altitude  float64 // feet
	Pitch     float64 // degrees
	Bank      float64 // degrees
	Heading   float64 // degrees
	OnGround  uint32  // 1=force to be on the ground
	Airspeed  uint32  // knots
}

type SIMCONNECT_DATA_MARKERSTATE struct {
	szMarkerName  [64]byte
	dwMarkerState uint32
}

type SIMCONNECT_DATA_WAYPOINT struct {
	Latitude        float64 // degrees
	Longitude       float64 // degrees
	Altitude        float64 // feet
	Flags           uint64
	ktsSpeed        float64 // knots
	percentThrottle float64
}

type SIMCONNECT_DATA_LATLONALT struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type SIMCONNECT_DATA_XYZ struct {
	x float64
	y float64
	z float64
}
