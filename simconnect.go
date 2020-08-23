package simconnect

type SimConnect struct {
}

//MapClientEventToSimEvent( EventID uint32, EventName string )
func (sc *SimConnect) MapClientEventToSimEvent(EventID uint32, EventName string) error {
	return nil
}

//TransmitClientEvent( ObjectID uint32, EventID uint32, dwData uint32,GroupID uint32,Flags uint32)
func (sc *SimConnect) TransmitClientEvent(ObjectID uint32, EventID uint32, dwData uint32, GroupID uint32, Flags uint32) error {
	return nil
}

//SetSystemEventState( EventID uint32,dwState uint32)
func (sc *SimConnect) SetSystemEventState(EventID uint32, dwState uint32) error {
	return nil
}

//AddClientEventToNotificationGroup(GroupID uint32, EventID uint32,bMaskable  uint32)
func (sc *SimConnect) AddClientEventToNotificationGroup(GroupID uint32, EventID uint32, bMaskable uint32) error {
	return nil
}

//RemoveClientEvent(GroupID uint32, EventID uint32)
func (sc *SimConnect) RemoveClientEvent(GroupID uint32, EventID uint32) error {
	return nil
}

//SetNotificationGroupPriority(GroupID uint32, uPriority uint32)
func (sc *SimConnect) SetNotificationGroupPriority(GroupID uint32, uPriority uint32) error {
	return nil
}

//ClearNotificationGroup(GroupID uint32)
func (sc *SimConnect) ClearNotificationGroup(GroupID uint32) error {
	return nil
}

//RequestNotificationGroup(GroupID uint32, dwReserved  uint32, Flags  uint32)
func (sc *SimConnect) RequestNotificationGroup(GroupID uint32, dwReserved uint32, Flags uint32) error {
	return nil
}

//AddToDataDefinition(DefineID uint32, DatumName string, UnitsName string,DatumType  uint32, fEpsilon  float32, DatumID  uint32)
func (sc *SimConnect) AddToDataDefinition(DefineID uint32, DatumName string, UnitsName string, DatumType uint32, fEpsilon float32, DatumID uint32) error {
	return nil
}

//ClearDataDefinition(DefineID uint32)
func (sc *SimConnect) ClearDataDefinition(DefineID uint32) error {
	return nil
}

//RequestDataOnSimObject(RequestID uint32,DefineID uint32, ObjectID uint32,Period uint32,Flags  uint32, origin  uint32, interval  uint32, limit  uint32)
func (sc *SimConnect) RequestDataOnSimObject(RequestID uint32, DefineID uint32, ObjectID uint32, Period uint32, Flags uint32, origin uint32, interval uint32, limit uint32) error {
	return nil
}

//RequestDataOnSimObjectType(RequestID uint32,DefineID uint32, dwRadiusMeters uint32,type uint32)
func (sc *SimConnect) RequestDataOnSimObjectType(RequestID uint32, DefineID uint32, dwRadiusMeters uint32, t uint32) error {
	return nil
}

//SetDataOnSimObject(DefineID uint32, ObjectID uint32,Flags uint32, ArrayCount uint32, cbUnitSize uint32, void * pDataSet)
func (sc *SimConnect) SetDataOnSimObject(DefineID uint32, ObjectID uint32, Flags uint32, ArrayCount uint32, cbUnitSize uint32, pDataSet *uint32) error {
	return nil
}

//MapInputEventToClientEvent(GroupID uint32, szInputDefinition string,DownEventID uint32, DownValue  uint32,UpEventID  uint32)SIMCONNECT_UNUSED, UpValue  uint32,bMaskable  uint32)
func (sc *SimConnect) MapInputEventToClientEvent(GroupID uint32, szInputDefinition string, DownEventID uint32, DownValue uint32, UpEventID uint32, UpValue uint32, bMaskable uint32) error {
	return nil
}

//SetInputGroupPriority(GroupID uint32, uPriority uint32)
func (sc *SimConnect) SetInputGroupPriority(GroupID uint32, uPriority uint32) error {
	return nil
}

//RemoveInputEvent(GroupID uint32, szInputDefinition string)
func (sc *SimConnect) RemoveInputEvent(GroupID uint32, szInputDefinition string) error {
	return nil
}

//ClearInputGroup(GroupID uint32)
func (sc *SimConnect) ClearInputGroup(GroupID uint32) error {
	return nil
}

//SetInputGroupState(GroupID uint32, dwState uint32)
func (sc *SimConnect) SetInputGroupState(GroupID uint32, dwState uint32) error {
	return nil
}

//RequestReservedKey( EventID uint32, szKeyChoice1 string , szKeyChoice2 string , szKeyChoice3 string )
func (sc *SimConnect) RequestReservedKey(EventID uint32, szKeyChoice1 string, szKeyChoice2 string, szKeyChoice3 string) error {
	return nil
}

//SubscribeToSystemEvent( EventID uint32, SystemEventName string)
func (sc *SimConnect) SubscribeToSystemEvent(EventID uint32, SystemEventName string) error {
	return nil
}

//UnsubscribeFromSystemEvent( EventID uint32)
func (sc *SimConnect) UnsubscribeFromSystemEvent(EventID uint32) error {
	return nil
}

//WeatherRequestInterpolatedObservation(RequestID uint32, lat float32, lon float32, alt float32)
func (sc *SimConnect) WeatherRequestInterpolatedObservation(RequestID uint32, lat float32, lon float32, alt float32) error {
	return nil
}

//WeatherRequestObservationAtStation(RequestID uint32, szICAO string)
func (sc *SimConnect) WeatherRequestObservationAtStation(RequestID uint32, szICAO string) error {
	return nil
}

//WeatherRequestObservationAtNearestStation(RequestID uint32, lat float32, lon float32)
func (sc *SimConnect) WeatherRequestObservationAtNearestStation(RequestID uint32, lat float32, lon float32) error {
	return nil
}

//WeatherCreateStation(RequestID uint32, szICAO string, szName string, lat float32, lon float32, alt float32)
func (sc *SimConnect) WeatherCreateStation(RequestID uint32, szICAO string, szName string, lat float32, lon float32, alt float32) error {
	return nil
}

//WeatherRemoveStation(RequestID uint32, szICAO string)
func (sc *SimConnect) WeatherRemoveStation(RequestID uint32, szICAO string) error {
	return nil
}

//WeatherSetObservation( Seconds uint32, szMETAR string)
func (sc *SimConnect) WeatherSetObservation(Seconds uint32, szMETAR string) error {
	return nil
}

//WeatherSetModeServer( dwPort uint32, dwSeconds uint32)
func (sc *SimConnect) WeatherSetModeServer(dwPort uint32, dwSeconds uint32) error {
	return nil
}

//WeatherSetModeTheme( szThemeName string)
func (sc *SimConnect) WeatherSetModeTheme(szThemeName string) error {
	return nil
}

//WeatherSetModeGlobal()
func (sc *SimConnect) WeatherSetModeGlobal() error {
	return nil
}

//WeatherSetModeCustom()
func (sc *SimConnect) WeatherSetModeCustom() error {
	return nil
}

//WeatherSetDynamicUpdateRate( dwRate uint32)
func (sc *SimConnect) WeatherSetDynamicUpdateRate(dwRate uint32) error {
	return nil
}

//WeatherRequestCloudState(RequestID uint32, minLat float32, minLon float32, minAlt float32, maxLat float32, maxLon float32, maxAlt float32, dwFlags  uint32)
func (sc *SimConnect) WeatherRequestCloudState(RequestID uint32, minLat float32, minLon float32, minAlt float32, maxLat float32, maxLon float32, maxAlt float32, dwFlags uint32) error {
	return nil
}

//WeatherCreateThermal(RequestID uint32, lat float32, lon float32, alt float32, radius float32, height float32, coreRate  float32, coreTurbulence  float32, sinkRate  float32, sinkTurbulence  float32, coreSize  float32, coreTransitionSize  float32, sinkLayerSize  float32, sinkTransitionSize  float32)
func (sc *SimConnect) WeatherCreateThermal(RequestID uint32, lat float32, lon float32, alt float32, radius float32, height float32, coreRate float32, coreTurbulence float32, sinkRate float32, sinkTurbulence float32, coreSize float32, coreTransitionSize float32, sinkLayerSize float32, sinkTransitionSize float32) error {
	return nil
}

//WeatherRemoveThermal( ObjectID uint32)
func (sc *SimConnect) WeatherRemoveThermal(ObjectID uint32) error {
	return nil
}

//AICreateParkedATCAircraft( szContainerTitle string, szTailNumber string, szAirportID string,RequestID uint32)
func (sc *SimConnect) AICreateParkedATCAircraft(szContainerTitle string, szTailNumber string, szAirportID string, RequestID uint32) error {
	return nil
}

//AICreateEnrouteATCAircraft( szContainerTitle string, szTailNumber string, int iFlightNumber, szFlightPlanPath string, dFlightPlanPosition float64,bTouchAndGo uint32,RequestID uint32)
func (sc *SimConnect) AICreateEnrouteATCAircraft(szContainerTitle string, szTailNumber string, iFlightNumber int, szFlightPlanPath string, dFlightPlanPosition float64, bTouchAndGo uint32, RequestID uint32) error {
	return nil
}

//AICreateNonATCAircraft( szContainerTitle string, szTailNumber string,InitPos uint32,RequestID uint32)
func (sc *SimConnect) AICreateNonATCAircraft(szContainerTitle string, szTailNumber string, InitPos uint32, RequestID uint32) error {
	return nil
}

//AICreateSimulatedObject( szContainerTitle string,InitPos uint32,RequestID uint32)
func (sc *SimConnect) AICreateSimulatedObject(szContainerTitle string, InitPos uint32, RequestID uint32) error {
	return nil
}

//AIReleaseControl( ObjectID uint32,RequestID uint32)
func (sc *SimConnect) AIReleaseControl(ObjectID uint32, RequestID uint32) error {
	return nil
}

//AIRemoveObject( ObjectID uint32,RequestID uint32)
func (sc *SimConnect) AIRemoveObject(ObjectID uint32, RequestID uint32) error {
	return nil
}

//AISetAircraftFlightPlan( ObjectID uint32, szFlightPlanPath string,RequestID uint32)
func (sc *SimConnect) AISetAircraftFlightPlan(ObjectID uint32, szFlightPlanPath string, RequestID uint32) error {
	return nil
}

//ExecuteMissionAction( constguidInstanceId uint32)
func (sc *SimConnect) ExecuteMissionAction(constguidInstanceId uint32) error {
	return nil
}

//CompleteCustomMissionAction( constguidInstanceId uint32)
func (sc *SimConnect) CompleteCustomMissionAction(constguidInstanceId uint32) error {
	return nil
}

//Close()
func (sc *SimConnect) Close() error {
	return nil
}

//RetrieveString(SIMCONNECT_RECV * pData, cbData uint32, void * pStringV, char ** pszString, * pcbString uint32)
func (sc *SimConnect) RetrieveString(pData *uint32, cbData uint32, pStringV string, pszString string, pcbString *uint32) error {
	return nil
}

//GetLastSentPacketID( * pdwError uint32)
func (sc *SimConnect) GetLastSentPacketID(pdwError *uint32) error {
	return nil
}

//Open(HANDLE * phSimConnect,szName uint32,hWnd uint32, UserEventWin32 uint32,hEventHandle uint32, ConfigIndex uint32)
func (sc *SimConnect) Open() error {
	return nil
}

//CallDispatch( DispatchProc pfcnDispatch, void * pContext)
//func (sc *SimConnect) CallDispatch( DispatchProc pfcnDispatch, void * pContext) error{
//}
//GetNextDispatch(** ppData uint32, * pcbData uint32)
func (sc *SimConnect) GetNextDispatch(ppData *uint32, pcbData *uint32) error {
	return nil
}

//RequestResponseTimes( nCount uint32, * fElapsedSeconds float32)
func (sc *SimConnect) RequestResponseTimes(nCount uint32, fElapsedSeconds *float32) error {
	return nil
}

//InsertString(char * pDest, cbDest uint32, void ** ppEnd, * pcbStringV uint32, pSource string)
func (sc *SimConnect) InsertString(pDest string, cbDest uint32, ppEnd *uint32, pcbStringV *uint32, pSource string) error {
	return nil
}

//CameraSetRelative6DOF( fDeltaX float32, fDeltaY float32, fDeltaZ float32, fPitchDeg float32, fBankDeg float32, fHeadingDeg float32)
func (sc *SimConnect) CameraSetRelative6DOF(fDeltaX float32, fDeltaY float32, fDeltaZ float32, fPitchDeg float32, fBankDeg float32, fHeadingDeg float32) error {
	return nil
}

//MenuAddItem( szMenuItem string,MenuEventID uint32, dwData uint32)
func (sc *SimConnect) MenuAddItem(szMenuItem string, MenuEventID uint32, dwData uint32) error {
	return nil
}

//MenuDeleteItem(MenuEventID uint32)
func (sc *SimConnect) MenuDeleteItem(MenuEventID uint32) error {
	return nil
}

//MenuAddSubItem(MenuEventID uint32, szMenuItem string,SubMenuEventID uint32, dwData uint32)
func (sc *SimConnect) MenuAddSubItem(MenuEventID uint32, szMenuItem string, SubMenuEventID uint32, dwData uint32) error {
	return nil
}

//MenuDeleteSubItem(MenuEventID uint32, constSubMenuEventID uint32)
func (sc *SimConnect) MenuDeleteSubItem(MenuEventID uint32, constSubMenuEventID uint32) error {
	return nil
}

//RequestSystemState(RequestID uint32, szState string)
func (sc *SimConnect) RequestSystemState(RequestID uint32, szState string) error {
	return nil
}

//SetSystemState( szState string, dwInteger uint32, fFloat float32, szString string)
func (sc *SimConnect) SetSystemState(szState string, dwInteger uint32, fFloat float32, szString string) error {
	return nil
}

//MapClientDataNameToID( szClientDataName string,ClientDataID uint32)
func (sc *SimConnect) MapClientDataNameToID(szClientDataName string, ClientDataID uint32) error {
	return nil
}

//CreateClientData(ClientDataID uint32, dwSize uint32,Flags uint32)
func (sc *SimConnect) CreateClientData(ClientDataID uint32, dwSize uint32, Flags uint32) error {
	return nil
}

//AddToClientDataDefinition(DefineID uint32, dwOffset uint32, dwSizeOrType uint32, fEpsilon  float32, DatumID  uint32)
func (sc *SimConnect) AddToClientDataDefinition(DefineID uint32, dwOffset uint32, dwSizeOrType uint32, fEpsilon float32, DatumID uint32) error {
	return nil
}

//ClearClientDataDefinition(DefineID uint32)
func (sc *SimConnect) ClearClientDataDefinition(DefineID uint32) error {
	return nil
}

//RequestClientData(ClientDataID uint32,RequestID uint32,DefineID uint32,Period  uint32,Flags  uint32, origin  uint32, interval  uint32, limit  uint32)
func (sc *SimConnect) RequestClientData(ClientDataID uint32, RequestID uint32, DefineID uint32, Period uint32, Flags uint32, origin uint32, interval uint32, limit uint32) error {
	return nil
}

//SetClientData(ClientDataID uint32,DefineID uint32,Flags uint32, dwReserved uint32, cbUnitSize uint32, void * pDataSet)
func (sc *SimConnect) SetClientData(ClientDataID uint32, DefineID uint32, Flags uint32, dwReserved uint32, cbUnitSize uint32, pDataSet *uint32) error {
	return nil
}

//FlightLoad( szFileName string)
func (sc *SimConnect) FlightLoad(szFileName string) error {
	return nil
}

//FlightSave( szFileName string, szTitle string, szDescription string, Flags uint32)
func (sc *SimConnect) FlightSave(szFileName string, szTitle string, szDescription string, Flags uint32) error {
	return nil
}

//FlightPlanLoad( szFileName string)
func (sc *SimConnect) FlightPlanLoad(szFileName string) error {
	return nil
}

//Text(type uint32, fTimeSeconds float32, EventID uint32, cbUnitSize uint32, void * pDataSet)
func (sc *SimConnect) Text(t uint32, fTimeSeconds float32, EventID uint32, cbUnitSize uint32, pDataSet *uint32) error {
	return nil
}

//SubscribeToFacilities(type uint32,RequestID uint32)
func (sc *SimConnect) SubscribeToFacilities(t uint32, RequestID uint32) error {
	return nil
}

//UnsubscribeToFacilities(type uint32)
func (sc *SimConnect) UnsubscribeToFacilities(t uint32) error {
	return nil
}

//RequestFacilitiesList(type uint32,RequestID uint32)
func (sc *SimConnect) RequestFacilitiesList(t uint32, RequestID uint32) error {
	return nil
}
