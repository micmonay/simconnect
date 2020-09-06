package simconnect

//KeySimEvent is a string
type KeySimEvent string

// Dcumentation based on http://www.prepar3d.com/SDKv3/LearningCenter/utilities/variables/event_ids.html
const (
	//KeySlingPickupRelease Toggle between pickup and release mode. Hold mode is automatic and cannot be selected. Refer to the document Notes on Aircraft Systems.
	KeySlingPickupRelease KeySimEvent = "SLING_PICKUP_RELEASE"
	//KeyHoistSwitchExtend The rate at which a hoist cable extends is set in the Aircraft Configuration File.
	KeyHoistSwitchExtend KeySimEvent = "HOIST_SWITCH_EXTEND"
	//KeyHoistSwitchRetract The rate at which a hoist cable retracts is set in the Aircraft Configuration File.
	KeyHoistSwitchRetract KeySimEvent = "HOIST_SWITCH_RETRACT"
	//KeyHoistSwitchSet The data value should be set to one of: <0 up =0 off >0 down
	KeyHoistSwitchSet KeySimEvent = "HOIST_SWITCH_SET"
	//KeyHoistDeployToggle Toggles the hoist arm switch, extend or retract.
	KeyHoistDeployToggle KeySimEvent = "HOIST_DEPLOY_TOGGLE"
	//KeyHoistDeploySet The data value should be set to: 0 - set hoist switch to retract the arm 1 - set hoist switch to extend the arm
	KeyHoistDeploySet KeySimEvent = "HOIST_DEPLOY_SET"
	//KeyToggleAntidetonationTankValve Toggle the antidetonation valve. Pass a value to determine which tank, if there are multiple tanks, to use. Tanks are indexed from 1. Refer to the document Notes on Aircraft Systems.
	KeyToggleAntidetonationTankValve KeySimEvent = "ANTIDETONATION_TANK_VALVE_TOGGLE"
	//KeyToggleNitrousTankValve Toggle the nitrous valve. Pass a value to determine which tank, if there are multiple tanks, to use. Tanks are indexed from 1.
	KeyToggleNitrousTankValve KeySimEvent = "NITROUS_TANK_VALVE_TOGGLE"
	//KeyToggleRaceresultsWindow Show or hide multiplayer race results.	Disabled
	KeyToggleRaceresultsWindow KeySimEvent = "TOGGLE_RACERESULTS_WINDOW"
	//KeyTakeoffAssistArmToggle Deploy or remove the assist arm. Refer to the document Notes on Aircraft Systems.
	KeyTakeoffAssistArmToggle KeySimEvent = "TAKEOFF_ASSIST_ARM_TOGGLE"
	//KeyTakeoffAssistArmSet Value: TRUE request set FALSE request unset
	KeyTakeoffAssistArmSet KeySimEvent = "TAKEOFF_ASSIST_ARM_SET"
	//KeyTakeoffAssistFire If everything is set up correctly. Launch from the catapult.
	KeyTakeoffAssistFire KeySimEvent = "TAKEOFF_ASSIST_FIRE"
	//KeyToggleLaunchBarSwitch Toggle the request for the launch bar to be installed or removed.
	KeyToggleLaunchBarSwitch KeySimEvent = "TOGGLE_LAUNCH_BAR_SWITCH"
	//KeySetLaunchbarSwitch Value: TRUE request set FALSE request unset
	KeySetLaunchbarSwitch KeySimEvent = "SET_LAUNCH_BAR_SWITCH"
	//KeyRepairAndRefuel Fully repair and refuel the user aircraft. Ignored if flight realism is enforced.
	KeyRepairAndRefuel KeySimEvent = "REPAIR_AND_REFUEL"
	//KeyDmeSelect Selects one of the two DME systems (1,2).
	KeyDmeSelect KeySimEvent = "DME_SELECT"
	//KeyFuelDumpToggle Turns on or off the fuel dump switch.
	KeyFuelDumpToggle KeySimEvent = "FUEL_DUMP_TOGGLE"
	//KeyViewCockpitForward Switch immediately to the forward view, in 2D mode.
	KeyViewCockpitForward KeySimEvent = "VIEW_COCKPIT_FORWARD"
	//KeyViewVirtualCockpitForward Switch immediately to the forward view, in virtual cockpit mode.
	KeyViewVirtualCockpitForward KeySimEvent = "VIEW_VIRTUAL_COCKPIT_FORWARD"
	//KeyTowPlaneRelease Release a towed aircraft, usually a glider.
	KeyTowPlaneRelease KeySimEvent = "TOW_PLANE_RELEASE"
	//KeyRequestTowPlane Request a tow plane. The user aircraft must be tow-able, stationary, on the ground and not already attached for this to succeed.
	KeyRequestTowPlane KeySimEvent = "TOW_PLANE_REQUEST"
	//KeyRequestFuel Request a fuel truck. The aircraft must be in a parking spot for this to be successful.	 Fuel Selection Keys
	KeyRequestFuel KeySimEvent = "REQUEST_FUEL_KEY"
	//KeyReleaseDroppableObjects Release one droppable object. Multiple key events will release multiple objects.
	KeyReleaseDroppableObjects KeySimEvent = "RELEASE_DROPPABLE_OBJECTS"
	//KeyViewPanelAlphaSet Sets the alpha-blending value for the panel. Takes a parameter in the range 0 to 255. The alpha-blending can be changed from the keyboard using Ctrl-Shift-T, and the plus and minus keys.
	KeyViewPanelAlphaSet KeySimEvent = "VIEW_PANEL_ALPHA_SET"
	//KeyViewPanelAlphaSelect Sets the mode to change the alpha-blending, so the keys KEY_PLUS and KEY_MINUS increment and decrement the value.
	KeyViewPanelAlphaSelect KeySimEvent = "VIEW_PANEL_ALPHA_SELECT"
	//KeyViewPanelAlphaInc Increment alpha-blending for the panel.
	KeyViewPanelAlphaInc KeySimEvent = "VIEW_PANEL_ALPHA_INC"
	//KeyViewPanelAlphaDec Decrement alpha-blending for the panel.
	KeyViewPanelAlphaDec KeySimEvent = "VIEW_PANEL_ALPHA_DEC"
	//KeyViewLinkingSet Links all the views from one camera together, so that panning the view will change the view of all the linked cameras.
	KeyViewLinkingSet KeySimEvent = "VIEW_LINKING_SET"
	//KeyViewLinkingToggle Turns view linking on or off.
	KeyViewLinkingToggle KeySimEvent = "VIEW_LINKING_TOGGLE"
	//KeyRadioSelectedDmeIdentEnable Turns on the identification sound for the selected DME.
	KeyRadioSelectedDmeIdentEnable KeySimEvent = "RADIO_SELECTED_DME_IDENT_ENABLE"
	//KeyRadioSelectedDmeIdentDisable Turns off the identification sound for the selected DME.
	KeyRadioSelectedDmeIdentDisable KeySimEvent = "RADIO_SELECTED_DME_IDENT_DISABLE"
	//KeyRadioSelectedDmeIdentSet Sets the DME identification sound to the given filename.
	KeyRadioSelectedDmeIdentSet KeySimEvent = "RADIO_SELECTED_DME_IDENT_SET"
	//KeyRadioSelectedDmeIdentToggle Turns on or off the identification sound for the selected DME.
	KeyRadioSelectedDmeIdentToggle KeySimEvent = "RADIO_SELECTED_DME_IDENT_TOGGLE"
	//KeyGaugeKeystroke Enables a keystroke to be sent to a gauge that is in focus. The keystrokes can only be in the range 0 to 9, A to Z, and the four keys: plus, minus, comma and period. This is typically used to allow some keyboard entry to a complex device such as a GPS to enter such things as ICAO codes using the keyboard, rather than turning dials.
	KeyGaugeKeystroke      KeySimEvent = "GAUGE_KEYSTROKE"
	KeySimuiWindowHideshow KeySimEvent = "SIMUI_WINDOW_HIDESHOW"
	//KeyToggleVariometerSwitch Turn the variometer on or off.
	KeyToggleVariometerSwitch KeySimEvent = "TOGGLE_VARIOMETER_SWITCH"
	//KeyToggleTurnIndicatorSwitch Turn the turn indicator on or off.
	KeyToggleTurnIndicatorSwitch KeySimEvent = "TOGGLE_TURN_INDICATOR_SWITCH"
	//KeyWindowTitlesToggle Turn window titles on or off.
	KeyWindowTitlesToggle KeySimEvent = "VIEW_WINDOW_TITLES_TOGGLE"
	//KeyAxisPanPitch Sets the pitch of the axis. Requires an angle.
	KeyAxisPanPitch KeySimEvent = "AXIS_PAN_PITCH"
	//KeyAxisPanHeading Sets the heading of the axis. Requires an angle.
	KeyAxisPanHeading KeySimEvent = "AXIS_PAN_HEADING"
	//KeyAxisPanTilt Sets the tilt of the axis. Requires an angle.
	KeyAxisPanTilt KeySimEvent = "AXIS_PAN_TILT"
	//KeyAxisIndicatorCycle Step through the view axes.
	KeyAxisIndicatorCycle KeySimEvent = "VIEW_AXIS_INDICATOR_CYCLE"
	//KeyMapOrientationCycle Step through the map orientations.
	KeyMapOrientationCycle KeySimEvent = "VIEW_MAP_ORIENTATION_CYCLE"
	//KeyToggleJetway Requests a jetway, which will only be answered if the aircraft is at a parking spot.
	KeyToggleJetway KeySimEvent = "TOGGLE_JETWAY"
	//KeyRetractFloatSwitchDec If the plane has retractable floats, moves the retract position from Extend to Neutral, or Neutral to Retract.
	KeyRetractFloatSwitchDec KeySimEvent = "RETRACT_FLOAT_SWITCH_DEC"
	//KeyRetractFloatSwitchInc If the plane has retractable floats, moves the retract position from Retract to Neutral, or Neutral to Extend.
	KeyRetractFloatSwitchInc KeySimEvent = "RETRACT_FLOAT_SWITCH_INC"
	//KeyToggleWaterBallastValve Turn the water ballast valve on or off.
	KeyToggleWaterBallastValve KeySimEvent = "TOGGLE_WATER_BALLAST_VALVE"
	//KeyViewChaseDistanceAdd Increments the distance of the view camera from the chase object (such as in Spot Plane view, or viewing an AI controlled aircraft).
	KeyViewChaseDistanceAdd KeySimEvent = "VIEW_CHASE_DISTANCE_ADD"
	//KeyViewChaseDistanceSub Decrements the distance of the view camera from the chase object.
	KeyViewChaseDistanceSub KeySimEvent = "VIEW_CHASE_DISTANCE_SUB"
	//KeyApuStarter Start up the auxiliary power unit (APU).
	KeyApuStarter KeySimEvent = "APU_STARTER"
	//KeyApuOffSwitch Turn the APU off.
	KeyApuOffSwitch KeySimEvent = "APU_OFF_SWITCH"
	//KeyApuGeneratorSwitchToggle Turn the auxiliary generator on or off.
	KeyApuGeneratorSwitchToggle KeySimEvent = "APU_GENERATOR_SWITCH_TOGGLE"
	//KeyApuGeneratorSwitchSet Set the auxiliary generator switch (0,1).
	KeyApuGeneratorSwitchSet KeySimEvent = "APU_GENERATOR_SWITCH_SET"
	//KeyExtinguishEngineFire Takes a two digit argument.  The first digit represents the fire extinguisher index, and the second represents the engine index.  For example,  11 would represent using bottle 1 on engine 1.  21 would represent using bottle 2 on engine 1.  Typical entries for a twin engine aircraft would be 11 and 22.
	KeyExtinguishEngineFire KeySimEvent = "EXTINGUISH_ENGINE_FIRE"
	//KeyApMaxBankInc Autopilot max bank angle increment.
	KeyApMaxBankInc KeySimEvent = "AP_MAX_BANK_INC"
	//KeyApMaxBankDec Autopilot max bank angle decrement.
	KeyApMaxBankDec KeySimEvent = "AP_MAX_BANK_DEC"
	//KeyApN1Hold Autopilot, hold the N1 percentage at its current level.
	KeyApN1Hold KeySimEvent = "AP_N1_HOLD"
	//KeyApN1RefInc Increment the autopilot N1 reference.
	KeyApN1RefInc KeySimEvent = "AP_N1_REF_INC"
	//KeyApN1RefDec Decrement the autopilot N1 reference.
	KeyApN1RefDec KeySimEvent = "AP_N1_REF_DEC"
	//KeyApN1RefSet Sets the autopilot N1 reference.
	KeyApN1RefSet KeySimEvent = "AP_N1_REF_SET"
	//KeyHydraulicSwitchToggle Turn the hydraulic switch on or off.
	KeyHydraulicSwitchToggle KeySimEvent = "HYDRAULIC_SWITCH_TOGGLE"
	//KeyBleedAirSourceControlInc Increases the bleed air source control.
	KeyBleedAirSourceControlInc KeySimEvent = "BLEED_AIR_SOURCE_CONTROL_INC"
	//KeyBleedAirSourceControlDec Decreases the bleed air source control.
	KeyBleedAirSourceControlDec KeySimEvent = "BLEED_AIR_SOURCE_CONTROL_DEC"
	//KeyTurbineIgnitionSwitchToggle Toggles the turbine ignition switch between OFF and AUTO.
	KeyTurbineIgnitionSwitchToggle KeySimEvent = "TURBINE_IGNITION_SWITCH_TOGGLE"
	//KeyCabinNoSmokingAlertSwitchToggle Turn the "No smoking" alert on or off.
	KeyCabinNoSmokingAlertSwitchToggle KeySimEvent = "CABIN_NO_SMOKING_ALERT_SWITCH_TOGGLE"
	//KeyCabinSeatbeltsAlertSwitchToggle Turn the "Fasten seatbelts" alert on or off.
	KeyCabinSeatbeltsAlertSwitchToggle KeySimEvent = "CABIN_SEATBELTS_ALERT_SWITCH_TOGGLE"
	//KeyAntiskidBrakesToggle Turn the anti-skid braking system on or off.
	KeyAntiskidBrakesToggle KeySimEvent = "ANTISKID_BRAKES_TOGGLE"
	//KeyGpwsSwitchToggle Turn the g round proximity warning system (GPWS) on or off.
	KeyGpwsSwitchToggle KeySimEvent = "GPWS_SWITCH_TOGGLE"
	//KeyVideoRecordToggle Turn on or off the video recording feature. This records uncompressed AVI format files to: %USERPROFILE%\Documents\My Videos
	KeyVideoRecordToggle KeySimEvent = "VIDEO_RECORD_TOGGLE"
	//KeyToggleAirportNameDisplay Turn on or off the airport name.
	KeyToggleAirportNameDisplay KeySimEvent = "TOGGLE_AIRPORT_NAME_DISPLAY"
	//KeyCaptureScreenshot Capture the current view as a screenshot. Which will be saved to a bmp file in: %USERPROFILE%\Documents\My Pictures
	KeyCaptureScreenshot KeySimEvent = "CAPTURE_SCREENSHOT"
	//KeyMouseLookToggle Switch Mouse Look mode on or off. Mouse Look mode enables a user to control their view using the mouse, and holding down the space bar.
	KeyMouseLookToggle KeySimEvent = "MOUSE_LOOK_TOGGLE"
	//KeyYaxisInvertToggle Switch inversion of Y axis controls on or off.
	KeyYaxisInvertToggle KeySimEvent = "YAXIS_INVERT_TOGGLE"
	//KeyAutocoordToggle Turn the automatic rudder control feature on or off.	 Freezing position
	KeyAutocoordToggle KeySimEvent = "AUTORUDDER_TOGGLE"
	//KeyFlyByWireElacToggle Turn on or off the fly by wire Elevators and Ailerons computer.
	KeyFlyByWireElacToggle KeySimEvent = "FLY_BY_WIRE_ELAC_TOGGLE"
	//KeyFlyByWireFacToggle Turn on or off the fly by wire Flight Augmentation computer.
	KeyFlyByWireFacToggle KeySimEvent = "FLY_BY_WIRE_FAC_TOGGLE"
	//KeyFlyByWireSecToggle Turn on or off the fly by wire Spoilers and Elevators computer.	 G1000 Keys (Primary Flight Display)
	KeyFlyByWireSecToggle KeySimEvent = "FLY_BY_WIRE_SEC_TOGGLE"
	//KeyManualFuelPressurePump Activate the manual fuel pressure pump.	 Nose wheel steering
	KeyManualFuelPressurePump KeySimEvent = "MANUAL_FUEL_PRESSURE_PUMP"
	//KeySteeringInc Increments the nose wheel steering position by 5 percent.
	KeySteeringInc KeySimEvent = "STEERING_INC"
	//KeySteeringDec Decrements the nose wheel steering position by 5 percent.
	KeySteeringDec KeySimEvent = "STEERING_DEC"
	//KeySteeringSet Sets the value of the nose wheel steering position. Zero is straight ahead (-16383, far left +16383, far right).	 Cabin pressurization
	KeySteeringSet KeySimEvent = "STEERING_SET"
	//KeyFreezeLatitudeLongitudeToggle Turns the freezing of the lat/lon position of the aircraft (either user or AI controlled) on or off. If this key event is set, it means that the latitude and longitude of the aircraft are not being controlled by Prepar3D, so enabling, for example, a SimConnect client to control the position of the aircraft. This can also apply to altitude and attitude. Refer to the simulation variables: IS LATITUDE LONGITUDE FREEZE ON, IS ALTITUDE FREEZE ON, and IS ATTITUDE FREEZE ON Refer also to the SimConnect_AIReleaseControl function.
	KeyFreezeLatitudeLongitudeToggle KeySimEvent = "FREEZE_LATITUDE_LONGITUDE_TOGGLE"
	//KeyFreezeLatitudeLongitudeSet Freezes the lat/lon position of the aircraft.
	KeyFreezeLatitudeLongitudeSet KeySimEvent = "FREEZE_LATITUDE_LONGITUDE_SET"
	//KeyFreezeAltitudeToggle Turns the freezing of the altitude of the aircraft on or off.
	KeyFreezeAltitudeToggle KeySimEvent = "FREEZE_ALTITUDE_TOGGLE"
	//KeyFreezeAltitudeSet Freezes the altitude of the aircraft..
	KeyFreezeAltitudeSet KeySimEvent = "FREEZE_ALTITUDE_SET"
	//KeyFreezeAttitudeToggle Turns the freezing of the attitude (pitch, bank and heading) of the aircraft on or off.
	KeyFreezeAttitudeToggle KeySimEvent = "FREEZE_ATTITUDE_TOGGLE"
	//KeyFreezeAttitudeSet Freezes the attitude (pitch, bank and heading) of the aircraft.
	KeyFreezeAttitudeSet KeySimEvent = "FREEZE_ATTITUDE_SET"
	//KeyPressurizationPressureAltInc Increases the altitude that the cabin is pressurized to.
	KeyPressurizationPressureAltInc KeySimEvent = "PRESSURIZATION_PRESSURE_ALT_INC"
	//KeyPressurizationPressureAltDec Decreases the altitude that the cabin is pressurized to.
	KeyPressurizationPressureAltDec KeySimEvent = "PRESSURIZATION_PRESSURE_ALT_DEC"
	//KeyPressurizationClimbRateInc Sets the rate at which cabin pressurization is increased.
	KeyPressurizationClimbRateInc KeySimEvent = "PRESSURIZATION_CLIMB_RATE_INC"
	//KeyPressurizationClimbRateDec Sets the rate at which cabin pressurization is decreased.
	KeyPressurizationClimbRateDec KeySimEvent = "PRESSURIZATION_CLIMB_RATE_DEC"
	//KeyPressurizationPressureDumpSwtich Sets the cabin pressure to the outside air pressure.	 Catapult launches
	KeyPressurizationPressureDumpSwtich KeySimEvent = "PRESSURIZATION_PRESSURE_DUMP_SWTICH"
	//KeyFuelSelectorLeftMain Sets the fuel selector. Fuel will be taken in the order left tip, left aux, then main fuel tanks.
	KeyFuelSelectorLeftMain KeySimEvent = "FUEL_SELECTOR_LEFT_MAIN"
	//KeyFuelSelector2LeftMain Sets the fuel selector for engine 2.
	KeyFuelSelector2LeftMain KeySimEvent = "FUEL_SELECTOR_2_LEFT_MAIN"
	//KeyFuelSelector3LeftMain Sets the fuel selector for engine 3.
	KeyFuelSelector3LeftMain KeySimEvent = "FUEL_SELECTOR_3_LEFT_MAIN"
	//KeyFuelSelector4LeftMain Sets the fuel selector for engine 4.
	KeyFuelSelector4LeftMain KeySimEvent = "FUEL_SELECTOR_4_LEFT_MAIN"
	//KeyFuelSelectorRightMain Sets the fuel selector. Fuel will be taken in the order right tip, right aux, then main fuel tanks.
	KeyFuelSelectorRightMain KeySimEvent = "FUEL_SELECTOR_RIGHT_MAIN"
	//KeyFuelSelector2RightMain Sets the fuel selector for engine 2.
	KeyFuelSelector2RightMain KeySimEvent = "FUEL_SELECTOR_2_RIGHT_MAIN"
	//KeyFuelSelector3RightMain Sets the fuel selector for engine 3.
	KeyFuelSelector3RightMain KeySimEvent = "FUEL_SELECTOR_3_RIGHT_MAIN"
	//KeyFuelSelector4RightMain Sets the fuel selector for engine 4.
	KeyFuelSelector4RightMain KeySimEvent = "FUEL_SELECTOR_4_RIGHT_MAIN"
	//KeyPointOfInterestTogglePointer Turn the point-of-interest indicator (often a light beam) on or off. Refer to the SimDirector documentation.
	KeyPointOfInterestTogglePointer KeySimEvent = "POINT_OF_INTEREST_TOGGLE_POINTER"
	//KeyPointOfInterestCyclePrevious Change the current point-of-interest to the previous point-of-interest.
	KeyPointOfInterestCyclePrevious KeySimEvent = "POINT_OF_INTEREST_CYCLE_PREVIOUS"
	//KeyPointOfInterestCycleNext Change the current point-of-interest to the next point-of-interest.
	KeyPointOfInterestCycleNext KeySimEvent = "POINT_OF_INTEREST_CYCLE_NEXT"
	//KeyG1000PfdFlightplanButton The primary flight display (PFD) should display its current flight plan.
	KeyG1000PfdFlightplanButton KeySimEvent = "G1000_PFD_FLIGHTPLAN_BUTTON"
	//KeyG1000PfdProcedureButton Turn to the Procedure page.
	KeyG1000PfdProcedureButton KeySimEvent = "G1000_PFD_PROCEDURE_BUTTON"
	//KeyG1000PfdZoominButton Zoom in on the current map.
	KeyG1000PfdZoominButton KeySimEvent = "G1000_PFD_ZOOMIN_BUTTON"
	//KeyG1000PfdZoomoutButton Zoom out on the current map.
	KeyG1000PfdZoomoutButton KeySimEvent = "G1000_PFD_ZOOMOUT_BUTTON"
	//KeyG1000PfdDirecttoButton Turn to the Direct To page.
	KeyG1000PfdDirecttoButton KeySimEvent = "G1000_PFD_DIRECTTO_BUTTON"
	//KeyG1000PfdMenuButton If a segmented flight plan is highlighted, activates the associated menu.
	KeyG1000PfdMenuButton KeySimEvent = "G1000_PFD_MENU_BUTTON"
	//KeyG1000PfdClearButton Clears the current input.
	KeyG1000PfdClearButton KeySimEvent = "G1000_PFD_CLEAR_BUTTON"
	//KeyG1000PfdEnterButton Enters the current input.
	KeyG1000PfdEnterButton KeySimEvent = "G1000_PFD_ENTER_BUTTON"
	//KeyG1000PfdCursorButton Turns on or off a screen cursor.
	KeyG1000PfdCursorButton KeySimEvent = "G1000_PFD_CURSOR_BUTTON"
	//KeyG1000PfdGroupKnobInc Step up through the page groups.
	KeyG1000PfdGroupKnobInc KeySimEvent = "G1000_PFD_GROUP_KNOB_INC"
	//KeyG1000PfdGroupKnobDec Step down through the page groups.
	KeyG1000PfdGroupKnobDec KeySimEvent = "G1000_PFD_GROUP_KNOB_DEC"
	//KeyG1000PfdPageKnobInc Step up through the individual pages.
	KeyG1000PfdPageKnobInc KeySimEvent = "G1000_PFD_PAGE_KNOB_INC"
	//KeyG1000PfdPageKnobDec Step down through the individual pages.
	KeyG1000PfdPageKnobDec KeySimEvent = "G1000_PFD_PAGE_KNOB_DEC"
	//KeyG1000PfdSoftkey1 Initiate the action for the icon displayed in the softkey position.	G1000 (Multi-function Display)
	KeyG1000PfdSoftkey1  KeySimEvent = "G1000_PFD_SOFTKEY1"
	KeyG1000PfdSoftkey2  KeySimEvent = "G1000_PFD_SOFTKEY2"
	KeyG1000PfdSoftkey3  KeySimEvent = "G1000_PFD_SOFTKEY3"
	KeyG1000PfdSoftkey4  KeySimEvent = "G1000_PFD_SOFTKEY4"
	KeyG1000PfdSoftkey5  KeySimEvent = "G1000_PFD_SOFTKEY5"
	KeyG1000PfdSoftkey6  KeySimEvent = "G1000_PFD_SOFTKEY6"
	KeyG1000PfdSoftkey7  KeySimEvent = "G1000_PFD_SOFTKEY7"
	KeyG1000PfdSoftkey8  KeySimEvent = "G1000_PFD_SOFTKEY8"
	KeyG1000PfdSoftkey9  KeySimEvent = "G1000_PFD_SOFTKEY9"
	KeyG1000PfdSoftkey10 KeySimEvent = "G1000_PFD_SOFTKEY10"
	KeyG1000PfdSoftkey11 KeySimEvent = "G1000_PFD_SOFTKEY11"
	KeyG1000PfdSoftkey12 KeySimEvent = "G1000_PFD_SOFTKEY12"
	//KeyG1000MfdFlightplanButton The multifunction display (MFD) should display its current flight plan.
	KeyG1000MfdFlightplanButton KeySimEvent = "G1000_MFD_FLIGHTPLAN_BUTTON"
	//KeyG1000MfdProcedureButton Turn to the Procedure page.
	KeyG1000MfdProcedureButton KeySimEvent = "G1000_MFD_PROCEDURE_BUTTON"
	//KeyG1000MfdZoominButton Zoom in on the current map.
	KeyG1000MfdZoominButton KeySimEvent = "G1000_MFD_ZOOMIN_BUTTON"
	//KeyG1000MfdZoomoutButton Zoom out on the current map.
	KeyG1000MfdZoomoutButton KeySimEvent = "G1000_MFD_ZOOMOUT_BUTTON"
	//KeyG1000MfdDirecttoButton Turn to the Direct To page.
	KeyG1000MfdDirecttoButton KeySimEvent = "G1000_MFD_DIRECTTO_BUTTON"
	//KeyG1000MfdMenuButton If a segmented flight plan is highlighted, activates the associated menu.
	KeyG1000MfdMenuButton KeySimEvent = "G1000_MFD_MENU_BUTTON"
	//KeyG1000MfdClearButton Clears the current input.
	KeyG1000MfdClearButton KeySimEvent = "G1000_MFD_CLEAR_BUTTON"
	//KeyG1000MfdEnterButton Enters the current input.
	KeyG1000MfdEnterButton KeySimEvent = "G1000_MFD_ENTER_BUTTON"
	//KeyG1000MfdCursorButton Turns on or off a screen cursor.
	KeyG1000MfdCursorButton KeySimEvent = "G1000_MFD_CURSOR_BUTTON"
	//KeyG1000MfdGroupKnobInc Step up through the page groups.
	KeyG1000MfdGroupKnobInc KeySimEvent = "G1000_MFD_GROUP_KNOB_INC"
	//KeyG1000MfdGroupKnobDec Step down through the page groups.
	KeyG1000MfdGroupKnobDec KeySimEvent = "G1000_MFD_GROUP_KNOB_DEC"
	//KeyG1000MfdPageKnobInc Step up through the individual pages.
	KeyG1000MfdPageKnobInc KeySimEvent = "G1000_MFD_PAGE_KNOB_INC"
	//KeyG1000MfdPageKnobDec Step down through the individual pages.
	KeyG1000MfdPageKnobDec KeySimEvent = "G1000_MFD_PAGE_KNOB_DEC"
	//KeyG1000MfdSoftkey1 Initiate the action for the icon displayed in the softkey position.
	KeyG1000MfdSoftkey1  KeySimEvent = "G1000_MFD_SOFTKEY1"
	KeyG1000MfdSoftkey2  KeySimEvent = "G1000_MFD_SOFTKEY2"
	KeyG1000MfdSoftkey3  KeySimEvent = "G1000_MFD_SOFTKEY3"
	KeyG1000MfdSoftkey4  KeySimEvent = "G1000_MFD_SOFTKEY4"
	KeyG1000MfdSoftkey5  KeySimEvent = "G1000_MFD_SOFTKEY5"
	KeyG1000MfdSoftkey6  KeySimEvent = "G1000_MFD_SOFTKEY6"
	KeyG1000MfdSoftkey7  KeySimEvent = "G1000_MFD_SOFTKEY7"
	KeyG1000MfdSoftkey8  KeySimEvent = "G1000_MFD_SOFTKEY8"
	KeyG1000MfdSoftkey9  KeySimEvent = "G1000_MFD_SOFTKEY9"
	KeyG1000MfdSoftkey10 KeySimEvent = "G1000_MFD_SOFTKEY10"
	KeyG1000MfdSoftkey11 KeySimEvent = "G1000_MFD_SOFTKEY11"
	KeyG1000MfdSoftkey12 KeySimEvent = "G1000_MFD_SOFTKEY12"
	//KeyThrottleFull Set throttles max
	KeyThrottleFull KeySimEvent = "THROTTLE_FULL"
	//KeyThrottleIncr Increment throttles
	KeyThrottleIncr KeySimEvent = "THROTTLE_INCR"
	//KeyThrottleIncrSmall Increment throttles small
	KeyThrottleIncrSmall KeySimEvent = "THROTTLE_INCR_SMALL"
	//KeyThrottleDecr Decrement throttles
	KeyThrottleDecr KeySimEvent = "THROTTLE_DECR"
	//KeyThrottleDecrSmall Decrease throttles small
	KeyThrottleDecrSmall KeySimEvent = "THROTTLE_DECR_SMALL"
	//KeyThrottleCut Set throttles to idle
	KeyThrottleCut KeySimEvent = "THROTTLE_CUT"
	//KeyIncreaseThrottle Increment throttles
	KeyIncreaseThrottle KeySimEvent = "INCREASE_THROTTLE"
	//KeyDecreaseThrottle Decrement throttles
	KeyDecreaseThrottle KeySimEvent = "DECREASE_THROTTLE"
	//KeyThrottleSet Set throttles exactly (0- 16383)
	KeyThrottleSet KeySimEvent = "THROTTLE_SET"
	//KeyAxisThrottleSet Set throttles (0- 16383)	 (Pilot only, transmitted to Co-pilot if in a helicopter, not-transmitted otherwise).
	KeyAxisThrottleSet KeySimEvent = "AXIS_THROTTLE_SET"
	//KeyThrottle1Set Set throttle 1 exactly (0 to 16383)
	KeyThrottle1Set KeySimEvent = "THROTTLE1_SET"
	//KeyThrottle2Set Set throttle 2 exactly (0 to 16383)
	KeyThrottle2Set KeySimEvent = "THROTTLE2_SET"
	//KeyThrottle3Set Set throttle 3 exactly (0 to 16383)
	KeyThrottle3Set KeySimEvent = "THROTTLE3_SET"
	//KeyThrottle4Set Set throttle 4 exactly (0 to 16383)
	KeyThrottle4Set KeySimEvent = "THROTTLE4_SET"
	//KeyThrottle1Full Set throttle 1 max
	KeyThrottle1Full KeySimEvent = "THROTTLE1_FULL"
	//KeyThrottle1Incr Increment throttle 1
	KeyThrottle1Incr KeySimEvent = "THROTTLE1_INCR"
	//KeyThrottle1IncrSmall Increment throttle 1 small
	KeyThrottle1IncrSmall KeySimEvent = "THROTTLE1_INCR_SMALL"
	//KeyThrottle1Decr Decrement throttle 1
	KeyThrottle1Decr KeySimEvent = "THROTTLE1_DECR"
	//KeyThrottle1Cut Set throttle 1 to idle
	KeyThrottle1Cut KeySimEvent = "THROTTLE1_CUT"
	//KeyThrottle2Full Set throttle 2 max
	KeyThrottle2Full KeySimEvent = "THROTTLE2_FULL"
	//KeyThrottle2Incr Increment throttle 2
	KeyThrottle2Incr KeySimEvent = "THROTTLE2_INCR"
	//KeyThrottle2IncrSmall Increment throttle 2 small
	KeyThrottle2IncrSmall KeySimEvent = "THROTTLE2_INCR_SMALL"
	//KeyThrottle2Decr Decrement throttle 2
	KeyThrottle2Decr KeySimEvent = "THROTTLE2_DECR"
	//KeyThrottle2Cut Set throttle 2 to idle
	KeyThrottle2Cut KeySimEvent = "THROTTLE2_CUT"
	//KeyThrottle3Full Set throttle 3 max
	KeyThrottle3Full KeySimEvent = "THROTTLE3_FULL"
	//KeyThrottle3Incr Increment throttle 3
	KeyThrottle3Incr KeySimEvent = "THROTTLE3_INCR"
	//KeyThrottle3IncrSmall Increment throttle 3 small
	KeyThrottle3IncrSmall KeySimEvent = "THROTTLE3_INCR_SMALL"
	//KeyThrottle3Decr Decrement throttle 3
	KeyThrottle3Decr KeySimEvent = "THROTTLE3_DECR"
	//KeyThrottle3Cut Set throttle 3 to idle
	KeyThrottle3Cut KeySimEvent = "THROTTLE3_CUT"
	//KeyThrottle4Full Set throttle 1 max
	KeyThrottle4Full KeySimEvent = "THROTTLE4_FULL"
	//KeyThrottle4Incr Increment throttle 4
	KeyThrottle4Incr KeySimEvent = "THROTTLE4_INCR"
	//KeyThrottle4IncrSmall Increment throttle 4 small
	KeyThrottle4IncrSmall KeySimEvent = "THROTTLE4_INCR_SMALL"
	//KeyThrottle4Decr Decrement throttle 4
	KeyThrottle4Decr KeySimEvent = "THROTTLE4_DECR"
	//KeyThrottle4Cut Set throttle 4 to idle
	KeyThrottle4Cut KeySimEvent = "THROTTLE4_CUT"
	//KeyThrottle10 Set throttles to 10%
	KeyThrottle10 KeySimEvent = "THROTTLE_10"
	//KeyThrottle20 Set throttles to 20%
	KeyThrottle20 KeySimEvent = "THROTTLE_20"
	//KeyThrottle30 Set throttles to 30%
	KeyThrottle30 KeySimEvent = "THROTTLE_30"
	//KeyThrottle40 Set throttles to 40%
	KeyThrottle40 KeySimEvent = "THROTTLE_40"
	//KeyThrottle50 Set throttles to 50%
	KeyThrottle50 KeySimEvent = "THROTTLE_50"
	//KeyThrottle60 Set throttles to 60%
	KeyThrottle60 KeySimEvent = "THROTTLE_60"
	//KeyThrottle70 Set throttles to 70%
	KeyThrottle70 KeySimEvent = "THROTTLE_70"
	//KeyThrottle80 Set throttles to 80%
	KeyThrottle80 KeySimEvent = "THROTTLE_80"
	//KeyThrottle90 Set throttles to 90%
	KeyThrottle90 KeySimEvent = "THROTTLE_90"
	//KeyAxisThrottle1Set Set throttle 1 exactly (-16383 - +16383)
	KeyAxisThrottle1Set KeySimEvent = "AXIS_THROTTLE1_SET"
	//KeyAxisThrottle2Set Set throttle 2 exactly (-16383 - +16383)
	KeyAxisThrottle2Set KeySimEvent = "AXIS_THROTTLE2_SET"
	//KeyAxisThrottle3Set Set throttle 3 exactly (-16383 - +16383)
	KeyAxisThrottle3Set KeySimEvent = "AXIS_THROTTLE3_SET"
	//KeyAxisThrottle4Set Set throttle 4 exactly (-16383 - +16383)
	KeyAxisThrottle4Set KeySimEvent = "AXIS_THROTTLE4_SET"
	//KeyThrottle1DecrSmall Decrease throttle 1 small
	KeyThrottle1DecrSmall KeySimEvent = "THROTTLE1_DECR_SMALL"
	//KeyThrottle2DecrSmall Decrease throttle 2 small
	KeyThrottle2DecrSmall KeySimEvent = "THROTTLE2_DECR_SMALL"
	//KeyThrottle3DecrSmall Decrease throttle 3 small
	KeyThrottle3DecrSmall KeySimEvent = "THROTTLE3_DECR_SMALL"
	//KeyThrottle4DecrSmall Decrease throttle 4 small
	KeyThrottle4DecrSmall KeySimEvent = "THROTTLE4_DECR_SMALL"
	//KeyPropPitchDecrSmall Decrease prop levers small
	KeyPropPitchDecrSmall KeySimEvent = "PROP_PITCH_DECR_SMALL"
	//KeyPropPitch1DecrSmall Decrease prop lever 1 small
	KeyPropPitch1DecrSmall KeySimEvent = "PROP_PITCH1_DECR_SMALL"
	//KeyPropPitch2DecrSmall Decrease prop lever 2 small
	KeyPropPitch2DecrSmall KeySimEvent = "PROP_PITCH2_DECR_SMALL"
	//KeyPropPitch3DecrSmall Decrease prop lever 3 small
	KeyPropPitch3DecrSmall KeySimEvent = "PROP_PITCH3_DECR_SMALL"
	//KeyPropPitch4DecrSmall Decrease prop lever 4 small
	KeyPropPitch4DecrSmall KeySimEvent = "PROP_PITCH4_DECR_SMALL"
	//KeyMixture1Rich Set mixture lever 1 to max rich
	KeyMixture1Rich KeySimEvent = "MIXTURE1_RICH"
	//KeyMixture1Incr Increment mixture lever 1
	KeyMixture1Incr KeySimEvent = "MIXTURE1_INCR"
	//KeyMixture1IncrSmall Increment mixture lever 1 small
	KeyMixture1IncrSmall KeySimEvent = "MIXTURE1_INCR_SMALL"
	//KeyMixture1Decr Decrement mixture lever 1
	KeyMixture1Decr KeySimEvent = "MIXTURE1_DECR"
	//KeyMixture1Lean Set mixture lever 1 to max lean
	KeyMixture1Lean KeySimEvent = "MIXTURE1_LEAN"
	//KeyMixture2Rich Set mixture lever 2 to max rich
	KeyMixture2Rich KeySimEvent = "MIXTURE2_RICH"
	//KeyMixture2Incr Increment mixture lever 2
	KeyMixture2Incr KeySimEvent = "MIXTURE2_INCR"
	//KeyMixture2IncrSmall Increment mixture lever 2 small
	KeyMixture2IncrSmall KeySimEvent = "MIXTURE2_INCR_SMALL"
	//KeyMixture2Decr Decrement mixture lever 2
	KeyMixture2Decr KeySimEvent = "MIXTURE2_DECR"
	//KeyMixture2Lean Set mixture lever 2 to max lean
	KeyMixture2Lean KeySimEvent = "MIXTURE2_LEAN"
	//KeyMixture3Rich Set mixture lever 3 to max rich
	KeyMixture3Rich KeySimEvent = "MIXTURE3_RICH"
	//KeyMixture3Incr Increment mixture lever 3
	KeyMixture3Incr KeySimEvent = "MIXTURE3_INCR"
	//KeyMixture3IncrSmall Increment mixture lever 3 small
	KeyMixture3IncrSmall KeySimEvent = "MIXTURE3_INCR_SMALL"
	//KeyMixture3Decr Decrement mixture lever 3
	KeyMixture3Decr KeySimEvent = "MIXTURE3_DECR"
	//KeyMixture3Lean Set mixture lever 3 to max lean
	KeyMixture3Lean KeySimEvent = "MIXTURE3_LEAN"
	//KeyMixture4Rich Set mixture lever 4 to max rich
	KeyMixture4Rich KeySimEvent = "MIXTURE4_RICH"
	//KeyMixture4Incr Increment mixture lever 4
	KeyMixture4Incr KeySimEvent = "MIXTURE4_INCR"
	//KeyMixture4IncrSmall Increment mixture lever 4 small
	KeyMixture4IncrSmall KeySimEvent = "MIXTURE4_INCR_SMALL"
	//KeyMixture4Decr Decrement mixture lever 4
	KeyMixture4Decr KeySimEvent = "MIXTURE4_DECR"
	//KeyMixture4Lean Set mixture lever 4 to max lean
	KeyMixture4Lean KeySimEvent = "MIXTURE4_LEAN"
	//KeyMixtureSet Set mixture levers to exact value (0 to 16383)
	KeyMixtureSet KeySimEvent = "MIXTURE_SET"
	//KeyMixtureRich Set mixture levers to max rich
	KeyMixtureRich KeySimEvent = "MIXTURE_RICH"
	//KeyMixtureIncr Increment mixture levers
	KeyMixtureIncr KeySimEvent = "MIXTURE_INCR"
	//KeyMixtureIncrSmall Increment mixture levers small
	KeyMixtureIncrSmall KeySimEvent = "MIXTURE_INCR_SMALL"
	//KeyMixtureDecr Decrement mixture levers
	KeyMixtureDecr KeySimEvent = "MIXTURE_DECR"
	//KeyMixtureLean Set mixture levers to max lean
	KeyMixtureLean KeySimEvent = "MIXTURE_LEAN"
	//KeyMixture1Set Set mixture lever 1 exact value (0 to 16383)
	KeyMixture1Set KeySimEvent = "MIXTURE1_SET"
	//KeyMixture2Set Set mixture lever 2 exact value (0 to 16383)
	KeyMixture2Set KeySimEvent = "MIXTURE2_SET"
	//KeyMixture3Set Set mixture lever 3 exact value (0 to 16383)
	KeyMixture3Set KeySimEvent = "MIXTURE3_SET"
	//KeyMixture4Set Set mixture lever 4 exact value (0 to 16383)
	KeyMixture4Set KeySimEvent = "MIXTURE4_SET"
	//KeyAxisMixtureSet Set mixture lever 1 exact value (-16383 to +16383)
	KeyAxisMixtureSet KeySimEvent = "AXIS_MIXTURE_SET"
	//KeyAxisMixture1Set Set mixture lever 1 exact value (-16383 to +16383)
	KeyAxisMixture1Set KeySimEvent = "AXIS_MIXTURE1_SET"
	//KeyAxisMixture2Set Set mixture lever 2 exact value (-16383 to +16383)
	KeyAxisMixture2Set KeySimEvent = "AXIS_MIXTURE2_SET"
	//KeyAxisMixture3Set Set mixture lever 3 exact value (-16383 to +16383)
	KeyAxisMixture3Set KeySimEvent = "AXIS_MIXTURE3_SET"
	//KeyAxisMixture4Set Set mixture lever 4 exact value (-16383 to +16383)
	KeyAxisMixture4Set KeySimEvent = "AXIS_MIXTURE4_SET"
	//KeyMixtureSetBest Set mixture levers to current best power setting
	KeyMixtureSetBest KeySimEvent = "MIXTURE_SET_BEST"
	//KeyMixtureDecrSmall Decrement mixture levers small
	KeyMixtureDecrSmall KeySimEvent = "MIXTURE_DECR_SMALL"
	//KeyMixture1DecrSmall Decrement mixture lever 1 small
	KeyMixture1DecrSmall KeySimEvent = "MIXTURE1_DECR_SMALL"
	//KeyMixture2DecrSmall Decrement mixture lever 4 small
	KeyMixture2DecrSmall KeySimEvent = "MIXTURE2_DECR_SMALL"
	//KeyMixture3DecrSmall Decrement mixture lever 4 small
	KeyMixture3DecrSmall KeySimEvent = "MIXTURE3_DECR_SMALL"
	//KeyMixture4DecrSmall Decrement mixture lever 4 small
	KeyMixture4DecrSmall KeySimEvent = "MIXTURE4_DECR_SMALL"
	//KeyPropPitchSet Set prop pitch levers (0 to 16383)
	KeyPropPitchSet KeySimEvent = "PROP_PITCH_SET"
	//KeyPropPitchLo Set prop pitch levers max (lo pitch)
	KeyPropPitchLo KeySimEvent = "PROP_PITCH_LO"
	//KeyPropPitchIncr Increment prop pitch levers
	KeyPropPitchIncr KeySimEvent = "PROP_PITCH_INCR"
	//KeyPropPitchIncrSmall Increment prop pitch levers small
	KeyPropPitchIncrSmall KeySimEvent = "PROP_PITCH_INCR_SMALL"
	//KeyPropPitchDecr Decrement prop pitch levers
	KeyPropPitchDecr KeySimEvent = "PROP_PITCH_DECR"
	//KeyPropPitchHi Set prop pitch levers min (hi pitch)
	KeyPropPitchHi KeySimEvent = "PROP_PITCH_HI"
	//KeyPropPitch1Set Set prop pitch lever 1 exact value (0 to 16383)
	KeyPropPitch1Set KeySimEvent = "PROP_PITCH1_SET"
	//KeyPropPitch2Set Set prop pitch lever 2 exact value (0 to 16383)
	KeyPropPitch2Set KeySimEvent = "PROP_PITCH2_SET"
	//KeyPropPitch3Set Set prop pitch lever 3 exact value (0 to 16383)
	KeyPropPitch3Set KeySimEvent = "PROP_PITCH3_SET"
	//KeyPropPitch4Set Set prop pitch lever 4 exact value (0 to 16383)
	KeyPropPitch4Set KeySimEvent = "PROP_PITCH4_SET"
	//KeyPropPitch1Lo Set prop pitch lever 1 max (lo pitch)
	KeyPropPitch1Lo KeySimEvent = "PROP_PITCH1_LO"
	//KeyPropPitch1Incr Increment prop pitch lever 1
	KeyPropPitch1Incr KeySimEvent = "PROP_PITCH1_INCR"
	//KeyPropPitch1IncrSmall Increment prop pitch lever 1 small
	KeyPropPitch1IncrSmall KeySimEvent = "PROP_PITCH1_INCR_SMALL"
	//KeyPropPitch1Decr Decrement prop pitch lever 1
	KeyPropPitch1Decr KeySimEvent = "PROP_PITCH1_DECR"
	//KeyPropPitch1Hi Set prop pitch lever 1 min (hi pitch)
	KeyPropPitch1Hi KeySimEvent = "PROP_PITCH1_HI"
	//KeyPropPitch2Lo Set prop pitch lever 2 max (lo pitch)
	KeyPropPitch2Lo KeySimEvent = "PROP_PITCH2_LO"
	//KeyPropPitch2Incr Increment prop pitch lever 2
	KeyPropPitch2Incr KeySimEvent = "PROP_PITCH2_INCR"
	//KeyPropPitch2IncrSmall Increment prop pitch lever 2 small
	KeyPropPitch2IncrSmall KeySimEvent = "PROP_PITCH2_INCR_SMALL"
	//KeyPropPitch2Decr Decrement prop pitch lever 2
	KeyPropPitch2Decr KeySimEvent = "PROP_PITCH2_DECR"
	//KeyPropPitch2Hi Set prop pitch lever 2 min (hi pitch)
	KeyPropPitch2Hi KeySimEvent = "PROP_PITCH2_HI"
	//KeyPropPitch3Lo Set prop pitch lever 3 max (lo pitch)
	KeyPropPitch3Lo KeySimEvent = "PROP_PITCH3_LO"
	//KeyPropPitch3Incr Increment prop pitch lever 3
	KeyPropPitch3Incr KeySimEvent = "PROP_PITCH3_INCR"
	//KeyPropPitch3IncrSmall Increment prop pitch lever 3 small
	KeyPropPitch3IncrSmall KeySimEvent = "PROP_PITCH3_INCR_SMALL"
	//KeyPropPitch3Decr Decrement prop pitch lever 3
	KeyPropPitch3Decr KeySimEvent = "PROP_PITCH3_DECR"
	//KeyPropPitch3Hi Set prop pitch lever 3 min (hi pitch)
	KeyPropPitch3Hi KeySimEvent = "PROP_PITCH3_HI"
	//KeyPropPitch4Lo Set prop pitch lever 4 max (lo pitch)
	KeyPropPitch4Lo KeySimEvent = "PROP_PITCH4_LO"
	//KeyPropPitch4Incr Increment prop pitch lever 4
	KeyPropPitch4Incr KeySimEvent = "PROP_PITCH4_INCR"
	//KeyPropPitch4IncrSmall Increment prop pitch lever 4 small
	KeyPropPitch4IncrSmall KeySimEvent = "PROP_PITCH4_INCR_SMALL"
	//KeyPropPitch4Decr Decrement prop pitch lever 4
	KeyPropPitch4Decr KeySimEvent = "PROP_PITCH4_DECR"
	//KeyPropPitch4Hi Set prop pitch lever 4 min (hi pitch)
	KeyPropPitch4Hi KeySimEvent = "PROP_PITCH4_HI"
	//KeyAxisPropellerSet Set propeller levers exact value (-16383 to +16383)
	KeyAxisPropellerSet KeySimEvent = "AXIS_PROPELLER_SET"
	//KeyAxisPropeller1Set Set propeller lever 1 exact value (-16383 to +16383)
	KeyAxisPropeller1Set KeySimEvent = "AXIS_PROPELLER1_SET"
	//KeyAxisPropeller2Set Set propeller lever 2 exact value (-16383 to +16383)
	KeyAxisPropeller2Set KeySimEvent = "AXIS_PROPELLER2_SET"
	//KeyAxisPropeller3Set Set propeller lever 3 exact value (-16383 to +16383)
	KeyAxisPropeller3Set KeySimEvent = "AXIS_PROPELLER3_SET"
	//KeyAxisPropeller4Set Set propeller lever 4 exact value (-16383 to +16383)
	KeyAxisPropeller4Set KeySimEvent = "AXIS_PROPELLER4_SET"
	//KeyJetStarter Selects jet engine starter (for +/- sequence)
	KeyJetStarter KeySimEvent = "JET_STARTER"
	//KeyStarterSet Sets magnetos (0,1)
	KeyStarterSet KeySimEvent = "MAGNETO_SET"
	//KeyToggleStarter1 Toggle starter 1
	KeyToggleStarter1 KeySimEvent = "TOGGLE_STARTER1"
	//KeyToggleStarter2 Toggle starter 2
	KeyToggleStarter2 KeySimEvent = "TOGGLE_STARTER2"
	//KeyToggleStarter3 Toggle starter 3
	KeyToggleStarter3 KeySimEvent = "TOGGLE_STARTER3"
	//KeyToggleStarter4 Toggle starter 4
	KeyToggleStarter4 KeySimEvent = "TOGGLE_STARTER4"
	//KeyToggleAllStarters Toggle starters
	KeyToggleAllStarters KeySimEvent = "TOGGLE_ALL_STARTERS"
	//KeyEngineAutoStart Triggers auto-start
	KeyEngineAutoStart KeySimEvent = "ENGINE_AUTO_START"
	//KeyEngineAutoShutdown Triggers auto-shutdown
	KeyEngineAutoShutdown KeySimEvent = "ENGINE_AUTO_SHUTDOWN"
	//KeyMagneto Selects magnetos (for +/- sequence)
	KeyMagneto KeySimEvent = "MAGNETO"
	//KeyMagnetoDecr Decrease magneto switches positions
	KeyMagnetoDecr KeySimEvent = "MAGNETO_DECR"
	//KeyMagnetoIncr Increase magneto switches positions
	KeyMagnetoIncr KeySimEvent = "MAGNETO_INCR"
	//KeyMagneto1Off Set engine 1 magnetos off
	KeyMagneto1Off KeySimEvent = "MAGNETO1_OFF"
	//KeyMagneto1Right Toggle engine 1 right magneto	All aircraft
	KeyMagneto1Right KeySimEvent = "MAGNETO1_RIGHT"
	//KeyMagneto1Left Toggle engine 1 left magneto	All aircraft
	KeyMagneto1Left KeySimEvent = "MAGNETO1_LEFT"
	//KeyMagneto1Both Set engine 1 magnetos on
	KeyMagneto1Both KeySimEvent = "MAGNETO1_BOTH"
	//KeyMagneto1Start Set engine 1 magnetos on and toggle starter
	KeyMagneto1Start KeySimEvent = "MAGNETO1_START"
	//KeyMagneto2Off Set engine 2 magnetos off
	KeyMagneto2Off KeySimEvent = "MAGNETO2_OFF"
	//KeyMagneto2Right Toggle engine 2 right magneto	All aircraft
	KeyMagneto2Right KeySimEvent = "MAGNETO2_RIGHT"
	//KeyMagneto2Left Toggle engine 2 left magneto	All aircraft
	KeyMagneto2Left KeySimEvent = "MAGNETO2_LEFT"
	//KeyMagneto2Both Set engine 2 magnetos on
	KeyMagneto2Both KeySimEvent = "MAGNETO2_BOTH"
	//KeyMagneto2Start Set engine 2 magnetos on and toggle starter
	KeyMagneto2Start KeySimEvent = "MAGNETO2_START"
	//KeyMagneto3Off Set engine 3 magnetos off
	KeyMagneto3Off KeySimEvent = "MAGNETO3_OFF"
	//KeyMagneto3Right Toggle engine 3 right magneto	All aircraft
	KeyMagneto3Right KeySimEvent = "MAGNETO3_RIGHT"
	//KeyMagneto3Left Toggle engine 3 left magneto	All aircraft
	KeyMagneto3Left KeySimEvent = "MAGNETO3_LEFT"
	//KeyMagneto3Both Set engine 3 magnetos on
	KeyMagneto3Both KeySimEvent = "MAGNETO3_BOTH"
	//KeyMagneto3Start Set engine 3 magnetos on and toggle starter
	KeyMagneto3Start KeySimEvent = "MAGNETO3_START"
	//KeyMagneto4Off Set engine 4 magnetos off
	KeyMagneto4Off KeySimEvent = "MAGNETO4_OFF"
	//KeyMagneto4Right Toggle engine 4 right magneto	All aircraft
	KeyMagneto4Right KeySimEvent = "MAGNETO4_RIGHT"
	//KeyMagneto4Left Toggle engine 4 left magneto	All aircraft
	KeyMagneto4Left KeySimEvent = "MAGNETO4_LEFT"
	//KeyMagneto4Both Set engine 4 magnetos on
	KeyMagneto4Both KeySimEvent = "MAGNETO4_BOTH"
	//KeyMagneto4Start Set engine 4 magnetos on and toggle starter
	KeyMagneto4Start KeySimEvent = "MAGNETO4_START"
	//KeyMagnetoOff Set engine magnetos off
	KeyMagnetoOff KeySimEvent = "MAGNETO_OFF"
	//KeyMagnetoRight Set engine right magnetos on
	KeyMagnetoRight KeySimEvent = "MAGNETO_RIGHT"
	//KeyMagnetoLeft Set engine left magnetos on
	KeyMagnetoLeft KeySimEvent = "MAGNETO_LEFT"
	//KeyMagnetoBoth Set engine magnetos on
	KeyMagnetoBoth KeySimEvent = "MAGNETO_BOTH"
	//KeyMagnetoStart Set engine magnetos on and toggle starters
	KeyMagnetoStart KeySimEvent = "MAGNETO_START"
	//KeyMagneto1Decr Decrease engine 1 magneto switch position
	KeyMagneto1Decr KeySimEvent = "MAGNETO1_DECR"
	//KeyMagneto1Incr Increase engine 1 magneto switch position
	KeyMagneto1Incr KeySimEvent = "MAGNETO1_INCR"
	//KeyMagneto2Decr Decrease engine 2 magneto switch position
	KeyMagneto2Decr KeySimEvent = "MAGNETO2_DECR"
	//KeyMagneto2Incr Increase engine 2 magneto switch position
	KeyMagneto2Incr KeySimEvent = "MAGNETO2_INCR"
	//KeyMagneto3Decr Decrease engine 3 magneto switch position
	KeyMagneto3Decr KeySimEvent = "MAGNETO3_DECR"
	//KeyMagneto3Incr Increase engine 3 magneto switch position
	KeyMagneto3Incr KeySimEvent = "MAGNETO3_INCR"
	//KeyMagneto4Decr Decrease engine 4 magneto switch position
	KeyMagneto4Decr KeySimEvent = "MAGNETO4_DECR"
	//KeyMagneto4Incr Increase engine 4 magneto switch position
	KeyMagneto4Incr KeySimEvent = "MAGNETO4_INCR"
	//KeyMagneto1Set Set engine 1 magneto switch
	KeyMagneto1Set KeySimEvent = "MAGNETO1_SET"
	//KeyMagneto2Set Set engine 2 magneto switch
	KeyMagneto2Set KeySimEvent = "MAGNETO2_SET"
	//KeyMagneto3Set Set engine 3 magneto switch
	KeyMagneto3Set KeySimEvent = "MAGNETO3_SET"
	//KeyMagneto4Set Set engine 4 magneto switch
	KeyMagneto4Set KeySimEvent = "MAGNETO4_SET"
	//KeyAntiIceOn Sets anti-ice switches on
	KeyAntiIceOn KeySimEvent = "ANTI_ICE_ON"
	//KeyAntiIceOff Sets anti-ice switches off
	KeyAntiIceOff KeySimEvent = "ANTI_ICE_OFF"
	//KeyAntiIceSet Sets anti-ice switches from argument (0,1)
	KeyAntiIceSet KeySimEvent = "ANTI_ICE_SET"
	//KeyAntiIceToggle Toggle anti-ice switches
	KeyAntiIceToggle KeySimEvent = "ANTI_ICE_TOGGLE"
	//KeyAntiIceToggleEng1 Toggle engine 1 anti-ice switch
	KeyAntiIceToggleEng1 KeySimEvent = "ANTI_ICE_TOGGLE_ENG1"
	//KeyAntiIceToggleEng2 Toggle engine 2 anti-ice switch
	KeyAntiIceToggleEng2 KeySimEvent = "ANTI_ICE_TOGGLE_ENG2"
	//KeyAntiIceToggleEng3 Toggle engine 3 anti-ice switch
	KeyAntiIceToggleEng3 KeySimEvent = "ANTI_ICE_TOGGLE_ENG3"
	//KeyAntiIceToggleEng4 Toggle engine 4 anti-ice switch
	KeyAntiIceToggleEng4 KeySimEvent = "ANTI_ICE_TOGGLE_ENG4"
	//KeyAntiIceSetEng1 Sets engine 1 anti-ice switch (0,1)
	KeyAntiIceSetEng1 KeySimEvent = "ANTI_ICE_SET_ENG1"
	//KeyAntiIceSetEng2 Sets engine 2 anti-ice switch (0,1)
	KeyAntiIceSetEng2 KeySimEvent = "ANTI_ICE_SET_ENG2"
	//KeyAntiIceSetEng3 Sets engine 3 anti-ice switch (0,1)
	KeyAntiIceSetEng3 KeySimEvent = "ANTI_ICE_SET_ENG3"
	//KeyAntiIceSetEng4 Sets engine 4 anti-ice switch (0,1)
	KeyAntiIceSetEng4 KeySimEvent = "ANTI_ICE_SET_ENG4"
	//KeyToggleFuelValveAll Toggle engine fuel valves
	KeyToggleFuelValveAll KeySimEvent = "TOGGLE_FUEL_VALVE_ALL"
	//KeyToggleFuelValveEng1 Toggle engine 1 fuel valve	All aircraft
	KeyToggleFuelValveEng1 KeySimEvent = "TOGGLE_FUEL_VALVE_ENG1"
	//KeyToggleFuelValveEng2 Toggle engine 2 fuel valve	All aircraft
	KeyToggleFuelValveEng2 KeySimEvent = "TOGGLE_FUEL_VALVE_ENG2"
	//KeyToggleFuelValveEng3 Toggle engine 3 fuel valve	All aircraft
	KeyToggleFuelValveEng3 KeySimEvent = "TOGGLE_FUEL_VALVE_ENG3"
	//KeyToggleFuelValveEng4 Toggle engine 4 fuel valve	All aircraft
	KeyToggleFuelValveEng4 KeySimEvent = "TOGGLE_FUEL_VALVE_ENG4"
	//KeyCowlflap1Set Sets engine 1 cowl flap lever position (0 to 16383)
	KeyCowlflap1Set KeySimEvent = "COWLFLAP1_SET"
	//KeyCowlflap2Set Sets engine 2 cowl flap lever position (0 to 16383)
	KeyCowlflap2Set KeySimEvent = "COWLFLAP2_SET"
	//KeyCowlflap3Set Sets engine 3 cowl flap lever position (0 to 16383)
	KeyCowlflap3Set KeySimEvent = "COWLFLAP3_SET"
	//KeyCowlflap4Set Sets engine 4 cowl flap lever position (0 to 16383)
	KeyCowlflap4Set KeySimEvent = "COWLFLAP4_SET"
	//KeyIncCowlFlaps Increment cowl flap levers
	KeyIncCowlFlaps KeySimEvent = "INC_COWL_FLAPS"
	//KeyDecCowlFlaps Decrement cowl flap levers
	KeyDecCowlFlaps KeySimEvent = "DEC_COWL_FLAPS"
	//KeyIncCowlFlaps1 Increment engine 1 cowl flap lever
	KeyIncCowlFlaps1 KeySimEvent = "INC_COWL_FLAPS1"
	//KeyDecCowlFlaps1 Decrement engine 1 cowl flap lever
	KeyDecCowlFlaps1 KeySimEvent = "DEC_COWL_FLAPS1"
	//KeyIncCowlFlaps2 Increment engine 2 cowl flap lever
	KeyIncCowlFlaps2 KeySimEvent = "INC_COWL_FLAPS2"
	//KeyDecCowlFlaps2 Decrement engine 2 cowl flap lever
	KeyDecCowlFlaps2 KeySimEvent = "DEC_COWL_FLAPS2"
	//KeyIncCowlFlaps3 Increment engine 3 cowl flap lever
	KeyIncCowlFlaps3 KeySimEvent = "INC_COWL_FLAPS3"
	//KeyDecCowlFlaps3 Decrement engine 3 cowl flap lever
	KeyDecCowlFlaps3 KeySimEvent = "DEC_COWL_FLAPS3"
	//KeyIncCowlFlaps4 Increment engine 4 cowl flap lever
	KeyIncCowlFlaps4 KeySimEvent = "INC_COWL_FLAPS4"
	//KeyDecCowlFlaps4 Decrement engine 4 cowl flap lever
	KeyDecCowlFlaps4 KeySimEvent = "DEC_COWL_FLAPS4"
	//KeyFuelPump Toggle electric fuel pumps
	KeyFuelPump KeySimEvent = "FUEL_PUMP"
	//KeyToggleElectFuelPump Toggle electric fuel pumps
	KeyToggleElectFuelPump KeySimEvent = "TOGGLE_ELECT_FUEL_PUMP"
	//KeyToggleElectFuelPump1 Toggle engine 1 electric fuel pump	All aircraft
	KeyToggleElectFuelPump1 KeySimEvent = "TOGGLE_ELECT_FUEL_PUMP1"
	//KeyToggleElectFuelPump2 Toggle engine 2 electric fuel pump	All aircraft
	KeyToggleElectFuelPump2 KeySimEvent = "TOGGLE_ELECT_FUEL_PUMP2"
	//KeyToggleElectFuelPump3 Toggle engine 3 electric fuel pump	All aircraft
	KeyToggleElectFuelPump3 KeySimEvent = "TOGGLE_ELECT_FUEL_PUMP3"
	//KeyToggleElectFuelPump4 Toggle engine 4 electric fuel pump	All aircraft
	KeyToggleElectFuelPump4 KeySimEvent = "TOGGLE_ELECT_FUEL_PUMP4"
	//KeyEnginePrimer Trigger engine primers
	KeyEnginePrimer KeySimEvent = "ENGINE_PRIMER"
	//KeyTogglePrimer Trigger engine primers
	KeyTogglePrimer KeySimEvent = "TOGGLE_PRIMER"
	//KeyTogglePrimer1 Trigger engine 1 primer
	KeyTogglePrimer1 KeySimEvent = "TOGGLE_PRIMER1"
	//KeyTogglePrimer2 Trigger engine 2 primer
	KeyTogglePrimer2 KeySimEvent = "TOGGLE_PRIMER2"
	//KeyTogglePrimer3 Trigger engine 3 primer
	KeyTogglePrimer3 KeySimEvent = "TOGGLE_PRIMER3"
	//KeyTogglePrimer4 Trigger engine 4 primer
	KeyTogglePrimer4 KeySimEvent = "TOGGLE_PRIMER4"
	//KeyToggleFeatherSwitches Trigger propeller switches
	KeyToggleFeatherSwitches KeySimEvent = "TOGGLE_FEATHER_SWITCHES"
	//KeyToggleFeatherSwitch1 Trigger propeller 1 switch
	KeyToggleFeatherSwitch1 KeySimEvent = "TOGGLE_FEATHER_SWITCH_1"
	//KeyToggleFeatherSwitch2 Trigger propeller 2 switch
	KeyToggleFeatherSwitch2 KeySimEvent = "TOGGLE_FEATHER_SWITCH_2"
	//KeyToggleFeatherSwitch3 Trigger propeller 3 switch
	KeyToggleFeatherSwitch3 KeySimEvent = "TOGGLE_FEATHER_SWITCH_3"
	//KeyToggleFeatherSwitch4 Trigger propeller 4 switch
	KeyToggleFeatherSwitch4 KeySimEvent = "TOGGLE_FEATHER_SWITCH_4"
	//KeyTogglePropSync Turns propeller synchronization switch on
	KeyTogglePropSync KeySimEvent = "TOGGLE_PROPELLER_SYNC"
	//KeyToggleArmAutofeather Turns auto-feather arming switch on.
	KeyToggleArmAutofeather KeySimEvent = "TOGGLE_AUTOFEATHER_ARM"
	//KeyToggleAfterburner Toggles afterburners
	KeyToggleAfterburner KeySimEvent = "TOGGLE_AFTERBURNER"
	//KeyToggleAfterburner1 Toggles engine 1 afterburner
	KeyToggleAfterburner1 KeySimEvent = "TOGGLE_AFTERBURNER1"
	//KeyToggleAfterburner2 Toggles engine 2 afterburner
	KeyToggleAfterburner2 KeySimEvent = "TOGGLE_AFTERBURNER2"
	//KeyToggleAfterburner3 Toggles engine 3 afterburner
	KeyToggleAfterburner3 KeySimEvent = "TOGGLE_AFTERBURNER3"
	//KeyToggleAfterburner4 Toggles engine 4 afterburner
	KeyToggleAfterburner4 KeySimEvent = "TOGGLE_AFTERBURNER4"
	//KeyEngine Sets engines for 1,2,3,4 selection (to be followed by SELECT_n)
	KeyEngine KeySimEvent = "ENGINE"
	//KeySpoilersToggle Toggles spoiler handle 	All aircraft
	KeySpoilersToggle KeySimEvent = "SPOILERS_TOGGLE"
	//KeyFlapsUp Sets flap handle to full retract position	All aircraft
	KeyFlapsUp KeySimEvent = "FLAPS_UP"
	//KeyFlaps1 Sets flap handle to first extension position	All aircraft
	KeyFlaps1 KeySimEvent = "FLAPS_1"
	//KeyFlaps2 Sets flap handle to second extension position	All aircraft
	KeyFlaps2 KeySimEvent = "FLAPS_2"
	//KeyFlaps3 Sets flap handle to third extension position	All aircraft
	KeyFlaps3 KeySimEvent = "FLAPS_3"
	//KeyFlapsDown Sets flap handle to full extension position	All aircraft
	KeyFlapsDown KeySimEvent = "FLAPS_DOWN"
	//KeyElevTrimDn Increments elevator trim down
	KeyElevTrimDn KeySimEvent = "ELEV_TRIM_DN"
	//KeyElevDown Increments elevator down	 (Pilot only).
	KeyElevDown KeySimEvent = "ELEV_DOWN"
	//KeyAileronsLeft Increments ailerons left	 (Pilot only).
	KeyAileronsLeft KeySimEvent = "AILERONS_LEFT"
	//KeyCenterAilerRudder Centers aileron and rudder positions
	KeyCenterAilerRudder KeySimEvent = "CENTER_AILER_RUDDER"
	//KeyAileronsRight Increments ailerons right	 (Pilot only).
	KeyAileronsRight KeySimEvent = "AILERONS_RIGHT"
	//KeyElevTrimUp Increment elevator trim up
	KeyElevTrimUp KeySimEvent = "ELEV_TRIM_UP"
	//KeyElevUp Increments elevator up	 (Pilot only).
	KeyElevUp KeySimEvent = "ELEV_UP"
	//KeyRudderLeft Increments rudder left
	KeyRudderLeft KeySimEvent = "RUDDER_LEFT"
	//KeyRudderCenter Centers rudder position
	KeyRudderCenter KeySimEvent = "RUDDER_CENTER"
	//KeyRudderRight Increments rudder right
	KeyRudderRight KeySimEvent = "RUDDER_RIGHT"
	//KeyElevatorSet Sets elevator position (-16383 - +16383)
	KeyElevatorSet KeySimEvent = "ELEVATOR_SET"
	//KeyAileronSet Sets aileron position (-16383 - +16383)
	KeyAileronSet KeySimEvent = "AILERON_SET"
	//KeyRudderSet Sets rudder position (-16383 - +16383)
	KeyRudderSet KeySimEvent = "RUDDER_SET"
	//KeyFlapsIncr Increments flap handle position	All aircraft
	KeyFlapsIncr KeySimEvent = "FLAPS_INCR"
	//KeyFlapsDecr Decrements flap handle position	All aircraft
	KeyFlapsDecr KeySimEvent = "FLAPS_DECR"
	//KeyAxisElevatorSet Sets elevator position (-16383 - +16383)	 (Pilot only, and not transmitted to Co-pilot)
	KeyAxisElevatorSet KeySimEvent = "AXIS_ELEVATOR_SET"
	//KeyAxisAileronsSet Sets aileron position (-16383 - +16383)	 (Pilot only, and not transmitted to Co-pilot)
	KeyAxisAileronsSet KeySimEvent = "AXIS_AILERONS_SET"
	//KeyAxisRudderSet Sets rudder position (-16383 - +16383)	 (Pilot only, and not transmitted to Co-pilot)
	KeyAxisRudderSet KeySimEvent = "AXIS_RUDDER_SET"
	//KeyAxisElevTrimSet Sets elevator trim position (-16383 - +16383)
	KeyAxisElevTrimSet KeySimEvent = "AXIS_ELEV_TRIM_SET"
	//KeySpoilersSet Sets spoiler handle position (0 to 16383)	All aircraft
	KeySpoilersSet KeySimEvent = "SPOILERS_SET"
	//KeySpoilersArmToggle Toggles arming of auto-spoilers	All aircraft
	KeySpoilersArmToggle KeySimEvent = "SPOILERS_ARM_TOGGLE"
	//KeySpoilersOn Sets spoiler handle to full extend position	All aircraft
	KeySpoilersOn KeySimEvent = "SPOILERS_ON"
	//KeySpoilersOff Sets spoiler handle to full retract position	All aircraft
	KeySpoilersOff KeySimEvent = "SPOILERS_OFF"
	//KeySpoilersArmOn Sets auto-spoiler arming on	All aircraft
	KeySpoilersArmOn KeySimEvent = "SPOILERS_ARM_ON"
	//KeySpoilersArmOff Sets auto-spoiler arming off	All aircraft
	KeySpoilersArmOff KeySimEvent = "SPOILERS_ARM_OFF"
	//KeySpoilersArmSet Sets auto-spoiler arming (0,1)	All aircraft
	KeySpoilersArmSet KeySimEvent = "SPOILERS_ARM_SET"
	//KeyAileronTrimLeft Increments aileron trim left
	KeyAileronTrimLeft KeySimEvent = "AILERON_TRIM_LEFT"
	//KeyAileronTrimRight Increments aileron trim right
	KeyAileronTrimRight KeySimEvent = "AILERON_TRIM_RIGHT"
	//KeyRudderTrimLeft Increments rudder trim left
	KeyRudderTrimLeft KeySimEvent = "RUDDER_TRIM_LEFT"
	//KeyRudderTrimRight Increments aileron trim right
	KeyRudderTrimRight KeySimEvent = "RUDDER_TRIM_RIGHT"
	//KeyAxisSpoilerSet Sets spoiler handle position (-16383 - +16383)	All aircraft
	KeyAxisSpoilerSet KeySimEvent = "AXIS_SPOILER_SET"
	//KeyFlapsSet Sets flap handle to closest increment (0 to 16383)	All aircraft
	KeyFlapsSet KeySimEvent = "FLAPS_SET"
	//KeyElevatorTrimSet Sets elevator trim position (0 to 16383)
	KeyElevatorTrimSet KeySimEvent = "ELEVATOR_TRIM_SET"
	//KeyAxisFlapsSet Sets flap handle to closest increment (-16383 - +16383)
	KeyAxisFlapsSet KeySimEvent = "AXIS_FLAPS_SET"
	//KeyApMaster Toggles AP on/off
	KeyApMaster KeySimEvent = "AP_MASTER"
	//KeyAutopilotOff Turns AP off
	KeyAutopilotOff KeySimEvent = "AUTOPILOT_OFF"
	//KeyAutopilotOn Turns AP on
	KeyAutopilotOn KeySimEvent = "AUTOPILOT_ON"
	//KeyYawDamperToggle Toggles yaw damper on/off
	KeyYawDamperToggle KeySimEvent = "YAW_DAMPER_TOGGLE"
	//KeyApPanelHeadingHold Toggles heading hold mode on/off
	KeyApPanelHeadingHold KeySimEvent = "AP_PANEL_HEADING_HOLD"
	//KeyApPanelAltitudeHold Toggles altitude hold mode on/off
	KeyApPanelAltitudeHold KeySimEvent = "AP_PANEL_ALTITUDE_HOLD"
	//KeyApAttHoldOn Turns on AP wing leveler and pitch hold mode
	KeyApAttHoldOn KeySimEvent = "AP_ATT_HOLD_ON"
	//KeyApLocHoldOn Turns AP localizer hold on/armed and glide-slope hold mode off
	KeyApLocHoldOn KeySimEvent = "AP_LOC_HOLD_ON"
	//KeyApAprHoldOn Turns both AP localizer and glide-slope modes on/armed
	KeyApAprHoldOn KeySimEvent = "AP_APR_HOLD_ON"
	//KeyApHdgHoldOn Turns heading hold mode on
	KeyApHdgHoldOn KeySimEvent = "AP_HDG_HOLD_ON"
	//KeyApAltHoldOn Turns altitude hold mode on
	KeyApAltHoldOn KeySimEvent = "AP_ALT_HOLD_ON"
	//KeyApWingLevelerOn Turns wing leveler mode on
	KeyApWingLevelerOn KeySimEvent = "AP_WING_LEVELER_ON"
	//KeyApBcHoldOn Turns localizer back course hold mode on/armed
	KeyApBcHoldOn KeySimEvent = "AP_BC_HOLD_ON"
	//KeyApNav1HoldOn Turns lateral hold mode on
	KeyApNav1HoldOn KeySimEvent = "AP_NAV1_HOLD_ON"
	//KeyApAttHoldOff Turns off attitude hold mode
	KeyApAttHoldOff KeySimEvent = "AP_ATT_HOLD_OFF"
	//KeyApLocHoldOff Turns off localizer hold mode
	KeyApLocHoldOff KeySimEvent = "AP_LOC_HOLD_OFF"
	//KeyApAprHoldOff Turns off approach hold mode
	KeyApAprHoldOff KeySimEvent = "AP_APR_HOLD_OFF"
	//KeyApHdgHoldOff Turns off heading hold mode
	KeyApHdgHoldOff KeySimEvent = "AP_HDG_HOLD_OFF"
	//KeyApAltHoldOff Turns off altitude hold mode
	KeyApAltHoldOff KeySimEvent = "AP_ALT_HOLD_OFF"
	//KeyApWingLevelerOff Turns off wing leveler mode
	KeyApWingLevelerOff KeySimEvent = "AP_WING_LEVELER_OFF"
	//KeyApBcHoldOff Turns off backcourse mode for localizer hold
	KeyApBcHoldOff KeySimEvent = "AP_BC_HOLD_OFF"
	//KeyApNav1HoldOff Turns off nav hold mode
	KeyApNav1HoldOff KeySimEvent = "AP_NAV1_HOLD_OFF"
	//KeyApAirspeedHold Toggles airspeed hold mode
	KeyApAirspeedHold KeySimEvent = "AP_AIRSPEED_HOLD"
	//KeyAutoThrottleArm Toggles autothrottle arming mode
	KeyAutoThrottleArm KeySimEvent = "AUTO_THROTTLE_ARM"
	//KeyAutoThrottleToGa Toggles Takeoff/Go Around mode
	KeyAutoThrottleToGa KeySimEvent = "AUTO_THROTTLE_TO_GA"
	//KeyHeadingBugInc Increments heading hold reference bug
	KeyHeadingBugInc KeySimEvent = "HEADING_BUG_INC"
	//KeyHeadingBugDec Decrements heading hold reference bug
	KeyHeadingBugDec KeySimEvent = "HEADING_BUG_DEC"
	//KeyHeadingBugSet Set heading hold reference bug (degrees)
	KeyHeadingBugSet KeySimEvent = "HEADING_BUG_SET"
	//KeyApPanelSpeedHold Toggles airspeed hold mode
	KeyApPanelSpeedHold KeySimEvent = "AP_PANEL_SPEED_HOLD"
	//KeyApAltVarInc Increments reference altitude
	KeyApAltVarInc KeySimEvent = "AP_ALT_VAR_INC"
	//KeyApAltVarDec Decrements reference altitude
	KeyApAltVarDec KeySimEvent = "AP_ALT_VAR_DEC"
	//KeyApVsVarInc Increments vertical speed reference
	KeyApVsVarInc KeySimEvent = "AP_VS_VAR_INC"
	//KeyApVsVarDec Decrements vertical speed reference
	KeyApVsVarDec KeySimEvent = "AP_VS_VAR_DEC"
	//KeyApSpdVarInc Increments airspeed hold reference
	KeyApSpdVarInc KeySimEvent = "AP_SPD_VAR_INC"
	//KeyApSpdVarDec Decrements airspeed hold reference
	KeyApSpdVarDec KeySimEvent = "AP_SPD_VAR_DEC"
	//KeyApPanelMachHold Toggles mach hold
	KeyApPanelMachHold KeySimEvent = "AP_PANEL_MACH_HOLD"
	//KeyApMachVarInc Increments reference mach
	KeyApMachVarInc KeySimEvent = "AP_MACH_VAR_INC"
	//KeyApMachVarDec Decrements reference mach
	KeyApMachVarDec KeySimEvent = "AP_MACH_VAR_DEC"
	//KeyApMachHold Toggles mach hold
	KeyApMachHold KeySimEvent = "AP_MACH_HOLD"
	//KeyApAltVarSetMetric Sets reference altitude in meters
	KeyApAltVarSetMetric KeySimEvent = "AP_ALT_VAR_SET_METRIC"
	//KeyApVsVarSetEnglish Sets reference vertical speed in feet per minute
	KeyApVsVarSetEnglish KeySimEvent = "AP_VS_VAR_SET_ENGLISH"
	//KeyApSpdVarSet Sets airspeed reference in knots
	KeyApSpdVarSet KeySimEvent = "AP_SPD_VAR_SET"
	//KeyApMachVarSet Sets mach reference
	KeyApMachVarSet KeySimEvent = "AP_MACH_VAR_SET"
	//KeyYawDamperOn Turns yaw damper on
	KeyYawDamperOn KeySimEvent = "YAW_DAMPER_ON"
	//KeyYawDamperOff Turns yaw damper off
	KeyYawDamperOff KeySimEvent = "YAW_DAMPER_OFF"
	//KeyYawDamperSet Sets yaw damper on/off (1,0)
	KeyYawDamperSet KeySimEvent = "YAW_DAMPER_SET"
	//KeyApAirspeedOn Turns airspeed hold on
	KeyApAirspeedOn KeySimEvent = "AP_AIRSPEED_ON"
	//KeyApAirspeedOff Turns airspeed hold off
	KeyApAirspeedOff KeySimEvent = "AP_AIRSPEED_OFF"
	//KeyApAirspeedSet Sets airspeed hold on/off (1,0)
	KeyApAirspeedSet KeySimEvent = "AP_AIRSPEED_SET"
	//KeyApMachOn Turns mach hold on
	KeyApMachOn KeySimEvent = "AP_MACH_ON"
	//KeyApMachOff Turns mach hold off
	KeyApMachOff KeySimEvent = "AP_MACH_OFF"
	//KeyApMachSet Sets mach hold on/off (1,0)
	KeyApMachSet KeySimEvent = "AP_MACH_SET"
	//KeyApPanelAltitudeOn Turns altitude hold mode on (without capturing current altitude)
	KeyApPanelAltitudeOn KeySimEvent = "AP_PANEL_ALTITUDE_ON"
	//KeyApPanelAltitudeOff Turns altitude hold mode off
	KeyApPanelAltitudeOff KeySimEvent = "AP_PANEL_ALTITUDE_OFF"
	//KeyApPanelAltitudeSet Sets altitude hold mode on/off (1,0)
	KeyApPanelAltitudeSet KeySimEvent = "AP_PANEL_ALTITUDE_SET"
	//KeyApPanelHeadingOn Turns heading mode on (without capturing current heading)
	KeyApPanelHeadingOn KeySimEvent = "AP_PANEL_HEADING_ON"
	//KeyApPanelHeadingOff Turns heading mode off
	KeyApPanelHeadingOff KeySimEvent = "AP_PANEL_HEADING_OFF"
	//KeyApPanelHeadingSet Set heading mode on/off (1,0)
	KeyApPanelHeadingSet KeySimEvent = "AP_PANEL_HEADING_SET"
	//KeyApPanelMachOn Turns on mach hold
	KeyApPanelMachOn KeySimEvent = "AP_PANEL_MACH_ON"
	//KeyApPanelMachOff Turns off mach hold
	KeyApPanelMachOff KeySimEvent = "AP_PANEL_MACH_OFF"
	//KeyApPanelMachSet Sets mach hold on/off (1,0)
	KeyApPanelMachSet KeySimEvent = "AP_PANEL_MACH_SET"
	//KeyApPanelSpeedOn Turns on speed hold mode
	KeyApPanelSpeedOn KeySimEvent = "AP_PANEL_SPEED_ON"
	//KeyApPanelSpeedOff Turns off speed hold mode
	KeyApPanelSpeedOff KeySimEvent = "AP_PANEL_SPEED_OFF"
	//KeyApPanelSpeedSet Set speed hold mode on/off (1,0)
	KeyApPanelSpeedSet KeySimEvent = "AP_PANEL_SPEED_SET"
	//KeyApAltVarSetEnglish Sets altitude reference in feet
	KeyApAltVarSetEnglish KeySimEvent = "AP_ALT_VAR_SET_ENGLISH"
	//KeyApVsVarSetMetric Sets vertical speed reference in meters per minute
	KeyApVsVarSetMetric KeySimEvent = "AP_VS_VAR_SET_METRIC"
	//KeyToggleFlightDirector Toggles flight director on/off
	KeyToggleFlightDirector KeySimEvent = "TOGGLE_FLIGHT_DIRECTOR"
	//KeySyncFlightDirectorPitch Synchronizes flight director pitch with current aircraft pitch
	KeySyncFlightDirectorPitch KeySimEvent = "SYNC_FLIGHT_DIRECTOR_PITCH"
	//KeyIncAutobrakeControl Increments autobrake level
	KeyIncAutobrakeControl KeySimEvent = "INCREASE_AUTOBRAKE_CONTROL"
	//KeyDecAutobrakeControl Decrements autobrake level
	KeyDecAutobrakeControl KeySimEvent = "DECREASE_AUTOBRAKE_CONTROL"
	//KeyAutopilotAirspeedHoldCurrent Turns airspeed hold mode on with current airspeed
	KeyAutopilotAirspeedHoldCurrent KeySimEvent = "AP_PANEL_SPEED_HOLD_TOGGLE"
	//KeyAutopilotMachHoldCurrent Sets mach hold reference to current mach
	KeyAutopilotMachHoldCurrent KeySimEvent = "AP_PANEL_MACH_HOLD_TOGGLE"
	//KeyApNavSelectSet Sets the nav (1 or 2) which is used by the Nav hold modes
	KeyApNavSelectSet KeySimEvent = "AP_NAV_SELECT_SET"
	//KeyHeadingBugSelect Selects the heading bug for use with +/-
	KeyHeadingBugSelect KeySimEvent = "HEADING_BUG_SELECT"
	//KeyAltitudeBugSelect Selects the altitude reference for use with +/-
	KeyAltitudeBugSelect KeySimEvent = "ALTITUDE_BUG_SELECT"
	//KeyVsiBugSelect Selects the vertical speed reference for use with +/-
	KeyVsiBugSelect KeySimEvent = "VSI_BUG_SELECT"
	//KeyAirspeedBugSelect Selects the airspeed reference for use with +/-
	KeyAirspeedBugSelect KeySimEvent = "AIRSPEED_BUG_SELECT"
	//KeyApPitchRefIncUp Increments the pitch reference for pitch hold mode
	KeyApPitchRefIncUp KeySimEvent = "AP_PITCH_REF_INC_UP"
	//KeyApPitchRefIncDn Decrements the pitch reference for pitch hold mode
	KeyApPitchRefIncDn KeySimEvent = "AP_PITCH_REF_INC_DN"
	//KeyApPitchRefSelect Selects pitch reference for use with +/-
	KeyApPitchRefSelect KeySimEvent = "AP_PITCH_REF_SELECT"
	//KeyApAttHold Toggle attitude hold mode
	KeyApAttHold KeySimEvent = "AP_ATT_HOLD"
	//KeyApLocHold Toggles localizer (only) hold mode
	KeyApLocHold KeySimEvent = "AP_LOC_HOLD"
	//KeyApAprHold Toggles approach hold (localizer and glide-slope)
	KeyApAprHold KeySimEvent = "AP_APR_HOLD"
	//KeyApHdgHold Toggles heading hold mode
	KeyApHdgHold KeySimEvent = "AP_HDG_HOLD"
	//KeyApAltHold Toggles altitude hold mode
	KeyApAltHold KeySimEvent = "AP_ALT_HOLD"
	//KeyApWingLeveler Toggles wing leveler mode
	KeyApWingLeveler KeySimEvent = "AP_WING_LEVELER"
	//KeyApBcHold Toggles the backcourse mode for the localizer hold
	KeyApBcHold KeySimEvent = "AP_BC_HOLD"
	//KeyApNav1Hold Toggles the nav hold mode
	KeyApNav1Hold KeySimEvent = "AP_NAV1_HOLD"
	//KeyFuelSelectorOff Turns selector 1 to OFF position
	KeyFuelSelectorOff KeySimEvent = "FUEL_SELECTOR_OFF"
	//KeyFuelSelectorAll Turns selector 1 to ALL position
	KeyFuelSelectorAll KeySimEvent = "FUEL_SELECTOR_ALL"
	//KeyFuelSelectorLeft Turns selector 1 to LEFT position (burns from tip then aux then main)
	KeyFuelSelectorLeft KeySimEvent = "FUEL_SELECTOR_LEFT"
	//KeyFuelSelectorRight Turns selector 1 to RIGHT position (burns from tip then aux then main)
	KeyFuelSelectorRight KeySimEvent = "FUEL_SELECTOR_RIGHT"
	//KeyFuelSelectorLeftAux Turns selector 1 to LEFT AUX position
	KeyFuelSelectorLeftAux KeySimEvent = "FUEL_SELECTOR_LEFT_AUX"
	//KeyFuelSelectorRightAux Turns selector 1 to RIGHT AUX position
	KeyFuelSelectorRightAux KeySimEvent = "FUEL_SELECTOR_RIGHT_AUX"
	//KeyFuelSelectorCenter Turns selector 1 to CENTER position
	KeyFuelSelectorCenter KeySimEvent = "FUEL_SELECTOR_CENTER"
	//KeyFuelSelectorSet Sets selector 1 position (see code list below)
	KeyFuelSelectorSet KeySimEvent = "FUEL_SELECTOR_SET"
	//KeyFuelSelector2Off Turns selector 2 to OFF position
	KeyFuelSelector2Off KeySimEvent = "FUEL_SELECTOR_2_OFF"
	//KeyFuelSelector2All Turns selector 2 to ALL position
	KeyFuelSelector2All KeySimEvent = "FUEL_SELECTOR_2_ALL"
	//KeyFuelSelector2Left Turns selector 2 to LEFT position (burns from tip then aux then main)
	KeyFuelSelector2Left KeySimEvent = "FUEL_SELECTOR_2_LEFT"
	//KeyFuelSelector2Right Turns selector 2 to RIGHT position (burns from tip then aux then main)
	KeyFuelSelector2Right KeySimEvent = "FUEL_SELECTOR_2_RIGHT"
	//KeyFuelSelector2LeftAux Turns selector 2 to LEFT AUX position
	KeyFuelSelector2LeftAux KeySimEvent = "FUEL_SELECTOR_2_LEFT_AUX"
	//KeyFuelSelector2RightAux Turns selector 2 to RIGHT AUX position
	KeyFuelSelector2RightAux KeySimEvent = "FUEL_SELECTOR_2_RIGHT_AUX"
	//KeyFuelSelector2Center Turns selector 2 to CENTER position
	KeyFuelSelector2Center KeySimEvent = "FUEL_SELECTOR_2_CENTER"
	//KeyFuelSelector2Set Sets selector 2 position (see code list below)
	KeyFuelSelector2Set KeySimEvent = "FUEL_SELECTOR_2_SET"
	//KeyFuelSelector3Off Turns selector 3 to OFF position
	KeyFuelSelector3Off KeySimEvent = "FUEL_SELECTOR_3_OFF"
	//KeyFuelSelector3All Turns selector 3 to ALL position
	KeyFuelSelector3All KeySimEvent = "FUEL_SELECTOR_3_ALL"
	//KeyFuelSelector3Left Turns selector 3 to LEFT position (burns from tip then aux then main)
	KeyFuelSelector3Left KeySimEvent = "FUEL_SELECTOR_3_LEFT"
	//KeyFuelSelector3Right Turns selector 3 to RIGHT position (burns from tip then aux then main)
	KeyFuelSelector3Right KeySimEvent = "FUEL_SELECTOR_3_RIGHT"
	//KeyFuelSelector3LeftAux Turns selector 3 to LEFT AUX position
	KeyFuelSelector3LeftAux KeySimEvent = "FUEL_SELECTOR_3_LEFT_AUX"
	//KeyFuelSelector3RightAux Turns selector 3 to RIGHT AUX position
	KeyFuelSelector3RightAux KeySimEvent = "FUEL_SELECTOR_3_RIGHT_AUX"
	//KeyFuelSelector3Center Turns selector 3 to CENTER position
	KeyFuelSelector3Center KeySimEvent = "FUEL_SELECTOR_3_CENTER"
	//KeyFuelSelector3Set Sets selector 3 position (see code list below)
	KeyFuelSelector3Set KeySimEvent = "FUEL_SELECTOR_3_SET"
	//KeyFuelSelector4Off Turns selector 4 to OFF position
	KeyFuelSelector4Off KeySimEvent = "FUEL_SELECTOR_4_OFF"
	//KeyFuelSelector4All Turns selector 4 to ALL position
	KeyFuelSelector4All KeySimEvent = "FUEL_SELECTOR_4_ALL"
	//KeyFuelSelector4Left Turns selector 4 to LEFT position (burns from tip then aux then main)
	KeyFuelSelector4Left KeySimEvent = "FUEL_SELECTOR_4_LEFT"
	//KeyFuelSelector4Right Turns selector 4 to RIGHT position (burns from tip then aux then main)
	KeyFuelSelector4Right KeySimEvent = "FUEL_SELECTOR_4_RIGHT"
	//KeyFuelSelector4LeftAux Turns selector 4 to LEFT AUX position
	KeyFuelSelector4LeftAux KeySimEvent = "FUEL_SELECTOR_4_LEFT_AUX"
	//KeyFuelSelector4RightAux Turns selector 4 to RIGHT AUX position
	KeyFuelSelector4RightAux KeySimEvent = "FUEL_SELECTOR_4_RIGHT_AUX"
	//KeyFuelSelector4Center Turns selector 4 to CENTER position
	KeyFuelSelector4Center KeySimEvent = "FUEL_SELECTOR_4_CENTER"
	//KeyFuelSelector4Set Sets selector 4 position (see code list below)
	KeyFuelSelector4Set KeySimEvent = "FUEL_SELECTOR_4_SET"
	//KeyCrossFeedOpen Opens cross feed valve (when used in conjunction with "isolate" tank)
	KeyCrossFeedOpen KeySimEvent = "CROSS_FEED_OPEN"
	//KeyCrossFeedToggle Toggles crossfeed valve (when used in conjunction with "isolate" tank)
	KeyCrossFeedToggle KeySimEvent = "CROSS_FEED_TOGGLE"
	//KeyCrossFeedOff Closes crossfeed valve (when used in conjunction with "isolate" tank)
	KeyCrossFeedOff KeySimEvent = "CROSS_FEED_OFF"
	//KeyXpndr Sequentially selects the transponder digits for use with +/-.
	KeyXpndr KeySimEvent = "XPNDR"
	//KeyAdf Sequentially selects the ADF tuner digits for use with +/-. Follow by KEY_SELECT_2 for ADF 2.
	KeyAdf KeySimEvent = "ADF"
	//KeyDme Selects the DME for use with +/-
	KeyDme KeySimEvent = "DME"
	//KeyComRadio Sequentially selects the COM tuner digits for use with +/-. Follow by KEY_SELECT_2 for COM 2.	All aircraft
	KeyComRadio KeySimEvent = "COM_RADIO"
	//KeyVorObs Sequentially selects the VOR OBS for use with +/-. Follow by KEY_SELECT_2 for VOR 2.
	KeyVorObs KeySimEvent = "VOR_OBS"
	//KeyNavRadio Sequentially selects the NAV tuner digits for use with +/-. Follow by KEY_SELECT_2 for NAV 2.
	KeyNavRadio KeySimEvent = "NAV_RADIO"
	//KeyComRadioWholeDec Decrements COM by one MHz	All aircraft
	KeyComRadioWholeDec KeySimEvent = "COM_RADIO_WHOLE_DEC"
	//KeyComRadioWholeInc Increments COM by one MHz	All aircraft
	KeyComRadioWholeInc KeySimEvent = "COM_RADIO_WHOLE_INC"
	//KeyComRadioFractDec Decrements COM by 25 KHz	All aircraft
	KeyComRadioFractDec KeySimEvent = "COM_RADIO_FRACT_DEC"
	//KeyComRadioFractInc Increments COM by 25 KHz	All aircraft
	KeyComRadioFractInc KeySimEvent = "COM_RADIO_FRACT_INC"
	//KeyNav1RadioWholeDec Decrements Nav 1 by one MHz
	KeyNav1RadioWholeDec KeySimEvent = "NAV1_RADIO_WHOLE_DEC"
	//KeyNav1RadioWholeInc Increments Nav 1 by one MHz
	KeyNav1RadioWholeInc KeySimEvent = "NAV1_RADIO_WHOLE_INC"
	//KeyNav1RadioFractDec Decrements Nav 1 by 25 KHz
	KeyNav1RadioFractDec KeySimEvent = "NAV1_RADIO_FRACT_DEC"
	//KeyNav1RadioFractInc Increments Nav 1 by 25 KHz
	KeyNav1RadioFractInc KeySimEvent = "NAV1_RADIO_FRACT_INC"
	//KeyNav2RadioWholeDec Decrements Nav 2 by one MHz
	KeyNav2RadioWholeDec KeySimEvent = "NAV2_RADIO_WHOLE_DEC"
	//KeyNav2RadioWholeInc Increments Nav 2 by one MHz
	KeyNav2RadioWholeInc KeySimEvent = "NAV2_RADIO_WHOLE_INC"
	//KeyNav2RadioFractDec Decrements Nav 2 by 25 KHz
	KeyNav2RadioFractDec KeySimEvent = "NAV2_RADIO_FRACT_DEC"
	//KeyNav2RadioFractInc Increments Nav 2 by 25 KHz
	KeyNav2RadioFractInc KeySimEvent = "NAV2_RADIO_FRACT_INC"
	//KeyAdf100Inc Increments ADF by 100 KHz
	KeyAdf100Inc KeySimEvent = "ADF_100_INC"
	//KeyAdf10Inc Increments ADF by 10 KHz
	KeyAdf10Inc KeySimEvent = "ADF_10_INC"
	//KeyAdf1Inc Increments ADF by 1 KHz
	KeyAdf1Inc KeySimEvent = "ADF_1_INC"
	//KeyXpndr1000Inc Increments first digit of transponder	All Aircraft
	KeyXpndr1000Inc KeySimEvent = "XPNDR_1000_INC"
	//KeyXpndr100Inc Increments second digit of transponder	All Aircraft
	KeyXpndr100Inc KeySimEvent = "XPNDR_100_INC"
	//KeyXpndr10Inc Increments third digit of transponder	All Aircraft
	KeyXpndr10Inc KeySimEvent = "XPNDR_10_INC"
	//KeyXpndr1Inc Increments fourth digit of transponder	All Aircraft
	KeyXpndr1Inc KeySimEvent = "XPNDR_1_INC"
	//KeyVor1ObiDec Decrements the VOR 1 OBS setting
	KeyVor1ObiDec KeySimEvent = "VOR1_OBI_DEC"
	//KeyVor1ObiInc Increments the VOR 1 OBS setting
	KeyVor1ObiInc KeySimEvent = "VOR1_OBI_INC"
	//KeyVor2ObiDec Decrements the VOR 2 OBS setting
	KeyVor2ObiDec KeySimEvent = "VOR2_OBI_DEC"
	//KeyVor2ObiInc Increments the VOR 2 OBS setting
	KeyVor2ObiInc KeySimEvent = "VOR2_OBI_INC"
	//KeyAdf100Dec Decrements ADF by 100 KHz
	KeyAdf100Dec KeySimEvent = "ADF_100_DEC"
	//KeyAdf10Dec Decrements ADF by 10 KHz
	KeyAdf10Dec KeySimEvent = "ADF_10_DEC"
	//KeyAdf1Dec Decrements ADF by 1 KHz
	KeyAdf1Dec KeySimEvent = "ADF_1_DEC"
	//KeyComRadioSet Sets COM frequency (BCD Hz)	All aircraft
	KeyComRadioSet KeySimEvent = "COM_RADIO_SET"
	//KeyNav1RadioSet Sets NAV 1 frequency (BCD Hz)
	KeyNav1RadioSet KeySimEvent = "NAV1_RADIO_SET"
	//KeyNav2RadioSet Sets NAV 2 frequency (BCD Hz)
	KeyNav2RadioSet KeySimEvent = "NAV2_RADIO_SET"
	//KeyAdfSet Sets ADF frequency (BCD Hz)
	KeyAdfSet KeySimEvent = "ADF_SET"
	//KeyXpndrSet Sets transponder code (BCD)	All aircraft
	KeyXpndrSet KeySimEvent = "XPNDR_SET"
	//KeyVor1Set Sets OBS 1 (0 to 360)
	KeyVor1Set KeySimEvent = "VOR1_SET"
	//KeyVor2Set Sets OBS 2 (0 to 360)
	KeyVor2Set KeySimEvent = "VOR2_SET"
	//KeyDme1Toggle Sets DME display to Nav 1
	KeyDme1Toggle KeySimEvent = "DME1_TOGGLE"
	//KeyDme2Toggle Sets DME display to Nav 2
	KeyDme2Toggle KeySimEvent = "DME2_TOGGLE"
	//KeyRadioVor1IdentDisable Turns NAV 1 ID off
	KeyRadioVor1IdentDisable KeySimEvent = "RADIO_VOR1_IDENT_DISABLE"
	//KeyRadioVor2IdentDisable Turns NAV 2 ID off
	KeyRadioVor2IdentDisable KeySimEvent = "RADIO_VOR2_IDENT_DISABLE"
	//KeyRadioDme1IdentDisable Turns DME 1 ID off
	KeyRadioDme1IdentDisable KeySimEvent = "RADIO_DME1_IDENT_DISABLE"
	//KeyRadioDme2IdentDisable Turns DME 2 ID off
	KeyRadioDme2IdentDisable KeySimEvent = "RADIO_DME2_IDENT_DISABLE"
	//KeyRadioAdfIdentDisable Turns ADF 1 ID off
	KeyRadioAdfIdentDisable KeySimEvent = "RADIO_ADF_IDENT_DISABLE"
	//KeyRadioVor1IdentEnable Turns NAV 1 ID on
	KeyRadioVor1IdentEnable KeySimEvent = "RADIO_VOR1_IDENT_ENABLE"
	//KeyRadioVor2IdentEnable Turns NAV 2 ID on
	KeyRadioVor2IdentEnable KeySimEvent = "RADIO_VOR2_IDENT_ENABLE"
	//KeyRadioDme1IdentEnable Turns DME 1 ID on
	KeyRadioDme1IdentEnable KeySimEvent = "RADIO_DME1_IDENT_ENABLE"
	//KeyRadioDme2IdentEnable Turns DME 2 ID on
	KeyRadioDme2IdentEnable KeySimEvent = "RADIO_DME2_IDENT_ENABLE"
	//KeyRadioAdfIdentEnable Turns ADF 1 ID on
	KeyRadioAdfIdentEnable KeySimEvent = "RADIO_ADF_IDENT_ENABLE"
	//KeyRadioVor1IdentToggle Toggles NAV 1 ID
	KeyRadioVor1IdentToggle KeySimEvent = "RADIO_VOR1_IDENT_TOGGLE"
	//KeyRadioVor2IdentToggle Toggles NAV 2 ID
	KeyRadioVor2IdentToggle KeySimEvent = "RADIO_VOR2_IDENT_TOGGLE"
	//KeyRadioDme1IdentToggle Toggles DME 1 ID
	KeyRadioDme1IdentToggle KeySimEvent = "RADIO_DME1_IDENT_TOGGLE"
	//KeyRadioDme2IdentToggle Toggles DME 2 ID
	KeyRadioDme2IdentToggle KeySimEvent = "RADIO_DME2_IDENT_TOGGLE"
	//KeyRadioAdfIdentToggle Toggles ADF 1 ID
	KeyRadioAdfIdentToggle KeySimEvent = "RADIO_ADF_IDENT_TOGGLE"
	//KeyRadioVor1IdentSet Sets NAV 1 ID (on/off)
	KeyRadioVor1IdentSet KeySimEvent = "RADIO_VOR1_IDENT_SET"
	//KeyRadioVor2IdentSet Sets NAV 2 ID (on/off)
	KeyRadioVor2IdentSet KeySimEvent = "RADIO_VOR2_IDENT_SET"
	//KeyRadioDme1IdentSet Sets DME 1 ID (on/off)
	KeyRadioDme1IdentSet KeySimEvent = "RADIO_DME1_IDENT_SET"
	//KeyRadioDme2IdentSet Sets DME 2 ID (on/off)
	KeyRadioDme2IdentSet KeySimEvent = "RADIO_DME2_IDENT_SET"
	//KeyRadioAdfIdentSet Sets ADF 1 ID (on/off)
	KeyRadioAdfIdentSet KeySimEvent = "RADIO_ADF_IDENT_SET"
	//KeyAdfCardInc Increments ADF card
	KeyAdfCardInc KeySimEvent = "ADF_CARD_INC"
	//KeyAdfCardDec Decrements ADF card
	KeyAdfCardDec KeySimEvent = "ADF_CARD_DEC"
	//KeyAdfCardSet Sets ADF card (0-360)
	KeyAdfCardSet KeySimEvent = "ADF_CARD_SET"
	//KeyDmeToggle Toggles between NAV 1 and NAV 2
	KeyDmeToggle KeySimEvent = "TOGGLE_DME"
	//KeyAvionicsMasterSet Sets the avionics master switch	All aircraft
	KeyAvionicsMasterSet KeySimEvent = "AVIONICS_MASTER_SET"
	//KeyToggleAvionicsMaster Toggles the avionics master switch	All aircraft
	KeyToggleAvionicsMaster KeySimEvent = "TOGGLE_AVIONICS_MASTER"
	//KeyComStbyRadioSet Sets COM 1 standby frequency (BCD Hz)	All aircraft
	KeyComStbyRadioSet      KeySimEvent = "COM_STBY_RADIO_SET"
	KeyComStbyRadioSwitchTo KeySimEvent = "COM_STBY_RADIO_SWAP"
	//KeyComRadioSwap Swaps COM 1 frequency with standby	All aircraft
	KeyComRadioSwap KeySimEvent = "COM_STBY_RADIO_SWAP"
	//KeyComRadioFractDecCarry Decrement COM 1 frequency by 25 KHz, and carry when digit wraps	All aircraft
	KeyComRadioFractDecCarry KeySimEvent = "COM_RADIO_FRACT_DEC_CARRY"
	//KeyComRadioFractIncCarry Increment COM 1 frequency by 25 KHz, and carry when digit wraps	All aircraft
	KeyComRadioFractIncCarry KeySimEvent = "COM_RADIO_FRACT_INC_CARRY"
	//KeyCom2RadioWholeDec Decrement COM 2 frequency by 1 MHz, with no carry when digit wraps	All aircraft
	KeyCom2RadioWholeDec KeySimEvent = "COM2_RADIO_WHOLE_DEC"
	//KeyCom2RadioWholeInc Increment COM 2 frequency by 1 MHz, with no carry when digit wraps	All aircraft
	KeyCom2RadioWholeInc KeySimEvent = "COM2_RADIO_WHOLE_INC"
	//KeyCom2RadioFractDec Decrement COM 2 frequency by 25 KHz, with no carry when digit wraps	All aircraft
	KeyCom2RadioFractDec KeySimEvent = "COM2_RADIO_FRACT_DEC"
	//KeyCom2RadioFractDecCarry Decrement COM 2 frequency by 25 KHz, and carry when digit wraps	All aircraft
	KeyCom2RadioFractDecCarry KeySimEvent = "COM2_RADIO_FRACT_DEC_CARRY"
	//KeyCom2RadioFractInc Increment COM 2 frequency by 25 KHz, with no carry when digit wraps	All aircraft
	KeyCom2RadioFractInc KeySimEvent = "COM2_RADIO_FRACT_INC"
	//KeyCom2RadioFractIncCarry Increment COM 2 frequency by 25 KHz, and carry when digit wraps	All aircraft
	KeyCom2RadioFractIncCarry KeySimEvent = "COM2_RADIO_FRACT_INC_CARRY"
	//KeyCom2RadioSet Sets COM 2 frequency (BCD Hz)	All aircraft
	KeyCom2RadioSet KeySimEvent = "COM2_RADIO_SET"
	//KeyCom2StbyRadioSet Sets COM 2 standby frequency (BCD Hz)	All aircraft
	KeyCom2StbyRadioSet KeySimEvent = "COM2_STBY_RADIO_SET"
	//KeyCom2RadioSwap Swaps COM 2 frequency with standby	All aircraft
	KeyCom2RadioSwap KeySimEvent = "COM2_RADIO_SWAP"
	//KeyNav1RadioFractDecCarry Decrement NAV 1 frequency by 50 KHz, and carry when digit wraps
	KeyNav1RadioFractDecCarry KeySimEvent = "NAV1_RADIO_FRACT_DEC_CARRY"
	//KeyNav1RadioFractIncCarry Increment NAV 1 frequency by 50 KHz, and carry when digit wraps
	KeyNav1RadioFractIncCarry KeySimEvent = "NAV1_RADIO_FRACT_INC_CARRY"
	//KeyNav1StbySet Sets NAV 1 standby frequency (BCD Hz)
	KeyNav1StbySet KeySimEvent = "NAV1_STBY_SET"
	//KeyNav1RadioSwap Swaps NAV 1 frequency with standby
	KeyNav1RadioSwap KeySimEvent = "NAV1_RADIO_SWAP"
	//KeyNav2RadioFractDecCarry Decrement NAV 2 frequency by 50 KHz, and carry when digit wraps
	KeyNav2RadioFractDecCarry KeySimEvent = "NAV2_RADIO_FRACT_DEC_CARRY"
	//KeyNav2RadioFractIncCarry Increment NAV 2 frequency by 50 KHz, and carry when digit wraps
	KeyNav2RadioFractIncCarry KeySimEvent = "NAV2_RADIO_FRACT_INC_CARRY"
	//KeyNav2StbySet Sets NAV 2 standby frequency (BCD Hz)
	KeyNav2StbySet KeySimEvent = "NAV2_STBY_SET"
	//KeyNav2RadioSwap Swaps NAV 2 frequency with standby
	KeyNav2RadioSwap KeySimEvent = "NAV2_RADIO_SWAP"
	//KeyAdf1RadioTenthsDec Decrements ADF 1 by 0.1 KHz.
	KeyAdf1RadioTenthsDec KeySimEvent = "ADF1_RADIO_TENTHS_DEC"
	//KeyAdf1RadioTenthsInc Increments ADF 1 by 0.1 KHz.
	KeyAdf1RadioTenthsInc KeySimEvent = "ADF1_RADIO_TENTHS_INC"
	//KeyXpndr1000Dec Decrements first digit of transponder	All Aircraft
	KeyXpndr1000Dec KeySimEvent = "XPNDR_1000_DEC"
	//KeyXpndr100Dec Decrements second digit of transponder	All Aircraft
	KeyXpndr100Dec KeySimEvent = "XPNDR_100_DEC"
	//KeyXpndr10Dec Decrements third digit of transponder	All Aircraft
	KeyXpndr10Dec KeySimEvent = "XPNDR_10_DEC"
	//KeyXpndr1Dec Decrements fourth digit of transponder	All Aircraft
	KeyXpndr1Dec KeySimEvent = "XPNDR_1_DEC"
	//KeyXpndrDecCarry Decrements fourth digit of transponder, and with carry.	All Aircraft
	KeyXpndrDecCarry KeySimEvent = "XPNDR_DEC_CARRY"
	//KeyXpndrIncCarry Increments fourth digit of transponder, and with carry.	All Aircraft
	KeyXpndrIncCarry KeySimEvent = "XPNDR_INC_CARRY"
	//KeyAdfFractDecCarry Decrements ADF 1 frequency by 0.1 KHz, with carry
	KeyAdfFractDecCarry KeySimEvent = "ADF_FRACT_DEC_CARRY"
	//KeyAdfFractIncCarry Increments ADF 1 frequency by 0.1 KHz, with carry
	KeyAdfFractIncCarry KeySimEvent = "ADF_FRACT_INC_CARRY"
	//KeyCom1TransmitSelect Selects COM 1 to transmit	All aircraft
	KeyCom1TransmitSelect KeySimEvent = "COM1_TRANSMIT_SELECT"
	//KeyCom2TransmitSelect Selects COM 2 to transmit	All aircraft
	KeyCom2TransmitSelect KeySimEvent = "COM2_TRANSMIT_SELECT"
	//KeyComReceiveAllToggle Toggles all COM radios to receive on	All aircraft
	KeyComReceiveAllToggle KeySimEvent = "COM_RECEIVE_ALL_TOGGLE"
	//KeyComReceiveAllSet Sets whether to receive on all COM radios (1,0)	All aircraft
	KeyComReceiveAllSet KeySimEvent = "COM_RECEIVE_ALL_SET"
	//KeyMarkerSoundToggle Toggles marker beacon sound on/off
	KeyMarkerSoundToggle KeySimEvent = "MARKER_SOUND_TOGGLE"
	//KeyAdfCompleteSet Sets ADF 1 frequency - standby if configured, otherwise primary (BCD Hz)
	KeyAdfCompleteSet KeySimEvent = "ADF_COMPLETE_SET"
	//KeyAdfWholeInc Increments ADF 1 by 1 KHz, with carry as digits wrap.
	KeyAdfWholeInc KeySimEvent = "ADF1_WHOLE_INC"
	//KeyAdfWholeDec Decrements ADF 1 by 1 KHz, with carry as digits wrap.
	KeyAdfWholeDec KeySimEvent = "ADF1_WHOLE_DEC"
	//KeyAdf2100Inc Increments the ADF 2 frequency 100 digit, with wrapping
	KeyAdf2100Inc KeySimEvent = "ADF2_100_INC"
	//KeyAdf210Inc Increments the ADF 2 frequency 10 digit, with wrapping
	KeyAdf210Inc KeySimEvent = "ADF2_10_INC"
	//KeyAdf21Inc Increments the ADF 2 frequency 1 digit, with wrapping
	KeyAdf21Inc KeySimEvent = "ADF2_1_INC"
	//KeyAdf2RadioTenthsInc Increments ADF 2 frequency 1/10 digit, with wrapping
	KeyAdf2RadioTenthsInc KeySimEvent = "ADF2_RADIO_TENTHS_INC"
	//KeyAdf2100Dec Decrements the ADF 2 frequency 100 digit, with wrapping
	KeyAdf2100Dec KeySimEvent = "ADF2_100_DEC"
	//KeyAdf210Dec Decrements the ADF 2 frequency 10 digit, with wrapping
	KeyAdf210Dec KeySimEvent = "ADF2_10_DEC"
	//KeyAdf21Dec Decrements the ADF 2 frequency 1 digit, with wrapping
	KeyAdf21Dec KeySimEvent = "ADF2_1_DEC"
	//KeyAdf2RadioTenthsDec Decrements ADF 2 frequency 1/10 digit, with wrapping
	KeyAdf2RadioTenthsDec KeySimEvent = "ADF2_RADIO_TENTHS_DEC"
	//KeyAdf2WholeInc Increments ADF 2 by 1 KHz, with carry as digits wrap.
	KeyAdf2WholeInc KeySimEvent = "ADF2_WHOLE_INC"
	//KeyAdf2WholeDec Decrements ADF 2 by 1 KHz, with carry as digits wrap.
	KeyAdf2WholeDec KeySimEvent = "ADF2_WHOLE_DEC"
	//KeyAdf2FractIncCarry Decrements ADF 2 frequency by 0.1 KHz, with carry
	KeyAdf2FractIncCarry KeySimEvent = "ADF2_FRACT_DEC_CARRY"
	//KeyAdf2FractDecCarry Increments ADF 2 frequency by 0.1 KHz, with carry
	KeyAdf2FractDecCarry KeySimEvent = "ADF2_FRACT_INC_CARRY"
	//KeyAdf2CompleteSet Sets ADF 2 frequency - standby if configured, otherwise primary (BCD Hz)
	KeyAdf2CompleteSet KeySimEvent = "ADF2_COMPLETE_SET"
	//KeyRadioAdf2IdentDisable Turns ADF 2 ID off
	KeyRadioAdf2IdentDisable KeySimEvent = "RADIO_ADF2_IDENT_DISABLE"
	//KeyRadioAdf2IdentEnable Turns ADF 2 ID on
	KeyRadioAdf2IdentEnable KeySimEvent = "RADIO_ADF2_IDENT_ENABLE"
	//KeyRadioAdf2IdentToggle Toggles ADF 2 ID
	KeyRadioAdf2IdentToggle KeySimEvent = "RADIO_ADF2_IDENT_TOGGLE"
	//KeyRadioAdf2IdentSet Sets ADF 2 ID on/off (1,0)
	KeyRadioAdf2IdentSet KeySimEvent = "RADIO_ADF2_IDENT_SET"
	//KeyFrequencySwap Swaps frequency with standby on whichever NAV or COM radio is selected.
	KeyFrequencySwap KeySimEvent = "FREQUENCY_SWAP"
	//KeyToggleGpsDrivesNav1 Toggles between GPS and NAV 1 driving NAV 1 OBS display (and AP)
	KeyToggleGpsDrivesNav1 KeySimEvent = "TOGGLE_GPS_DRIVES_NAV1"
	//KeyGpsPowerButton Toggles power button
	KeyGpsPowerButton KeySimEvent = "GPS_POWER_BUTTON"
	//KeyGpsNearestButton Selects Nearest Airport Page
	KeyGpsNearestButton KeySimEvent = "GPS_NEAREST_BUTTON"
	//KeyGpsObsButton Toggles automatic sequencing of waypoints
	KeyGpsObsButton KeySimEvent = "GPS_OBS_BUTTON"
	//KeyGpsMsgButton Toggles the Message Page
	KeyGpsMsgButton KeySimEvent = "GPS_MSG_BUTTON"
	//KeyGpsMsgButtonDown Triggers the pressing of the message button.
	KeyGpsMsgButtonDown KeySimEvent = "GPS_MSG_BUTTON_DOWN"
	//KeyGpsMsgButtonUp Triggers the release of the message button
	KeyGpsMsgButtonUp KeySimEvent = "GPS_MSG_BUTTON_UP"
	//KeyGpsFlightplanButton Displays the programmed flightplan.
	KeyGpsFlightplanButton KeySimEvent = "GPS_FLIGHTPLAN_BUTTON"
	//KeyGpsTerrainButton Displays terrain information on default display
	KeyGpsTerrainButton KeySimEvent = "GPS_TERRAIN_BUTTON"
	//KeyGpsProcedureButton Displays the approach procedure page.
	KeyGpsProcedureButton KeySimEvent = "GPS_PROCEDURE_BUTTON"
	//KeyGpsZoominButton Zooms in default display
	KeyGpsZoominButton KeySimEvent = "GPS_ZOOMIN_BUTTON"
	//KeyGpsZoomoutButton Zooms out default display
	KeyGpsZoomoutButton KeySimEvent = "GPS_ZOOMOUT_BUTTON"
	//KeyGpsDirecttoButton Brings up the "Direct To" page
	KeyGpsDirecttoButton KeySimEvent = "GPS_DIRECTTO_BUTTON"
	//KeyGpsMenuButton Brings up page to select active legs in a flightplan.
	KeyGpsMenuButton KeySimEvent = "GPS_MENU_BUTTON"
	//KeyGpsClearButton Clears entered data on a page
	KeyGpsClearButton KeySimEvent = "GPS_CLEAR_BUTTON"
	//KeyGpsClearAllButton Clears all data immediately
	KeyGpsClearAllButton KeySimEvent = "GPS_CLEAR_ALL_BUTTON"
	//KeyGpsClearButtonDown Triggers the pressing of the Clear button
	KeyGpsClearButtonDown KeySimEvent = "GPS_CLEAR_BUTTON_DOWN"
	//KeyGpsClearButtonUp Triggers the release of the Clear button.
	KeyGpsClearButtonUp KeySimEvent = "GPS_CLEAR_BUTTON_UP"
	//KeyGpsEnterButton Approves entered data.
	KeyGpsEnterButton KeySimEvent = "GPS_ENTER_BUTTON"
	//KeyGpsCursorButton Selects GPS cursor
	KeyGpsCursorButton KeySimEvent = "GPS_CURSOR_BUTTON"
	//KeyGpsGroupKnobInc Increments cursor
	KeyGpsGroupKnobInc KeySimEvent = "GPS_GROUP_KNOB_INC"
	//KeyGpsGroupKnobDec Decrements cursor
	KeyGpsGroupKnobDec KeySimEvent = "GPS_GROUP_KNOB_DEC"
	//KeyGpsPageKnobInc Increments through pages
	KeyGpsPageKnobInc KeySimEvent = "GPS_PAGE_KNOB_INC"
	//KeyGpsPageKnobDec Decrements through pages
	KeyGpsPageKnobDec KeySimEvent = "GPS_PAGE_KNOB_DEC"
	//KeyEgt Selects EGT bug for +/-
	KeyEgt KeySimEvent = "EGT"
	//KeyEgtInc Increments EGT bugs
	KeyEgtInc KeySimEvent = "EGT_INC"
	//KeyEgtDec Decrements EGT bugs
	KeyEgtDec KeySimEvent = "EGT_DEC"
	//KeyEgtSet Sets EGT bugs (0 to 32767)
	KeyEgtSet KeySimEvent = "EGT_SET"
	//KeyBarometric Syncs altimeter setting to sea level pressure, or 29.92 if above 18000 feet
	KeyBarometric KeySimEvent = "BAROMETRIC"
	//KeyGyroDriftInc Increments heading indicator
	KeyGyroDriftInc KeySimEvent = "GYRO_DRIFT_INC"
	//KeyGyroDriftDec Decrements heading indicator
	KeyGyroDriftDec KeySimEvent = "GYRO_DRIFT_DEC"
	//KeyKohlsmanInc Increments altimeter setting
	KeyKohlsmanInc KeySimEvent = "KOHLSMAN_INC"
	//KeyKohlsmanDec Decrements altimeter setting
	KeyKohlsmanDec KeySimEvent = "KOHLSMAN_DEC"
	//KeyKohlsmanSet Sets altimeter setting (Millibars * 16)
	KeyKohlsmanSet KeySimEvent = "KOHLSMAN_SET"
	//KeyTrueAirspeedCalibrateInc Increments airspeed indicators true airspeed reference card
	KeyTrueAirspeedCalibrateInc KeySimEvent = "TRUE_AIRSPEED_CAL_INC"
	//KeyTrueAirspeedCalibrateDec Decrements airspeed indicators true airspeed reference card
	KeyTrueAirspeedCalibrateDec KeySimEvent = "TRUE_AIRSPEED_CAL_DEC"
	//KeyTrueAirspeedCalSet Sets airspeed indicators true airspeed reference card (degrees, where 0 is standard sea level conditions)
	KeyTrueAirspeedCalSet KeySimEvent = "TRUE_AIRSPEED_CAL_SET"
	//KeyEgt1Inc Increments EGT bug 1
	KeyEgt1Inc KeySimEvent = "EGT1_INC"
	//KeyEgt1Dec Decrements EGT bug 1
	KeyEgt1Dec KeySimEvent = "EGT1_DEC"
	//KeyEgt1Set Sets EGT bug 1 (0 to 32767)
	KeyEgt1Set KeySimEvent = "EGT1_SET"
	//KeyEgt2Inc Increments EGT bug 2
	KeyEgt2Inc KeySimEvent = "EGT2_INC"
	//KeyEgt2Dec Decrements EGT bug 2
	KeyEgt2Dec KeySimEvent = "EGT2_DEC"
	//KeyEgt2Set Sets EGT bug 2 (0 to 32767)
	KeyEgt2Set KeySimEvent = "EGT2_SET"
	//KeyEgt3Inc Increments EGT bug 3
	KeyEgt3Inc KeySimEvent = "EGT3_INC"
	//KeyEgt3Dec Decrements EGT bug 3
	KeyEgt3Dec KeySimEvent = "EGT3_DEC"
	//KeyEgt3Set Sets EGT bug 3 (0 to 32767)
	KeyEgt3Set KeySimEvent = "EGT3_SET"
	//KeyEgt4Inc Increments EGT bug 4
	KeyEgt4Inc KeySimEvent = "EGT4_INC"
	//KeyEgt4Dec Decrements EGT bug 4
	KeyEgt4Dec KeySimEvent = "EGT4_DEC"
	//KeyEgt4Set Sets EGT bug 4 (0 to 32767)
	KeyEgt4Set KeySimEvent = "EGT4_SET"
	//KeyAttitudeBarsPositionInc Increments attitude indicator pitch reference bars
	KeyAttitudeBarsPositionInc KeySimEvent = "ATTITUDE_BARS_POSITION_UP"
	//KeyAttitudeBarsPositionDec Decrements attitude indicator pitch reference bars
	KeyAttitudeBarsPositionDec KeySimEvent = "ATTITUDE_BARS_POSITION_DOWN"
	//KeyToggleAttitudeCage Cages attitude indicator at 0 pitch and bank
	KeyToggleAttitudeCage KeySimEvent = "ATTITUDE_CAGE_BUTTON"
	//KeyResetGForceIndicator Resets max/min indicated G force to 1.0.
	KeyResetGForceIndicator KeySimEvent = "RESET_G_FORCE_INDICATOR"
	//KeyResetMaxRpmIndicator Reset max indicated engine rpm to 0.
	KeyResetMaxRpmIndicator KeySimEvent = "RESET_MAX_RPM_INDICATOR"
	//KeyHeadingGyroSet Sets heading indicator to 0 drift error.
	KeyHeadingGyroSet KeySimEvent = "HEADING_GYRO_SET"
	//KeyGyroDriftSet Sets heading indicator drift angle (degrees).
	KeyGyroDriftSet KeySimEvent = "GYRO_DRIFT_SET"
	//KeyStrobesToggle Toggle strobe lights 	All aircraft
	KeyStrobesToggle KeySimEvent = "STROBES_TOGGLE"
	//KeyAllLightsToggle Toggle all lights
	KeyAllLightsToggle KeySimEvent = "ALL_LIGHTS_TOGGLE"
	//KeyPanelLightsToggle Toggle panel lights	All aircraft
	KeyPanelLightsToggle KeySimEvent = "PANEL_LIGHTS_TOGGLE"
	//KeyLandingLightsToggle Toggle landing lights	All aircraft
	KeyLandingLightsToggle KeySimEvent = "LANDING_LIGHTS_TOGGLE"
	//KeyLandingLightUp Rotate landing light up
	KeyLandingLightUp KeySimEvent = "LANDING_LIGHT_UP"
	//KeyLandingLightDown Rotate landing light down
	KeyLandingLightDown KeySimEvent = "LANDING_LIGHT_DOWN"
	//KeyLandingLightLeft Rotate landing light left
	KeyLandingLightLeft KeySimEvent = "LANDING_LIGHT_LEFT"
	//KeyLandingLightRight Rotate landing light right
	KeyLandingLightRight KeySimEvent = "LANDING_LIGHT_RIGHT"
	//KeyLandingLightHome Return landing light to default position
	KeyLandingLightHome KeySimEvent = "LANDING_LIGHT_HOME"
	//KeyStrobesOn Turn strobe lights on	All aircraft
	KeyStrobesOn KeySimEvent = "STROBES_ON"
	//KeyStrobesOff Turn strobe light off	All aircraft
	KeyStrobesOff KeySimEvent = "STROBES_OFF"
	//KeyStrobesSet Set strobe lights on/off (1,0)	All aircraft
	KeyStrobesSet KeySimEvent = "STROBES_SET"
	//KeyPanelLightsOn Turn panel lights on	All aircraft
	KeyPanelLightsOn KeySimEvent = "PANEL_LIGHTS_ON"
	//KeyPanelLightsOff Turn panel lights off	All aircraft
	KeyPanelLightsOff KeySimEvent = "PANEL_LIGHTS_OFF"
	//KeyPanelLightsSet Set panel lights on/off (1,0)	All aircraft
	KeyPanelLightsSet KeySimEvent = "PANEL_LIGHTS_SET"
	//KeyLandingLightsOn Turn landing lights on	All aircraft
	KeyLandingLightsOn KeySimEvent = "LANDING_LIGHTS_ON"
	//KeyLandingLightsOff Turn landing lights off	All aircraft
	KeyLandingLightsOff KeySimEvent = "LANDING_LIGHTS_OFF"
	//KeyLandingLightsSet Set landing lights on/off (1,0)	All aircraft
	KeyLandingLightsSet KeySimEvent = "LANDING_LIGHTS_SET"
	//KeyToggleBeaconLights Toggle beacon lights	All aircraft
	KeyToggleBeaconLights KeySimEvent = "TOGGLE_BEACON_LIGHTS"
	//KeyToggleTaxiLights Toggle taxi lights	All aircraft
	KeyToggleTaxiLights KeySimEvent = "TOGGLE_TAXI_LIGHTS"
	//KeyToggleLogoLights Toggle logo lights	All aircraft
	KeyToggleLogoLights KeySimEvent = "TOGGLE_LOGO_LIGHTS"
	//KeyToggleRecognitionLights Toggle recognition lights	All aircraft
	KeyToggleRecognitionLights KeySimEvent = "TOGGLE_RECOGNITION_LIGHTS"
	//KeyToggleWingLights Toggle wing lights	All aircraft
	KeyToggleWingLights KeySimEvent = "TOGGLE_WING_LIGHTS"
	//KeyToggleNavLights Toggle navigation lights	All aircraft
	KeyToggleNavLights KeySimEvent = "TOGGLE_NAV_LIGHTS"
	//KeyToggleCabinLights Toggle cockpit/cabin lights	All aircraft
	KeyToggleCabinLights KeySimEvent = "TOGGLE_CABIN_LIGHTS"
	//KeyToggleVacuumFailure Toggle vacuum system failure
	KeyToggleVacuumFailure KeySimEvent = "TOGGLE_VACUUM_FAILURE"
	//KeyToggleElectricalFailure Toggle electrical system failure
	KeyToggleElectricalFailure KeySimEvent = "TOGGLE_ELECTRICAL_FAILURE"
	//KeyTogglePitotBlockage Toggles blocked pitot tube
	KeyTogglePitotBlockage KeySimEvent = "TOGGLE_PITOT_BLOCKAGE"
	//KeyToggleStaticPortBlockage  Toggles blocked static port
	KeyToggleStaticPortBlockage KeySimEvent = "TOGGLE_STATIC_PORT_BLOCKAGE"
	//KeyToggleHydraulicFailure Toggles hydraulic system failure
	KeyToggleHydraulicFailure KeySimEvent = "TOGGLE_HYDRAULIC_FAILURE"
	//KeyToggleTotalBrakeFailure Toggles brake failure (both)
	KeyToggleTotalBrakeFailure KeySimEvent = "TOGGLE_TOTAL_BRAKE_FAILURE"
	//KeyToggleLeftBrakeFailure Toggles left brake failure
	KeyToggleLeftBrakeFailure KeySimEvent = "TOGGLE_LEFT_BRAKE_FAILURE"
	//KeyToggleRightBrakeFailure Toggles right brake failure
	KeyToggleRightBrakeFailure KeySimEvent = "TOGGLE_RIGHT_BRAKE_FAILURE"
	//KeyToggleEngine1Failure Toggle engine 1 failure
	KeyToggleEngine1Failure KeySimEvent = "TOGGLE_ENGINE1_FAILURE"
	//KeyToggleEngine2Failure Toggle engine 2 failure
	KeyToggleEngine2Failure KeySimEvent = "TOGGLE_ENGINE2_FAILURE"
	//KeyToggleEngine3Failure Toggle engine 3 failure
	KeyToggleEngine3Failure KeySimEvent = "TOGGLE_ENGINE3_FAILURE"
	//KeyToggleEngine4Failure Toggle engine 4 failure
	KeyToggleEngine4Failure KeySimEvent = "TOGGLE_ENGINE4_FAILURE"
	//KeySmokeToggle Toggle smoke system switch	All aircraft
	KeySmokeToggle KeySimEvent = "SMOKE_TOGGLE"
	//KeyGearToggle Toggle gear handle	All aircraft
	KeyGearToggle KeySimEvent = "GEAR_TOGGLE"
	//KeyBrakes Increment brake pressure  Note: These are simulated spring-loaded toe brakes, which will bleed back to zero over time.
	KeyBrakes KeySimEvent = "BRAKES"
	//KeyGearSet Sets gear handle position up/down (0,1)	All aircraft
	KeyGearSet KeySimEvent = "GEAR_SET"
	//KeyBrakesLeft Increments left brake pressure. Note: This is a simulated spring-loaded toe brake, which will bleed back to zero over time.
	KeyBrakesLeft KeySimEvent = "BRAKES_LEFT"
	//KeyBrakesRight Increments right brake pressure. Note: This is a simulated spring-loaded toe brake, which will bleed back to zero over time.
	KeyBrakesRight KeySimEvent = "BRAKES_RIGHT"
	//KeyParkingBrakes Toggles parking brake on/off
	KeyParkingBrakes KeySimEvent = "PARKING_BRAKES"
	//KeyGearPump Increments emergency gear extension
	KeyGearPump KeySimEvent = "GEAR_PUMP"
	//KeyPitotHeatToggle Toggles pitot heat switch	All aircraft
	KeyPitotHeatToggle KeySimEvent = "PITOT_HEAT_TOGGLE"
	//KeySmokeOn Turns smoke system on	All aircraft
	KeySmokeOn KeySimEvent = "SMOKE_ON"
	//KeySmokeOff Turns smoke system off	All aircraft
	KeySmokeOff KeySimEvent = "SMOKE_OFF"
	//KeySmokeSet Sets smoke system on/off (1,0)	All aircraft
	KeySmokeSet KeySimEvent = "SMOKE_SET"
	//KeyPitotHeatOn Turns pitot heat switch on
	KeyPitotHeatOn KeySimEvent = "PITOT_HEAT_ON"
	//KeyPitotHeatOff Turns pitot heat switch off
	KeyPitotHeatOff KeySimEvent = "PITOT_HEAT_OFF"
	//KeyPitotHeatSet Sets pitot heat switch on/off (1,0)
	KeyPitotHeatSet KeySimEvent = "PITOT_HEAT_SET"
	//KeyGearUp Sets gear handle in UP position	All aircraft
	KeyGearUp KeySimEvent = "GEAR_UP"
	//KeyGearDown Sets gear handle in DOWN position	All aircraft
	KeyGearDown KeySimEvent = "GEAR_DOWN"
	//KeyToggleMasterBattery Toggles main battery switch	All aircraft
	KeyToggleMasterBattery KeySimEvent = "TOGGLE_MASTER_BATTERY"
	//KeyToggleMasterAlternator Toggles main alternator/generator switch	All aircraft
	KeyToggleMasterAlternator KeySimEvent = "TOGGLE_MASTER_ALTERNATOR"
	//KeyToggleElectricVacuumPump Toggles backup electric vacuum pump
	KeyToggleElectricVacuumPump KeySimEvent = "TOGGLE_ELECTRIC_VACUUM_PUMP"
	//KeyToggleAlternateStatic Toggles alternate static pressure port	All aircraft
	KeyToggleAlternateStatic KeySimEvent = "TOGGLE_ALTERNATE_STATIC"
	//KeyDecisionHeightDec Decrements decision height reference
	KeyDecisionHeightDec KeySimEvent = "DECREASE_DECISION_HEIGHT"
	//KeyDecisionHeightInc Increments decision height reference
	KeyDecisionHeightInc KeySimEvent = "INCREASE_DECISION_HEIGHT"
	//KeyToggleStructuralDeice Toggles structural deice switch
	KeyToggleStructuralDeice KeySimEvent = "TOGGLE_STRUCTURAL_DEICE"
	//KeyTogglePropellerDeice Toggles propeller deice switch
	KeyTogglePropellerDeice KeySimEvent = "TOGGLE_PROPELLER_DEICE"
	//KeyToggleAlternator1 Toggles alternator/generator 1 switch	All aircraft
	KeyToggleAlternator1 KeySimEvent = "TOGGLE_ALTERNATOR1"
	//KeyToggleAlternator2 Toggles alternator/generator 2 switch	All aircraft
	KeyToggleAlternator2 KeySimEvent = "TOGGLE_ALTERNATOR2"
	//KeyToggleAlternator3 Toggles alternator/generator 3 switch	All aircraft
	KeyToggleAlternator3 KeySimEvent = "TOGGLE_ALTERNATOR3"
	//KeyToggleAlternator4 Toggles alternator/generator 4 switch	All aircraft
	KeyToggleAlternator4 KeySimEvent = "TOGGLE_ALTERNATOR4"
	//KeyToggleMasterBatteryAlternator Toggles master battery and alternator switch
	KeyToggleMasterBatteryAlternator KeySimEvent = "TOGGLE_MASTER_BATTERY_ALTERNATOR"
	//KeyAxisLeftBrakeSet Sets left brake position from axis controller (e.g. joystick). -16383 (0 brakes) to +16383 (max brakes)
	KeyAxisLeftBrakeSet KeySimEvent = "AXIS_LEFT_BRAKE_SET"
	//KeyAxisRightBrakeSet Sets right brake position from axis controller (e.g. joystick). -16383 (0 brakes) to +16383 (max brakes)
	KeyAxisRightBrakeSet KeySimEvent = "AXIS_RIGHT_BRAKE_SET"
	//KeyToggleAircraftExit Toggles primary door open/close. Follow by KEY_SELECT_2, etc for subsequent doors.
	KeyToggleAircraftExit KeySimEvent = "TOGGLE_AIRCRAFT_EXIT"
	//KeyToggleWingFold Toggles wing folding
	KeyToggleWingFold KeySimEvent = "TOGGLE_WING_FOLD"
	//KeySetWingFold Sets the wings into the folded position suitable for storage, typically on a carrier. Takes a value: 1 - fold wings, 0 - unfold wings
	KeySetWingFold KeySimEvent = "SET_WING_FOLD"
	//KeyToggleTailHookHandle Toggles tail hook
	KeyToggleTailHookHandle KeySimEvent = "TOGGLE_TAIL_HOOK_HANDLE"
	//KeySetTailHookHandle Sets the tail hook handle. Takes a value: 1 - set tail hook, 0 - retract tail hook
	KeySetTailHookHandle KeySimEvent = "SET_TAIL_HOOK_HANDLE"
	//KeyToggleWaterRudder Toggles water rudders
	KeyToggleWaterRudder KeySimEvent = "TOGGLE_WATER_RUDDER"
	//KeyPushbackSet Toggles pushback.
	KeyPushbackSet KeySimEvent = "TOGGLE_PUSHBACK"
	//KeyTugHeading Triggers tug and sets the desired heading. The units are a 32 bit integer (0 to 4294967295) which represent 0 to 360 degrees. To set a 45 degree angle, for example, set the value to 4294967295 / 8.
	KeyTugHeading KeySimEvent = "KeyTugHeading"
	//KeyTugSpeed Triggers tug, and sets desired speed, in feet per second. The speed can be both positive (forward movement) and negative (backward movement).
	KeyTugSpeed KeySimEvent = "KeyTugSpeed"
	//KeyTugDisable Disables tug
	KeyTugDisable KeySimEvent = "TUG_DISABLE"
	//KeyToggleMasterIgnitionSwitch Toggles master ignition switch
	KeyToggleMasterIgnitionSwitch KeySimEvent = "TOGGLE_MASTER_IGNITION_SWITCH"
	//KeyToggleTailwheelLock Toggles tail wheel lock
	KeyToggleTailwheelLock KeySimEvent = "TOGGLE_TAILWHEEL_LOCK"
	//KeyAddFuelQuantity Adds fuel to the aircraft, 25% of capacity by default. 0 to 65535 (max fuel) can be passed.
	KeyAddFuelQuantity KeySimEvent = "ADD_FUEL_QUANTITY"
	//KeyRotorBrake  Triggers rotor braking input
	KeyRotorBrake KeySimEvent = "ROTOR_BRAKE"
	//KeyRotorClutchSwitchToggle Toggles on electric rotor clutch switch
	KeyRotorClutchSwitchToggle KeySimEvent = "ROTOR_CLUTCH_SWITCH_TOGGLE"
	//KeyRotorClutchSwitchSet Sets electric rotor clutch switch on/off (1,0)
	KeyRotorClutchSwitchSet KeySimEvent = "ROTOR_CLUTCH_SWITCH_SET"
	//KeyRotorGovSwitchToggle Toggles the electric rotor governor switch
	KeyRotorGovSwitchToggle KeySimEvent = "ROTOR_GOV_SWITCH_TOGGLE"
	//KeyRotorGovSwitchSet Sets the electric rotor governor switch on/off (1,0)
	KeyRotorGovSwitchSet KeySimEvent = "ROTOR_GOV_SWITCH_SET"
	//KeyRotorLateralTrimInc Increments the lateral (right) rotor trim
	KeyRotorLateralTrimInc KeySimEvent = "ROTOR_LATERAL_TRIM_INC"
	//KeyRotorLateralTrimDec Decrements the lateral (right) rotor trim
	KeyRotorLateralTrimDec KeySimEvent = "ROTOR_LATERAL_TRIM_DEC"
	//KeyRotorLateralTrimSet Sets the lateral (right) rotor trim (0 to 16383)	 Slings and Hoists
	KeyRotorLateralTrimSet KeySimEvent = "ROTOR_LATERAL_TRIM_SET"
	//KeySlewToggle Toggles slew on/off	 (Pilot only)
	KeySlewToggle KeySimEvent = "SLEW_TOGGLE"
	//KeySlewOff Turns slew off	 (Pilot only)
	KeySlewOff KeySimEvent = "SLEW_OFF"
	//KeySlewOn Turns slew on	 (Pilot only)
	KeySlewOn KeySimEvent = "SLEW_ON"
	//KeySlewSet Sets slew on/off (1,0)	 (Pilot only)
	KeySlewSet KeySimEvent = "SLEW_SET"
	//KeySlewReset Stop slew and reset pitch, bank, and heading all to zero.	 (Pilot only)
	KeySlewReset KeySimEvent = "SLEW_RESET"
	//KeySlewAltitUpFast Slew upward fast	 (Pilot only)
	KeySlewAltitUpFast KeySimEvent = "SLEW_ALTIT_UP_FAST"
	//KeySlewAltitUpSlow Slew upward slow	 (Pilot only)
	KeySlewAltitUpSlow KeySimEvent = "SLEW_ALTIT_UP_SLOW"
	//KeySlewAltitFreeze Stop vertical slew	 (Pilot only)
	KeySlewAltitFreeze KeySimEvent = "SLEW_ALTIT_FREEZE"
	//KeySlewAltitDnSlow Slew downward slow	 (Pilot only)
	KeySlewAltitDnSlow KeySimEvent = "SLEW_ALTIT_DN_SLOW"
	//KeySlewAltitDnFast Slew downward fast	 (Pilot only)
	KeySlewAltitDnFast KeySimEvent = "SLEW_ALTIT_DN_FAST"
	//KeySlewAltitPlus Increase upward slew	 (Pilot only)
	KeySlewAltitPlus KeySimEvent = "SLEW_ALTIT_PLUS"
	//KeySlewAltitMinus Decrease upward slew 	 (Pilot only)
	KeySlewAltitMinus KeySimEvent = "SLEW_ALTIT_MINUS"
	//KeySlewPitchDnFast Slew pitch downward fast	 (Pilot only)
	KeySlewPitchDnFast KeySimEvent = "SLEW_PITCH_DN_FAST"
	//KeySlewPitchDnSlow Slew pitch downward slow	 (Pilot only)
	KeySlewPitchDnSlow KeySimEvent = "SLEW_PITCH_DN_SLOW"
	//KeySlewPitchFreeze Stop pitch slew	 (Pilot only)
	KeySlewPitchFreeze KeySimEvent = "SLEW_PITCH_FREEZE"
	//KeySlewPitchUpSlow Slew pitch up slow	 (Pilot only)
	KeySlewPitchUpSlow KeySimEvent = "SLEW_PITCH_UP_SLOW"
	//KeySlewPitchUpFast Slew pitch upward fast	 (Pilot only)
	KeySlewPitchUpFast KeySimEvent = "SLEW_PITCH_UP_FAST"
	//KeySlewPitchPlus Increase pitch up slew	 (Pilot only)
	KeySlewPitchPlus KeySimEvent = "SLEW_PITCH_PLUS"
	//KeySlewPitchMinus Decrease pitch up slew	 (Pilot only)
	KeySlewPitchMinus KeySimEvent = "SLEW_PITCH_MINUS"
	//KeySlewBankMinus Increase left bank slew	 (Pilot only)
	KeySlewBankMinus KeySimEvent = "SLEW_BANK_MINUS"
	//KeySlewAheadPlus Increase forward slew	 (Pilot only)
	KeySlewAheadPlus KeySimEvent = "SLEW_AHEAD_PLUS"
	//KeySlewBankPlus Increase right bank slew	 (Pilot only)
	KeySlewBankPlus KeySimEvent = "SLEW_BANK_PLUS"
	//KeySlewLeft Slew to the left	 (Pilot only)
	KeySlewLeft KeySimEvent = "SLEW_LEFT"
	//KeySlewFreeze Stop all slew	 (Pilot only)
	KeySlewFreeze KeySimEvent = "SLEW_FREEZE"
	//KeySlewRight Slew to the right	 (Pilot only)
	KeySlewRight KeySimEvent = "SLEW_RIGHT"
	//KeySlewHeadingMinus Increase slew heading to the left	 (Pilot only)
	KeySlewHeadingMinus KeySimEvent = "SLEW_HEADING_MINUS"
	//KeySlewAheadMinus Decrease forward slew	 (Pilot only)
	KeySlewAheadMinus KeySimEvent = "SLEW_AHEAD_MINUS"
	//KeySlewHeadingPlus Increase slew heading to the right	 (Pilot only)
	KeySlewHeadingPlus KeySimEvent = "SLEW_HEADING_PLUS"
	//KeyAxisSlewAheadSet Sets forward slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewAheadSet KeySimEvent = "AXIS_SLEW_AHEAD_SET"
	//KeyAxisSlewSidewaysSet Sets sideways slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewSidewaysSet KeySimEvent = "AXIS_SLEW_SIDEWAYS_SET"
	//KeyAxisSlewHeadingSet Sets heading slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewHeadingSet KeySimEvent = "AXIS_SLEW_HEADING_SET"
	//KeyAxisSlewAltSet Sets vertical slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewAltSet KeySimEvent = "AXIS_SLEW_ALT_SET"
	//KeyAxisSlewBankSet Sets roll slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewBankSet KeySimEvent = "AXIS_SLEW_BANK_SET"
	//KeyAxisSlewPitchSet Sets pitch slew (+/- 16383)	 (Pilot only)
	KeyAxisSlewPitchSet KeySimEvent = "AXIS_SLEW_PITCH_SET"
	//KeyViewMode Selects next view
	KeyViewMode KeySimEvent = "VIEW_MODE"
	//KeyViewWindowToFront Sets active window to front
	KeyViewWindowToFront KeySimEvent = "VIEW_WINDOW_TO_FRONT"
	//KeyViewReset Reset view forward
	KeyViewReset KeySimEvent = "VIEW_RESET"
	//KeyViewAlwaysPanUp SimEvent
	KeyViewAlwaysPanUp KeySimEvent = "VIEW_ALWAYS_PAN_UP"
	//KeyViewAlwaysPanDown SimEvent
	KeyViewAlwaysPanDown KeySimEvent = "VIEW_ALWAYS_PAN_DOWN"
	//KeyNextSubView SimEvent
	KeyNextSubView KeySimEvent = "NEXT_SUB_VIEW"
	//KeyPrevSubView SimEvent
	KeyPrevSubView KeySimEvent = "PREV_SUB_VIEW"
	//KeyViewTrackPanToggle SimEvent
	KeyViewTrackPanToggle KeySimEvent = "VIEW_TRACK_PAN_TOGGLE"
	//KeyViewPreviousToggle SimEvent
	KeyViewPreviousToggle KeySimEvent = "VIEW_PREVIOUS_TOGGLE"
	//KeyViewCameraSelectStarting SimEvent
	KeyViewCameraSelectStarting KeySimEvent = "VIEW_CAMERA_SELECT_START"
	//KeyPanelHudNext SimEvent
	KeyPanelHudNext KeySimEvent = "PANEL_HUD_NEXT"
	//KeyPanelHudPrevious SimEvent
	KeyPanelHudPrevious KeySimEvent = "PANEL_HUD_PREVIOUS"
	//KeyZoomIn Zooms view in
	KeyZoomIn KeySimEvent = "ZOOM_IN"
	//KeyZoomOut Zooms view out
	KeyZoomOut KeySimEvent = "ZOOM_OUT"
	//KeyMapZoomFineIn Fine zoom in map view
	KeyMapZoomFineIn KeySimEvent = "MAP_ZOOM_FINE_IN"
	//KeyPanLeft Pans view left
	KeyPanLeft KeySimEvent = "PAN_LEFT"
	//KeyPanRight Pans view right
	KeyPanRight KeySimEvent = "PAN_RIGHT"
	//KeyMapZoomFineOut Fine zoom out in map view
	KeyMapZoomFineOut KeySimEvent = "MAP_ZOOM_FINE_OUT"
	//KeyViewForward Sets view direction forward
	KeyViewForward KeySimEvent = "VIEW_FORWARD"
	//KeyViewForwardRight Sets view direction forward and right
	KeyViewForwardRight KeySimEvent = "VIEW_FORWARD_RIGHT"
	//KeyViewRight Sets view direction to the right
	KeyViewRight KeySimEvent = "VIEW_RIGHT"
	//KeyViewRearRight Sets view direction to the rear and right
	KeyViewRearRight KeySimEvent = "VIEW_REAR_RIGHT"
	//KeyViewRear Sets view direction to the rear
	KeyViewRear KeySimEvent = "VIEW_REAR"
	//KeyViewRearLeft Sets view direction to the rear and left
	KeyViewRearLeft KeySimEvent = "VIEW_REAR_LEFT"
	//KeyViewLeft Sets view direction to the left
	KeyViewLeft KeySimEvent = "VIEW_LEFT"
	//KeyViewForwardLeft Sets view direction forward and left
	KeyViewForwardLeft KeySimEvent = "VIEW_FORWARD_LEFT"
	//KeyViewDown Sets view direction down
	KeyViewDown KeySimEvent = "VIEW_DOWN"
	//KeyZoomMinus Decreases zoom
	KeyZoomMinus KeySimEvent = "ZOOM_MINUS"
	//KeyZoomPlus Increase zoom
	KeyZoomPlus KeySimEvent = "ZOOM_PLUS"
	//KeyPanUp Pan view up
	KeyPanUp KeySimEvent = "PAN_UP"
	//KeyPanDown Pan view down
	KeyPanDown KeySimEvent = "PAN_DOWN"
	//KeyViewModeRev Reverse view cycle
	KeyViewModeRev KeySimEvent = "VIEW_MODE_REV"
	//KeyZoomInFine Zoom in fine
	KeyZoomInFine KeySimEvent = "ZOOM_IN_FINE"
	//KeyZoomOutFine Zoom out fine
	KeyZoomOutFine KeySimEvent = "ZOOM_OUT_FINE"
	//KeyCloseView Close current view
	KeyCloseView KeySimEvent = "CLOSE_VIEW"
	//KeyNewView Open new view
	KeyNewView KeySimEvent = "NEW_VIEW"
	//KeyNextView Select next view
	KeyNextView KeySimEvent = "NEXT_VIEW"
	//KeyPrevView Select previous view
	KeyPrevView KeySimEvent = "PREV_VIEW"
	//KeyPanLeftUp Pan view left
	KeyPanLeftUp KeySimEvent = "PAN_LEFT_UP"
	//KeyPanLeftDown Pan view left and down
	KeyPanLeftDown KeySimEvent = "PAN_LEFT_DOWN"
	//KeyPanRightUp Pan view right and up
	KeyPanRightUp KeySimEvent = "PAN_RIGHT_UP"
	//KeyPanRightDown Pan view right and down
	KeyPanRightDown KeySimEvent = "PAN_RIGHT_DOWN"
	//KeyPanTiltLeft Tilt view left
	KeyPanTiltLeft KeySimEvent = "PAN_TILT_LEFT"
	//KeyPanTiltRight Tilt view right
	KeyPanTiltRight KeySimEvent = "PAN_TILT_RIGHT"
	//KeyPanReset Reset view to forward
	KeyPanReset KeySimEvent = "PAN_RESET"
	//KeyViewForwardUp Sets view forward and up
	KeyViewForwardUp KeySimEvent = "VIEW_FORWARD_UP"
	//KeyViewForwardRightUp Sets view forward, right, and up
	KeyViewForwardRightUp KeySimEvent = "VIEW_FORWARD_RIGHT_UP"
	//KeyViewRightUp Sets view right and up
	KeyViewRightUp KeySimEvent = "VIEW_RIGHT_UP"
	//KeyViewRearRightUp Sets view rear, right, and up
	KeyViewRearRightUp KeySimEvent = "VIEW_REAR_RIGHT_UP"
	//KeyViewRearUp Sets view rear and up
	KeyViewRearUp KeySimEvent = "VIEW_REAR_UP"
	//KeyViewRearLeftUp Sets view rear left and up
	KeyViewRearLeftUp KeySimEvent = "VIEW_REAR_LEFT_UP"
	//KeyViewLeftUp Sets view left and up
	KeyViewLeftUp KeySimEvent = "VIEW_LEFT_UP"
	//KeyViewForwardLeftUp Sets view forward left and up
	KeyViewForwardLeftUp KeySimEvent = "VIEW_FORWARD_LEFT_UP"
	//KeyViewUp Sets view up
	KeyViewUp KeySimEvent = "VIEW_UP"
	//KeyPanResetCockpit Reset panning to forward, if in cockpit view
	KeyPanResetCockpit KeySimEvent = "PAN_RESET_COCKPIT"
	//KeyChaseViewNext Cycle view to next target
	KeyChaseViewNext KeySimEvent = "KeyChaseViewNext"
	//KeyChaseViewPrev Cycle view to previous target
	KeyChaseViewPrev KeySimEvent = "KeyChaseViewPrev"
	//KeyChaseViewToggle Toggles chase view on/off
	KeyChaseViewToggle KeySimEvent = "CHASE_VIEW_TOGGLE"
	//KeyEyepointUp Move eyepoint up
	KeyEyepointUp KeySimEvent = "EYEPOINT_UP"
	//KeyEyepointDown Move eyepoint down
	KeyEyepointDown KeySimEvent = "EYEPOINT_DOWN"
	//KeyEyepointRight Move eyepoint right
	KeyEyepointRight KeySimEvent = "EYEPOINT_RIGHT"
	//KeyEyepointLeft Move eyepoint left
	KeyEyepointLeft KeySimEvent = "EYEPOINT_LEFT"
	//KeyEyepointForward Move eyepoint forward
	KeyEyepointForward KeySimEvent = "EYEPOINT_FORWARD"
	//KeyEyepointBack Move eyepoint backward
	KeyEyepointBack KeySimEvent = "EYEPOINT_BACK"
	//KeyEyepointReset Move eyepoint to default position
	KeyEyepointReset KeySimEvent = "EYEPOINT_RESET"
	//KeyNewMap Opens new map view
	KeyNewMap KeySimEvent = "NEW_MAP"
	//KeyPauseToggle Toggles pause on/off	Disabled
	KeyPauseToggle KeySimEvent = "PAUSE_TOGGLE"
	//KeyPauseOn Turns pause on	Disabled
	KeyPauseOn KeySimEvent = "PAUSE_ON"
	//KeyPauseOff Turns pause off	Disabled
	KeyPauseOff KeySimEvent = "PAUSE_OFF"
	//KeyPauseSet Sets pause on/off (1,0)	Disabled
	KeyPauseSet KeySimEvent = "PAUSE_SET"
	//KeyDemoStop Stops demo system playback
	KeyDemoStop KeySimEvent = "DEMO_STOP"
	//KeySelect1 Sets "selected" index (for other events) to 1
	KeySelect1 KeySimEvent = "SELECT_1"
	//KeySelect2 Sets "selected" index (for other events) to 2
	KeySelect2 KeySimEvent = "SELECT_2"
	//KeySelect3 Sets "selected" index (for other events) to 3
	KeySelect3 KeySimEvent = "SELECT_3"
	//KeySelect4 Sets "selected" index (for other events) to 4
	KeySelect4 KeySimEvent = "SELECT_4"
	//KeyMinus Used in conjunction with "selected" parameters to decrease their value (e.g., radio frequency)
	KeyMinus KeySimEvent = "MINUS"
	//KeyPlus Used in conjunction with "selected" parameters to increase their value (e.g., radio frequency)
	KeyPlus KeySimEvent = "PLUS"
	//KeyZoom1x Sets zoom level to 1
	KeyZoom1x KeySimEvent = "ZOOM_1X"
	//KeySoundToggle Toggles sound on/off
	KeySoundToggle KeySimEvent = "SOUND_TOGGLE"
	//KeySimRate Selects simulation rate (use KEY_MINUS, KEY_PLUS to change)
	KeySimRate KeySimEvent = "SIM_RATE"
	//KeyJoystickCalibrate Toggles joystick on/off
	KeyJoystickCalibrate KeySimEvent = "JOYSTICK_CALIBRATE"
	//KeySituationSave Saves scenario
	KeySituationSave KeySimEvent = "SITUATION_SAVE"
	//KeySituationReset Resets scenario
	KeySituationReset KeySimEvent = "SITUATION_RESET"
	//KeySoundSet Sets sound on/off (1,0)
	KeySoundSet KeySimEvent = "SOUND_SET"
	//KeyExit Quit Prepar3D with a message
	KeyExit KeySimEvent = "EXIT"
	//KeyAbort Quit Prepar3D without a message
	KeyAbort KeySimEvent = "ABORT"
	//KeyReadoutsSlew Cycle through information readouts while in slew
	KeyReadoutsSlew KeySimEvent = "READOUTS_SLEW"
	//KeyReadoutsFlight Cycle through information readouts
	KeyReadoutsFlight KeySimEvent = "READOUTS_FLIGHT"
	//KeyMinusShift Used with other events
	KeyMinusShift KeySimEvent = "MINUS_SHIFT"
	//KeyPlusShift Used with other events
	KeyPlusShift KeySimEvent = "PLUS_SHIFT"
	//KeySimRateIncr Increase sim rate
	KeySimRateIncr KeySimEvent = "SIM_RATE_INCR"
	//KeySimRateDecr Decrease sim rate
	KeySimRateDecr KeySimEvent = "SIM_RATE_DECR"
	//KeyKneeboard Toggles kneeboard
	KeyKneeboard KeySimEvent = "KNEEBOARD_VIEW"
	//KeyPanel1 Toggles panel 1
	KeyPanel1 KeySimEvent = "PANEL_1"
	//KeyPanel2 Toggles panel 2
	KeyPanel2 KeySimEvent = "PANEL_2"
	//KeyPanel3 Toggles panel 3
	KeyPanel3 KeySimEvent = "PANEL_3"
	//KeyPanel4 Toggles panel 4
	KeyPanel4 KeySimEvent = "PANEL_4"
	//KeyPanel5 Toggles panel 5
	KeyPanel5 KeySimEvent = "PANEL_5"
	//KeyPanel6 Toggles panel 6
	KeyPanel6 KeySimEvent = "PANEL_6"
	//KeyPanel7 Toggles panel 7
	KeyPanel7 KeySimEvent = "PANEL_7"
	//KeyPanel8 Toggles panel 8
	KeyPanel8 KeySimEvent = "PANEL_8"
	//KeyPanel9 Toggles panel 9
	KeyPanel9 KeySimEvent = "PANEL_9"
	//KeySoundOn Turns sound on
	KeySoundOn KeySimEvent = "SOUND_ON"
	//KeySoundOff Turns sound off
	KeySoundOff KeySimEvent = "SOUND_OFF"
	//KeyInvokeHelp Brings up Help system
	KeyInvokeHelp KeySimEvent = "INVOKE_HELP"
	//KeyToggleAircraftLabels Toggles aircraft labels
	KeyToggleAircraftLabels KeySimEvent = "TOGGLE_AIRCRAFT_LABELS"
	//KeyFlightMap Brings up flight map
	KeyFlightMap KeySimEvent = "FLIGHT_MAP"
	//KeyReloadPanels Reload panel data
	KeyReloadPanels  KeySimEvent = "RELOAD_PANELS"
	KeyPanelIDToggle KeySimEvent = "PANEL_ID_TOGGLE"
	KeyPanelIDOpen   KeySimEvent = "PANEL_ID_OPEN"
	KeyPanelIDClose  KeySimEvent = "PANEL_ID_CLOSE"
	//KeyControlReloadUserAircraft Reloads the user aircraft data (from cache if same type loaded as an AI, otherwise from disk)
	KeyControlReloadUserAircraft KeySimEvent = "RELOAD_USER_AIRCRAFT"
	//KeySimReset Resets aircraft state
	KeySimReset KeySimEvent = "SIM_RESET"
	//KeyVirtualCopilotToggle Turns User Tips on/off
	KeyVirtualCopilotToggle KeySimEvent = "VIRTUAL_COPILOT_TOGGLE"
	//KeyVirtualCopilotSet Sets User Tips on/off (1,0)
	KeyVirtualCopilotSet KeySimEvent = "VIRTUAL_COPILOT_SET"
	//KeyVirtualCopilotAction Triggers action noted in User Tips
	KeyVirtualCopilotAction KeySimEvent = "VIRTUAL_COPILOT_ACTION"
	//KeyRefreshScenery Reloads scenery
	KeyRefreshScenery KeySimEvent = "REFRESH_SCENERY"
	//KeyClockHoursDec Decrements time by hours
	KeyClockHoursDec KeySimEvent = "CLOCK_HOURS_DEC"
	//KeyClockHoursInc Increments time by hours
	KeyClockHoursInc KeySimEvent = "CLOCK_HOURS_INC"
	//KeyClockMinutesDec Decrements time by minutes
	KeyClockMinutesDec KeySimEvent = "CLOCK_MINUTES_DEC"
	//KeyClockMinutesInc Increments time by minutes
	KeyClockMinutesInc KeySimEvent = "CLOCK_MINUTES_INC"
	//KeyClockSecondsZero Zeros seconds
	KeyClockSecondsZero KeySimEvent = "CLOCK_SECONDS_ZERO"
	//KeyClockHoursSet Sets hour of day
	KeyClockHoursSet KeySimEvent = "CLOCK_HOURS_SET"
	//KeyClockMinutesSet Sets minutes of the hour
	KeyClockMinutesSet KeySimEvent = "CLOCK_MINUTES_SET"
	//KeyZuluHoursSet Sets hours, zulu time
	KeyZuluHoursSet KeySimEvent = "ZULU_HOURS_SET"
	//KeyZuluMinutesSet Sets minutes, in zulu time
	KeyZuluMinutesSet KeySimEvent = "ZULU_MINUTES_SET"
	//KeyZuluDaySet Sets day, in zulu time
	KeyZuluDaySet KeySimEvent = "ZULU_DAY_SET"
	//KeyZuluYearSet Sets year, in zulu time
	KeyZuluYearSet KeySimEvent = "ZULU_YEAR_SET"
	//KeyAtc Activates ATC window
	KeyAtc KeySimEvent = "ATC"
	//KeyAtcMenu1 Selects ATC option 1
	KeyAtcMenu1 KeySimEvent = "ATC_MENU_1"
	//KeyAtcMenu2 Selects ATC option 2
	KeyAtcMenu2 KeySimEvent = "ATC_MENU_2"
	//KeyAtcMenu3 Selects ATC option 3
	KeyAtcMenu3 KeySimEvent = "ATC_MENU_3"
	//KeyAtcMenu4 Selects ATC option 4
	KeyAtcMenu4 KeySimEvent = "ATC_MENU_4"
	//KeyAtcMenu5 Selects ATC option 5
	KeyAtcMenu5 KeySimEvent = "ATC_MENU_5"
	//KeyAtcMenu6 Selects ATC option 6
	KeyAtcMenu6 KeySimEvent = "ATC_MENU_6"
	//KeyAtcMenu7 Selects ATC option 7
	KeyAtcMenu7 KeySimEvent = "ATC_MENU_7"
	//KeyAtcMenu8 Selects ATC option 8
	KeyAtcMenu8 KeySimEvent = "ATC_MENU_8"
	//KeyAtcMenu9 Selects ATC option 9
	KeyAtcMenu9 KeySimEvent = "ATC_MENU_9"
	//KeyAtcMenu0 Selects ATC option 10
	KeyAtcMenu0 KeySimEvent = "ATC_MENU_0"
	//KeyMultiplayerTransferControl Toggle to the next player to track	-
	KeyMultiplayerTransferControl KeySimEvent = "MP_TRANSFER_CONTROL"
	//KeyMultiplayerPlayerCycle Cycle through the current user aircraft.
	KeyMultiplayerPlayerCycle KeySimEvent = "MP_PLAYER_CYCLE"
	//KeyMultiplayerPlayerFollow Set the view to follow the selected user aircraft.
	KeyMultiplayerPlayerFollow KeySimEvent = "MP_PLAYER_FOLLOW"
	//KeyMultiplayerChat Toggles chat window visible/invisible
	KeyMultiplayerChat KeySimEvent = "MP_CHAT"
	//KeyMultiplayerActivateChat Activates chat window
	KeyMultiplayerActivateChat KeySimEvent = "MP_ACTIVATE_CHAT"
	//KeyMultiplayerVoiceCaptureStart Start capturing audio from the users computer and transmitting it to all other players in the multiplayer session who are turned to the same radio frequency.
	KeyMultiplayerVoiceCaptureStart KeySimEvent = "MP_VOICE_CAPTURE_START"
	//KeyMultiplayerVoiceCaptureStop Stop capturing radio audio.
	KeyMultiplayerVoiceCaptureStop KeySimEvent = "MP_VOICE_CAPTURE_STOP"
	//KeyMultiplayerBroadcastVoiceCaptureStart Start capturing audio from the users computer and transmitting it to all other players in the multiplayer session.
	KeyMultiplayerBroadcastVoiceCaptureStart KeySimEvent = "MP_BROADCAST_VOICE_CAPTURE_START"
	//KeyMultiplayerBroadcastVoiceCaptureStop Stop capturing broadcast audio.
	KeyMultiplayerBroadcastVoiceCaptureStop KeySimEvent = "MP_BROADCAST_VOICE_CAPTURE_STOP"
)
