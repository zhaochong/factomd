package testHelper

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdtestHelperNewPrivKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewPrivKeyString_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewPrivKeyString",
	})
	factomdtestHelperNewPrivKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewPrivKey_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewPrivKey",
	})
	factomdtestHelperNewFullPrivKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewFullPrivKey_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewFullPrivKey",
	})
	factomdtestHelperNewPrimitivesPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewPrimitivesPrivateKey_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewPrimitivesPrivateKey",
	})
	factomdtestHelperNewFactoidAddressStrings = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewFactoidAddressStrings_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewFactoidAddressStrings",
	})
	factomdtestHelperNewFactoidAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewFactoidAddress_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewFactoidAddress",
	})
	factomdtestHelperNewFactoidRCDAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewFactoidRCDAddressString_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewFactoidRCDAddressString",
	})
	factomdtestHelperNewFactoidRCDAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewFactoidRCDAddress_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewFactoidRCDAddress",
	})
	factomdtestHelperNewECAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewECAddress_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewECAddress",
	})
	factomdtestHelperNewECAddressPublicKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewECAddressPublicKeyString_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewECAddressPublicKeyString",
	})
	factomdtestHelperNewECAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewECAddressString_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewECAddressString",
	})
	factomdtestHelperPrivateKeyToEDPub = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_PrivateKeyToEDPub_Summary",
		Help: "Summary of calls to  factomd_testHelper_PrivateKeyToEDPub",
	})
	factomdtestHelperCreateTestDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestDirBlockInfo",
	})
	factomdtestHelperIntToByteSlice = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_IntToByteSlice_Summary",
		Help: "Summary of calls to  factomd_testHelper_IntToByteSlice",
	})
	factomdtestHelperCreateTestEntryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestEntryBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestEntryBlock",
	})
	factomdtestHelperCreateTestEntryBlockWithContentN = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestEntryBlockWithContentN_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestEntryBlockWithContentN",
	})
	factomdtestHelperCreateTestAnchorEntryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestAnchorEntryBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestAnchorEntryBlock",
	})
	factomdtestHelperGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_GetChainID_Summary",
		Help: "Summary of calls to  factomd_testHelper_GetChainID",
	})
	factomdtestHelperGetAnchorChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_GetAnchorChainID_Summary",
		Help: "Summary of calls to  factomd_testHelper_GetAnchorChainID",
	})
	factomdtestHelperCreateFirstTestEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateFirstTestEntry_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateFirstTestEntry",
	})
	factomdtestHelperCreateFirstAnchorEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateFirstAnchorEntry_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateFirstAnchorEntry",
	})
	factomdtestHelperCreateTestEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestEntry_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestEntry",
	})
	factomdtestHelperCreateTestAnchorEnry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestAnchorEnry_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestAnchorEnry",
	})
	factomdtestHelpercreateECEntriesfromBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_createECEntriesfromBlocks_Summary",
		Help: "Summary of calls to  factomd_testHelper_createECEntriesfromBlocks",
	})
	factomdtestHelperNewCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewCommitEntry_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewCommitEntry",
	})
	factomdtestHelperNewCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewCommitChain_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewCommitChain",
	})
	factomdtestHelperCreateTestEntryCreditBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestEntryCreditBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestEntryCreditBlock",
	})
	factomdtestHelperSignCommit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_SignCommit_Summary",
		Help: "Summary of calls to  factomd_testHelper_SignCommit",
	})
	factomdtestHelperCreateTestFactoidBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestFactoidBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestFactoidBlock",
	})
	factomdtestHelperSignFactoidTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_SignFactoidTransaction_Summary",
		Help: "Summary of calls to  factomd_testHelper_SignFactoidTransaction",
	})
	factomdtestHelperCreateTestFactoidBlockWithCoinbase = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestFactoidBlockWithCoinbase_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestFactoidBlockWithCoinbase",
	})
	factomdtestHelperNewRepeatingHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_NewRepeatingHash_Summary",
		Help: "Summary of calls to  factomd_testHelper_NewRepeatingHash",
	})
	factomdtestHelperMakeFEREntryWithHeightFromContent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_MakeFEREntryWithHeightFromContent_Summary",
		Help: "Summary of calls to  factomd_testHelper_MakeFEREntryWithHeightFromContent",
	})
	factomdtestHelperCreateAndPopulateTestStateForFER = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateAndPopulateTestStateForFER_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateAndPopulateTestStateForFER",
	})
	factomdtestHelperCreateAndPopulateTestDatabaseOverlayForFER = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateAndPopulateTestDatabaseOverlayForFER_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateAndPopulateTestDatabaseOverlayForFER",
	})
	factomdtestHelperCreateTestBlockSetForFER = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestBlockSetForFER_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestBlockSetForFER",
	})
	factomdtestHelperCreateTestEntryBlockForFER = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestEntryBlockForFER_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestEntryBlockForFER",
	})
	factomdtestHelperCreateTestFEREntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestFEREntry_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestFEREntry",
	})
	factomdtestHelperCreateEmptyTestState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateEmptyTestState_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateEmptyTestState",
	})
	factomdtestHelperCreateAndPopulateTestState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateAndPopulateTestState_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateAndPopulateTestState",
	})
	factomdtestHelperCreateTestDBStateList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestDBStateList_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestDBStateList",
	})
	factomdtestHelperCreateTestLogFileString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestLogFileString_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestLogFileString",
	})
	factomdtestHelperMakeSureAnchorValidationKeyIsPresent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_MakeSureAnchorValidationKeyIsPresent_Summary",
		Help: "Summary of calls to  factomd_testHelper_MakeSureAnchorValidationKeyIsPresent",
	})
	factomdtestHelperPopulateTestDatabaseOverlay = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_PopulateTestDatabaseOverlay_Summary",
		Help: "Summary of calls to  factomd_testHelper_PopulateTestDatabaseOverlay",
	})
	factomdtestHelperCreateAndPopulateTestDatabaseOverlay = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateAndPopulateTestDatabaseOverlay_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateAndPopulateTestDatabaseOverlay",
	})
	factomdtestHelpernewBlockSet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_newBlockSet_Summary",
		Help: "Summary of calls to  factomd_testHelper_newBlockSet",
	})
	factomdtestHelperCreateFullTestBlockSet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateFullTestBlockSet_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateFullTestBlockSet",
	})
	factomdtestHelperCreateTestBlockSet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestBlockSet_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestBlockSet",
	})
	factomdtestHelperCreateEmptyTestDatabaseOverlay = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateEmptyTestDatabaseOverlay_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateEmptyTestDatabaseOverlay",
	})
	factomdtestHelperCreateTestAdminBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestAdminBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestAdminBlock",
	})
	factomdtestHelperCreateTestAdminHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestAdminHeader_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestAdminHeader",
	})
	factomdtestHelperCreateTestDirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestDirectoryBlock_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestDirectoryBlock",
	})
	factomdtestHelperCreateTestDirectoryBlockHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateTestDirectoryBlockHeader_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateTestDirectoryBlockHeader",
	})
	factomdtestHelperGetRespMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_GetRespMap_Summary",
		Help: "Summary of calls to  factomd_testHelper_GetRespMap",
	})
	factomdtestHelperUnmarshalResp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_UnmarshalResp_Summary",
		Help: "Summary of calls to  factomd_testHelper_UnmarshalResp",
	})
	factomdtestHelperUnmarshalRespDirectly = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_UnmarshalRespDirectly_Summary",
		Help: "Summary of calls to  factomd_testHelper_UnmarshalRespDirectly",
	})
	factomdtestHelperGetRespText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_GetRespText_Summary",
		Help: "Summary of calls to  factomd_testHelper_GetRespText",
	})
	factomdtestHelperClearContextResponseWriter = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_ClearContextResponseWriter_Summary",
		Help: "Summary of calls to  factomd_testHelper_ClearContextResponseWriter",
	})
	factomdtestHelperCreateWebContext = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_CreateWebContext_Summary",
		Help: "Summary of calls to  factomd_testHelper_CreateWebContext",
	})
	factomdtestHelperTestResponseWriterHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_TestResponseWriter_Header_Summary",
		Help: "Summary of calls to  factomd_testHelper_TestResponseWriter_Header",
	})
	factomdtestHelperTestResponseWriterWriteHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_TestResponseWriter_WriteHeader_Summary",
		Help: "Summary of calls to  factomd_testHelper_TestResponseWriter_WriteHeader",
	})
	factomdtestHelperTestResponseWriterWrite = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_TestResponseWriter_Write_Summary",
		Help: "Summary of calls to  factomd_testHelper_TestResponseWriter_Write",
	})
	factomdtestHelperGetBody = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_testHelper_GetBody_Summary",
		Help: "Summary of calls to  factomd_testHelper_GetBody",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdtestHelperNewPrivKeyString)
	prometheus.MustRegister(factomdtestHelperNewPrivKey)
	prometheus.MustRegister(factomdtestHelperNewFullPrivKey)
	prometheus.MustRegister(factomdtestHelperNewPrimitivesPrivateKey)
	prometheus.MustRegister(factomdtestHelperNewFactoidAddressStrings)
	prometheus.MustRegister(factomdtestHelperNewFactoidAddress)
	prometheus.MustRegister(factomdtestHelperNewFactoidRCDAddressString)
	prometheus.MustRegister(factomdtestHelperNewFactoidRCDAddress)
	prometheus.MustRegister(factomdtestHelperNewECAddress)
	prometheus.MustRegister(factomdtestHelperNewECAddressPublicKeyString)
	prometheus.MustRegister(factomdtestHelperNewECAddressString)
	prometheus.MustRegister(factomdtestHelperPrivateKeyToEDPub)
	prometheus.MustRegister(factomdtestHelperCreateTestDirBlockInfo)
	prometheus.MustRegister(factomdtestHelperIntToByteSlice)
	prometheus.MustRegister(factomdtestHelperCreateTestEntryBlock)
	prometheus.MustRegister(factomdtestHelperCreateTestEntryBlockWithContentN)
	prometheus.MustRegister(factomdtestHelperCreateTestAnchorEntryBlock)
	prometheus.MustRegister(factomdtestHelperGetChainID)
	prometheus.MustRegister(factomdtestHelperGetAnchorChainID)
	prometheus.MustRegister(factomdtestHelperCreateFirstTestEntry)
	prometheus.MustRegister(factomdtestHelperCreateFirstAnchorEntry)
	prometheus.MustRegister(factomdtestHelperCreateTestEntry)
	prometheus.MustRegister(factomdtestHelperCreateTestAnchorEnry)
	prometheus.MustRegister(factomdtestHelpercreateECEntriesfromBlocks)
	prometheus.MustRegister(factomdtestHelperNewCommitEntry)
	prometheus.MustRegister(factomdtestHelperNewCommitChain)
	prometheus.MustRegister(factomdtestHelperCreateTestEntryCreditBlock)
	prometheus.MustRegister(factomdtestHelperSignCommit)
	prometheus.MustRegister(factomdtestHelperCreateTestFactoidBlock)
	prometheus.MustRegister(factomdtestHelperSignFactoidTransaction)
	prometheus.MustRegister(factomdtestHelperCreateTestFactoidBlockWithCoinbase)
	prometheus.MustRegister(factomdtestHelperNewRepeatingHash)
	prometheus.MustRegister(factomdtestHelperMakeFEREntryWithHeightFromContent)
	prometheus.MustRegister(factomdtestHelperCreateAndPopulateTestStateForFER)
	prometheus.MustRegister(factomdtestHelperCreateAndPopulateTestDatabaseOverlayForFER)
	prometheus.MustRegister(factomdtestHelperCreateTestBlockSetForFER)
	prometheus.MustRegister(factomdtestHelperCreateTestEntryBlockForFER)
	prometheus.MustRegister(factomdtestHelperCreateTestFEREntry)
	prometheus.MustRegister(factomdtestHelperCreateEmptyTestState)
	prometheus.MustRegister(factomdtestHelperCreateAndPopulateTestState)
	prometheus.MustRegister(factomdtestHelperCreateTestDBStateList)
	prometheus.MustRegister(factomdtestHelperCreateTestLogFileString)
	prometheus.MustRegister(factomdtestHelperMakeSureAnchorValidationKeyIsPresent)
	prometheus.MustRegister(factomdtestHelperPopulateTestDatabaseOverlay)
	prometheus.MustRegister(factomdtestHelperCreateAndPopulateTestDatabaseOverlay)
	prometheus.MustRegister(factomdtestHelpernewBlockSet)
	prometheus.MustRegister(factomdtestHelperCreateFullTestBlockSet)
	prometheus.MustRegister(factomdtestHelperCreateTestBlockSet)
	prometheus.MustRegister(factomdtestHelperCreateEmptyTestDatabaseOverlay)
	prometheus.MustRegister(factomdtestHelperCreateTestAdminBlock)
	prometheus.MustRegister(factomdtestHelperCreateTestAdminHeader)
	prometheus.MustRegister(factomdtestHelperCreateTestDirectoryBlock)
	prometheus.MustRegister(factomdtestHelperCreateTestDirectoryBlockHeader)
	prometheus.MustRegister(factomdtestHelperGetRespMap)
	prometheus.MustRegister(factomdtestHelperUnmarshalResp)
	prometheus.MustRegister(factomdtestHelperUnmarshalRespDirectly)
	prometheus.MustRegister(factomdtestHelperGetRespText)
	prometheus.MustRegister(factomdtestHelperClearContextResponseWriter)
	prometheus.MustRegister(factomdtestHelperCreateWebContext)
	prometheus.MustRegister(factomdtestHelperTestResponseWriterHeader)
	prometheus.MustRegister(factomdtestHelperTestResponseWriterWriteHeader)
	prometheus.MustRegister(factomdtestHelperTestResponseWriterWrite)
	prometheus.MustRegister(factomdtestHelperGetBody)
}
