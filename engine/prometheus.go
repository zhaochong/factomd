package engine

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdengineMsgLoginit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_MsgLog_init_Summary",
		Help: "Summary of calls to  factomd_engine_MsgLog_init",
	})
	factomdengineMsgLogadd2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_MsgLog_add2_Summary",
		Help: "Summary of calls to  factomd_engine_MsgLog_add2",
	})
	factomdengineMsgLogPrtMsgs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_MsgLog_PrtMsgs_Summary",
		Help: "Summary of calls to  factomd_engine_MsgLog_PrtMsgs",
	})
	factomdengineNetStart = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_NetStart_Summary",
		Help: "Summary of calls to  factomd_engine_NetStart",
	})
	factomdenginemakeServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_makeServer_Summary",
		Help: "Summary of calls to  factomd_engine_makeServer",
	})
	factomdenginestartServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_startServers_Summary",
		Help: "Summary of calls to  factomd_engine_startServers",
	})
	factomdenginesetupFirstAuthority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_setupFirstAuthority_Summary",
		Help: "Summary of calls to  factomd_engine_setupFirstAuthority",
	})
	factomdenginenetworkHousekeeping = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_networkHousekeeping_Summary",
		Help: "Summary of calls to  factomd_engine_networkHousekeeping",
	})
	factomdengineNetworkProcessorNet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_NetworkProcessorNet_Summary",
		Help: "Summary of calls to  factomd_engine_NetworkProcessorNet",
	})
	factomdenginePeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_Peers_Summary",
		Help: "Summary of calls to  factomd_engine_Peers",
	})
	factomdengineNetworkOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_NetworkOutputs_Summary",
		Help: "Summary of calls to  factomd_engine_NetworkOutputs",
	})
	factomdengineInvalidOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_InvalidOutputs_Summary",
		Help: "Summary of calls to  factomd_engine_InvalidOutputs",
	})
	factomdenginemainInterruptHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_mainInterruptHandler_Summary",
		Help: "Summary of calls to  factomd_engine_mainInterruptHandler",
	})
	factomdengineAddInterruptHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_AddInterruptHandler_Summary",
		Help: "Summary of calls to  factomd_engine_AddInterruptHandler",
	})
	factomdengineSimPeerBytesOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_BytesOut_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_BytesOut",
	})
	factomdengineSimPeerBytesIn = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_BytesIn_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_BytesIn",
	})
	factomdengineWeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_Weight_Summary",
		Help: "Summary of calls to  factomd_engine_Weight",
	})
	factomdengineSimPeerEquals = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_Equals_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_Equals",
	})
	factomdengineSimPeerLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_Len_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_Len",
	})
	factomdengineSimPeerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_Init_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_Init",
	})
	factomdengineSimPeerGetNameFrom = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_GetNameFrom_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_GetNameFrom",
	})
	factomdengineSimPeerGetNameTo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_GetNameTo_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_GetNameTo",
	})
	factomdengineSimPeercomputeBandwidth = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_computeBandwidth_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_computeBandwidth",
	})
	factomdengineSimPeerSend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_Send_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_Send",
	})
	factomdengineSimPeerRecieve = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimPeer_Recieve_Summary",
		Help: "Summary of calls to  factomd_engine_SimPeer_Recieve",
	})
	factomdengineAddSimPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_AddSimPeer_Summary",
		Help: "Summary of calls to  factomd_engine_AddSimPeer",
	})
	factomdengineFactomd = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_Factomd_Summary",
		Help: "Summary of calls to  factomd_engine_Factomd",
	})
	factomdengineisCompilerVersionOK = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_isCompilerVersionOK_Summary",
		Help: "Summary of calls to  factomd_engine_isCompilerVersionOK",
	})
	factomdenginewaitToKill = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_waitToKill_Summary",
		Help: "Summary of calls to  factomd_engine_waitToKill",
	})
	factomdenginebringback = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_bringback_Summary",
		Help: "Summary of calls to  factomd_engine_bringback",
	})
	factomdengineofflineReport = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_offlineReport_Summary",
		Help: "Summary of calls to  factomd_engine_offlineReport",
	})
	factomdenginefaultTest = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_faultTest_Summary",
		Help: "Summary of calls to  factomd_engine_faultTest",
	})
	factomdengineLoadJournal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_LoadJournal_Summary",
		Help: "Summary of calls to  factomd_engine_LoadJournal",
	})
	factomdengineLoadJournalFromString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_LoadJournalFromString_Summary",
		Help: "Summary of calls to  factomd_engine_LoadJournalFromString",
	})
	factomdengineLoadJournalFromReader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_LoadJournalFromReader_Summary",
		Help: "Summary of calls to  factomd_engine_LoadJournalFromReader",
	})
	factomdenginefactomMessageJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_factomMessage_JSONByte_Summary",
		Help: "Summary of calls to  factomd_engine_factomMessage_JSONByte",
	})
	factomdenginefactomMessageJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_factomMessage_JSONString_Summary",
		Help: "Summary of calls to  factomd_engine_factomMessage_JSONString",
	})
	factomdenginefactomMessageString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_factomMessage_String_Summary",
		Help: "Summary of calls to  factomd_engine_factomMessage_String",
	})
	factomdengineP2PProxyWeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Weight_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Weight",
	})
	factomdengineP2PProxySetWeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_SetWeight_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_SetWeight",
	})
	factomdengineP2PProxyBytesOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_BytesOut_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_BytesOut",
	})
	factomdengineP2PProxyBytesIn = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_BytesIn_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_BytesIn",
	})
	factomdengineP2PProxyInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Init_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Init",
	})
	factomdengineP2PProxySetDebugMode = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_SetDebugMode_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_SetDebugMode",
	})
	factomdengineP2PProxyGetNameFrom = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_GetNameFrom_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_GetNameFrom",
	})
	factomdengineP2PProxyGetNameTo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_GetNameTo_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_GetNameTo",
	})
	factomdengineP2PProxySend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Send_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Send",
	})
	factomdengineP2PProxyRecieve = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Recieve_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Recieve",
	})
	factomdengineP2PProxyEquals = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Equals_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Equals",
	})
	factomdengineP2PProxyLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_Len_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_Len",
	})
	factomdengineP2PProxyStartProxy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_StartProxy_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_StartProxy",
	})
	factomdengineP2PProxystopProxy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_stopProxy_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_stopProxy",
	})
	factomdenginemessageLogJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_messageLog_JSONByte_Summary",
		Help: "Summary of calls to  factomd_engine_messageLog_JSONByte",
	})
	factomdenginemessageLogJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_messageLog_JSONString_Summary",
		Help: "Summary of calls to  factomd_engine_messageLog_JSONString",
	})
	factomdenginemessageLogString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_messageLog_String_Summary",
		Help: "Summary of calls to  factomd_engine_messageLog_String",
	})
	factomdengineP2PProxylogMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_logMessage_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_logMessage",
	})
	factomdengineP2PProxyManageLogging = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_ManageLogging_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_ManageLogging",
	})
	factomdengineP2PProxyManageOutChannel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_ManageOutChannel_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_ManageOutChannel",
	})
	factomdengineP2PProxyManageInChannel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_ManageInChannel_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_ManageInChannel",
	})
	factomdengineP2PProxytrace = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_trace_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_trace",
	})
	factomdengineP2PProxyPeriodicStatusReport = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_P2PProxy_PeriodicStatusReport_Summary",
		Help: "Summary of calls to  factomd_engine_P2PProxy_PeriodicStatusReport",
	})
	factomdengineprintSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_printSummary_Summary",
		Help: "Summary of calls to  factomd_engine_printSummary",
	})
	factomdenginesystemFaults = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_systemFaults_Summary",
		Help: "Summary of calls to  factomd_engine_systemFaults",
	})
	factomdenginefaultSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_faultSummary_Summary",
		Help: "Summary of calls to  factomd_engine_faultSummary",
	})
	factomdengineStartProfiler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_StartProfiler_Summary",
		Help: "Summary of calls to  factomd_engine_StartProfiler",
	})
	factomdengineSimControl = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_SimControl_Summary",
		Help: "Summary of calls to  factomd_engine_SimControl",
	})
	factomdenginereturnStatString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_returnStatString_Summary",
		Help: "Summary of calls to  factomd_engine_returnStatString",
	})
	factomdenginerotateWSAPI = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_rotateWSAPI_Summary",
		Help: "Summary of calls to  factomd_engine_rotateWSAPI",
	})
	factomdengineprintProcessList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_printProcessList_Summary",
		Help: "Summary of calls to  factomd_engine_printProcessList",
	})
	factomdengineprintMessages = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_printMessages_Summary",
		Help: "Summary of calls to  factomd_engine_printMessages",
	})
	factomdengineTimer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_Timer_Summary",
		Help: "Summary of calls to  factomd_engine_Timer",
	})
	factomdenginePrintBusy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_engine_PrintBusy_Summary",
		Help: "Summary of calls to  factomd_engine_PrintBusy",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdengineMsgLoginit)
	prometheus.MustRegister(factomdengineMsgLogadd2)
	prometheus.MustRegister(factomdengineMsgLogPrtMsgs)
	prometheus.MustRegister(factomdengineNetStart)
	prometheus.MustRegister(factomdenginemakeServer)
	prometheus.MustRegister(factomdenginestartServers)
	prometheus.MustRegister(factomdenginesetupFirstAuthority)
	prometheus.MustRegister(factomdenginenetworkHousekeeping)
	prometheus.MustRegister(factomdengineNetworkProcessorNet)
	prometheus.MustRegister(factomdenginePeers)
	prometheus.MustRegister(factomdengineNetworkOutputs)
	prometheus.MustRegister(factomdengineInvalidOutputs)
	prometheus.MustRegister(factomdenginemainInterruptHandler)
	prometheus.MustRegister(factomdengineAddInterruptHandler)
	prometheus.MustRegister(factomdengineSimPeerBytesOut)
	prometheus.MustRegister(factomdengineSimPeerBytesIn)
	prometheus.MustRegister(factomdengineWeight)
	prometheus.MustRegister(factomdengineSimPeerEquals)
	prometheus.MustRegister(factomdengineSimPeerLen)
	prometheus.MustRegister(factomdengineSimPeerInit)
	prometheus.MustRegister(factomdengineSimPeerGetNameFrom)
	prometheus.MustRegister(factomdengineSimPeerGetNameTo)
	prometheus.MustRegister(factomdengineSimPeercomputeBandwidth)
	prometheus.MustRegister(factomdengineSimPeerSend)
	prometheus.MustRegister(factomdengineSimPeerRecieve)
	prometheus.MustRegister(factomdengineAddSimPeer)
	prometheus.MustRegister(factomdengineFactomd)
	prometheus.MustRegister(factomdengineisCompilerVersionOK)
	prometheus.MustRegister(factomdenginewaitToKill)
	prometheus.MustRegister(factomdenginebringback)
	prometheus.MustRegister(factomdengineofflineReport)
	prometheus.MustRegister(factomdenginefaultTest)
	prometheus.MustRegister(factomdengineLoadJournal)
	prometheus.MustRegister(factomdengineLoadJournalFromString)
	prometheus.MustRegister(factomdengineLoadJournalFromReader)
	prometheus.MustRegister(factomdenginefactomMessageJSONByte)
	prometheus.MustRegister(factomdenginefactomMessageJSONString)
	prometheus.MustRegister(factomdenginefactomMessageString)
	prometheus.MustRegister(factomdengineP2PProxyWeight)
	prometheus.MustRegister(factomdengineP2PProxySetWeight)
	prometheus.MustRegister(factomdengineP2PProxyBytesOut)
	prometheus.MustRegister(factomdengineP2PProxyBytesIn)
	prometheus.MustRegister(factomdengineP2PProxyInit)
	prometheus.MustRegister(factomdengineP2PProxySetDebugMode)
	prometheus.MustRegister(factomdengineP2PProxyGetNameFrom)
	prometheus.MustRegister(factomdengineP2PProxyGetNameTo)
	prometheus.MustRegister(factomdengineP2PProxySend)
	prometheus.MustRegister(factomdengineP2PProxyRecieve)
	prometheus.MustRegister(factomdengineP2PProxyEquals)
	prometheus.MustRegister(factomdengineP2PProxyLen)
	prometheus.MustRegister(factomdengineP2PProxyStartProxy)
	prometheus.MustRegister(factomdengineP2PProxystopProxy)
	prometheus.MustRegister(factomdenginemessageLogJSONByte)
	prometheus.MustRegister(factomdenginemessageLogJSONString)
	prometheus.MustRegister(factomdenginemessageLogString)
	prometheus.MustRegister(factomdengineP2PProxylogMessage)
	prometheus.MustRegister(factomdengineP2PProxyManageLogging)
	prometheus.MustRegister(factomdengineP2PProxyManageOutChannel)
	prometheus.MustRegister(factomdengineP2PProxyManageInChannel)
	prometheus.MustRegister(factomdengineP2PProxytrace)
	prometheus.MustRegister(factomdengineP2PProxyPeriodicStatusReport)
	prometheus.MustRegister(factomdengineprintSummary)
	prometheus.MustRegister(factomdenginesystemFaults)
	prometheus.MustRegister(factomdenginefaultSummary)
	prometheus.MustRegister(factomdengineStartProfiler)
	prometheus.MustRegister(factomdengineSimControl)
	prometheus.MustRegister(factomdenginereturnStatString)
	prometheus.MustRegister(factomdenginerotateWSAPI)
	prometheus.MustRegister(factomdengineprintProcessList)
	prometheus.MustRegister(factomdengineprintMessages)
	prometheus.MustRegister(factomdengineTimer)
	prometheus.MustRegister(factomdenginePrintBusy)
}
