package factoid

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdfactoidRandomAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RandomAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_RandomAddress",
	})
	factomdfactoidAddressCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Address_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_Address_CustomMarshalText",
	})
	factomdfactoidNewAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_NewAddress",
	})
	factomdfactoidCreateAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_CreateAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_CreateAddress",
	})
	factomdfactoidUpdateAmount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_UpdateAmount_Summary",
		Help: "Summary of calls to  factomd_factoid_UpdateAmount",
	})
	factomdfactoidGetCoinbase = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_GetCoinbase_Summary",
		Help: "Summary of calls to  factomd_factoid_GetCoinbase",
	})
	factomdfactoidPublicKeyStringToECAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyStringToECAddressString_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyStringToECAddressString",
	})
	factomdfactoidPublicKeyStringToECAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyStringToECAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyStringToECAddress",
	})
	factomdfactoidPublicKeyToECAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyToECAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyToECAddress",
	})
	factomdfactoidPublicKeyStringToFactoidAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyStringToFactoidAddressString_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyStringToFactoidAddressString",
	})
	factomdfactoidPublicKeyToFactoidAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyToFactoidAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyToFactoidAddress",
	})
	factomdfactoidPublicKeyStringToFactoidAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyStringToFactoidAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyStringToFactoidAddress",
	})
	factomdfactoidPublicKeyStringToFactoidRCDAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PublicKeyStringToFactoidRCDAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_PublicKeyStringToFactoidRCDAddress",
	})
	factomdfactoidHumanReadiblePrivateKeyStringToEverythingString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_HumanReadiblePrivateKeyStringToEverythingString_Summary",
		Help: "Summary of calls to  factomd_factoid_HumanReadiblePrivateKeyStringToEverythingString",
	})
	factomdfactoidPrivateKeyStringToEverythingString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_PrivateKeyStringToEverythingString_Summary",
		Help: "Summary of calls to  factomd_factoid_PrivateKeyStringToEverythingString",
	})
	factomdfactoidFBlockGetEntryHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetEntryHashes_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetEntryHashes",
	})
	factomdfactoidFBlockGetTransactionByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetTransactionByHash_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetTransactionByHash",
	})
	factomdfactoidFBlockGetEntrySigHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetEntrySigHashes_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetEntrySigHashes",
	})
	factomdfactoidFBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_New_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_New",
	})
	factomdfactoidFBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_DatabasePrimaryIndex",
	})
	factomdfactoidFBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_DatabaseSecondaryIndex",
	})
	factomdfactoidFBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetDatabaseHeight",
	})
	factomdfactoidFBlockGetCoinbaseTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetCoinbaseTimestamp_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetCoinbaseTimestamp",
	})
	factomdfactoidFBlockEndOfPeriod = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_EndOfPeriod_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_EndOfPeriod",
	})
	factomdfactoidFBlockGetTransactions = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetTransactions_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetTransactions",
	})
	factomdfactoidFBlockGetNewInstance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetNewInstance_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetNewInstance",
	})
	factomdfactoidFBlockGetEndOfPeriod = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetEndOfPeriod_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetEndOfPeriod",
	})
	factomdfactoidFBlockMarshalTrans = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_MarshalTrans_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_MarshalTrans",
	})
	factomdfactoidFBlockMarshalHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_MarshalHeader_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_MarshalHeader",
	})
	factomdfactoidFBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_MarshalBinary",
	})
	factomdfactoidUnmarshalFBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_UnmarshalFBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_UnmarshalFBlock",
	})
	factomdfactoidFBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_UnmarshalBinaryData",
	})
	factomdfactoidFBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_UnmarshalBinary",
	})
	factomdfactoidFBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetChainID_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetChainID",
	})
	factomdfactoidFBlockGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetKeyMR",
	})
	factomdfactoidFBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetHash",
	})
	factomdfactoidFBlockGetLedgerKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetLedgerKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetLedgerKeyMR",
	})
	factomdfactoidFBlockGetLedgerMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetLedgerMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetLedgerMR",
	})
	factomdfactoidFBlockGetBodyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetBodyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetBodyMR",
	})
	factomdfactoidFBlockGetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetPrevKeyMR",
	})
	factomdfactoidFBlockSetPrevKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_SetPrevKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_SetPrevKeyMR",
	})
	factomdfactoidFBlockGetPrevLedgerKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetPrevLedgerKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetPrevLedgerKeyMR",
	})
	factomdfactoidFBlockSetPrevLedgerKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_SetPrevLedgerKeyMR_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_SetPrevLedgerKeyMR",
	})
	factomdfactoidFBlockCalculateHashes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_CalculateHashes_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_CalculateHashes",
	})
	factomdfactoidFBlockSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_SetDBHeight_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_SetDBHeight",
	})
	factomdfactoidFBlockGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetDBHeight",
	})
	factomdfactoidFBlockSetExchRate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_SetExchRate_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_SetExchRate",
	})
	factomdfactoidFBlockGetExchRate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_GetExchRate_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_GetExchRate",
	})
	factomdfactoidFBlockValidateTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_ValidateTransaction_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_ValidateTransaction",
	})
	factomdfactoidFBlockValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_Validate_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_Validate",
	})
	factomdfactoidFBlockAddCoinbase = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_AddCoinbase_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_AddCoinbase",
	})
	factomdfactoidFBlockAddTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_AddTransaction_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_AddTransaction",
	})
	factomdfactoidFBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_String_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_String",
	})
	factomdfactoidFBlockCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_CustomMarshalText",
	})
	factomdfactoidFBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_JSONByte",
	})
	factomdfactoidFBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_JSONString",
	})
	factomdfactoidFBlockMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FBlock_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_factoid_FBlock_MarshalJSON",
	})
	factomdfactoidNewFBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewFBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_NewFBlock",
	})
	factomdfactoidCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_CheckBlockPairIntegrity_Summary",
		Help: "Summary of calls to  factomd_factoid_CheckBlockPairIntegrity",
	})
	factomdfactoidGetGenesisFBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_GetGenesisFBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_GetGenesisFBlock",
	})
	factomdfactoidUnmarshalBinaryAuth = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_UnmarshalBinaryAuth_Summary",
		Help: "Summary of calls to  factomd_factoid_UnmarshalBinaryAuth",
	})
	factomdfactoidNewRCD_1 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewRCD_1_Summary",
		Help: "Summary of calls to  factomd_factoid_NewRCD_1",
	})
	factomdfactoidNewRCD_2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewRCD_2_Summary",
		Help: "Summary of calls to  factomd_factoid_NewRCD_2",
	})
	factomdfactoidCreateRCD = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_CreateRCD_Summary",
		Help: "Summary of calls to  factomd_factoid_CreateRCD",
	})
	factomdfactoidRCD_1IsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_IsSameAs",
	})
	factomdfactoidRCD_1UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_UnmarshalBinary",
	})
	factomdfactoidRCD_1JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_JSONByte",
	})
	factomdfactoidRCD_1JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_JSONString",
	})
	factomdfactoidRCD_1String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_String_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_String",
	})
	factomdfactoidRCD_1MarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_MarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_MarshalText",
	})
	factomdfactoidRCD_1CheckSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_CheckSig_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_CheckSig",
	})
	factomdfactoidRCD_1Clone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_Clone_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_Clone",
	})
	factomdfactoidRCD_1GetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_GetAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_GetAddress",
	})
	factomdfactoidRCD_1GetPublicKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_GetPublicKey_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_GetPublicKey",
	})
	factomdfactoidRCD_1NumberOfSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_NumberOfSignatures_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_NumberOfSignatures",
	})
	factomdfactoidRCD_1UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_UnmarshalBinaryData",
	})
	factomdfactoidRCD_1MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_MarshalBinary",
	})
	factomdfactoidRCD_1CustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_1_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_1_CustomMarshalText",
	})
	factomdfactoidRCD_2GetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_GetAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_GetAddress",
	})
	factomdfactoidRCD_2NumberOfSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_NumberOfSignatures_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_NumberOfSignatures",
	})
	factomdfactoidRCD_2IsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_IsSameAs",
	})
	factomdfactoidRCD_2UnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_UnmarshalBinary",
	})
	factomdfactoidRCD_2CheckSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_CheckSig_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_CheckSig",
	})
	factomdfactoidRCD_2JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_JSONByte",
	})
	factomdfactoidRCD_2JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_JSONString",
	})
	factomdfactoidRCD_2String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_String_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_String",
	})
	factomdfactoidRCD_2Clone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_Clone_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_Clone",
	})
	factomdfactoidRCD_2UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_UnmarshalBinaryData",
	})
	factomdfactoidRCD_2MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_MarshalBinary",
	})
	factomdfactoidRCD_2CustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RCD_2_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_RCD_2_CustomMarshalText",
	})
	factomdfactoidFactoidSignatureIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_IsSameAs",
	})
	factomdfactoidFactoidSignatureVerify = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_Verify_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_Verify",
	})
	factomdfactoidFactoidSignatureBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_Bytes_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_Bytes",
	})
	factomdfactoidFactoidSignatureGetKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_GetKey_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_GetKey",
	})
	factomdfactoidFactoidSignatureMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_MarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_MarshalText",
	})
	factomdfactoidFactoidSignatureJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_JSONByte",
	})
	factomdfactoidFactoidSignatureJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_JSONString",
	})
	factomdfactoidFactoidSignatureString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_String_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_String",
	})
	factomdfactoidFactoidSignatureSetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_SetSignature_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_SetSignature",
	})
	factomdfactoidFactoidSignatureGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_GetSignature_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_GetSignature",
	})
	factomdfactoidFactoidSignatureMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_MarshalBinary",
	})
	factomdfactoidFactoidSignatureCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_CustomMarshalText",
	})
	factomdfactoidFactoidSignatureUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_UnmarshalBinaryData",
	})
	factomdfactoidFactoidSignatureUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_FactoidSignature_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_FactoidSignature_UnmarshalBinary",
	})
	factomdfactoidNewED25519Signature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewED25519Signature_Summary",
		Help: "Summary of calls to  factomd_factoid_NewED25519Signature",
	})
	factomdfactoidSignatureBlockIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_IsSameAs",
	})
	factomdfactoidSignatureBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_UnmarshalBinary",
	})
	factomdfactoidSignatureBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_JSONByte",
	})
	factomdfactoidSignatureBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_JSONString",
	})
	factomdfactoidSignatureBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_String_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_String",
	})
	factomdfactoidSignatureBlockAddSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_AddSignature_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_AddSignature",
	})
	factomdfactoidSignatureBlockGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_GetSignature_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_GetSignature",
	})
	factomdfactoidSignatureBlockGetSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_GetSignatures_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_GetSignatures",
	})
	factomdfactoidSignatureBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_MarshalBinary",
	})
	factomdfactoidSignatureBlockCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_CustomMarshalText",
	})
	factomdfactoidSignatureBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_SignatureBlock_UnmarshalBinaryData",
	})
	factomdfactoidNewSingleSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewSingleSignatureBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_NewSingleSignatureBlock",
	})
	factomdfactoidTransactionIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_IsSameAs",
	})
	factomdfactoidTransactionNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_New_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_New",
	})
	factomdfactoidTransactionSetBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_SetBlockHeight_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_SetBlockHeight",
	})
	factomdfactoidTransactionGetBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetBlockHeight_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetBlockHeight",
	})
	factomdfactoidTransactionclearCaches = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_clearCaches_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_clearCaches",
	})
	factomdfactoidGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_GetVersion(__Summary",
		Help: "Summary of calls to  factomd_factoid_GetVersion(_",
	})
	factomdfactoidTransactionGetTxID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetTxID_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetTxID",
	})
	factomdfactoidTransactionGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetHash_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetHash",
	})
	factomdfactoidTransactionGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetFullHash_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetFullHash",
	})
	factomdfactoidTransactionGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSigHash_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetSigHash",
	})
	factomdfactoidTransactionString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_String_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_String",
	})
	factomdfactoidTransactionGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetTimestamp",
	})
	factomdfactoidTransactionSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_SetTimestamp_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_SetTimestamp",
	})
	factomdfactoidTransactionSetSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_SetSignatureBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_SetSignatureBlock",
	})
	factomdfactoidTransactionGetSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSignatureBlock_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetSignatureBlock",
	})
	factomdfactoidTransactionAddRCD = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddRCD_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_AddRCD",
	})
	factomdfactoidTransactionCalculateFee = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_CalculateFee_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_CalculateFee",
	})
	factomdfactoidValidateAmounts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_ValidateAmounts_Summary",
		Help: "Summary of calls to  factomd_factoid_ValidateAmounts",
	})
	factomdfactoidTransactionTotalInputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalInputs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_TotalInputs",
	})
	factomdfactoidTransactionTotalOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalOutputs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_TotalOutputs",
	})
	factomdfactoidTransactionTotalECs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalECs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_TotalECs",
	})
	factomdfactoidTransactionValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_Validate_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_Validate",
	})
	factomdfactoidTransactionValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_ValidateSignatures_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_ValidateSignatures",
	})
	factomdfactoidTransactionGetInputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetInputs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetInputs",
	})
	factomdfactoidTransactionGetOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetOutputs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetOutputs",
	})
	factomdfactoidTransactionGetECOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetECOutputs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetECOutputs",
	})
	factomdfactoidTransactionGetRCDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetRCDs_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetRCDs",
	})
	factomdfactoidTransactionGetSignatureBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSignatureBlocks_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetSignatureBlocks",
	})
	factomdfactoidTransactionGetInput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetInput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetInput",
	})
	factomdfactoidTransactionGetOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetOutput",
	})
	factomdfactoidTransactionGetECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetECOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetECOutput",
	})
	factomdfactoidTransactionGetRCD = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetRCD_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_GetRCD",
	})
	factomdfactoidTransactionUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_UnmarshalBinaryData",
	})
	factomdfactoidTransactionUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_UnmarshalBinary",
	})
	factomdfactoidTransactionMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_MarshalBinarySig_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_MarshalBinarySig",
	})
	factomdfactoidTransactionMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_MarshalBinary",
	})
	factomdfactoidTransactionAddInput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddInput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_AddInput",
	})
	factomdfactoidTransactionAddOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_AddOutput",
	})
	factomdfactoidTransactionAddECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddECOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_AddECOutput",
	})
	factomdfactoidTransactionCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_CustomMarshalText_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_CustomMarshalText",
	})
	factomdfactoidTransactionAddAuthorization = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddAuthorization_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_AddAuthorization",
	})
	factomdfactoidTransactionJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_JSONByte",
	})
	factomdfactoidTransactionJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_JSONString",
	})
	factomdfactoidTransactionHasUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_HasUserAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_Transaction_HasUserAddress",
	})
	factomdfactoidRandomTransAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_RandomTransAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_RandomTransAddress",
	})
	factomdfactoidTransAddressSetUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetUserAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_SetUserAddress",
	})
	factomdfactoidTransAddressGetUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetUserAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_GetUserAddress",
	})
	factomdfactoidTransAddressUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_UnmarshalBinary",
	})
	factomdfactoidTransAddressJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_JSONByte_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_JSONByte",
	})
	factomdfactoidTransAddressJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_JSONString_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_JSONString",
	})
	factomdfactoidTransAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_String_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_String",
	})
	factomdfactoidTransAddressIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_IsSameAs",
	})
	factomdfactoidTransAddressUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_UnmarshalBinaryData",
	})
	factomdfactoidTransAddressMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_MarshalBinary",
	})
	factomdfactoidTransAddressGetName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetName_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_GetName",
	})
	factomdfactoidTransAddressGetAmount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetAmount_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_GetAmount",
	})
	factomdfactoidTransAddressSetAmount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetAmount_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_SetAmount",
	})
	factomdfactoidTransAddressGetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_GetAddress",
	})
	factomdfactoidTransAddressSetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_SetAddress",
	})
	factomdfactoidTransAddressCustomMarshalTextAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextAll_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalTextAll",
	})
	factomdfactoidTransAddressCustomMarshalText2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalText2_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalText2",
	})
	factomdfactoidTransAddressCustomMarshalTextEC2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextEC2_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalTextEC2",
	})
	factomdfactoidTransAddressCustomMarshalTextInput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextInput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalTextInput",
	})
	factomdfactoidTransAddressStringInput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_StringInput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_StringInput",
	})
	factomdfactoidTransAddressCustomMarshalTextOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalTextOutput",
	})
	factomdfactoidTransAddressStringOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_StringOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_StringOutput",
	})
	factomdfactoidTransAddressCustomMarshalTextECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextECOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_CustomMarshalTextECOutput",
	})
	factomdfactoidTransAddressStringECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_StringECOutput_Summary",
		Help: "Summary of calls to  factomd_factoid_TransAddress_StringECOutput",
	})
	factomdfactoidNewOutECAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewOutECAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_NewOutECAddress",
	})
	factomdfactoidNewOutAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewOutAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_NewOutAddress",
	})
	factomdfactoidNewInAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewInAddress_Summary",
		Help: "Summary of calls to  factomd_factoid_NewInAddress",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdfactoidRandomAddress)
	prometheus.MustRegister(factomdfactoidAddressCustomMarshalText)
	prometheus.MustRegister(factomdfactoidNewAddress)
	prometheus.MustRegister(factomdfactoidCreateAddress)
	prometheus.MustRegister(factomdfactoidUpdateAmount)
	prometheus.MustRegister(factomdfactoidGetCoinbase)
	prometheus.MustRegister(factomdfactoidPublicKeyStringToECAddressString)
	prometheus.MustRegister(factomdfactoidPublicKeyStringToECAddress)
	prometheus.MustRegister(factomdfactoidPublicKeyToECAddress)
	prometheus.MustRegister(factomdfactoidPublicKeyStringToFactoidAddressString)
	prometheus.MustRegister(factomdfactoidPublicKeyToFactoidAddress)
	prometheus.MustRegister(factomdfactoidPublicKeyStringToFactoidAddress)
	prometheus.MustRegister(factomdfactoidPublicKeyStringToFactoidRCDAddress)
	prometheus.MustRegister(factomdfactoidHumanReadiblePrivateKeyStringToEverythingString)
	prometheus.MustRegister(factomdfactoidPrivateKeyStringToEverythingString)
	prometheus.MustRegister(factomdfactoidFBlockGetEntryHashes)
	prometheus.MustRegister(factomdfactoidFBlockGetTransactionByHash)
	prometheus.MustRegister(factomdfactoidFBlockGetEntrySigHashes)
	prometheus.MustRegister(factomdfactoidFBlockNew)
	prometheus.MustRegister(factomdfactoidFBlockDatabasePrimaryIndex)
	prometheus.MustRegister(factomdfactoidFBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(factomdfactoidFBlockGetDatabaseHeight)
	prometheus.MustRegister(factomdfactoidFBlockGetCoinbaseTimestamp)
	prometheus.MustRegister(factomdfactoidFBlockEndOfPeriod)
	prometheus.MustRegister(factomdfactoidFBlockGetTransactions)
	prometheus.MustRegister(factomdfactoidFBlockGetNewInstance)
	prometheus.MustRegister(factomdfactoidFBlockGetEndOfPeriod)
	prometheus.MustRegister(factomdfactoidFBlockMarshalTrans)
	prometheus.MustRegister(factomdfactoidFBlockMarshalHeader)
	prometheus.MustRegister(factomdfactoidFBlockMarshalBinary)
	prometheus.MustRegister(factomdfactoidUnmarshalFBlock)
	prometheus.MustRegister(factomdfactoidFBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidFBlockUnmarshalBinary)
	prometheus.MustRegister(factomdfactoidFBlockGetChainID)
	prometheus.MustRegister(factomdfactoidFBlockGetKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockGetHash)
	prometheus.MustRegister(factomdfactoidFBlockGetLedgerKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockGetLedgerMR)
	prometheus.MustRegister(factomdfactoidFBlockGetBodyMR)
	prometheus.MustRegister(factomdfactoidFBlockGetPrevKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockSetPrevKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockGetPrevLedgerKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockSetPrevLedgerKeyMR)
	prometheus.MustRegister(factomdfactoidFBlockCalculateHashes)
	prometheus.MustRegister(factomdfactoidFBlockSetDBHeight)
	prometheus.MustRegister(factomdfactoidFBlockGetDBHeight)
	prometheus.MustRegister(factomdfactoidFBlockSetExchRate)
	prometheus.MustRegister(factomdfactoidFBlockGetExchRate)
	prometheus.MustRegister(factomdfactoidFBlockValidateTransaction)
	prometheus.MustRegister(factomdfactoidFBlockValidate)
	prometheus.MustRegister(factomdfactoidFBlockAddCoinbase)
	prometheus.MustRegister(factomdfactoidFBlockAddTransaction)
	prometheus.MustRegister(factomdfactoidFBlockString)
	prometheus.MustRegister(factomdfactoidFBlockCustomMarshalText)
	prometheus.MustRegister(factomdfactoidFBlockJSONByte)
	prometheus.MustRegister(factomdfactoidFBlockJSONString)
	prometheus.MustRegister(factomdfactoidFBlockMarshalJSON)
	prometheus.MustRegister(factomdfactoidNewFBlock)
	prometheus.MustRegister(factomdfactoidCheckBlockPairIntegrity)
	prometheus.MustRegister(factomdfactoidGetGenesisFBlock)
	prometheus.MustRegister(factomdfactoidUnmarshalBinaryAuth)
	prometheus.MustRegister(factomdfactoidNewRCD_1)
	prometheus.MustRegister(factomdfactoidNewRCD_2)
	prometheus.MustRegister(factomdfactoidCreateRCD)
	prometheus.MustRegister(factomdfactoidRCD_1IsSameAs)
	prometheus.MustRegister(factomdfactoidRCD_1UnmarshalBinary)
	prometheus.MustRegister(factomdfactoidRCD_1JSONByte)
	prometheus.MustRegister(factomdfactoidRCD_1JSONString)
	prometheus.MustRegister(factomdfactoidRCD_1String)
	prometheus.MustRegister(factomdfactoidRCD_1MarshalText)
	prometheus.MustRegister(factomdfactoidRCD_1CheckSig)
	prometheus.MustRegister(factomdfactoidRCD_1Clone)
	prometheus.MustRegister(factomdfactoidRCD_1GetAddress)
	prometheus.MustRegister(factomdfactoidRCD_1GetPublicKey)
	prometheus.MustRegister(factomdfactoidRCD_1NumberOfSignatures)
	prometheus.MustRegister(factomdfactoidRCD_1UnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidRCD_1MarshalBinary)
	prometheus.MustRegister(factomdfactoidRCD_1CustomMarshalText)
	prometheus.MustRegister(factomdfactoidRCD_2GetAddress)
	prometheus.MustRegister(factomdfactoidRCD_2NumberOfSignatures)
	prometheus.MustRegister(factomdfactoidRCD_2IsSameAs)
	prometheus.MustRegister(factomdfactoidRCD_2UnmarshalBinary)
	prometheus.MustRegister(factomdfactoidRCD_2CheckSig)
	prometheus.MustRegister(factomdfactoidRCD_2JSONByte)
	prometheus.MustRegister(factomdfactoidRCD_2JSONString)
	prometheus.MustRegister(factomdfactoidRCD_2String)
	prometheus.MustRegister(factomdfactoidRCD_2Clone)
	prometheus.MustRegister(factomdfactoidRCD_2UnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidRCD_2MarshalBinary)
	prometheus.MustRegister(factomdfactoidRCD_2CustomMarshalText)
	prometheus.MustRegister(factomdfactoidFactoidSignatureIsSameAs)
	prometheus.MustRegister(factomdfactoidFactoidSignatureVerify)
	prometheus.MustRegister(factomdfactoidFactoidSignatureBytes)
	prometheus.MustRegister(factomdfactoidFactoidSignatureGetKey)
	prometheus.MustRegister(factomdfactoidFactoidSignatureMarshalText)
	prometheus.MustRegister(factomdfactoidFactoidSignatureJSONByte)
	prometheus.MustRegister(factomdfactoidFactoidSignatureJSONString)
	prometheus.MustRegister(factomdfactoidFactoidSignatureString)
	prometheus.MustRegister(factomdfactoidFactoidSignatureSetSignature)
	prometheus.MustRegister(factomdfactoidFactoidSignatureGetSignature)
	prometheus.MustRegister(factomdfactoidFactoidSignatureMarshalBinary)
	prometheus.MustRegister(factomdfactoidFactoidSignatureCustomMarshalText)
	prometheus.MustRegister(factomdfactoidFactoidSignatureUnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidFactoidSignatureUnmarshalBinary)
	prometheus.MustRegister(factomdfactoidNewED25519Signature)
	prometheus.MustRegister(factomdfactoidSignatureBlockIsSameAs)
	prometheus.MustRegister(factomdfactoidSignatureBlockUnmarshalBinary)
	prometheus.MustRegister(factomdfactoidSignatureBlockJSONByte)
	prometheus.MustRegister(factomdfactoidSignatureBlockJSONString)
	prometheus.MustRegister(factomdfactoidSignatureBlockString)
	prometheus.MustRegister(factomdfactoidSignatureBlockAddSignature)
	prometheus.MustRegister(factomdfactoidSignatureBlockGetSignature)
	prometheus.MustRegister(factomdfactoidSignatureBlockGetSignatures)
	prometheus.MustRegister(factomdfactoidSignatureBlockMarshalBinary)
	prometheus.MustRegister(factomdfactoidSignatureBlockCustomMarshalText)
	prometheus.MustRegister(factomdfactoidSignatureBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidNewSingleSignatureBlock)
	prometheus.MustRegister(factomdfactoidTransactionIsSameAs)
	prometheus.MustRegister(factomdfactoidTransactionNew)
	prometheus.MustRegister(factomdfactoidTransactionSetBlockHeight)
	prometheus.MustRegister(factomdfactoidTransactionGetBlockHeight)
	prometheus.MustRegister(factomdfactoidTransactionclearCaches)
	prometheus.MustRegister(factomdfactoidGetVersion)
	prometheus.MustRegister(factomdfactoidTransactionGetTxID)
	prometheus.MustRegister(factomdfactoidTransactionGetHash)
	prometheus.MustRegister(factomdfactoidTransactionGetFullHash)
	prometheus.MustRegister(factomdfactoidTransactionGetSigHash)
	prometheus.MustRegister(factomdfactoidTransactionString)
	prometheus.MustRegister(factomdfactoidTransactionGetTimestamp)
	prometheus.MustRegister(factomdfactoidTransactionSetTimestamp)
	prometheus.MustRegister(factomdfactoidTransactionSetSignatureBlock)
	prometheus.MustRegister(factomdfactoidTransactionGetSignatureBlock)
	prometheus.MustRegister(factomdfactoidTransactionAddRCD)
	prometheus.MustRegister(factomdfactoidTransactionCalculateFee)
	prometheus.MustRegister(factomdfactoidValidateAmounts)
	prometheus.MustRegister(factomdfactoidTransactionTotalInputs)
	prometheus.MustRegister(factomdfactoidTransactionTotalOutputs)
	prometheus.MustRegister(factomdfactoidTransactionTotalECs)
	prometheus.MustRegister(factomdfactoidTransactionValidate)
	prometheus.MustRegister(factomdfactoidTransactionValidateSignatures)
	prometheus.MustRegister(factomdfactoidTransactionGetInputs)
	prometheus.MustRegister(factomdfactoidTransactionGetOutputs)
	prometheus.MustRegister(factomdfactoidTransactionGetECOutputs)
	prometheus.MustRegister(factomdfactoidTransactionGetRCDs)
	prometheus.MustRegister(factomdfactoidTransactionGetSignatureBlocks)
	prometheus.MustRegister(factomdfactoidTransactionGetInput)
	prometheus.MustRegister(factomdfactoidTransactionGetOutput)
	prometheus.MustRegister(factomdfactoidTransactionGetECOutput)
	prometheus.MustRegister(factomdfactoidTransactionGetRCD)
	prometheus.MustRegister(factomdfactoidTransactionUnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidTransactionUnmarshalBinary)
	prometheus.MustRegister(factomdfactoidTransactionMarshalBinarySig)
	prometheus.MustRegister(factomdfactoidTransactionMarshalBinary)
	prometheus.MustRegister(factomdfactoidTransactionAddInput)
	prometheus.MustRegister(factomdfactoidTransactionAddOutput)
	prometheus.MustRegister(factomdfactoidTransactionAddECOutput)
	prometheus.MustRegister(factomdfactoidTransactionCustomMarshalText)
	prometheus.MustRegister(factomdfactoidTransactionAddAuthorization)
	prometheus.MustRegister(factomdfactoidTransactionJSONByte)
	prometheus.MustRegister(factomdfactoidTransactionJSONString)
	prometheus.MustRegister(factomdfactoidTransactionHasUserAddress)
	prometheus.MustRegister(factomdfactoidRandomTransAddress)
	prometheus.MustRegister(factomdfactoidTransAddressSetUserAddress)
	prometheus.MustRegister(factomdfactoidTransAddressGetUserAddress)
	prometheus.MustRegister(factomdfactoidTransAddressUnmarshalBinary)
	prometheus.MustRegister(factomdfactoidTransAddressJSONByte)
	prometheus.MustRegister(factomdfactoidTransAddressJSONString)
	prometheus.MustRegister(factomdfactoidTransAddressString)
	prometheus.MustRegister(factomdfactoidTransAddressIsSameAs)
	prometheus.MustRegister(factomdfactoidTransAddressUnmarshalBinaryData)
	prometheus.MustRegister(factomdfactoidTransAddressMarshalBinary)
	prometheus.MustRegister(factomdfactoidTransAddressGetName)
	prometheus.MustRegister(factomdfactoidTransAddressGetAmount)
	prometheus.MustRegister(factomdfactoidTransAddressSetAmount)
	prometheus.MustRegister(factomdfactoidTransAddressGetAddress)
	prometheus.MustRegister(factomdfactoidTransAddressSetAddress)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalTextAll)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalText2)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalTextEC2)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalTextInput)
	prometheus.MustRegister(factomdfactoidTransAddressStringInput)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalTextOutput)
	prometheus.MustRegister(factomdfactoidTransAddressStringOutput)
	prometheus.MustRegister(factomdfactoidTransAddressCustomMarshalTextECOutput)
	prometheus.MustRegister(factomdfactoidTransAddressStringECOutput)
	prometheus.MustRegister(factomdfactoidNewOutECAddress)
	prometheus.MustRegister(factomdfactoidNewOutAddress)
	prometheus.MustRegister(factomdfactoidNewInAddress)
}
