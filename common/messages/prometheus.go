package messages

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdmessagesresend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_resend_Summary",
		Help: "Summary of calls to  factomd_messages_resend",
	})
	factomdmessagesMessageBaseSendOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SendOut_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SendOut",
	})
	factomdmessagesMessageBaseGetNoResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetNoResend_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetNoResend",
	})
	factomdmessagesMessageBaseSetNoResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetNoResend_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetNoResend",
	})
	factomdmessagesMessageBaseIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsValid_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_IsValid",
	})
	factomdmessagesMessageBaseSetValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetValid_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetValid",
	})
	factomdmessagesMessageBaseMarkSentInvalid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_MarkSentInvalid_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_MarkSentInvalid",
	})
	factomdmessagesMessageBaseSentInvlaid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SentInvlaid_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SentInvlaid",
	})
	factomdmessagesMessageBaseResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_Resend_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_Resend",
	})
	factomdmessagesMessageBaseExpire = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_Expire_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_Expire",
	})
	factomdmessagesMessageBaseIsStalled = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsStalled_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_IsStalled",
	})
	factomdmessagesMessageBaseSetStall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetStall_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetStall",
	})
	factomdmessagesMessageBaseGetFullMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetFullMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetFullMsgHash",
	})
	factomdmessagesMessageBaseSetFullMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetFullMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetFullMsgHash",
	})
	factomdmessagesMessageBaseGetOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetOrigin_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetOrigin",
	})
	factomdmessagesMessageBaseSetOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetOrigin_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetOrigin",
	})
	factomdmessagesMessageBaseGetNetworkOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetNetworkOrigin_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetNetworkOrigin",
	})
	factomdmessagesMessageBaseSetNetworkOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetNetworkOrigin_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetNetworkOrigin",
	})
	factomdmessagesMessageBaseIsPeer2Peer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsPeer2Peer_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_IsPeer2Peer",
	})
	factomdmessagesMessageBaseSetPeer2Peer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetPeer2Peer_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetPeer2Peer",
	})
	factomdmessagesMessageBaseIsLocal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsLocal_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_IsLocal",
	})
	factomdmessagesMessageBaseSetLocal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetLocal_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetLocal",
	})
	factomdmessagesMessageBaseGetLeaderChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetLeaderChainID_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetLeaderChainID",
	})
	factomdmessagesMessageBaseSetLeaderChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetLeaderChainID_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetLeaderChainID",
	})
	factomdmessagesMessageBaseGetVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetVMIndex",
	})
	factomdmessagesMessageBaseSetVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetVMIndex",
	})
	factomdmessagesMessageBaseGetVMHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetVMHash_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetVMHash",
	})
	factomdmessagesMessageBaseSetVMHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetVMHash_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetVMHash",
	})
	factomdmessagesMessageBaseGetMinute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetMinute_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_GetMinute",
	})
	factomdmessagesMessageBaseSetMinute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetMinute_Summary",
		Help: "Summary of calls to  factomd_messages_MessageBase_SetMinute",
	})
	factomdmessagesAckGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_GetRepeatHash",
	})
	factomdmessagesAckGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_GetHash",
	})
	factomdmessagesAckGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_GetMsgHash",
	})
	factomdmessagesAckType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_Type_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_Type",
	})
	factomdmessagesAckGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_GetTimestamp",
	})
	factomdmessagesAckVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_VerifySignature",
	})
	factomdmessagesAckValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_Validate",
	})
	factomdmessagesAckComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_ComputeVMIndex",
	})
	factomdmessagesAckLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_LeaderExecute",
	})
	factomdmessagesAckFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_FollowerExecute",
	})
	factomdmessagesAckProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_Process_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_Process",
	})
	factomdmessagesAckJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_JSONByte",
	})
	factomdmessagesAckJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_JSONString",
	})
	factomdmessagesAckSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_Sign",
	})
	factomdmessagesAckGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_GetSignature",
	})
	factomdmessagesAckUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_UnmarshalBinaryData",
	})
	factomdmessagesAckUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_UnmarshalBinary",
	})
	factomdmessagesAckMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_MarshalForSignature",
	})
	factomdmessagesAckMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_MarshalBinary",
	})
	factomdmessagesAckString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_String_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_String",
	})
	factomdmessagesAckIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Ack_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_Ack_IsSameAs",
	})
	factomdmessagesAddServerMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_GetRepeatHash",
	})
	factomdmessagesAddServerMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_GetHash",
	})
	factomdmessagesAddServerMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_GetMsgHash",
	})
	factomdmessagesAddServerMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_Type",
	})
	factomdmessagesAddServerMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_GetTimestamp",
	})
	factomdmessagesAddServerMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_Validate",
	})
	factomdmessagesAddServerMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_ComputeVMIndex",
	})
	factomdmessagesAddServerMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_LeaderExecute",
	})
	factomdmessagesAddServerMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_FollowerExecute",
	})
	factomdmessagesAddServerMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_Process",
	})
	factomdmessagesAddServerMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_JSONByte",
	})
	factomdmessagesAddServerMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_JSONString",
	})
	factomdmessagesAddServerMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_Sign",
	})
	factomdmessagesAddServerMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_GetSignature",
	})
	factomdmessagesAddServerMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_VerifySignature",
	})
	factomdmessagesAddServerMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_UnmarshalBinaryData",
	})
	factomdmessagesAddServerMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_UnmarshalBinary",
	})
	factomdmessagesAddServerMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_MarshalForSignature",
	})
	factomdmessagesAddServerMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_MarshalBinary",
	})
	factomdmessagesAddServerMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_String",
	})
	factomdmessagesAddServerMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServerMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_AddServerMsg_IsSameAs",
	})
	factomdmessagesNewAddServerMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewAddServerMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewAddServerMsg",
	})
	factomdmessagesNewAddServerByHashMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewAddServerByHashMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewAddServerByHashMsg",
	})
	factomdmessagesAuditServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_GetRepeatHash",
	})
	factomdmessagesAuditServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_IsSameAs",
	})
	factomdmessagesAuditServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_Sign",
	})
	factomdmessagesAuditServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_GetSignature",
	})
	factomdmessagesAuditServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_VerifySignature",
	})
	factomdmessagesAuditServerFaultProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Process_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_Process",
	})
	factomdmessagesAuditServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_GetTimestamp",
	})
	factomdmessagesAuditServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_GetHash",
	})
	factomdmessagesAuditServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_GetMsgHash",
	})
	factomdmessagesAuditServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Type_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_Type",
	})
	factomdmessagesAuditServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_UnmarshalBinaryData",
	})
	factomdmessagesAuditServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_UnmarshalBinary",
	})
	factomdmessagesAuditServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_MarshalForSignature",
	})
	factomdmessagesAuditServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_MarshalBinary",
	})
	factomdmessagesAuditServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_String_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_String",
	})
	factomdmessagesAuditServerFaultDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_DBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_DBHeight",
	})
	factomdmessagesAuditServerFaultChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_ChainID",
	})
	factomdmessagesAuditServerFaultListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_ListHeight",
	})
	factomdmessagesAuditServerFaultSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_SerialHash",
	})
	factomdmessagesAuditServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_Validate",
	})
	factomdmessagesAuditServerFaultComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_ComputeVMIndex",
	})
	factomdmessagesAuditServerFaultLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_LeaderExecute",
	})
	factomdmessagesAuditServerFaultFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_FollowerExecute",
	})
	factomdmessagesAuditServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_JSONByte",
	})
	factomdmessagesAuditServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_AuditServerFault_JSONString",
	})
	factomdmessagesBounceAddData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_AddData_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_AddData",
	})
	factomdmessagesBounceGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_GetRepeatHash",
	})
	factomdmessagesBounceGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_GetHash",
	})
	factomdmessagesBounceSizeOf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_SizeOf_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_SizeOf",
	})
	factomdmessagesBounceGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_GetMsgHash",
	})
	factomdmessagesBounceType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_Type_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_Type",
	})
	factomdmessagesBounceGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_GetTimestamp",
	})
	factomdmessagesBounceVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_VerifySignature",
	})
	factomdmessagesBounceValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_Validate",
	})
	factomdmessagesBounceComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_ComputeVMIndex",
	})
	factomdmessagesBounceLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_LeaderExecute",
	})
	factomdmessagesBounceFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_FollowerExecute",
	})
	factomdmessagesBounceProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_Process_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_Process",
	})
	factomdmessagesBounceJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_JSONByte",
	})
	factomdmessagesBounceJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_JSONString",
	})
	factomdmessagesBounceSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_Sign",
	})
	factomdmessagesBounceGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_GetSignature",
	})
	factomdmessagesBounceUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_UnmarshalBinaryData",
	})
	factomdmessagesBounceUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_UnmarshalBinary",
	})
	factomdmessagesBounceMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_MarshalForSignature",
	})
	factomdmessagesBounceMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_MarshalBinary",
	})
	factomdmessagesBounceString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_String_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_String",
	})
	factomdmessagesBounceIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_Bounce_IsSameAs",
	})
	factomdmessagesBounceReplyGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_GetRepeatHash",
	})
	factomdmessagesBounceReplyGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_GetHash",
	})
	factomdmessagesBounceReplySizeOf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_SizeOf_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_SizeOf",
	})
	factomdmessagesBounceReplyGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_GetMsgHash",
	})
	factomdmessagesBounceReplyType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_Type_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_Type",
	})
	factomdmessagesBounceReplyGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_GetTimestamp",
	})
	factomdmessagesBounceReplyVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_VerifySignature",
	})
	factomdmessagesBounceReplyValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_Validate",
	})
	factomdmessagesBounceReplyComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_ComputeVMIndex",
	})
	factomdmessagesBounceReplyLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_LeaderExecute",
	})
	factomdmessagesBounceReplyFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_FollowerExecute",
	})
	factomdmessagesBounceReplyProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_Process_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_Process",
	})
	factomdmessagesBounceReplyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_JSONByte",
	})
	factomdmessagesBounceReplyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_JSONString",
	})
	factomdmessagesBounceReplySign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_Sign",
	})
	factomdmessagesBounceReplyGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_GetSignature",
	})
	factomdmessagesBounceReplyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_UnmarshalBinaryData",
	})
	factomdmessagesBounceReplyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_UnmarshalBinary",
	})
	factomdmessagesBounceReplyMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_MarshalForSignature",
	})
	factomdmessagesBounceReplyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_MarshalBinary",
	})
	factomdmessagesBounceReplyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_String_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_String",
	})
	factomdmessagesBounceReplyIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_BounceReply_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_BounceReply_IsSameAs",
	})
	factomdmessagesChangeServerKeyMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_GetRepeatHash",
	})
	factomdmessagesChangeServerKeyMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_GetHash",
	})
	factomdmessagesChangeServerKeyMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_GetMsgHash",
	})
	factomdmessagesChangeServerKeyMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_Type",
	})
	factomdmessagesChangeServerKeyMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_GetTimestamp",
	})
	factomdmessagesChangeServerKeyMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_Validate",
	})
	factomdmessagesChangeServerKeyMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_ComputeVMIndex",
	})
	factomdmessagesChangeServerKeyMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_LeaderExecute",
	})
	factomdmessagesChangeServerKeyMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_FollowerExecute",
	})
	factomdmessagesChangeServerKeyMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_Process",
	})
	factomdmessagesChangeServerKeyMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_JSONByte",
	})
	factomdmessagesChangeServerKeyMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_JSONString",
	})
	factomdmessagesChangeServerKeyMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_Sign",
	})
	factomdmessagesChangeServerKeyMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_GetSignature",
	})
	factomdmessagesChangeServerKeyMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_VerifySignature",
	})
	factomdmessagesChangeServerKeyMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_UnmarshalBinaryData",
	})
	factomdmessagesChangeServerKeyMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_UnmarshalBinary",
	})
	factomdmessagesChangeServerKeyMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_MarshalForSignature",
	})
	factomdmessagesChangeServerKeyMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_MarshalBinary",
	})
	factomdmessagesChangeServerKeyMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_String",
	})
	factomdmessagesChangeServerKeyMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ChangeServerKeyMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_ChangeServerKeyMsg_IsSameAs",
	})
	factomdmessagesNewChangeServerKeyMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewChangeServerKeyMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewChangeServerKeyMsg",
	})
	factomdmessagesCommitChainMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_IsSameAs",
	})
	factomdmessagesCommitChainMsgGetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetCount",
	})
	factomdmessagesCommitChainMsgIncCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_IncCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_IncCount",
	})
	factomdmessagesCommitChainMsgSetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_SetCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_SetCount",
	})
	factomdmessagesCommitChainMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_Process",
	})
	factomdmessagesCommitChainMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetRepeatHash",
	})
	factomdmessagesCommitChainMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetHash",
	})
	factomdmessagesCommitChainMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetMsgHash",
	})
	factomdmessagesCommitChainMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetTimestamp",
	})
	factomdmessagesCommitChainMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_Type",
	})
	factomdmessagesCommitChainMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_Validate",
	})
	factomdmessagesCommitChainMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_ComputeVMIndex",
	})
	factomdmessagesCommitChainMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_LeaderExecute",
	})
	factomdmessagesCommitChainMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_FollowerExecute",
	})
	factomdmessagesCommitChainMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_JSONByte",
	})
	factomdmessagesCommitChainMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_JSONString",
	})
	factomdmessagesCommitChainMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_Sign",
	})
	factomdmessagesCommitChainMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_GetSignature",
	})
	factomdmessagesCommitChainMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_VerifySignature",
	})
	factomdmessagesCommitChainMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_UnmarshalBinaryData",
	})
	factomdmessagesCommitChainMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_UnmarshalBinary",
	})
	factomdmessagesCommitChainMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_MarshalForSignature",
	})
	factomdmessagesCommitChainMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_MarshalBinary",
	})
	factomdmessagesCommitChainMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitChainMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_CommitChainMsg_String",
	})
	factomdmessagesCommitEntryMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_IsSameAs",
	})
	factomdmessagesCommitEntryMsgGetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetCount",
	})
	factomdmessagesCommitEntryMsgIncCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_IncCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_IncCount",
	})
	factomdmessagesCommitEntryMsgSetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_SetCount_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_SetCount",
	})
	factomdmessagesCommitEntryMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_Process",
	})
	factomdmessagesCommitEntryMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetRepeatHash",
	})
	factomdmessagesCommitEntryMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetHash",
	})
	factomdmessagesCommitEntryMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetMsgHash",
	})
	factomdmessagesCommitEntryMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetTimestamp",
	})
	factomdmessagesCommitEntryMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_Type",
	})
	factomdmessagesCommitEntryMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_Sign",
	})
	factomdmessagesCommitEntryMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_GetSignature",
	})
	factomdmessagesCommitEntryMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_VerifySignature",
	})
	factomdmessagesCommitEntryMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_UnmarshalBinaryData",
	})
	factomdmessagesCommitEntryMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_UnmarshalBinary",
	})
	factomdmessagesCommitEntryMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_MarshalForSignature",
	})
	factomdmessagesCommitEntryMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_MarshalBinary",
	})
	factomdmessagesCommitEntryMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_String",
	})
	factomdmessagesCommitEntryMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_Validate",
	})
	factomdmessagesCommitEntryMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_ComputeVMIndex",
	})
	factomdmessagesCommitEntryMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_LeaderExecute",
	})
	factomdmessagesCommitEntryMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_FollowerExecute",
	})
	factomdmessagesCommitEntryMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_JSONByte",
	})
	factomdmessagesCommitEntryMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_CommitEntryMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_CommitEntryMsg_JSONString",
	})
	factomdmessagesNewCommitEntryMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewCommitEntryMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewCommitEntryMsg",
	})
	factomdmessagesDataResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_IsSameAs",
	})
	factomdmessagesDataResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_GetRepeatHash",
	})
	factomdmessagesDataResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_GetHash",
	})
	factomdmessagesDataResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_GetMsgHash",
	})
	factomdmessagesDataResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_Type_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_Type",
	})
	factomdmessagesDataResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_GetTimestamp",
	})
	factomdmessagesDataResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_Validate",
	})
	factomdmessagesDataResponseComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_ComputeVMIndex",
	})
	factomdmessagesDataResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_LeaderExecute",
	})
	factomdmessagesDataResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_FollowerExecute",
	})
	factomdmessagesDataResponseProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_Process_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_Process",
	})
	factomdmessagesDataResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_JSONByte",
	})
	factomdmessagesDataResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_JSONString",
	})
	factomdmessagesDataResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_UnmarshalBinaryData",
	})
	factomdmessagesDataResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_UnmarshalBinary",
	})
	factomdmessagesattemptEntryUnmarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_attemptEntryUnmarshal_Summary",
		Help: "Summary of calls to  factomd_messages_attemptEntryUnmarshal",
	})
	factomdmessagesattemptEBlockUnmarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_attemptEBlockUnmarshal_Summary",
		Help: "Summary of calls to  factomd_messages_attemptEBlockUnmarshal",
	})
	factomdmessagesDataResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_MarshalBinary",
	})
	factomdmessagesDataResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DataResponse_String_Summary",
		Help: "Summary of calls to  factomd_messages_DataResponse_String",
	})
	factomdmessagesNewDataResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewDataResponse_Summary",
		Help: "Summary of calls to  factomd_messages_NewDataResponse",
	})
	factomdmessagesDBStateMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_IsSameAs",
	})
	factomdmessagesDBStateMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_GetRepeatHash",
	})
	factomdmessagesDBStateMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_GetHash",
	})
	factomdmessagesDBStateMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_GetMsgHash",
	})
	factomdmessagesDBStateMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_Type",
	})
	factomdmessagesDBStateMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_GetTimestamp",
	})
	factomdmessagesDBStateMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_Validate",
	})
	factomdmessagesDBStateMsgSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_SigTally_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_SigTally",
	})
	factomdmessagesDBStateMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_ComputeVMIndex",
	})
	factomdmessagesDBStateMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_LeaderExecute",
	})
	factomdmessagesDBStateMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_FollowerExecute",
	})
	factomdmessagesDBStateMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_Process",
	})
	factomdmessagesDBStateMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_JSONByte",
	})
	factomdmessagesDBStateMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_JSONString",
	})
	factomdmessagesDBStateMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_UnmarshalBinaryData",
	})
	factomdmessagesDBStateMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_UnmarshalBinary",
	})
	factomdmessagesDBStateMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_MarshalBinary",
	})
	factomdmessagesDBStateMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMsg_String",
	})
	factomdmessagesNewDBStateMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewDBStateMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewDBStateMsg",
	})
	factomdmessagesDBStateMissingIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_IsSameAs",
	})
	factomdmessagesDBStateMissingGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_GetRepeatHash",
	})
	factomdmessagesDBStateMissingGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_GetHash",
	})
	factomdmessagesDBStateMissingGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_GetMsgHash",
	})
	factomdmessagesDBStateMissingType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Type_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_Type",
	})
	factomdmessagesDBStateMissingGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_GetTimestamp",
	})
	factomdmessagesDBStateMissingValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_Validate",
	})
	factomdmessagesDBStateMissingComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_ComputeVMIndex",
	})
	factomdmessagesDBStateMissingLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_LeaderExecute",
	})
	factomdmessagesDBStateMissingsend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_send_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_send",
	})
	factomdmessagesDBStateMissingFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_FollowerExecute",
	})
	factomdmessagesDBStateMissingProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Process_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_Process",
	})
	factomdmessagesDBStateMissingJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_JSONByte",
	})
	factomdmessagesDBStateMissingJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_JSONString",
	})
	factomdmessagesDBStateMissingUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_UnmarshalBinaryData",
	})
	factomdmessagesDBStateMissingUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_UnmarshalBinary",
	})
	factomdmessagesDBStateMissingMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_MarshalForSignature",
	})
	factomdmessagesDBStateMissingMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_MarshalBinary",
	})
	factomdmessagesDBStateMissingString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_String_Summary",
		Help: "Summary of calls to  factomd_messages_DBStateMissing_String",
	})
	factomdmessagesNewDBStateMissing = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewDBStateMissing_Summary",
		Help: "Summary of calls to  factomd_messages_NewDBStateMissing",
	})
	factomdmessagesDirectoryBlockSignatureIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_IsSameAs",
	})
	factomdmessagesDirectoryBlockSignatureProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Process_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_Process",
	})
	factomdmessagesDirectoryBlockSignatureGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_GetRepeatHash",
	})
	factomdmessagesDirectoryBlockSignatureGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_GetHash",
	})
	factomdmessagesDirectoryBlockSignatureGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_GetMsgHash",
	})
	factomdmessagesDirectoryBlockSignatureGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_GetTimestamp",
	})
	factomdmessagesDirectoryBlockSignatureType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Type_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_Type",
	})
	factomdmessagesDirectoryBlockSignatureValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_Validate",
	})
	factomdmessagesDirectoryBlockSignatureComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_ComputeVMIndex",
	})
	factomdmessagesDirectoryBlockSignatureLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_LeaderExecute",
	})
	factomdmessagesDirectoryBlockSignatureFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_FollowerExecute",
	})
	factomdmessagesDirectoryBlockSignatureSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_Sign",
	})
	factomdmessagesDirectoryBlockSignatureSignHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_SignHeader_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_SignHeader",
	})
	factomdmessagesDirectoryBlockSignatureGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_GetSignature",
	})
	factomdmessagesDirectoryBlockSignatureVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_VerifySignature",
	})
	factomdmessagesDirectoryBlockSignatureUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_UnmarshalBinaryData",
	})
	factomdmessagesDirectoryBlockSignatureUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_UnmarshalBinary",
	})
	factomdmessagesDirectoryBlockSignatureMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_MarshalForSignature",
	})
	factomdmessagesDirectoryBlockSignatureMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_MarshalBinary",
	})
	factomdmessagesDirectoryBlockSignatureString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_String_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_String",
	})
	factomdmessagesDirectoryBlockSignatureJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_JSONByte",
	})
	factomdmessagesDirectoryBlockSignatureJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_DirectoryBlockSignature_JSONString",
	})
	factomdmessagesEntryBlockResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_IsSameAs",
	})
	factomdmessagesEntryBlockResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_GetRepeatHash",
	})
	factomdmessagesEntryBlockResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_GetHash",
	})
	factomdmessagesEntryBlockResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_GetMsgHash",
	})
	factomdmessagesEntryBlockResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Type_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_Type",
	})
	factomdmessagesEntryBlockResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_GetTimestamp",
	})
	factomdmessagesEntryBlockResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_Validate",
	})
	factomdmessagesEntryBlockResponseComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_ComputeVMIndex",
	})
	factomdmessagesEntryBlockResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_LeaderExecute",
	})
	factomdmessagesEntryBlockResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_FollowerExecute",
	})
	factomdmessagesEntryBlockResponseProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Process_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_Process",
	})
	factomdmessagesEntryBlockResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_JSONByte",
	})
	factomdmessagesEntryBlockResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_JSONString",
	})
	factomdmessagesEntryBlockResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_UnmarshalBinaryData",
	})
	factomdmessagesEntryBlockResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_UnmarshalBinary",
	})
	factomdmessagesEntryBlockResponseMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_MarshalForSignature",
	})
	factomdmessagesEntryBlockResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_MarshalBinary",
	})
	factomdmessagesEntryBlockResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_String_Summary",
		Help: "Summary of calls to  factomd_messages_EntryBlockResponse_String",
	})
	factomdmessagesNewEntryBlockResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewEntryBlockResponse_Summary",
		Help: "Summary of calls to  factomd_messages_NewEntryBlockResponse",
	})
	factomdmessagesEOMIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_IsSameAs",
	})
	factomdmessagesEOMProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Process_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_Process",
	})
	factomdmessagesEOMGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_GetRepeatHash",
	})
	factomdmessagesEOMGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_GetHash",
	})
	factomdmessagesEOMGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_GetMsgHash",
	})
	factomdmessagesEOMGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_GetTimestamp",
	})
	factomdmessagesEOMType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Type_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_Type",
	})
	factomdmessagesEOMValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_Validate",
	})
	factomdmessagesEOMComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_ComputeVMIndex",
	})
	factomdmessagesEOMLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_LeaderExecute",
	})
	factomdmessagesEOMFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_FollowerExecute",
	})
	factomdmessagesEOMJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_JSONByte",
	})
	factomdmessagesEOMJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_JSONString",
	})
	factomdmessagesEOMSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_Sign",
	})
	factomdmessagesEOMGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_GetSignature",
	})
	factomdmessagesEOMVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_VerifySignature",
	})
	factomdmessagesEOMUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_UnmarshalBinaryData",
	})
	factomdmessagesEOMUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_UnmarshalBinary",
	})
	factomdmessagesEOMMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_MarshalForSignature",
	})
	factomdmessagesEOMMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_MarshalBinary",
	})
	factomdmessagesEOMString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_String_Summary",
		Help: "Summary of calls to  factomd_messages_EOM_String",
	})
	factomdmessagesEOMTimeoutIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_IsSameAs",
	})
	factomdmessagesEOMTimeoutSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_Sign",
	})
	factomdmessagesEOMTimeoutGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_GetSignature",
	})
	factomdmessagesEOMTimeoutVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_VerifySignature",
	})
	factomdmessagesEOMTimeoutProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Process_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_Process",
	})
	factomdmessagesEOMTimeoutGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_GetRepeatHash",
	})
	factomdmessagesEOMTimeoutGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_GetHash",
	})
	factomdmessagesEOMTimeoutGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_GetMsgHash",
	})
	factomdmessagesEOMTimeoutGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_GetTimestamp",
	})
	factomdmessagesEOMTimeoutType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Type_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_Type",
	})
	factomdmessagesEOMTimeoutUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_UnmarshalBinaryData",
	})
	factomdmessagesEOMTimeoutUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_UnmarshalBinary",
	})
	factomdmessagesEOMTimeoutMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_MarshalForSignature",
	})
	factomdmessagesEOMTimeoutMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_MarshalBinary",
	})
	factomdmessagesEOMTimeoutString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_String_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_String",
	})
	factomdmessagesEOMTimeoutDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_DBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_DBHeight",
	})
	factomdmessagesEOMTimeoutChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_ChainID",
	})
	factomdmessagesEOMTimeoutListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_ListHeight",
	})
	factomdmessagesEOMTimeoutSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_SerialHash",
	})
	factomdmessagesEOMTimeoutValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_Validate",
	})
	factomdmessagesEOMTimeoutComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_ComputeVMIndex",
	})
	factomdmessagesEOMTimeoutLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_LeaderExecute",
	})
	factomdmessagesEOMTimeoutFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_FollowerExecute",
	})
	factomdmessagesEOMTimeoutJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_JSONByte",
	})
	factomdmessagesEOMTimeoutJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_EOMTimeout_JSONString",
	})
	factomdmessagesMessageErrorError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_Error_Summary",
		Help: "Summary of calls to  factomd_messages_MessageError_Error",
	})
	factomdmessagesmessageError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_messageError_Summary",
		Help: "Summary of calls to  factomd_messages_messageError",
	})
	factomdmessagesFactoidTransactionIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_IsSameAs",
	})
	factomdmessagesFactoidTransactionGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_GetRepeatHash",
	})
	factomdmessagesFactoidTransactionGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_GetHash",
	})
	factomdmessagesFactoidTransactionGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_GetMsgHash",
	})
	factomdmessagesFactoidTransactionGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_GetTimestamp",
	})
	factomdmessagesFactoidTransactionGetTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_GetTransaction_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_GetTransaction",
	})
	factomdmessagesFactoidTransactionSetTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_SetTransaction_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_SetTransaction",
	})
	factomdmessagesFactoidTransactionType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_Type_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_Type",
	})
	factomdmessagesFactoidTransactionValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_Validate",
	})
	factomdmessagesFactoidTransactionComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_ComputeVMIndex",
	})
	factomdmessagesFactoidTransactionLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_LeaderExecute",
	})
	factomdmessagesFactoidTransactionFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_FollowerExecute",
	})
	factomdmessagesFactoidTransactionProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_Process_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_Process",
	})
	factomdmessagesFactoidTransactionUnmarshalTransData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_UnmarshalTransData_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_UnmarshalTransData",
	})
	factomdmessagesFactoidTransactionUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_UnmarshalBinaryData",
	})
	factomdmessagesFactoidTransactionUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_UnmarshalBinary",
	})
	factomdmessagesFactoidTransactionMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_MarshalBinary",
	})
	factomdmessagesFactoidTransactionString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_String_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_String",
	})
	factomdmessagesFactoidTransactionJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_JSONByte",
	})
	factomdmessagesFactoidTransactionJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FactoidTransaction_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_FactoidTransaction_JSONString",
	})
	factomdmessagesFullServerFaultGetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetAmINegotiator_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetAmINegotiator",
	})
	factomdmessagesFullServerFaultSetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetAmINegotiator_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SetAmINegotiator",
	})
	factomdmessagesFullServerFaultGetMyVoteTallied = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetMyVoteTallied_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetMyVoteTallied",
	})
	factomdmessagesFullServerFaultSetMyVoteTallied = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetMyVoteTallied_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SetMyVoteTallied",
	})
	factomdmessagesFullServerFaultGetPledgeDone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetPledgeDone_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetPledgeDone",
	})
	factomdmessagesFullServerFaultSetPledgeDone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetPledgeDone_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SetPledgeDone",
	})
	factomdmessagesFullServerFaultGetLastMatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetLastMatch_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetLastMatch",
	})
	factomdmessagesFullServerFaultSetLastMatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetLastMatch_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SetLastMatch",
	})
	factomdmessagesFullServerFaultIsNil = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_IsNil_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_IsNil",
	})
	factomdmessagesFullServerFaultAddFaultVote = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_AddFaultVote_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_AddFaultVote",
	})
	factomdmessagesFullServerFaultPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Priority_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_Priority",
	})
	factomdmessagesFullServerFaultGetSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetSerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetSerialHash",
	})
	factomdmessagesFullServerFaultProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Process_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_Process",
	})
	factomdmessagesFullServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetRepeatHash",
	})
	factomdmessagesFullServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetHash",
	})
	factomdmessagesFullServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetMsgHash",
	})
	factomdmessagesFullServerFaultGetCoreHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetCoreHash_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetCoreHash",
	})
	factomdmessagesFullServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetTimestamp",
	})
	factomdmessagesFullServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Type_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_Type",
	})
	factomdmessagesFullServerFaultMarshalCore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalCore_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_MarshalCore",
	})
	factomdmessagesFullServerFaultMarshalForSF = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalForSF_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_MarshalForSF",
	})
	factomdmessagesFullServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_MarshalForSignature",
	})
	factomdmessagesSigListMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SigList_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_SigList_MarshalBinary",
	})
	factomdmessagesSigListUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SigList_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_SigList_UnmarshalBinaryData",
	})
	factomdmessagesFullServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_MarshalBinary",
	})
	factomdmessagesUnmarshall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Unmarshall_Summary",
		Help: "Summary of calls to  factomd_messages_Unmarshall",
	})
	factomdmessagesFullServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_UnmarshalBinaryData",
	})
	factomdmessagesFullServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_UnmarshalBinary",
	})
	factomdmessagesFullServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetSignature",
	})
	factomdmessagesFullServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_VerifySignature",
	})
	factomdmessagesFullServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_Sign",
	})
	factomdmessagesFullServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_String_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_String",
	})
	factomdmessagesFullServerFaultStringWithSigCnt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_StringWithSigCnt_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_StringWithSigCnt",
	})
	factomdmessagesFullServerFaultGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetDBHeight",
	})
	factomdmessagesFullServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_Validate",
	})
	factomdmessagesFullServerFaultSetAlreadyProcessed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetAlreadyProcessed_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SetAlreadyProcessed",
	})
	factomdmessagesFullServerFaultGetAlreadyProcessed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetAlreadyProcessed_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_GetAlreadyProcessed",
	})
	factomdmessagesFullServerFaultHasEnoughSigs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_HasEnoughSigs_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_HasEnoughSigs",
	})
	factomdmessagesFullServerFaultSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SigTally_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_SigTally",
	})
	factomdmessagesFullServerFaultComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_ComputeVMIndex",
	})
	factomdmessagesFullServerFaultLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_LeaderExecute",
	})
	factomdmessagesFullServerFaultFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_FollowerExecute",
	})
	factomdmessagesFullServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_JSONByte",
	})
	factomdmessagesFullServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_JSONString",
	})
	factomdmessagesFullServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_IsSameAs",
	})
	factomdmessagesFullServerFaultToAdminBlockEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_ToAdminBlockEntry_Summary",
		Help: "Summary of calls to  factomd_messages_FullServerFault_ToAdminBlockEntry",
	})
	factomdmessagesNewFullServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewFullServerFault_Summary",
		Help: "Summary of calls to  factomd_messages_NewFullServerFault",
	})
	factomdmessagesUnmarshalMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_UnmarshalMessage_Summary",
		Help: "Summary of calls to  factomd_messages_UnmarshalMessage",
	})
	factomdmessagesUnmarshalMessageData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_UnmarshalMessageData_Summary",
		Help: "Summary of calls to  factomd_messages_UnmarshalMessageData",
	})
	factomdmessagesMessageName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageName_Summary",
		Help: "Summary of calls to  factomd_messages_MessageName",
	})
	factomdmessagesSignSignable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignSignable_Summary",
		Help: "Summary of calls to  factomd_messages_SignSignable",
	})
	factomdmessagesVerifyMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_VerifyMessage_Summary",
		Help: "Summary of calls to  factomd_messages_VerifyMessage",
	})
	factomdmessagesHeartbeatIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_IsSameAs",
	})
	factomdmessagesHeartbeatProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Process_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_Process",
	})
	factomdmessagesHeartbeatGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_GetRepeatHash",
	})
	factomdmessagesHeartbeatGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_GetHash",
	})
	factomdmessagesHeartbeatGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_GetMsgHash",
	})
	factomdmessagesHeartbeatGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_GetTimestamp",
	})
	factomdmessagesHeartbeatType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Type_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_Type",
	})
	factomdmessagesHeartbeatUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_UnmarshalBinaryData",
	})
	factomdmessagesHeartbeatUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_UnmarshalBinary",
	})
	factomdmessagesHeartbeatMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_MarshalForSignature",
	})
	factomdmessagesHeartbeatMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_MarshalBinary",
	})
	factomdmessagesHeartbeatString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_String_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_String",
	})
	factomdmessagesHeartbeatChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_ChainID",
	})
	factomdmessagesHeartbeatListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_ListHeight",
	})
	factomdmessagesHeartbeatSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_SerialHash",
	})
	factomdmessagesHeartbeatValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_Validate",
	})
	factomdmessagesHeartbeatComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_ComputeVMIndex",
	})
	factomdmessagesHeartbeatLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_LeaderExecute",
	})
	factomdmessagesHeartbeatFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_FollowerExecute",
	})
	factomdmessagesHeartbeatJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_JSONByte",
	})
	factomdmessagesHeartbeatJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_JSONString",
	})
	factomdmessagesHeartbeatSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_Sign",
	})
	factomdmessagesHeartbeatGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_GetSignature",
	})
	factomdmessagesHeartbeatVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_Heartbeat_VerifySignature",
	})
	factomdmessagesInvalidDirectoryBlockIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_IsSameAs",
	})
	factomdmessagesInvalidDirectoryBlockSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_Sign",
	})
	factomdmessagesInvalidDirectoryBlockGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_GetSignature",
	})
	factomdmessagesInvalidDirectoryBlockVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_VerifySignature",
	})
	factomdmessagesInvalidDirectoryBlockProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Process_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_Process",
	})
	factomdmessagesInvalidDirectoryBlockGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_GetRepeatHash",
	})
	factomdmessagesInvalidDirectoryBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_GetHash",
	})
	factomdmessagesInvalidDirectoryBlockGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_GetMsgHash",
	})
	factomdmessagesInvalidDirectoryBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_GetTimestamp",
	})
	factomdmessagesInvalidDirectoryBlockType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Type_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_Type",
	})
	factomdmessagesInvalidDirectoryBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_UnmarshalBinaryData",
	})
	factomdmessagesInvalidDirectoryBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_UnmarshalBinary",
	})
	factomdmessagesInvalidDirectoryBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_MarshalBinary",
	})
	factomdmessagesInvalidDirectoryBlockMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_MarshalForSignature",
	})
	factomdmessagesInvalidDirectoryBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_String_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_String",
	})
	factomdmessagesInvalidDirectoryBlockDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_DBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_DBHeight",
	})
	factomdmessagesInvalidDirectoryBlockChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_ChainID",
	})
	factomdmessagesInvalidDirectoryBlockListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_ListHeight",
	})
	factomdmessagesInvalidDirectoryBlockSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_SerialHash",
	})
	factomdmessagesInvalidDirectoryBlockValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_Validate",
	})
	factomdmessagesInvalidDirectoryBlockComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_ComputeVMIndex",
	})
	factomdmessagesInvalidDirectoryBlockLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_LeaderExecute",
	})
	factomdmessagesInvalidDirectoryBlockFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_FollowerExecute",
	})
	factomdmessagesInvalidDirectoryBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_JSONByte",
	})
	factomdmessagesInvalidDirectoryBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_InvalidDirectoryBlock_JSONString",
	})
	factomdmessagesMissingDataIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_IsSameAs",
	})
	factomdmessagesMissingDataProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Process_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_Process",
	})
	factomdmessagesMissingDataGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_GetRepeatHash",
	})
	factomdmessagesMissingDataGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_GetHash",
	})
	factomdmessagesMissingDataGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_GetMsgHash",
	})
	factomdmessagesMissingDataGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_GetTimestamp",
	})
	factomdmessagesMissingDataType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Type_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_Type",
	})
	factomdmessagesMissingDataUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_UnmarshalBinaryData",
	})
	factomdmessagesMissingDataUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_UnmarshalBinary",
	})
	factomdmessagesMissingDataMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_MarshalBinary",
	})
	factomdmessagesMissingDataString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_String_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_String",
	})
	factomdmessagesMissingDataValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_Validate",
	})
	factomdmessagesMissingDataComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_ComputeVMIndex",
	})
	factomdmessagesMissingDataLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_LeaderExecute",
	})
	factomdmessagesMissingDataFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_FollowerExecute",
	})
	factomdmessagesMissingDataJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_JSONByte",
	})
	factomdmessagesMissingDataJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_MissingData_JSONString",
	})
	factomdmessagesNewMissingData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewMissingData_Summary",
		Help: "Summary of calls to  factomd_messages_NewMissingData",
	})
	factomdmessagesMissingEntryBlocksIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_IsSameAs",
	})
	factomdmessagesMissingEntryBlocksGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_GetRepeatHash",
	})
	factomdmessagesMissingEntryBlocksGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_GetHash",
	})
	factomdmessagesMissingEntryBlocksGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_GetMsgHash",
	})
	factomdmessagesMissingEntryBlocksType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_Type_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_Type",
	})
	factomdmessagesMissingEntryBlocksGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_GetTimestamp",
	})
	factomdmessagesMissingEntryBlocksValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_Validate",
	})
	factomdmessagesMissingEntryBlocksComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_ComputeVMIndex",
	})
	factomdmessagesMissingEntryBlocksLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_LeaderExecute",
	})
	factomdmessagesMissingEntryBlocksFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_FollowerExecute",
	})
	factomdmessagesMissingEntryBlocksProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_Process_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_Process",
	})
	factomdmessagesMissingEntryBlocksJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_JSONByte",
	})
	factomdmessagesMissingEntryBlocksJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_JSONString",
	})
	factomdmessagesMissingEntryBlocksUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_UnmarshalBinaryData",
	})
	factomdmessagesMissingEntryBlocksUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_UnmarshalBinary",
	})
	factomdmessagesMissingEntryBlocksMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_MarshalForSignature",
	})
	factomdmessagesMissingEntryBlocksMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_MarshalBinary",
	})
	factomdmessagesMissingEntryBlocksString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_String_Summary",
		Help: "Summary of calls to  factomd_messages_MissingEntryBlocks_String",
	})
	factomdmessagesNewMissingEntryBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewMissingEntryBlocks_Summary",
		Help: "Summary of calls to  factomd_messages_NewMissingEntryBlocks",
	})
	factomdmessagesMissingMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_IsSameAs",
	})
	factomdmessagesMissingMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_Process",
	})
	factomdmessagesMissingMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_GetRepeatHash",
	})
	factomdmessagesMissingMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_GetHash",
	})
	factomdmessagesMissingMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_GetMsgHash",
	})
	factomdmessagesMissingMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_GetTimestamp",
	})
	factomdmessagesMissingMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_Type",
	})
	factomdmessagesMissingMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_UnmarshalBinaryData",
	})
	factomdmessagesMissingMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_UnmarshalBinary",
	})
	factomdmessagesMissingMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_MarshalBinary",
	})
	factomdmessagesMissingMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_String",
	})
	factomdmessagesMissingMsgChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_ChainID",
	})
	factomdmessagesMissingMsgListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_ListHeight",
	})
	factomdmessagesMissingMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_Validate",
	})
	factomdmessagesMissingMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_ComputeVMIndex",
	})
	factomdmessagesMissingMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_LeaderExecute",
	})
	factomdmessagesMissingMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_FollowerExecute",
	})
	factomdmessagesMissingMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_JSONByte",
	})
	factomdmessagesMissingMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_JSONString",
	})
	factomdmessagesMissingMsgAddHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_AddHeight_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsg_AddHeight",
	})
	factomdmessagesNewMissingMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewMissingMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewMissingMsg",
	})
	factomdmessagesMissingMsgResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_IsSameAs",
	})
	factomdmessagesMissingMsgResponseProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Process_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_Process",
	})
	factomdmessagesMissingMsgResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_GetRepeatHash",
	})
	factomdmessagesMissingMsgResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_GetHash",
	})
	factomdmessagesMissingMsgResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_GetMsgHash",
	})
	factomdmessagesMissingMsgResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_GetTimestamp",
	})
	factomdmessagesMissingMsgResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Type_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_Type",
	})
	factomdmessagesMissingMsgResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_UnmarshalBinaryData",
	})
	factomdmessagesMissingMsgResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_UnmarshalBinary",
	})
	factomdmessagesMissingMsgResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_MarshalBinary",
	})
	factomdmessagesMissingMsgResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_String_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_String",
	})
	factomdmessagesMissingMsgResponseChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_ChainID",
	})
	factomdmessagesMissingMsgResponseListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_ListHeight",
	})
	factomdmessagesMissingMsgResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_Validate",
	})
	factomdmessagesMissingMsgResponseComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_ComputeVMIndex",
	})
	factomdmessagesMissingMsgResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_LeaderExecute",
	})
	factomdmessagesMissingMsgResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_FollowerExecute",
	})
	factomdmessagesMissingMsgResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_JSONByte",
	})
	factomdmessagesMissingMsgResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_MissingMsgResponse_JSONString",
	})
	factomdmessagesNewMissingMsgResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewMissingMsgResponse_Summary",
		Help: "Summary of calls to  factomd_messages_NewMissingMsgResponse",
	})
	factomdmessagesRemoveServerMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_GetRepeatHash",
	})
	factomdmessagesRemoveServerMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_GetHash",
	})
	factomdmessagesRemoveServerMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_GetMsgHash",
	})
	factomdmessagesRemoveServerMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_Type",
	})
	factomdmessagesRemoveServerMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_GetTimestamp",
	})
	factomdmessagesRemoveServerMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_Validate",
	})
	factomdmessagesRemoveServerMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_ComputeVMIndex",
	})
	factomdmessagesRemoveServerMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_LeaderExecute",
	})
	factomdmessagesRemoveServerMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_FollowerExecute",
	})
	factomdmessagesRemoveServerMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_Process",
	})
	factomdmessagesRemoveServerMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_JSONByte",
	})
	factomdmessagesRemoveServerMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_JSONString",
	})
	factomdmessagesRemoveServerMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_Sign",
	})
	factomdmessagesRemoveServerMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_GetSignature",
	})
	factomdmessagesRemoveServerMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_VerifySignature",
	})
	factomdmessagesRemoveServerMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_UnmarshalBinaryData",
	})
	factomdmessagesRemoveServerMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_UnmarshalBinary",
	})
	factomdmessagesRemoveServerMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_MarshalForSignature",
	})
	factomdmessagesRemoveServerMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_MarshalBinary",
	})
	factomdmessagesRemoveServerMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_String",
	})
	factomdmessagesRemoveServerMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RemoveServerMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_RemoveServerMsg_IsSameAs",
	})
	factomdmessagesNewRemoveServerMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewRemoveServerMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewRemoveServerMsg",
	})
	factomdmessagesRequestBlockIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_IsSameAs",
	})
	factomdmessagesRequestBlockProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_Process_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_Process",
	})
	factomdmessagesRequestBlockGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_GetRepeatHash",
	})
	factomdmessagesRequestBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_GetHash",
	})
	factomdmessagesRequestBlockGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_GetMsgHash",
	})
	factomdmessagesRequestBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_GetTimestamp",
	})
	factomdmessagesRequestBlockType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_Type_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_Type",
	})
	factomdmessagesRequestBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_UnmarshalBinaryData",
	})
	factomdmessagesRequestBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_UnmarshalBinary",
	})
	factomdmessagesRequestBlockMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_MarshalForSignature",
	})
	factomdmessagesRequestBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_MarshalBinary",
	})
	factomdmessagesRequestBlockString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_String_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_String",
	})
	factomdmessagesRequestBlockDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_DBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_DBHeight",
	})
	factomdmessagesRequestBlockChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_ChainID",
	})
	factomdmessagesRequestBlockListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_ListHeight",
	})
	factomdmessagesRequestBlockSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_SerialHash",
	})
	factomdmessagesRequestBlockSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_Signature_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_Signature",
	})
	factomdmessagesRequestBlockValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_Validate",
	})
	factomdmessagesRequestBlockComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_ComputeVMIndex",
	})
	factomdmessagesRequestBlockLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_LeaderExecute",
	})
	factomdmessagesRequestBlockFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_FollowerExecute",
	})
	factomdmessagesRequestBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_JSONByte",
	})
	factomdmessagesRequestBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_RequestBlock_JSONString",
	})
	factomdmessagesRevealEntryMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_IsSameAs",
	})
	factomdmessagesRevealEntryMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Process_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_Process",
	})
	factomdmessagesRevealEntryMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_GetRepeatHash",
	})
	factomdmessagesRevealEntryMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_GetHash",
	})
	factomdmessagesRevealEntryMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_GetMsgHash",
	})
	factomdmessagesRevealEntryMsgGetChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetChainIDHash_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_GetChainIDHash",
	})
	factomdmessagesRevealEntryMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_GetTimestamp",
	})
	factomdmessagesRevealEntryMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Type_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_Type",
	})
	factomdmessagesRevealEntryMsgValidateRTN = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_ValidateRTN_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_ValidateRTN",
	})
	factomdmessagesRevealEntryMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_Validate",
	})
	factomdmessagesRevealEntryMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_ComputeVMIndex",
	})
	factomdmessagesRevealEntryMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_LeaderExecute",
	})
	factomdmessagesRevealEntryMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_FollowerExecute",
	})
	factomdmessagesRevealEntryMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_JSONByte",
	})
	factomdmessagesRevealEntryMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_JSONString",
	})
	factomdmessagesNewRevealEntryMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewRevealEntryMsg_Summary",
		Help: "Summary of calls to  factomd_messages_NewRevealEntryMsg",
	})
	factomdmessagesRevealEntryMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_UnmarshalBinaryData",
	})
	factomdmessagesRevealEntryMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_UnmarshalBinary",
	})
	factomdmessagesRevealEntryMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_MarshalBinary",
	})
	factomdmessagesRevealEntryMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_String_Summary",
		Help: "Summary of calls to  factomd_messages_RevealEntryMsg_String",
	})
	factomdmessagesServerFaultProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Process_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_Process",
	})
	factomdmessagesServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetRepeatHash",
	})
	factomdmessagesServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetHash",
	})
	factomdmessagesServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetMsgHash",
	})
	factomdmessagesServerFaultGetCoreHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetCoreHash_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetCoreHash",
	})
	factomdmessagesServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetTimestamp",
	})
	factomdmessagesServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Type_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_Type",
	})
	factomdmessagesServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_MarshalForSignature",
	})
	factomdmessagesServerFaultPreMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_PreMarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_PreMarshalBinary",
	})
	factomdmessagesServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_MarshalBinary",
	})
	factomdmessagesServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_UnmarshalBinaryData",
	})
	factomdmessagesServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_UnmarshalBinary",
	})
	factomdmessagesServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetSignature",
	})
	factomdmessagesServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_VerifySignature",
	})
	factomdmessagesServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_Sign",
	})
	factomdmessagesServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_String_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_String",
	})
	factomdmessagesServerFaultGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetDBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_GetDBHeight",
	})
	factomdmessagesServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_Validate",
	})
	factomdmessagesServerFaultComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_ComputeVMIndex",
	})
	factomdmessagesServerFaultLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_LeaderExecute",
	})
	factomdmessagesServerFaultFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_FollowerExecute",
	})
	factomdmessagesServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_JSONByte",
	})
	factomdmessagesServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_JSONString",
	})
	factomdmessagesServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_ServerFault_IsSameAs",
	})
	factomdmessagesNewServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_NewServerFault_Summary",
		Help: "Summary of calls to  factomd_messages_NewServerFault",
	})
	factomdmessagesSignatureTimeoutIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_IsSameAs",
	})
	factomdmessagesSignatureTimeoutProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Process_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_Process",
	})
	factomdmessagesSignatureTimeoutGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetRepeatHash_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_GetRepeatHash",
	})
	factomdmessagesSignatureTimeoutGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetHash_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_GetHash",
	})
	factomdmessagesSignatureTimeoutGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetMsgHash_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_GetMsgHash",
	})
	factomdmessagesSignatureTimeoutGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_GetTimestamp",
	})
	factomdmessagesSignatureTimeoutType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Type_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_Type",
	})
	factomdmessagesSignatureTimeoutUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_UnmarshalBinaryData_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_UnmarshalBinaryData",
	})
	factomdmessagesSignatureTimeoutUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_UnmarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_UnmarshalBinary",
	})
	factomdmessagesSignatureTimeoutMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_MarshalForSignature_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_MarshalForSignature",
	})
	factomdmessagesSignatureTimeoutMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_MarshalBinary_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_MarshalBinary",
	})
	factomdmessagesSignatureTimeoutGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetSignature_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_GetSignature",
	})
	factomdmessagesSignatureTimeoutVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_VerifySignature",
	})
	factomdmessagesSignatureTimeoutSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Sign_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_Sign",
	})
	factomdmessagesSignatureTimeoutString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_String_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_String",
	})
	factomdmessagesSignatureTimeoutDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_DBHeight_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_DBHeight",
	})
	factomdmessagesSignatureTimeoutChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_ChainID_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_ChainID",
	})
	factomdmessagesSignatureTimeoutListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_ListHeight_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_ListHeight",
	})
	factomdmessagesSignatureTimeoutSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_SerialHash_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_SerialHash",
	})
	factomdmessagesSignatureTimeoutValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Validate_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_Validate",
	})
	factomdmessagesSignatureTimeoutComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_ComputeVMIndex",
	})
	factomdmessagesSignatureTimeoutLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_LeaderExecute",
	})
	factomdmessagesSignatureTimeoutFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_FollowerExecute_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_FollowerExecute",
	})
	factomdmessagesSignatureTimeoutJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_JSONByte_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_JSONByte",
	})
	factomdmessagesSignatureTimeoutJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_JSONString_Summary",
		Help: "Summary of calls to  factomd_messages_SignatureTimeout_JSONString",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdmessagesresend)
	prometheus.MustRegister(factomdmessagesMessageBaseSendOut)
	prometheus.MustRegister(factomdmessagesMessageBaseGetNoResend)
	prometheus.MustRegister(factomdmessagesMessageBaseSetNoResend)
	prometheus.MustRegister(factomdmessagesMessageBaseIsValid)
	prometheus.MustRegister(factomdmessagesMessageBaseSetValid)
	prometheus.MustRegister(factomdmessagesMessageBaseMarkSentInvalid)
	prometheus.MustRegister(factomdmessagesMessageBaseSentInvlaid)
	prometheus.MustRegister(factomdmessagesMessageBaseResend)
	prometheus.MustRegister(factomdmessagesMessageBaseExpire)
	prometheus.MustRegister(factomdmessagesMessageBaseIsStalled)
	prometheus.MustRegister(factomdmessagesMessageBaseSetStall)
	prometheus.MustRegister(factomdmessagesMessageBaseGetFullMsgHash)
	prometheus.MustRegister(factomdmessagesMessageBaseSetFullMsgHash)
	prometheus.MustRegister(factomdmessagesMessageBaseGetOrigin)
	prometheus.MustRegister(factomdmessagesMessageBaseSetOrigin)
	prometheus.MustRegister(factomdmessagesMessageBaseGetNetworkOrigin)
	prometheus.MustRegister(factomdmessagesMessageBaseSetNetworkOrigin)
	prometheus.MustRegister(factomdmessagesMessageBaseIsPeer2Peer)
	prometheus.MustRegister(factomdmessagesMessageBaseSetPeer2Peer)
	prometheus.MustRegister(factomdmessagesMessageBaseIsLocal)
	prometheus.MustRegister(factomdmessagesMessageBaseSetLocal)
	prometheus.MustRegister(factomdmessagesMessageBaseGetLeaderChainID)
	prometheus.MustRegister(factomdmessagesMessageBaseSetLeaderChainID)
	prometheus.MustRegister(factomdmessagesMessageBaseGetVMIndex)
	prometheus.MustRegister(factomdmessagesMessageBaseSetVMIndex)
	prometheus.MustRegister(factomdmessagesMessageBaseGetVMHash)
	prometheus.MustRegister(factomdmessagesMessageBaseSetVMHash)
	prometheus.MustRegister(factomdmessagesMessageBaseGetMinute)
	prometheus.MustRegister(factomdmessagesMessageBaseSetMinute)
	prometheus.MustRegister(factomdmessagesAckGetRepeatHash)
	prometheus.MustRegister(factomdmessagesAckGetHash)
	prometheus.MustRegister(factomdmessagesAckGetMsgHash)
	prometheus.MustRegister(factomdmessagesAckType)
	prometheus.MustRegister(factomdmessagesAckGetTimestamp)
	prometheus.MustRegister(factomdmessagesAckVerifySignature)
	prometheus.MustRegister(factomdmessagesAckValidate)
	prometheus.MustRegister(factomdmessagesAckComputeVMIndex)
	prometheus.MustRegister(factomdmessagesAckLeaderExecute)
	prometheus.MustRegister(factomdmessagesAckFollowerExecute)
	prometheus.MustRegister(factomdmessagesAckProcess)
	prometheus.MustRegister(factomdmessagesAckJSONByte)
	prometheus.MustRegister(factomdmessagesAckJSONString)
	prometheus.MustRegister(factomdmessagesAckSign)
	prometheus.MustRegister(factomdmessagesAckGetSignature)
	prometheus.MustRegister(factomdmessagesAckUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesAckUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesAckMarshalForSignature)
	prometheus.MustRegister(factomdmessagesAckMarshalBinary)
	prometheus.MustRegister(factomdmessagesAckString)
	prometheus.MustRegister(factomdmessagesAckIsSameAs)
	prometheus.MustRegister(factomdmessagesAddServerMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesAddServerMsgGetHash)
	prometheus.MustRegister(factomdmessagesAddServerMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesAddServerMsgType)
	prometheus.MustRegister(factomdmessagesAddServerMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesAddServerMsgValidate)
	prometheus.MustRegister(factomdmessagesAddServerMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesAddServerMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesAddServerMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesAddServerMsgProcess)
	prometheus.MustRegister(factomdmessagesAddServerMsgJSONByte)
	prometheus.MustRegister(factomdmessagesAddServerMsgJSONString)
	prometheus.MustRegister(factomdmessagesAddServerMsgSign)
	prometheus.MustRegister(factomdmessagesAddServerMsgGetSignature)
	prometheus.MustRegister(factomdmessagesAddServerMsgVerifySignature)
	prometheus.MustRegister(factomdmessagesAddServerMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesAddServerMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesAddServerMsgMarshalForSignature)
	prometheus.MustRegister(factomdmessagesAddServerMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesAddServerMsgString)
	prometheus.MustRegister(factomdmessagesAddServerMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesNewAddServerMsg)
	prometheus.MustRegister(factomdmessagesNewAddServerByHashMsg)
	prometheus.MustRegister(factomdmessagesAuditServerFaultGetRepeatHash)
	prometheus.MustRegister(factomdmessagesAuditServerFaultIsSameAs)
	prometheus.MustRegister(factomdmessagesAuditServerFaultSign)
	prometheus.MustRegister(factomdmessagesAuditServerFaultGetSignature)
	prometheus.MustRegister(factomdmessagesAuditServerFaultVerifySignature)
	prometheus.MustRegister(factomdmessagesAuditServerFaultProcess)
	prometheus.MustRegister(factomdmessagesAuditServerFaultGetTimestamp)
	prometheus.MustRegister(factomdmessagesAuditServerFaultGetHash)
	prometheus.MustRegister(factomdmessagesAuditServerFaultGetMsgHash)
	prometheus.MustRegister(factomdmessagesAuditServerFaultType)
	prometheus.MustRegister(factomdmessagesAuditServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesAuditServerFaultUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesAuditServerFaultMarshalForSignature)
	prometheus.MustRegister(factomdmessagesAuditServerFaultMarshalBinary)
	prometheus.MustRegister(factomdmessagesAuditServerFaultString)
	prometheus.MustRegister(factomdmessagesAuditServerFaultDBHeight)
	prometheus.MustRegister(factomdmessagesAuditServerFaultChainID)
	prometheus.MustRegister(factomdmessagesAuditServerFaultListHeight)
	prometheus.MustRegister(factomdmessagesAuditServerFaultSerialHash)
	prometheus.MustRegister(factomdmessagesAuditServerFaultValidate)
	prometheus.MustRegister(factomdmessagesAuditServerFaultComputeVMIndex)
	prometheus.MustRegister(factomdmessagesAuditServerFaultLeaderExecute)
	prometheus.MustRegister(factomdmessagesAuditServerFaultFollowerExecute)
	prometheus.MustRegister(factomdmessagesAuditServerFaultJSONByte)
	prometheus.MustRegister(factomdmessagesAuditServerFaultJSONString)
	prometheus.MustRegister(factomdmessagesBounceAddData)
	prometheus.MustRegister(factomdmessagesBounceGetRepeatHash)
	prometheus.MustRegister(factomdmessagesBounceGetHash)
	prometheus.MustRegister(factomdmessagesBounceSizeOf)
	prometheus.MustRegister(factomdmessagesBounceGetMsgHash)
	prometheus.MustRegister(factomdmessagesBounceType)
	prometheus.MustRegister(factomdmessagesBounceGetTimestamp)
	prometheus.MustRegister(factomdmessagesBounceVerifySignature)
	prometheus.MustRegister(factomdmessagesBounceValidate)
	prometheus.MustRegister(factomdmessagesBounceComputeVMIndex)
	prometheus.MustRegister(factomdmessagesBounceLeaderExecute)
	prometheus.MustRegister(factomdmessagesBounceFollowerExecute)
	prometheus.MustRegister(factomdmessagesBounceProcess)
	prometheus.MustRegister(factomdmessagesBounceJSONByte)
	prometheus.MustRegister(factomdmessagesBounceJSONString)
	prometheus.MustRegister(factomdmessagesBounceSign)
	prometheus.MustRegister(factomdmessagesBounceGetSignature)
	prometheus.MustRegister(factomdmessagesBounceUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesBounceUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesBounceMarshalForSignature)
	prometheus.MustRegister(factomdmessagesBounceMarshalBinary)
	prometheus.MustRegister(factomdmessagesBounceString)
	prometheus.MustRegister(factomdmessagesBounceIsSameAs)
	prometheus.MustRegister(factomdmessagesBounceReplyGetRepeatHash)
	prometheus.MustRegister(factomdmessagesBounceReplyGetHash)
	prometheus.MustRegister(factomdmessagesBounceReplySizeOf)
	prometheus.MustRegister(factomdmessagesBounceReplyGetMsgHash)
	prometheus.MustRegister(factomdmessagesBounceReplyType)
	prometheus.MustRegister(factomdmessagesBounceReplyGetTimestamp)
	prometheus.MustRegister(factomdmessagesBounceReplyVerifySignature)
	prometheus.MustRegister(factomdmessagesBounceReplyValidate)
	prometheus.MustRegister(factomdmessagesBounceReplyComputeVMIndex)
	prometheus.MustRegister(factomdmessagesBounceReplyLeaderExecute)
	prometheus.MustRegister(factomdmessagesBounceReplyFollowerExecute)
	prometheus.MustRegister(factomdmessagesBounceReplyProcess)
	prometheus.MustRegister(factomdmessagesBounceReplyJSONByte)
	prometheus.MustRegister(factomdmessagesBounceReplyJSONString)
	prometheus.MustRegister(factomdmessagesBounceReplySign)
	prometheus.MustRegister(factomdmessagesBounceReplyGetSignature)
	prometheus.MustRegister(factomdmessagesBounceReplyUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesBounceReplyUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesBounceReplyMarshalForSignature)
	prometheus.MustRegister(factomdmessagesBounceReplyMarshalBinary)
	prometheus.MustRegister(factomdmessagesBounceReplyString)
	prometheus.MustRegister(factomdmessagesBounceReplyIsSameAs)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgGetHash)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgType)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgValidate)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgProcess)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgJSONByte)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgJSONString)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgSign)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgGetSignature)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgVerifySignature)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgMarshalForSignature)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgString)
	prometheus.MustRegister(factomdmessagesChangeServerKeyMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesNewChangeServerKeyMsg)
	prometheus.MustRegister(factomdmessagesCommitChainMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetCount)
	prometheus.MustRegister(factomdmessagesCommitChainMsgIncCount)
	prometheus.MustRegister(factomdmessagesCommitChainMsgSetCount)
	prometheus.MustRegister(factomdmessagesCommitChainMsgProcess)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetHash)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesCommitChainMsgType)
	prometheus.MustRegister(factomdmessagesCommitChainMsgValidate)
	prometheus.MustRegister(factomdmessagesCommitChainMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesCommitChainMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesCommitChainMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesCommitChainMsgJSONByte)
	prometheus.MustRegister(factomdmessagesCommitChainMsgJSONString)
	prometheus.MustRegister(factomdmessagesCommitChainMsgSign)
	prometheus.MustRegister(factomdmessagesCommitChainMsgGetSignature)
	prometheus.MustRegister(factomdmessagesCommitChainMsgVerifySignature)
	prometheus.MustRegister(factomdmessagesCommitChainMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesCommitChainMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesCommitChainMsgMarshalForSignature)
	prometheus.MustRegister(factomdmessagesCommitChainMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesCommitChainMsgString)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetCount)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgIncCount)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgSetCount)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgProcess)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetHash)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgType)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgSign)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgGetSignature)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgVerifySignature)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgMarshalForSignature)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgString)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgValidate)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgJSONByte)
	prometheus.MustRegister(factomdmessagesCommitEntryMsgJSONString)
	prometheus.MustRegister(factomdmessagesNewCommitEntryMsg)
	prometheus.MustRegister(factomdmessagesDataResponseIsSameAs)
	prometheus.MustRegister(factomdmessagesDataResponseGetRepeatHash)
	prometheus.MustRegister(factomdmessagesDataResponseGetHash)
	prometheus.MustRegister(factomdmessagesDataResponseGetMsgHash)
	prometheus.MustRegister(factomdmessagesDataResponseType)
	prometheus.MustRegister(factomdmessagesDataResponseGetTimestamp)
	prometheus.MustRegister(factomdmessagesDataResponseValidate)
	prometheus.MustRegister(factomdmessagesDataResponseComputeVMIndex)
	prometheus.MustRegister(factomdmessagesDataResponseLeaderExecute)
	prometheus.MustRegister(factomdmessagesDataResponseFollowerExecute)
	prometheus.MustRegister(factomdmessagesDataResponseProcess)
	prometheus.MustRegister(factomdmessagesDataResponseJSONByte)
	prometheus.MustRegister(factomdmessagesDataResponseJSONString)
	prometheus.MustRegister(factomdmessagesDataResponseUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesDataResponseUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesattemptEntryUnmarshal)
	prometheus.MustRegister(factomdmessagesattemptEBlockUnmarshal)
	prometheus.MustRegister(factomdmessagesDataResponseMarshalBinary)
	prometheus.MustRegister(factomdmessagesDataResponseString)
	prometheus.MustRegister(factomdmessagesNewDataResponse)
	prometheus.MustRegister(factomdmessagesDBStateMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesDBStateMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesDBStateMsgGetHash)
	prometheus.MustRegister(factomdmessagesDBStateMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesDBStateMsgType)
	prometheus.MustRegister(factomdmessagesDBStateMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesDBStateMsgValidate)
	prometheus.MustRegister(factomdmessagesDBStateMsgSigTally)
	prometheus.MustRegister(factomdmessagesDBStateMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesDBStateMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesDBStateMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesDBStateMsgProcess)
	prometheus.MustRegister(factomdmessagesDBStateMsgJSONByte)
	prometheus.MustRegister(factomdmessagesDBStateMsgJSONString)
	prometheus.MustRegister(factomdmessagesDBStateMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesDBStateMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesDBStateMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesDBStateMsgString)
	prometheus.MustRegister(factomdmessagesNewDBStateMsg)
	prometheus.MustRegister(factomdmessagesDBStateMissingIsSameAs)
	prometheus.MustRegister(factomdmessagesDBStateMissingGetRepeatHash)
	prometheus.MustRegister(factomdmessagesDBStateMissingGetHash)
	prometheus.MustRegister(factomdmessagesDBStateMissingGetMsgHash)
	prometheus.MustRegister(factomdmessagesDBStateMissingType)
	prometheus.MustRegister(factomdmessagesDBStateMissingGetTimestamp)
	prometheus.MustRegister(factomdmessagesDBStateMissingValidate)
	prometheus.MustRegister(factomdmessagesDBStateMissingComputeVMIndex)
	prometheus.MustRegister(factomdmessagesDBStateMissingLeaderExecute)
	prometheus.MustRegister(factomdmessagesDBStateMissingsend)
	prometheus.MustRegister(factomdmessagesDBStateMissingFollowerExecute)
	prometheus.MustRegister(factomdmessagesDBStateMissingProcess)
	prometheus.MustRegister(factomdmessagesDBStateMissingJSONByte)
	prometheus.MustRegister(factomdmessagesDBStateMissingJSONString)
	prometheus.MustRegister(factomdmessagesDBStateMissingUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesDBStateMissingUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesDBStateMissingMarshalForSignature)
	prometheus.MustRegister(factomdmessagesDBStateMissingMarshalBinary)
	prometheus.MustRegister(factomdmessagesDBStateMissingString)
	prometheus.MustRegister(factomdmessagesNewDBStateMissing)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureIsSameAs)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureProcess)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureGetRepeatHash)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureGetHash)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureGetMsgHash)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureGetTimestamp)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureType)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureValidate)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureComputeVMIndex)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureLeaderExecute)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureFollowerExecute)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureSign)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureSignHeader)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureGetSignature)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureVerifySignature)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureMarshalForSignature)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureMarshalBinary)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureString)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureJSONByte)
	prometheus.MustRegister(factomdmessagesDirectoryBlockSignatureJSONString)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseIsSameAs)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseGetRepeatHash)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseGetHash)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseGetMsgHash)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseType)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseGetTimestamp)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseValidate)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseComputeVMIndex)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseLeaderExecute)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseFollowerExecute)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseProcess)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseJSONByte)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseJSONString)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseMarshalForSignature)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseMarshalBinary)
	prometheus.MustRegister(factomdmessagesEntryBlockResponseString)
	prometheus.MustRegister(factomdmessagesNewEntryBlockResponse)
	prometheus.MustRegister(factomdmessagesEOMIsSameAs)
	prometheus.MustRegister(factomdmessagesEOMProcess)
	prometheus.MustRegister(factomdmessagesEOMGetRepeatHash)
	prometheus.MustRegister(factomdmessagesEOMGetHash)
	prometheus.MustRegister(factomdmessagesEOMGetMsgHash)
	prometheus.MustRegister(factomdmessagesEOMGetTimestamp)
	prometheus.MustRegister(factomdmessagesEOMType)
	prometheus.MustRegister(factomdmessagesEOMValidate)
	prometheus.MustRegister(factomdmessagesEOMComputeVMIndex)
	prometheus.MustRegister(factomdmessagesEOMLeaderExecute)
	prometheus.MustRegister(factomdmessagesEOMFollowerExecute)
	prometheus.MustRegister(factomdmessagesEOMJSONByte)
	prometheus.MustRegister(factomdmessagesEOMJSONString)
	prometheus.MustRegister(factomdmessagesEOMSign)
	prometheus.MustRegister(factomdmessagesEOMGetSignature)
	prometheus.MustRegister(factomdmessagesEOMVerifySignature)
	prometheus.MustRegister(factomdmessagesEOMUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesEOMUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesEOMMarshalForSignature)
	prometheus.MustRegister(factomdmessagesEOMMarshalBinary)
	prometheus.MustRegister(factomdmessagesEOMString)
	prometheus.MustRegister(factomdmessagesEOMTimeoutIsSameAs)
	prometheus.MustRegister(factomdmessagesEOMTimeoutSign)
	prometheus.MustRegister(factomdmessagesEOMTimeoutGetSignature)
	prometheus.MustRegister(factomdmessagesEOMTimeoutVerifySignature)
	prometheus.MustRegister(factomdmessagesEOMTimeoutProcess)
	prometheus.MustRegister(factomdmessagesEOMTimeoutGetRepeatHash)
	prometheus.MustRegister(factomdmessagesEOMTimeoutGetHash)
	prometheus.MustRegister(factomdmessagesEOMTimeoutGetMsgHash)
	prometheus.MustRegister(factomdmessagesEOMTimeoutGetTimestamp)
	prometheus.MustRegister(factomdmessagesEOMTimeoutType)
	prometheus.MustRegister(factomdmessagesEOMTimeoutUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesEOMTimeoutUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesEOMTimeoutMarshalForSignature)
	prometheus.MustRegister(factomdmessagesEOMTimeoutMarshalBinary)
	prometheus.MustRegister(factomdmessagesEOMTimeoutString)
	prometheus.MustRegister(factomdmessagesEOMTimeoutDBHeight)
	prometheus.MustRegister(factomdmessagesEOMTimeoutChainID)
	prometheus.MustRegister(factomdmessagesEOMTimeoutListHeight)
	prometheus.MustRegister(factomdmessagesEOMTimeoutSerialHash)
	prometheus.MustRegister(factomdmessagesEOMTimeoutValidate)
	prometheus.MustRegister(factomdmessagesEOMTimeoutComputeVMIndex)
	prometheus.MustRegister(factomdmessagesEOMTimeoutLeaderExecute)
	prometheus.MustRegister(factomdmessagesEOMTimeoutFollowerExecute)
	prometheus.MustRegister(factomdmessagesEOMTimeoutJSONByte)
	prometheus.MustRegister(factomdmessagesEOMTimeoutJSONString)
	prometheus.MustRegister(factomdmessagesMessageErrorError)
	prometheus.MustRegister(factomdmessagesmessageError)
	prometheus.MustRegister(factomdmessagesFactoidTransactionIsSameAs)
	prometheus.MustRegister(factomdmessagesFactoidTransactionGetRepeatHash)
	prometheus.MustRegister(factomdmessagesFactoidTransactionGetHash)
	prometheus.MustRegister(factomdmessagesFactoidTransactionGetMsgHash)
	prometheus.MustRegister(factomdmessagesFactoidTransactionGetTimestamp)
	prometheus.MustRegister(factomdmessagesFactoidTransactionGetTransaction)
	prometheus.MustRegister(factomdmessagesFactoidTransactionSetTransaction)
	prometheus.MustRegister(factomdmessagesFactoidTransactionType)
	prometheus.MustRegister(factomdmessagesFactoidTransactionValidate)
	prometheus.MustRegister(factomdmessagesFactoidTransactionComputeVMIndex)
	prometheus.MustRegister(factomdmessagesFactoidTransactionLeaderExecute)
	prometheus.MustRegister(factomdmessagesFactoidTransactionFollowerExecute)
	prometheus.MustRegister(factomdmessagesFactoidTransactionProcess)
	prometheus.MustRegister(factomdmessagesFactoidTransactionUnmarshalTransData)
	prometheus.MustRegister(factomdmessagesFactoidTransactionUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesFactoidTransactionUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesFactoidTransactionMarshalBinary)
	prometheus.MustRegister(factomdmessagesFactoidTransactionString)
	prometheus.MustRegister(factomdmessagesFactoidTransactionJSONByte)
	prometheus.MustRegister(factomdmessagesFactoidTransactionJSONString)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetAmINegotiator)
	prometheus.MustRegister(factomdmessagesFullServerFaultSetAmINegotiator)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetMyVoteTallied)
	prometheus.MustRegister(factomdmessagesFullServerFaultSetMyVoteTallied)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetPledgeDone)
	prometheus.MustRegister(factomdmessagesFullServerFaultSetPledgeDone)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetLastMatch)
	prometheus.MustRegister(factomdmessagesFullServerFaultSetLastMatch)
	prometheus.MustRegister(factomdmessagesFullServerFaultIsNil)
	prometheus.MustRegister(factomdmessagesFullServerFaultAddFaultVote)
	prometheus.MustRegister(factomdmessagesFullServerFaultPriority)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetSerialHash)
	prometheus.MustRegister(factomdmessagesFullServerFaultProcess)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetRepeatHash)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetHash)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetMsgHash)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetCoreHash)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetTimestamp)
	prometheus.MustRegister(factomdmessagesFullServerFaultType)
	prometheus.MustRegister(factomdmessagesFullServerFaultMarshalCore)
	prometheus.MustRegister(factomdmessagesFullServerFaultMarshalForSF)
	prometheus.MustRegister(factomdmessagesFullServerFaultMarshalForSignature)
	prometheus.MustRegister(factomdmessagesSigListMarshalBinary)
	prometheus.MustRegister(factomdmessagesSigListUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesFullServerFaultMarshalBinary)
	prometheus.MustRegister(factomdmessagesUnmarshall)
	prometheus.MustRegister(factomdmessagesFullServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesFullServerFaultUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetSignature)
	prometheus.MustRegister(factomdmessagesFullServerFaultVerifySignature)
	prometheus.MustRegister(factomdmessagesFullServerFaultSign)
	prometheus.MustRegister(factomdmessagesFullServerFaultString)
	prometheus.MustRegister(factomdmessagesFullServerFaultStringWithSigCnt)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetDBHeight)
	prometheus.MustRegister(factomdmessagesFullServerFaultValidate)
	prometheus.MustRegister(factomdmessagesFullServerFaultSetAlreadyProcessed)
	prometheus.MustRegister(factomdmessagesFullServerFaultGetAlreadyProcessed)
	prometheus.MustRegister(factomdmessagesFullServerFaultHasEnoughSigs)
	prometheus.MustRegister(factomdmessagesFullServerFaultSigTally)
	prometheus.MustRegister(factomdmessagesFullServerFaultComputeVMIndex)
	prometheus.MustRegister(factomdmessagesFullServerFaultLeaderExecute)
	prometheus.MustRegister(factomdmessagesFullServerFaultFollowerExecute)
	prometheus.MustRegister(factomdmessagesFullServerFaultJSONByte)
	prometheus.MustRegister(factomdmessagesFullServerFaultJSONString)
	prometheus.MustRegister(factomdmessagesFullServerFaultIsSameAs)
	prometheus.MustRegister(factomdmessagesFullServerFaultToAdminBlockEntry)
	prometheus.MustRegister(factomdmessagesNewFullServerFault)
	prometheus.MustRegister(factomdmessagesUnmarshalMessage)
	prometheus.MustRegister(factomdmessagesUnmarshalMessageData)
	prometheus.MustRegister(factomdmessagesMessageName)
	prometheus.MustRegister(factomdmessagesSignSignable)
	prometheus.MustRegister(factomdmessagesVerifyMessage)
	prometheus.MustRegister(factomdmessagesHeartbeatIsSameAs)
	prometheus.MustRegister(factomdmessagesHeartbeatProcess)
	prometheus.MustRegister(factomdmessagesHeartbeatGetRepeatHash)
	prometheus.MustRegister(factomdmessagesHeartbeatGetHash)
	prometheus.MustRegister(factomdmessagesHeartbeatGetMsgHash)
	prometheus.MustRegister(factomdmessagesHeartbeatGetTimestamp)
	prometheus.MustRegister(factomdmessagesHeartbeatType)
	prometheus.MustRegister(factomdmessagesHeartbeatUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesHeartbeatUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesHeartbeatMarshalForSignature)
	prometheus.MustRegister(factomdmessagesHeartbeatMarshalBinary)
	prometheus.MustRegister(factomdmessagesHeartbeatString)
	prometheus.MustRegister(factomdmessagesHeartbeatChainID)
	prometheus.MustRegister(factomdmessagesHeartbeatListHeight)
	prometheus.MustRegister(factomdmessagesHeartbeatSerialHash)
	prometheus.MustRegister(factomdmessagesHeartbeatValidate)
	prometheus.MustRegister(factomdmessagesHeartbeatComputeVMIndex)
	prometheus.MustRegister(factomdmessagesHeartbeatLeaderExecute)
	prometheus.MustRegister(factomdmessagesHeartbeatFollowerExecute)
	prometheus.MustRegister(factomdmessagesHeartbeatJSONByte)
	prometheus.MustRegister(factomdmessagesHeartbeatJSONString)
	prometheus.MustRegister(factomdmessagesHeartbeatSign)
	prometheus.MustRegister(factomdmessagesHeartbeatGetSignature)
	prometheus.MustRegister(factomdmessagesHeartbeatVerifySignature)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockIsSameAs)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockSign)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockGetSignature)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockVerifySignature)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockProcess)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockGetRepeatHash)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockGetHash)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockGetMsgHash)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockGetTimestamp)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockType)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockMarshalBinary)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockMarshalForSignature)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockString)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockDBHeight)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockChainID)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockListHeight)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockSerialHash)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockValidate)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockComputeVMIndex)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockLeaderExecute)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockFollowerExecute)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockJSONByte)
	prometheus.MustRegister(factomdmessagesInvalidDirectoryBlockJSONString)
	prometheus.MustRegister(factomdmessagesMissingDataIsSameAs)
	prometheus.MustRegister(factomdmessagesMissingDataProcess)
	prometheus.MustRegister(factomdmessagesMissingDataGetRepeatHash)
	prometheus.MustRegister(factomdmessagesMissingDataGetHash)
	prometheus.MustRegister(factomdmessagesMissingDataGetMsgHash)
	prometheus.MustRegister(factomdmessagesMissingDataGetTimestamp)
	prometheus.MustRegister(factomdmessagesMissingDataType)
	prometheus.MustRegister(factomdmessagesMissingDataUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesMissingDataUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingDataMarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingDataString)
	prometheus.MustRegister(factomdmessagesMissingDataValidate)
	prometheus.MustRegister(factomdmessagesMissingDataComputeVMIndex)
	prometheus.MustRegister(factomdmessagesMissingDataLeaderExecute)
	prometheus.MustRegister(factomdmessagesMissingDataFollowerExecute)
	prometheus.MustRegister(factomdmessagesMissingDataJSONByte)
	prometheus.MustRegister(factomdmessagesMissingDataJSONString)
	prometheus.MustRegister(factomdmessagesNewMissingData)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksIsSameAs)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksGetRepeatHash)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksGetHash)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksGetMsgHash)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksType)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksGetTimestamp)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksValidate)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksComputeVMIndex)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksLeaderExecute)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksFollowerExecute)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksProcess)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksJSONByte)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksJSONString)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksMarshalForSignature)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksMarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingEntryBlocksString)
	prometheus.MustRegister(factomdmessagesNewMissingEntryBlocks)
	prometheus.MustRegister(factomdmessagesMissingMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesMissingMsgProcess)
	prometheus.MustRegister(factomdmessagesMissingMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesMissingMsgGetHash)
	prometheus.MustRegister(factomdmessagesMissingMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesMissingMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesMissingMsgType)
	prometheus.MustRegister(factomdmessagesMissingMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesMissingMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingMsgString)
	prometheus.MustRegister(factomdmessagesMissingMsgChainID)
	prometheus.MustRegister(factomdmessagesMissingMsgListHeight)
	prometheus.MustRegister(factomdmessagesMissingMsgValidate)
	prometheus.MustRegister(factomdmessagesMissingMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesMissingMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesMissingMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesMissingMsgJSONByte)
	prometheus.MustRegister(factomdmessagesMissingMsgJSONString)
	prometheus.MustRegister(factomdmessagesMissingMsgAddHeight)
	prometheus.MustRegister(factomdmessagesNewMissingMsg)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseIsSameAs)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseProcess)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseGetRepeatHash)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseGetHash)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseGetMsgHash)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseGetTimestamp)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseType)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseMarshalBinary)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseString)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseChainID)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseListHeight)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseValidate)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseComputeVMIndex)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseLeaderExecute)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseFollowerExecute)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseJSONByte)
	prometheus.MustRegister(factomdmessagesMissingMsgResponseJSONString)
	prometheus.MustRegister(factomdmessagesNewMissingMsgResponse)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgGetHash)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgType)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgValidate)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgProcess)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgJSONByte)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgJSONString)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgSign)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgGetSignature)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgVerifySignature)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgMarshalForSignature)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgString)
	prometheus.MustRegister(factomdmessagesRemoveServerMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesNewRemoveServerMsg)
	prometheus.MustRegister(factomdmessagesRequestBlockIsSameAs)
	prometheus.MustRegister(factomdmessagesRequestBlockProcess)
	prometheus.MustRegister(factomdmessagesRequestBlockGetRepeatHash)
	prometheus.MustRegister(factomdmessagesRequestBlockGetHash)
	prometheus.MustRegister(factomdmessagesRequestBlockGetMsgHash)
	prometheus.MustRegister(factomdmessagesRequestBlockGetTimestamp)
	prometheus.MustRegister(factomdmessagesRequestBlockType)
	prometheus.MustRegister(factomdmessagesRequestBlockUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesRequestBlockUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesRequestBlockMarshalForSignature)
	prometheus.MustRegister(factomdmessagesRequestBlockMarshalBinary)
	prometheus.MustRegister(factomdmessagesRequestBlockString)
	prometheus.MustRegister(factomdmessagesRequestBlockDBHeight)
	prometheus.MustRegister(factomdmessagesRequestBlockChainID)
	prometheus.MustRegister(factomdmessagesRequestBlockListHeight)
	prometheus.MustRegister(factomdmessagesRequestBlockSerialHash)
	prometheus.MustRegister(factomdmessagesRequestBlockSignature)
	prometheus.MustRegister(factomdmessagesRequestBlockValidate)
	prometheus.MustRegister(factomdmessagesRequestBlockComputeVMIndex)
	prometheus.MustRegister(factomdmessagesRequestBlockLeaderExecute)
	prometheus.MustRegister(factomdmessagesRequestBlockFollowerExecute)
	prometheus.MustRegister(factomdmessagesRequestBlockJSONByte)
	prometheus.MustRegister(factomdmessagesRequestBlockJSONString)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgIsSameAs)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgProcess)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgGetRepeatHash)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgGetHash)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgGetMsgHash)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgGetChainIDHash)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgGetTimestamp)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgType)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgValidateRTN)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgValidate)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgComputeVMIndex)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgLeaderExecute)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgFollowerExecute)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgJSONByte)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgJSONString)
	prometheus.MustRegister(factomdmessagesNewRevealEntryMsg)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgMarshalBinary)
	prometheus.MustRegister(factomdmessagesRevealEntryMsgString)
	prometheus.MustRegister(factomdmessagesServerFaultProcess)
	prometheus.MustRegister(factomdmessagesServerFaultGetRepeatHash)
	prometheus.MustRegister(factomdmessagesServerFaultGetHash)
	prometheus.MustRegister(factomdmessagesServerFaultGetMsgHash)
	prometheus.MustRegister(factomdmessagesServerFaultGetCoreHash)
	prometheus.MustRegister(factomdmessagesServerFaultGetTimestamp)
	prometheus.MustRegister(factomdmessagesServerFaultType)
	prometheus.MustRegister(factomdmessagesServerFaultMarshalForSignature)
	prometheus.MustRegister(factomdmessagesServerFaultPreMarshalBinary)
	prometheus.MustRegister(factomdmessagesServerFaultMarshalBinary)
	prometheus.MustRegister(factomdmessagesServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesServerFaultUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesServerFaultGetSignature)
	prometheus.MustRegister(factomdmessagesServerFaultVerifySignature)
	prometheus.MustRegister(factomdmessagesServerFaultSign)
	prometheus.MustRegister(factomdmessagesServerFaultString)
	prometheus.MustRegister(factomdmessagesServerFaultGetDBHeight)
	prometheus.MustRegister(factomdmessagesServerFaultValidate)
	prometheus.MustRegister(factomdmessagesServerFaultComputeVMIndex)
	prometheus.MustRegister(factomdmessagesServerFaultLeaderExecute)
	prometheus.MustRegister(factomdmessagesServerFaultFollowerExecute)
	prometheus.MustRegister(factomdmessagesServerFaultJSONByte)
	prometheus.MustRegister(factomdmessagesServerFaultJSONString)
	prometheus.MustRegister(factomdmessagesServerFaultIsSameAs)
	prometheus.MustRegister(factomdmessagesNewServerFault)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutIsSameAs)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutProcess)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutGetRepeatHash)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutGetHash)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutGetMsgHash)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutGetTimestamp)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutType)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutUnmarshalBinaryData)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutUnmarshalBinary)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutMarshalForSignature)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutMarshalBinary)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutGetSignature)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutVerifySignature)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutSign)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutString)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutDBHeight)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutChainID)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutListHeight)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutSerialHash)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutValidate)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutComputeVMIndex)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutLeaderExecute)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutFollowerExecute)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutJSONByte)
	prometheus.MustRegister(factomdmessagesSignatureTimeoutJSONString)
}
