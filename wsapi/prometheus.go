package wsapi

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdwsapiHandleV2FactoidACK = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2FactoidACK_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2FactoidACK",
	})
	factomdwsapiHandleV2EntryACK = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2EntryACK_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2EntryACK",
	})
	factomdwsapiDecodeTransactionToHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_DecodeTransactionToHashes_Summary",
		Help: "Summary of calls to  factomd_wsapi_DecodeTransactionToHashes",
	})
	factomdwsapiNewParseError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewParseError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewParseError",
	})
	factomdwsapiNewInvalidRequestError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidRequestError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidRequestError",
	})
	factomdwsapiNewMethodNotFoundError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewMethodNotFoundError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewMethodNotFoundError",
	})
	factomdwsapiNewInvalidParamsError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidParamsError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidParamsError",
	})
	factomdwsapiNewInternalError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInternalError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInternalError",
	})
	factomdwsapiNewCustomInternalError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewCustomInternalError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewCustomInternalError",
	})
	factomdwsapiNewCustomInvalidParamsError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewCustomInvalidParamsError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewCustomInvalidParamsError",
	})
	factomdwsapiNewInvalidAddressError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidAddressError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidAddressError",
	})
	factomdwsapiNewUnableToDecodeTransactionError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewUnableToDecodeTransactionError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewUnableToDecodeTransactionError",
	})
	factomdwsapiNewInvalidTransactionError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidTransactionError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidTransactionError",
	})
	factomdwsapiNewInvalidHashError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidHashError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidHashError",
	})
	factomdwsapiNewInvalidEntryError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidEntryError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidEntryError",
	})
	factomdwsapiNewInvalidCommitChainError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidCommitChainError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidCommitChainError",
	})
	factomdwsapiNewInvalidCommitEntryError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidCommitEntryError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidCommitEntryError",
	})
	factomdwsapiNewInvalidDataPassedError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInvalidDataPassedError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInvalidDataPassedError",
	})
	factomdwsapiNewInternalDatabaseError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewInternalDatabaseError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewInternalDatabaseError",
	})
	factomdwsapiNewBlockNotFoundError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewBlockNotFoundError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewBlockNotFoundError",
	})
	factomdwsapiNewEntryNotFoundError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewEntryNotFoundError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewEntryNotFoundError",
	})
	factomdwsapiNewObjectNotFoundError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewObjectNotFoundError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewObjectNotFoundError",
	})
	factomdwsapiNewMissingChainHeadError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewMissingChainHeadError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewMissingChainHeadError",
	})
	factomdwsapiNewReceiptError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_NewReceiptError_Summary",
		Help: "Summary of calls to  factomd_wsapi_NewReceiptError",
	})
	factomdwsapiInitLogs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_InitLogs_Summary",
		Help: "Summary of calls to  factomd_wsapi_InitLogs",
	})
	factomdwsapiStart = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_Start_Summary",
		Help: "Summary of calls to  factomd_wsapi_Start",
	})
	factomdwsapiSetState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_SetState_Summary",
		Help: "Summary of calls to  factomd_wsapi_SetState",
	})
	factomdwsapiStop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_Stop_Summary",
		Help: "Summary of calls to  factomd_wsapi_Stop",
	})
	factomdwsapihandleV1Error = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_handleV1Error_Summary",
		Help: "Summary of calls to  factomd_wsapi_handleV1Error",
	})
	factomdwsapireturnV1 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_returnV1_Summary",
		Help: "Summary of calls to  factomd_wsapi_returnV1",
	})
	factomdwsapiHandleDBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleDBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleDBlockByHeight",
	})
	factomdwsapiHandleECBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleECBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleECBlockByHeight",
	})
	factomdwsapiHandleFBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleFBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleFBlockByHeight",
	})
	factomdwsapiHandleABlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleABlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleABlockByHeight",
	})
	factomdwsapiHandleCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleCommitChain_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleCommitChain",
	})
	factomdwsapiHandleRevealChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleRevealChain_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleRevealChain",
	})
	factomdwsapiHandleCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleCommitEntry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleCommitEntry",
	})
	factomdwsapiHandleRevealEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleRevealEntry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleRevealEntry",
	})
	factomdwsapiHandleDirectoryBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleDirectoryBlockHead_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleDirectoryBlockHead",
	})
	factomdwsapiHandleGetRaw = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleGetRaw_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleGetRaw",
	})
	factomdwsapiHandleGetReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleGetReceipt_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleGetReceipt",
	})
	factomdwsapiHandleDirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleDirectoryBlock_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleDirectoryBlock",
	})
	factomdwsapiHandleDirectoryBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleDirectoryBlockHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleDirectoryBlockHeight",
	})
	factomdwsapiHandleEntryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleEntryBlock_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleEntryBlock",
	})
	factomdwsapiHandleEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleEntry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleEntry",
	})
	factomdwsapiHandleChainHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleChainHead_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleChainHead",
	})
	factomdwsapiHandleEntryCreditBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleEntryCreditBalance_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleEntryCreditBalance",
	})
	factomdwsapiHandleGetFee = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleGetFee_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleGetFee",
	})
	factomdwsapiHandleFactoidSubmit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleFactoidSubmit_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleFactoidSubmit",
	})
	factomdwsapiHandleFactoidBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleFactoidBalance_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleFactoidBalance",
	})
	factomdwsapiHandleProperties = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleProperties_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleProperties",
	})
	factomdwsapiHandleHeights = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleHeights_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleHeights",
	})
	factomdwsapireturnMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_returnMsg_Summary",
		Help: "Summary of calls to  factomd_wsapi_returnMsg",
	})
	factomdwsapireturnV1Msg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_returnV1Msg_Summary",
		Help: "Summary of calls to  factomd_wsapi_returnV1Msg",
	})
	factomdwsapihttpBasicAuth = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_httpBasicAuth_Summary",
		Help: "Summary of calls to  factomd_wsapi_httpBasicAuth",
	})
	factomdwsapicheckAuthHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_checkAuthHeader_Summary",
		Help: "Summary of calls to  factomd_wsapi_checkAuthHeader",
	})
	factomdwsapicheckHttpPasswordOkV1 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_checkHttpPasswordOkV1_Summary",
		Help: "Summary of calls to  factomd_wsapi_checkHttpPasswordOkV1",
	})
	factomdwsapifileExists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_fileExists_Summary",
		Help: "Summary of calls to  factomd_wsapi_fileExists",
	})
	factomdwsapigenCertPair = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_genCertPair_Summary",
		Help: "Summary of calls to  factomd_wsapi_genCertPair",
	})
	factomdwsapiDBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_DBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_wsapi_DBlock_JSONString",
	})
	factomdwsapiEBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_EBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_wsapi_EBlock_JSONString",
	})
	factomdwsapiEntryStructJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_EntryStruct_JSONString_Summary",
		Help: "Summary of calls to  factomd_wsapi_EntryStruct_JSONString",
	})
	factomdwsapiCHeadJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_CHead_JSONString_Summary",
		Help: "Summary of calls to  factomd_wsapi_CHead_JSONString",
	})
	factomdwsapiHandleV2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2",
	})
	factomdwsapiHandleV2Request = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Request_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Request",
	})
	factomdwsapiHandleV2DBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2DBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2DBlockByHeight",
	})
	factomdwsapiHandleV2ECBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2ECBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2ECBlockByHeight",
	})
	factomdwsapiHandleV2FBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2FBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2FBlockByHeight",
	})
	factomdwsapiHandleV2ABlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2ABlockByHeight_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2ABlockByHeight",
	})
	factomdwsapiHandleV2Error = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Error_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Error",
	})
	factomdwsapiMapToObject = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_MapToObject_Summary",
		Help: "Summary of calls to  factomd_wsapi_MapToObject",
	})
	factomdwsapiJStructMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_JStruct_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_wsapi_JStruct_MarshalJSON",
	})
	factomdwsapiJStructUnmarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_JStruct_UnmarshalJSON_Summary",
		Help: "Summary of calls to  factomd_wsapi_JStruct_UnmarshalJSON",
	})
	factomdwsapiObjectToJStruct = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_ObjectToJStruct_Summary",
		Help: "Summary of calls to  factomd_wsapi_ObjectToJStruct",
	})
	factomdwsapiHandleV2CommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2CommitChain_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2CommitChain",
	})
	factomdwsapiHandleV2RevealChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2RevealChain_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2RevealChain",
	})
	factomdwsapiHandleV2CommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2CommitEntry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2CommitEntry",
	})
	factomdwsapiHandleV2RevealEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2RevealEntry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2RevealEntry",
	})
	factomdwsapiHandleV2DirectoryBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2DirectoryBlockHead_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2DirectoryBlockHead",
	})
	factomdwsapiHandleV2RawData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2RawData_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2RawData",
	})
	factomdwsapiHandleV2Receipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Receipt_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Receipt",
	})
	factomdwsapiHandleV2DirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2DirectoryBlock_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2DirectoryBlock",
	})
	factomdwsapiHandleV2EntryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2EntryBlock_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2EntryBlock",
	})
	factomdwsapiHandleV2Entry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Entry_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Entry",
	})
	factomdwsapiHandleV2ChainHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2ChainHead_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2ChainHead",
	})
	factomdwsapiHandleV2EntryCreditBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2EntryCreditBalance_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2EntryCreditBalance",
	})
	factomdwsapiHandleV2EntryCreditRate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2EntryCreditRate_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2EntryCreditRate",
	})
	factomdwsapiHandleV2FactoidSubmit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2FactoidSubmit_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2FactoidSubmit",
	})
	factomdwsapiHandleV2FactoidBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2FactoidBalance_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2FactoidBalance",
	})
	factomdwsapiHandleV2Heights = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Heights_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Heights",
	})
	factomdwsapiHandleV2GetPendingEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2GetPendingEntries_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2GetPendingEntries",
	})
	factomdwsapiHandleV2GetPendingTransactions = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2GetPendingTransactions_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2GetPendingTransactions",
	})
	factomdwsapiHandleV2Properties = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2Properties_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2Properties",
	})
	factomdwsapiHandleV2SendRawMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2SendRawMessage_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2SendRawMessage",
	})
	factomdwsapiHandleV2GetTranasction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_HandleV2GetTranasction_Summary",
		Help: "Summary of calls to  factomd_wsapi_HandleV2GetTranasction",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdwsapiHandleV2FactoidACK)
	prometheus.MustRegister(factomdwsapiHandleV2EntryACK)
	prometheus.MustRegister(factomdwsapiDecodeTransactionToHashes)
	prometheus.MustRegister(factomdwsapiNewParseError)
	prometheus.MustRegister(factomdwsapiNewInvalidRequestError)
	prometheus.MustRegister(factomdwsapiNewMethodNotFoundError)
	prometheus.MustRegister(factomdwsapiNewInvalidParamsError)
	prometheus.MustRegister(factomdwsapiNewInternalError)
	prometheus.MustRegister(factomdwsapiNewCustomInternalError)
	prometheus.MustRegister(factomdwsapiNewCustomInvalidParamsError)
	prometheus.MustRegister(factomdwsapiNewInvalidAddressError)
	prometheus.MustRegister(factomdwsapiNewUnableToDecodeTransactionError)
	prometheus.MustRegister(factomdwsapiNewInvalidTransactionError)
	prometheus.MustRegister(factomdwsapiNewInvalidHashError)
	prometheus.MustRegister(factomdwsapiNewInvalidEntryError)
	prometheus.MustRegister(factomdwsapiNewInvalidCommitChainError)
	prometheus.MustRegister(factomdwsapiNewInvalidCommitEntryError)
	prometheus.MustRegister(factomdwsapiNewInvalidDataPassedError)
	prometheus.MustRegister(factomdwsapiNewInternalDatabaseError)
	prometheus.MustRegister(factomdwsapiNewBlockNotFoundError)
	prometheus.MustRegister(factomdwsapiNewEntryNotFoundError)
	prometheus.MustRegister(factomdwsapiNewObjectNotFoundError)
	prometheus.MustRegister(factomdwsapiNewMissingChainHeadError)
	prometheus.MustRegister(factomdwsapiNewReceiptError)
	prometheus.MustRegister(factomdwsapiInitLogs)
	prometheus.MustRegister(factomdwsapiStart)
	prometheus.MustRegister(factomdwsapiSetState)
	prometheus.MustRegister(factomdwsapiStop)
	prometheus.MustRegister(factomdwsapihandleV1Error)
	prometheus.MustRegister(factomdwsapireturnV1)
	prometheus.MustRegister(factomdwsapiHandleDBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleECBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleFBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleABlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleCommitChain)
	prometheus.MustRegister(factomdwsapiHandleRevealChain)
	prometheus.MustRegister(factomdwsapiHandleCommitEntry)
	prometheus.MustRegister(factomdwsapiHandleRevealEntry)
	prometheus.MustRegister(factomdwsapiHandleDirectoryBlockHead)
	prometheus.MustRegister(factomdwsapiHandleGetRaw)
	prometheus.MustRegister(factomdwsapiHandleGetReceipt)
	prometheus.MustRegister(factomdwsapiHandleDirectoryBlock)
	prometheus.MustRegister(factomdwsapiHandleDirectoryBlockHeight)
	prometheus.MustRegister(factomdwsapiHandleEntryBlock)
	prometheus.MustRegister(factomdwsapiHandleEntry)
	prometheus.MustRegister(factomdwsapiHandleChainHead)
	prometheus.MustRegister(factomdwsapiHandleEntryCreditBalance)
	prometheus.MustRegister(factomdwsapiHandleGetFee)
	prometheus.MustRegister(factomdwsapiHandleFactoidSubmit)
	prometheus.MustRegister(factomdwsapiHandleFactoidBalance)
	prometheus.MustRegister(factomdwsapiHandleProperties)
	prometheus.MustRegister(factomdwsapiHandleHeights)
	prometheus.MustRegister(factomdwsapireturnMsg)
	prometheus.MustRegister(factomdwsapireturnV1Msg)
	prometheus.MustRegister(factomdwsapihttpBasicAuth)
	prometheus.MustRegister(factomdwsapicheckAuthHeader)
	prometheus.MustRegister(factomdwsapicheckHttpPasswordOkV1)
	prometheus.MustRegister(factomdwsapifileExists)
	prometheus.MustRegister(factomdwsapigenCertPair)
	prometheus.MustRegister(factomdwsapiDBlockJSONString)
	prometheus.MustRegister(factomdwsapiEBlockJSONString)
	prometheus.MustRegister(factomdwsapiEntryStructJSONString)
	prometheus.MustRegister(factomdwsapiCHeadJSONString)
	prometheus.MustRegister(factomdwsapiHandleV2)
	prometheus.MustRegister(factomdwsapiHandleV2Request)
	prometheus.MustRegister(factomdwsapiHandleV2DBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleV2ECBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleV2FBlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleV2ABlockByHeight)
	prometheus.MustRegister(factomdwsapiHandleV2Error)
	prometheus.MustRegister(factomdwsapiMapToObject)
	prometheus.MustRegister(factomdwsapiJStructMarshalJSON)
	prometheus.MustRegister(factomdwsapiJStructUnmarshalJSON)
	prometheus.MustRegister(factomdwsapiObjectToJStruct)
	prometheus.MustRegister(factomdwsapiHandleV2CommitChain)
	prometheus.MustRegister(factomdwsapiHandleV2RevealChain)
	prometheus.MustRegister(factomdwsapiHandleV2CommitEntry)
	prometheus.MustRegister(factomdwsapiHandleV2RevealEntry)
	prometheus.MustRegister(factomdwsapiHandleV2DirectoryBlockHead)
	prometheus.MustRegister(factomdwsapiHandleV2RawData)
	prometheus.MustRegister(factomdwsapiHandleV2Receipt)
	prometheus.MustRegister(factomdwsapiHandleV2DirectoryBlock)
	prometheus.MustRegister(factomdwsapiHandleV2EntryBlock)
	prometheus.MustRegister(factomdwsapiHandleV2Entry)
	prometheus.MustRegister(factomdwsapiHandleV2ChainHead)
	prometheus.MustRegister(factomdwsapiHandleV2EntryCreditBalance)
	prometheus.MustRegister(factomdwsapiHandleV2EntryCreditRate)
	prometheus.MustRegister(factomdwsapiHandleV2FactoidSubmit)
	prometheus.MustRegister(factomdwsapiHandleV2FactoidBalance)
	prometheus.MustRegister(factomdwsapiHandleV2Heights)
	prometheus.MustRegister(factomdwsapiHandleV2GetPendingEntries)
	prometheus.MustRegister(factomdwsapiHandleV2GetPendingTransactions)
	prometheus.MustRegister(factomdwsapiHandleV2Properties)
	prometheus.MustRegister(factomdwsapiHandleV2SendRawMessage)
	prometheus.MustRegister(factomdwsapiHandleV2GetTranasction)
}
