// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package adminBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	//adminBlock.go variables
	adminBlockErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_anchor_error_responses",
		Help: "Number of error responses from anchor objects",
	})    
	adminBlockString = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_String_Summary",
		Help: "Summary of adminblock.String call",
	}) 
	adminBlockUpdateState = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_UpdateState_Summary",
		Help: "Summary of adminblock.UpdateState call",
	}) 
	adminBlockAddDBSig = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddDBSig_Summary",
		Help: "Summary of adminblock.AddDBSig call",
	}) 
	adminBlockAddFedServer = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddFedServer_Summary",
		Help: "Summary of adminblock.AddFedServer call",
	}) 
	adminBlockAddAuditServer = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddAuditServer_Summary",
		Help: "Summary of adminblock.AddAuditServer call",
	}) 
	adminBlockRemoveFederatedServer = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_RemoveFederatedServer_Summary",
		Help: "Summary of adminblock.RemoveFederatedServer call",
    })
	adminBlockAddMatryoshkaHash = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddMatryoshkaHash_Summary",
		Help: "Summary of adminblock.AddMatryoshkaHash call",
	}) 
	adminBlockAddFederatedServerSigningKey = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddFederatedServerSigningKey_Summary",
		Help: "Summary of adminblock.AddFederatedServerSigningKey call",
	}) 
	adminBlockAddFederatedServerBitcoinAnchorKey = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddFederatedServerBitcoinAnchorKey_Summary",
		Help: "Summary of adminblock.AddFederatedServerBitcoinAnchorKey call",
	}) 
	adminBlockAddEntry = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddEntry_Summary",
		Help: "Summary of adminblock.AddEntry call",
	}) 
	adminBlockAddServerFault = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddServerFault_Summary",
		Help: "Summary of adminblock.AddServerFault call",
	}) 
	adminBlockGetHeader = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetHeader_Summary",
		Help: "Summary of adminblock.GetHeader call",
	}) 
	adminBlockSetHeader = prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_SetHeader_Summary",
		Help: "Summary of adminblock.SetHeader call",
	}) 
	adminBlockGetABEntries= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetABEntries_Summary",
		Help: "Summary of adminblock.GetABEntries call",
	}) 
	adminBlockGetDBHeight= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetDBHeight_Summary",
		Help: "Summary of adminblock.GetDBHeight call",
	}) 
	adminBlockSetABEntries= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_SetABEntries_Summary",
		Help: "Summary of adminblock.SetABEntries call",
	}) 
	adminBlockNew= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_New_Summary",
		Help: "Summary of adminblock.New call",
	}) 
	adminBlockGetDatabaseHeight= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetDatabaseHeight_Summary",
		Help: "Summary of adminblock.GetDatabaseHeight call",
	}) 
	adminBlockGetChainID= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetChainID_Summary",
		Help: "Summary of adminblock.GetChainID call",
	}) 
	adminBlockDatabasePrimaryIndex= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_DatabasePrimaryIndex_Summary",
		Help: "Summary of adminblock.DatabasePrimaryIndex call",
	}) 
	adminBlockDatabaseSecondaryIndex= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of adminblock.DatabaseSecondaryIndex call",
	}) 
	adminBlockGetHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetHash_Summary",
		Help: "Summary of adminblock.GetHash call",
	}) 
	adminBlockGetKeyMR= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetKeyMR_Summary",
		Help: "Summary of adminblock.GetKeyMR call",
	}) 
	adminBlockBackReferenceHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_BackReferenceHash_Summary",
		Help: "Summary of adminblock.BackReferenceHash call",
	}) 
	adminBlockLookupHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_LookupHash_Summary",
		Help: "Summary of adminblock.LookupHash call",
	}) 
	adminBlockAddABEntry= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddABEntry_Summary",
		Help: "Summary of adminblock.AddABEntry call",
	}) 
	adminBlockAddFirstABEntry= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_AddFirstABEntry_Summary",
		Help: "Summary of adminblock.AddFirstABEntry call",
	}) 
	adminBlockMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_MarshalBinary_Summary",
		Help: "Summary of adminblock.MarshalBinary call",
	}) 
	adminBlockUnmarshalABlock= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_UnmarshalABlock_Summary",
		Help: "Summary of adminblock.UnmarshalABlock call",
	}) 
	adminBlockUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_UnmarshalBinaryData_Summary",
		Help: "Summary of adminblock.UnmarshalBinaryData call",
	}) 
	adminBlockUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_UnmarshalBinary_Summary",
		Help: "Summary of adminblock.UnmarshalBinary call",
	}) 
	adminBlockGetDBSignature= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_GetDBSignature_Summary",
		Help: "Summary of adminblock.GetDBSignature call",
	}) 
	adminBlockJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_JSONByte_Summary",
		Help: "Summary of adminblock.JSONByte call",
	}) 
	adminBlockJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_JSONString_Summary",
		Help: "Summary of adminblock.JSONString call",
	}) 
	adminBlockJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_JSONBuffer_Summary",
		Help: "Summary of adminblock.JSONBuffer call",
	}) 
	adminBlockMarshalJSON= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_MarshalJSON_Summary",
		Help: "Summary of adminblock.MarshalJSON call",
	}) 
	adminBlockNewAdminBlock= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_NewAdminBlock_Summary",
		Help: "Summary of adminblock.NewAdminBlock call",
	}) 
	adminBlockCheckBlockPairIntegrity= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of adminblock.CheckBlockPairIntegrity call",
	}) 

    // adminBlockHeader.go variables
 	adminBlockHeaderString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_String_Summary",
		Help: "Summary of adminblockheader.String call",
	}) 
 	adminBlockHeaderGetMessageCount= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetMessageCount_Summary",
		Help: "Summary of adminblockheader.GetMessageCount call",
	}) 
 	adminBlockHeaderSetMessageCount= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_SetMessageCount_Summary",
		Help: "Summary of adminblockheader.SetMessageCount call",
	}) 
 	adminBlockHeaderGetBodySize= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetBodySize_Summary",
		Help: "Summary of adminblockheader.GetBodySize call",
	}) 
 	adminBlockHeaderSetBodySize= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_SetBodySize_Summary",
		Help: "Summary of adminblockheader.SetBodySize call",
	}) 
	adminBlockHeaderGetAdminChainID= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetAdminChainID_Summary",
		Help: "Summary of adminblockheader.GetAdminChainID call",
	}) 
	adminBlockHeaderGetDBHeight= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetDBHeight_Summary",
		Help: "Summary of adminblockheader.GetDBHeight call",
	}) 
	adminBlockHeaderGetHeaderExpansionArea= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetHeaderExpansionArea_Summary",
		Help: "Summary of adminblockheader.GetHeaderExpansionArea call",
	}) 
 	adminBlockHeaderGetHeaderExpansionSize= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetHeaderExpansionSize_Summary",
		Help: "Summary of adminblockheader.GetHeaderExpansionSize call",
	}) 
	adminBlockHeaderGetPrevBackRefHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_GetPrevBackRefHash_Summary",
		Help: "Summary of adminblockheader.GetPrevBackRefHash call",
	}) 
 	adminBlockHeaderSetDBHeight= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_SetDBHeight_Summary",
		Help: "Summary of adminblockheader.SetDBHeight call",
	}) 
 	adminBlockHeaderSetHeaderExpansionArea= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_SetHeaderExpansionArea_Summary",
		Help: "Summary of adminblockheader.SetHeaderExpansionArea call",
	}) 
	adminBlockHeaderSetPrevBackRefHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_SetPrevBackRefHash_Summary",
		Help: "Summary of adminblockheader.SetPrevBackRefHash call",
	}) 
	adminBlockHeaderMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_MarshalBinary_Summary",
		Help: "Summary of adminblockheader.MarshalBinary call",
	}) 
	adminBlockHeaderUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_UnmarshalBinaryData_Summary",
		Help: "Summary of adminblockheader.UnmarshalBinaryData call",
	}) 
	adminBlockHeaderUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_UnmarshalBinary_Summary",
		Help: "Summary of adminblockheader.UnmarshalBinary call",
	}) 
	adminBlockHeaderJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_JSONByte_Summary",
		Help: "Summary of adminblockheader.JSONByte call",
	}) 
 	adminBlockHeaderJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_JSONString_Summary",
		Help: "Summary of adminblockheader.JSONString call",
	}) 
 	adminBlockHeaderJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_JSONBuffer_Summary",
		Help: "Summary of adminblockheader.JSONBuffer call",
	}) 
 	adminBlockHeaderMarshalJSON= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_adminblockheader_MarshalJSON_Summary",
		Help: "Summary of adminblockheader.MarshalJSON call",
	}) 

   // EntryAddAuditServer.go variables
 	entryAddAuditServerString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_String_Summary",
		Help: "Summary of entryAddAuditServer.String call",
	}) 
	entryAddAuditServerUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_UpdateState_Summary",
		Help: "Summary of entryAddAuditServer.UpdateState call",
	}) 
	entryAddAuditServerNewAddAuditServer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_NewAddAuditServer_Summary",
		Help: "Summary of entryAddAuditServer.NewAddAuditServer call",
	}) 
	entryAddAuditServerType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_Type_Summary",
		Help: "Summary of entryAddAuditServer.Type call",
	}) 
	entryAddAuditServerMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_MarshalBinary_Summary",
		Help: "Summary of entryAddAuditServer.MarshalBinary call",
	}) 
	entryAddAuditServerUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_UnmarshalBinaryData_Summary",
		Help: "Summary of entryAddAuditServer.UnmarshalBinaryData call",
	})  
	entryAddAuditServerUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_UnmarshalBinary_Summary",
		Help: "Summary of entryAddAuditServer.UnmarshalBinary call",
	}) 
	entryAddAuditServerJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_JSONByte_Summary",
		Help: "Summary of entryAddAuditServer.JSONByte call",
	}) 
	entryAddAuditServerJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_JSONString_Summary",
		Help: "Summary of entryAddAuditServer.JSONString call",
	}) 
	entryAddAuditServerJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_JSONBuffer_Summary",
		Help: "Summary of entryAddAuditServer.JSONBuffer call",
	}) 
	entryAddAuditServerHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddAuditServer_Hash_Summary",
		Help: "Summary of entryAddAuditServer.Hash call",
	}) 

    // EntryAddFederatedServer.go variables
 	entryAddFederatedServerString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_String_Summary",
		Help: "Summary of entryAddFederatedServer.String call",
	}) 
	entryAddFederatedServerUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_UpdateState_Summary",
		Help: "Summary of entryAddFederatedServer.UpdateState call",
	}) 
	entryAddFederatedServerNewAddFederatedServer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_NewAddFederatedServer_Summary",
		Help: "Summary of entryAddFederatedServer.NewAddFederatedServer call",
	}) 
	entryAddFederatedServerType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_Type_Summary",
		Help: "Summary of entryAddFederatedServer.Type call",
	}) 
	entryAddFederatedServerMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_MarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServer.MarshalBinary call",
	}) 
	entryAddFederatedServerUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_UnmarshalBinaryData_Summary",
		Help: "Summary of entryAddFederatedServer.UnmarshalBinaryData call",
	})  
	entryAddFederatedServerUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_UnmarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServer.UnmarshalBinary call",
	}) 
	entryAddFederatedServerJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_JSONByte_Summary",
		Help: "Summary of entryAddFederatedServer.JSONByte call",
	}) 
	entryAddFederatedServerJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_JSONString_Summary",
		Help: "Summary of entryAddFederatedServer.JSONString call",
	}) 
	entryAddFederatedServerJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_JSONBuffer_Summary",
		Help: "Summary of entryAddFederatedServer.JSONBuffer call",
	}) 
	entryAddFederatedServerHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServer_Hash_Summary",
		Help: "Summary of entryAddFederatedServer.Hash call",
	}) 

    // EntryAddFederatedServerBitcoinAnchorKey.go variables
	entryAddFederatedServerBitcoinAnchorKeyString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_String_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.String call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_UpdateState_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.UpdateState call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyNewAddFederatedServerBitcoinAnchorKey= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_NewAddFederatedServerBitcoinAnchorKey_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.NewAddFederatedServerBitcoinAnchorKey call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_Type_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.Type call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_MarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.MarshalBinary call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_UnmarshalBinaryData_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.UnmarshalBinaryData call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_UnmarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.UnmarshalBinary call",
	}) 
 	entryAddFederatedServerBitcoinAnchorKeyJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_JSONString_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.JSONString call",
	}) 
 	entryAddFederatedServerBitcoinAnchorKeyJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_JSONBuffer_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.JSONBuffer call",
	}) 
	entryAddFederatedServerBitcoinAnchorKeyHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerBitcoinAnchorKey_Hash_Summary",
		Help: "Summary of entryAddFederatedServerBitcoinAnchorKey.Hash call",
	}) 
   
    // EntryAddFederatedServerSignerKey.go variables
	entryAddFederatedServerSigningKeyString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_String_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.String call",
	}) 
	entryAddFederatedServerSigningKeyUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_UpdateState_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.UpdateState call",
	}) 
	entryAddFederatedServerSigningKeyNewAddFederatedServerSigningKey= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.NewAddFederatedServerSigningKey call",
	}) 
	entryAddFederatedServerSigningKeyType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_Type_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.Type call",
	}) 
	entryAddFederatedServerSigningKeyMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_MarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.MarshalBinary call",
	}) 
	entryAddFederatedServerSigningKeyUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_UnmarshalBinaryData_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.UnmarshalBinaryData call",
	}) 
	entryAddFederatedServerSigningKeyUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_UnmarshalBinary_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.UnmarshalBinary call",
	}) 
 	entryAddFederatedServerSigningKeyJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_JSONByte_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.JSONSByte call",
	}) 
 	entryAddFederatedServerSigningKeyJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_JSONString_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.JSONString call",
	}) 
 	entryAddFederatedServerSigningKeyJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_JSONBuffer_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.JSONBuffer call",
	}) 
	entryAddFederatedServerSigningKeyHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddFederatedServerSigningKey_Hash_Summary",
		Help: "Summary of entryAddFederatedServerSigningKey.Hash call",
	}) 
    

    // entryDBSignature.go variables
    entryDBSignatureUpdateStateString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_String_Summary",
		Help: "Summary of entryDBSignature.String call",
	}) 
	entryDBSignatureUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_UpdateState_Summary",
		Help: "Summary of entryDBSignature.UpdateState call",
	}) 
	entryDBSignatureNewDBSignatureEntry= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryDBSignature.NewAddFederatedServerSigningKey call",
	}) 
	entryDBSignatureType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_Type_Summary",
		Help: "Summary of entryDBSignature.Type call",
	}) 
	entryDBSignatureMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_MarshalBinary_Summary",
		Help: "Summary of entryDBSignature.MarshalBinary call",
	}) 
	entryDBSignatureUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_UnmarshalBinaryData_Summary",
		Help: "Summary of entryDBSignature.UnmarshalBinaryData call",
	}) 
	entryDBSignatureUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_UnmarshalBinary_Summary",
		Help: "Summary of entryDBSignature.UnmarshalBinary call",
	}) 
 	entryDBSignatureJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_JSONByte_Summary",
		Help: "Summary of entryDBSignature.JSONSByte call",
	}) 
 	entryDBSignatureJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_JSONString_Summary",
		Help: "Summary of entryDBSignature.JSONString call",
	}) 
 	entryDBSignatureJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_JSONBuffer_Summary",
		Help: "Summary of entryDBSignature.JSONBuffer call",
	}) 
	entryDBSignatureHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryDBSignature_Hash_Summary",
		Help: "Summary of entryDBSignature.Hash call",
	}) 
    

    // EntryAddReplaceMatryoshkaHash.go variables
	entryAddReplaceMatryoshkaHashString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_String_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.String call",
	}) 
	entryAddReplaceMatryoshkaHashUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_UpdateState_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.UpdateState call",
	}) 
	entryAddReplaceMatryoshkaHashNewAddReplaceMatryoshkaHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.NewAddFederatedServerSigningKey call",
	}) 
	entryAddReplaceMatryoshkaHashType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_Type_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.Type call",
	}) 
	entryAddReplaceMatryoshkaHashMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_MarshalBinary_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.MarshalBinary call",
	}) 
	entryAddReplaceMatryoshkaHashUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_UnmarshalBinaryData_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.UnmarshalBinaryData call",
	}) 
	entryAddReplaceMatryoshkaHashUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_UnmarshalBinary_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.UnmarshalBinary call",
	}) 
 	entryAddReplaceMatryoshkaHashJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_JSONByte_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.JSONSByte call",
	}) 
 	entryAddReplaceMatryoshkaHashJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_JSONString_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.JSONString call",
	}) 
 	entryAddReplaceMatryoshkaHashJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_JSONBuffer_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.JSONBuffer call",
	}) 
	entryAddReplaceMatryoshkaHashHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryAddReplaceMatryoshkaHash_Hash_Summary",
		Help: "Summary of entryAddReplaceMatryoshkaHash.Hash call",
	}) 
    



    // entryEndOfMinute.go variables
	entryEndOfMinuteString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_String_Summary",
		Help: "Summary of entryEndOfMinute.String call",
	}) 
	entryEndOfMinuteUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_UpdateState_Summary",
		Help: "Summary of entryEndOfMinute.UpdateState call",
	}) 
	entryEndOfMinuteNewentryEndOfMinute= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryEndOfMinute.NewAddFederatedServerSigningKey call",
	}) 
	entryEndOfMinuteType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_Type_Summary",
		Help: "Summary of entryEndOfMinute.Type call",
	}) 
	entryEndOfMinuteMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_MarshalBinary_Summary",
		Help: "Summary of entryEndOfMinute.MarshalBinary call",
	}) 
	entryEndOfMinuteUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_UnmarshalBinaryData_Summary",
		Help: "Summary of entryEndOfMinute.UnmarshalBinaryData call",
	}) 
	entryEndOfMinuteUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_UnmarshalBinary_Summary",
		Help: "Summary of entryEndOfMinute.UnmarshalBinary call",
	}) 
 	entryEndOfMinuteJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_JSONByte_Summary",
		Help: "Summary of entryEndOfMinute.JSONSByte call",
	}) 
 	entryEndOfMinuteJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_JSONString_Summary",
		Help: "Summary of entryEndOfMinute.JSONString call",
	}) 
 	entryEndOfMinuteJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_JSONBuffer_Summary",
		Help: "Summary of entryEndOfMinute.JSONBuffer call",
	}) 
	entryEndOfMinuteHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryEndOfMinute_Hash_Summary",
		Help: "Summary of entryEndOfMinute.Hash call",
	}) 
        


    // entryIncreaseServerCount.go variables
	entryIncreaseServerCountString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_String_Summary",
		Help: "Summary of entryIncreaseServerCount.String call",
	}) 
	entryIncreaseServerCountUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_UpdateState_Summary",
		Help: "Summary of entryIncreaseServerCount.UpdateState call",
	}) 
	entryIncreaseServerCountNewentryIncreaseServerCount= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryIncreaseServerCount.NewAddFederatedServerSigningKey call",
	}) 
	entryIncreaseServerCountType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_Type_Summary",
		Help: "Summary of entryIncreaseServerCount.Type call",
	}) 
	entryIncreaseServerCountMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_MarshalBinary_Summary",
		Help: "Summary of entryIncreaseServerCount.MarshalBinary call",
	}) 
	entryIncreaseServerCountUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_UnmarshalBinaryData_Summary",
		Help: "Summary of entryIncreaseServerCount.UnmarshalBinaryData call",
	}) 
	entryIncreaseServerCountUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_UnmarshalBinary_Summary",
		Help: "Summary of entryIncreaseServerCount.UnmarshalBinary call",
	}) 
 	entryIncreaseServerCountJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_JSONByte_Summary",
		Help: "Summary of entryIncreaseServerCount.JSONSByte call",
	}) 
 	entryIncreaseServerCountJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_JSONString_Summary",
		Help: "Summary of entryIncreaseServerCount.JSONString call",
	}) 
 	entryIncreaseServerCountJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_JSONBuffer_Summary",
		Help: "Summary of entryIncreaseServerCount.JSONBuffer call",
	}) 
	entryIncreaseServerCountHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryIncreaseServerCount_Hash_Summary",
		Help: "Summary of entryIncreaseServerCount.Hash call",
	}) 



    // entryRemoveFederatedServer.go variables
	entryRemoveFederatedServerString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_String_Summary",
		Help: "Summary of entryRemoveFederatedServer.String call",
	}) 
	entryRemoveFederatedServerUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_UpdateState_Summary",
		Help: "Summary of entryRemoveFederatedServer.UpdateState call",
	}) 
	entryRemoveFederatedServerNewRemoveFederatedServer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryRemoveFederatedServer.NewAddFederatedServerSigningKey call",
	}) 
	entryRemoveFederatedServerType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_Type_Summary",
		Help: "Summary of entryRemoveFederatedServer.Type call",
	}) 
	entryRemoveFederatedServerMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_MarshalBinary_Summary",
		Help: "Summary of entryRemoveFederatedServer.MarshalBinary call",
	}) 
	entryRemoveFederatedServerUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_UnmarshalBinaryData_Summary",
		Help: "Summary of entryRemoveFederatedServer.UnmarshalBinaryData call",
	}) 
	entryRemoveFederatedServerUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_UnmarshalBinary_Summary",
		Help: "Summary of entryRemoveFederatedServer.UnmarshalBinary call",
	}) 
 	entryRemoveFederatedServerJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_JSONByte_Summary",
		Help: "Summary of entryRemoveFederatedServer.JSONSByte call",
	}) 
 	entryRemoveFederatedServerJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_JSONString_Summary",
		Help: "Summary of entryRemoveFederatedServer.JSONString call",
	}) 
 	entryRemoveFederatedServerJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_JSONBuffer_Summary",
		Help: "Summary of entryRemoveFederatedServer.JSONBuffer call",
	}) 
	entryRemoveFederatedServerHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRemoveFederatedServer_Hash_Summary",
		Help: "Summary of entryRemoveFederatedServer.Hash call",
	}) 
           


    // entryRevealMatryoshkaHash.go variables
	entryRevealMatryoshkaHashString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_String_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.String call",
	}) 
	
    entryRevealMatryoshkaHashNewRevealMatryoshkaHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.NewAddFederatedServerSigningKey call",
	}) 
	entryRevealMatryoshkaHashType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_Type_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.Type call",
	}) 
	entryRevealMatryoshkaHashMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_MarshalBinary_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.MarshalBinary call",
	}) 
	entryRevealMatryoshkaHashUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_UnmarshalBinaryData_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.UnmarshalBinaryData call",
	}) 
	entryRevealMatryoshkaHashUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_UnmarshalBinary_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.UnmarshalBinary call",
	}) 
 	entryRevealMatryoshkaHashJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_JSONByte_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.JSONSByte call",
	}) 
 	entryRevealMatryoshkaHashJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_JSONString_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.JSONString call",
	}) 
 	entryRevealMatryoshkaHashJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_JSONBuffer_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.JSONBuffer call",
	}) 
	entryRevealMatryoshkaHashHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryRevealMatryoshkaHash_Hash_Summary",
		Help: "Summary of entryRevealMatryoshkaHash.Hash call",
	}) 




    // entryServerFault.go variables
	entryServerFaultString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_String_Summary",
		Help: "Summary of entryServerFault.String call",
	}) 
	entryServerFaultUpdateState= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_UpdateState_Summary",
		Help: "Summary of entryServerFault.UpdateState call",
	}) 
    entryServerFaultNewServerFault= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of entryServerFault.NewAddFederatedServerSigningKey call",
	}) 
	entryServerFaultType= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_Type_Summary",
		Help: "Summary of entryServerFault.Type call",
	}) 
	entryServerFaultMarshalCore= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_MarshalCore_Summary",
		Help: "Summary of entryServerFault.MarshalCore call",
	}) 
	entryServerFaultMarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_MarshalBinary_Summary",
		Help: "Summary of entryServerFault.MarshalBinary call",
	}) 
	entryServerFaultUnmarshalBinaryData= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of entryServerFault.UnmarshalBinaryData call",
	}) 
	entryServerFaultUnmarshalBinary= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_UnmarshalBinary_Summary",
		Help: "Summary of entryServerFault.UnmarshalBinary call",
	}) 
 	entryServerFaultJSONByte= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_JSONByte_Summary",
		Help: "Summary of entryServerFault.JSONSByte call",
	}) 
 	entryServerFaultJSONString= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_JSONString_Summary",
		Help: "Summary of entryServerFault.JSONString call",
	}) 
 	entryServerFaultJSONBuffer= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_JSONBuffer_Summary",
		Help: "Summary of entryServerFault.JSONBuffer call",
	}) 
	entryServerFaultHash= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_Hash_Summary",
		Help: "Summary of entryServerFault.Hash call",
	}) 
    entryServerFaultCompare= prometheus.NewSummary(prometheus.CounterOpts{
		Name: "factomd_adminblock_entryServerFault_Compare_Summary",
		Help: "Summary of entryServerFault.Compare call",
	}) 



    

       
)



func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(AnchorErrors)
	//adminBlock.go variables
    prometheus.MustRegister(adminBlockErrors) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockString) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockUpdateState) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddDBSig) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddFedServer) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddAuditServer) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockRemoveFederatedServer) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddMatryoshkaHash) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddFederatedServerSigningKey) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddFederatedServerBitcoinAnchorKey) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddEntry) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddServerFault) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetHeader) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockSetHeader) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetABEntries) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetDBHeight) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockSetABEntries) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockNew) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetDatabaseHeight) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetChainID) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockDatabasePrimaryIndex) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockDatabaseSecondaryIndex) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetHash) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetKeyMR) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockBackReferenceHash) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockLookupHash) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddABEntry) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockAddFirstABEntry) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockMarshalBinary) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockUnmarshalABlock) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockUnmarshalBinaryData) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockUnmarshalBinary) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockGetDBSignature) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockJSONByte) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockJSONString) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockJSONBuffer) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockMarshalJSON) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockNewAdminBlock) = prometheus.NewCounter(prometheus.CounterOpts{
	prometheus.MustRegister(adminBlockCheckBlockPairIntegrity) 

    // adminBlockHeader.go variables
    prometheus.MustRegister(adminBlockHeaderString)
    prometheus.MustRegister(adminBlockHeaderGetMessageCount)
    prometheus.MustRegister(adminBlockHeaderSetMessageCount)
    prometheus.MustRegister(adminBlockHeaderGetBodySize)
    prometheus.MustRegister(adminBlockHeaderSetBodySize)
    prometheus.MustRegister(adminBlockHeaderGetAdminChainID)
    prometheus.MustRegister(adminBlockHeaderGetDBHeight)
    prometheus.MustRegister(adminBlockHeaderGetHeaderExpansionArea)
    prometheus.MustRegister(adminBlockHeaderGetHeaderExpansionArea)
    prometheus.MustRegister(adminBlockHeaderGetHeaderExpansionSize)
    prometheus.MustRegister(adminBlockHeaderGetPrevBackRefHash)
    prometheus.MustRegister(adminBlockHeaderSetDBHeight)
    prometheus.MustRegister(adminBlockHeaderSetHeaderExpansionArea)
    prometheus.MustRegister(adminBlockHeaderSetPrevBackRefHash)
    prometheus.MustRegister(adminBlockHeaderMarshalBinary)
    prometheus.MustRegister(adminBlockHeaderUnmarshalBinaryData)
    prometheus.MustRegister(adminBlockHeaderUnmarshalBinary)
    prometheus.MustRegister(adminBlockHeaderJSONByte)
    prometheus.MustRegister(adminBlockHeaderJSONString)
    prometheus.MustRegister(adminBlockHeaderJSONBuffer)
    prometheus.MustRegister(adminBlockHeaderMarshalJSON)

   // EntryAddAuditServer.go variables
    prometheus.MustRegister(entryAddAuditServerString)
    prometheus.MustRegister(entryAddAuditServerUpdateState)
    prometheus.MustRegister(entryAddAuditServerNewAddAuditServer)
    prometheus.MustRegister(entryAddAuditServerType)
    prometheus.MustRegister(entryAddAuditServerMarshalBinary)
    prometheus.MustRegister(entryAddAuditServerUnmarshalBinaryData)
    prometheus.MustRegister(entryAddAuditServerUnmarshalBinary)
    prometheus.MustRegister(entryAddAuditServerJSONByte)
    prometheus.MustRegister(entryAddAuditServerJSONString)
    prometheus.MustRegister(entryAddAuditServerJSONBuffer)
    prometheus.MustRegister(entryAddAuditServerHash)

    // EntryAddFederatedServer.go variables)
    prometheus.MustRegister(entryAddFederatedServerString)
    prometheus.MustRegister(entryAddFederatedServerUpdateState)
    prometheus.MustRegister(entryAddFederatedServerNewAddFederatedServer)
    prometheus.MustRegister(entryAddFederatedServerType)
    prometheus.MustRegister(entryAddFederatedServerMarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerUnmarshalBinaryData)
    prometheus.MustRegister(entryAddFederatedServerUnmarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerJSONByte)
    prometheus.MustRegister(entryAddFederatedServerJSONString)
    prometheus.MustRegister(entryAddFederatedServerJSONBuffer)
    prometheus.MustRegister(entryAddFederatedServerHash)

    // EntryAddFederatedServerBitcoinAnchorKey.go variables)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyString)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyUpdateState)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyNewAddFederatedServerBitcoinAnchorKey)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyType)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyMarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyUnmarshalBinaryData)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyUnmarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyJSONString)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyJSONBuffer)
    prometheus.MustRegister(entryAddFederatedServerBitcoinAnchorKeyHash)

    // EntryAddFederatedServerSignerKey.go variables)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyString)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyUpdateState)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyNewAddFederatedServerSigningKey)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyType)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyMarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyUnmarshalBinaryData)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyUnmarshalBinary)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyJSONByte)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyJSONString)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyJSONBuffer)
    prometheus.MustRegister(entryAddFederatedServerSigningKeyHash) 
    

    // entryDBSignature.go variables)
    prometheus.MustRegister(entryDBSignatureUpdateStateString)
    prometheus.MustRegister(entryDBSignatureUpdateState)
    prometheus.MustRegister(entryDBSignatureNewDBSignatureEntry)
    prometheus.MustRegister(entryDBSignatureType)
    prometheus.MustRegister(entryDBSignatureMarshalBinary)
    prometheus.MustRegister(entryDBSignatureUnmarshalBinaryData)
    prometheus.MustRegister(entryDBSignatureUnmarshalBinary)
    prometheus.MustRegister(entryDBSignatureJSONByte)
    prometheus.MustRegister(entryDBSignatureJSONString)
    prometheus.MustRegister(entryDBSignatureJSONBuffer)
    prometheus.MustRegister(entryDBSignatureHash) 
    

    // EntryAddReplaceMatryoshkaHash.go variables)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashString)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashUpdateState)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashNewAddReplaceMatryoshkaHash)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashType)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashMarshalBinary)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashUnmarshalBinaryData)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashUnmarshalBinary)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashJSONByte)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashJSONString)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashJSONBuffer)
    prometheus.MustRegister(entryAddReplaceMatryoshkaHashHash) 
    



    // entryEndOfMinute.go variables)
    prometheus.MustRegister(entryEndOfMinuteString)
    prometheus.MustRegister(entryEndOfMinuteUpdateState)
    prometheus.MustRegister(entryEndOfMinuteNewentryEndOfMinute)
    prometheus.MustRegister(entryEndOfMinuteType)
    prometheus.MustRegister(entryEndOfMinuteMarshalBinary)
    prometheus.MustRegister(entryEndOfMinuteUnmarshalBinaryData)
    prometheus.MustRegister(entryEndOfMinuteUnmarshalBinary)
    prometheus.MustRegister(entryEndOfMinuteJSONByte)
    prometheus.MustRegister(entryEndOfMinuteJSONString)
    prometheus.MustRegister(entryEndOfMinuteJSONBuffer)
    prometheus.MustRegister(entryEndOfMinuteHash) 
        


    // entryIncreaseServerCount.go variables)
    prometheus.MustRegister(entryIncreaseServerCountString)
    prometheus.MustRegister(entryIncreaseServerCountUpdateState)
    prometheus.MustRegister(entryIncreaseServerCountNewentryIncreaseServerCount)
    prometheus.MustRegister(entryIncreaseServerCountType)
    prometheus.MustRegister(entryIncreaseServerCountMarshalBinary)
    prometheus.MustRegister(entryIncreaseServerCountUnmarshalBinaryData)
    prometheus.MustRegister(entryIncreaseServerCountUnmarshalBinary)
    prometheus.MustRegister(entryIncreaseServerCountJSONByte)
    prometheus.MustRegister(entryIncreaseServerCountJSONString)
    prometheus.MustRegister(entryIncreaseServerCountJSONBuffer)
    prometheus.MustRegister(entryIncreaseServerCountHash) 



    // entryRemoveFederatedServer.go variables)
    prometheus.MustRegister(entryRemoveFederatedServerString)
    prometheus.MustRegister(entryRemoveFederatedServerUpdateState)
    prometheus.MustRegister(entryRemoveFederatedServerNewRemoveFederatedServer)
    prometheus.MustRegister(entryRemoveFederatedServerType)
    prometheus.MustRegister(entryRemoveFederatedServerMarshalBinary)
    prometheus.MustRegister(entryRemoveFederatedServerUnmarshalBinaryData)
    prometheus.MustRegister(entryRemoveFederatedServerUnmarshalBinary)
    prometheus.MustRegister(entryRemoveFederatedServerJSONByte)
    prometheus.MustRegister(entryRemoveFederatedServerJSONString)
    prometheus.MustRegister(entryRemoveFederatedServerJSONBuffer)
    prometheus.MustRegister(entryRemoveFederatedServerHash) 
           


    // entryRevealMatryoshkaHash.go variables)
    prometheus.MustRegister(entryRevealMatryoshkaHashString)
    prometheus.MustRegister(entryRevealMatryoshkaHashNewRevealMatryoshkaHash)
    prometheus.MustRegister(entryRevealMatryoshkaHashType)
    prometheus.MustRegister(entryRevealMatryoshkaHashMarshalBinary)
    prometheus.MustRegister(entryRevealMatryoshkaHashUnmarshalBinaryData)
    prometheus.MustRegister(entryRevealMatryoshkaHashUnmarshalBinary)
    prometheus.MustRegister(entryRevealMatryoshkaHashJSONByte)
    prometheus.MustRegister(entryRevealMatryoshkaHashJSONString)
    prometheus.MustRegister(entryRevealMatryoshkaHashJSONBuffer)
    prometheus.MustRegister(entryRevealMatryoshkaHashHash) 




    // entryServerFault.go variables)
    prometheus.MustRegister(entryServerFaultString)
    prometheus.MustRegister(entryServerFaultUpdateState)
    prometheus.MustRegister(entryServerFaultNewServerFault)
    prometheus.MustRegister(entryServerFaultType)
    prometheus.MustRegister(entryServerFaultMarshalCore)
    prometheus.MustRegister(entryServerFaultMarshalBinary)
    prometheus.MustRegister(entryServerFaultUnmarshalBinaryData)
    prometheus.MustRegister(entryServerFaultUnmarshalBinary)
    prometheus.MustRegister(entryServerFaultJSONByte)
    prometheus.MustRegister(entryServerFaultJSONString)
    prometheus.MustRegister(entryServerFaultJSONBuffer)
    prometheus.MustRegister(entryServerFaultHash)
    prometheus.MustRegister(entryServerFaultCompare) 



/* these are what is being added to each function to measure function time
	callTime := time.Now().UnixNano()
	defer adminBlockString.Observe(float64(time.Now().UnixNano() - callTime))
*/


}