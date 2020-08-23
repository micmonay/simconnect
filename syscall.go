package simconnect

import (
	"errors"
	"syscall"
)

type SyscallSC struct {
	hSimConnect                                uintptr
	pMapClientEventToSimEvent                  *syscall.Proc
	pTransmitClientEvent                       *syscall.Proc
	pSetSystemEventState                       *syscall.Proc
	pAddClientEventToNotificationGroup         *syscall.Proc
	pRemoveClientEvent                         *syscall.Proc
	pSetNotificationGroupPriority              *syscall.Proc
	pClearNotificationGroup                    *syscall.Proc
	pRequestNotificationGroup                  *syscall.Proc
	pAddToDataDefinition                       *syscall.Proc
	pClearDataDefinition                       *syscall.Proc
	pRequestDataOnSimObject                    *syscall.Proc
	pRequestDataOnSimObjectType                *syscall.Proc
	pSetDataOnSimObject                        *syscall.Proc
	pMapInputEventToClientEvent                *syscall.Proc
	pSetInputGroupPriority                     *syscall.Proc
	pRemoveInputEvent                          *syscall.Proc
	pClearInputGroup                           *syscall.Proc
	pSetInputGroupState                        *syscall.Proc
	pRequestReservedKey                        *syscall.Proc
	pSubscribeToSystemEvent                    *syscall.Proc
	pUnsubscribeFromSystemEvent                *syscall.Proc
	pWeatherRequestInterpolatedObservation     *syscall.Proc
	pWeatherRequestObservationAtStation        *syscall.Proc
	pWeatherRequestObservationAtNearestStation *syscall.Proc
	pWeatherCreateStation                      *syscall.Proc
	pWeatherRemoveStation                      *syscall.Proc
	pWeatherSetObservation                     *syscall.Proc
	pWeatherSetModeServer                      *syscall.Proc
	pWeatherSetModeTheme                       *syscall.Proc
	pWeatherSetModeGlobal                      *syscall.Proc
	pWeatherSetModeCustom                      *syscall.Proc
	pWeatherSetDynamicUpdateRate               *syscall.Proc
	pWeatherRequestCloudState                  *syscall.Proc
	pWeatherCreateThermal                      *syscall.Proc
	pWeatherRemoveThermal                      *syscall.Proc
	pAICreateParkedATCAircraft                 *syscall.Proc
	pAICreateEnrouteATCAircraft                *syscall.Proc
	pAICreateNonATCAircraft                    *syscall.Proc
	pAICreateSimulatedObject                   *syscall.Proc
	pAIReleaseControl                          *syscall.Proc
	pAIRemoveObject                            *syscall.Proc
	pAISetAircraftFlightPlan                   *syscall.Proc
	pExecuteMissionAction                      *syscall.Proc
	pCompleteCustomMissionAction               *syscall.Proc
	pClose                                     *syscall.Proc
	pRetrieveString                            *syscall.Proc
	pGetLastSentPacketID                       *syscall.Proc
	pOpen                                      *syscall.Proc
	pCallDispatch                              *syscall.Proc
	pGetNextDispatch                           *syscall.Proc
	pRequestResponseTimes                      *syscall.Proc
	pInsertString                              *syscall.Proc
	pCameraSetRelative6DOF                     *syscall.Proc
	pMenuAddItem                               *syscall.Proc
	pMenuDeleteItem                            *syscall.Proc
	pMenuAddSubItem                            *syscall.Proc
	pMenuDeleteSubItem                         *syscall.Proc
	pRequestSystemState                        *syscall.Proc
	pSetSystemState                            *syscall.Proc
	pMapClientDataNameToID                     *syscall.Proc
	pCreateClientData                          *syscall.Proc
	pAddToClientDataDefinition                 *syscall.Proc
	pClearClientDataDefinition                 *syscall.Proc
	pRequestClientData                         *syscall.Proc
	pSetClientData                             *syscall.Proc
	pFlightLoad                                *syscall.Proc
	pFlightSave                                *syscall.Proc
	pFlightPlanLoad                            *syscall.Proc
	pText                                      *syscall.Proc
	pSubscribeToFacilities                     *syscall.Proc
	pUnsubscribeToFacilities                   *syscall.Proc
	pRequestFacilitiesList                     *syscall.Proc
}

// NewsyscallSC.pinit all syscall
func NewSyscallSC() (*SyscallSC, error) {
	simDLL, err := syscall.LoadDLL("SimConnect.dll")
	if err != nil {
		return nil, err
	}
	syscallSC := &SyscallSC{}
	syscallSC.pMapClientEventToSimEvent, err = simDLL.FindProc("SimConnect_MapClientEventToSimEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pTransmitClientEvent, err = simDLL.FindProc("SimConnect_TransmitClientEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetSystemEventState, err = simDLL.FindProc("SimConnect_SetSystemEventState")
	if err != nil {
		return nil, err
	}
	syscallSC.pAddClientEventToNotificationGroup, err = simDLL.FindProc("SimConnect_AddClientEventToNotificationGroup")
	if err != nil {
		return nil, err
	}
	syscallSC.pRemoveClientEvent, err = simDLL.FindProc("SimConnect_RemoveClientEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetNotificationGroupPriority, err = simDLL.FindProc("SimConnect_SetNotificationGroupPriority")
	if err != nil {
		return nil, err
	}
	syscallSC.pClearNotificationGroup, err = simDLL.FindProc("SimConnect_ClearNotificationGroup")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestNotificationGroup, err = simDLL.FindProc("SimConnect_RequestNotificationGroup")
	if err != nil {
		return nil, err
	}
	syscallSC.pAddToDataDefinition, err = simDLL.FindProc("SimConnect_AddToDataDefinition")
	if err != nil {
		return nil, err
	}
	syscallSC.pClearDataDefinition, err = simDLL.FindProc("SimConnect_ClearDataDefinition")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestDataOnSimObject, err = simDLL.FindProc("SimConnect_RequestDataOnSimObject")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestDataOnSimObjectType, err = simDLL.FindProc("SimConnect_RequestDataOnSimObjectType")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetDataOnSimObject, err = simDLL.FindProc("SimConnect_SetDataOnSimObject")
	if err != nil {
		return nil, err
	}
	syscallSC.pMapInputEventToClientEvent, err = simDLL.FindProc("SimConnect_MapInputEventToClientEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetInputGroupPriority, err = simDLL.FindProc("SimConnect_SetInputGroupPriority")
	if err != nil {
		return nil, err
	}
	syscallSC.pRemoveInputEvent, err = simDLL.FindProc("SimConnect_RemoveInputEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pClearInputGroup, err = simDLL.FindProc("SimConnect_ClearInputGroup")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetInputGroupState, err = simDLL.FindProc("SimConnect_SetInputGroupState")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestReservedKey, err = simDLL.FindProc("SimConnect_RequestReservedKey")
	if err != nil {
		return nil, err
	}
	syscallSC.pSubscribeToSystemEvent, err = simDLL.FindProc("SimConnect_SubscribeToSystemEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pUnsubscribeFromSystemEvent, err = simDLL.FindProc("SimConnect_UnsubscribeFromSystemEvent")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRequestInterpolatedObservation, err = simDLL.FindProc("SimConnect_WeatherRequestInterpolatedObservation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRequestObservationAtStation, err = simDLL.FindProc("SimConnect_WeatherRequestObservationAtStation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRequestObservationAtNearestStation, err = simDLL.FindProc("SimConnect_WeatherRequestObservationAtNearestStation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherCreateStation, err = simDLL.FindProc("SimConnect_WeatherCreateStation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRemoveStation, err = simDLL.FindProc("SimConnect_WeatherRemoveStation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetObservation, err = simDLL.FindProc("SimConnect_WeatherSetObservation")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetModeServer, err = simDLL.FindProc("SimConnect_WeatherSetModeServer")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetModeTheme, err = simDLL.FindProc("SimConnect_WeatherSetModeTheme")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetModeGlobal, err = simDLL.FindProc("SimConnect_WeatherSetModeGlobal")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetModeCustom, err = simDLL.FindProc("SimConnect_WeatherSetModeCustom")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherSetDynamicUpdateRate, err = simDLL.FindProc("SimConnect_WeatherSetDynamicUpdateRate")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRequestCloudState, err = simDLL.FindProc("SimConnect_WeatherRequestCloudState")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherCreateThermal, err = simDLL.FindProc("SimConnect_WeatherCreateThermal")
	if err != nil {
		return nil, err
	}
	syscallSC.pWeatherRemoveThermal, err = simDLL.FindProc("SimConnect_WeatherRemoveThermal")
	if err != nil {
		return nil, err
	}
	syscallSC.pAICreateParkedATCAircraft, err = simDLL.FindProc("SimConnect_AICreateParkedATCAircraft")
	if err != nil {
		return nil, err
	}
	syscallSC.pAICreateEnrouteATCAircraft, err = simDLL.FindProc("SimConnect_AICreateEnrouteATCAircraft")
	if err != nil {
		return nil, err
	}
	syscallSC.pAICreateNonATCAircraft, err = simDLL.FindProc("SimConnect_AICreateNonATCAircraft")
	if err != nil {
		return nil, err
	}
	syscallSC.pAICreateSimulatedObject, err = simDLL.FindProc("SimConnect_AICreateSimulatedObject")
	if err != nil {
		return nil, err
	}
	syscallSC.pAIReleaseControl, err = simDLL.FindProc("SimConnect_AIReleaseControl")
	if err != nil {
		return nil, err
	}
	syscallSC.pAIRemoveObject, err = simDLL.FindProc("SimConnect_AIRemoveObject")
	if err != nil {
		return nil, err
	}
	syscallSC.pAISetAircraftFlightPlan, err = simDLL.FindProc("SimConnect_AISetAircraftFlightPlan")
	if err != nil {
		return nil, err
	}
	syscallSC.pExecuteMissionAction, err = simDLL.FindProc("SimConnect_ExecuteMissionAction")
	if err != nil {
		return nil, err
	}
	syscallSC.pCompleteCustomMissionAction, err = simDLL.FindProc("SimConnect_CompleteCustomMissionAction")
	if err != nil {
		return nil, err
	}
	syscallSC.pClose, err = simDLL.FindProc("SimConnect_Close")
	if err != nil {
		return nil, err
	}
	syscallSC.pRetrieveString, err = simDLL.FindProc("SimConnect_RetrieveString")
	if err != nil {
		return nil, err
	}
	syscallSC.pGetLastSentPacketID, err = simDLL.FindProc("SimConnect_GetLastSentPacketID")
	if err != nil {
		return nil, err
	}
	syscallSC.pOpen, err = simDLL.FindProc("SimConnect_Open")
	if err != nil {
		return nil, err
	}
	syscallSC.pCallDispatch, err = simDLL.FindProc("SimConnect_CallDispatch")
	if err != nil {
		return nil, err
	}
	syscallSC.pGetNextDispatch, err = simDLL.FindProc("SimConnect_GetNextDispatch")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestResponseTimes, err = simDLL.FindProc("SimConnect_RequestResponseTimes")
	if err != nil {
		return nil, err
	}
	syscallSC.pInsertString, err = simDLL.FindProc("SimConnect_InsertString")
	if err != nil {
		return nil, err
	}
	syscallSC.pCameraSetRelative6DOF, err = simDLL.FindProc("SimConnect_CameraSetRelative6DOF")
	if err != nil {
		return nil, err
	}
	syscallSC.pMenuAddItem, err = simDLL.FindProc("SimConnect_MenuAddItem")
	if err != nil {
		return nil, err
	}
	syscallSC.pMenuDeleteItem, err = simDLL.FindProc("SimConnect_MenuDeleteItem")
	if err != nil {
		return nil, err
	}
	syscallSC.pMenuAddSubItem, err = simDLL.FindProc("SimConnect_MenuAddSubItem")
	if err != nil {
		return nil, err
	}
	syscallSC.pMenuDeleteSubItem, err = simDLL.FindProc("SimConnect_MenuDeleteSubItem")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestSystemState, err = simDLL.FindProc("SimConnect_RequestSystemState")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetSystemState, err = simDLL.FindProc("SimConnect_SetSystemState")
	if err != nil {
		return nil, err
	}
	syscallSC.pMapClientDataNameToID, err = simDLL.FindProc("SimConnect_MapClientDataNameToID")
	if err != nil {
		return nil, err
	}
	syscallSC.pCreateClientData, err = simDLL.FindProc("SimConnect_CreateClientData")
	if err != nil {
		return nil, err
	}
	syscallSC.pAddToClientDataDefinition, err = simDLL.FindProc("SimConnect_AddToClientDataDefinition")
	if err != nil {
		return nil, err
	}
	syscallSC.pClearClientDataDefinition, err = simDLL.FindProc("SimConnect_ClearClientDataDefinition")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestClientData, err = simDLL.FindProc("SimConnect_RequestClientData")
	if err != nil {
		return nil, err
	}
	syscallSC.pSetClientData, err = simDLL.FindProc("SimConnect_SetClientData")
	if err != nil {
		return nil, err
	}
	syscallSC.pFlightLoad, err = simDLL.FindProc("SimConnect_FlightLoad")
	if err != nil {
		return nil, err
	}
	syscallSC.pFlightSave, err = simDLL.FindProc("SimConnect_FlightSave")
	if err != nil {
		return nil, err
	}
	syscallSC.pFlightPlanLoad, err = simDLL.FindProc("SimConnect_FlightPlanLoad")
	if err != nil {
		return nil, err
	}
	syscallSC.pText, err = simDLL.FindProc("SimConnect_Text")
	if err != nil {
		return nil, err
	}
	syscallSC.pSubscribeToFacilities, err = simDLL.FindProc("SimConnect_SubscribeToFacilities")
	if err != nil {
		return nil, err
	}
	syscallSC.pUnsubscribeToFacilities, err = simDLL.FindProc("SimConnect_UnsubscribeToFacilities")
	if err != nil {
		return nil, err
	}
	syscallSC.pRequestFacilitiesList, err = simDLL.FindProc("SimConnect_RequestFacilitiesList")
	if err != nil {
		return nil, err
	}
	return syscallSC, nil
}
func (syscallSC *SyscallSC) MapClientEventToSimEvent(hSimConnect uintptr, EventID uintptr, EventName uintptr) error {
	r1, _, _ := syscallSC.pMapClientEventToSimEvent.Call(hSimConnect, EventID, EventName)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) TransmitClientEvent(hSimConnect uintptr, ObjectID uintptr, EventID uintptr, dwData uintptr, GroupID uintptr, Flags uintptr) error {
	r1, _, _ := syscallSC.pTransmitClientEvent.Call(hSimConnect, ObjectID, EventID, dwData, GroupID, Flags)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetSystemEventState(hSimConnect uintptr, EventID uintptr, dwState uintptr) error {
	r1, _, _ := syscallSC.pSetSystemEventState.Call(hSimConnect, EventID, dwState)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AddClientEventToNotificationGroup(hSimConnect uintptr, GroupID uintptr, EventID uintptr, bMaskable uintptr) error {
	r1, _, _ := syscallSC.pAddClientEventToNotificationGroup.Call(hSimConnect, GroupID, EventID, bMaskable)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RemoveClientEvent(hSimConnect uintptr, GroupID uintptr, EventID uintptr) error {
	r1, _, _ := syscallSC.pRemoveClientEvent.Call(hSimConnect, GroupID, EventID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetNotificationGroupPriority(hSimConnect uintptr, GroupID uintptr, uPriority uintptr) error {
	r1, _, _ := syscallSC.pSetNotificationGroupPriority.Call(hSimConnect, GroupID, uPriority)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) ClearNotificationGroup(hSimConnect uintptr, GroupID uintptr) error {
	r1, _, _ := syscallSC.pClearNotificationGroup.Call(hSimConnect, GroupID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestNotificationGroup(hSimConnect uintptr, GroupID uintptr, dwReserved uintptr, Flags uintptr) error {
	r1, _, _ := syscallSC.pRequestNotificationGroup.Call(hSimConnect, GroupID, dwReserved, Flags)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AddToDataDefinition(hSimConnect uintptr, DefineID uintptr, DatumName uintptr, UnitsName uintptr, DatumType uintptr, fEpsilon uintptr, DatumID uintptr) error {
	r1, _, _ := syscallSC.pAddToDataDefinition.Call(hSimConnect, DefineID, DatumName, UnitsName, DatumType, fEpsilon, DatumID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) ClearDataDefinition(hSimConnect uintptr, DefineID uintptr) error {
	r1, _, _ := syscallSC.pClearDataDefinition.Call(hSimConnect, DefineID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestDataOnSimObject(hSimConnect uintptr, RequestID uintptr, DefineID uintptr, ObjectID uintptr, Period uintptr, Flags uintptr, origin uintptr, interval uintptr, limit uintptr) error {
	r1, _, _ := syscallSC.pRequestDataOnSimObject.Call(hSimConnect, RequestID, DefineID, ObjectID, Period, Flags, origin, interval, limit)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestDataOnSimObjectType(hSimConnect uintptr, RequestID uintptr, DefineID uintptr, dwRadiusMeters uintptr, t uintptr) error {
	r1, _, _ := syscallSC.pRequestDataOnSimObjectType.Call(hSimConnect, RequestID, DefineID, dwRadiusMeters, t)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetDataOnSimObject(hSimConnect uintptr, DefineID uintptr, ObjectID uintptr, Flags uintptr, ArrayCount uintptr, cbUnitSize uintptr, pDataSet uintptr) error {
	r1, _, _ := syscallSC.pSetDataOnSimObject.Call(hSimConnect, DefineID, ObjectID, Flags, ArrayCount, cbUnitSize, pDataSet)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MapInputEventToClientEvent(hSimConnect uintptr, GroupID uintptr, szInputDefinition uintptr, DownEventID uintptr, DownValue uintptr, UpEventID uintptr, UpValue uintptr, bMaskable uintptr) error {
	r1, _, _ := syscallSC.pMapInputEventToClientEvent.Call(hSimConnect, GroupID, szInputDefinition, DownEventID, DownValue, UpEventID, UpValue, bMaskable)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetInputGroupPriority(hSimConnect uintptr, GroupID uintptr, uPriority uintptr) error {
	r1, _, _ := syscallSC.pSetInputGroupPriority.Call(hSimConnect, GroupID, uPriority)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RemoveInputEvent(hSimConnect uintptr, GroupID uintptr, szInputDefinition uintptr) error {
	r1, _, _ := syscallSC.pRemoveInputEvent.Call(hSimConnect, GroupID, szInputDefinition)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) ClearInputGroup(hSimConnect uintptr, GroupID uintptr) error {
	r1, _, _ := syscallSC.pClearInputGroup.Call(hSimConnect, GroupID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetInputGroupState(hSimConnect uintptr, GroupID uintptr, dwState uintptr) error {
	r1, _, _ := syscallSC.pSetInputGroupState.Call(hSimConnect, GroupID, dwState)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestReservedKey(hSimConnect uintptr, EventID uintptr, szKeyChoice1 uintptr, szKeyChoice2 uintptr, szKeyChoice3 uintptr) error {
	r1, _, _ := syscallSC.pRequestReservedKey.Call(hSimConnect, EventID, szKeyChoice1, szKeyChoice2, szKeyChoice3)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SubscribeToSystemEvent(hSimConnect uintptr, EventID uintptr, SystemEventName uintptr) error {
	r1, _, _ := syscallSC.pSubscribeToSystemEvent.Call(hSimConnect, EventID, SystemEventName)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) UnsubscribeFromSystemEvent(hSimConnect uintptr, EventID uintptr) error {
	r1, _, _ := syscallSC.pUnsubscribeFromSystemEvent.Call(hSimConnect, EventID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRequestInterpolatedObservation(hSimConnect uintptr, RequestID uintptr, lat uintptr, lon uintptr, alt uintptr) error {
	r1, _, _ := syscallSC.pWeatherRequestInterpolatedObservation.Call(hSimConnect, RequestID, lat, lon, alt)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRequestObservationAtStation(hSimConnect uintptr, RequestID uintptr, szICAO uintptr) error {
	r1, _, _ := syscallSC.pWeatherRequestObservationAtStation.Call(hSimConnect, RequestID, szICAO)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRequestObservationAtNearestStation(hSimConnect uintptr, RequestID uintptr, lat uintptr, lon uintptr) error {
	r1, _, _ := syscallSC.pWeatherRequestObservationAtNearestStation.Call(hSimConnect, RequestID, lat, lon)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherCreateStation(hSimConnect uintptr, RequestID uintptr, szICAO uintptr, szName uintptr, lat uintptr, lon uintptr, alt uintptr) error {
	r1, _, _ := syscallSC.pWeatherCreateStation.Call(hSimConnect, RequestID, szICAO, szName, lat, lon, alt)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRemoveStation(hSimConnect uintptr, RequestID uintptr, szICAO uintptr) error {
	r1, _, _ := syscallSC.pWeatherRemoveStation.Call(hSimConnect, RequestID, szICAO)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetObservation(hSimConnect uintptr, Seconds uintptr, szMETAR uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetObservation.Call(hSimConnect, Seconds, szMETAR)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetModeServer(hSimConnect uintptr, dwPort uintptr, dwSeconds uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetModeServer.Call(hSimConnect, dwPort, dwSeconds)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetModeTheme(hSimConnect uintptr, szThemeName uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetModeTheme.Call(hSimConnect, szThemeName)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetModeGlobal(hSimConnect uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetModeGlobal.Call(hSimConnect)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetModeCustom(hSimConnect uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetModeCustom.Call(hSimConnect)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherSetDynamicUpdateRate(hSimConnect uintptr, dwRate uintptr) error {
	r1, _, _ := syscallSC.pWeatherSetDynamicUpdateRate.Call(hSimConnect, dwRate)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRequestCloudState(hSimConnect uintptr, RequestID uintptr, minLat uintptr, minLon uintptr, minAlt uintptr, maxLat uintptr, maxLon uintptr, maxAlt uintptr, dwFlags uintptr) error {
	r1, _, _ := syscallSC.pWeatherRequestCloudState.Call(hSimConnect, RequestID, minLat, minLon, minAlt, maxLat, maxLon, maxAlt, dwFlags)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherCreateThermal(hSimConnect uintptr, RequestID uintptr, lat uintptr, lon uintptr, alt uintptr, radius uintptr, height uintptr, coreRate uintptr, coreTurbulence uintptr, sinkRate uintptr, sinkTurbulence uintptr, coreSize uintptr, coreTransitionSize uintptr, sinkLayerSize uintptr, sinkTransitionSize uintptr) error {
	r1, _, _ := syscallSC.pWeatherCreateThermal.Call(hSimConnect, RequestID, lat, lon, alt, radius, height, coreRate, coreTurbulence, sinkRate, sinkTurbulence, coreSize, coreTransitionSize, sinkLayerSize, sinkTransitionSize)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) WeatherRemoveThermal(hSimConnect uintptr, ObjectID uintptr) error {
	r1, _, _ := syscallSC.pWeatherRemoveThermal.Call(hSimConnect, ObjectID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AICreateParkedATCAircraft(hSimConnect uintptr, szContainerTitle uintptr, szTailNumber uintptr, szAirportID uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAICreateParkedATCAircraft.Call(hSimConnect, szContainerTitle, szTailNumber, szAirportID, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AICreateEnrouteATCAircraft(hSimConnect uintptr, szContainerTitle uintptr, szTailNumber uintptr, iFlightNumber uintptr, szFlightPlanPath uintptr, dFlightPlanPosition uintptr, bTouchAndGo uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAICreateEnrouteATCAircraft.Call(hSimConnect, szContainerTitle, szTailNumber, iFlightNumber, szFlightPlanPath, dFlightPlanPosition, bTouchAndGo, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AICreateNonATCAircraft(hSimConnect uintptr, szContainerTitle uintptr, szTailNumber uintptr, InitPos uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAICreateNonATCAircraft.Call(hSimConnect, szContainerTitle, szTailNumber, InitPos, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AICreateSimulatedObject(hSimConnect uintptr, szContainerTitle uintptr, InitPos uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAICreateSimulatedObject.Call(hSimConnect, szContainerTitle, InitPos, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AIReleaseControl(hSimConnect uintptr, ObjectID uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAIReleaseControl.Call(hSimConnect, ObjectID, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AIRemoveObject(hSimConnect uintptr, ObjectID uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAIRemoveObject.Call(hSimConnect, ObjectID, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AISetAircraftFlightPlan(hSimConnect uintptr, ObjectID uintptr, szFlightPlanPath uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pAISetAircraftFlightPlan.Call(hSimConnect, ObjectID, szFlightPlanPath, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) ExecuteMissionAction(hSimConnect uintptr, guidInstanceId uintptr) error {
	r1, _, _ := syscallSC.pExecuteMissionAction.Call(hSimConnect, guidInstanceId)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) CompleteCustomMissionAction(hSimConnect uintptr, guidInstanceId uintptr) error {
	r1, _, _ := syscallSC.pCompleteCustomMissionAction.Call(hSimConnect, guidInstanceId)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) Close(hSimConnect uintptr) error {
	r1, _, _ := syscallSC.pClose.Call(hSimConnect)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RetrieveString(pData uintptr, cbData uintptr, pStringV uintptr, pszString uintptr, pcbString uintptr) error {
	r1, _, _ := syscallSC.pRetrieveString.Call(pData, cbData, pStringV, pszString, pcbString)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) GetLastSentPacketID(hSimConnect uintptr, pdwError uintptr) error {
	r1, _, _ := syscallSC.pGetLastSentPacketID.Call(hSimConnect, pdwError)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) Open(phSimConnect uintptr, szName uintptr, hWnd uintptr, UserEventWin uintptr, hEventHandle uintptr, ConfigIndex uintptr) error {
	r1, _, _ := syscallSC.pOpen.Call(phSimConnect, szName, hWnd, UserEventWin, hEventHandle, ConfigIndex)
	if r1 != 0 {
		return errors.New("r1 error")
	}
	return nil
}
func (syscallSC *SyscallSC) CallDispatch(hSimConnect uintptr, pfcnDispatch uintptr, pContext uintptr) error {
	r1, _, _ := syscallSC.pCallDispatch.Call(hSimConnect, pfcnDispatch, pContext)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) GetNextDispatch(hSimConnect uintptr, ppData uintptr, pcbData uintptr) error {
	r1, _, _ := syscallSC.pGetNextDispatch.Call(hSimConnect, ppData, pcbData)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestResponseTimes(hSimConnect uintptr, nCount uintptr, fElapsedSeconds uintptr) error {
	r1, _, _ := syscallSC.pRequestResponseTimes.Call(hSimConnect, nCount, fElapsedSeconds)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) InsertString(pDest uintptr, cbDest uintptr, ppEnd uintptr, pcbStringV uintptr, pSource uintptr) error {
	r1, _, _ := syscallSC.pInsertString.Call(pDest, cbDest, ppEnd, pcbStringV, pSource)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) CameraSetRelative6DOF(hSimConnect uintptr, fDeltaX uintptr, fDeltaY uintptr, fDeltaZ uintptr, fPitchDeg uintptr, fBankDeg uintptr, fHeadingDeg uintptr) error {
	r1, _, _ := syscallSC.pCameraSetRelative6DOF.Call(hSimConnect, fDeltaX, fDeltaY, fDeltaZ, fPitchDeg, fBankDeg, fHeadingDeg)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MenuAddItem(hSimConnect uintptr, szMenuItem uintptr, MenuEventID uintptr, dwData uintptr) error {
	r1, _, _ := syscallSC.pMenuAddItem.Call(hSimConnect, szMenuItem, MenuEventID, dwData)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MenuDeleteItem(hSimConnect uintptr, MenuEventID uintptr) error {
	r1, _, _ := syscallSC.pMenuDeleteItem.Call(hSimConnect, MenuEventID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MenuAddSubItem(hSimConnect uintptr, MenuEventID uintptr, szMenuItem uintptr, SubMenuEventID uintptr, dwData uintptr) error {
	r1, _, _ := syscallSC.pMenuAddSubItem.Call(hSimConnect, MenuEventID, szMenuItem, SubMenuEventID, dwData)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MenuDeleteSubItem(hSimConnect uintptr, MenuEventID uintptr, SubMenuEventID uintptr) error {
	r1, _, _ := syscallSC.pMenuDeleteSubItem.Call(hSimConnect, MenuEventID, SubMenuEventID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestSystemState(hSimConnect uintptr, RequestID uintptr, szState uintptr) error {
	r1, _, _ := syscallSC.pRequestSystemState.Call(hSimConnect, RequestID, szState)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetSystemState(hSimConnect uintptr, szState uintptr, dwInteger uintptr, fFloat uintptr, szString uintptr) error {
	r1, _, _ := syscallSC.pSetSystemState.Call(hSimConnect, szState, dwInteger, fFloat, szString)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) MapClientDataNameToID(hSimConnect uintptr, szClientDataName uintptr, ClientDataID uintptr) error {
	r1, _, _ := syscallSC.pMapClientDataNameToID.Call(hSimConnect, szClientDataName, ClientDataID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) CreateClientData(hSimConnect uintptr, ClientDataID uintptr, dwSize uintptr, Flags uintptr) error {
	r1, _, _ := syscallSC.pCreateClientData.Call(hSimConnect, ClientDataID, dwSize, Flags)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) AddToClientDataDefinition(hSimConnect uintptr, DefineID uintptr, dwOffset uintptr, dwSizeOrType uintptr, fEpsilon uintptr, DatumID uintptr) error {
	r1, _, _ := syscallSC.pAddToClientDataDefinition.Call(hSimConnect, DefineID, dwOffset, dwSizeOrType, fEpsilon, DatumID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) ClearClientDataDefinition(hSimConnect uintptr, DefineID uintptr) error {
	r1, _, _ := syscallSC.pClearClientDataDefinition.Call(hSimConnect, DefineID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestClientData(hSimConnect uintptr, ClientDataID uintptr, RequestID uintptr, DefineID uintptr, Period uintptr, Flags uintptr, origin uintptr, interval uintptr, limit uintptr) error {
	r1, _, _ := syscallSC.pRequestClientData.Call(hSimConnect, ClientDataID, RequestID, DefineID, Period, Flags, origin, interval, limit)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SetClientData(hSimConnect uintptr, ClientDataID uintptr, DefineID uintptr, Flags uintptr, dwReserved uintptr, cbUnitSize uintptr, pDataSet uintptr) error {
	r1, _, _ := syscallSC.pSetClientData.Call(hSimConnect, ClientDataID, DefineID, Flags, dwReserved, cbUnitSize, pDataSet)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) FlightLoad(hSimConnect uintptr, szFileName uintptr) error {
	r1, _, _ := syscallSC.pFlightLoad.Call(hSimConnect, szFileName)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) FlightSave(hSimConnect uintptr, szFileName uintptr, szTitle uintptr, szDescription uintptr, Flags uintptr) error {
	r1, _, _ := syscallSC.pFlightSave.Call(hSimConnect, szFileName, szTitle, szDescription, Flags)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) FlightPlanLoad(hSimConnect uintptr, szFileName uintptr) error {
	r1, _, _ := syscallSC.pFlightPlanLoad.Call(hSimConnect, szFileName)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) Text(hSimConnect uintptr, t uintptr, fTimeSeconds uintptr, EventID uintptr, cbUnitSize uintptr, pDataSet uintptr) error {
	r1, _, _ := syscallSC.pText.Call(hSimConnect, t, fTimeSeconds, EventID, cbUnitSize, pDataSet)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) SubscribeToFacilities(hSimConnect uintptr, t uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pSubscribeToFacilities.Call(hSimConnect, t, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) UnsubscribeToFacilities(hSimConnect uintptr, t uintptr) error {
	r1, _, _ := syscallSC.pUnsubscribeToFacilities.Call(hSimConnect, t)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
func (syscallSC *SyscallSC) RequestFacilitiesList(hSimConnect uintptr, t uintptr, RequestID uintptr) error {
	r1, _, _ := syscallSC.pRequestFacilitiesList.Call(hSimConnect, t, RequestID)
	if r1 != 0 {
		return errors.New("r1 error")
	}

	return nil
}
