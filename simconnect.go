package simconnect

import (
	"errors"
	"unsafe"
)

// convert string to const char *
func cChar(str string) uintptr {
	b := []byte(str + "\x00")
	return uintptr(unsafe.Pointer(&b[0]))
}

// SimConnect golang interface
type SimConnect struct {
	hSimConnect uintptr
	syscallSC   *SyscallSC
}

// NewSimConnect get instance of SimConnect
func NewSimConnect() (*SimConnect, error) {
	var err error
	simConnect := &SimConnect{}
	simConnect.syscallSC, err = NewSyscallSC()
	if err != nil {
		return nil, err
	}
	return simConnect, nil
}

// MapClientEventToSimEvent SimConnect_MapClientEventToSimEvent(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID EventID, const char * EventName = "")
func (sc *SimConnect) MapClientEventToSimEvent(EventID uint32, EventName string) error {
	return errors.New("not implemented")
}

// TransmitClientEvent SimConnect_TransmitClientEvent(HANDLE hSimConnect, SIMCONNECT_OBJECT_ID ObjectID, SIMCONNECT_CLIENT_EVENT_ID EventID, DWORD dwData, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID, SIMCONNECT_EVENT_FLAG Flags);
func (sc *SimConnect) TransmitClientEvent(ObjectID uint32, EventID uint32, dwData uint32, GroupID uint32, Flags uint32) error {
	return errors.New("not implemented")
}

// SetSystemEventState SimConnect_SetSystemEventState(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID EventID, SIMCONNECT_STATE dwState);
func (sc *SimConnect) SetSystemEventState(EventID uint32, dwState uint32) error {
	return errors.New("not implemented")
}

// AddClientEventToNotificationGroup SimConnect_AddClientEventToNotificationGroup(HANDLE hSimConnect, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID, SIMCONNECT_CLIENT_EVENT_ID EventID, BOOL bMaskable = FALSE);
func (sc *SimConnect) AddClientEventToNotificationGroup(GroupID uint32, EventID uint32, bMaskable uint32) error {
	return errors.New("not implemented")
}

// RemoveClientEvent SimConnect_RemoveClientEvent(HANDLE hSimConnect, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID, SIMCONNECT_CLIENT_EVENT_ID EventID);
func (sc *SimConnect) RemoveClientEvent(GroupID uint32, EventID uint32) error {
	return errors.New("not implemented")
}

// SetNotificationGroupPriority SimConnect_SetNotificationGroupPriority(HANDLE hSimConnect, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID, DWORD uPriority);
func (sc *SimConnect) SetNotificationGroupPriority(GroupID uint32, uPriority uint32) error {
	return errors.New("not implemented")
}

// ClearNotificationGroup SimConnect_ClearNotificationGroup(HANDLE hSimConnect, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID);
func (sc *SimConnect) ClearNotificationGroup(GroupID uint32) error {
	return errors.New("not implemented")
}

// RequestNotificationGroup SimConnect_RequestNotificationGroup(HANDLE hSimConnect, SIMCONNECT_NOTIFICATION_GROUP_ID GroupID, DWORD dwReserved = 0, DWORD Flags = 0);
func (sc *SimConnect) RequestNotificationGroup(GroupID uint32, dwReserved uint32, Flags uint32) error {
	return errors.New("not implemented")
}

// AddToDataDefinition SimConnect_AddToDataDefinition(HANDLE hSimConnect, SIMCONNECT_DATA_DEFINITION_ID DefineID, const char * DatumName, const char * UnitsName, SIMCONNECT_DATATYPE DatumType = SIMCONNECT_DATATYPE_FLOAT64, float fEpsilon = 0, DWORD DatumID = SIMCONNECT_UNUSED);
func (sc *SimConnect) AddToDataDefinition(DefineID uint32, DatumName string, UnitsName string, DatumType uint32, fEpsilon float32, DatumID uint32) error {
	err := sc.syscallSC.AddToDataDefinition(sc.hSimConnect, uintptr(DefineID), cChar(DatumName), cChar(UnitsName), uintptr(DatumType), uintptr(fEpsilon), uintptr(DatumID))
	return err
}

// ClearDataDefinition SimConnect_ClearDataDefinition(HANDLE hSimConnect, SIMCONNECT_DATA_DEFINITION_ID DefineID);
func (sc *SimConnect) ClearDataDefinition(DefineID uint32) error {
	return sc.syscallSC.ClearDataDefinition(sc.hSimConnect, uintptr(DefineID))
}

// RequestDataOnSimObject SimConnect_RequestDataOnSimObject(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, SIMCONNECT_DATA_DEFINITION_ID DefineID, SIMCONNECT_OBJECT_ID ObjectID, SIMCONNECT_PERIOD Period, SIMCONNECT_DATA_REQUEST_FLAG Flags = 0, DWORD origin = 0, DWORD interval = 0, DWORD limit = 0);
func (sc *SimConnect) RequestDataOnSimObject(RequestID uint32, DefineID uint32, ObjectID uint32, Period uint32, Flags uint32, origin uint32, interval uint32, limit uint32) error {
	return errors.New("not implemented")
}

// RequestDataOnSimObjectType SimConnect_RequestDataOnSimObjectType(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, SIMCONNECT_DATA_DEFINITION_ID DefineID, DWORD dwRadiusMeters, SIMCONNECT_SIMOBJECT_TYPE type);
func (sc *SimConnect) RequestDataOnSimObjectType(RequestID uint32, DefineID uint32, dwRadiusMeters uint32, t uint32) error {
	err := sc.syscallSC.RequestDataOnSimObjectType(sc.hSimConnect, uintptr(RequestID), uintptr(DefineID), uintptr(dwRadiusMeters), uintptr(t))
	return err
}

// SetDataOnSimObject SimConnect_SetDataOnSimObject(HANDLE hSimConnect, SIMCONNECT_DATA_DEFINITION_ID DefineID, SIMCONNECT_OBJECT_ID ObjectID, SIMCONNECT_DATA_SET_FLAG Flags, DWORD ArrayCount, DWORD cbUnitSize, void * pDataSet);
func (sc *SimConnect) SetDataOnSimObject(DefineID uint32, ObjectID uint32, Flags uint32, ArrayCount uint32, cbUnitSize uint32, pDataSet []byte) error {
	if len(pDataSet) < 0 {
		return errors.New("Your pDataSet is too short on SetDataOnSimObject")
	}
	return sc.syscallSC.SetDataOnSimObject(sc.hSimConnect, uintptr(DefineID), uintptr(ObjectID), uintptr(Flags), uintptr(ArrayCount), uintptr(cbUnitSize), uintptr(unsafe.Pointer(&pDataSet[0])))
}

// MapInputEventToClientEvent SimConnect_MapInputEventToClientEvent(HANDLE hSimConnect, SIMCONNECT_INPUT_GROUP_ID GroupID, const char * szInputDefinition, SIMCONNECT_CLIENT_EVENT_ID DownEventID, DWORD DownValue = 0, SIMCONNECT_CLIENT_EVENT_ID UpEventID = (SIMCONNECT_CLIENT_EVENT_ID)SIMCONNECT_UNUSED, DWORD UpValue = 0, BOOL bMaskable = FALSE);
func (sc *SimConnect) MapInputEventToClientEvent(GroupID uint32, szInputDefinition string, DownEventID uint32, DownValue uint32, UpEventID uint32, UpValue uint32, bMaskable uint32) error {
	return errors.New("not implemented")
}

// SetInputGroupPriority SimConnect_SetInputGroupPriority(HANDLE hSimConnect, SIMCONNECT_INPUT_GROUP_ID GroupID, DWORD uPriority);
func (sc *SimConnect) SetInputGroupPriority(GroupID uint32, uPriority uint32) error {
	return errors.New("not implemented")
}

// RemoveInputEvent SimConnect_RemoveInputEvent(HANDLE hSimConnect, SIMCONNECT_INPUT_GROUP_ID GroupID, const char * szInputDefinition);
func (sc *SimConnect) RemoveInputEvent(GroupID uint32, szInputDefinition string) error {
	return errors.New("not implemented")
}

// ClearInputGroup SimConnect_ClearInputGroup(HANDLE hSimConnect, SIMCONNECT_INPUT_GROUP_ID GroupID);
func (sc *SimConnect) ClearInputGroup(GroupID uint32) error {
	return errors.New("not implemented")
}

// SetInputGroupState SimConnect_SetInputGroupState(HANDLE hSimConnect, SIMCONNECT_INPUT_GROUP_ID GroupID, DWORD dwState);
func (sc *SimConnect) SetInputGroupState(GroupID uint32, dwState uint32) error {
	return errors.New("not implemented")
}

// RequestReservedKey SimConnect_RequestReservedKey(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID EventID, const char * szKeyChoice1 = "", const char * szKeyChoice2 = "", const char * szKeyChoice3 = "");
func (sc *SimConnect) RequestReservedKey(EventID uint32, szKeyChoice1 string, szKeyChoice2 string, szKeyChoice3 string) error {
	return errors.New("not implemented")
}

// SubscribeToSystemEvent SimConnect_SubscribeToSystemEvent(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID EventID, const char * SystemEventName);
func (sc *SimConnect) SubscribeToSystemEvent(EventID uint32, SystemEventName string) error {
	err := sc.syscallSC.SubscribeToSystemEvent(sc.hSimConnect, uintptr(EventID), cChar(SystemEventName))
	return err
}

// UnsubscribeFromSystemEvent SimConnect_UnsubscribeFromSystemEvent(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID EventID);
func (sc *SimConnect) UnsubscribeFromSystemEvent(EventID uint32) error {
	return errors.New("not implemented")
}

// WeatherRequestInterpolatedObservation SimConnect_WeatherRequestInterpolatedObservation(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, float lat, float lon, float alt);
func (sc *SimConnect) WeatherRequestInterpolatedObservation(RequestID uint32, lat float32, lon float32, alt float32) error {
	return errors.New("not implemented")
}

// WeatherRequestObservationAtStation SimConnect_WeatherRequestObservationAtStation(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, const char * szICAO);
func (sc *SimConnect) WeatherRequestObservationAtStation(RequestID uint32, szICAO string) error {
	return errors.New("not implemented")
}

// WeatherRequestObservationAtNearestStation SimConnect_WeatherRequestObservationAtNearestStation(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, float lat, float lon);
func (sc *SimConnect) WeatherRequestObservationAtNearestStation(RequestID uint32, lat float32, lon float32) error {
	return errors.New("not implemented")
}

// WeatherCreateStation SimConnect_WeatherCreateStation(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, const char * szICAO, const char * szName, float lat, float lon, float alt);
func (sc *SimConnect) WeatherCreateStation(RequestID uint32, szICAO string, szName string, lat float32, lon float32, alt float32) error {
	return errors.New("not implemented")
}

// WeatherRemoveStation SimConnect_WeatherRemoveStation(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, const char * szICAO);
func (sc *SimConnect) WeatherRemoveStation(RequestID uint32, szICAO string) error {
	return errors.New("not implemented")
}

// WeatherSetObservation SimConnect_WeatherSetObservation(HANDLE hSimConnect, DWORD Seconds, const char * szMETAR);
func (sc *SimConnect) WeatherSetObservation(Seconds uint32, szMETAR string) error {
	return errors.New("not implemented")
}

// WeatherSetModeServer SimConnect_WeatherSetModeServer(HANDLE hSimConnect, DWORD dwPort, DWORD dwSeconds);
func (sc *SimConnect) WeatherSetModeServer(dwPort uint32, dwSeconds uint32) error {
	return errors.New("not implemented")
}

// WeatherSetModeTheme SimConnect_WeatherSetModeTheme(HANDLE hSimConnect, const char * szThemeName);
func (sc *SimConnect) WeatherSetModeTheme(szThemeName string) error {
	return errors.New("not implemented")
}

// WeatherSetModeGlobal SimConnect_WeatherSetModeGlobal(HANDLE hSimConnect);
func (sc *SimConnect) WeatherSetModeGlobal() error {
	return errors.New("not implemented")
}

// WeatherSetModeCustom SimConnect_WeatherSetModeCustom(HANDLE hSimConnect);
func (sc *SimConnect) WeatherSetModeCustom() error {
	return errors.New("not implemented")
}

// WeatherSetDynamicUpdateRate SimConnect_WeatherSetDynamicUpdateRate(HANDLE hSimConnect, DWORD dwRate);
func (sc *SimConnect) WeatherSetDynamicUpdateRate(dwRate uint32) error {
	return errors.New("not implemented")
}

// WeatherRequestCloudState SimConnect_WeatherRequestCloudState(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, float minLat, float minLon, float minAlt, float maxLat, float maxLon, float maxAlt, DWORD dwFlags = 0);
func (sc *SimConnect) WeatherRequestCloudState(RequestID uint32, minLat float32, minLon float32, minAlt float32, maxLat float32, maxLon float32, maxAlt float32, dwFlags uint32) error {
	return errors.New("not implemented")
}

// WeatherCreateThermal SimConnect_WeatherCreateThermal(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, float lat, float lon, float alt, float radius, float height, float coreRate = 3.0f, float coreTurbulence = 0.05f, float sinkRate = 3.0f, float sinkTurbulence = 0.2f, float coreSize = 0.4f, float coreTransitionSize = 0.1f, float sinkLayerSize = 0.4f, float sinkTransitionSize = 0.1f);
func (sc *SimConnect) WeatherCreateThermal(RequestID uint32, lat float32, lon float32, alt float32, radius float32, height float32, coreRate float32, coreTurbulence float32, sinkRate float32, sinkTurbulence float32, coreSize float32, coreTransitionSize float32, sinkLayerSize float32, sinkTransitionSize float32) error {
	return errors.New("not implemented")
}

// WeatherRemoveThermal SimConnect_WeatherRemoveThermal(HANDLE hSimConnect, SIMCONNECT_OBJECT_ID ObjectID);
func (sc *SimConnect) WeatherRemoveThermal(ObjectID uint32) error {
	return errors.New("not implemented")
}

// AICreateParkedATCAircraft SimConnect_AICreateParkedATCAircraft(HANDLE hSimConnect, const char * szContainerTitle, const char * szTailNumber, const char * szAirportID, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AICreateParkedATCAircraft(szContainerTitle string, szTailNumber string, szAirportID string, RequestID uint32) error {
	return errors.New("not implemented")
}

// AICreateEnrouteATCAircraft SimConnect_AICreateEnrouteATCAircraft(HANDLE hSimConnect, const char * szContainerTitle, const char * szTailNumber, int iFlightNumber, const char * szFlightPlanPath, double dFlightPlanPosition, BOOL bTouchAndGo, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AICreateEnrouteATCAircraft(szContainerTitle string, szTailNumber string, iFlightNumber int, szFlightPlanPath string, dFlightPlanPosition float64, bTouchAndGo uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// AICreateNonATCAircraft SimConnect_AICreateNonATCAircraft(HANDLE hSimConnect, const char * szContainerTitle, const char * szTailNumber, SIMCONNECT_DATA_INITPOSITION InitPos, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AICreateNonATCAircraft(szContainerTitle string, szTailNumber string, InitPos uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// AICreateSimulatedObject SimConnect_AICreateSimulatedObject(HANDLE hSimConnect, const char * szContainerTitle, SIMCONNECT_DATA_INITPOSITION InitPos, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AICreateSimulatedObject(szContainerTitle string, InitPos uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// AIReleaseControl SimConnect_AIReleaseControl(HANDLE hSimConnect, SIMCONNECT_OBJECT_ID ObjectID, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AIReleaseControl(ObjectID uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// AIRemoveObject SimConnect_AIRemoveObject(HANDLE hSimConnect, SIMCONNECT_OBJECT_ID ObjectID, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AIRemoveObject(ObjectID uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// AISetAircraftFlightPlan SimConnect_AISetAircraftFlightPlan(HANDLE hSimConnect, SIMCONNECT_OBJECT_ID ObjectID, const char * szFlightPlanPath, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) AISetAircraftFlightPlan(ObjectID uint32, szFlightPlanPath string, RequestID uint32) error {
	return errors.New("not implemented")
}

// ExecuteMissionAction SimConnect_ExecuteMissionAction(HANDLE hSimConnect, const GUID guidInstanceId);
func (sc *SimConnect) ExecuteMissionAction(guidInstanceID GUID) error {
	return errors.New("not implemented")
}

// CompleteCustomMissionAction SimConnect_CompleteCustomMissionAction(HANDLE hSimConnect, const GUID guidInstanceId);
func (sc *SimConnect) CompleteCustomMissionAction(guidInstanceID GUID) error {
	return errors.New("not implemented")
}

// Close SimConnect_Close(HANDLE hSimConnect);
func (sc *SimConnect) Close() error {
	err := sc.syscallSC.Close(sc.hSimConnect)
	return err
}

// RetrieveString SimConnect_RetrieveString(SIMCONNECT_RECV * pData, DWORD cbData, void * pStringV, char ** pszString, DWORD * pcbString);
func (sc *SimConnect) RetrieveString(pData *uint32, cbData uint32, pStringV string, pszString **string, pcbString *uint32) error {
	return errors.New("not implemented")
}

// GetLastSentPacketID SimConnect_GetLastSentPacketID(HANDLE hSimConnect, DWORD * pdwError);
func (sc *SimConnect) GetLastSentPacketID(pdwError *uint32) error {
	return errors.New("not implemented")
}

// Open SimConnect_Open(HANDLE * phSimConnect, LPCSTR szName, HWND hWnd, DWORD UserEventWin32, HANDLE hEventHandle, DWORD ConfigIndex);
func (sc *SimConnect) Open(appTitle string) error {
	err := sc.syscallSC.Open(uintptr(unsafe.Pointer(&sc.hSimConnect)), cChar(appTitle), uintptr(unsafe.Pointer(nil)), 0, 0, 0)
	if err != nil {
		return errors.New("No connected")
	}
	return err
}

// CallDispatch SimConnect_CallDispatch(HANDLE hSimConnect, DispatchProc pfcnDispatch, void * pContext);
//func (sc *SimConnect) CallDispatch( DispatchProc pfcnDispatch, void * pContext) error{
//}

// GetNextDispatch SimConnect_GetNextDispatch(HANDLE hSimConnect, SIMCONNECT_RECV ** ppData, DWORD * pcbData);
func (sc *SimConnect) GetNextDispatch(ppData *unsafe.Pointer, pcbData *uint32) error {
	err := sc.syscallSC.GetNextDispatch(sc.hSimConnect, uintptr(unsafe.Pointer(ppData)), uintptr(unsafe.Pointer(pcbData)))
	return err
}

// RequestResponseTimes SimConnect_RequestResponseTimes(HANDLE hSimConnect, DWORD nCount, float * fElapsedSeconds);
func (sc *SimConnect) RequestResponseTimes(nCount uint32, fElapsedSeconds *float32) error {
	return errors.New("not implemented")
}

// InsertString SimConnect_InsertString(char * pDest, DWORD cbDest, void ** ppEnd, DWORD * pcbStringV, const char * pSource);
func (sc *SimConnect) InsertString(pDest string, cbDest uint32, ppEnd *uint32, pcbStringV *uint32, pSource string) error {
	return errors.New("not implemented")
}

// CameraSetRelative6DOF SimConnect_CameraSetRelative6DOF(HANDLE hSimConnect, float fDeltaX, float fDeltaY, float fDeltaZ, float fPitchDeg, float fBankDeg, float fHeadingDeg);
func (sc *SimConnect) CameraSetRelative6DOF(fDeltaX float32, fDeltaY float32, fDeltaZ float32, fPitchDeg float32, fBankDeg float32, fHeadingDeg float32) error {
	return errors.New("not implemented")
}

// MenuAddItem SimConnect_MenuAddItem(HANDLE hSimConnect, const char * szMenuItem, SIMCONNECT_CLIENT_EVENT_ID MenuEventID, DWORD dwData);
func (sc *SimConnect) MenuAddItem(szMenuItem string, MenuEventID uint32, dwData uint32) error {
	return errors.New("not implemented")
}

// MenuDeleteItem SimConnect_MenuDeleteItem(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID MenuEventID);
func (sc *SimConnect) MenuDeleteItem(MenuEventID uint32) error {
	return errors.New("not implemented")
}

// MenuAddSubItem SimConnect_MenuAddSubItem(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID MenuEventID, const char * szMenuItem, SIMCONNECT_CLIENT_EVENT_ID SubMenuEventID, DWORD dwData);
func (sc *SimConnect) MenuAddSubItem(MenuEventID uint32, szMenuItem string, SubMenuEventID uint32, dwData uint32) error {
	return errors.New("not implemented")
}

// MenuDeleteSubItem SimConnect_MenuDeleteSubItem(HANDLE hSimConnect, SIMCONNECT_CLIENT_EVENT_ID MenuEventID, const SIMCONNECT_CLIENT_EVENT_ID SubMenuEventID);
func (sc *SimConnect) MenuDeleteSubItem(MenuEventID uint32, constSubMenuEventID uint32) error {
	return errors.New("not implemented")
}

// RequestSystemState SimConnect_RequestSystemState(HANDLE hSimConnect, SIMCONNECT_DATA_REQUEST_ID RequestID, const char * szState);
func (sc *SimConnect) RequestSystemState(RequestID uint32, szState string) error {
	return errors.New("not implemented")
}

// SetSystemState SimConnect_SetSystemState(HANDLE hSimConnect, const char * szState, DWORD dwInteger, float fFloat, const char * szString);
func (sc *SimConnect) SetSystemState(szState string, dwInteger uint32, fFloat float32, szString string) error {
	return errors.New("not implemented")
}

// MapClientDataNameToID SimConnect_MapClientDataNameToID(HANDLE hSimConnect, const char * szClientDataName, SIMCONNECT_CLIENT_DATA_ID ClientDataID);
func (sc *SimConnect) MapClientDataNameToID(szClientDataName string, ClientDataID uint32) error {
	return errors.New("not implemented")
}

// CreateClientData SimConnect_CreateClientData(HANDLE hSimConnect, SIMCONNECT_CLIENT_DATA_ID ClientDataID, DWORD dwSize, SIMCONNECT_CREATE_CLIENT_DATA_FLAG Flags);
func (sc *SimConnect) CreateClientData(ClientDataID uint32, dwSize uint32, Flags uint32) error {
	return errors.New("not implemented")
}

// AddToClientDataDefinition SimConnect_AddToClientDataDefinition(HANDLE hSimConnect, SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID, DWORD dwOffset, DWORD dwSizeOrType, float fEpsilon = 0, DWORD DatumID = SIMCONNECT_UNUSED);
func (sc *SimConnect) AddToClientDataDefinition(DefineID uint32, dwOffset uint32, dwSizeOrType uint32, fEpsilon float32, DatumID uint32) error {
	return errors.New("not implemented")
}

// ClearClientDataDefinition SimConnect_ClearClientDataDefinition(HANDLE hSimConnect, SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID);
func (sc *SimConnect) ClearClientDataDefinition(DefineID uint32) error {
	return errors.New("not implemented")
}

// RequestClientData SimConnect_RequestClientData(HANDLE hSimConnect, SIMCONNECT_CLIENT_DATA_ID ClientDataID, SIMCONNECT_DATA_REQUEST_ID RequestID, SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID, SIMCONNECT_CLIENT_DATA_PERIOD Period = SIMCONNECT_CLIENT_DATA_PERIOD_ONCE, SIMCONNECT_CLIENT_DATA_REQUEST_FLAG Flags = 0, DWORD origin = 0, DWORD interval = 0, DWORD limit = 0);
func (sc *SimConnect) RequestClientData(ClientDataID uint32, RequestID uint32, DefineID uint32, Period uint32, Flags uint32, origin uint32, interval uint32, limit uint32) error {
	return errors.New("not implemented")
}

// SetClientData SimConnect_SetClientData(HANDLE hSimConnect, SIMCONNECT_CLIENT_DATA_ID ClientDataID, SIMCONNECT_CLIENT_DATA_DEFINITION_ID DefineID, SIMCONNECT_CLIENT_DATA_SET_FLAG Flags, DWORD dwReserved, DWORD cbUnitSize, void * pDataSet);
func (sc *SimConnect) SetClientData(ClientDataID uint32, DefineID uint32, Flags uint32, dwReserved uint32, cbUnitSize uint32, pDataSet *uint32) error {
	return errors.New("not implemented")
}

// FlightLoad SimConnect_FlightLoad(HANDLE hSimConnect, const char * szFileName);
func (sc *SimConnect) FlightLoad(szFileName string) error {
	return errors.New("not implemented")
}

// FlightSave SimConnect_FlightSave(HANDLE hSimConnect, const char * szFileName, const char * szTitle, const char * szDescription, DWORD Flags);
func (sc *SimConnect) FlightSave(szFileName string, szTitle string, szDescription string, Flags uint32) error {
	return errors.New("not implemented")
}

// FlightPlanLoad SimConnect_FlightPlanLoad(HANDLE hSimConnect, const char * szFileName);
func (sc *SimConnect) FlightPlanLoad(szFileName string) error {
	return errors.New("not implemented")
}

// Text SimConnect_Text(HANDLE hSimConnect, SIMCONNECT_TEXT_TYPE type, float fTimeSeconds, SIMCONNECT_CLIENT_EVENT_ID EventID, DWORD cbUnitSize, void * pDataSet);
func (sc *SimConnect) Text(t uint32, fTimeSeconds float32, EventID uint32, cbUnitSize uint32, pDataSet *uint32) error {
	return errors.New("not implemented")
}

// SubscribeToFacilities SimConnect_SubscribeToFacilities(HANDLE hSimConnect, SIMCONNECT_FACILITY_LIST_TYPE type, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) SubscribeToFacilities(t uint32, RequestID uint32) error {
	return errors.New("not implemented")
}

// UnsubscribeToFacilities SimConnect_UnsubscribeToFacilities(HANDLE hSimConnect, SIMCONNECT_FACILITY_LIST_TYPE type);
func (sc *SimConnect) UnsubscribeToFacilities(t uint32) error {
	return errors.New("not implemented")
}

// RequestFacilitiesList SimConnect_RequestFacilitiesList(HANDLE hSimConnect, SIMCONNECT_FACILITY_LIST_TYPE type, SIMCONNECT_DATA_REQUEST_ID RequestID);
func (sc *SimConnect) RequestFacilitiesList(t uint32, RequestID uint32) error {
	return errors.New("not implemented")
}
