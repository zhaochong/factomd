package adminBlock

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdadminBlockAddAuditServerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_Init",
	})
	factomdadminBlockAddAuditServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_String",
	})
	factomdadminBlockAddAuditServerUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_UpdateState",
	})
	factomdadminBlockNewAddAuditServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAddAuditServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAddAuditServer",
	})
	factomdadminBlockAddAuditServerType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_Type",
	})
	factomdadminBlockAddAuditServerMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_MarshalBinary",
	})
	factomdadminBlockAddAuditServerUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_UnmarshalBinaryData",
	})
	factomdadminBlockAddAuditServerUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_UnmarshalBinary",
	})
	factomdadminBlockAddAuditServerJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_JSONByte",
	})
	factomdadminBlockAddAuditServerJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_JSONString",
	})
	factomdadminBlockAddAuditServerIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_IsInterpretable",
	})
	factomdadminBlockAddAuditServerInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_Interpret",
	})
	factomdadminBlockAddAuditServerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddAuditServer_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddAuditServer_Hash",
	})
	factomdadminBlockAddFederatedServerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_Init",
	})
	factomdadminBlockAddFederatedServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_String",
	})
	factomdadminBlockAddFederatedServerUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_UpdateState",
	})
	factomdadminBlockNewAddFederatedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAddFederatedServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAddFederatedServer",
	})
	factomdadminBlockAddFederatedServerType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_Type",
	})
	factomdadminBlockAddFederatedServerMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_MarshalBinary",
	})
	factomdadminBlockAddFederatedServerUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_UnmarshalBinaryData",
	})
	factomdadminBlockAddFederatedServerUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_UnmarshalBinary",
	})
	factomdadminBlockAddFederatedServerJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_JSONByte",
	})
	factomdadminBlockAddFederatedServerJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_JSONString",
	})
	factomdadminBlockAddFederatedServerIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_IsInterpretable",
	})
	factomdadminBlockAddFederatedServerInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_Interpret",
	})
	factomdadminBlockAddFederatedServerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServer_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServer_Hash",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Init",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_String",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UpdateState",
	})
	factomdadminBlockNewAddFederatedServerBitcoinAnchorKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAddFederatedServerBitcoinAnchorKey_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAddFederatedServerBitcoinAnchorKey",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Type",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_MarshalBinary",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UnmarshalBinaryData",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_UnmarshalBinary",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_JSONByte",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_JSONString",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_IsInterpretable",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Interpret",
	})
	factomdadminBlockAddFederatedServerBitcoinAnchorKeyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerBitcoinAnchorKey_Hash",
	})
	factomdadminBlockAddFederatedServerSigningKeyInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_Init",
	})
	factomdadminBlockAddFederatedServerSigningKeyUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_UpdateState",
	})
	factomdadminBlockAddFederatedServerSigningKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_String",
	})
	factomdadminBlockNewAddFederatedServerSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAddFederatedServerSigningKey_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAddFederatedServerSigningKey",
	})
	factomdadminBlockAddFederatedServerSigningKeyType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_Type",
	})
	factomdadminBlockAddFederatedServerSigningKeyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_MarshalBinary",
	})
	factomdadminBlockAddFederatedServerSigningKeyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_UnmarshalBinaryData",
	})
	factomdadminBlockAddFederatedServerSigningKeyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_UnmarshalBinary",
	})
	factomdadminBlockAddFederatedServerSigningKeyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_JSONByte",
	})
	factomdadminBlockAddFederatedServerSigningKeyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_JSONString",
	})
	factomdadminBlockAddFederatedServerSigningKeyIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_IsInterpretable",
	})
	factomdadminBlockAddFederatedServerSigningKeyInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_Interpret",
	})
	factomdadminBlockAddFederatedServerSigningKeyHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddFederatedServerSigningKey_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddFederatedServerSigningKey_Hash",
	})
	factomdadminBlockAddReplaceMatryoshkaHashInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_Init",
	})
	factomdadminBlockAddReplaceMatryoshkaHashString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_String",
	})
	factomdadminBlockAddReplaceMatryoshkaHashType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_Type",
	})
	factomdadminBlockAddReplaceMatryoshkaHashUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_UpdateState",
	})
	factomdadminBlockNewAddReplaceMatryoshkaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAddReplaceMatryoshkaHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAddReplaceMatryoshkaHash",
	})
	factomdadminBlockAddReplaceMatryoshkaHashMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_MarshalBinary",
	})
	factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_UnmarshalBinaryData",
	})
	factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_UnmarshalBinary",
	})
	factomdadminBlockAddReplaceMatryoshkaHashJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_JSONByte",
	})
	factomdadminBlockAddReplaceMatryoshkaHashJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_JSONString",
	})
	factomdadminBlockAddReplaceMatryoshkaHashIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_IsInterpretable",
	})
	factomdadminBlockAddReplaceMatryoshkaHashInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_Interpret",
	})
	factomdadminBlockAddReplaceMatryoshkaHashHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AddReplaceMatryoshkaHash_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AddReplaceMatryoshkaHash_Hash",
	})
	factomdadminBlockDBSignatureEntryInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_Init",
	})
	factomdadminBlockDBSignatureEntryUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_UpdateState",
	})
	factomdadminBlockNewDBSignatureEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewDBSignatureEntry_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewDBSignatureEntry",
	})
	factomdadminBlockDBSignatureEntryType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_Type",
	})
	factomdadminBlockDBSignatureEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_MarshalBinary",
	})
	factomdadminBlockDBSignatureEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_UnmarshalBinaryData",
	})
	factomdadminBlockDBSignatureEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_UnmarshalBinary",
	})
	factomdadminBlockDBSignatureEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_JSONByte",
	})
	factomdadminBlockDBSignatureEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_JSONString",
	})
	factomdadminBlockDBSignatureEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_String",
	})
	factomdadminBlockDBSignatureEntryIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_IsInterpretable",
	})
	factomdadminBlockDBSignatureEntryInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_Interpret",
	})
	factomdadminBlockDBSignatureEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_DBSignatureEntry_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_DBSignatureEntry_Hash",
	})
	factomdadminBlockEndOfMinuteEntryType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_Type",
	})
	factomdadminBlockEndOfMinuteEntryUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_UpdateState",
	})
	factomdadminBlockNewEndOfMinuteEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewEndOfMinuteEntry_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewEndOfMinuteEntry",
	})
	factomdadminBlockEndOfMinuteEntryMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_MarshalBinary",
	})
	factomdadminBlockEndOfMinuteEntryUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_UnmarshalBinaryData",
	})
	factomdadminBlockEndOfMinuteEntryUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_UnmarshalBinary",
	})
	factomdadminBlockEndOfMinuteEntryJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_JSONByte",
	})
	factomdadminBlockEndOfMinuteEntryJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_JSONString",
	})
	factomdadminBlockEndOfMinuteEntryString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_String",
	})
	factomdadminBlockEndOfMinuteEntryIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_IsInterpretable",
	})
	factomdadminBlockEndOfMinuteEntryInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_Interpret",
	})
	factomdadminBlockEndOfMinuteEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_EndOfMinuteEntry_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_EndOfMinuteEntry_Hash",
	})
	factomdadminBlockNewIncreaseSererCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewIncreaseSererCount_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewIncreaseSererCount",
	})
	factomdadminBlockIncreaseServerCountUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_UpdateState",
	})
	factomdadminBlockIncreaseServerCountType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_Type",
	})
	factomdadminBlockIncreaseServerCountMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_MarshalBinary",
	})
	factomdadminBlockIncreaseServerCountUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_UnmarshalBinaryData",
	})
	factomdadminBlockIncreaseServerCountUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_UnmarshalBinary",
	})
	factomdadminBlockIncreaseServerCountJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_JSONByte",
	})
	factomdadminBlockIncreaseServerCountJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_JSONString",
	})
	factomdadminBlockIncreaseServerCountString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_String",
	})
	factomdadminBlockIncreaseServerCountIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_IsInterpretable",
	})
	factomdadminBlockIncreaseServerCountInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_Interpret",
	})
	factomdadminBlockIncreaseServerCountHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_IncreaseServerCount_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_IncreaseServerCount_Hash",
	})
	factomdadminBlockRemoveFederatedServerInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_Init",
	})
	factomdadminBlockRemoveFederatedServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_String",
	})
	factomdadminBlockRemoveFederatedServerUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_UpdateState",
	})
	factomdadminBlockNewRemoveFederatedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewRemoveFederatedServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewRemoveFederatedServer",
	})
	factomdadminBlockRemoveFederatedServerType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_Type",
	})
	factomdadminBlockRemoveFederatedServerMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_MarshalBinary",
	})
	factomdadminBlockRemoveFederatedServerUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_UnmarshalBinaryData",
	})
	factomdadminBlockRemoveFederatedServerUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_UnmarshalBinary",
	})
	factomdadminBlockRemoveFederatedServerJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_JSONByte",
	})
	factomdadminBlockRemoveFederatedServerJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_JSONString",
	})
	factomdadminBlockRemoveFederatedServerIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_IsInterpretable",
	})
	factomdadminBlockRemoveFederatedServerInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_Interpret",
	})
	factomdadminBlockRemoveFederatedServerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RemoveFederatedServer_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RemoveFederatedServer_Hash",
	})
	factomdadminBlockRevealMatryoshkaHashInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_Init",
	})
	factomdadminBlockRevealMatryoshkaHashType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_Type",
	})
	factomdadminBlockNewRevealMatryoshkaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewRevealMatryoshkaHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewRevealMatryoshkaHash",
	})
	factomdadminBlockRevealMatryoshkaHashUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_UpdateState",
	})
	factomdadminBlockRevealMatryoshkaHashMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_MarshalBinary",
	})
	factomdadminBlockRevealMatryoshkaHashUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_UnmarshalBinaryData",
	})
	factomdadminBlockRevealMatryoshkaHashUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_UnmarshalBinary",
	})
	factomdadminBlockRevealMatryoshkaHashJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_JSONByte",
	})
	factomdadminBlockRevealMatryoshkaHashJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_JSONString",
	})
	factomdadminBlockRevealMatryoshkaHashString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_String",
	})
	factomdadminBlockRevealMatryoshkaHashIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_IsInterpretable",
	})
	factomdadminBlockRevealMatryoshkaHashInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_Interpret",
	})
	factomdadminBlockRevealMatryoshkaHashHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_RevealMatryoshkaHash_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_RevealMatryoshkaHash_Hash",
	})
	factomdadminBlockServerFaultInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_Init",
	})
	factomdadminBlockSigListMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_SigList_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_SigList_MarshalBinary",
	})
	factomdadminBlockSigListUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_SigList_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_SigList_UnmarshalBinaryData",
	})
	factomdadminBlockServerFaultUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_UpdateState",
	})
	factomdadminBlockServerFaultMarshalCore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_MarshalCore_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_MarshalCore",
	})
	factomdadminBlockServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_MarshalBinary",
	})
	factomdadminBlockServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_UnmarshalBinaryData",
	})
	factomdadminBlockServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_UnmarshalBinary",
	})
	factomdadminBlockServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_JSONByte",
	})
	factomdadminBlockServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_JSONString",
	})
	factomdadminBlockServerFaultIsInterpretable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_IsInterpretable_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_IsInterpretable",
	})
	factomdadminBlockServerFaultInterpret = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_Interpret_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_Interpret",
	})
	factomdadminBlockServerFaultHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_Hash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_Hash",
	})
	factomdadminBlockServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_String",
	})
	factomdadminBlockServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_Type_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_Type",
	})
	factomdadminBlockServerFaultCompare = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ServerFault_Compare_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ServerFault_Compare",
	})
	factomdadminBlockAdminBlockInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_Init",
	})
	factomdadminBlockAdminBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_String",
	})
	factomdadminBlockAdminBlockUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_UpdateState_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_UpdateState",
	})
	factomdadminBlockAdminBlockAddDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddDBSig_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddDBSig",
	})
	factomdadminBlockAdminBlockAddFedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddFedServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddFedServer",
	})
	factomdadminBlockAdminBlockAddAuditServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddAuditServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddAuditServer",
	})
	factomdadminBlockAdminBlockRemoveFederatedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_RemoveFederatedServer_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_RemoveFederatedServer",
	})
	factomdadminBlockAdminBlockAddMatryoshkaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddMatryoshkaHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddMatryoshkaHash",
	})
	factomdadminBlockAdminBlockAddFederatedServerSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddFederatedServerSigningKey_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddFederatedServerSigningKey",
	})
	factomdadminBlockAdminBlockAddFederatedServerBitcoinAnchorKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddFederatedServerBitcoinAnchorKey_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddFederatedServerBitcoinAnchorKey",
	})
	factomdadminBlockAdminBlockAddEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddEntry_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddEntry",
	})
	factomdadminBlockAdminBlockAddServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddServerFault_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddServerFault",
	})
	factomdadminBlockAdminBlockGetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetHeader_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetHeader",
	})
	factomdadminBlockAdminBlockSetHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_SetHeader_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_SetHeader",
	})
	factomdadminBlockAdminBlockGetABEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetABEntries_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetABEntries",
	})
	factomdadminBlockAdminBlockGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetDBHeight",
	})
	factomdadminBlockAdminBlockSetABEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_SetABEntries_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_SetABEntries",
	})
	factomdadminBlockAdminBlockNew = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_New_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_New",
	})
	factomdadminBlockAdminBlockGetDatabaseHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetDatabaseHeight_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetDatabaseHeight",
	})
	factomdadminBlockAdminBlockGetChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetChainID_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetChainID",
	})
	factomdadminBlockAdminBlockDatabasePrimaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_DatabasePrimaryIndex_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_DatabasePrimaryIndex",
	})
	factomdadminBlockAdminBlockDatabaseSecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_DatabaseSecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_DatabaseSecondaryIndex",
	})
	factomdadminBlockAdminBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetHash",
	})
	factomdadminBlockAdminBlockGetKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetKeyMR_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetKeyMR",
	})
	factomdadminBlockAdminBlockBackReferenceHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_BackReferenceHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_BackReferenceHash",
	})
	factomdadminBlockAdminBlockLookupHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_LookupHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_LookupHash",
	})
	factomdadminBlockAdminBlockAddABEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddABEntry_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddABEntry",
	})
	factomdadminBlockAdminBlockAddFirstABEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_AddFirstABEntry_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_AddFirstABEntry",
	})
	factomdadminBlockAdminBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_MarshalBinary",
	})
	factomdadminBlockUnmarshalABlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_UnmarshalABlock_Summary",
		Help: "Summary of calls to  factomd_adminBlock_UnmarshalABlock",
	})
	factomdadminBlockAdminBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_UnmarshalBinaryData",
	})
	factomdadminBlockAdminBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_UnmarshalBinary",
	})
	factomdadminBlockAdminBlockGetDBSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_GetDBSignature_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_GetDBSignature",
	})
	factomdadminBlockAdminBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_JSONByte",
	})
	factomdadminBlockAdminBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_JSONString",
	})
	factomdadminBlockAdminBlockMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_AdminBlock_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_adminBlock_AdminBlock_MarshalJSON",
	})
	factomdadminBlockNewAdminBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_NewAdminBlock_Summary",
		Help: "Summary of calls to  factomd_adminBlock_NewAdminBlock",
	})
	factomdadminBlockCheckBlockPairIntegrity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_CheckBlockPairIntegrity_Summary",
		Help: "Summary of calls to  factomd_adminBlock_CheckBlockPairIntegrity",
	})
	factomdadminBlockABlockHeaderInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_Init_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_Init",
	})
	factomdadminBlockABlockHeaderString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_String_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_String",
	})
	factomdadminBlockABlockHeaderGetMessageCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetMessageCount_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetMessageCount",
	})
	factomdadminBlockABlockHeaderSetMessageCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_SetMessageCount_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_SetMessageCount",
	})
	factomdadminBlockABlockHeaderGetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetBodySize_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetBodySize",
	})
	factomdadminBlockABlockHeaderSetBodySize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_SetBodySize_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_SetBodySize",
	})
	factomdadminBlockABlockHeaderGetAdminChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetAdminChainID_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetAdminChainID",
	})
	factomdadminBlockABlockHeaderGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetDBHeight",
	})
	factomdadminBlockABlockHeaderGetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetHeaderExpansionArea_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetHeaderExpansionArea",
	})
	factomdadminBlockABlockHeaderGetHeaderExpansionSize = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetHeaderExpansionSize_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetHeaderExpansionSize",
	})
	factomdadminBlockABlockHeaderGetPrevBackRefHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_GetPrevBackRefHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_GetPrevBackRefHash",
	})
	factomdadminBlockABlockHeaderSetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_SetDBHeight_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_SetDBHeight",
	})
	factomdadminBlockABlockHeaderSetHeaderExpansionArea = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_SetHeaderExpansionArea_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_SetHeaderExpansionArea",
	})
	factomdadminBlockABlockHeaderSetPrevBackRefHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_SetPrevBackRefHash_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_SetPrevBackRefHash",
	})
	factomdadminBlockABlockHeaderMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_MarshalBinary",
	})
	factomdadminBlockABlockHeaderUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_UnmarshalBinaryData",
	})
	factomdadminBlockABlockHeaderUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_UnmarshalBinary",
	})
	factomdadminBlockABlockHeaderJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_JSONByte_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_JSONByte",
	})
	factomdadminBlockABlockHeaderJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_JSONString_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_JSONString",
	})
	factomdadminBlockABlockHeaderMarshalJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_adminBlock_ABlockHeader_MarshalJSON_Summary",
		Help: "Summary of calls to  factomd_adminBlock_ABlockHeader_MarshalJSON",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdadminBlockAddAuditServerInit)
	prometheus.MustRegister(factomdadminBlockAddAuditServerString)
	prometheus.MustRegister(factomdadminBlockAddAuditServerUpdateState)
	prometheus.MustRegister(factomdadminBlockNewAddAuditServer)
	prometheus.MustRegister(factomdadminBlockAddAuditServerType)
	prometheus.MustRegister(factomdadminBlockAddAuditServerMarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddAuditServerUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAddAuditServerUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddAuditServerJSONByte)
	prometheus.MustRegister(factomdadminBlockAddAuditServerJSONString)
	prometheus.MustRegister(factomdadminBlockAddAuditServerIsInterpretable)
	prometheus.MustRegister(factomdadminBlockAddAuditServerInterpret)
	prometheus.MustRegister(factomdadminBlockAddAuditServerHash)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerInit)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerString)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerUpdateState)
	prometheus.MustRegister(factomdadminBlockNewAddFederatedServer)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerType)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerMarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerJSONByte)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerJSONString)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerIsInterpretable)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerInterpret)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerHash)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyInit)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyString)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyUpdateState)
	prometheus.MustRegister(factomdadminBlockNewAddFederatedServerBitcoinAnchorKey)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyType)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyMarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyJSONByte)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyJSONString)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyIsInterpretable)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyInterpret)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerBitcoinAnchorKeyHash)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyInit)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyUpdateState)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyString)
	prometheus.MustRegister(factomdadminBlockNewAddFederatedServerSigningKey)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyType)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyMarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyJSONByte)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyJSONString)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyIsInterpretable)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyInterpret)
	prometheus.MustRegister(factomdadminBlockAddFederatedServerSigningKeyHash)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashInit)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashString)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashType)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashUpdateState)
	prometheus.MustRegister(factomdadminBlockNewAddReplaceMatryoshkaHash)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashMarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashJSONByte)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashJSONString)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashIsInterpretable)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashInterpret)
	prometheus.MustRegister(factomdadminBlockAddReplaceMatryoshkaHashHash)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryInit)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryUpdateState)
	prometheus.MustRegister(factomdadminBlockNewDBSignatureEntry)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryType)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryMarshalBinary)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryJSONByte)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryJSONString)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryString)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryIsInterpretable)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryInterpret)
	prometheus.MustRegister(factomdadminBlockDBSignatureEntryHash)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryType)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryUpdateState)
	prometheus.MustRegister(factomdadminBlockNewEndOfMinuteEntry)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryMarshalBinary)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryJSONByte)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryJSONString)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryString)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryIsInterpretable)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryInterpret)
	prometheus.MustRegister(factomdadminBlockEndOfMinuteEntryHash)
	prometheus.MustRegister(factomdadminBlockNewIncreaseSererCount)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountUpdateState)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountType)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountMarshalBinary)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountJSONByte)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountJSONString)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountString)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountIsInterpretable)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountInterpret)
	prometheus.MustRegister(factomdadminBlockIncreaseServerCountHash)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerInit)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerString)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerUpdateState)
	prometheus.MustRegister(factomdadminBlockNewRemoveFederatedServer)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerType)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerMarshalBinary)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerJSONByte)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerJSONString)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerIsInterpretable)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerInterpret)
	prometheus.MustRegister(factomdadminBlockRemoveFederatedServerHash)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashInit)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashType)
	prometheus.MustRegister(factomdadminBlockNewRevealMatryoshkaHash)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashUpdateState)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashMarshalBinary)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashJSONByte)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashJSONString)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashString)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashIsInterpretable)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashInterpret)
	prometheus.MustRegister(factomdadminBlockRevealMatryoshkaHashHash)
	prometheus.MustRegister(factomdadminBlockServerFaultInit)
	prometheus.MustRegister(factomdadminBlockSigListMarshalBinary)
	prometheus.MustRegister(factomdadminBlockSigListUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockServerFaultUpdateState)
	prometheus.MustRegister(factomdadminBlockServerFaultMarshalCore)
	prometheus.MustRegister(factomdadminBlockServerFaultMarshalBinary)
	prometheus.MustRegister(factomdadminBlockServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockServerFaultUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockServerFaultJSONByte)
	prometheus.MustRegister(factomdadminBlockServerFaultJSONString)
	prometheus.MustRegister(factomdadminBlockServerFaultIsInterpretable)
	prometheus.MustRegister(factomdadminBlockServerFaultInterpret)
	prometheus.MustRegister(factomdadminBlockServerFaultHash)
	prometheus.MustRegister(factomdadminBlockServerFaultString)
	prometheus.MustRegister(factomdadminBlockServerFaultType)
	prometheus.MustRegister(factomdadminBlockServerFaultCompare)
	prometheus.MustRegister(factomdadminBlockAdminBlockInit)
	prometheus.MustRegister(factomdadminBlockAdminBlockString)
	prometheus.MustRegister(factomdadminBlockAdminBlockUpdateState)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddDBSig)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddFedServer)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddAuditServer)
	prometheus.MustRegister(factomdadminBlockAdminBlockRemoveFederatedServer)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddMatryoshkaHash)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddFederatedServerSigningKey)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddFederatedServerBitcoinAnchorKey)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddEntry)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddServerFault)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetHeader)
	prometheus.MustRegister(factomdadminBlockAdminBlockSetHeader)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetABEntries)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetDBHeight)
	prometheus.MustRegister(factomdadminBlockAdminBlockSetABEntries)
	prometheus.MustRegister(factomdadminBlockAdminBlockNew)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetDatabaseHeight)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetChainID)
	prometheus.MustRegister(factomdadminBlockAdminBlockDatabasePrimaryIndex)
	prometheus.MustRegister(factomdadminBlockAdminBlockDatabaseSecondaryIndex)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetHash)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetKeyMR)
	prometheus.MustRegister(factomdadminBlockAdminBlockBackReferenceHash)
	prometheus.MustRegister(factomdadminBlockAdminBlockLookupHash)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddABEntry)
	prometheus.MustRegister(factomdadminBlockAdminBlockAddFirstABEntry)
	prometheus.MustRegister(factomdadminBlockAdminBlockMarshalBinary)
	prometheus.MustRegister(factomdadminBlockUnmarshalABlock)
	prometheus.MustRegister(factomdadminBlockAdminBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockAdminBlockUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockAdminBlockGetDBSignature)
	prometheus.MustRegister(factomdadminBlockAdminBlockJSONByte)
	prometheus.MustRegister(factomdadminBlockAdminBlockJSONString)
	prometheus.MustRegister(factomdadminBlockAdminBlockMarshalJSON)
	prometheus.MustRegister(factomdadminBlockNewAdminBlock)
	prometheus.MustRegister(factomdadminBlockCheckBlockPairIntegrity)
	prometheus.MustRegister(factomdadminBlockABlockHeaderInit)
	prometheus.MustRegister(factomdadminBlockABlockHeaderString)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetMessageCount)
	prometheus.MustRegister(factomdadminBlockABlockHeaderSetMessageCount)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetBodySize)
	prometheus.MustRegister(factomdadminBlockABlockHeaderSetBodySize)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetAdminChainID)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetDBHeight)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetHeaderExpansionArea)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetHeaderExpansionSize)
	prometheus.MustRegister(factomdadminBlockABlockHeaderGetPrevBackRefHash)
	prometheus.MustRegister(factomdadminBlockABlockHeaderSetDBHeight)
	prometheus.MustRegister(factomdadminBlockABlockHeaderSetHeaderExpansionArea)
	prometheus.MustRegister(factomdadminBlockABlockHeaderSetPrevBackRefHash)
	prometheus.MustRegister(factomdadminBlockABlockHeaderMarshalBinary)
	prometheus.MustRegister(factomdadminBlockABlockHeaderUnmarshalBinaryData)
	prometheus.MustRegister(factomdadminBlockABlockHeaderUnmarshalBinary)
	prometheus.MustRegister(factomdadminBlockABlockHeaderJSONByte)
	prometheus.MustRegister(factomdadminBlockABlockHeaderJSONString)
	prometheus.MustRegister(factomdadminBlockABlockHeaderMarshalJSON)
}
