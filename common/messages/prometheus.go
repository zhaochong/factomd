// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import "github.com/prometheus/client_golang/prometheus"

var (
	//ack.go  variables
	messageErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_message_error_responses",
		Help: "Number of error responses from message objects",
	})
	messagesAckGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_GetRepeatHash_Summary",
		Help: "Summary of messages.ack.GetRepeatHash call",
	})
    messagesAckGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_GetHash_Summary",
		Help: "Summary of messages.ack.GetHash call",
	})
    messagesAckGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_GetMsgHash_Summary",
		Help: "Summary of messages.ack.GetMsgHash call",
	})
    messagesAckType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Type_Summary",
		Help: "Summary of messages.ack.Type call",
	})
    messagesAckInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Int_Summary",
		Help: "Summary of messages.ack.Int call",
	})
    messagesAckBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Bytes_Summary",
		Help: "Summary of messages.ack.Bytes call",
	})
    messagesAckGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_GetTimestamp_Summary",
		Help: "Summary of messages.ack.GetTimestamp call",
	})
    messagesAckVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_VerifySignature_Summary",
		Help: "Summary of messages.ack.VerifySignature call",
	})
    messagesAckValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Validate_Summary",
		Help: "Summary of messages.ack.Validate call",
	})
    messagesAckLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_LeaderExecute_Summary",
		Help: "Summary of messages.ack.LeaderExecute call",
	})
    messagesAckFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_FollowerExecute_Summary",
		Help: "Summary of messages.ack.FollowerExecute call",
	})
    messagesAckProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Process_Summary",
		Help: "Summary of messages.ack.Process call",
	})
    messagesAckJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_JSONByte_Summary",
		Help: "Summary of messages.ack.JSONByte call",
	})
    messagesAckJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_JSONString_Summary",
		Help: "Summary of messages.ack.JSONString call",
	})
    messagesAckJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_JSONBuffer_Summary",
		Help: "Summary of messages.ack.JSONBuffer call",
	})
    messagesAckSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_Sign_Summary",
		Help: "Summary of messages.ack.Sign call",
	})
    messagesAckGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_GetSignature_Summary",
		Help: "Summary of messages.ack.GetSignature call",
	})
    messagesAckUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.ack.UnmarshalBinaryData call",
	})
    messagesAckUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_UnmarshalBinary_Summary",
		Help: "Summary of messages.ack.UnmarshalBinary call",
	})
    messagesAckMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_MarshalForSignature_Summary",
		Help: "Summary of messages.ack.MarshalForSignature call",
	})
    messagesAckMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_MarshalBinary_Summary",
		Help: "Summary of messages.ack.MarshalBinary call",
	})
    messagesAckString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_String_Summary",
		Help: "Summary of messages.ack.String call",
	})
    messagesAckIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ack_IsSameAs_Summary",
		Help: "Summary of messages.ack.IsSameAs call",
	})
// addServer variables
   messagesAddServerGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_GetRepeatHash_Summary",
		Help: "Summary of messages.addServer.GetRepeatHash call",
	})
   messagesAddServerGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_GetHash_Summary",
		Help: "Summary of messages.addServer.GetHash call",
	})
   messagesAddServerGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_GetMsgHash_Summary",
		Help: "Summary of messages.addServer.GetMsgHash call",
	})
   messagesAddServerType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Type_Summary",
		Help: "Summary of messages.addServer.Type call",
	})
   messagesAddServerInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Int_Summary",
		Help: "Summary of messages.addServer.Int call",
	})
   messagesAddServerBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Bytes_Summary",
		Help: "Summary of messages.addServer.Bytes call",
	})
   messagesAddServerGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_GetTimestamp_Summary",
		Help: "Summary of messages.addServer.GetTimestamp call",
	})
   messagesAddServerValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Validate_Summary",
		Help: "Summary of messages.addServer.Validate call",
	})
   messagesAddServerComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_ComputeVMIndex_Summary",
		Help: "Summary of messages.addServer.ComputeVMIndex call",
	})
   messagesAddServerLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_LeaderExecute_Summary",
		Help: "Summary of messages.addServer.LeaderExecute call",
	})
   messagesAddServerFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_FollowerExecute_Summary",
		Help: "Summary of messages.addServer.FollowerExecute call",
	})
   messagesAddServerProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Process_Summary",
		Help: "Summary of messages.addServer.Process call",
	})
   messagesAddServerJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_JSONByte_Summary",
		Help: "Summary of messages.addServer.JSONByte call",
	})
   messagesAddServerJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_JSONString_Summary",
		Help: "Summary of messages.addServer.JSONString call",
	})
   messagesAddServerJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_JSONBuffer_Summary",
		Help: "Summary of messages.addServer.JSONBuffer call",
	})
   messagesAddServerSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_Sign_Summary",
		Help: "Summary of messages.addServer.Sign call",
	})
   messagesAddServerGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_GetSignature_Summary",
		Help: "Summary of messages.addServer.GetSignature call",
	})
   messagesAddServerVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_VerifySignature_Summary",
		Help: "Summary of messages.addServer.VerifySignature call",
	})
   messagesAddServerUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.addServer.UnmarshalBinaryData call",
	})
   messagesAddServerUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_UnmarshalBinary_Summary",
		Help: "Summary of messages.addServer.UnmarshalBinary call",
	})
   messagesAddServerMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_MarshalForSignature_Summary",
		Help: "Summary of messages.addServer.MarshalForSignature call",
	})
   messagesAddServerMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_MarshalBinary_Summary",
		Help: "Summary of messages.addServer.MarshalBinary call",
	})
   messagesAddServerString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_String_Summary",
		Help: "Summary of messages.addServer.String call",
	})
   messagesAddServerIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_IsSameAs_Summary",
		Help: "Summary of messages.addServer.IsSameAs call",
	})
   messagesAddServerNewAddServerMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_NewAddServerMsg_Summary",
		Help: "Summary of messages.addServer.NewAddServerMsg call",
	})
  messagesAddServerNewAddServerByHashMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AddServer_NewAddServerByHashMsg_Summary",
		Help: "Summary of messages.addServer.NewAddServerByHashMsg call",
	})
// AddServerFault
  messagesAuditServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetRepeatHash_Summary",
		Help: "Summary of messages.AuditServerFault.GetRepeatHash call",
	})
  messagesAuditServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_IsSameAs_Summary",
		Help: "Summary of messages.AuditServerFault.IsSameAs call",
	})
  messagesAuditServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Sign_Summary",
		Help: "Summary of messages.AuditServerFault.Sign call",
	})
  messagesAuditServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetSignature_Summary",
		Help: "Summary of messages.AuditServerFault.GetSignature call",
	})
  messagesAuditServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_VerifySignature_Summary",
		Help: "Summary of messages.AuditServerFault.VerifySignature call",
	})
  messagesAuditServerFaultProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Process_Summary",
		Help: "Summary of messages.AuditServerFault.Process call",
	})
  messagesAuditServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetTimestamp_Summary",
		Help: "Summary of messages.AuditServerFault.GetTimestamp call",
	})
  messagesAuditServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetHash_Summary",
		Help: "Summary of messages.AuditServerFault.GetHash call",
	})
  messagesAuditServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_GetMsgHash_Summary",
		Help: "Summary of messages.AuditServerFault.GetMsgHash call",
	})
  messagesAuditServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Type_Summary",
		Help: "Summary of messages.AuditServerFault.Type call",
	})
  messagesAuditServerFaultInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Int_Summary",
		Help: "Summary of messages.AuditServerFault.Int call",
	})
  messagesAuditServerFaultBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Bytes_Summary",
		Help: "Summary of messages.AuditServerFault.Bytes call",
	})
  messagesAuditServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.AuditServerFault.UnmarshalBinaryData call",
	})
  messagesAuditServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_UnmarshalBinary_Summary",
		Help: "Summary of messages.AuditServerFault.UnmarshalBinary call",
	})
  messagesAuditServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_MarshalForSignature_Summary",
		Help: "Summary of messages.AuditServerFault.MarshalForSignature call",
	})
  messagesAuditServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_MarshalBinary_Summary",
		Help: "Summary of messages.AuditServerFault.MarshalBinary call",
	})
  messagesAuditServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_String_Summary",
		Help: "Summary of messages.AuditServerFault.String call",
	})
  messagesAuditServerFaultDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_DBHeight_Summary",
		Help: "Summary of messages.AuditServerFault.DBHeight call",
	})
  messagesAuditServerFaultChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_ChainID_Summary",
		Help: "Summary of messages.AuditServerFault.ChainID call",
	})
  messagesAuditServerFaultListHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_ListHeight_Summary",
		Help: "Summary of messages.AuditServerFault.ListHeight call",
	})
  messagesAuditServerFaultSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_SerialHash_Summary",
		Help: "Summary of messages.AuditServerFault.SerialHash call",
	})
  messagesAuditServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_Validate_Summary",
		Help: "Summary of messages.AuditServerFault.Validate call",
	})
  messagesAuditServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_JSONByte_Summary",
		Help: "Summary of messages.AuditServerFault.JSONByte call",
	})
  messagesAuditServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_JSONString_Summary",
		Help: "Summary of messages.AuditServerFault.JSONString call",
	})
  messagesAuditServerFaultJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_AuditServerFault_JSONBuffer_Summary",
		Help: "Summary of messages.AuditServerFault.JSONBuffer call",
	})
//boucne.go values
  messagesBounceAddData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_AddData_Summary",
		Help: "Summary of messages.Bounce.AddData call",
	})
  messagesBounceGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetRepeatHash_Summary",
		Help: "Summary of messages.Bounce.GetRepeatHash call",
	})
  messagesBounceGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetHash_Summary",
		Help: "Summary of messages.Bounce.GetHash call",
	})
  messagesBounceSizeOf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_SizeOf_Summary",
		Help: "Summary of messages.Bounce.SizeOf call",
	})
  messagesBounceGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetMsgHash_Summary",
		Help: "Summary of messages.Bounce.GetMsgHash call",
	})
  messagesBounceType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_Type_Summary",
		Help: "Summary of messages.Bounce.Type call",
	})
  messagesBounceGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_GetTimestamp_Summary",
		Help: "Summary of messages.Bounce.GetTimestamp call",
	})
  messagesBounceVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_VerifySignature_Summary",
		Help: "Summary of messages.Bounce.VerifySignature call",
	})
  messagesBounceJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_JSONByte_Summary",
		Help: "Summary of messages.Bounce.JSONByte call",
	})
  messagesBounceJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_JSONString_Summary",
		Help: "Summary of messages.Bounce.JSONString call",
	})
  messagesBounceJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_JSONBuffer_Summary",
		Help: "Summary of messages.Bounce.JSONBuffer call",
	})
  messagesBounceUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.Bounce.UnmarshalBinaryData call",
	})
  messagesBounceUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_UnmarshalBinary_Summary",
		Help: "Summary of messages.Bounce.UnmarshalBinary call",
	})
  messagesBounceMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_MarshalForSignature_Summary",
		Help: "Summary of messages.Bounce.MarshalForSignature call",
	})
  messagesBounceMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_MarshalBinary_Summary",
		Help: "Summary of messages.Bounce.MarshalBinary call",
	})
  messagesBounceString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Bounce_String_Summary",
		Help: "Summary of messages.Bounce.String call",
	})
//boucnereply
  messagesBounceReplyGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_GetRepeatHash_Summary",
		Help: "Summary of messages.BounceReply.GetRepeatHash call",
	})
  messagesBounceReplyGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_GetHash_Summary",
		Help: "Summary of messages.BounceReply.GetHash call",
	})
  messagesBounceReplySizeOf = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_SizeOf_Summary",
		Help: "Summary of messages.BounceReply.SizeOf call",
	})
  messagesBounceReplyGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_GetMsgHash_Summary",
		Help: "Summary of messages.BounceReply.GetMsgHash call",
	})
  messagesBounceReplyType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_Type_Summary",
		Help: "Summary of messages.BounceReply.Type call",
	})
  messagesBounceReplyGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_GetTimestamp_Summary",
		Help: "Summary of messages.BounceReply.GetTimestamp call",
	})
  messagesBounceReplyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_JSONByte_Summary",
		Help: "Summary of messages.BounceReply.JSONByte call",
	})
  messagesBounceReplyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_JSONString_Summary",
		Help: "Summary of messages.BounceReply.JSONString call",
	})
  messagesBounceReplyJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_JSONBuffer_Summary",
		Help: "Summary of messages.BounceReply.JSONBuffer call",
	})
  messagesBounceReplyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.BounceReply.UnmarshalBinaryData call",
	})
  messagesBounceReplyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_UnmarshalBinary_Summary",
		Help: "Summary of messages.BounceReply.UnmarshalBinary call",
	})
  messagesBounceReplyMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_MarshalForSignature_Summary",
		Help: "Summary of messages.BounceReply.MarshalForSignature call",
	})
  messagesBounceReplyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_MarshalBinary_Summary",
		Help: "Summary of messages.BounceReply.MarshalBinary call",
	})
  messagesBounceReplyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_bouncereply_String_Summary",
		Help: "Summary of messages.BounceReply.String call",
	})
//changeServerKey
  messagesChangeServerKeyGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_GetRepeatHash_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.GetRepeatHash call",
	})
  messagesChangeServerKeyGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_GetHash_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.GetHash call",
	})
  messagesChangeServerKeyGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_GetMsgHash_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.GetMsgHash call",
	})
  messagesChangeServerKeyType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_Type_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.Type call",
	})
  messagesChangeServerKeyGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_GetTimestamp_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.GetTimestamp call",
	})
  messagesChangeServerKeyValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_Validate_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.Validate call",
	})
  messagesChangeServerKeyComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_ComputeVMIndex_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.ComputeVMIndex call",
	})
  messagesChangeServerKeyLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_LeaderExecute_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.LeaderExecute call",
	})
  messagesChangeServerKeyFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_FollowerExecute_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.FollowerExecute call",
	})
  messagesChangeServerKeyProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_Process_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.Process call",
	})
  messagesChangeServerKeyJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_JSONString_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.JSONString call",
	})
  messagesChangeServerKeyJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_JSONBuffer_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.JSONBuffer call",
	})
  messagesChangeServerKeyJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_JSONByte_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.JSONByte call",
	})
  messagesChangeServerKeySign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_Sign_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.Sign call",
	})
  messagesChangeServerKeyVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_VerifySignature_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.VerifySignature call",
	})
  messagesChangeServerKeyUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg. UnmarshalBinaryDatacall",
	})
  messagesChangeServerKeyUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.UnmarshalBinary call",
	})
  messagesChangeServerKeyMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_MarshalForSignature_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.MarshalForSignature call",
	})
  messagesChangeServerKeyMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_MarshalBinary_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.MarshalBinary call",
	})
  messagesChangeServerKeyString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_String_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.String call",
	})
  messagesChangeServerKeyIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_IsSameAs_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.IsSameAs call",
	})
  messagesChangeServerKeyNewChangeServerKeyMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_changeserverkeymsg_NewChangeServerKeyMsg_Summary",
		Help: "Summary of messages.ChangeServerKeyMsg.NewChangeServerKeyMsg call",
	})
 
 //commitChain.go
  messagesCommitChainMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_IsSameAs_Summary",
		Help: "Summary of messages.CommitChainMsg.IsSameAs call",
	})
  messagesCommitChainMsgGetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetCount_Summary",
		Help: "Summary of messages.CommitChainMsg.GetCount call",
	})
  messagesCommitChainMsgIncCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_IncCount_Summary",
		Help: "Summary of messages.CommitChainMsg.IncCount call",
	})
  messagesCommitChainMsgSetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_SetCount_Summary",
		Help: "Summary of messages.CommitChainMsg.SetCount call",
	})
  messagesCommitChainMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_Process_Summary",
		Help: "Summary of messages.CommitChainMsg.Process call",
	})
  messagesCommitChainMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetRepeatHash_Summary",
		Help: "Summary of messages.CommitChainMsg.GetRepeatHash call",
	})
  messagesCommitChainMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetHash_Summary",
		Help: "Summary of messages.CommitChainMsg.GetHash call",
	})
  messagesCommitChainMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetMsgHash_Summary",
		Help: "Summary of messages.CommitChainMsg.GetMsgHash call",
	})
  messagesCommitChainMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetTimestamp_Summary",
		Help: "Summary of messages.CommitChainMsg.GetTimestamp call",
	})
  messagesCommitChainMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_Type_Summary",
		Help: "Summary of messages.CommitChainMsg. Typecall",
	})
  messagesCommitChainMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_Validate_Summary",
		Help: "Summary of messages.CommitChainMsg.Validate call",
	})
  messagesCommitChainMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_ComputeVMIndex_Summary",
		Help: "Summary of messages.CommitChainMsg.ComputeVMIndex call",
	})
  messagesCommitChainMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_LeaderExecute_Summary",
		Help: "Summary of messages.CommitChainMsg.LeaderExecute call",
	})
  messagesCommitChainMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_FollowerExecute_Summary",
		Help: "Summary of messages.CommitChainMsg.FollowerExecute call",
	})
  messagesCommitChainMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_JSONByte_Summary",
		Help: "Summary of messages.CommitChainMsg.JSONByte call",
	})
  messagesCommitChainMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_JSONString_Summary",
		Help: "Summary of messages.CommitChainMsg.JSONString call",
	})
  messagesCommitChainMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_JSONBuffer_Summary",
		Help: "Summary of messages.CommitChainMsg.JSONBuffer call",
	})
  messagesCommitChainMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_Sign_Summary",
		Help: "Summary of messages.CommitChainMsg.Sign call",
	})
  messagesCommitChainMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_GetSignature_Summary",
		Help: "Summary of messages.CommitChainMsg.GetSignature call",
	})
  messagesCommitChainMsgVerifySignature= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_VerifySignature_Summary",
		Help: "Summary of messages.CommitChainMsg.VerifySignature call",
	})
  messagesCommitChainMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.CommitChainMsg.UnmarshalBinaryData call",
	})
  messagesCommitChainMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.CommitChainMsg.UnmarshalBinary call",
	})
  messagesCommitChainMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_MarshalForSignature_Summary",
		Help: "Summary of messages.CommitChainMsg.MarshalForSignature call",
	})
  messagesCommitChainMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_MarshalBinary_Summary",
		Help: "Summary of messages.CommitChainMsg.MarshalBinary call",
	})
  messagesCommitChainMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitchainmsg_String_Summary",
		Help: "Summary of messages.CommitChainMsg.String call",
	})
//commit entry
  messagesCommitEntryMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_IsSameAs_Summary",
		Help: "Summary of messages.CommitEntryMsg.IsSameAs call",
	})
  messagesCommitEntryMsgGetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetCount_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetCount call",
	})
  messagesCommitEntryMsgSetCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_SetCount_Summary",
		Help: "Summary of messages.CommitEntryMsg.SetCount call",
	})
  messagesCommitEntryMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_Process_Summary",
		Help: "Summary of messages.CommitEntryMsg.Process call",
	})
  messagesCommitEntryMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetRepeatHash_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetRepeatHash call",
	})
  messagesCommitEntryMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetHash_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetHash call",
	})
  messagesCommitEntryMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetMsgHash_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetMsgHash call",
	})
  messagesCommitEntryMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetTimestamp_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetTimestamp call",
	})
  messagesCommitEntryMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_Type_Summary",
		Help: "Summary of messages.CommitEntryMsg.Typecall",
	})
  messagesCommitEntryMsgInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_Int_Summary",
		Help: "Summary of messages.CommitEntryMsg.Int call",
	})
  messagesCommitEntryMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_Sign_Summary",
		Help: "Summary of messages.CommitEntryMsg.Sign call",
	})
  messagesCommitEntryMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_GetSignature_Summary",
		Help: "Summary of messages.CommitEntryMsg.GetSignature call",
	})
  messagesCommitEntryMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_VerifySignature_Summary",
		Help: "Summary of messages.CommitEntryMsg.VerifySignature call",
	})
  messagesCommitEntryMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.CommitEntryMsg.UnmarshalBinaryData call",
	})
  messagesCommitEntryMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.CommitEntryMsg.UnmarshalBinary call",
	})
  messagesCommitEntryMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_MarshalForSignature_Summary",
		Help: "Summary of messages.CommitEntryMsg.MarshalForSignature call",
	})
  messagesCommitEntryMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_MarshalBinary_Summary",
		Help: "Summary of messages.CommitEntryMsg.MarshalBinary call",
	})
  messagesCommitEntryMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_String_Summary",
		Help: "Summary of messages.CommitEntryMsg.String call",
	})
  messagesCommitEntryMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_Validate_Summary",
		Help: "Summary of messages.CommitEntryMsg.Validate call",
	})
  messagesCommitEntryMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_ComputeVMIndex_Summary",
		Help: "Summary of messages.CommitEntryMsg.ComputeVMIndex call",
	})
  messagesCommitEntryMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_LeaderExecute_Summary",
		Help: "Summary of messages.CommitEntryMsg.LeaderExecute call",
	})
  messagesCommitEntryMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_FollowerExecute_Summary",
		Help: "Summary of messages.CommitEntryMsg.FollowerExecute call",
	})
  messagesCommitEntryMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_JSONByte_Summary",
		Help: "Summary of messages.CommitEntryMsg.JSONByte call",
	})
  messagesCommitEntryMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_JSONString_Summary",
		Help: "Summary of messages.CommitEntryMsg.JSONString call",
	})
  messagesCommitEntryMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_JSONBuffer_Summary",
		Help: "Summary of messages.CommitEntryMsg.JSONBuffer call",
	})
  messagesCommitEntryMsgNewCommitEntryMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_commitentrymsg_NewCommitEntryMsg_Summary",
		Help: "Summary of messages.CommitEntryMsg.NewCommitEntryMsg call",
	})

// dataResponse    
  messagesdataResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_IsSameAs_Summary",
		Help: "Summary of messages.dataResponse.IsSameAs call",
	})
  messagesdataResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_GetRepeatHash_Summary",
		Help: "Summary of messages.dataResponse.GetRepeatHash call",
	})
  messagesdataResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_GetHash_Summary",
		Help: "Summary of messages.dataResponse.GetHash call",
	})
  messagesdataResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_GetMsgHash_Summary",
		Help: "Summary of messages.dataResponse.GetMsgHash call",
	})
  messagesdataResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_Type_Summary",
		Help: "Summary of messages.dataResponse.Type call",
	})
  messagesdataResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_GetTimestamp_Summary",
		Help: "Summary of messages.dataResponse.GetTimestamp call",
	})
  messagesdataResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_Validate_Summary",
		Help: "Summary of messages.dataResponse.Validate call",
	})
  messagesdataResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_LeaderExecute_Summary",
		Help: "Summary of messages.dataResponse.LeaderExecute call",
	})
  messagesdataResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_FollowerExecute_Summary",
		Help: "Summary of messages.dataResponse.FollowerExecute call",
	})
  messagesdataResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_JSONByte_Summary",
		Help: "Summary of messages.dataResponse.JSONByte call",
	})
  messagesdataResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_JSONString_Summary",
		Help: "Summary of messages.dataResponse.JSONString call",
	})
  messagesdataResponseJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_JSONBuffer_Summary",
		Help: "Summary of messages.dataResponse.JSONBuffer call",
	})
  messagesdataResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.dataResponse.UnmarshalBinaryData call",
	})
  messagesdataResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_UnmarshalBinary_Summary",
		Help: "Summary of messages.dataResponse.UnmarshalBinary call",
	})
  messagesdataResponseattemptEntryUnmarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_attemptEntryUnmarshal_Summary",
		Help: "Summary of messages.dataResponse.attemptEntryUnmarshal call",
	})
  messagesdataResponseattemptEBlockUnmarshal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_attemptEBlockUnmarshal_Summary",
		Help: "Summary of messages.dataResponse.attemptEBlockUnmarshal call",
	})
  messagesdataResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_MarshalBinary_Summary",
		Help: "Summary of messages.dataResponse.MarshalBinary call",
	})
  messagesdataResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_String_Summary",
		Help: "Summary of messages.dataResponse.String call",
	})
  messagesdataResponseNewDataResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_dataresponse_NewDataResponse_Summary",
		Help: "Summary of messages.dataResponse.NewDataResponse call",
	})

    //dbstate.go
  messagesDBStateMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_IsSameAs_Summary",
		Help: "Summary of messages.DBStateMsg.IsSameAs call",
	})
  messagesDBStateMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetRepeatHash_Summary",
		Help: "Summary of messages.DBStateMsg.GetRepeatHash call",
	})
  messagesDBStateMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetHash_Summary",
		Help: "Summary of messages.DBStateMsg.GetHash call",
	})
  messagesDBStateMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetMsgHash_Summary",
		Help: "Summary of messages.DBStateMsg.GetMsgHash call",
	})
  messagesDBStateMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Type_Summary",
		Help: "Summary of messages.DBStateMsg.Type call",
	})
  messagesDBStateMsgInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Int_Summary",
		Help: "Summary of messages.DBStateMsg.Int call",
	})
  messagesDBStateMsgBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Bytes_Summary",
		Help: "Summary of messages.DBStateMsg.Bytes call",
	})
  messagesDBStateMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_GetTimestamp_Summary",
		Help: "Summary of messages.DBStateMsg.GetTimestamp call",
	})
  messagesDBStateMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Validate_Summary",
		Help: "Summary of messages.DBStateMsg.Validate call",
	})
  messagesDBStateMsgSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_SigTally_Summary",
		Help: "Summary of messages.DBStateMsg.SigTally call",
	})
  messagesDBStateMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_ComputeVMIndex_Summary",
		Help: "Summary of messages.DBStateMsg.ComputeVMIndex call",
	})
  messagesDBStateMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_LeaderExecute_Summary",
		Help: "Summary of messages.DBStateMsg.LeaderExecute call",
	})
  messagesDBStateMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_Process_Summary",
		Help: "Summary of messages.DBStateMsg.Process call",
	})
  messagesDBStateMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_JSONByte_Summary",
		Help: "Summary of messages.DBStateMsg.JSONByte call",
	})
  messagesDBStateMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_JSONString_Summary",
		Help: "Summary of messages.DBStateMsg.JSONString call",
	})
  messagesDBStateMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_JSONBuffer_Summary",
		Help: "Summary of messages.DBStateMsg. call",
	})
  messagesDBStateMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.DBStateMsg.UnmarshalBinaryData call",
	})
  messagesDBStateMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.DBStateMsg.UnmarshalBinary call",
	})
  messagesDBStateMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_MarshalBinary_Summary",
		Help: "Summary of messages.DBStateMsg.MarshalBinary call",
	})
  messagesDBStateMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_String_Summary",
		Help: "Summary of messages.DBStateMsg.String call",
	})
  messagesDBStateMsgNewDBStateMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMsg_NewDBStateMsg_Summary",
		Help: "Summary of messages.DBStateMsg.NewDBStateMsg call",
	})

    //dbstatemissing
  messagesDBStateMissingIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_IsSameAs_Summary",
		Help: "Summary of messages.DBStateMissing.IsSameAs call",
	})
  messagesDBStateMissingGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetRepeatHash_Summary",
		Help: "Summary of messages.DBStateMissing.GetRepeatHash call",
	})
  messagesDBStateMissingGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetHash_Summary",
		Help: "Summary of messages.DBStateMissing.GetHash call",
	})
  messagesDBStateMissingGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetMsgHash_Summary",
		Help: "Summary of messages.DBStateMissing.GetMsgHash call",
	})
  messagesDBStateMissingType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Type_Summary",
		Help: "Summary of messages.DBStateMissing.Type call",
	})
  messagesDBStateMissingInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Int_Summary",
		Help: "Summary of messages.DBStateMissing.Int call",
	})
  messagesDBStateMissingBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Bytes_Summary",
		Help: "Summary of messages.DBStateMissing.Bytes call",
	})
  messagesDBStateMissingGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_GetTimestamp_Summary",
		Help: "Summary of messages.DBStateMissing.GetTimestamp call",
	})
  messagesDBStateMissingValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_Validate_Summary",
		Help: "Summary of messages.DBStateMissing.Validate call",
	})
  messagesDBStateMissingLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_LeaderExecute_Summary",
		Help: "Summary of messages.DBStateMissing.LeaderExecute call",
	})
  messagesDBStateMissingsend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_send_Summary",
		Help: "Summary of messages.DBStateMissing.send call",
	})
  messagesDBStateMissingFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_FollowerExecute_Summary",
		Help: "Summary of messages.DBStateMissing.FollowerExecute call",
	})
  messagesDBStateMissingJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_JSONByte_Summary",
		Help: "Summary of messages.DBStateMissing.JSONByte call",
	})
  messagesDBStateMissingJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_JSONString_Summary",
		Help: "Summary of messages.DBStateMissing.JSONString call",
	})
  messagesDBStateMissingJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_JSONBuffer_Summary",
		Help: "Summary of messages.DBStateMissing.JSONBuffer call",
	})
  messagesDBStateMissingUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.DBStateMissing.UnmarshalBinaryData call",
	})
  messagesDBStateMissingUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_UnmarshalBinary_Summary",
		Help: "Summary of messages.DBStateMissing.UnmarshalBinary call",
	})
  messagesDBStateMissingMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_MarshalForSignature_Summary",
		Help: "Summary of messages.DBStateMissing.MarshalForSignature call",
	})
  messagesDBStateMissingMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_MarshalBinary_Summary",
		Help: "Summary of messages.DBStateMissing.MarshalBinary call",
	})
  messagesDBStateMissingString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DBStateMissing_String_Summary",
		Help: "Summary of messages.DBStateMissing.String call",
	})
  messagesDBStateMissingNewDBStateMissing = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_NewDBStateMissing_Summary",
		Help: "Summary of messages.DBStateMissing.NewDBStateMissing call",
	})

  //directoryBlockSignature  
  messagesDirectoryBlockSignatureIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_IsSameAs_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.IsSameAs call",
	})
  messagesDirectoryBlockSignatureProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Process_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.Process call",
	})
  messagesDirectoryBlockSignatureGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetRepeatHash_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.GetRepeatHash call",
	})
  messagesDirectoryBlockSignatureGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetHash_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.GetHash call",
	})
  messagesDirectoryBlockSignatureGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetTimestamp_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.GetTimestamp call",
	})
  messagesDirectoryBlockSignatureType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Type_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.Type call",
	})
  messagesDirectoryBlockSignatureValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Validate_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.Validate call",
	})
  messagesDirectoryBlockSignatureLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_LeaderExecute_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.LeaderExecute call",
	})
  messagesDirectoryBlockSignatureFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_FollowerExecute_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.FollowerExecute call",
	})
  messagesDirectoryBlockSignatureSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_Sign_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.Sign call",
	})
  messagesDirectoryBlockSignatureSignHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_SignHeader_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.SignHeader call",
	})
  messagesDirectoryBlockSignatureGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_GetSignature_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.GetSignature call",
	})
  messagesDirectoryBlockSignatureVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_VerifySignature_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.VerifySignature call",
	})
  messagesDirectoryBlockSignatureUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.UnmarshalBinaryData call",
	})
  messagesDirectoryBlockSignatureUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_UnmarshalBinary_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.UnmarshalBinary call",
	})
  messagesDirectoryBlockSignatureMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_MarshalForSignature_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.MarshalForSignature call",
	})
  messagesDirectoryBlockSignatureMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_MarshalBinary_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.MarshalBinary call",
	})
  messagesDirectoryBlockSignatureString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_String_Summary",
		Help: "Summary of messages.DirectoryBlockSignature. call",
	})
  messagesDirectoryBlockSignatureJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_JSONByte_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.JSONByte call",
	})
  messagesDirectoryBlockSignatureJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_JSONString_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.JSONString call",
	})
  messagesDirectoryBlockSignatureJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_DirectoryBlockSignature_JSONBuffer_Summary",
		Help: "Summary of messages.DirectoryBlockSignature.JSONBuffer call",
	})

    //entryBlockResponse
  messagesEntryBlockResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_IsSameAs_Summary",
		Help: "Summary of messages.EntryBlockResponse.IsSameAs call",
	})
  messagesEntryBlockResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetRepeatHash_Summary",
		Help: "Summary of messages.EntryBlockResponse.GetRepeatHash call",
	})
  messagesEntryBlockResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetHash_Summary",
		Help: "Summary of messages.EntryBlockResponse.GetHash call",
	})
  messagesEntryBlockResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetMsgHash_Summary",
		Help: "Summary of messages.EntryBlockResponse.GetMsgHash call",
	})
  messagesEntryBlockResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Type_Summary",
		Help: "Summary of messages.EntryBlockResponse.Type call",
	})
  messagesEntryBlockResponseInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Int_Summary",
		Help: "Summary of messages.EntryBlockResponse.Int call",
	})
  messagesEntryBlockResponseBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Bytes_Summary",
		Help: "Summary of messages.EntryBlockResponse.Bytes call",
	})
  messagesEntryBlockResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_GetTimestamp_Summary",
		Help: "Summary of messages.EntryBlockResponse.GetTimestamp call",
	})
  messagesEntryBlockResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_Validate_Summary",
		Help: "Summary of messages.EntryBlockResponse.Validate call",
	})
  messagesEntryBlockResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_LeaderExecute_Summary",
		Help: "Summary of messages.EntryBlockResponse.LeaderExecute call",
	})
  messagesEntryBlockResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_FollowerExecute_Summary",
		Help: "Summary of messages.EntryBlockResponse.FollowerExecute call",
	})
  messagesEntryBlockResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_JSONByte_Summary",
		Help: "Summary of messages.EntryBlockResponse.JSONByte call",
	})
  messagesEntryBlockResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_JSONString_Summary",
		Help: "Summary of messages.EntryBlockResponse.JSONString call",
	})
  messagesEntryBlockResponseJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_JSONBuffer_Summary",
		Help: "Summary of messages.EntryBlockResponse.JSONBuffer call",
	})
  messagesEntryBlockResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.EntryBlockResponse.UnmarshalBinaryData call",
	})
  messagesEntryBlockResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_UnmarshalBinary_Summary",
		Help: "Summary of messages.EntryBlockResponse.UnmarshalBinary call",
	})
  messagesEntryBlockResponseMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_MarshalForSignature_Summary",
		Help: "Summary of messages.EntryBlockResponse.MarshalForSignature call",
	})
  messagesEntryBlockResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_MarshalBinary_Summary",
		Help: "Summary of messages.EntryBlockResponse.MarshalBinary call",
	})
  messagesEntryBlockResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_String_Summary",
		Help: "Summary of messages.EntryBlockResponse.String call",
	})
  messagesEntryBlockResponseNewEntryBlockResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EntryBlockResponse_NewEntryBlockResponse_Summary",
		Help: "Summary of messages.EntryBlockResponse.NewEntryBlockResponse call",
	})

// 

  messagesEOMIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_IsSameAs_Summary",
		Help: "Summary of messages.EOM.IsSameAs call",
	})
  messagesEOMProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Process_Summary",
		Help: "Summary of messages.EOM.Process call",
	})
  messagesEOMGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetRepeatHash_Summary",
		Help: "Summary of messages.EOM.GetRepeatHash call",
	})
  messagesEOMGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetHash_Summary",
		Help: "Summary of messages.EOM.GetHash call",
	})
  messagesEOMGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetMsgHash_Summary",
		Help: "Summary of messages.EOM.GetMsgHash call",
	})
  messagesEOMGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetTimestamp_Summary",
		Help: "Summary of messages.EOM.GetTimestamp call",
	})
  messagesEOMInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Int_Summary",
		Help: "Summary of messages.EOM.Int call",
	})
  messagesEOMBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Bytes_Summary",
		Help: "Summary of messages.EOM.Bytes call",
	})
  messagesEOMType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Type_Summary",
		Help: "Summary of messages.EOM.Type call",
	})
  messagesEOMValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Validate_Summary",
		Help: "Summary of messages.EOM.Validate call",
	})
  messagesEOMLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_LeaderExecute_Summary",
		Help: "Summary of messages.EOM.LeaderExecute call",
	})
  messagesEOMFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_FollowerExecute_Summary",
		Help: "Summary of messages.EOM.FollowerExecute call",
	})
  messagesEOMJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_JSONByte_Summary",
		Help: "Summary of messages.EOM.JSONByte call",
	})
  messagesEOMJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_JSONString_Summary",
		Help: "Summary of messages.EOM.JSONString call",
	})
  messagesEOMJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_JSONBuffer_Summary",
		Help: "Summary of messages.EOM.JSONBuffer call",
	})
  messagesEOMSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_Sign_Summary",
		Help: "Summary of messages.EOM.Sign call",
	})
  messagesEOMGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_GetSignature_Summary",
		Help: "Summary of messages.EOM.GetSignature call",
	})
  messagesEOMVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_VerifySignature_Summary",
		Help: "Summary of messages.EOM.VerifySignature call",
	})
  messagesEOMUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.EOM.UnmarshalBinaryData call",
	})
  messagesEOMUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_UnmarshalBinary_Summary",
		Help: "Summary of messages.EOM.UnmarshalBinary call",
	})
  messagesEOMMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_MarshalForSignature_Summary",
		Help: "Summary of messages.EOM.MarshalForSignature call",
	})
  messagesEOMMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_MarshalBinary_Summary",
		Help: "Summary of messages.EOM.MarshalBinary call",
	})
  messagesEOMString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOM_String_Summary",
		Help: "Summary of messages.EOM.String call",
	})

// EOMTimeout
  messagesEOMTimeoutIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_IsSameAs_Summary",
		Help: "Summary of messages.EOMTimeout.IsSameAs call",
	})
 messagesEOMTimeoutSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Sign_Summary",
		Help: "Summary of messages.EOMTimeout.Sign call",
	})
 messagesEOMTimeoutGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetSignature_Summary",
		Help: "Summary of messages.EOMTimeout.GetSignature call",
	})
 messagesEOMTimeoutVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_VerifySignature_Summary",
		Help: "Summary of messages.EOMTimeout.VerifySignature call",
	})
 messagesEOMTimeoutGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetRepeatHash_Summary",
		Help: "Summary of messages.EOMTimeout.GetRepeatHash call",
	})
 messagesEOMTimeoutGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetHash_Summary",
		Help: "Summary of messages.EOMTimeout.GetHash call",
	})
 messagesEOMTimeoutGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetMsgHash_Summary",
		Help: "Summary of messages.EOMTimeout.GetMsgHash call",
	})
 messagesEOMTimeoutGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_GetTimestamp_Summary",
		Help: "Summary of messages.EOMTimeout.GetTimestamp call",
	})
 messagesEOMTimeoutType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Type_Summary",
		Help: "Summary of messages.EOMTimeout.Type call",
	})
 messagesEOMTimeoutInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Int_Summary",
		Help: "Summary of messages.EOMTimeout.Int call",
	})
 messagesEOMTimeoutBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_Bytes_Summary",
		Help: "Summary of messages.EOMTimeout.Bytes call",
	})
 messagesEOMTimeoutUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.EOMTimeout.UnmarshalBinaryData call",
	})
 messagesEOMTimeoutUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_UnmarshalBinary_Summary",
		Help: "Summary of messages.EOMTimeout.UnmarshalBinary call",
	})
 messagesEOMTimeoutMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_MarshalForSignature_Summary",
		Help: "Summary of messages.EOMTimeout.MarshalForSignature call",
	})
 messagesEOMTimeoutMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_MarshalBinary_Summary",
		Help: "Summary of messages.EOMTimeout.MarshalBinary call",
	})
 messagesEOMTimeoutJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_JSONByte_Summary",
		Help: "Summary of messages.EOMTimeout.JSONByte call",
	})
 messagesEOMTimeoutJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_JSONString_Summary",
		Help: "Summary of messages.EOMTimeout.JSONString call",
	})
 messagesEOMTimeoutJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_EOMTimeout_JSONBuffer_Summary",
		Help: "Summary of messages.EOMTimeout.JSONBuffer call",
	})

//  error.go   
 messagesMessageErrorError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_Error_Summary",
		Help: "Summary of messages.MessageError.Error call",
	})
messagesMessageErrormessageError = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_messageError_Summary",
		Help: "Summary of messages.MessageError.messageError call",
	})


//factoidTransaction
messagesFactoidTransactionIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_IsSameAs_Summary",
		Help: "Summary of messages.MessageError.IsSameAs call",
	})
messagesFactoidTransactionGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_GetRepeatHash_Summary",
		Help: "Summary of messages.MessageError.GetRepeatHash call",
	})
messagesFactoidTransactionGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_GetHash_Summary",
		Help: "Summary of messages.MessageError.GetHash call",
	})
messagesFactoidTransactionGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_GetMsgHash_Summary",
		Help: "Summary of messages.MessageError.GetMsgHash call",
	})
messagesFactoidTransactionGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_GetTimestamp_Summary",
		Help: "Summary of messages.MessageError.GetTimestamp call",
	})
messagesFactoidTransactionGetTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_GetTransaction_Summary",
		Help: "Summary of messages.MessageError.GetTransaction call",
	})
messagesFactoidTransactionSetTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_SetTransaction_Summary",
		Help: "Summary of messages.MessageError.SetTransaction call",
	})
messagesFactoidTransactionType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_Type_Summary",
		Help: "Summary of messages.MessageError.Type call",
	})
messagesFactoidTransactionValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_Validate_Summary",
		Help: "Summary of messages.MessageError.Validate call",
	})
messagesFactoidTransactionComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_ComputeVMIndex_Summary",
		Help: "Summary of messages.MessageError.ComputeVMIndex call",
	})
messagesFactoidTransactionLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_LeaderExecute_Summary",
		Help: "Summary of messages.MessageError.LeaderExecute call",
	})
messagesFactoidTransactionFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_FollowerExecute_Summary",
		Help: "Summary of messages.MessageError.FollowerExecute call",
	})
messagesFactoidTransactionProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_Process_Summary",
		Help: "Summary of messages.MessageError.Process call",
	})
messagesFactoidTransactionUnmarshalTransData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_UnmarshalTransData_Summary",
		Help: "Summary of messages.MessageError.UnmarshalTransData call",
	})
messagesFactoidTransactionUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MessageError.UnmarshalBinaryData call",
	})
messagesFactoidTransactionUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_UnmarshalBinary_Summary",
		Help: "Summary of messages.MessageError.UnmarshalBinary call",
	})
messagesFactoidTransactionString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_String_Summary",
		Help: "Summary of messages.MessageError.String call",
	})
messagesFactoidTransactionJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_JSONByte_Summary",
		Help: "Summary of messages.MessageError.JSONByte call",
	})
messagesFactoidTransactionJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_JSONString_Summary",
		Help: "Summary of messages.MessageError.JSONString call",
	})
messagesFactoidTransactionJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageError_JSONBuffer_Summary",
		Help: "Summary of messages.MessageError.JSONBuffer call",
	})

//
messagesFullServerFaultGetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetAmINegotiator_Summary",
		Help: "Summary of messages.FullServerFault.GetAmINegotiator call",
	})
messagesFullServerFaultSetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetAmINegotiator_Summary",
		Help: "Summary of messages.FullServerFault.SetAmINegotiator call",
	})
messagesFullServerFaultGetMyVoteTallied = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetMyVoteTallied_Summary",
		Help: "Summary of messages.FullServerFault.GetMyVoteTallied call",
	})
messagesFullServerFaultSetMyVoteTallied = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetMyVoteTallied_Summary",
		Help: "Summary of messages.FullServerFault.SetMyVoteTallied call",
	})
messagesFullServerFaultGetPledgeDone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetPledgeDone_Summary",
		Help: "Summary of messages.FullServerFault.GetPledgeDone call",
	})
messagesFullServerFaultSetPledgeDone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetPledgeDone_Summary",
		Help: "Summary of messages.FullServerFault.SetPledgeDone call",
	})
messagesFullServerFaultGetLastMatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetLastMatch_Summary",
		Help: "Summary of messages.FullServerFault.GetLastMatch call",
	})
messagesFullServerFaultSetLastMatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetLastMatch_Summary",
		Help: "Summary of messages.FullServerFault.SetLastMatch call",
	})
messagesFullServerFaultIsNil = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_IsNil_Summary",
		Help: "Summary of messages.FullServerFault.IsNil call",
	})
messagesFullServerFaultAddFaultVote = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_AddFaultVote_Summary",
		Help: "Summary of messages.FullServerFault.AddFaultVote call",
	})
messagesFullServerFaultPriority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Priority_Summary",
		Help: "Summary of messages.FullServerFault.Priority call",
	})
messagesFullServerFaultGetSerialHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetSerialHash_Summary",
		Help: "Summary of messages.FullServerFault.GetSerialHash call",
	})
messagesFullServerFaultProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Process_Summary",
		Help: "Summary of messages.FullServerFault.Process call",
	})
messagesFullServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetRepeatHash_Summary",
		Help: "Summary of messages.FullServerFault.GetRepeatHash call",
	})
messagesFullServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetHash_Summary",
		Help: "Summary of messages.FullServerFault.GetHash call",
	})
messagesFullServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetMsgHash_Summary",
		Help: "Summary of messages.FullServerFault.GetMsgHash call",
	})
messagesFullServerFaultGetCoreHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetCoreHash_Summary",
		Help: "Summary of messages.FullServerFault.GetCoreHash call",
	})
messagesFullServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetTimestamp_Summary",
		Help: "Summary of messages.FullServerFault.GetTimestamp call",
	})
messagesFullServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Type_Summary",
		Help: "Summary of messages.FullServerFault.Type call",
	})
messagesFullServerFaultMarshalCore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalCore_Summary",
		Help: "Summary of messages.FullServerFault.MarshalCore call",
	})
messagesFullServerFaultMarshalForSF = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalForSF_Summary",
		Help: "Summary of messages.FullServerFault.MarshalForSF call",
	})
messagesFullServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalForSignature_Summary",
		Help: "Summary of messages.FullServerFault.MarshalForSignature call",
	})
messagesFullServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalBinary_Summary",
		Help: "Summary of messages.FullServerFault.MarshalBinary call",
	})
messagesFullServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.FullServerFault.UnmarshalBinaryData call",
	})
messagesFullServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_MarshalBinary_Summary",
		Help: "Summary of messages.FullServerFault.MarshalBinary call",
	})
messagesFullServerFaultUnmarshall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Unmarshall_Summary",
		Help: "Summary of messages.FullServerFault.Unmarshall call",
	})
messagesFullServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.FullServerFault.UnmarshalBinaryData call",
	})
messagesFullServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_UnmarshalBinary_Summary",
		Help: "Summary of messages.FullServerFault.UnmarshalBinary call",
	})
messagesFullServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetSignature_Summary",
		Help: "Summary of messages.FullServerFault.GetSignature call",
	})
messagesFullServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_VerifySignature_Summary",
		Help: "Summary of messages.FullServerFault.VerifySignature call",
	})
messagesFullServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Sign_Summary",
		Help: "Summary of messages.FullServerFault.Sign call",
	})
messagesFullServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_String_Summary",
		Help: "Summary of messages.FullServerFault.String call",
	})
messagesFullServerFaultStringWithSigCnt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_StringWithSigCnt_Summary",
		Help: "Summary of messages.FullServerFault.StringWithSigCnt call",
	})
messagesFullServerFaultGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetDBHeight_Summary",
		Help: "Summary of messages.FullServerFault.GetDBHeight call",
	})
messagesFullServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_Validate_Summary",
		Help: "Summary of messages.FullServerFault.Validate call",
	})
messagesFullServerFaultSetAlreadyProcessed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SetAlreadyProcessed_Summary",
		Help: "Summary of messages.FullServerFault.SetAlreadyProcessed call",
	})
messagesFullServerFaultGetAlreadyProcessed = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_GetAlreadyProcessed_Summary",
		Help: "Summary of messages.FullServerFault.GetAlreadyProcessed call",
	})
messagesFullServerFaultHasEnoughSigs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_HasEnoughSigs_Summary",
		Help: "Summary of messages.FullServerFault.HasEnoughSigs call",
	})
messagesFullServerFaultSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_SigTally_Summary",
		Help: "Summary of messages.FullServerFault.SigTally call",
	})
messagesFullServerFaultLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_LeaderExecute_Summary",
		Help: "Summary of messages.FullServerFault.LeaderExecute call",
	})
messagesFullServerFaultFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_FollowerExecute_Summary",
		Help: "Summary of messages.FullServerFault.FollowerExecute call",
	})
messagesFullServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_JSONByte_Summary",
		Help: "Summary of messages.FullServerFault. call",
	})
messagesFullServerFaultJSONString= prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_JSONString_Summary",
		Help: "Summary of messages.FullServerFault.JSONString call",
	})
messagesFullServerFaultJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_JSONBuffer_Summary",
		Help: "Summary of messages.FullServerFault.JSONBuffer call",
	})
messagesFullServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_IsSameAs_Summary",
		Help: "Summary of messages.FullServerFault.IsSameAs call",
	})
messagesFullServerFaultToAdminBlockEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_ToAdminBlockEntry_Summary",
		Help: "Summary of messages.FullServerFault.ToAdminBlockEntry call",
	})
messagesFullServerFaultNewFullServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_FullServerFault_NewFullServerFault_Summary",
		Help: "Summary of messages.FullServerFault.NewFullServerFault call",
	})


// general.go

messagesGeneralUnmarshalMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_General_UnmarshalMessage_Summary",
		Help: "Summary of messages.General.UnmarshalMessage call",
	})
messagesGeneralUnmarshalMessageData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_General_UnmarshalMessageData_Summary",
		Help: "Summary of messages.General.UnmarshalMessageData call",
	})
messagesGeneralMessageName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_General_MessageName_Summary",
		Help: "Summary of messages.General.MessageName call",
	})
messagesGeneralSignSignable = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_General_SignSignable_Summary",
		Help: "Summary of messages.General.SignSignable call",
	})
messagesGeneralVerifyMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_General_VerifyMessage_Summary",
		Help: "Summary of messages.General.VerifyMessage call",
	})

//
messagesHeartbeatIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_IsSameAs_Summary",
		Help: "Summary of messages.Heartbeat.IsSameAs call",
	})
messagesHeartbeatGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetRepeatHash_Summary",
		Help: "Summary of messages.Heartbeat.GetRepeatHash call",
	})
messagesHeartbeatGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetHash_Summary",
		Help: "Summary of messages.Heartbeat.GetHash call",
	})
messagesHeartbeatGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetMsgHash_Summary",
		Help: "Summary of messages.Heartbeat.GetMsgHash call",
	})
messagesHeartbeatGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetTimestamp_Summary",
		Help: "Summary of messages.Heartbeat.GetTimestamp call",
	})
messagesHeartbeatType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Type_Summary",
		Help: "Summary of messages.Heartbeat.Type call",
	})
messagesHeartbeatInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Int_Summary",
		Help: "Summary of messages.Heartbeat.Int call",
	})
messagesHeartbeatBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Bytes_Summary",
		Help: "Summary of messages.Heartbeat.Bytes call",
	})
messagesHeartbeatUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.Heartbeat.UnmarshalBinaryData call",
	})
messagesHeartbeatUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_UnmarshalBinary_Summary",
		Help: "Summary of messages.Heartbeat.UnmarshalBinary call",
	})
messagesHeartbeatMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_MarshalForSignature_Summary",
		Help: "Summary of messages.Heartbeat.MarshalForSignature call",
	})
messagesHeartbeatMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_MarshalBinary_Summary",
		Help: "Summary of messages.Heartbeat.MarshalBinary call",
	})
messagesHeartbeatString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_String_Summary",
		Help: "Summary of messages.Heartbeat.String call",
	})
messagesHeartbeatValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Validate_Summary",
		Help: "Summary of messages.Heartbeat.Validate call",
	})
messagesHeartbeatLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_LeaderExecute_Summary",
		Help: "Summary of messages.Heartbeat.LeaderExecute call",
	})
messagesHeartbeatJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_JSONByte_Summary",
		Help: "Summary of messages.Heartbeat.JSONByte call",
	})
messagesHeartbeatJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_JSONString_Summary",
		Help: "Summary of messages.Heartbeat.JSONString call",
	})
messagesHeartbeatJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_JSONBuffer_Summary",
		Help: "Summary of messages.Heartbeat.JSONBuffer call",
	})
messagesHeartbeatSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_Sign_Summary",
		Help: "Summary of messages.Heartbeat.Sign call",
	})
messagesHeartbeatGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_Heartbeat_GetSignature_Summary",
		Help: "Summary of messages.Heartbeat.GetSignature call",
	})


// invqlid
messagesInvalidDirectoryBlockIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_IsSameAs_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.IsSameAs call",
	})
messagesInvalidDirectoryBlockSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Sign_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.Sign call",
	})
messagesInvalidDirectoryBlockGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetSignature_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.GetSignature call",
	})
messagesInvalidDirectoryBlockVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_VerifySignature_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.VerifySignature call",
	})
messagesInvalidDirectoryBlockProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Process_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.Process call",
	})
messagesInvalidDirectoryBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetHash_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.GetHash call",
	})
messagesInvalidDirectoryBlockGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetMsgHash_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.GetMsgHash call",
	})
messagesInvalidDirectoryBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_GetTimestamp_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.GetTimestamp call",
	})
messagesInvalidDirectoryBlockType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_Type_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.Type call",
	})
messagesInvalidDirectoryBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.UnmarshalBinaryData call",
	})
messagesInvalidDirectoryBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_UnmarshalBinary_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.UnmarshalBinary call",
	})
messagesInvalidDirectoryBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_MarshalBinary_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.MarshalBinary call",
	})
messagesInvalidDirectoryBlockMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_MarshalForSignature_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.MarshalForSignature call",
	})
messagesInvalidDirectoryBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_JSONByte_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.JSONByte call",
	})
messagesInvalidDirectoryBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_JSONString_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.JSONString call",
	})
messagesInvalidDirectoryBlockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_InvalidDirectoryBlock_JSONBuffer_Summary",
		Help: "Summary of messages.InvalidDirectoryBlock.JSONBuffer call",
	})
    

// messageBase.go
messagesMessageBaseresend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_resend_Summary",
		Help: "Summary of messages.MessageBase.resend call",
	})
messagesMessageBaseSendOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SendOut_Summary",
		Help: "Summary of messages.MessageBase.SendOut call",
	})
messagesMessageBaseGetNoResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetNoResend_Summary",
		Help: "Summary of messages.MessageBase.GetNoResend call",
	})
messagesMessageBaseSetNoResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetNoResend_Summary",
		Help: "Summary of messages.MessageBase.SetNoResend call",
	})
messagesMessageBaseIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsValid_Summary",
		Help: "Summary of messages.MessageBase.IsValid call",
	})
messagesMessageBaseSetValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetValid_Summary",
		Help: "Summary of messages.MessageBase.SetValid call",
	})
messagesMessageBaseMarkSentInvalid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_MarkSentInvalid_Summary",
		Help: "Summary of messages.MessageBase.MarkSentInvalid call",
	})
messagesMessageBaseSentInvlaid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SentInvlaid_Summary",
		Help: "Summary of messages.MessageBase.SentInvlaid call",
	})
messagesMessageBaseResend = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_Resend_Summary",
		Help: "Summary of messages.MessageBase.Resend call",
	})
messagesMessageBaseExpire = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_Expire_Summary",
		Help: "Summary of messages.MessageBase.Expire call",
	})
messagesMessageBaseIsStalled = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsStalled_Summary",
		Help: "Summary of messages.MessageBase.IsStalled call",
	})
messagesMessageBaseSetStall = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetStall_Summary",
		Help: "Summary of messages.MessageBase.SetStall call",
	})
messagesMessageBaseGetFullMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetFullMsgHash_Summary",
		Help: "Summary of messages.MessageBase.GetFullMsgHash call",
	})
messagesMessageBaseSetFullMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetFullMsgHash_Summary",
		Help: "Summary of messages.MessageBase.SetFullMsgHash call",
	})
messagesMessageBaseGetOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetOrigin_Summary",
		Help: "Summary of messages.MessageBase.GetOrigin call",
	})
messagesMessageBaseSetOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetOrigin_Summary",
		Help: "Summary of messages.MessageBase.SetOrigin call",
	})
messagesMessageBaseGetNetworkOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetNetworkOrigin_Summary",
		Help: "Summary of messages.MessageBase.GetNetworkOrigin call",
	})
messagesMessageBaseSetNetworkOrigin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetNetworkOrigin_Summary",
		Help: "Summary of messages.MessageBase.SetNetworkOrigin call",
	})
messagesMessageBaseIsPeer2Peer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsPeer2Peer_Summary",
		Help: "Summary of messages.MessageBase.IsPeer2Peer call",
	})
messagesMessageBaseSetPeer2Peer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetPeer2Peer_Summary",
		Help: "Summary of messages.MessageBase.SetPeer2Peer call",
	})
messagesMessageBaseIsLocal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_IsLocal_Summary",
		Help: "Summary of messages.MessageBase.IsLocal call",
	})
messagesMessageBaseSetLocal = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetLocal_Summary",
		Help: "Summary of messages.MessageBase.SetLocal call",
	})
messagesMessageBaseGetLeaderChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetLeaderChainID_Summary",
		Help: "Summary of messages.MessageBase.GetLeaderChainID call",
	})
messagesMessageBaseSetLeaderChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetLeaderChainID_Summary",
		Help: "Summary of messages.MessageBase.SetLeaderChainID call",
	})
messagesMessageBaseGetVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetVMIndex_Summary",
		Help: "Summary of messages.MessageBase.GetVMIndex call",
	})
messagesMessageBaseSetVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetVMIndex_Summary",
		Help: "Summary of messages.MessageBase.SetVMIndex call",
	})
messagesMessageBaseGetVMHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetVMHash_Summary",
		Help: "Summary of messages.MessageBase.GetVMHash call",
	})
messagesMessageBaseSetVMHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetVMHash_Summary",
		Help: "Summary of messages.MessageBase.SetVMHash call",
	})
messagesMessageBaseGetMinute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_GetMinute_Summary",
		Help: "Summary of messages.MessageBase.GetMinute call",
	})
messagesMessageBaseSetMinute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MessageBase_SetMinute_Summary",
		Help: "Summary of messages.MessageBase.SetMinute call",
	})


//  missingData
messagesMissingDataIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_IsSameAs_Summary",
		Help: "Summary of messages.MissingData.IsSameAs call",
	})
messagesMissingDataProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Process_Summary",
		Help: "Summary of messages.MissingData.Process call",
	})
messagesMissingDataGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetRepeatHash_Summary",
		Help: "Summary of messages.MissingData.GetRepeatHash call",
	})
messagesMissingDataGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetHash_Summary",
		Help: "Summary of messages.MissingData.GetHash call",
	})
messagesMissingDataGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetMsgHash_Summary",
		Help: "Summary of messages.MissingData.GetMsgHash call",
	})
messagesMissingDataGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_GetTimestamp_Summary",
		Help: "Summary of messages.MissingData.GetTimestamp call",
	})
messagesMissingDataType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Type_Summary",
		Help: "Summary of messages.MissingData.Type call",
	})
messagesMissingDataInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Int_Summary",
		Help: "Summary of messages.MissingData.Int call",
	})
messagesMissingDataBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_Bytes_Summary",
		Help: "Summary of messages.MissingData.Bytes call",
	})
messagesMissingDataUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MissingData.UnmarshalBinaryData call",
	})
messagesMissingDataUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_UnmarshalBinary_Summary",
		Help: "Summary of messages.MissingData.UnmarshalBinary call",
	})
messagesMissingDataMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_MarshalBinary_Summary",
		Help: "Summary of messages.MissingData.MarshalBinary call",
	})
messagesMissingDataString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_String_Summary",
		Help: "Summary of messages.MissingData.String call",
	})
messagesMissingDataLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_LeaderExecute_Summary",
		Help: "Summary of messages.MissingData.LeaderExecute call",
	})
messagesMissingDataFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_FollowerExecute_Summary",
		Help: "Summary of messages.MissingData.FollowerExecute call",
	})
messagesMissingDataJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_JSONByte_Summary",
		Help: "Summary of messages.MissingData.JSONByte call",
	})
messagesMissingDataJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_JSONString_Summary",
		Help: "Summary of messages.MissingData.JSONString call",
	})
messagesMissingDataJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_JSONBuffer_Summary",
		Help: "Summary of messages.MissingData.JSONBuffer call",
	})
messagesMissingDataNewMissingData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingData_NewMissingData_Summary",
		Help: "Summary of messages.MissingData.NewMissingData call",
	})


//
messagesMissingEntryBlocksIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_IsSameAs_Summary",
		Help: "Summary of messages.MissingEntryBlocks.IsSameAs call",
	})
messagesMissingEntryBlocksGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetRepeatHash_Summary",
		Help: "Summary of messages.MissingEntryBlocks.GetRepeatHash call",
	})
messagesMissingEntryBlocksGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetHash_Summary",
		Help: "Summary of messages.MissingEntryBlocks.GetHash call",
	})
messagesMissingEntryBlocksGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetMsgHash_Summary",
		Help: "Summary of messages.MissingEntryBlocks.GetMsgHash call",
	})
messagesMissingEntryBlocksType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_Type_Summary",
		Help: "Summary of messages.MissingEntryBlocks.Type call",
	})
messagesMissingEntryBlocksGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_GetTimestamp_Summary",
		Help: "Summary of messages.MissingEntryBlocks.GetTimestamp call",
	})
messagesMissingEntryBlocksValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_Validate_Summary",
		Help: "Summary of messages.MissingEntryBlocks.Validate call",
	})
messagesMissingEntryBlocksLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_LeaderExecute_Summary",
		Help: "Summary of messages.MissingEntryBlocks.LeaderExecute call",
	})
messagesMissingEntryBlocksFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_FollowerExecute_Summary",
		Help: "Summary of messages.MissingEntryBlocks.FollowerExecute call",
	})
messagesMissingEntryBlocksJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_JSONByte_Summary",
		Help: "Summary of messages.MissingEntryBlocks.JSONByte call",
	})
messagesMissingEntryBlocksJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_JSONString_Summary",
		Help: "Summary of messages.MissingEntryBlocks.JSONString call",
	})
messagesMissingEntryBlocksJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_JSONBuffer_Summary",
		Help: "Summary of messages.MissingEntryBlocks.JSONBuffer call",
	})
messagesMissingEntryBlocksUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MissingEntryBlocks.UnmarshalBinaryData call",
	})
messagesMissingEntryBlocksUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_UnmarshalBinary_Summary",
		Help: "Summary of messages.MissingEntryBlocks.UnmarshalBinary call",
	})
messagesMissingEntryBlocksMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_MarshalForSignature_Summary",
		Help: "Summary of messages.MissingEntryBlocks.MarshalForSignature call",
	})
messagesMissingEntryBlocksMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_MarshalBinary_Summary",
		Help: "Summary of messages.MissingEntryBlocks.MarshalBinary call",
	})
messagesMissingEntryBlocksString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_String_Summary",
		Help: "Summary of messages.MissingEntryBlocks.String call",
	})
messagesMissingEntryBlocksNewMissingEntryBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingEntryBlocks_NewMissingEntryBlocks_Summary",
		Help: "Summary of messages.MissingEntryBlocks.NewMissingEntryBlocks call",
	})
    
    //MissingMsg
messagesMissingMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_IsSameAs_Summary",
		Help: "Summary of messages.MissingMsg.IsSameAs call",
	})
messagesMissingMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetRepeatHash_Summary",
		Help: "Summary of messages.MissingMsg.GetRepeatHash call",
	})
messagesMissingMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetHash_Summary",
		Help: "Summary of messages.MissingMsg.GetHash call",
	})
messagesMissingMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetMsgHash_Summary",
		Help: "Summary of messages.MissingMsg.GetMsgHash call",
	})
messagesMissingMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_GetTimestamp_Summary",
		Help: "Summary of messages.MissingMsg.GetTimestamp call",
	})
messagesMissingMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_Type_Summary",
		Help: "Summary of messages.MissingMsg.Type call",
	})
messagesMissingMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MissingMsg.UnmarshalBinaryData call",
	})
messagesMissingMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.MissingMsg.UnmarshalBinary call",
	})
messagesMissingMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_MarshalBinary_Summary",
		Help: "Summary of messages.MissingMsg.MarshalBinary call",
	})
messagesMissingMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_String_Summary",
		Help: "Summary of messages.MissingMsg.String call",
	})
messagesMissingMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_Validate_Summary",
		Help: "Summary of messages.MissingMsg.Validate call",
	})
messagesMissingMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_LeaderExecute_Summary",
		Help: "Summary of messages.MissingMsg.LeaderExecute call",
	})
messagesMissingMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_FollowerExecute_Summary",
		Help: "Summary of messages.MissingMsg.FollowerExecute call",
	})
messagesMissingMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_JSONByte_Summary",
		Help: "Summary of messages.MissingMsg.JSONByte call",
	})
messagesMissingMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_JSONString_Summary",
		Help: "Summary of messages.MissingMsg.JSONString call",
	})
messagesMissingMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_JSONBuffer_Summary",
		Help: "Summary of messages.MissingMsg.JSONBuffer call",
	})
messagesMissingMsgAddHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_AddHeight_Summary",
		Help: "Summary of messages.MissingMsg.AddHeight call",
	})
messagesMissingMsgNewMissingMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsg_NewMissingMsg_Summary",
		Help: "Summary of messages.MissingMsg.NewMissingMsg call",
	})
    

    // MissingMsg
messagesMissingMsgResponseIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_IsSameAs_Summary",
		Help: "Summary of messages.MissingMsgResponse.IsSameAs call",
	})
messagesMissingMsgResponseGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetRepeatHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetRepeatHash call",
	})
messagesMissingMsgResponseGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetHash call",
	})
messagesMissingMsgResponseGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetMsgHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetMsgHashcall",
	})
messagesMissingMsgResponseGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetTimestamp_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetTimestamp call",
	})
messagesMissingMsgResponseType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Type_Summary",
		Help: "Summary of messages.MissingMsgResponse.Type call",
	})
messagesMissingMsgResponseInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Int_Summary",
		Help: "Summary of messages.MissingMsgResponse.Int call",
	})
messagesMissingMsgResponseBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Bytes_Summary",
		Help: "Summary of messages.MissingMsgResponse.Bytes call",
	})
messagesMissingMsgResponseUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MissingMsgResponse.UnmarshalBinaryData call",
	})
messagesMissingMsgResponseUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinary_Summary",
		Help: "Summary of messages.MissingMsgResponse.UnmarshalBinary call",
	})
messagesMissingMsgResponseMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_MarshalBinary_Summary",
		Help: "Summary of messages.MissingMsgResponse.MarshalBinary call",
	})
messagesMissingMsgResponseString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_String_Summary",
		Help: "Summary of messages.MissingMsgResponse.String call",
	})
messagesMissingMsgResponseValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Validate_Summary",
		Help: "Summary of messages.MissingMsgResponse.Validate call",
	})
messagesMissingMsgResponseLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_LeaderExecute_Summary",
		Help: "Summary of messages.MissingMsgResponse.LeaderExecute call",
	})
messagesMissingMsgResponseFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_FollowerExecute_Summary",
		Help: "Summary of messages.MissingMsgResponse.FollowerExecute call",
	})
messagesMissingMsgResponseJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONByte_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONByte call",
	})
messagesMissingMsgResponseJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONString_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONString call",
	})
messagesMissingMsgResponseJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONBuffer_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONBuffer call",
	})
messagesMissingMsgResponseNewMissingMsgResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_NewMissingMsgResponse_Summary",
		Help: "Summary of messages.MissingMsgResponse.NewMissingMsgResponse call",
	})

    //removeServer
messagesRemoveServerMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetRepeatHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetRepeatHash call",
	})
messagesRemoveServerMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetHash call",
	})
messagesRemoveServerMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetMsgHash_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetMsgHash call",
	})
messagesRemoveServerMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Type_Summary",
		Help: "Summary of messages.MissingMsgResponse.Type call",
	})
messagesRemoveServerMsgInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Int_Summary",
		Help: "Summary of messages.MissingMsgResponse.Int call",
	})
messagesRemoveServerMsgBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Bytes_Summary",
		Help: "Summary of messages.MissingMsgResponse.Bytes call",
	})
messagesRemoveServerMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetTimestamp_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetTimestamp call",
	})
messagesRemoveServerMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Validate_Summary",
		Help: "Summary of messages.MissingMsgResponse.Validate call",
	})
messagesRemoveServerMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_ComputeVMIndex_Summary",
		Help: "Summary of messages.MissingMsgResponse.ComputeVMIndex call",
	})
messagesRemoveServerMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_LeaderExecute_Summary",
		Help: "Summary of messages.MissingMsgResponse.LeaderExecute call",
	})
messagesRemoveServerMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_FollowerExecute_Summary",
		Help: "Summary of messages.MissingMsgResponse.FollowerExecute call",
	})
messagesRemoveServerMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Process_Summary",
		Help: "Summary of messages.MissingMsgResponse.Process call",
	})
messagesRemoveServerMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONByte_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONByte call",
	})
messagesRemoveServerMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONString_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONString call",
	})
messagesRemoveServerMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_JSONBuffer_Summary",
		Help: "Summary of messages.MissingMsgResponse.JSONBuffer call",
	})
messagesRemoveServerMsgSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_Sign_Summary",
		Help: "Summary of messages.MissingMsgResponse.Sign call",
	})
messagesRemoveServerMsgGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_GetSignature_Summary",
		Help: "Summary of messages.MissingMsgResponse.GetSignature call",
	})
messagesRemoveServerMsgVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_VerifySignature_Summary",
		Help: "Summary of messages.MissingMsgResponse.VerifySignature call",
	})
messagesRemoveServerMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.MissingMsgResponse.UnmarshalBinaryData call",
	})
messagesRemoveServerMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_UnmarshalBinary_Summary",
		Help: "Summary of messages.MissingMsgResponse.UnmarshalBinary call",
	})
messagesRemoveServerMsgMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_MarshalForSignature_Summary",
		Help: "Summary of messages.MissingMsgResponse.MarshalForSignature call",
	})
messagesRemoveServerMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_MarshalBinary_Summary",
		Help: "Summary of messages.MissingMsgResponse.MarshalBinary call",
	})
messagesRemoveServerMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_String_Summary",
		Help: "Summary of messages.MissingMsgResponse.String call",
	})
messagesRemoveServerMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_IsSameAs_Summary",
		Help: "Summary of messages.MissingMsgResponse.IsSameAs call",
	})
messagesRemoveServerMsgNewRemoveServerMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_MissingMsgResponse_NewRemoveServerMsg_Summary",
		Help: "Summary of messages.MissingMsgResponse.NewRemoveServerMsg call",
	})

    // RequestBlock
messagesRequestBlockIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_IsSameAs_Summary",
		Help: "Summary of messages.RequestBlock.IsSameAs call",
	})
messagesRequestBlockGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetRepeatHash_Summary",
		Help: "Summary of messages.RequestBlock.GetRepeatHash call",
	})
messagesRequestBlockGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetHash_Summary",
		Help: "Summary of messages.RequestBlock.GetHash call",
	})
messagesRequestBlockGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetMsgHash_Summary",
		Help: "Summary of messages.RequestBlock.GetMsgHash call",
	})
messagesRequestBlockGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_GetTimestamp_Summary",
		Help: "Summary of messages.RequestBlock.GetTimestamp call",
	})
messagesRequestBlockType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_Type_Summary",
		Help: "Summary of messages.RequestBlock.Type call",
	})
messagesRequestBlockUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.RequestBlock.UnmarshalBinaryData call",
	})
messagesRequestBlockUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_UnmarshalBinary_Summary",
		Help: "Summary of messages.RequestBlock.UnmarshalBinary call",
	})
messagesRequestBlockMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_MarshalForSignature_Summary",
		Help: "Summary of messages.RequestBlock.MarshalForSignature call",
	})
messagesRequestBlockMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_MarshalBinary_Summary",
		Help: "Summary of messages.RequestBlock.MarshalBinary call",
	})
messagesRequestBlockJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_JSONByte_Summary",
		Help: "Summary of messages.RequestBlock.JSONByte call",
	})
messagesRequestBlockJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_JSONString_Summary",
		Help: "Summary of messages.RequestBlock.JSONString call",
	})
messagesRequestBlockJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RequestBlock_JSONBuffer_Summary",
		Help: "Summary of messages.RequestBlock.JSONBuffer call",
	})
// revealEntry

messagesRevealEntryMsgIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_IsSameAs_Summary",
		Help: "Summary of messages.RevealEntryMsg.IsSameAs call",
	})
)
messagesRevealEntryMsgProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Process_Summary",
		Help: "Summary of messages.RevealEntryMsg.Process call",
	})
)
messagesRevealEntryMsgGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetRepeatHash_Summary",
		Help: "Summary of messages.RevealEntryMsg.GetRepeatHash call",
	})
)
messagesRevealEntryMsgGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetHash_Summary",
		Help: "Summary of messages.RevealEntryMsg.GetHash call",
	})
)
messagesRevealEntryMsgGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetMsgHash_Summary",
		Help: "Summary of messages.RevealEntryMsg.GetMsgHash call",
	})
)
messagesRevealEntryMsgGetChainIDHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetChainIDHash_Summary",
		Help: "Summary of messages.RevealEntryMsg.GetChainIDHash call",
	})
)
messagesRevealEntryMsgGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_GetTimestamp_Summary",
		Help: "Summary of messages.RevealEntryMsg. call",
	})
)
messagesRevealEntryMsgType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Type_Summary",
		Help: "Summary of messages.RevealEntryMsg.Type call",
	})
)
messagesRevealEntryMsgValidateRTN = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_ValidateRTN_Summary",
		Help: "Summary of messages.RevealEntryMsg.ValidateRTN call",
	})
)
messagesRevealEntryMsgValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_Validate_Summary",
		Help: "Summary of messages.RevealEntryMsg.Validate call",
	})
)
messagesRevealEntryMsgComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_ComputeVMIndex_Summary",
		Help: "Summary of messages.RevealEntryMsg.ComputeVMIndex call",
	})
)
messagesRevealEntryMsgLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_LeaderExecute_Summary",
		Help: "Summary of messages.RevealEntryMsg.LeaderExecute call",
	})
)
messagesRevealEntryMsgFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_FollowerExecute_Summary",
		Help: "Summary of messages.RevealEntryMsg.FollowerExecute call",
	})
)
messagesRevealEntryMsgJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_JSONByte_Summary",
		Help: "Summary of messages.RevealEntryMsg.JSONByte call",
	})
)
messagesRevealEntryMsgJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_JSONString_Summary",
		Help: "Summary of messages.RevealEntryMsg.JSONString call",
	})
)
messagesRevealEntryMsgJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_JSONBuffer_Summary",
		Help: "Summary of messages.RevealEntryMsg.JSONBuffer call",
	})
)
messagesRevealEntryMsgNewRevealEntryMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_NewRevealEntryMsg_Summary",
		Help: "Summary of messages.RevealEntryMsg.NewRevealEntryMsg call",
	})
)
messagesRevealEntryMsgUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.RevealEntryMsg.UnmarshalBinaryData call",
	})
)
messagesRevealEntryMsgUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_UnmarshalBinary_Summary",
		Help: "Summary of messages.RevealEntryMsg.UnmarshalBinary call",
	})
)
messagesRevealEntryMsgMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_MarshalBinary_Summary",
		Help: "Summary of messages.RevealEntryMsg.MarshalBinary call",
	})
)
messagesRevealEntryMsgString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_RevealEntryMsg_String_Summary",
		Help: "Summary of messages.RevealEntryMsg.String call",
	})
)


//serverFault
messagesServerFaultGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetRepeatHash_Summary",
		Help: "Summary of messages.ServerFault.GetRepeatHash call",
	})
)
messagesServerFaultGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetHash_Summary",
		Help: "Summary of messages.ServerFault.GetHash call",
	})
)
messagesServerFaultGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetMsgHash_Summary",
		Help: "Summary of messages.ServerFault.GetMsgHash call",
	})
)
messagesServerFaultGetCoreHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetCoreHash_Summary",
		Help: "Summary of messages.ServerFault.GetCoreHash call",
	})
)
messagesServerFaultGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetTimestamp_Summary",
		Help: "Summary of messages.ServerFault.GetTimestamp call",
	})
)
messagesServerFaultType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Type_Summary",
		Help: "Summary of messages.ServerFault.Type call",
	})
)
messagesServerFaultMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_MarshalForSignature_Summary",
		Help: "Summary of messages.ServerFault.MarshalForSignature call",
	})
)
messagesServerFaultPreMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_PreMarshalBinary_Summary",
		Help: "Summary of messages.ServerFault.PreMarshalBinary call",
	})
)
messagesServerFaultMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_MarshalBinary_Summary",
		Help: "Summary of messages.ServerFault.MarshalBinary call",
	})
)
messagesServerFaultUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.ServerFault.UnmarshalBinaryData call",
	})
)
messagesServerFaultUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_UnmarshalBinary_Summary",
		Help: "Summary of messages.ServerFault.UnmarshalBinary call",
	})
)
messagesServerFaultGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetSignature_Summary",
		Help: "Summary of messages.ServerFault.GetSignature call",
	})
)
messagesServerFaultVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_VerifySignature_Summary",
		Help: "Summary of messages.ServerFault.VerifySignature call",
	})
)
messagesServerFaultSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Sign_Summary",
		Help: "Summary of messages.ServerFault.Sign call",
	})
)
messagesServerFaultString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_String_Summary",
		Help: "Summary of messages.ServerFault.String call",
	})
)
messagesServerFaultGetDBHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_GetDBHeight_Summary",
		Help: "Summary of messages.ServerFault.GetDBHeight call",
	})
)
messagesServerFaultValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_Validate_Summary",
		Help: "Summary of messages.ServerFault.Validate call",
	})
)
messagesServerFaultLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_LeaderExecute_Summary",
		Help: "Summary of messages.ServerFault.LeaderExecute call",
	})
)
messagesServerFaultFollowerExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_FollowerExecute_Summary",
		Help: "Summary of messages.ServerFault.FollowerExecute call",
	})
)
messagesServerFaultJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_JSONByte_Summary",
		Help: "Summary of messages.ServerFault.JSONByte call",
	})
)
messagesServerFaultJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_JSONString_Summary",
		Help: "Summary of messages.ServerFault.JSONString call",
	})
)
messagesServerFaultJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_JSONBuffer_Summary",
		Help: "Summary of messages.ServerFault.JSONBuffer call",
	})
)
messagesServerFaultIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_IsSameAs_Summary",
		Help: "Summary of messages.ServerFault.IsSameAs call",
	})
)
messagesServerFaultNewServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_ServerFault_NewServerFault_Summary",
		Help: "Summary of messages.ServerFault.NewServerFault call",
	})
)

//SignatureTimeout

messagesSignatureTimeoutIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_IsSameAs_Summary",
		Help: "Summary of messages.SignatureTimeout.IsSameAs call",
	})
)
messagesSignatureTimeoutGetRepeatHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetRepeatHash_Summary",
		Help: "Summary of messages.SignatureTimeout.GetRepeatHash call",
	})
)
messagesSignatureTimeoutGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetHash_Summary",
		Help: "Summary of messages.SignatureTimeout.GetHash call",
	})
)
messagesSignatureTimeoutGetMsgHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetMsgHash_Summary",
		Help: "Summary of messages.SignatureTimeout.GetMsgHash call",
	})
)
messagesSignatureTimeoutGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetTimestamp_Summary",
		Help: "Summary of messages.SignatureTimeout.GetTimestamp call",
	})
)
messagesSignatureTimeoutType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Type_Summary",
		Help: "Summary of messages.SignatureTimeout.Type call",
	})
)
messagesSignatureTimeoutInt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Int_Summary",
		Help: "Summary of messages.SignatureTimeout.Int call",
	})
)
messagesSignatureTimeoutBytes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Bytes_Summary",
		Help: "Summary of messages.SignatureTimeout.Bytes call",
	})
)
messagesSignatureTimeoutUnmarshalBinaryData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_UnmarshalBinaryData_Summary",
		Help: "Summary of messages.SignatureTimeout.UnmarshalBinaryData call",
	})
)
messagesSignatureTimeoutUnmarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_UnmarshalBinary_Summary",
		Help: "Summary of messages.SignatureTimeout.UnmarshalBinary call",
	})
)
messagesSignatureTimeoutMarshalForSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_MarshalForSignature_Summary",
		Help: "Summary of messages.SignatureTimeout.MarshalForSignature call",
	})
)
messagesSignatureTimeoutMarshalBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_MarshalBinary_Summary",
		Help: "Summary of messages.SignatureTimeout.MarshalBinary call",
	})
)
messagesSignatureTimeoutGetSignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_GetSignature_Summary",
		Help: "Summary of messages.SignatureTimeout.GetSignature call",
	})
)
messagesSignatureTimeoutVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_VerifySignature_Summary",
		Help: "Summary of messages.SignatureTimeout.VerifySignature call",
	})
)
messagesSignatureTimeoutSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_Sign_Summary",
		Help: "Summary of messages.SignatureTimeout.Sign call",
	})
)
messagesSignatureTimeoutJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_JSONByte_Summary",
		Help: "Summary of messages.SignatureTimeout.JSONByte call",
	})
)
messagesSignatureTimeoutJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_JSONString_Summary",
		Help: "Summary of messages.SignatureTimeout.JSONString call",
	})
)
messagesSignatureTimeoutJSONBuffer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_messages_SignatureTimeout_JSONBuffer_Summary",
		Help: "Summary of messages.SignatureTimeout.JSONBuffer call",
	})
)



func InitPrometheus() {
	//anchorRecord.go variables
	prometheus.MustRegister(messageErrors)
	prometheus.MustRegister(messagesAckGetRepeatHash)
	prometheus.MustRegister(messagesAckGetHash)
	prometheus.MustRegister(messagesAckGetMsgHash)
	prometheus.MustRegister(messagesAckType)
	prometheus.MustRegister(messagesAckInt)
	prometheus.MustRegister(messagesAckBytes)
	prometheus.MustRegister(messagesAckGetTimestamp)
	prometheus.MustRegister(messagesAckVerifySignature)
	prometheus.MustRegister(messagesAckValidate)
	prometheus.MustRegister(messagesAckLeaderExecute)
	prometheus.MustRegister(messagesAckFollowerExecute)
	prometheus.MustRegister(messagesAckProcess)
	prometheus.MustRegister(messagesAckJSONByte)
	prometheus.MustRegister(messagesAckJSONString)
	prometheus.MustRegister(messagesAckJSONBuffer)
	prometheus.MustRegister(messagesAckSign)
	prometheus.MustRegister(messagesAckGetSignature)
	prometheus.MustRegister(messagesAckUnmarshalBinaryData)
	prometheus.MustRegister(messagesAckUnmarshalBinary)
	prometheus.MustRegister(messagesAckMarshalForSignature)
	prometheus.MustRegister(messagesAckMarshalBinary)
	prometheus.MustRegister(messagesAckString)
	prometheus.MustRegister(messagesAckIsSameAs)
	prometheus.MustRegister(messagesAddServerGetRepeatHash)
	prometheus.MustRegister(messagesAddServerGetHash)
	prometheus.MustRegister(messagesAddServerGetMsgHash)
	prometheus.MustRegister(messagesAddServerType)
	prometheus.MustRegister(messagesAddServerInt)
	prometheus.MustRegister(messagesAddServerBytes)
	prometheus.MustRegister(messagesAddServerGetTimestamp)
	prometheus.MustRegister(messagesAddServerValidate)
	prometheus.MustRegister(messagesAddServerComputeVMIndex)
	prometheus.MustRegister(messagesAddServerLeaderExecute)
	prometheus.MustRegister(messagesAddServerFollowerExecute)
	prometheus.MustRegister(messagesAddServerProcess)
	prometheus.MustRegister(messagesAddServerJSONByte)
	prometheus.MustRegister(messagesAddServerJSONString)
	prometheus.MustRegister(messagesAddServerJSONBuffer)
	prometheus.MustRegister(messagesAddServerSign)
	prometheus.MustRegister(messagesAddServerGetSignature)
	prometheus.MustRegister(messagesAddServerVerifySignature)
	prometheus.MustRegister(messagesAddServerUnmarshalBinaryData)
	prometheus.MustRegister(messagesAddServerUnmarshalBinary)
	prometheus.MustRegister(messagesAddServerMarshalForSignature)
	prometheus.MustRegister(messagesAddServerMarshalBinary)
	prometheus.MustRegister(messagesAddServerString)
	prometheus.MustRegister(messagesAddServerIsSameAs)
	prometheus.MustRegister(messagesAddServerNewAddServerMsg)
	prometheus.MustRegister(messagesAddServerNewAddServerByHashMsg)
	prometheus.MustRegister(messagesAuditServerFaultGetRepeatHash)
	prometheus.MustRegister(messagesAuditServerFaultIsSameAs)
	prometheus.MustRegister(messagesAuditServerFaultSign)
	prometheus.MustRegister(messagesAuditServerFaultGetSignature)
	prometheus.MustRegister(messagesAuditServerFaultVerifySignature)
	prometheus.MustRegister(messagesAuditServerFaultProcess)
	prometheus.MustRegister(messagesAuditServerFaultGetTimestamp)
	prometheus.MustRegister(messagesAuditServerFaultGetHash)
	prometheus.MustRegister(messagesAuditServerFaultGetMsgHash)
	prometheus.MustRegister(messagesAuditServerFaultType)
	prometheus.MustRegister(messagesAuditServerFaultInt)
	prometheus.MustRegister(messagesAuditServerFaultBytes)
	prometheus.MustRegister(messagesAuditServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(messagesAuditServerFaultUnmarshalBinary)
	prometheus.MustRegister(messagesAuditServerFaultMarshalForSignature)
	prometheus.MustRegister(messagesAuditServerFaultMarshalBinary)
	prometheus.MustRegister(messagesAuditServerFaultString)
	prometheus.MustRegister(messagesAuditServerFaultDBHeight)
	prometheus.MustRegister(messagesAuditServerFaultChainID)
	prometheus.MustRegister(messagesAuditServerFaultListHeight)
	prometheus.MustRegister(messagesAuditServerFaultSerialHash)
	prometheus.MustRegister(messagesAuditServerFaultValidate)
	prometheus.MustRegister(messagesAuditServerFaultJSONByte)
	prometheus.MustRegister(messagesAuditServerFaultJSONString)
	prometheus.MustRegister(messagesAuditServerFaultJSONBuffer)
	prometheus.MustRegister(messagesBounceAddData)
	prometheus.MustRegister(messagesBounceGetRepeatHash)
	prometheus.MustRegister(messagesBounceGetHash)
	prometheus.MustRegister(messagesBounceSizeOf)
	prometheus.MustRegister(messagesBounceGetMsgHash)
	prometheus.MustRegister(messagesBounceType)
	prometheus.MustRegister(messagesBounceGetTimestamp)
	prometheus.MustRegister(messagesBounceVerifySignature)
	prometheus.MustRegister(messagesBounceJSONByte)
	prometheus.MustRegister(messagesBounceJSONString)
	prometheus.MustRegister(messagesBounceJSONBuffer)
	prometheus.MustRegister(messagesBounceUnmarshalBinaryData)
	prometheus.MustRegister(messagesBounceUnmarshalBinary)
	prometheus.MustRegister(messagesBounceMarshalForSignature)
	prometheus.MustRegister(messagesBounceMarshalBinary)
	prometheus.MustRegister(messagesBounceString)
	prometheus.MustRegister(messagesBounceReplyGetRepeatHash)
	prometheus.MustRegister(messagesBounceReplyGetHash)
	prometheus.MustRegister(messagesBounceReplySizeOf)
	prometheus.MustRegister(messagesBounceReplyGetMsgHash)
	prometheus.MustRegister(messagesBounceReplyType)
	prometheus.MustRegister(messagesBounceReplyGetTimestamp)
	prometheus.MustRegister(messagesBounceReplyJSONByte)
	prometheus.MustRegister(messagesBounceReplyJSONString)
	prometheus.MustRegister(messagesBounceReplyJSONBuffer)
	prometheus.MustRegister(messagesBounceReplyUnmarshalBinaryData)
	prometheus.MustRegister(messagesBounceReplyUnmarshalBinary)
	prometheus.MustRegister(messagesBounceReplyMarshalForSignature)
	prometheus.MustRegister(messagesBounceReplyMarshalBinary)
	prometheus.MustRegister(messagesBounceReplyString)
	prometheus.MustRegister(messagesChangeServerKeyGetRepeatHash)
	prometheus.MustRegister(messagesChangeServerKeyGetHash)
	prometheus.MustRegister(messagesChangeServerKeyGetMsgHash)
	prometheus.MustRegister(messagesChangeServerKeyType)
	prometheus.MustRegister(messagesChangeServerKeyGetTimestamp)
	prometheus.MustRegister(messagesChangeServerKeyValidate)
	prometheus.MustRegister(messagesChangeServerKeyComputeVMIndex)
	prometheus.MustRegister(messagesChangeServerKeyLeaderExecute)
	prometheus.MustRegister(messagesChangeServerKeyFollowerExecute)
	prometheus.MustRegister(messagesChangeServerKeyProcess)
	prometheus.MustRegister(messagesChangeServerKeyJSONString)
	prometheus.MustRegister(messagesChangeServerKeyJSONBuffer)
	prometheus.MustRegister(messagesChangeServerKeyJSONByte)
	prometheus.MustRegister(messagesChangeServerKeySign)
	prometheus.MustRegister(messagesChangeServerKeyVerifySignature)
	prometheus.MustRegister(messagesChangeServerKeyUnmarshalBinaryData)
	prometheus.MustRegister(messagesChangeServerKeyUnmarshalBinary)
	prometheus.MustRegister(messagesChangeServerKeyMarshalForSignature)
	prometheus.MustRegister(messagesChangeServerKeyMarshalBinary)
	prometheus.MustRegister(messagesChangeServerKeyString)
	prometheus.MustRegister(messagesChangeServerKeyIsSameAs)
	prometheus.MustRegister(messagesChangeServerKeyNewChangeServerKeyMsg)
	prometheus.MustRegister(messagesCommitChainMsgIsSameAs)
	prometheus.MustRegister(messagesCommitChainMsgGetCount)
	prometheus.MustRegister(messagesCommitChainMsgIncCount)
	prometheus.MustRegister(messagesCommitChainMsgSetCount)
	prometheus.MustRegister(messagesCommitChainMsgProcess)
	prometheus.MustRegister(messagesCommitChainMsgGetRepeatHash)
	prometheus.MustRegister(messagesCommitChainMsgGetHash)
	prometheus.MustRegister(messagesCommitChainMsgGetMsgHash)
	prometheus.MustRegister(messagesCommitChainMsgGetTimestamp)
	prometheus.MustRegister(messagesCommitChainMsgType)
	prometheus.MustRegister(messagesCommitChainMsgValidate)
	prometheus.MustRegister(messagesCommitChainMsgComputeVMIndex)
	prometheus.MustRegister(messagesCommitChainMsgLeaderExecute)
	prometheus.MustRegister(messagesCommitChainMsgFollowerExecute)
	prometheus.MustRegister(messagesCommitChainMsgJSONByte)
	prometheus.MustRegister(messagesCommitChainMsgJSONString)
	prometheus.MustRegister(messagesCommitChainMsgJSONBuffer)
	prometheus.MustRegister(messagesCommitChainMsgSign)
	prometheus.MustRegister(messagesCommitChainMsgGetSignature)
	prometheus.MustRegister(messagesCommitChainMsgVerifySignature)
	prometheus.MustRegister(messagesCommitChainMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesCommitChainMsgUnmarshalBinary)
	prometheus.MustRegister(messagesCommitChainMsgMarshalForSignature)
	prometheus.MustRegister(messagesCommitChainMsgMarshalBinary)
	prometheus.MustRegister(messagesCommitChainMsgString)
	prometheus.MustRegister(messagesCommitEntryMsgIsSameAs)
	prometheus.MustRegister(messagesCommitEntryMsgGetCount)
	prometheus.MustRegister(messagesCommitEntryMsgSetCount)
	prometheus.MustRegister(messagesCommitEntryMsgProcess)
	prometheus.MustRegister(messagesCommitEntryMsgGetRepeatHash)
	prometheus.MustRegister(messagesCommitEntryMsgGetHash)
	prometheus.MustRegister(messagesCommitEntryMsgGetMsgHash)
	prometheus.MustRegister(messagesCommitEntryMsgGetTimestamp)
	prometheus.MustRegister(messagesCommitEntryMsgType)
	prometheus.MustRegister(messagesCommitEntryMsgInt)
	prometheus.MustRegister(messagesCommitEntryMsgSign)
	prometheus.MustRegister(messagesCommitEntryMsgGetSignature)
	prometheus.MustRegister(messagesCommitEntryMsgVerifySignature)
	prometheus.MustRegister(messagesCommitEntryMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesCommitEntryMsgUnmarshalBinary)
	prometheus.MustRegister(messagesCommitEntryMsgMarshalForSignature)
	prometheus.MustRegister(messagesCommitEntryMsgMarshalBinary)
	prometheus.MustRegister(messagesCommitEntryMsgString)
	prometheus.MustRegister(messagesCommitEntryMsgValidate)
	prometheus.MustRegister(messagesCommitEntryMsgComputeVMIndex)
	prometheus.MustRegister(messagesCommitEntryMsgLeaderExecute)
	prometheus.MustRegister(messagesCommitEntryMsgFollowerExecute)
	prometheus.MustRegister(messagesCommitEntryMsgJSONByte)
	prometheus.MustRegister(messagesCommitEntryMsgJSONString)
	prometheus.MustRegister(messagesCommitEntryMsgJSONBuffer)
	prometheus.MustRegister(messagesCommitEntryMsgNewCommitEntryMsg)
	prometheus.MustRegister(messagesdataResponseIsSameAs)
	prometheus.MustRegister(messagesdataResponseGetRepeatHash)
	prometheus.MustRegister(messagesdataResponseGetHash)
	prometheus.MustRegister(messagesdataResponseGetMsgHash)
	prometheus.MustRegister(messagesdataResponseType)
	prometheus.MustRegister(messagesdataResponseGetTimestamp)
	prometheus.MustRegister(messagesdataResponseValidate)
	prometheus.MustRegister(messagesdataResponseLeaderExecute)
	prometheus.MustRegister(messagesdataResponseFollowerExecute)
	prometheus.MustRegister(messagesdataResponseJSONByte)
	prometheus.MustRegister(messagesdataResponseJSONString)
	prometheus.MustRegister(messagesdataResponseJSONBuffer)
	prometheus.MustRegister(messagesdataResponseUnmarshalBinaryData)
	prometheus.MustRegister(messagesdataResponseUnmarshalBinary)
	prometheus.MustRegister(messagesdataResponseattemptEntryUnmarshal)
	prometheus.MustRegister(messagesdataResponseattemptEBlockUnmarshal)
	prometheus.MustRegister(messagesdataResponseMarshalBinary)
	prometheus.MustRegister(messagesdataResponseString)
	prometheus.MustRegister(messagesdataResponseNewDataResponse)
	prometheus.MustRegister(messagesDBStateMsgIsSameAs)
	prometheus.MustRegister(messagesDBStateMsgGetRepeatHash)
	prometheus.MustRegister(messagesDBStateMsgGetHash)
	prometheus.MustRegister(messagesDBStateMsgGetMsgHash)
	prometheus.MustRegister(messagesDBStateMsgType)
	prometheus.MustRegister(messagesDBStateMsgInt)
	prometheus.MustRegister(messagesDBStateMsgBytes)
	prometheus.MustRegister(messagesDBStateMsgGetTimestamp)
	prometheus.MustRegister(messagesDBStateMsgValidate)
	prometheus.MustRegister(messagesDBStateMsgSigTally)
	prometheus.MustRegister(messagesDBStateMsgComputeVMIndex)
	prometheus.MustRegister(messagesDBStateMsgLeaderExecute)
	prometheus.MustRegister(messagesDBStateMsgProcess)
	prometheus.MustRegister(messagesDBStateMsgJSONByte)
	prometheus.MustRegister(messagesDBStateMsgJSONString)
	prometheus.MustRegister(messagesDBStateMsgJSONBuffer)
	prometheus.MustRegister(messagesDBStateMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesDBStateMsgUnmarshalBinary)
	prometheus.MustRegister(messagesDBStateMsgMarshalBinary)
	prometheus.MustRegister(messagesDBStateMsgString)
	prometheus.MustRegister(messagesDBStateMsgNewDBStateMsg)
	prometheus.MustRegister(messagesDBStateMissingIsSameAs)
	prometheus.MustRegister(messagesDBStateMissingGetRepeatHash)
	prometheus.MustRegister(messagesDBStateMissingGetHash)
	prometheus.MustRegister(messagesDBStateMissingGetMsgHash)
	prometheus.MustRegister(messagesDBStateMissingType)
	prometheus.MustRegister(messagesDBStateMissingInt)
	prometheus.MustRegister(messagesDBStateMissingBytes)
	prometheus.MustRegister(messagesDBStateMissingGetTimestamp)
	prometheus.MustRegister(messagesDBStateMissingValidate)
	prometheus.MustRegister(messagesDBStateMissingLeaderExecute)
	prometheus.MustRegister(messagesDBStateMissingsend)
	prometheus.MustRegister(messagesDBStateMissingFollowerExecute)
	prometheus.MustRegister(messagesDBStateMissingJSONByte)
	prometheus.MustRegister(messagesDBStateMissingJSONString)
	prometheus.MustRegister(messagesDBStateMissingJSONBuffer)
	prometheus.MustRegister(messagesDBStateMissingUnmarshalBinaryData)
	prometheus.MustRegister(messagesDBStateMissingUnmarshalBinary)
	prometheus.MustRegister(messagesDBStateMissingMarshalForSignature)
	prometheus.MustRegister(messagesDBStateMissingMarshalBinary)
	prometheus.MustRegister(messagesDBStateMissingString)
	prometheus.MustRegister(messagesDBStateMissingNewDBStateMissing)
	prometheus.MustRegister(messagesDirectoryBlockSignatureIsSameAs)
	prometheus.MustRegister(messagesDirectoryBlockSignatureProcess)
	prometheus.MustRegister(messagesDirectoryBlockSignatureGetRepeatHash)
	prometheus.MustRegister(messagesDirectoryBlockSignatureGetHash)
	prometheus.MustRegister(messagesDirectoryBlockSignatureGetTimestamp)
	prometheus.MustRegister(messagesDirectoryBlockSignatureType)
	prometheus.MustRegister(messagesDirectoryBlockSignatureValidate)
	prometheus.MustRegister(messagesDirectoryBlockSignatureLeaderExecute)
	prometheus.MustRegister(messagesDirectoryBlockSignatureFollowerExecute)
	prometheus.MustRegister(messagesDirectoryBlockSignatureSign)
	prometheus.MustRegister(messagesDirectoryBlockSignatureSignHeader)
	prometheus.MustRegister(messagesDirectoryBlockSignatureGetSignature)
	prometheus.MustRegister(messagesDirectoryBlockSignatureVerifySignature)
	prometheus.MustRegister(messagesDirectoryBlockSignatureUnmarshalBinaryData)
	prometheus.MustRegister(messagesDirectoryBlockSignatureUnmarshalBinary)
	prometheus.MustRegister(messagesDirectoryBlockSignatureMarshalForSignature)
	prometheus.MustRegister(messagesDirectoryBlockSignatureMarshalBinary)
	prometheus.MustRegister(messagesDirectoryBlockSignatureString)
	prometheus.MustRegister(messagesDirectoryBlockSignatureJSONByte)
	prometheus.MustRegister(messagesDirectoryBlockSignatureJSONString)
	prometheus.MustRegister(messagesDirectoryBlockSignatureJSONBuffer)
	prometheus.MustRegister(messagesEntryBlockResponseIsSameAs)
	prometheus.MustRegister(messagesEntryBlockResponseGetRepeatHash)
	prometheus.MustRegister(messagesEntryBlockResponseGetHash)
	prometheus.MustRegister(messagesEntryBlockResponseGetMsgHash)
	prometheus.MustRegister(messagesEntryBlockResponseType)
	prometheus.MustRegister(messagesEntryBlockResponseInt)
	prometheus.MustRegister(messagesEntryBlockResponseBytes)
	prometheus.MustRegister(messagesEntryBlockResponseGetTimestamp)
	prometheus.MustRegister(messagesEntryBlockResponseValidate)
	prometheus.MustRegister(messagesEntryBlockResponseLeaderExecute)
	prometheus.MustRegister(messagesEntryBlockResponseFollowerExecute)
	prometheus.MustRegister(messagesEntryBlockResponseJSONByte)
	prometheus.MustRegister(messagesEntryBlockResponseJSONString)
	prometheus.MustRegister(messagesEntryBlockResponseJSONBuffer)
	prometheus.MustRegister(messagesEntryBlockResponseUnmarshalBinaryData)
	prometheus.MustRegister(messagesEntryBlockResponseUnmarshalBinary)
	prometheus.MustRegister(messagesEntryBlockResponseMarshalForSignature)
	prometheus.MustRegister(messagesEntryBlockResponseMarshalBinary)
	prometheus.MustRegister(messagesEntryBlockResponseString)
	prometheus.MustRegister(messagesEntryBlockResponseNewEntryBlockResponse)
	prometheus.MustRegister(messagesEOMIsSameAs)
	prometheus.MustRegister(messagesEOMProcess)
	prometheus.MustRegister(messagesEOMGetRepeatHash)
	prometheus.MustRegister(messagesEOMGetHash)
	prometheus.MustRegister(messagesEOMGetMsgHash)
	prometheus.MustRegister(messagesEOMGetTimestamp)
	prometheus.MustRegister(messagesEOMInt)
	prometheus.MustRegister(messagesEOMBytes)
	prometheus.MustRegister(messagesEOMType)
	prometheus.MustRegister(messagesEOMValidate)
	prometheus.MustRegister(messagesEOMLeaderExecute)
	prometheus.MustRegister(messagesEOMFollowerExecute)
	prometheus.MustRegister(messagesEOMJSONByte)
	prometheus.MustRegister(messagesEOMJSONString)
	prometheus.MustRegister(messagesEOMJSONBuffer)
	prometheus.MustRegister(messagesEOMSign)
	prometheus.MustRegister(messagesEOMGetSignature)
	prometheus.MustRegister(messagesEOMVerifySignature)
	prometheus.MustRegister(messagesEOMUnmarshalBinaryData)
	prometheus.MustRegister(messagesEOMUnmarshalBinary)
	prometheus.MustRegister(messagesEOMMarshalForSignature)
	prometheus.MustRegister(messagesEOMMarshalBinary)
	prometheus.MustRegister(messagesEOMString)
	prometheus.MustRegister(messagesEOMTimeoutIsSameAs)
	prometheus.MustRegister(messagesEOMTimeoutSign)
	prometheus.MustRegister(messagesEOMTimeoutGetSignature)
	prometheus.MustRegister(messagesEOMTimeoutVerifySignature)
	prometheus.MustRegister(messagesEOMTimeoutGetRepeatHash)
	prometheus.MustRegister(messagesEOMTimeoutGetHash)
	prometheus.MustRegister(messagesEOMTimeoutGetMsgHash)
	prometheus.MustRegister(messagesEOMTimeoutGetTimestamp)
	prometheus.MustRegister(messagesEOMTimeoutType)
	prometheus.MustRegister(messagesEOMTimeoutInt)
	prometheus.MustRegister(messagesEOMTimeoutBytes)
	prometheus.MustRegister(messagesEOMTimeoutUnmarshalBinaryData)
	prometheus.MustRegister(messagesEOMTimeoutUnmarshalBinary)
	prometheus.MustRegister(messagesEOMTimeoutMarshalForSignature)
	prometheus.MustRegister(messagesEOMTimeoutMarshalBinary)
	prometheus.MustRegister(messagesEOMTimeoutJSONByte)
	prometheus.MustRegister(messagesEOMTimeoutJSONString)
	prometheus.MustRegister(messagesEOMTimeoutJSONBuffer)
	prometheus.MustRegister(messagesMessageErrorError)
	prometheus.MustRegister(messagesMessageErrormessageError)
	prometheus.MustRegister(messagesFactoidTransactionIsSameAs)
	prometheus.MustRegister(messagesFactoidTransactionGetRepeatHash)
	prometheus.MustRegister(messagesFactoidTransactionGetHash)
	prometheus.MustRegister(messagesFactoidTransactionGetMsgHash)
	prometheus.MustRegister(messagesFactoidTransactionGetTimestamp)
	prometheus.MustRegister(messagesFactoidTransactionGetTransaction)
	prometheus.MustRegister(messagesFactoidTransactionSetTransaction)
	prometheus.MustRegister(messagesFactoidTransactionType)
	prometheus.MustRegister(messagesFactoidTransactionValidate)
	prometheus.MustRegister(messagesFactoidTransactionComputeVMIndex)
	prometheus.MustRegister(messagesFactoidTransactionLeaderExecute)
	prometheus.MustRegister(messagesFactoidTransactionFollowerExecute)
	prometheus.MustRegister(messagesFactoidTransactionProcess)
	prometheus.MustRegister(messagesFactoidTransactionUnmarshalTransData)
	prometheus.MustRegister(messagesFactoidTransactionUnmarshalBinaryData)
	prometheus.MustRegister(messagesFactoidTransactionUnmarshalBinary)
	prometheus.MustRegister(messagesFactoidTransactionString)
	prometheus.MustRegister(messagesFactoidTransactionJSONByte)
	prometheus.MustRegister(messagesFactoidTransactionJSONString)
	prometheus.MustRegister(messagesFactoidTransactionJSONBuffer)
	prometheus.MustRegister(messagesFullServerFaultGetAmINegotiator)
	prometheus.MustRegister(messagesFullServerFaultSetAmINegotiator)
	prometheus.MustRegister(messagesFullServerFaultGetMyVoteTallied)
	prometheus.MustRegister(messagesFullServerFaultSetMyVoteTallied)
	prometheus.MustRegister(messagesFullServerFaultGetPledgeDone)
	prometheus.MustRegister(messagesFullServerFaultSetPledgeDone)
	prometheus.MustRegister(messagesFullServerFaultGetLastMatch)
	prometheus.MustRegister(messagesFullServerFaultSetLastMatch)
	prometheus.MustRegister(messagesFullServerFaultIsNil)
	prometheus.MustRegister(messagesFullServerFaultAddFaultVote)
	prometheus.MustRegister(messagesFullServerFaultPriority)
	prometheus.MustRegister(messagesFullServerFaultGetSerialHash)
	prometheus.MustRegister(messagesFullServerFaultProcess)
	prometheus.MustRegister(messagesFullServerFaultGetRepeatHash)
	prometheus.MustRegister(messagesFullServerFaultGetHash)
	prometheus.MustRegister(messagesFullServerFaultGetMsgHash)
	prometheus.MustRegister(messagesFullServerFaultGetCoreHash)
	prometheus.MustRegister(messagesFullServerFaultGetTimestamp)
	prometheus.MustRegister(messagesFullServerFaultType)
	prometheus.MustRegister(messagesFullServerFaultMarshalCore)
	prometheus.MustRegister(messagesFullServerFaultMarshalForSF)
	prometheus.MustRegister(messagesFullServerFaultMarshalForSignature)
	prometheus.MustRegister(messagesFullServerFaultMarshalBinary)
	prometheus.MustRegister(messagesFullServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(messagesFullServerFaultMarshalBinary)
	prometheus.MustRegister(messagesFullServerFaultUnmarshall)
	prometheus.MustRegister(messagesFullServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(messagesFullServerFaultUnmarshalBinary)
	prometheus.MustRegister(messagesFullServerFaultGetSignature)
	prometheus.MustRegister(messagesFullServerFaultVerifySignature)
	prometheus.MustRegister(messagesFullServerFaultSign)
	prometheus.MustRegister(messagesFullServerFaultString)
	prometheus.MustRegister(messagesFullServerFaultStringWithSigCnt)
	prometheus.MustRegister(messagesFullServerFaultGetDBHeight)
	prometheus.MustRegister(messagesFullServerFaultValidate)
	prometheus.MustRegister(messagesFullServerFaultSetAlreadyProcessed)
	prometheus.MustRegister(messagesFullServerFaultGetAlreadyProcessed)
	prometheus.MustRegister(messagesFullServerFaultHasEnoughSigs)
	prometheus.MustRegister(messagesFullServerFaultSigTally)
	prometheus.MustRegister(messagesFullServerFaultLeaderExecute)
	prometheus.MustRegister(messagesFullServerFaultFollowerExecute)
	prometheus.MustRegister(messagesFullServerFaultJSONByte)
	prometheus.MustRegister(messagesFullServerFaultJSONString)
	prometheus.MustRegister(messagesFullServerFaultJSONBuffer)
	prometheus.MustRegister(messagesFullServerFaultIsSameAs)
	prometheus.MustRegister(messagesFullServerFaultToAdminBlockEntry)
	prometheus.MustRegister(messagesFullServerFaultNewFullServerFault)
	prometheus.MustRegister(messagesGeneralUnmarshalMessage)
	prometheus.MustRegister(messagesGeneralUnmarshalMessageData)
	prometheus.MustRegister(messagesGeneralMessageName)
	prometheus.MustRegister(messagesGeneralSignSignable)
	prometheus.MustRegister(messagesGeneralVerifyMessage)
	prometheus.MustRegister(messagesHeartbeatIsSameAs)
	prometheus.MustRegister(messagesHeartbeatGetRepeatHash)
	prometheus.MustRegister(messagesHeartbeatGetHash)
	prometheus.MustRegister(messagesHeartbeatGetMsgHash)
	prometheus.MustRegister(messagesHeartbeatGetTimestamp)
	prometheus.MustRegister(messagesHeartbeatType)
	prometheus.MustRegister(messagesHeartbeatInt)
	prometheus.MustRegister(messagesHeartbeatBytes)
	prometheus.MustRegister(messagesHeartbeatUnmarshalBinaryData)
	prometheus.MustRegister(messagesHeartbeatUnmarshalBinary)
	prometheus.MustRegister(messagesHeartbeatMarshalForSignature)
	prometheus.MustRegister(messagesHeartbeatMarshalBinary)
	prometheus.MustRegister(messagesHeartbeatString)
	prometheus.MustRegister(messagesHeartbeatValidate)
	prometheus.MustRegister(messagesHeartbeatLeaderExecute)
	prometheus.MustRegister(messagesHeartbeatJSONByte)
	prometheus.MustRegister(messagesHeartbeatJSONString)
	prometheus.MustRegister(messagesHeartbeatJSONBuffer)
	prometheus.MustRegister(messagesHeartbeatSign)
	prometheus.MustRegister(messagesHeartbeatGetSignature)
	prometheus.MustRegister(messagesInvalidDirectoryBlockIsSameAs)
	prometheus.MustRegister(messagesInvalidDirectoryBlockSign)
	prometheus.MustRegister(messagesInvalidDirectoryBlockGetSignature)
	prometheus.MustRegister(messagesInvalidDirectoryBlockVerifySignature)
	prometheus.MustRegister(messagesInvalidDirectoryBlockProcess)
	prometheus.MustRegister(messagesInvalidDirectoryBlockGetHash)
	prometheus.MustRegister(messagesInvalidDirectoryBlockGetMsgHash)
	prometheus.MustRegister(messagesInvalidDirectoryBlockGetTimestamp)
	prometheus.MustRegister(messagesInvalidDirectoryBlockType)
	prometheus.MustRegister(messagesInvalidDirectoryBlockUnmarshalBinaryData)
	prometheus.MustRegister(messagesInvalidDirectoryBlockUnmarshalBinary)
	prometheus.MustRegister(messagesInvalidDirectoryBlockMarshalBinary)
	prometheus.MustRegister(messagesInvalidDirectoryBlockMarshalForSignature)
	prometheus.MustRegister(messagesInvalidDirectoryBlockJSONByte)
	prometheus.MustRegister(messagesInvalidDirectoryBlockJSONString)
	prometheus.MustRegister(messagesInvalidDirectoryBlockJSONBuffer)
	prometheus.MustRegister(messagesMessageBaseresend)
	prometheus.MustRegister(messagesMessageBaseSendOut)
	prometheus.MustRegister(messagesMessageBaseGetNoResend)
	prometheus.MustRegister(messagesMessageBaseSetNoResend)
	prometheus.MustRegister(messagesMessageBaseIsValid)
	prometheus.MustRegister(messagesMessageBaseSetValid)
	prometheus.MustRegister(messagesMessageBaseMarkSentInvalid)
	prometheus.MustRegister(messagesMessageBaseSentInvlaid)
	prometheus.MustRegister(messagesMessageBaseResend)
	prometheus.MustRegister(messagesMessageBaseExpire)
	prometheus.MustRegister(messagesMessageBaseIsStalled)
	prometheus.MustRegister(messagesMessageBaseSetStall)
	prometheus.MustRegister(messagesMessageBaseGetFullMsgHash)
	prometheus.MustRegister(messagesMessageBaseSetFullMsgHash)
	prometheus.MustRegister(messagesMessageBaseGetOrigin)
	prometheus.MustRegister(messagesMessageBaseSetOrigin)
	prometheus.MustRegister(messagesMessageBaseGetNetworkOrigin)
	prometheus.MustRegister(messagesMessageBaseSetNetworkOrigin)
	prometheus.MustRegister(messagesMessageBaseIsPeer2Peer)
	prometheus.MustRegister(messagesMessageBaseSetPeer2Peer)
	prometheus.MustRegister(messagesMessageBaseIsLocal)
	prometheus.MustRegister(messagesMessageBaseSetLocal)
	prometheus.MustRegister(messagesMessageBaseGetLeaderChainID)
	prometheus.MustRegister(messagesMessageBaseSetLeaderChainID)
	prometheus.MustRegister(messagesMessageBaseGetVMIndex)
	prometheus.MustRegister(messagesMessageBaseSetVMIndex)
	prometheus.MustRegister(messagesMessageBaseGetVMHash)
	prometheus.MustRegister(messagesMessageBaseSetVMHash)
	prometheus.MustRegister(messagesMessageBaseGetMinute)
	prometheus.MustRegister(messagesMessageBaseSetMinute)
	prometheus.MustRegister(messagesMissingDataIsSameAs)
	prometheus.MustRegister(messagesMissingDataProcess)
	prometheus.MustRegister(messagesMissingDataGetRepeatHash)
	prometheus.MustRegister(messagesMissingDataGetHash)
	prometheus.MustRegister(messagesMissingDataGetMsgHash)
	prometheus.MustRegister(messagesMissingDataGetTimestamp)
	prometheus.MustRegister(messagesMissingDataType)
	prometheus.MustRegister(messagesMissingDataInt)
	prometheus.MustRegister(messagesMissingDataBytes)
	prometheus.MustRegister(messagesMissingDataUnmarshalBinaryData)
	prometheus.MustRegister(messagesMissingDataUnmarshalBinary)
	prometheus.MustRegister(messagesMissingDataMarshalBinary)
	prometheus.MustRegister(messagesMissingDataString)
	prometheus.MustRegister(messagesMissingDataLeaderExecute)
	prometheus.MustRegister(messagesMissingDataFollowerExecute)
	prometheus.MustRegister(messagesMissingDataJSONByte)
	prometheus.MustRegister(messagesMissingDataJSONString)
	prometheus.MustRegister(messagesMissingDataJSONBuffer)
	prometheus.MustRegister(messagesMissingDataNewMissingData)
	prometheus.MustRegister(messagesMissingEntryBlocksIsSameAs)
	prometheus.MustRegister(messagesMissingEntryBlocksGetRepeatHash)
	prometheus.MustRegister(messagesMissingEntryBlocksGetHash)
	prometheus.MustRegister(messagesMissingEntryBlocksGetMsgHash)
	prometheus.MustRegister(messagesMissingEntryBlocksType)
	prometheus.MustRegister(messagesMissingEntryBlocksGetTimestamp)
	prometheus.MustRegister(messagesMissingEntryBlocksValidate)
	prometheus.MustRegister(messagesMissingEntryBlocksLeaderExecute)
	prometheus.MustRegister(messagesMissingEntryBlocksFollowerExecute)
	prometheus.MustRegister(messagesMissingEntryBlocksJSONByte)
	prometheus.MustRegister(messagesMissingEntryBlocksJSONString)
	prometheus.MustRegister(messagesMissingEntryBlocksJSONBuffer)
	prometheus.MustRegister(messagesMissingEntryBlocksUnmarshalBinaryData)
	prometheus.MustRegister(messagesMissingEntryBlocksUnmarshalBinary)
	prometheus.MustRegister(messagesMissingEntryBlocksMarshalForSignature)
	prometheus.MustRegister(messagesMissingEntryBlocksMarshalBinary)
	prometheus.MustRegister(messagesMissingEntryBlocksString)
	prometheus.MustRegister(messagesMissingEntryBlocksNewMissingEntryBlocks)
	prometheus.MustRegister(messagesMissingMsgIsSameAs)
	prometheus.MustRegister(messagesMissingMsgGetRepeatHash)
	prometheus.MustRegister(messagesMissingMsgGetHash)
	prometheus.MustRegister(messagesMissingMsgGetMsgHash)
	prometheus.MustRegister(messagesMissingMsgGetTimestamp)
	prometheus.MustRegister(messagesMissingMsgType)
	prometheus.MustRegister(messagesMissingMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesMissingMsgUnmarshalBinary)
	prometheus.MustRegister(messagesMissingMsgMarshalBinary)
	prometheus.MustRegister(messagesMissingMsgString)
	prometheus.MustRegister(messagesMissingMsgValidate)
	prometheus.MustRegister(messagesMissingMsgLeaderExecute)
	prometheus.MustRegister(messagesMissingMsgFollowerExecute)
	prometheus.MustRegister(messagesMissingMsgJSONByte)
	prometheus.MustRegister(messagesMissingMsgJSONString)
	prometheus.MustRegister(messagesMissingMsgJSONBuffer)
	prometheus.MustRegister(messagesMissingMsgAddHeight)
	prometheus.MustRegister(messagesMissingMsgNewMissingMsg)
	prometheus.MustRegister(messagesMissingMsgResponseIsSameAs)
	prometheus.MustRegister(messagesMissingMsgResponseGetRepeatHash)
	prometheus.MustRegister(messagesMissingMsgResponseGetHash)
	prometheus.MustRegister(messagesMissingMsgResponseGetMsgHash)
	prometheus.MustRegister(messagesMissingMsgResponseGetTimestamp)
	prometheus.MustRegister(messagesMissingMsgResponseType)
	prometheus.MustRegister(messagesMissingMsgResponseInt)
	prometheus.MustRegister(messagesMissingMsgResponseBytes)
	prometheus.MustRegister(messagesMissingMsgResponseUnmarshalBinaryData)
	prometheus.MustRegister(messagesMissingMsgResponseUnmarshalBinary)
	prometheus.MustRegister(messagesMissingMsgResponseMarshalBinary)
	prometheus.MustRegister(messagesMissingMsgResponseString)
	prometheus.MustRegister(messagesMissingMsgResponseValidate)
	prometheus.MustRegister(messagesMissingMsgResponseLeaderExecute)
	prometheus.MustRegister(messagesMissingMsgResponseFollowerExecute)
	prometheus.MustRegister(messagesMissingMsgResponseJSONByte)
	prometheus.MustRegister(messagesMissingMsgResponseJSONString)
	prometheus.MustRegister(messagesMissingMsgResponseJSONBuffer)
	prometheus.MustRegister(messagesMissingMsgResponseNewMissingMsgResponse)
	prometheus.MustRegister(messagesRemoveServerMsgGetRepeatHash)
	prometheus.MustRegister(messagesRemoveServerMsgGetHash)
	prometheus.MustRegister(messagesRemoveServerMsgGetMsgHash)
	prometheus.MustRegister(messagesRemoveServerMsgType)
	prometheus.MustRegister(messagesRemoveServerMsgInt)
	prometheus.MustRegister(messagesRemoveServerMsgBytes)
	prometheus.MustRegister(messagesRemoveServerMsgGetTimestamp)
	prometheus.MustRegister(messagesRemoveServerMsgValidate)
	prometheus.MustRegister(messagesRemoveServerMsgComputeVMIndex)
	prometheus.MustRegister(messagesRemoveServerMsgLeaderExecute)
	prometheus.MustRegister(messagesRemoveServerMsgFollowerExecute)
	prometheus.MustRegister(messagesRemoveServerMsgProcess)
	prometheus.MustRegister(messagesRemoveServerMsgJSONByte)
	prometheus.MustRegister(messagesRemoveServerMsgJSONString)
	prometheus.MustRegister(messagesRemoveServerMsgJSONBuffer)
	prometheus.MustRegister(messagesRemoveServerMsgSign)
	prometheus.MustRegister(messagesRemoveServerMsgGetSignature)
	prometheus.MustRegister(messagesRemoveServerMsgVerifySignature)
	prometheus.MustRegister(messagesRemoveServerMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesRemoveServerMsgUnmarshalBinary)
	prometheus.MustRegister(messagesRemoveServerMsgMarshalForSignature)
	prometheus.MustRegister(messagesRemoveServerMsgMarshalBinary)
	prometheus.MustRegister(messagesRemoveServerMsgString)
	prometheus.MustRegister(messagesRemoveServerMsgIsSameAs)
	prometheus.MustRegister(messagesRemoveServerMsgNewRemoveServerMsg)
	prometheus.MustRegister(messagesRequestBlockIsSameAs)
	prometheus.MustRegister(messagesRequestBlockGetRepeatHash)
	prometheus.MustRegister(messagesRequestBlockGetHash)
	prometheus.MustRegister(messagesRequestBlockGetMsgHash)
	prometheus.MustRegister(messagesRequestBlockGetTimestamp)
	prometheus.MustRegister(messagesRequestBlockType)
	prometheus.MustRegister(messagesRequestBlockUnmarshalBinaryData)
	prometheus.MustRegister(messagesRequestBlockUnmarshalBinary)
	prometheus.MustRegister(messagesRequestBlockMarshalForSignature)
	prometheus.MustRegister(messagesRequestBlockMarshalBinary)
	prometheus.MustRegister(messagesRequestBlockJSONByte)
	prometheus.MustRegister(messagesRequestBlockJSONString)
	prometheus.MustRegister(messagesRequestBlockJSONBuffer)
	prometheus.MustRegister(messagesRevealEntryMsgIsSameAs)
	prometheus.MustRegister(messagesRevealEntryMsgProcess)
	prometheus.MustRegister(messagesRevealEntryMsgGetRepeatHash)
	prometheus.MustRegister(messagesRevealEntryMsgGetHash)
	prometheus.MustRegister(messagesRevealEntryMsgGetMsgHash)
	prometheus.MustRegister(messagesRevealEntryMsgGetChainIDHash)
	prometheus.MustRegister(messagesRevealEntryMsgGetTimestamp)
	prometheus.MustRegister(messagesRevealEntryMsgType)
	prometheus.MustRegister(messagesRevealEntryMsgValidateRTN)
	prometheus.MustRegister(messagesRevealEntryMsgValidate)
	prometheus.MustRegister(messagesRevealEntryMsgComputeVMIndex)
	prometheus.MustRegister(messagesRevealEntryMsgLeaderExecute)
	prometheus.MustRegister(messagesRevealEntryMsgFollowerExecute)
	prometheus.MustRegister(messagesRevealEntryMsgJSONByte)
	prometheus.MustRegister(messagesRevealEntryMsgJSONString)
	prometheus.MustRegister(messagesRevealEntryMsgJSONBuffer)
	prometheus.MustRegister(messagesRevealEntryMsgNewRevealEntryMsg)
	prometheus.MustRegister(messagesRevealEntryMsgUnmarshalBinaryData)
	prometheus.MustRegister(messagesRevealEntryMsgUnmarshalBinary)
	prometheus.MustRegister(messagesRevealEntryMsgMarshalBinary)
	prometheus.MustRegister(messagesRevealEntryMsgString)
	prometheus.MustRegister(messagesServerFaultGetRepeatHash)
	prometheus.MustRegister(messagesServerFaultGetHash)
	prometheus.MustRegister(messagesServerFaultGetMsgHash)
	prometheus.MustRegister(messagesServerFaultGetCoreHash)
	prometheus.MustRegister(messagesServerFaultGetTimestamp)
	prometheus.MustRegister(messagesServerFaultType)
	prometheus.MustRegister(messagesServerFaultMarshalForSignature)
	prometheus.MustRegister(messagesServerFaultPreMarshalBinary)
	prometheus.MustRegister(messagesServerFaultMarshalBinary)
	prometheus.MustRegister(messagesServerFaultUnmarshalBinaryData)
	prometheus.MustRegister(messagesServerFaultUnmarshalBinary)
	prometheus.MustRegister(messagesServerFaultGetSignature)
	prometheus.MustRegister(messagesServerFaultVerifySignature)
	prometheus.MustRegister(messagesServerFaultSign)
	prometheus.MustRegister(messagesServerFaultString)
	prometheus.MustRegister(messagesServerFaultGetDBHeight)
	prometheus.MustRegister(messagesServerFaultValidate)
	prometheus.MustRegister(messagesServerFaultLeaderExecute)
	prometheus.MustRegister(messagesServerFaultFollowerExecute)
	prometheus.MustRegister(messagesServerFaultJSONByte)
	prometheus.MustRegister(messagesServerFaultJSONString)
	prometheus.MustRegister(messagesServerFaultJSONBuffer)
	prometheus.MustRegister(messagesServerFaultIsSameAs)
	prometheus.MustRegister(messagesServerFaultNewServerFault)
	prometheus.MustRegister(messagesSignatureTimeoutIsSameAs)
	prometheus.MustRegister(messagesSignatureTimeoutGetRepeatHash)
	prometheus.MustRegister(messagesSignatureTimeoutGetHash)
	prometheus.MustRegister(messagesSignatureTimeoutGetMsgHash)
	prometheus.MustRegister(messagesSignatureTimeoutGetTimestamp)
	prometheus.MustRegister(messagesSignatureTimeoutType)
	prometheus.MustRegister(messagesSignatureTimeoutInt)
	prometheus.MustRegister(messagesSignatureTimeoutBytes)
	prometheus.MustRegister(messagesSignatureTimeoutUnmarshalBinaryData)
	prometheus.MustRegister(messagesSignatureTimeoutUnmarshalBinary)
	prometheus.MustRegister(messagesSignatureTimeoutMarshalForSignature)
	prometheus.MustRegister(messagesSignatureTimeoutMarshalBinary)
	prometheus.MustRegister(messagesSignatureTimeoutGetSignature)
	prometheus.MustRegister(messagesSignatureTimeoutVerifySignature)
	prometheus.MustRegister(messagesSignatureTimeoutSign)
	prometheus.MustRegister(messagesSignatureTimeoutJSONByte)
	prometheus.MustRegister(messagesSignatureTimeoutJSONString)
	prometheus.MustRegister(messagesSignatureTimeoutJSONBuffer)
}