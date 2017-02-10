package p2p

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdp2pConnectionParcelJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionParcel_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionParcel_JSONByte",
	})
	factomdp2pConnectionParcelJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionParcel_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionParcel_JSONString",
	})
	factomdp2pConnectionParcelString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionParcel_String_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionParcel_String",
	})
	factomdp2pConnectionCommandJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionCommand_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionCommand_JSONByte",
	})
	factomdp2pConnectionCommandJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionCommand_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionCommand_JSONString",
	})
	factomdp2pConnectionCommandString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ConnectionCommand_String_Summary",
		Help: "Summary of calls to  factomd_p2p_ConnectionCommand_String",
	})
	factomdp2pConnectionInitWithConn = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_InitWithConn_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_InitWithConn",
	})
	factomdp2pConnectionInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_Init",
	})
	factomdp2pConnectionIsOutGoing = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_IsOutGoing_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_IsOutGoing",
	})
	factomdp2pConnectionIsOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_IsOnline_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_IsOnline",
	})
	factomdp2pConnectionStatusString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_StatusString_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_StatusString",
	})
	factomdp2pConnectionIsPersistent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_IsPersistent_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_IsPersistent",
	})
	factomdp2pConnectionNotes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_Notes_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_Notes",
	})
	factomdp2pConnectioncommonInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_commonInit_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_commonInit",
	})
	factomdp2pConnectionStart = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_Start_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_Start",
	})
	factomdp2pConnectionrunLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_runLoop_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_runLoop",
	})
	factomdp2pConnectionsetNotes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_setNotes_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_setNotes",
	})
	factomdp2pConnectiondialLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_dialLoop_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_dialLoop",
	})
	factomdp2pConnectiondial = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_dial_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_dial",
	})
	factomdp2pConnectiongoOnline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_goOnline_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_goOnline",
	})
	factomdp2pConnectiongoOffline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_goOffline_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_goOffline",
	})
	factomdp2pConnectiongoShutdown = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_goShutdown_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_goShutdown",
	})
	factomdp2pConnectionprocessSends = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_processSends_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_processSends",
	})
	factomdp2pConnectionhandleCommand = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_handleCommand_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_handleCommand",
	})
	factomdp2pConnectionsendParcel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_sendParcel_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_sendParcel",
	})
	factomdp2pConnectionprocessReceives = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_processReceives_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_processReceives",
	})
	factomdp2pConnectionhandleNetErrors = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_handleNetErrors_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_handleNetErrors",
	})
	factomdp2pConnectionhandleParcel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_handleParcel_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_handleParcel",
	})
	factomdp2pConnectionparcelValidity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_parcelValidity_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_parcelValidity",
	})
	factomdp2pConnectionhandleParcelTypes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_handleParcelTypes_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_handleParcelTypes",
	})
	factomdp2pConnectionpingPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_pingPeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_pingPeer",
	})
	factomdp2pConnectionupdatePeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_updatePeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_updatePeer",
	})
	factomdp2pConnectionupdateStats = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_updateStats_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_updateStats",
	})
	factomdp2pConnectionConnectionState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_ConnectionState_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_ConnectionState",
	})
	factomdp2pConnectionconnectionStatusReport = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Connection_connectionStatusReport_Summary",
		Help: "Summary of calls to  factomd_p2p_Connection_connectionStatusReport",
	})
	factomdp2pCommandAdjustPeerQualityJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandAdjustPeerQuality_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandAdjustPeerQuality_JSONByte",
	})
	factomdp2pCommandAdjustPeerQualityJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandAdjustPeerQuality_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandAdjustPeerQuality_JSONString",
	})
	factomdp2pCommandAdjustPeerQualityString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandAdjustPeerQuality_String_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandAdjustPeerQuality_String",
	})
	factomdp2pCommandBanJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandBan_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandBan_JSONByte",
	})
	factomdp2pCommandBanJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandBan_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandBan_JSONString",
	})
	factomdp2pCommandBanString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandBan_String_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandBan_String",
	})
	factomdp2pCommandDisconnectJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandDisconnect_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandDisconnect_JSONByte",
	})
	factomdp2pCommandDisconnectJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandDisconnect_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandDisconnect_JSONString",
	})
	factomdp2pCommandDisconnectString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandDisconnect_String_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandDisconnect_String",
	})
	factomdp2pCommandChangeLoggingJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandChangeLogging_JSONByte_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandChangeLogging_JSONByte",
	})
	factomdp2pCommandChangeLoggingJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandChangeLogging_JSONString_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandChangeLogging_JSONString",
	})
	factomdp2pCommandChangeLoggingString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_CommandChangeLogging_String_Summary",
		Help: "Summary of calls to  factomd_p2p_CommandChangeLogging_String",
	})
	factomdp2pControllerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_Init",
	})
	factomdp2pControllerStartNetwork = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_StartNetwork_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_StartNetwork",
	})
	factomdp2pControllerDialSpecialPeersString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_DialSpecialPeersString_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_DialSpecialPeersString",
	})
	factomdp2pControllerStartLogging = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_StartLogging_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_StartLogging",
	})
	factomdp2pControllerStopLogging = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_StopLogging_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_StopLogging",
	})
	factomdp2pControllerChangeLogLevel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_ChangeLogLevel_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_ChangeLogLevel",
	})
	factomdp2pControllerDialPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_DialPeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_DialPeer",
	})
	factomdp2pControllerAddPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_AddPeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_AddPeer",
	})
	factomdp2pControllerNetworkStop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_NetworkStop_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_NetworkStop",
	})
	factomdp2pControllerAdjustPeerQuality = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_AdjustPeerQuality_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_AdjustPeerQuality",
	})
	factomdp2pControllerBan = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_Ban_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_Ban",
	})
	factomdp2pControllerDisconnect = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_Disconnect_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_Disconnect",
	})
	factomdp2pControllerGetNumberConnections = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_GetNumberConnections_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_GetNumberConnections",
	})
	factomdp2pControllerlisten = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_listen_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_listen",
	})
	factomdp2pControlleracceptLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_acceptLoop_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_acceptLoop",
	})
	factomdp2pControllerrunloop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_runloop_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_runloop",
	})
	factomdp2pControllerroute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_route_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_route",
	})
	factomdp2pControllerdoDirectedSend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_doDirectedSend_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_doDirectedSend",
	})
	factomdp2pControllerhandleParcelReceive = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_handleParcelReceive_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_handleParcelReceive",
	})
	factomdp2pControllerhandleConnectionCommand = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_handleConnectionCommand_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_handleConnectionCommand",
	})
	factomdp2pControllerhandleCommand = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_handleCommand_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_handleCommand",
	})
	factomdp2pControllerapplicationPeerUpdate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_applicationPeerUpdate_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_applicationPeerUpdate",
	})
	factomdp2pControllermanagePeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_managePeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_managePeers",
	})
	factomdp2pControllerupdateConnectionCounts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_updateConnectionCounts_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_updateConnectionCounts",
	})
	factomdp2pControllerupdateConnectionAddressMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_updateConnectionAddressMap_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_updateConnectionAddressMap",
	})
	factomdp2pControllerweAreNotAlreadyConnectedTo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_weAreNotAlreadyConnectedTo_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_weAreNotAlreadyConnectedTo",
	})
	factomdp2pControllerfillOutgoingSlots = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_fillOutgoingSlots_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_fillOutgoingSlots",
	})
	factomdp2pControllerupdateMetrics = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_updateMetrics_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_updateMetrics",
	})
	factomdp2pControllershutdown = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_shutdown_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_shutdown",
	})
	factomdp2pControllernetworkStatusReport = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Controller_networkStatusReport_Summary",
		Help: "Summary of calls to  factomd_p2p_Controller_networkStatusReport",
	})
	factomdp2pDiscoveryInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_Init",
	})
	factomdp2pDiscoveryupdatePeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_updatePeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_updatePeer",
	})
	factomdp2pDiscoverygetPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_getPeer_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_getPeer",
	})
	factomdp2pDiscoveryisPeerPresent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_isPeerPresent_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_isPeerPresent",
	})
	factomdp2pDiscoveryLoadPeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_LoadPeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_LoadPeers",
	})
	factomdp2pDiscoverySavePeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_SavePeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_SavePeers",
	})
	factomdp2pDiscoveryLearnPeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_LearnPeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_LearnPeers",
	})
	factomdp2pDiscoveryupdatePeerSource = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_updatePeerSource_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_updatePeerSource",
	})
	factomdp2pDiscoveryfilterPeersFromOtherNetworks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_filterPeersFromOtherNetworks_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_filterPeersFromOtherNetworks",
	})
	factomdp2pDiscoveryfilterForUniqueIPAdresses = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_filterForUniqueIPAdresses_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_filterForUniqueIPAdresses",
	})
	factomdp2pDiscoveryGetOutgoingPeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_GetOutgoingPeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_GetOutgoingPeers",
	})
	factomdp2pDiscoverySharePeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_SharePeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_SharePeers",
	})
	factomdp2pDiscoverygetPeerSelection = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_getPeerSelection_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_getPeerSelection",
	})
	factomdp2pDiscoveryDiscoverPeersFromSeed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_DiscoverPeersFromSeed_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_DiscoverPeersFromSeed",
	})
	factomdp2pDiscoveryPrintPeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Discovery_PrintPeers_Summary",
		Help: "Summary of calls to  factomd_p2p_Discovery_PrintPeers",
	})
	factomdp2pNewParcel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_NewParcel_Summary",
		Help: "Summary of calls to  factomd_p2p_NewParcel",
	})
	factomdp2pParcelsForPayload = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ParcelsForPayload_Summary",
		Help: "Summary of calls to  factomd_p2p_ParcelsForPayload",
	})
	factomdp2pReassembleParcel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ReassembleParcel_Summary",
		Help: "Summary of calls to  factomd_p2p_ReassembleParcel",
	})
	factomdp2pParcelHeaderInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ParcelHeader_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_ParcelHeader_Init",
	})
	factomdp2pParcelInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_Init",
	})
	factomdp2pParcelUpdateHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_UpdateHeader_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_UpdateHeader",
	})
	factomdp2pParcelTrace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_Trace_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_Trace",
	})
	factomdp2pParcelHeaderPrint = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_ParcelHeader_Print_Summary",
		Help: "Summary of calls to  factomd_p2p_ParcelHeader_Print",
	})
	factomdp2pParcelPrint = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_Print_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_Print",
	})
	factomdp2pParcelMessageType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_MessageType_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_MessageType",
	})
	factomdp2pParcelPrintMessageType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_PrintMessageType_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_PrintMessageType",
	})
	factomdp2pParcelString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Parcel_String_Summary",
		Help: "Summary of calls to  factomd_p2p_Parcel_String",
	})
	factomdp2pPartsAssemblerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PartsAssembler_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_PartsAssembler_Init",
	})
	factomdp2pPartsAssemblerhandlePart = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PartsAssembler_handlePart_Summary",
		Help: "Summary of calls to  factomd_p2p_PartsAssembler_handlePart",
	})
	factomdp2pvalidateParcelPart = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_validateParcelPart_Summary",
		Help: "Summary of calls to  factomd_p2p_validateParcelPart",
	})
	factomdp2pPartsAssemblercleanupOldPartialMessages = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PartsAssembler_cleanupOldPartialMessages_Summary",
		Help: "Summary of calls to  factomd_p2p_PartsAssembler_cleanupOldPartialMessages",
	})
	factomdp2pcreateNewPartialMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_createNewPartialMessage_Summary",
		Help: "Summary of calls to  factomd_p2p_createNewPartialMessage",
	})
	factomdp2ptryReassemblingMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_tryReassemblingMessage_Summary",
		Help: "Summary of calls to  factomd_p2p_tryReassemblingMessage",
	})
	factomdp2pPeerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_Init_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_Init",
	})
	factomdp2pPeergeneratePeerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_generatePeerHash_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_generatePeerHash",
	})
	factomdp2pPeerAddressPort = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_AddressPort_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_AddressPort",
	})
	factomdp2pPeerPeerIdent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_PeerIdent_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_PeerIdent",
	})
	factomdp2pPeerPeerFixedIdent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_PeerFixedIdent_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_PeerFixedIdent",
	})
	factomdp2pPeerLocationFromAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_LocationFromAddress_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_LocationFromAddress",
	})
	factomdp2pPeermerit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_merit_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_merit",
	})
	factomdp2pPeerdemerit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_Peer_demerit_Summary",
		Help: "Summary of calls to  factomd_p2p_Peer_demerit",
	})
	factomdp2pPeerQualitySortLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerQualitySort_Len_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerQualitySort_Len",
	})
	factomdp2pPeerQualitySortSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerQualitySort_Swap_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerQualitySort_Swap",
	})
	factomdp2pPeerQualitySortLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerQualitySort_Less_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerQualitySort_Less",
	})
	factomdp2pPeerDistanceSortLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerDistanceSort_Len_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerDistanceSort_Len",
	})
	factomdp2pPeerDistanceSortSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerDistanceSort_Swap_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerDistanceSort_Swap",
	})
	factomdp2pPeerDistanceSortLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_PeerDistanceSort_Less_Summary",
		Help: "Summary of calls to  factomd_p2p_PeerDistanceSort_Less",
	})
	factomdp2pBlockFreeChannelSend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_BlockFreeChannelSend_Summary",
		Help: "Summary of calls to  factomd_p2p_BlockFreeChannelSend",
	})
	factomdp2pNetworkIDString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_NetworkID_String_Summary",
		Help: "Summary of calls to  factomd_p2p_NetworkID_String",
	})
	factomdp2pdot = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_dot_Summary",
		Help: "Summary of calls to  factomd_p2p_dot",
	})
	factomdp2psilence = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_silence_Summary",
		Help: "Summary of calls to  factomd_p2p_silence",
	})
	factomdp2psignificant = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_significant_Summary",
		Help: "Summary of calls to  factomd_p2p_significant",
	})
	factomdp2plogfatal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_logfatal_Summary",
		Help: "Summary of calls to  factomd_p2p_logfatal",
	})
	factomdp2plogerror = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_logerror_Summary",
		Help: "Summary of calls to  factomd_p2p_logerror",
	})
	factomdp2pnote = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_note_Summary",
		Help: "Summary of calls to  factomd_p2p_note",
	})
	factomdp2pdebug = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_debug_Summary",
		Help: "Summary of calls to  factomd_p2p_debug",
	})
	factomdp2pverbose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_verbose_Summary",
		Help: "Summary of calls to  factomd_p2p_verbose",
	})
	factomdp2plog = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_p2p_log_Summary",
		Help: "Summary of calls to  factomd_p2p_log",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdp2pConnectionParcelJSONByte)
	prometheus.MustRegister(factomdp2pConnectionParcelJSONString)
	prometheus.MustRegister(factomdp2pConnectionParcelString)
	prometheus.MustRegister(factomdp2pConnectionCommandJSONByte)
	prometheus.MustRegister(factomdp2pConnectionCommandJSONString)
	prometheus.MustRegister(factomdp2pConnectionCommandString)
	prometheus.MustRegister(factomdp2pConnectionInitWithConn)
	prometheus.MustRegister(factomdp2pConnectionInit)
	prometheus.MustRegister(factomdp2pConnectionIsOutGoing)
	prometheus.MustRegister(factomdp2pConnectionIsOnline)
	prometheus.MustRegister(factomdp2pConnectionStatusString)
	prometheus.MustRegister(factomdp2pConnectionIsPersistent)
	prometheus.MustRegister(factomdp2pConnectionNotes)
	prometheus.MustRegister(factomdp2pConnectioncommonInit)
	prometheus.MustRegister(factomdp2pConnectionStart)
	prometheus.MustRegister(factomdp2pConnectionrunLoop)
	prometheus.MustRegister(factomdp2pConnectionsetNotes)
	prometheus.MustRegister(factomdp2pConnectiondialLoop)
	prometheus.MustRegister(factomdp2pConnectiondial)
	prometheus.MustRegister(factomdp2pConnectiongoOnline)
	prometheus.MustRegister(factomdp2pConnectiongoOffline)
	prometheus.MustRegister(factomdp2pConnectiongoShutdown)
	prometheus.MustRegister(factomdp2pConnectionprocessSends)
	prometheus.MustRegister(factomdp2pConnectionhandleCommand)
	prometheus.MustRegister(factomdp2pConnectionsendParcel)
	prometheus.MustRegister(factomdp2pConnectionprocessReceives)
	prometheus.MustRegister(factomdp2pConnectionhandleNetErrors)
	prometheus.MustRegister(factomdp2pConnectionhandleParcel)
	prometheus.MustRegister(factomdp2pConnectionparcelValidity)
	prometheus.MustRegister(factomdp2pConnectionhandleParcelTypes)
	prometheus.MustRegister(factomdp2pConnectionpingPeer)
	prometheus.MustRegister(factomdp2pConnectionupdatePeer)
	prometheus.MustRegister(factomdp2pConnectionupdateStats)
	prometheus.MustRegister(factomdp2pConnectionConnectionState)
	prometheus.MustRegister(factomdp2pConnectionconnectionStatusReport)
	prometheus.MustRegister(factomdp2pCommandAdjustPeerQualityJSONByte)
	prometheus.MustRegister(factomdp2pCommandAdjustPeerQualityJSONString)
	prometheus.MustRegister(factomdp2pCommandAdjustPeerQualityString)
	prometheus.MustRegister(factomdp2pCommandBanJSONByte)
	prometheus.MustRegister(factomdp2pCommandBanJSONString)
	prometheus.MustRegister(factomdp2pCommandBanString)
	prometheus.MustRegister(factomdp2pCommandDisconnectJSONByte)
	prometheus.MustRegister(factomdp2pCommandDisconnectJSONString)
	prometheus.MustRegister(factomdp2pCommandDisconnectString)
	prometheus.MustRegister(factomdp2pCommandChangeLoggingJSONByte)
	prometheus.MustRegister(factomdp2pCommandChangeLoggingJSONString)
	prometheus.MustRegister(factomdp2pCommandChangeLoggingString)
	prometheus.MustRegister(factomdp2pControllerInit)
	prometheus.MustRegister(factomdp2pControllerStartNetwork)
	prometheus.MustRegister(factomdp2pControllerDialSpecialPeersString)
	prometheus.MustRegister(factomdp2pControllerStartLogging)
	prometheus.MustRegister(factomdp2pControllerStopLogging)
	prometheus.MustRegister(factomdp2pControllerChangeLogLevel)
	prometheus.MustRegister(factomdp2pControllerDialPeer)
	prometheus.MustRegister(factomdp2pControllerAddPeer)
	prometheus.MustRegister(factomdp2pControllerNetworkStop)
	prometheus.MustRegister(factomdp2pControllerAdjustPeerQuality)
	prometheus.MustRegister(factomdp2pControllerBan)
	prometheus.MustRegister(factomdp2pControllerDisconnect)
	prometheus.MustRegister(factomdp2pControllerGetNumberConnections)
	prometheus.MustRegister(factomdp2pControllerlisten)
	prometheus.MustRegister(factomdp2pControlleracceptLoop)
	prometheus.MustRegister(factomdp2pControllerrunloop)
	prometheus.MustRegister(factomdp2pControllerroute)
	prometheus.MustRegister(factomdp2pControllerdoDirectedSend)
	prometheus.MustRegister(factomdp2pControllerhandleParcelReceive)
	prometheus.MustRegister(factomdp2pControllerhandleConnectionCommand)
	prometheus.MustRegister(factomdp2pControllerhandleCommand)
	prometheus.MustRegister(factomdp2pControllerapplicationPeerUpdate)
	prometheus.MustRegister(factomdp2pControllermanagePeers)
	prometheus.MustRegister(factomdp2pControllerupdateConnectionCounts)
	prometheus.MustRegister(factomdp2pControllerupdateConnectionAddressMap)
	prometheus.MustRegister(factomdp2pControllerweAreNotAlreadyConnectedTo)
	prometheus.MustRegister(factomdp2pControllerfillOutgoingSlots)
	prometheus.MustRegister(factomdp2pControllerupdateMetrics)
	prometheus.MustRegister(factomdp2pControllershutdown)
	prometheus.MustRegister(factomdp2pControllernetworkStatusReport)
	prometheus.MustRegister(factomdp2pDiscoveryInit)
	prometheus.MustRegister(factomdp2pDiscoveryupdatePeer)
	prometheus.MustRegister(factomdp2pDiscoverygetPeer)
	prometheus.MustRegister(factomdp2pDiscoveryisPeerPresent)
	prometheus.MustRegister(factomdp2pDiscoveryLoadPeers)
	prometheus.MustRegister(factomdp2pDiscoverySavePeers)
	prometheus.MustRegister(factomdp2pDiscoveryLearnPeers)
	prometheus.MustRegister(factomdp2pDiscoveryupdatePeerSource)
	prometheus.MustRegister(factomdp2pDiscoveryfilterPeersFromOtherNetworks)
	prometheus.MustRegister(factomdp2pDiscoveryfilterForUniqueIPAdresses)
	prometheus.MustRegister(factomdp2pDiscoveryGetOutgoingPeers)
	prometheus.MustRegister(factomdp2pDiscoverySharePeers)
	prometheus.MustRegister(factomdp2pDiscoverygetPeerSelection)
	prometheus.MustRegister(factomdp2pDiscoveryDiscoverPeersFromSeed)
	prometheus.MustRegister(factomdp2pDiscoveryPrintPeers)
	prometheus.MustRegister(factomdp2pNewParcel)
	prometheus.MustRegister(factomdp2pParcelsForPayload)
	prometheus.MustRegister(factomdp2pReassembleParcel)
	prometheus.MustRegister(factomdp2pParcelHeaderInit)
	prometheus.MustRegister(factomdp2pParcelInit)
	prometheus.MustRegister(factomdp2pParcelUpdateHeader)
	prometheus.MustRegister(factomdp2pParcelTrace)
	prometheus.MustRegister(factomdp2pParcelHeaderPrint)
	prometheus.MustRegister(factomdp2pParcelPrint)
	prometheus.MustRegister(factomdp2pParcelMessageType)
	prometheus.MustRegister(factomdp2pParcelPrintMessageType)
	prometheus.MustRegister(factomdp2pParcelString)
	prometheus.MustRegister(factomdp2pPartsAssemblerInit)
	prometheus.MustRegister(factomdp2pPartsAssemblerhandlePart)
	prometheus.MustRegister(factomdp2pvalidateParcelPart)
	prometheus.MustRegister(factomdp2pPartsAssemblercleanupOldPartialMessages)
	prometheus.MustRegister(factomdp2pcreateNewPartialMessage)
	prometheus.MustRegister(factomdp2ptryReassemblingMessage)
	prometheus.MustRegister(factomdp2pPeerInit)
	prometheus.MustRegister(factomdp2pPeergeneratePeerHash)
	prometheus.MustRegister(factomdp2pPeerAddressPort)
	prometheus.MustRegister(factomdp2pPeerPeerIdent)
	prometheus.MustRegister(factomdp2pPeerPeerFixedIdent)
	prometheus.MustRegister(factomdp2pPeerLocationFromAddress)
	prometheus.MustRegister(factomdp2pPeermerit)
	prometheus.MustRegister(factomdp2pPeerdemerit)
	prometheus.MustRegister(factomdp2pPeerQualitySortLen)
	prometheus.MustRegister(factomdp2pPeerQualitySortSwap)
	prometheus.MustRegister(factomdp2pPeerQualitySortLess)
	prometheus.MustRegister(factomdp2pPeerDistanceSortLen)
	prometheus.MustRegister(factomdp2pPeerDistanceSortSwap)
	prometheus.MustRegister(factomdp2pPeerDistanceSortLess)
	prometheus.MustRegister(factomdp2pBlockFreeChannelSend)
	prometheus.MustRegister(factomdp2pNetworkIDString)
	prometheus.MustRegister(factomdp2pdot)
	prometheus.MustRegister(factomdp2psilence)
	prometheus.MustRegister(factomdp2psignificant)
	prometheus.MustRegister(factomdp2plogfatal)
	prometheus.MustRegister(factomdp2plogerror)
	prometheus.MustRegister(factomdp2pnote)
	prometheus.MustRegister(factomdp2pdebug)
	prometheus.MustRegister(factomdp2pverbose)
	prometheus.MustRegister(factomdp2plog)
}
