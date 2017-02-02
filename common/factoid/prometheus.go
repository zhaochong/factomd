// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import "github.com/prometheus/client_golang/prometheus"

var (
	//address  variables
	factoidErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_factoid_error_responses",
		Help: "Number of error responses from anchor objects",
	})
	factoidAddressCustomMarshalTextSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_CustomMarshalText_Summary",
		Help: "Summary of factoid.CustomMarshalText call",
	})
	factoidAddressNewAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_NewAddress_Summary",
		Help: "Summary of factoid.NewAddress call",
	})
	factoidAddressCreateAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_address_CreateAddress_Summary",
		Help: "Summary of factoid.address.CreateAddress call",
	})
	// coinbase
	factoidCoinbaseUpdateAmountSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_coinbase_UpdateAmount_Summary",
		Help: "Summary of factoid.coinbase.UpdateAmount call",
	})
	factoidCoinbaseGetCoinbaseSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_coinbase_GetCoinbase_Summary",
		Help: "Summary of factoid.coinbase.GetCoinbase call",
	})
	//conversions
	factoidConversionsPublicKeyStringToECAddressStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyStringToECAddressString_Summary",
		Help: "Summary of factoid.conversions.PublicKeyStringToECAddressString call",
	})
	factoidConversionsPublicKeyToECAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyToECAddress_Summary",
		Help: "Summary of factoid.conversions.PublicKeyToECAddress call",
	})
	factoidConversionsPublicKeyStringToFactoidAddressStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyStringToFactoidAddressString_Summary",
		Help: "Summary of factoid.conversions.PublicKeyStringToFactoidAddressString  call",
	})
	factoidConversionsPublicKeyToFactoidAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyToFactoidAddress_Summary",
		Help: "Summary of factoid.conversions.PublicKeyToFactoidAddress  call",
	})
	factoidConversionsPublicKeyStringToFactoidAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyStringToFactoidAddress_Summary",
		Help: "Summary of factoid.conversions.PublicKeyStringToFactoidAddress  call",
	})
	factoidConversionsPublicKeyStringToFactoidRCDAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PublicKeyStringToFactoidRCDAddress_Summary",
		Help: "Summary of factoid.conversions.PublicKeyStringToFactoidRCDAddress  call",
	})
	factoidConversionsHumanReadiblePrivateKeyStringToEverythingStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_HumanReadiblePrivateKeyStringToEverythingString_Summary",
		Help: "Summary of factoid.conversions.HumanReadiblePrivateKeyStringToEverythingString  call",
	})
	factoidConversionsPrivateKeyStringToEverythingStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_conversions_PrivateKeyStringToEverythingString_Summary",
		Help: "Summary of factoid.conversions.PrivateKeyStringToEverythingString  call",
	})
	// factoi
	factoidFactoidFactoidTx_VersionCheckSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_factoid_FactoidTx_VersionCheck_Summary",
		Help: "Summary of factoid.factoid.FactoidTx_VersionCheck  call",
	})
	factoidFactoidFactoidTx_LocktimeCheckSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_factoid_FactoidTx_LocktimeCheck_Summary",
		Help: "Summary of factoid.factoid.FactoidTx_LocktimeCheck  call",
	})
	factoidFactoidFactoidTx_RCDVersionCheckSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_factoid_FactoidTx_RCDVersionCheck_Summary",
		Help: "Summary of factoid.factoid.FactoidTx_RCDVersionCheck call",
	})
	factoidFactoidFactoidTx_RCDTypeCheckSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_factoid_FactoidTx_RCDTypeCheck_Summary",
		Help: "Summary of factoid.factoid.FactoidTx_RCDTypeCheck call",
	})
	//fblock
	factoidFBlockGetEntryHashesSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetEntryHashes_Summary",
		Help: "Summary of factoid.fblock.GetEntryHashes call",
	})
	factoidFBlockGetTransactionByHashSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetTransactionByHash_Summary",
		Help: "Summary of factoid.fblock.GetTransactionByHash call",
	})
	factoidFBlockGetEntrySigHashesSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetEntrySigHashes_Summary",
		Help: "Summary of factoid.fblock.GetEntrySigHashes call",
	})
	factoidFBlockNewSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_New_Summary",
		Help: "Summary of factoid.fblock.New call",
	})
	factoidFBlockDatabasePrimaryIndexSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_DatabasePrimaryIndex_Summary",
		Help: "Summary of factoid.fblock.DatabasePrimaryIndex call",
	})
	factoidFBlockDatabaseSecondaryIndexSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of factoid.fblock.DatabaseSecondaryIndex call",
	})
	factoidFBlockGetDatabaseHeightSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetDatabaseHeight_Summary",
		Help: "Summary of factoid.fblock.GetDatabaseHeight call",
	})
	factoidFBlockGetCoinbaseTimestampSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetCoinbaseTimestamp_Summary",
		Help: "Summary of factoid.fblock.GetCoinbaseTimestamp call",
	})
	factoidFBlockEndOfPeriodSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_EndOfPeriod_Summary",
		Help: "Summary of factoid.fblock.EndOfPeriod call",
	})
	factoidFBlockGetTransactionsSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetTransactions_Summary",
		Help: "Summary of factoid.fblock.GetTransactions call",
	})
	factoidFBlockGetNewInstanceSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetNewInstance_Summary",
		Help: "Summary of factoid.fblock.GetNewInstance call",
	})
	factoidFBlockGetEndOfPeriodSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetEndOfPeriod_Summary",
		Help: "Summary of factoid.fblock.GetEndOfPeriod call",
	})
	factoidFBlockMarshalTransSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_MarshalTrans_Summary",
		Help: "Summary of factoid.fblock.MarshalTrans call",
	})
	factoidFBlockMarshalHeaderSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_MarshalHeader_Summary",
		Help: "Summary of factoid.fblock.MarshalHeader call",
	})
	factoidFBlockMarshalBinarySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_MarshalBinary_Summary",
		Help: "Summary of factoid.fblock.MarshalBinary call",
	})
	factoidFBlockUnmarshalFBlockSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_UnmarshalFBlock_Summary",
		Help: "Summary of factoid.fblock.UnmarshalFBlock call",
	})
	factoidFBlockUnmarshalBinaryDataSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.fblock.UnmarshalBinaryData call",
	})
	factoidFBlockUnmarshalBinarySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_UnmarshalBinary_Summary",
		Help: "Summary of factoid.fblock.UnmarshalBinary call",
	})
	factoidFBlockIsEqualSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_IsEqual_Summary",
		Help: "Summary of factoid.fblock.IsEqual call",
	})
	factoidFBlockGetChainIDSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetChainID_Summary",
		Help: "Summary of factoid.fblock.GetChainID call",
	})
	factoidFBlockGetKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetKeyMR_Summary",
		Help: "Summary of factoid.fblock.GetKeyMR call",
	})
	factoidFBlockGetHashSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetHash_Summary",
		Help: "Summary of factoid.fblock.GetHash call",
	})
	factoidFBlockGetLedgerKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetLedgerKeyMR_Summary",
		Help: "Summary of factoid.fblock.GetLedgerKeyMR call",
	})
	factoidFBlockGetLedgerMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetLedgerMR_Summary",
		Help: "Summary of factoid.fblock.GetLedgerMR call",
	})
	factoidFBlockGetBodyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetBodyMR_Summary",
		Help: "Summary of factoid.fblock.GetBodyMR call",
	})
	factoidFBlockGetPrevKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetPrevKeyMR_Summary",
		Help: "Summary of factoid.fblock.GetPrevKeyMR call",
	})
	factoidFBlockSetPrevKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_SetPrevKeyMR_Summary",
		Help: "Summary of factoid.fblock.SetPrevKeyMR call",
	})
	factoidFBlockGetPrevLedgerKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetPrevLedgerKeyMR_Summary",
		Help: "Summary of factoid.fblock.GetPrevLedgerKeyMR call",
	})
	factoidFBlockSetPrevLedgerKeyMRSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_SetPrevLedgerKeyMR_Summary",
		Help: "Summary of factoid.fblock.SetPrevLedgerKeyMR call",
	})
	factoidFBlockCalculateHashesSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_CalculateHashes_Summary",
		Help: "Summary of factoid.fblock.CalculateHashes call",
	})
	factoidFBlockSetDBHeightSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_SetDBHeight_Summary",
		Help: "Summary of factoid.fblock.SetDBHeight call",
	})
	factoidFBlockGetDBHeightSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetDBHeight_Summary",
		Help: "Summary of factoid.fblock.GetDBHeight call",
	})
	factoidFBlockSetExchRateSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_SetExchRate_Summary",
		Help: "Summary of factoid.fblock.SetExchRate call",
	})
	factoidFBlockGetExchRateSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_GetExchRate_Summary",
		Help: "Summary of factoid.fblock.GetExchRate call",
	})
	factoidFBlockValidateTransactionSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_ValidateTransaction_Summary",
		Help: "Summary of factoid.fblock.ValidateTransactiont call",
	})
	factoidFBlockValidateSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_Validate_Summary",
		Help: "Summary of factoid.fblock.Validate call",
	})
	factoidFBlockAddCoinbaseSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_AddCoinbase_Summary",
		Help: "Summary of factoid.fblock.AddCoinbase call",
	})
	factoidFBlockAddTransactionSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_AddTransaction_Summary",
		Help: "Summary of factoid.fblock.AddTransaction call",
	})
	factoidFBlockStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_String_Summary",
		Help: "Summary of factoid.fblock.String call",
	})
	factoidFBlockCustomMarshalTextSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_CustomMarshalText_Summary",
		Help: "Summary of factoid.fblock.CustomMarshalText call",
	})
	factoidFBlockJSONByteSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_JSONByte_Summary",
		Help: "Summary of factoid.fblock.JSONByte call",
	})
	factoidFBlockJSONStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_JSONString_Summary",
		Help: "Summary of factoid.fblock.JSONString call",
	})
	factoidFBlockJSONBufferSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_JSONBuffer_Summary",
		Help: "Summary of factoid.fblock.JSONBuffer call",
	})
	factoidFBlockMarshalJSONSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_MarshalJSON_Summary",
		Help: "Summary of factoid.fblock.MarshalJSON call",
	})
	factoidFBlockNewFBlockSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_NewFBlock_Summary",
		Help: "Summary of factoid.fblock.NewFBlock call",
	})
	factoidFBlockCheckBlockPairIntegritySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_fblock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of factoid.fblock.CheckBlockPairIntegrity call",
	})
	//genesisBlockNe
	factoidgenesisBlockNewGetGenesisFBlockSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_genesisBlockNew_GetGenesisFBlock_Summary",
		Help: "Summary of factoid.genesisBlockNew.GetGenesisFBlock call",
	})
	//inaddres
	factoidInAddressStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_InAddress_String_Summary",
		Help: "Summary of factoid.InAddress.String call",
	})
	factoidInAddressCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_InAddress_CustomMarshalText_Summary",
		Help: "Summary of factoid.InAddress.CustomMarshalText call",
	})
	factoidInAddressNewInAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_InAddress_NewInAddress_Summary",
		Help: "Summary of factoid.InAddress.NewInAddress call",
	})
	// outaddress
	factoidOutAddressStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutAddress_String_Summary",
		Help: "Summary of factoid.OutAddress.String call",
	})
	factoidOutAddressCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutAddress_CustomMarshalText_Summary",
		Help: "Summary of factoid.OutAddress.CustomMarshalText call",
	})
	factoidOutAddressNewOutAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutAddress_NewOutAddress_Summary",
		Help: "Summary of factoid.OutAddress.NewInAddress call",
	})
	// outecaddress
	factoidOutECAddressStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutECAddress_String_Summary",
		Help: "Summary of factoid.OutECAddress.String call",
	})
	factoidOutECAddressCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutCAddress_CustomMarshalText_Summary",
		Help: "Summary of factoid.OutECAddress.CustomMarshalText call",
	})
	factoidOutECAddressNewOutAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_OutECAddress_NewOutECAddress_Summary",
		Help: "Summary of factoid.OutECAddress.NewInAddress call",
	})
	//rcd.go variables
	factoidrcdUnmarshalBinaryAuthSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd_UnmarshalBinaryAuth_Summary",
		Help: "Summary of factoid.rcd.UnmarshalBinaryAuth call",
	})
	factoidrcdNewRCD_1SummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd_NewRCD_1_Summary",
		Help: "Summary of factoid.rcd.NewRCD_1 call",
	})
	factoidrcdNewRCD_2SummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd_NewRCD_2_Summary",
		Help: "Summary of factoid.rcd.NewRCD_2 call",
	})
	factoidrcdCreateRCD2SummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd_CreateRCD_Summary",
		Help: "Summary of factoid.rcd.CreateRCD call",
	})
	// rct1.go
	factoidrcd1GetHashSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_GetHash_Summary",
		Help: "Summary of factoid.rcd1.GetHash call",
	})
	factoidrcd1UnmarshalBinarySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_UnmarshalBinary_Summary",
		Help: "Summary of factoid.rcd1.UnmarshalBinary call",
	})
	factoidrcd1JSONByteSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_JSONByte_Summary",
		Help: "Summary of factoid.rcd1.JSONByte call",
	})
	factoidrcd1JSONStringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_JSONString_Summary",
		Help: "Summary of factoid.rcd1.JSONString call",
	})
	factoidrcd1JSONBufferSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_JSONBuffer_Summary",
		Help: "Summary of factoid.rcd1.JSONBuffer call",
	})
	factoidrcd1StringSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_String_Summary",
		Help: "Summary of factoid.rcd1.String call",
	})
	factoidrcd1MarshalTextSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_MarshalText_Summary",
		Help: "Summary of factoid.rcd1.MarshalText call",
	})
	factoidrcd1CheckSigSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_CheckSig_Summary",
		Help: "Summary of factoid.rcd1.CheckSig call",
	})
	factoidrcd1CloneSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_Clone_Summary",
		Help: "Summary of factoid.rcd1.Clone call",
	})
	factoidrcd1GetAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_GetAddress_Summary",
		Help: "Summary of factoid.rcd1.GetAddress call",
	})
	factoidrcd1GetPublicKeySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_GetPublicKey_Summary",
		Help: "Summary of factoid.rcd1.GetPublicKey call",
	})
	factoidrcd1NumberOfSignaturesSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_NumberOfSignatures_Summary",
		Help: "Summary of factoid.rcd1.NumberOfSignatures call",
	})
	factoidrcd1IsEqualSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_IsEqual_Summary",
		Help: "Summary of factoid.rcd1.IsEqual call",
	})
	factoidrcd1UnmarshalBinaryDataSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.rcd1.UnmarshalBinaryData call",
	})
	factoidrcd1MarshalBinarySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_MarshalBinary_Summary",
		Help: "Summary of factoid.rcd1.MarshalBinary call",
	})
	factoidrcd1CustomMarshalTextSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd1_CustomMarshalText_Summary",
		Help: "Summary of factoid.rcd1.CustomMarshalText call",
	})
	//rcd2 variables
	factoidrcd2GetAddressSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_GetAddress_Summary",
		Help: "Summary of factoid.rcd2.GetAddress call",
	})
	factoidrcd2GetHashSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_GetHash_Summary",
		Help: "Summary of factoid.rcd2.GetHash call",
	})
	factoidrcd2NumberOfSignaturesSummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_NumberOfSignatures_Summary",
		Help: "Summary of factoid.rcd2.NumberOfSignatures call",
	})
	factoidrcd2UnmarshalBinarySummaryOpts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_UnmarshalBinary_Summary",
		Help: "Summary of factoid.rcd2.UnmarshalBinary call",
	})
	factoidrcd2CheckSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_CheckSig_Summary",
		Help: "Summary of factoid.rcd2.CheckSig call",
	})
	factoidrcd2JSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_JSONByte_Summary",
		Help: "Summary of factoid.rcd2.JSONByte call",
	})
	factoidrcd2JSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_JSONString_Summary",
		Help: "Summary of factoid.rcd2.JSONString call",
	})
	factoidrcd2JSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_JSONBuffer_Summary",
		Help: "Summary of factoid.rcd2.JSONBuffer call",
	})
	factoidrcd2String = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_String_Summary",
		Help: "Summary of factoid.rcd2.String call",
	})
	factoidrcd2Clone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_Clone_Summary",
		Help: "Summary of factoid.rcd2.Clone call",
	})
	factoidrcd2IsEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_IsEqual_Summary",
		Help: "Summary of factoid.rcd2.IsEqual call",
	})
	factoidrcd2UnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.rcd2.UnmarshalBinaryData call",
	})
	factoidrcd2MarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_2MarshalBinary_Summary",
		Help: "Summary of factoid.rcd2.factoidrcd2MarshalBinary call",
	})
	factoidrcd2CustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_rcd2_CustomMarshalText_Summary",
		Help: "Summary of factoid.rcd2.CustomMarshalText call",
	})
	//signature.go variable =prometheus.NewSummary(prometheus.SummaryOpts{
	factoidSignatureVerify = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_Verify_Summary",
		Help: "Summary of factoid.Signature.Verify call",
	})
	factoidSignatureBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_Bytes_Summary",
		Help: "Summary of factoid.Signature.Bytes call",
	})
	factoidSignatureGetKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_GetKey_Summary",
		Help: "Summary of factoid.Signature.GetKey call",
	})
	factoidSignatureGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_GetHash_Summary",
		Help: "Summary of factoid.Signature.GetHash call",
	})
	factoidSignatureMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_MarshalText_Summary",
		Help: "Summary of factoid.Signature.MarshalText call",
	})
	factoidSignatureJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_JSONByte_Summary",
		Help: "Summary of factoid.Signature.JSONByte call",
	})
	factoidSignatureJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_JSONString_Summary",
		Help: "Summary of factoid.Signature.JSONString call",
	})
	factoidSignatureJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_JSONBuffer_Summary",
		Help: "Summary of factoid.Signature.JSONBuffer call",
	})
	factoidSignatureString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_String_Summary",
		Help: "Summary of factoid.Signature.String call",
	})
	factoidSignatureIsEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_IsEqual_Summary",
		Help: "Summary of factoid.Signature.IsEqual call",
	})
	factoidSignatureSetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_SetSignature_Summary",
		Help: "Summary of factoid.Signature.SetSignature call",
	})
	factoidSignatureGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_GetSignature_Summary",
		Help: "Summary of factoid.Signature.GetSignature call",
	})
	factoidSignatureMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_MarshalBinary_Summary",
		Help: "Summary of factoid.Signature.MarshalBinary call",
	})
	factoidSignatureCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_CustomMarshalText_Summary",
		Help: "Summary of factoid.Signature.CustomMarshalText call",
	})
	factoidSignatureUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.Signature.UnmarshalBinaryData call",
	})
	factoidSignatureUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_UnmarshalBinary_Summary",
		Help: "Summary of factoid.Signature.UnmarshalBinary call",
	})
	factoidSignatureNewED25519Signature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Signature_NewED25519Signature_Summary",
		Help: "Summary of factoid.Signature.NewED25519Signature call",
	})
	// signatureblock cod =prometheus.NewSummary(prometheus.SummaryOpts{
	factoidSignatureBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_UnmarshalBinary_Summary",
		Help: "Summary of factoid.SignatureBlock.UnmarshalBinary call",
	})
	factoidSignatureBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_JSONByte_Summary",
		Help: "Summary of factoid.SignatureBlock.JSONByte call",
	})
	factoidSignatureBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_JSONString_Summary",
		Help: "Summary of factoid.SignatureBlock.JSONString call",
	})
	factoidSignatureBlockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_JSONBuffer_Summary",
		Help: "Summary of factoid.SignatureBlock.JSONBuffer call",
	})
	factoidSignatureBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_String_Summary",
		Help: "Summary of factoid.SignatureBlock.String call",
	})
	factoidSignatureBlockIsEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_IsEqual_Summary",
		Help: "Summary of factoid.SignatureBlock.IsEqual call",
	})
	factoidSignatureBlockAddSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_AddSignature_Summary",
		Help: "Summary of factoid.SignatureBlock.AddSignature call",
	})
	factoidSignatureBlockGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_GetSignature_Summary",
		Help: "Summary of factoid.SignatureBlock.GetSignature call",
	})
	factoidSignatureBlockGetSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_GetSignatures_Summary",
		Help: "Summary of factoid.SignatureBlock.GetSignatures call",
	})
	factoidSignatureBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_MarshalBinary_Summary",
		Help: "Summary of factoid.SignatureBlock.MarshalBinary call",
	})
	factoidSignatureBlockCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_CustomMarshalText_Summary",
		Help: "Summary of factoid.SignatureBlock.CustomMarshalText call",
	})
	factoidSignatureBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.SignatureBlock.UnmarshalBinaryData call",
	})
	factoidSignatureBlockNewSingleSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_SignatureBlock_NewSingleSignatureBlock_Summary",
		Help: "Summary of factoid.SignatureBlock.NewSingleSignatureBlock call",
	})
	// transaction.go
	factoidTransactionNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_New_Summary",
		Help: "Summary of factoid.Transaction.New call",
	})
	factoidTransactionSetBlockHeight = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_factoid_Transaction_SetBlockHeight_Summary",
		Help: "Summary of factoid.Transaction.SetBlockHeight call",
	})
	factoidTransactionGetBlockHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetBlockHeight_Summary",
		Help: "Summary of factoid.Transaction.GetBlockHeight call",
	})
	factoidTransactionclearCaches = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_clearCaches_Summary",
		Help: "Summary of factoid.Transaction.clearCaches call",
	})
	factoidTransactionGetVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetVersion_Summary",
		Help: "Summary of factoid.Transaction.GetVersion call",
	})
	factoidTransactionGetTxID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetTxID_Summary",
		Help: "Summary of factoid.Transaction.GetTxID call",
	})
	factoidTransactionGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetHash_Summary",
		Help: "Summary of factoid.Transaction.GetHash call",
	})
	factoidTransactionGetFullHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetFullHash_Summary",
		Help: "Summary of factoid.Transaction.GetFullHash call",
	})
	factoidTransactionGetSigHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSigHash_Summary",
		Help: "Summary of factoid.Transaction.GetSigHash call",
	})
	factoidTransactionSetString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_String_Summary",
		Help: "Summary of factoid.Transaction.String call",
	})
	factoidTransactionGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetTimestamp_Summary",
		Help: "Summary of factoid.Transaction.GetTimestamp call",
	})
	factoidTransactionSetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_SetTimestamp_Summary",
		Help: "Summary of factoid.Transaction.SetTimestamp call",
	})
	factoidTransactionSetSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_SetSignatureBlock_Summary",
		Help: "Summary of factoid.Transaction.SetSignatureBlock call",
	})
	factoidTransactionGetSignatureBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSignatureBlock_Summary",
		Help: "Summary of factoid.Transaction.GetSignatureBlock call",
	})
	factoidTransactionAddRCD = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddRCD_Summary",
		Help: "Summary of factoid.Transaction.AddRCD call",
	})
	factoidTransactionCalculateFee = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_CalculateFee_Summary",
		Help: "Summary of factoid.Transaction.CalculateFee call",
	})
	factoidTransactionValidateAmounts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_ValidateAmounts_Summary",
		Help: "Summary of factoid.Transaction.ValidateAmounts call",
	})
	factoidTransactionSTotalInputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalInputs_Summary",
		Help: "Summary of factoid.Transaction.TotalInputs call",
	})
	factoidTransactionTotalOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalOutputs_Summary",
		Help: "Summary of factoid.Transaction.TotalOutputs call",
	})
	factoidTransactionTotalECs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_TotalECs_Summary",
		Help: "Summary of factoid.Transaction.TotalECs call",
	})
	factoidTransactionValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_Validate_Summary",
		Help: "Summary of factoid.Transaction.Validate call",
	})
	factoidTransactionValidateSignatures = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_ValidateSignatures_Summary",
		Help: "Summary of factoid.Transaction.ValidateSignatures call",
	})
	factoidTransactionIsEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_IsEqualt_Summary",
		Help: "Summary of factoid.Transaction.IsEqual call",
	})
	factoidTransactionGetInputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetInputs_Summary",
		Help: "Summary of factoid.Transaction.GetInputs call",
	})
	factoidTransactionGetOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetOutputs_Summary",
		Help: "Summary of factoid.Transaction.GetOutputs call",
	})
	factoidTransactionGetECOutputs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetECOutputs_Summary",
		Help: "Summary of factoid.Transaction.GetECOutputs call",
	})
	factoidTransactionGetRCDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetRCDs_Summary",
		Help: "Summary of factoid.Transaction.GetRCDs call",
	})
	factoidTransactionGetSignatureBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetSignatureBlocks_Summary",
		Help: "Summary of factoid.Transaction.GetSignatureBlocks call",
	})
	factoidTransactionGetInput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetInput_Summary",
		Help: "Summary of factoid.Transaction.GetInput call",
	})
	factoidTransactionGetOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetOutput_Summary",
		Help: "Summary of factoid.Transaction.GetOutput call",
	})
	factoidTransactionGetECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetECOutput_Summary",
		Help: "Summary of factoid.Transaction.GetECOutput call",
	})
	factoidTransactionGetRCD = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_GetRCD_Summary",
		Help: "Summary of factoid.Transaction.GetRCD call",
	})
	factoidTransactionUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.Transaction.UnmarshalBinaryData call",
	})
	factoidTransactionUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_UnmarshalBinary_Summary",
		Help: "Summary of factoid.Transaction.UnmarshalBinary call",
	})
	factoidTransactionMarshalBinarySig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_MarshalBinarySig_Summary",
		Help: "Summary of factoid.Transaction.MarshalBinarySig call",
	})
	factoidTransactionMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_MarshalBinary_Summary",
		Help: "Summary of factoid.Transaction.MarshalBinary call",
	})
	factoidTransactionAddOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddOutput_Summary",
		Help: "Summary of factoid.Transaction.AddOutput call",
	})
	factoidTransactionAddECOutput = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddECOutput_Summary",
		Help: "Summary of factoid.Transaction.AddECOutput call",
	})
	factoidTransactionCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_CustomMarshalText_Summary",
		Help: "Summary of factoid.Transaction.CustomMarshalText call",
	})
	factoidTransactionAddAuthorization = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_AddAuthorization_Summary",
		Help: "Summary of factoid.Transaction.AddAuthorization call",
	})
	factoidTransactionJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_JSONByte_Summary",
		Help: "Summary of factoid.Transaction.JSONByte call",
	})
	factoidTransactionJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_JSONString_Summary",
		Help: "Summary of factoid.Transaction.JSONString call",
	})
	factoidTransactionJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_JSONString_Summary",
		Help: "Summary of factoid.Transaction.JSONString call",
	})
	factoidTransactionHasUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_Transaction_HasUserAddress_Summary",
		Help: "Summary of factoid.Transaction.HasUserAddress call",
	})
	//transaddress
	factoidTransAddressSetUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetUserAddress_Summary",
		Help: "Summary of factoid.TransAddress.SetUserAddress call",
	})
	factoidTransAddressGetUserAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetUserAddress_Summary",
		Help: "Summary of factoid.TransAddress.GetUserAddress call",
	})
	factoidTransAddressGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetHash_Summary",
		Help: "Summary of factoid.TransAddress.GetHash call",
	})
	factoidTransAddressUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_UnmarshalBinary_Summary",
		Help: "Summary of factoid.TransAddress.UnmarshalBinary call",
	})
	factoidTransAddressCustomMarshalText = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalText_Summary",
		Help: "Summary of factoid.TransAddress.CustomMarshalText call",
	})
	factoidTransAddressJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_JSONByte_Summary",
		Help: "Summary of factoid.TransAddress.JSONByte call",
	})
	factoidTransAddressJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_JSONString_Summary",
		Help: "Summary of factoid.TransAddress.JSONString call",
	})
	factoidTransAddressJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_JSONBuffer_Summary",
		Help: "Summary of factoid.TransAddress.JSONBuffer call",
	})
	factoidTransAddressString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_String_Summary",
		Help: "Summary of factoid.TransAddress.String call",
	})
	factoidTransAddressIsEqual = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_IsEqual_Summary",
		Help: "Summary of factoid.TransAddress.IsEqual call",
	})
	factoidTransAddressUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_UnmarshalBinaryData_Summary",
		Help: "Summary of factoid.TransAddress.UnmarshalBinaryData call",
	})
	factoidTransAddressMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_MarshalBinary_Summary",
		Help: "Summary of factoid.TransAddress.MarshalBinary call",
	})
	factoidTransAddressGetName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetName_Summary",
		Help: "Summary of factoid.TransAddress.GetName call",
	})
	factoidTransAddressGetAmount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetAmount_Summary",
		Help: "Summary of factoid.TransAddress.GetAmount call",
	})
	factoidTransAddressSetAmount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetAmount_Summary",
		Help: "Summary of factoid.TransAddress.SetAmount call",
	})
	factoidTransAddressGetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_GetAddress_Summary",
		Help: "Summary of factoid.TransAddress.GetAddress call",
	})
	factoidTransAddressSetAddress = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_SetAddress_Summary",
		Help: "Summary of factoid.TransAddress.SetAddress call",
	})
	factoidTransAddressCustomMarshalTextAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextAll_Summary",
		Help: "Summary of factoid.TransAddress.CustomMarshalTextAll call",
	})
	factoidTransAddressCustomMarshalText2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalText2_Summary",
		Help: "Summary of factoid.TransAddress.CustomMarshalText2 call",
	})
	factoidTransAddressCustomMarshalTextEC2 = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_factoid_TransAddress_CustomMarshalTextEC2_Summary",
		Help: "Summary of factoid.TransAddress.CustomMarshalTextEC2 call",
	})
)

func InitPrometheus() {
	//anchorRecord.go variables
	prometheus.MustRegister(factoidErrors)
	prometheus.MustRegister(factoidAddressCustomMarshalText)
	prometheus.MustRegister(factoidAddressNewAddress)
	prometheus.MustRegister(factoidAddressCreateAddress)
	prometheus.MustRegister(factoidCoinbaseUpdateAmount)
	prometheus.MustRegister(factoidCoinbaseGetCoinbase)
	//conversions
	prometheus.MustRegister(factoidConversionsPublicKeyStringToECAddressString)
	prometheus.MustRegister(factoidConversionsPublicKeyToECAddress)
	prometheus.MustRegister(factoidConversionsPublicKeyStringToFactoidAddressStringSummaryOpts)
	prometheus.MustRegister(factoidConversionsPublicKeyToFactoidAddressSummaryOpts)
	prometheus.MustRegister(factoidConversionsPublicKeyStringToFactoidAddressSummaryOpts)
	prometheus.MustRegister(factoidConversionsPublicKeyStringToFactoidRCDAddressSummaryOpts)
	prometheus.MustRegister(factoidConversionsHumanReadiblePrivateKeyStringToEverythingStringSummaryOpts)
	prometheus.MustRegister(factoidConversionsPrivateKeyStringToEverythingStringSummaryOpts)
	// factoid
	prometheus.MustRegister(factoidFactoidFactoidTx_VersionCheckSummaryOpts)
	prometheus.MustRegister(factoidFactoidFactoidTx_LocktimeCheckSummaryOpts)
	prometheus.MustRegister(factoidFactoidFactoidTx_RCDVersionCheckSummaryOpts)
	prometheus.MustRegister(factoidFactoidFactoidTx_RCDTypeCheckSummaryOpts)
	//fblock
	prometheus.MustRegister(factoidFBlockGetEntryHashesSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetTransactionByHashSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetEntrySigHashesSummaryOpts)
	prometheus.MustRegister(factoidFBlockNewSummaryOpts)
	prometheus.MustRegister(factoidFBlockDatabasePrimaryIndexSummaryOpts)
	prometheus.MustRegister(factoidFBlockDatabaseSecondaryIndexSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetDatabaseHeightSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetCoinbaseTimestampSummaryOpts)
	prometheus.MustRegister(factoidFBlockEndOfPeriodSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetTransactionsSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetNewInstanceSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetEndOfPeriodSummaryOpts)
	prometheus.MustRegister(factoidFBlockMarshalTransSummaryOpts)
	prometheus.MustRegister(factoidFBlockMarshalHeaderSummaryOpts)
	prometheus.MustRegister(factoidFBlockMarshalBinarySummaryOpts)
	prometheus.MustRegister(factoidFBlockUnmarshalFBlockSummaryOpts)
	prometheus.MustRegister(factoidFBlockUnmarshalBinaryDataSummaryOpts)
	prometheus.MustRegister(factoidFBlockUnmarshalBinarySummaryOpts)
	prometheus.MustRegister(factoidFBlockIsEqualSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetChainIDSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetHashSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetLedgerKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetLedgerMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetBodyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetPrevKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockSetPrevKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetPrevLedgerKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockSetPrevLedgerKeyMRSummaryOpts)
	prometheus.MustRegister(factoidFBlockCalculateHashesSummaryOpts)
	prometheus.MustRegister(factoidFBlockSetDBHeightSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetDBHeightSummaryOpts)
	prometheus.MustRegister(factoidFBlockSetExchRateSummaryOpts)
	prometheus.MustRegister(factoidFBlockGetExchRateSummaryOpts)
	prometheus.MustRegister(factoidFBlockValidateTransactionSummaryOpts)
	prometheus.MustRegister(factoidFBlockValidateSummaryOpts)
	prometheus.MustRegister(factoidFBlockAddCoinbaseSummaryOpts)
	prometheus.MustRegister(factoidFBlockAddTransactionSummaryOpts)
	prometheus.MustRegister(factoidFBlockStringSummaryOpts)
	prometheus.MustRegister(factoidFBlockCustomMarshalTextSummaryOpts)
	prometheus.MustRegister(factoidFBlockJSONByteSummaryOpts)
	prometheus.MustRegister(factoidFBlockJSONStringSummaryOpts)
	prometheus.MustRegister(factoidFBlockJSONBufferSummaryOpts)
	prometheus.MustRegister(factoidFBlockMarshalJSONSummaryOpts)
	prometheus.MustRegister(factoidFBlockNewFBlockSummaryOpts)
	prometheus.MustRegister(factoidFBlockCheckBlockPairIntegritySummaryOpts)
	//genesisBlockNew
	prometheus.MustRegister(factoidgenesisBlockNewGetGenesisFBlockSummaryOpts)
	//inaddress
	prometheus.MustRegister(factoidInAddressStringSummaryOpts)
	prometheus.MustRegister(factoidInAddressCustomMarshalText)
	prometheus.MustRegister(factoidInAddressNewInAddress)
	// outaddress
	prometheus.MustRegister(factoidOutAddressStringSummaryOpts)
	prometheus.MustRegister(factoidOutAddressCustomMarshalText)
	prometheus.MustRegister(factoidOutAddressNewOutAddress)
	// outecaddress
	prometheus.MustRegister(factoidOutECAddressStringSummaryOpts)
	prometheus.MustRegister(factoidOutECAddressCustomMarshalText)
	prometheus.MustRegister(factoidOutECAddressNewOutAddress)
	//rcd.go variables
	prometheus.MustRegister(factoidrcdUnmarshalBinaryAuthSummaryOpts)
	prometheus.MustRegister(factoidrcdNewRCD_1SummaryOpts)
	prometheus.MustRegister(factoidrcdNewRCD_2SummaryOpts)
	prometheus.MustRegister(factoidrcdCreateRCD2SummaryOpts)
	// rct1.go
	prometheus.MustRegister(factoidrcd1GetHashSummaryOpts)
	prometheus.MustRegister(factoidrcd1UnmarshalBinarySummaryOpts)
	prometheus.MustRegister(factoidrcd1JSONByteSummaryOpts)
	prometheus.MustRegister(factoidrcd1JSONStringSummaryOpts)
	prometheus.MustRegister(factoidrcd1JSONBufferSummaryOpts)
	prometheus.MustRegister(factoidrcd1StringSummaryOpts)
	prometheus.MustRegister(factoidrcd1MarshalTextSummaryOpts)
	prometheus.MustRegister(factoidrcd1CheckSigSummaryOpts)
	prometheus.MustRegister(factoidrcd1CloneSummaryOpts)
	prometheus.MustRegister(factoidrcd1GetAddressSummaryOpts)
	prometheus.MustRegister(factoidrcd1GetPublicKeySummaryOpts)
	prometheus.MustRegister(factoidrcd1NumberOfSignaturesSummaryOpts)
	prometheus.MustRegister(factoidrcd1IsEqualSummaryOpts)
	prometheus.MustRegister(factoidrcd1UnmarshalBinaryDataSummaryOpts)
	prometheus.MustRegister(factoidrcd1MarshalBinarySummaryOpts)
	prometheus.MustRegister(factoidrcd1CustomMarshalTextSummaryOpts)
	//rcd2 variables
	prometheus.MustRegister(factoidrcd2GetAddressSummaryOpts)
	prometheus.MustRegister(factoidrcd2GetHashSummaryOpts)
	prometheus.MustRegister(factoidrcd2NumberOfSignaturesSummaryOpts)
	prometheus.MustRegister(factoidrcd2UnmarshalBinarySummaryOpts)
	prometheus.MustRegister(factoidrcd2CheckSig)
	prometheus.MustRegister(factoidrcd2JSONByte)
	prometheus.MustRegister(factoidrcd2JSONString)
	prometheus.MustRegister(factoidrcd2JSONBuffer)
	prometheus.MustRegister(factoidrcd2String)
	prometheus.MustRegister(factoidrcd2Clone)
	prometheus.MustRegister(factoidrcd2IsEqual)
	prometheus.MustRegister(factoidrcd2UnmarshalBinaryData)
	prometheus.MustRegister(factoidrcd2MarshalBinary)
	prometheus.MustRegister(factoidrcd2CustomMarshalText)
	//signature.go variables
	prometheus.MustRegister(factoidSignatureVerify)
	prometheus.MustRegister(factoidSignatureBytes)
	prometheus.MustRegister(factoidSignatureGetKey)
	prometheus.MustRegister(factoidSignatureGetHash)
	prometheus.MustRegister(factoidSignatureMarshalText)
	prometheus.MustRegister(factoidSignatureJSONByte)
	prometheus.MustRegister(factoidSignatureJSONString)
	prometheus.MustRegister(factoidSignatureJSONBuffer)
	prometheus.MustRegister(factoidSignatureString)
	prometheus.MustRegister(factoidSignatureIsEqual)
	prometheus.MustRegister(factoidSignatureSetSignature)
	prometheus.MustRegister(factoidSignatureGetSignature)
	prometheus.MustRegister(factoidSignatureMarshalBinary)
	prometheus.MustRegister(factoidSignatureCustomMarshalText)
	prometheus.MustRegister(factoidSignatureUnmarshalBinaryData)
	prometheus.MustRegister(factoidSignatureUnmarshalBinary)
	prometheus.MustRegister(factoidSignatureNewED25519Signature)
	// signatureblock code
	prometheus.MustRegister(factoidSignatureBlockUnmarshalBinary)
	prometheus.MustRegister(factoidSignatureBlockJSONByte)
	prometheus.MustRegister(factoidSignatureBlockJSONString)
	prometheus.MustRegister(factoidSignatureBlockJSONBuffer)
	prometheus.MustRegister(factoidSignatureBlockString)
	prometheus.MustRegister(factoidSignatureBlockIsEqual)
	prometheus.MustRegister(factoidSignatureBlockAddSignature)
	prometheus.MustRegister(factoidSignatureBlockGetSignature)
	prometheus.MustRegister(factoidSignatureBlockGetSignatures)
	prometheus.MustRegister(factoidSignatureBlockMarshalBinary)
	prometheus.MustRegister(factoidSignatureBlockCustomMarshalText)
	prometheus.MustRegister(factoidSignatureBlockUnmarshalBinaryData)
	prometheus.MustRegister(factoidSignatureBlockNewSingleSignatureBlock)
	// transaction.go
	prometheus.MustRegister(factoidTransactionNew)
	prometheus.MustRegister(factoidTransactionSetBlockHeight)
	prometheus.MustRegister(factoidTransactionGetBlockHeight)
	prometheus.MustRegister(factoidTransactionclearCaches)
	prometheus.MustRegister(factoidTransactionGetVersion)
	prometheus.MustRegister(factoidTransactionGetTxID)
	prometheus.MustRegister(factoidTransactionGetHash)
	prometheus.MustRegister(factoidTransactionGetFullHash)
	prometheus.MustRegister(factoidTransactionGetSigHash)
	prometheus.MustRegister(factoidTransactionSetString)
	prometheus.MustRegister(factoidTransactionGetTimestamp)
	prometheus.MustRegister(factoidTransactionSetTimestamp)
	prometheus.MustRegister(factoidTransactionSetSignatureBlock)
	prometheus.MustRegister(factoidTransactionGetSignatureBlock)
	prometheus.MustRegister(factoidTransactionAddRCD)
	prometheus.MustRegister(factoidTransactionCalculateFee)
	prometheus.MustRegister(factoidTransactionValidateAmounts)
	prometheus.MustRegister(factoidTransactionSTotalInputs)
	prometheus.MustRegister(factoidTransactionTotalOutputs)
	prometheus.MustRegister(factoidTransactionTotalECs)
	prometheus.MustRegister(factoidTransactionValidate)
	prometheus.MustRegister(factoidTransactionValidateSignatures)
	prometheus.MustRegister(factoidTransactionIsEqual)
	prometheus.MustRegister(factoidTransactionGetInputs)
	prometheus.MustRegister(factoidTransactionGetOutputs)
	prometheus.MustRegister(factoidTransactionGetECOutputs)
	prometheus.MustRegister(factoidTransactionGetRCDs)
	prometheus.MustRegister(factoidTransactionGetSignatureBlocks)
	prometheus.MustRegister(factoidTransactionGetInput)
	prometheus.MustRegister(factoidTransactionGetOutput)
	prometheus.MustRegister(factoidTransactionGetECOutput)
	prometheus.MustRegister(factoidTransactionGetRCD)
	prometheus.MustRegister(factoidTransactionUnmarshalBinaryData)
	prometheus.MustRegister(factoidTransactionUnmarshalBinary)
	prometheus.MustRegister(factoidTransactionMarshalBinarySig)
	prometheus.MustRegister(factoidTransactionMarshalBinary)
	prometheus.MustRegister(factoidTransactionAddOutput)
	prometheus.MustRegister(factoidTransactionAddECOutput)
	prometheus.MustRegister(factoidTransactionCustomMarshalText)
	prometheus.MustRegister(factoidTransactionAddAuthorization)
	prometheus.MustRegister(factoidTransactionJSONByte)
	prometheus.MustRegister(factoidTransactionJSONByte)
	prometheus.MustRegister(factoidTransactionJSONString)
	prometheus.MustRegister(factoidTransactionJSONBuffer)
	prometheus.MustRegister(factoidTransactionHasUserAddress)
	//transaddress
	prometheus.MustRegister(factoidTransAddressSetUserAddress)
	prometheus.MustRegister(factoidTransAddressGetUserAddress)
	prometheus.MustRegister(factoidTransAddressGetHash)
	prometheus.MustRegister(factoidTransAddressUnmarshalBinary)
	prometheus.MustRegister(factoidTransAddressCustomMarshalText)
	prometheus.MustRegister(factoidTransAddressJSONByte)
	prometheus.MustRegister(factoidTransAddressJSONString)
	prometheus.MustRegister(factoidTransAddressJSONBuffer)
	prometheus.MustRegister(factoidTransAddressString)
	prometheus.MustRegister(factoidTransAddressIsEqual)
	prometheus.MustRegister(factoidTransAddressUnmarshalBinaryData)
	prometheus.MustRegister(factoidTransAddressMarshalBinary)
	prometheus.MustRegister(factoidTransAddressGetName)
	prometheus.MustRegister(factoidTransAddressGetAmount)
	prometheus.MustRegister(factoidTransAddressSetAmount)
	prometheus.MustRegister(factoidTransAddressGetAddress)
	prometheus.MustRegister(factoidTransAddressSetAddress)
	prometheus.MustRegister(factoidTransAddressCustomMarshalTextAll)
	prometheus.MustRegister(factoidTransAddressCustomMarshalText2)
	prometheus.MustRegister(factoidTransAddressCustomMarshalTextEC2)

	/* these are what is being added to each function to measure function time
	       callTime := time.Now().UnixNano()

	   	runTime := time.Now().UnixNano() - callTime
	   	v1DBlockByHeightSummary.Observe(float64(runTime))
	*/

}
