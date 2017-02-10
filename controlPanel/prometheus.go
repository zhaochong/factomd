package controlPanel

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdcontrolPanelNewAllConnectionTotals = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_NewAllConnectionTotals_Summary",
		Help: "Summary of calls to  factomd_controlPanel_NewAllConnectionTotals",
	})
	factomdcontrolPanelNewConnectionsMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_NewConnectionsMap_Summary",
		Help: "Summary of calls to  factomd_controlPanel_NewConnectionsMap",
	})
	factomdcontrolPanelConnectionsMapTallyTotals = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_TallyTotals_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_TallyTotals",
	})
	factomdcontrolPanelConnectionsMapUpdateConnections = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_UpdateConnections_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_UpdateConnections",
	})
	factomdcontrolPanelhashPeerAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_hashPeerAddress_Summary",
		Help: "Summary of calls to  factomd_controlPanel_hashPeerAddress",
	})
	factomdcontrolPanelConnectionsMapAddConnection = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_AddConnection_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_AddConnection",
	})
	factomdcontrolPanelConnectionsMapRemoveConnection = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_RemoveConnection_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_RemoveConnection",
	})
	factomdcontrolPanelConnectionsMapConnect = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_Connect_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_Connect",
	})
	factomdcontrolPanelConnectionsMapGetConnection = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_GetConnection_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_GetConnection",
	})
	factomdcontrolPanelConnectionsMapGetConnectedCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_GetConnectedCopy_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_GetConnectedCopy",
	})
	factomdcontrolPanelConnectionsMapGetDisconnectedCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_GetDisconnectedCopy_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_GetDisconnectedCopy",
	})
	factomdcontrolPanelConnectionsMapDisconnect = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_Disconnect_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_Disconnect",
	})
	factomdcontrolPanelConnectionsMapCleanDisconnected = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_CleanDisconnected_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_CleanDisconnected",
	})
	factomdcontrolPanelConnectionInfoArrayLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionInfoArray_Len_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionInfoArray_Len",
	})
	factomdcontrolPanelConnectionInfoArrayLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionInfoArray_Less_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionInfoArray_Less",
	})
	factomdcontrolPanelConnectionInfoArraySwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionInfoArray_Swap_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionInfoArray_Swap",
	})
	factomdcontrolPanelConnectionsMapSortedConnections = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ConnectionsMap_SortedConnections_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ConnectionsMap_SortedConnections",
	})
	factomdcontrolPanelFormatDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_FormatDuration_Summary",
		Help: "Summary of calls to  factomd_controlPanel_FormatDuration",
	})
	factomdcontrolPanelmanageConnections = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_manageConnections_Summary",
		Help: "Summary of calls to  factomd_controlPanel_manageConnections",
	})
	factomdcontrolPaneldirectoryExists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_directoryExists_Summary",
		Help: "Summary of calls to  factomd_controlPanel_directoryExists",
	})
	factomdcontrolPanelDisplayStateDrain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_DisplayStateDrain_Summary",
		Help: "Summary of calls to  factomd_controlPanel_DisplayStateDrain",
	})
	factomdcontrolPanelServeControlPanel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_ServeControlPanel_Summary",
		Help: "Summary of calls to  factomd_controlPanel_ServeControlPanel",
	})
	factomdcontrolPanelnoStaticFilesFoundHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_noStaticFilesFoundHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_noStaticFilesFoundHandler",
	})
	factomdcontrolPanelstatic = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_static_Summary",
		Help: "Summary of calls to  factomd_controlPanel_static",
	})
	factomdcontrolPanelindexHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_indexHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_indexHandler",
	})
	factomdcontrolPanelpostHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_postHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_postHandler",
	})
	factomdcontrolPanelsearchHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_searchHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_searchHandler",
	})
	factomdcontrolPanelfactomdBatchHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_factomdBatchHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_factomdBatchHandler",
	})
	factomdcontrolPanelfactomdHandler = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_factomdHandler_Summary",
		Help: "Summary of calls to  factomd_controlPanel_factomdHandler",
	})
	factomdcontrolPanelrequestData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_requestData_Summary",
		Help: "Summary of calls to  factomd_controlPanel_requestData",
	})
	factomdcontrolPanelfactomdQuery = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_factomdQuery_Summary",
		Help: "Summary of calls to  factomd_controlPanel_factomdQuery",
	})
	factomdcontrolPaneldisconnectPeer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_disconnectPeer_Summary",
		Help: "Summary of calls to  factomd_controlPanel_disconnectPeer",
	})
	factomdcontrolPanelgetPeers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getPeers_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getPeers",
	})
	factomdcontrolPanelgetPeetTotals = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getPeetTotals_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getPeetTotals",
	})
	factomdcontrolPanelLastDirectoryBlockTransactionsContainsEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_LastDirectoryBlockTransactions_ContainsEntry_Summary",
		Help: "Summary of calls to  factomd_controlPanel_LastDirectoryBlockTransactions_ContainsEntry",
	})
	factomdcontrolPanelLastDirectoryBlockTransactionsContainsTrans = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_LastDirectoryBlockTransactions_ContainsTrans_Summary",
		Help: "Summary of calls to  factomd_controlPanel_LastDirectoryBlockTransactions_ContainsTrans",
	})
	factomdcontrolPaneltoggleDCT = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_toggleDCT_Summary",
		Help: "Summary of calls to  factomd_controlPanel_toggleDCT",
	})
	factomdcontrolPanelgetRecentTransactions = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getRecentTransactions_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getRecentTransactions",
	})
	factomdcontrolPanelgetPastEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getPastEntries_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getPastEntries",
	})
	factomdcontrolPaneldoEvery = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_doEvery_Summary",
		Help: "Summary of calls to  factomd_controlPanel_doEvery",
	})
	factomdcontrolPanelcheckControlPanelPassword = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_checkControlPanelPassword_Summary",
		Help: "Summary of calls to  factomd_controlPanel_checkControlPanelPassword",
	})
	factomdcontrolPanelcheckAuthHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_checkAuthHeader_Summary",
		Help: "Summary of calls to  factomd_controlPanel_checkAuthHeader",
	})
	factomdcontrolPanelgetDataDumps = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getDataDumps_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getDataDumps",
	})
	factomdcontrolPanelSortedConnectionString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_SortedConnectionString_Summary",
		Help: "Summary of calls to  factomd_controlPanel_SortedConnectionString",
	})
	factomdcontrolPanelAllConnectionsString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_AllConnectionsString_Summary",
		Help: "Summary of calls to  factomd_controlPanel_AllConnectionsString",
	})
	factomdcontrolPanelnewSearchResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_newSearchResponse_Summary",
		Help: "Summary of calls to  factomd_controlPanel_newSearchResponse",
	})
	factomdcontrolPanelsearchDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_searchDB_Summary",
		Help: "Summary of calls to  factomd_controlPanel_searchDB",
	})
	factomdcontrolPanelhandleSearchResult = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_handleSearchResult_Summary",
		Help: "Summary of calls to  factomd_controlPanel_handleSearchResult",
	})
	factomdcontrolPanelgetEcTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getEcTransaction_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getEcTransaction",
	})
	factomdcontrolPanelgetFactTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getFactTransaction_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getFactTransaction",
	})
	factomdcontrolPanelgetFactoidAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getFactoidAck_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getFactoidAck",
	})
	factomdcontrolPanelgetEntryAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getEntryAck_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getEntryAck",
	})
	factomdcontrolPanelgetECblock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getECblock_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getECblock",
	})
	factomdcontrolPanelgetFblock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getFblock_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getFblock",
	})
	factomdcontrolPanelgetAblock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getAblock_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getAblock",
	})
	factomdcontrolPanelgetEblock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getEblock_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getEblock",
	})
	factomdcontrolPanelgetDblock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getDblock_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getDblock",
	})
	factomdcontrolPanelgetEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getEntry_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getEntry",
	})
	factomdcontrolPanelgetAllChainEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_getAllChainEntries_Summary",
		Help: "Summary of calls to  factomd_controlPanel_getAllChainEntries",
	})
	factomdcontrolPanelHeightToJsonStruct = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_controlPanel_HeightToJsonStruct_Summary",
		Help: "Summary of calls to  factomd_controlPanel_HeightToJsonStruct",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdcontrolPanelNewAllConnectionTotals)
	prometheus.MustRegister(factomdcontrolPanelNewConnectionsMap)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapTallyTotals)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapUpdateConnections)
	prometheus.MustRegister(factomdcontrolPanelhashPeerAddress)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapAddConnection)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapRemoveConnection)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapConnect)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapGetConnection)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapGetConnectedCopy)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapGetDisconnectedCopy)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapDisconnect)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapCleanDisconnected)
	prometheus.MustRegister(factomdcontrolPanelConnectionInfoArrayLen)
	prometheus.MustRegister(factomdcontrolPanelConnectionInfoArrayLess)
	prometheus.MustRegister(factomdcontrolPanelConnectionInfoArraySwap)
	prometheus.MustRegister(factomdcontrolPanelConnectionsMapSortedConnections)
	prometheus.MustRegister(factomdcontrolPanelFormatDuration)
	prometheus.MustRegister(factomdcontrolPanelmanageConnections)
	prometheus.MustRegister(factomdcontrolPaneldirectoryExists)
	prometheus.MustRegister(factomdcontrolPanelDisplayStateDrain)
	prometheus.MustRegister(factomdcontrolPanelServeControlPanel)
	prometheus.MustRegister(factomdcontrolPanelnoStaticFilesFoundHandler)
	prometheus.MustRegister(factomdcontrolPanelstatic)
	prometheus.MustRegister(factomdcontrolPanelindexHandler)
	prometheus.MustRegister(factomdcontrolPanelpostHandler)
	prometheus.MustRegister(factomdcontrolPanelsearchHandler)
	prometheus.MustRegister(factomdcontrolPanelfactomdBatchHandler)
	prometheus.MustRegister(factomdcontrolPanelfactomdHandler)
	prometheus.MustRegister(factomdcontrolPanelrequestData)
	prometheus.MustRegister(factomdcontrolPanelfactomdQuery)
	prometheus.MustRegister(factomdcontrolPaneldisconnectPeer)
	prometheus.MustRegister(factomdcontrolPanelgetPeers)
	prometheus.MustRegister(factomdcontrolPanelgetPeetTotals)
	prometheus.MustRegister(factomdcontrolPanelLastDirectoryBlockTransactionsContainsEntry)
	prometheus.MustRegister(factomdcontrolPanelLastDirectoryBlockTransactionsContainsTrans)
	prometheus.MustRegister(factomdcontrolPaneltoggleDCT)
	prometheus.MustRegister(factomdcontrolPanelgetRecentTransactions)
	prometheus.MustRegister(factomdcontrolPanelgetPastEntries)
	prometheus.MustRegister(factomdcontrolPaneldoEvery)
	prometheus.MustRegister(factomdcontrolPanelcheckControlPanelPassword)
	prometheus.MustRegister(factomdcontrolPanelcheckAuthHeader)
	prometheus.MustRegister(factomdcontrolPanelgetDataDumps)
	prometheus.MustRegister(factomdcontrolPanelSortedConnectionString)
	prometheus.MustRegister(factomdcontrolPanelAllConnectionsString)
	prometheus.MustRegister(factomdcontrolPanelnewSearchResponse)
	prometheus.MustRegister(factomdcontrolPanelsearchDB)
	prometheus.MustRegister(factomdcontrolPanelhandleSearchResult)
	prometheus.MustRegister(factomdcontrolPanelgetEcTransaction)
	prometheus.MustRegister(factomdcontrolPanelgetFactTransaction)
	prometheus.MustRegister(factomdcontrolPanelgetFactoidAck)
	prometheus.MustRegister(factomdcontrolPanelgetEntryAck)
	prometheus.MustRegister(factomdcontrolPanelgetECblock)
	prometheus.MustRegister(factomdcontrolPanelgetFblock)
	prometheus.MustRegister(factomdcontrolPanelgetAblock)
	prometheus.MustRegister(factomdcontrolPanelgetEblock)
	prometheus.MustRegister(factomdcontrolPanelgetDblock)
	prometheus.MustRegister(factomdcontrolPanelgetEntry)
	prometheus.MustRegister(factomdcontrolPanelgetAllChainEntries)
	prometheus.MustRegister(factomdcontrolPanelHeightToJsonStruct)
}
