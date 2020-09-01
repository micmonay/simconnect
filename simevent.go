package simconnect

const (
	//SimEvent1sec Request a notification every second.
	SimEvent1sec = "1sec"
	//SimEvent4sec Request a notification every four seconds.
	SimEvent4sec = "4sec"
	//SimEvent6Hz Request notifications six times per second. This is the same rate that joystick movement events are transmitted.
	SimEvent6Hz = "6Hz"
	//SimEventAircraftLoaded Request a notification when the aircraft flight dynamics file is changed. These files have a .AIR extension. The filename is returned in a SIMCONNECT_RECV_EVENT_FILENAME structure.
	SimEventAircraftLoaded = "AircraftLoaded"
	//SimEventCrashed Request a notification if the user aircraft crashes.
	SimEventCrashed = "Crashed"
	//SimEventCrashReset Request a notification when the crash cut-scene has completed.
	SimEventCrashReset = "CrashReset"
	//SimEventFlightLoaded Request a notification when a flight is loaded. Note that when a flight is ended, a default flight is typically loaded, so these events will occur when flights and missions are started and finished. The filename of the flight loaded is returned in a SIMCONNECT_RECV_EVENT_FILENAME structure.
	SimEventFlightLoaded = "FlightLoaded"
	//SimEventFlightSaved Request a notification when a flight is saved correctly. The filename of the flight saved is returned in a SIMCONNECT_RECV_EVENT_FILENAME structure.
	SimEventFlightSaved = "FlightSaved"
	//SimEventFlightPlanActivated Request a notification when a new flight plan is activated. The filename of the activated flight plan is returned in a SIMCONNECT_RECV_EVENT_FILENAME structure.
	SimEventFlightPlanActivated = "FlightPlanActivated"
	//SimEventFlightPlanDeactivated Request a notification when the active flight plan is de-activated.
	SimEventFlightPlanDeactivated = "FlightPlanDeactivated"
	//SimEventFrame Request notifications every visual frame. Information is returned in a SIMCONNECT_RECV_EVENT_FRAME structure.
	SimEventFrame = "Frame"
	//SimEventPause Request notifications when the flight is paused or unpaused, and also immediately returns the current pause state (1 = paused or 0 = unpaused). The state is returned in the dwData parameter.
	SimEventPause = "Pause"
	//SimEventPaused Request a notification when the flight is paused.
	SimEventPaused = "Paused"
	//SimEventPauseFrame Request notifications for every visual frame that the simulation is paused. Information is returned in a SIMCONNECT_RECV_EVENT_FRAME structure.
	SimEventPauseFrame = "PauseFrame"
	//SimEventPositionChanged Request a notification when the user changes the position of their aircraft through a dialog.
	SimEventPositionChanged = "PositionChanged"
	//SimEventSim Request notifications when the flight is running or not, and also immediately returns the current state (1 = running or 0 = not running). The state is returned in the dwData parameter.
	SimEventSim = "Sim"
	//SimEventSimStart The simulator is running. Typically the user is actively controlling the aircraft on the ground or in the air. However, in some cases additional pairs of SimStart/SimStop events are sent. For example, when a flight is reset the events that are sent are SimStop, SimStart, SimStop, SimStart. Also when a flight is started with the SHOW_OPENING_SCREEN value (defined in the FSX.CFG file) set to zero, then an additional SimStart/SimStop pair are sent before a second SimStart event is sent when the scenery is fully loaded. The opening screen provides the options to change aircraft, departure airport, and so on.
	SimEventSimStart = "SimStart"
	//SimEventSimStop The simulator is not running. Typically the user is loading a flight, navigating the shell or in a dialog.
	SimEventSimStop = "SimStop"
	//SimEventSound Requests a notification when the master sound switch is changed. This request will also return the current state of the master sound switch immediately. A flag is returned in the dwData parameter, 0 if the switch is off, SIMCONNECT_SOUND_SYSTEM_EVENT_DATA_MASTER (0x1) if the switch is on.
	SimEventSound = "Sound"
	//SimEventUnpaused Request a notification when the flight is un-paused.
	SimEventUnpaused = "Unpaused"
	//SimEventView Requests a notification when the user aircraft view is changed. This request will also return the current view immediately. A flag is returned in the dwData parameter, one of: SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_2D SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_VIRTUAL SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_ORTHOGONAL (the map view).
	SimEventView = "View"
	//SimEventWeatherModeChanged Request a notification when the weather mode is changed.
	SimEventWeatherModeChanged = "WeatherModeChanged"
	//SimEventObjectAdded Request a notification when an AI object is added to the simulation. Refer also to the SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE structure.
	SimEventObjectAdded = "ObjectAdded"
	//SimEventObjectRemoved Request a notification when an AI object is removed from the simulation. Refer also to the SIMCONNECT_RECV_EVENT_OBJECT_ADDREMOVE structure.
	SimEventObjectRemoved = "ObjectRemoved"
	//SimEventMissionCompleted Request a notification when the user has completed a mission. Refer also to the SIMCONNECT_MISSION_END enum.
	SimEventMissionCompleted = "MissionCompleted"
	//SimEventCustomMissionActionExecuted Request a notification when a mission action has been executed. Refer also to the SimConnect_CompleteCustomMissionAction function.
	SimEventCustomMissionActionExecuted = "CustomMissionActionExecuted"
	//SimEventMultiplayerClientStarted Used by a client to request a notification that they have successfully joined a multiplayer race. The event is returned as a SIMCONNECT_RECV_EVENT_MULTIPLAYER_CLIENT_STARTED structure. This event is only sent to the client, not the host of the session.
	SimEventMultiplayerClientStarted = "MultiplayerClientStarted"
	//SimEventMultiplayerServerStarted Used by a host of a multiplayer race to request a notification when the race is open to other players in the lobby. The event is returned in a SIMCONNECT_RECV_EVENT_MULTIPLAYER_SERVER_STARTED structure.
	SimEventMultiplayerServerStarted = "MultiplayerServerStarted"
	//SimEventMultiplayerSessionEnded Request a notification when the mutliplayer race session is terminated. The event is returned in a SIMCONNECT_RECV_EVENT_MULTIPLAYER_SESSION_ENDED structure. If a client player leaves a race, this event wil be returned just to the client. If a host leaves or terminates the session, then all players will receive this event. This is the only event that will be broadcast to all players.
	SimEventMultiplayerSessionEnded = "MultiplayerSessionEnded"
	//SimEventRaceEnd Request a notification of the race results for each racer. The results will be returned in SIMCONNECT_RECV_EVENT_RACE_END structures, one for each player.
	SimEventRaceEnd = "RaceEnd"
	//SimEventRaceLap Request a notification of the race results for each racer. The results will be returned in SIMCONNECT_RECV_EVENT_RACE_LAP structures, one for each player.
	SimEventRaceLap = "RaceLap"
)
