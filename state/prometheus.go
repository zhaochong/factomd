package state

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdstateStateIsStateFullySynced = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IsStateFullySynced_Summary",
		Help: "Summary of calls to  factomd_state_State_IsStateFullySynced",
	})
	factomdstateStateGetACKStatus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetACKStatus_Summary",
		Help: "Summary of calls to  factomd_state_State_GetACKStatus",
	})
	factomdstateStateFetchHoldingMessageByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchHoldingMessageByHash_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchHoldingMessageByHash",
	})
	factomdstateStateFetchECTransactionByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchECTransactionByHash_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchECTransactionByHash",
	})
	factomdstateStateFetchFactoidTransactionByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchFactoidTransactionByHash_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchFactoidTransactionByHash",
	})
	factomdstateStateFetchPaidFor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchPaidFor_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchPaidFor",
	})
	factomdstateStateFetchEntryByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchEntryByHash_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchEntryByHash",
	})
	factomdstateAuthorityType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Authority_Type_Summary",
		Help: "Summary of calls to  factomd_state_Authority_Type",
	})
	factomdstateAuthorityVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Authority_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_state_Authority_VerifySignature",
	})
	factomdstateStateVerifyAuthoritySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_VerifyAuthoritySignature_Summary",
		Help: "Summary of calls to  factomd_state_State_VerifyAuthoritySignature",
	})
	factomdstateStateFastVerifyAuthoritySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FastVerifyAuthoritySignature_Summary",
		Help: "Summary of calls to  factomd_state_State_FastVerifyAuthoritySignature",
	})
	factomdstatepkEq = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_pkEq_Summary",
		Help: "Summary of calls to  factomd_state_pkEq",
	})
	factomdstateStateGetAuthority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuthority_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuthority",
	})
	factomdstateStateUpdateAuthSigningKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_UpdateAuthSigningKeys_Summary",
		Help: "Summary of calls to  factomd_state_State_UpdateAuthSigningKeys",
	})
	factomdstateStateUpdateAuthorityFromABEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_UpdateAuthorityFromABEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_UpdateAuthorityFromABEntry",
	})
	factomdstateStateGetAuthorityServerType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuthorityServerType_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuthorityServerType",
	})
	factomdstateStateAddAuthorityFromChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddAuthorityFromChainID_Summary",
		Help: "Summary of calls to  factomd_state_State_AddAuthorityFromChainID",
	})
	factomdstateStateRemoveAuthority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_RemoveAuthority_Summary",
		Help: "Summary of calls to  factomd_state_State_RemoveAuthority",
	})
	factomdstateStateisAuthorityChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_isAuthorityChain_Summary",
		Help: "Summary of calls to  factomd_state_State_isAuthorityChain",
	})
	factomdstateStatecreateAuthority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_createAuthority_Summary",
		Help: "Summary of calls to  factomd_state_State_createAuthority",
	})
	factomdstateStateRepairAuthorities = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_RepairAuthorities_Summary",
		Help: "Summary of calls to  factomd_state_State_RepairAuthorities",
	})
	factomdstateregisterAuthAnchor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_registerAuthAnchor_Summary",
		Help: "Summary of calls to  factomd_state_registerAuthAnchor",
	})
	factomdstateaddServerSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_addServerSigningKey_Summary",
		Help: "Summary of calls to  factomd_state_addServerSigningKey",
	})
	factomdstateDBStateValidNext = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBState_ValidNext_Summary",
		Help: "Summary of calls to  factomd_state_DBState_ValidNext",
	})
	factomdstateDBStateListString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_String_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_String",
	})
	factomdstateDBStateString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBState_String_Summary",
		Help: "Summary of calls to  factomd_state_DBState_String",
	})
	factomdstateDBStateListGetHighestCompletedBlk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_GetHighestCompletedBlk_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_GetHighestCompletedBlk",
	})
	factomdstateDBStateListGetHighestSignedBlk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_GetHighestSignedBlk_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_GetHighestSignedBlk",
	})
	factomdstateDBStateListGetHighestSavedBlk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_GetHighestSavedBlk_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_GetHighestSavedBlk",
	})
	factomdstateDBStateListCatchup = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_Catchup_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_Catchup",
	})
	factomdstatecontainsServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_containsServer_Summary",
		Help: "Summary of calls to  factomd_state_containsServer",
	})
	factomdstateDBStateListFixupLinks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_FixupLinks_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_FixupLinks",
	})
	factomdstateDBStateListProcessBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_ProcessBlocks_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_ProcessBlocks",
	})
	factomdstateDBStateListSignDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_SignDB_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_SignDB",
	})
	factomdstateDBStateListSaveDBStateToDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_SaveDBStateToDB_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_SaveDBStateToDB",
	})
	factomdstateDBStateListUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_UpdateState_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_UpdateState",
	})
	factomdstateDBStateListLast = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_Last_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_Last",
	})
	factomdstateDBStateListHighest = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_Highest_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_Highest",
	})
	factomdstateDBStateListPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_Put_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_Put",
	})
	factomdstateDBStateListGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_Get_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_Get",
	})
	factomdstateDBStateListNewDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DBStateList_NewDBState_Summary",
		Help: "Summary of calls to  factomd_state_DBStateList_NewDBState",
	})
	factomdstateStatesetTimersMakeRequests = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_setTimersMakeRequests_Summary",
		Help: "Summary of calls to  factomd_state_State_setTimersMakeRequests",
	})
	factomdstateStatesyncEntryBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_syncEntryBlocks_Summary",
		Help: "Summary of calls to  factomd_state_State_syncEntryBlocks",
	})
	factomdstateStatesyncEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_syncEntries_Summary",
		Help: "Summary of calls to  factomd_state_State_syncEntries",
	})
	factomdstateStatecatchupEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_catchupEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_State_catchupEBlocks",
	})
	factomdstateFactoidStateReset = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_Reset_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_Reset",
	})
	factomdstateFactoidStateEndOfPeriod = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_EndOfPeriod_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_EndOfPeriod",
	})
	factomdstateFactoidStateGetWallet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_GetWallet_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_GetWallet",
	})
	factomdstateFactoidStateSetWallet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_SetWallet_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_SetWallet",
	})
	factomdstateFactoidStateGetCurrentBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_GetCurrentBlock_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_GetCurrentBlock",
	})
	factomdstateFactoidStateAddTransactionBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_AddTransactionBlock_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_AddTransactionBlock",
	})
	factomdstateFactoidStateAddECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_AddECBlock_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_AddECBlock",
	})
	factomdstateFactoidStateValidateTransactionAge = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_ValidateTransactionAge_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_ValidateTransactionAge",
	})
	factomdstateFactoidStateAddTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_AddTransaction_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_AddTransaction",
	})
	factomdstateFactoidStateGetFactoidBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_GetFactoidBalance_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_GetFactoidBalance",
	})
	factomdstateFactoidStateGetECBalance = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_GetECBalance_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_GetECBalance",
	})
	factomdstateFactoidStateResetBalances = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_ResetBalances_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_ResetBalances",
	})
	factomdstateFactoidStateUpdateECTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_UpdateECTransaction_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_UpdateECTransaction",
	})
	factomdstateFactoidStateUpdateTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_UpdateTransaction_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_UpdateTransaction",
	})
	factomdstateFactoidStateProcessEndOfBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_ProcessEndOfBlock_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_ProcessEndOfBlock",
	})
	factomdstateFactoidStateValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FactoidState_Validate_Summary",
		Help: "Summary of calls to  factomd_state_FactoidState_Validate",
	})
	factomdstateFaultCoreGetHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FaultCore_GetHash_Summary",
		Help: "Summary of calls to  factomd_state_FaultCore_GetHash",
	})
	factomdstateFaultCoreMarshalCore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FaultCore_MarshalCore_Summary",
		Help: "Summary of calls to  factomd_state_FaultCore_MarshalCore",
	})
	factomdstatemarkFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_markFault_Summary",
		Help: "Summary of calls to  factomd_state_markFault",
	})
	factomdstatemarkNoFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_markNoFault_Summary",
		Help: "Summary of calls to  factomd_state_markNoFault",
	})
	factomdstateNegotiationCheck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_NegotiationCheck_Summary",
		Help: "Summary of calls to  factomd_state_NegotiationCheck",
	})
	factomdstateFaultCheck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_FaultCheck_Summary",
		Help: "Summary of calls to  factomd_state_FaultCheck",
	})
	factomdstateprecedingVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_precedingVMIndex_Summary",
		Help: "Summary of calls to  factomd_state_precedingVMIndex",
	})
	factomdstateToggleAuditOffline = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ToggleAuditOffline_Summary",
		Help: "Summary of calls to  factomd_state_ToggleAuditOffline",
	})
	factomdstatecouldIFullFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_couldIFullFault_Summary",
		Help: "Summary of calls to  factomd_state_couldIFullFault",
	})
	factomdstateCraftFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CraftFault_Summary",
		Help: "Summary of calls to  factomd_state_CraftFault",
	})
	factomdstateCraftFullFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CraftFullFault_Summary",
		Help: "Summary of calls to  factomd_state_CraftFullFault",
	})
	factomdstateStateFollowerExecuteSFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteSFault_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteSFault",
	})
	factomdstateExtractFaultCore = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ExtractFaultCore_Summary",
		Help: "Summary of calls to  factomd_state_ExtractFaultCore",
	})
	factomdstateisMyNegotiation = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_isMyNegotiation_Summary",
		Help: "Summary of calls to  factomd_state_isMyNegotiation",
	})
	factomdstateStatematchFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_matchFault_Summary",
		Help: "Summary of calls to  factomd_state_State_matchFault",
	})
	factomdstateStateFollowerExecuteFullFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteFullFault_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteFullFault",
	})
	factomdstateStatepledgedByAudit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_pledgedByAudit_Summary",
		Help: "Summary of calls to  factomd_state_State_pledgedByAudit",
	})
	factomdstateStateReset = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Reset_Summary",
		Help: "Summary of calls to  factomd_state_State_Reset",
	})
	factomdstateStateDoReset = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_DoReset_Summary",
		Help: "Summary of calls to  factomd_state_State_DoReset",
	})
	factomdstateStateGetNetworkSkeletonKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkSkeletonKey_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkSkeletonKey",
	})
	factomdstateStateIntiateNetworkSkeletonIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IntiateNetworkSkeletonIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_IntiateNetworkSkeletonIdentity",
	})
	factomdstateStateAddIdentityFromChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddIdentityFromChainID_Summary",
		Help: "Summary of calls to  factomd_state_State_AddIdentityFromChainID",
	})
	factomdstateStateRemoveIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_RemoveIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_RemoveIdentity",
	})
	factomdstateStateremoveIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_removeIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_removeIdentity",
	})
	factomdstateStateisIdentityChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_isIdentityChain_Summary",
		Help: "Summary of calls to  factomd_state_State_isIdentityChain",
	})
	factomdstateLoadIdentityByEntryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_LoadIdentityByEntryBlock_Summary",
		Help: "Summary of calls to  factomd_state_LoadIdentityByEntryBlock",
	})
	factomdstateLoadIdentityByEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_LoadIdentityByEntry_Summary",
		Help: "Summary of calls to  factomd_state_LoadIdentityByEntry",
	})
	factomdstateStateCreateBlankFactomIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_CreateBlankFactomIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_CreateBlankFactomIdentity",
	})
	factomdstateRegisterFactomIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_RegisterFactomIdentity_Summary",
		Help: "Summary of calls to  factomd_state_RegisterFactomIdentity",
	})
	factomdstateaddIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_addIdentity_Summary",
		Help: "Summary of calls to  factomd_state_addIdentity",
	})
	factomdstatecheckIdentityForFull = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_checkIdentityForFull_Summary",
		Help: "Summary of calls to  factomd_state_checkIdentityForFull",
	})
	factomdstateUpdateManagementKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_UpdateManagementKey_Summary",
		Help: "Summary of calls to  factomd_state_UpdateManagementKey",
	})
	factomdstateregisterIdentityAsServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_registerIdentityAsServer_Summary",
		Help: "Summary of calls to  factomd_state_registerIdentityAsServer",
	})
	factomdstateRegisterBlockSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_RegisterBlockSigningKey_Summary",
		Help: "Summary of calls to  factomd_state_RegisterBlockSigningKey",
	})
	factomdstateUpdateMatryoshkaHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_UpdateMatryoshkaHash_Summary",
		Help: "Summary of calls to  factomd_state_UpdateMatryoshkaHash",
	})
	factomdstateRegisterAnchorSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_RegisterAnchorSigningKey_Summary",
		Help: "Summary of calls to  factomd_state_RegisterAnchorSigningKey",
	})
	factomdstateProcessIdentityToAdminBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessIdentityToAdminBlock_Summary",
		Help: "Summary of calls to  factomd_state_ProcessIdentityToAdminBlock",
	})
	factomdstateStateVerifyIsAuthority = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_VerifyIsAuthority_Summary",
		Help: "Summary of calls to  factomd_state_State_VerifyIsAuthority",
	})
	factomdstateUpdateIdentityStatus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_UpdateIdentityStatus_Summary",
		Help: "Summary of calls to  factomd_state_UpdateIdentityStatus",
	})
	factomdstateIdentityFixMissingKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_FixMissingKeys_Summary",
		Help: "Summary of calls to  factomd_state_Identity_FixMissingKeys",
	})
	factomdstateIdentityVerifySignature = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_VerifySignature_Summary",
		Help: "Summary of calls to  factomd_state_Identity_VerifySignature",
	})
	factomdstateIdentityJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_JSONByte_Summary",
		Help: "Summary of calls to  factomd_state_Identity_JSONByte",
	})
	factomdstateIdentityJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_JSONString_Summary",
		Help: "Summary of calls to  factomd_state_Identity_JSONString",
	})
	factomdstateIdentityString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_String_Summary",
		Help: "Summary of calls to  factomd_state_Identity_String",
	})
	factomdstateCheckSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CheckSig_Summary",
		Help: "Summary of calls to  factomd_state_CheckSig",
	})
	factomdstateCheckExternalIDsLength = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CheckExternalIDsLength_Summary",
		Help: "Summary of calls to  factomd_state_CheckExternalIDsLength",
	})
	factomdstateCheckLength = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CheckLength_Summary",
		Help: "Summary of calls to  factomd_state_CheckLength",
	})
	factomdstateAppendExtIDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_AppendExtIDs_Summary",
		Help: "Summary of calls to  factomd_state_AppendExtIDs",
	})
	factomdstateCheckTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CheckTimestamp_Summary",
		Help: "Summary of calls to  factomd_state_CheckTimestamp",
	})
	factomdstatestatusIsFedOrAudit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_statusIsFedOrAudit_Summary",
		Help: "Summary of calls to  factomd_state_statusIsFedOrAudit",
	})
	factomdstateIdentityIsFull = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Identity_IsFull_Summary",
		Help: "Summary of calls to  factomd_state_Identity_IsFull",
	})
	factomdstateSetDBFinished = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_SetDBFinished_Summary",
		Help: "Summary of calls to  factomd_state_SetDBFinished",
	})
	factomdstateLoadDatabase = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_LoadDatabase_Summary",
		Help: "Summary of calls to  factomd_state_LoadDatabase",
	})
	factomdstateGenerateGenesisBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_GenerateGenesisBlocks_Summary",
		Help: "Summary of calls to  factomd_state_GenerateGenesisBlocks",
	})
	factomdstatePrintState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_PrintState_Summary",
		Help: "Summary of calls to  factomd_state_PrintState",
	})
	factomdstateRequestkey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Request_key_Summary",
		Help: "Summary of calls to  factomd_state_Request_key",
	})
	factomdstateProcessListGetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetAmINegotiator_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetAmINegotiator",
	})
	factomdstateProcessListSetAmINegotiator = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_SetAmINegotiator_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_SetAmINegotiator",
	})
	factomdstateProcessListClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_Clear_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_Clear",
	})
	factomdstateProcessListGetKeysNewEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetKeysNewEntries_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetKeysNewEntries",
	})
	factomdstateProcessListGetNewEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetNewEntry_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetNewEntry",
	})
	factomdstateProcessListLenNewEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_LenNewEntries_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_LenNewEntries",
	})
	factomdstateProcessListComplete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_Complete_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_Complete",
	})
	factomdstateProcessListVMIndexFor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_VMIndexFor_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_VMIndexFor",
	})
	factomdstateSortServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_SortServers_Summary",
		Help: "Summary of calls to  factomd_state_SortServers",
	})
	factomdstateProcessListSortFedServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_SortFedServers_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_SortFedServers",
	})
	factomdstateProcessListSortAuditServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_SortAuditServers_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_SortAuditServers",
	})
	factomdstateProcessListSortDBSigs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_SortDBSigs_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_SortDBSigs",
	})
	factomdstateProcessListFedServerFor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_FedServerFor_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_FedServerFor",
	})
	factomdstateProcessListGetVirtualServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetVirtualServers_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetVirtualServers",
	})
	factomdstateProcessListGetFedServerIndexHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetFedServerIndexHash_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetFedServerIndexHash",
	})
	factomdstateProcessListGetAuditServerIndexHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetAuditServerIndexHash_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetAuditServerIndexHash",
	})
	factomdstateProcessListMakeMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_MakeMap_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_MakeMap",
	})
	factomdstateProcessListPrintMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_PrintMap_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_PrintMap",
	})
	factomdstateProcessListAddFedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddFedServer_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddFedServer",
	})
	factomdstateProcessListAddAuditServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddAuditServer_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddAuditServer",
	})
	factomdstateProcessListRemoveFedServerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_RemoveFedServerHash_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_RemoveFedServerHash",
	})
	factomdstateProcessListRemoveAuditServerHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_RemoveAuditServerHash_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_RemoveAuditServerHash",
	})
	factomdstateProcessListGetAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetAck_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetAck",
	})
	factomdstateProcessListGetAckAt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetAckAt_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetAckAt",
	})
	factomdstateProcessListHasMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_HasMessage_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_HasMessage",
	})
	factomdstateProcessListAddOldMsgs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddOldMsgs_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddOldMsgs",
	})
	factomdstateProcessListDeleteOldMsgs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_DeleteOldMsgs_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_DeleteOldMsgs",
	})
	factomdstateProcessListGetOldMsgs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetOldMsgs_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetOldMsgs",
	})
	factomdstateProcessListAddNewEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddNewEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddNewEBlocks",
	})
	factomdstateProcessListGetNewEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetNewEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetNewEBlocks",
	})
	factomdstateProcessListDeleteEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_DeleteEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_DeleteEBlocks",
	})
	factomdstateProcessListAddNewEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddNewEntry_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddNewEntry",
	})
	factomdstateProcessListDeleteNewEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_DeleteNewEntry_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_DeleteNewEntry",
	})
	factomdstateProcessListCurrentFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_CurrentFault_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_CurrentFault",
	})
	factomdstateProcessListGetLeaderTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetLeaderTimestamp_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetLeaderTimestamp",
	})
	factomdstateProcessListResetDiffSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_ResetDiffSigTally_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_ResetDiffSigTally",
	})
	factomdstateProcessListIncrementDiffSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_IncrementDiffSigTally_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_IncrementDiffSigTally",
	})
	factomdstateProcessListCheckDiffSigTally = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_CheckDiffSigTally_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_CheckDiffSigTally",
	})
	factomdstateProcessListGetRequest = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_GetRequest_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_GetRequest",
	})
	factomdstateProcessListAskDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AskDBState_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AskDBState",
	})
	factomdstateProcessListAsk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_Ask_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_Ask",
	})
	factomdstategetLeaderMin = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_getLeaderMin_Summary",
		Help: "Summary of calls to  factomd_state_getLeaderMin",
	})
	factomdstateProcessListTrimVMList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_TrimVMList_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_TrimVMList",
	})
	factomdstateProcessListProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_Process_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_Process",
	})
	factomdstateProcessListAddToSystemList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddToSystemList_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddToSystemList",
	})
	factomdstateProcessListAddToProcessList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddToProcessList_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddToProcessList",
	})
	factomdstateProcessListContainsDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_ContainsDBSig_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_ContainsDBSig",
	})
	factomdstateProcessListAddDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_AddDBSig_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_AddDBSig",
	})
	factomdstateProcessListString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_String_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_String",
	})
	factomdstateProcessListReset = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessList_Reset_Summary",
		Help: "Summary of calls to  factomd_state_ProcessList_Reset",
	})
	factomdstateNewProcessList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_NewProcessList_Summary",
		Help: "Summary of calls to  factomd_state_NewProcessList",
	})
	factomdstateProcessListsLastList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessLists_LastList_Summary",
		Help: "Summary of calls to  factomd_state_ProcessLists_LastList",
	})
	factomdstateProcessListsUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessLists_UpdateState_Summary",
		Help: "Summary of calls to  factomd_state_ProcessLists_UpdateState",
	})
	factomdstateProcessListsGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessLists_Get_Summary",
		Help: "Summary of calls to  factomd_state_ProcessLists_Get",
	})
	factomdstateProcessListsString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_ProcessLists_String_Summary",
		Help: "Summary of calls to  factomd_state_ProcessLists_String",
	})
	factomdstateNewProcessLists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_NewProcessLists_Summary",
		Help: "Summary of calls to  factomd_state_NewProcessLists",
	})
	factomdstateReplaySave = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_Save_Summary",
		Help: "Summary of calls to  factomd_state_Replay_Save",
	})
	factomdstateMinutes = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Minutes_Summary",
		Help: "Summary of calls to  factomd_state_Minutes",
	})
	factomdstateReplayValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_Valid_Summary",
		Help: "Summary of calls to  factomd_state_Replay_Valid",
	})
	factomdstateReplayIsTSValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_IsTSValid_Summary",
		Help: "Summary of calls to  factomd_state_Replay_IsTSValid",
	})
	factomdstateReplayIsTSValid_ = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_IsTSValid__Summary",
		Help: "Summary of calls to  factomd_state_Replay_IsTSValid_",
	})
	factomdstateReplayIsHashUnique = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_IsHashUnique_Summary",
		Help: "Summary of calls to  factomd_state_Replay_IsHashUnique",
	})
	factomdstateReplaySetHashNow = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_SetHashNow_Summary",
		Help: "Summary of calls to  factomd_state_Replay_SetHashNow",
	})
	factomdstateReplayClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Replay_Clear_Summary",
		Help: "Summary of calls to  factomd_state_Replay_Clear",
	})
	factomdstateSaveFactomdState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_SaveFactomdState_Summary",
		Help: "Summary of calls to  factomd_state_SaveFactomdState",
	})
	factomdstateSaveStateTrimBack = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_SaveState_TrimBack_Summary",
		Help: "Summary of calls to  factomd_state_SaveState_TrimBack",
	})
	factomdstateSaveStateRestoreFactomdState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_SaveState_RestoreFactomdState_Summary",
		Help: "Summary of calls to  factomd_state_SaveState_RestoreFactomdState",
	})
	factomdstateStateSimSetNewKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SimSetNewKeys_Summary",
		Help: "Summary of calls to  factomd_state_State_SimSetNewKeys",
	})
	factomdstateStateSimGetSigKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SimGetSigKey_Summary",
		Help: "Summary of calls to  factomd_state_State_SimGetSigKey",
	})
	factomdstateStateClone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Clone_Summary",
		Help: "Summary of calls to  factomd_state_State_Clone",
	})
	factomdstateStateAddPrefix = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddPrefix_Summary",
		Help: "Summary of calls to  factomd_state_State_AddPrefix",
	})
	factomdstateStateGetFactomNodeName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFactomNodeName_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFactomNodeName",
	})
	factomdstateStateGetDBStatesSent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDBStatesSent_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDBStatesSent",
	})
	factomdstateStateSetDBStatesSent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetDBStatesSent_Summary",
		Help: "Summary of calls to  factomd_state_State_SetDBStatesSent",
	})
	factomdstateStateGetDropRate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDropRate_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDropRate",
	})
	factomdstateStateSetDropRate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetDropRate_Summary",
		Help: "Summary of calls to  factomd_state_State_SetDropRate",
	})
	factomdstateStateSetAuthoritySetString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetAuthoritySetString_Summary",
		Help: "Summary of calls to  factomd_state_State_SetAuthoritySetString",
	})
	factomdstateStateGetAuthoritySetString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuthoritySetString_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuthoritySetString",
	})
	factomdstateStateAddAuthorityDelta = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddAuthorityDelta_Summary",
		Help: "Summary of calls to  factomd_state_State_AddAuthorityDelta",
	})
	factomdstateStateGetAuthorityDeltas = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuthorityDeltas_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuthorityDeltas",
	})
	factomdstateStateGetNetStateOff = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetStateOff_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetStateOff",
	})
	factomdstateStateSetNetStateOff = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetNetStateOff_Summary",
		Help: "Summary of calls to  factomd_state_State_SetNetStateOff",
	})
	factomdstateStateGetRpcUser = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetRpcUser_Summary",
		Help: "Summary of calls to  factomd_state_State_GetRpcUser",
	})
	factomdstateStateGetRpcPass = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetRpcPass_Summary",
		Help: "Summary of calls to  factomd_state_State_GetRpcPass",
	})
	factomdstateStateSetRpcAuthHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetRpcAuthHash_Summary",
		Help: "Summary of calls to  factomd_state_State_SetRpcAuthHash",
	})
	factomdstateStateGetRpcAuthHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetRpcAuthHash_Summary",
		Help: "Summary of calls to  factomd_state_State_GetRpcAuthHash",
	})
	factomdstateStateGetTlsInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetTlsInfo_Summary",
		Help: "Summary of calls to  factomd_state_State_GetTlsInfo",
	})
	factomdstateStateGetFactomdLocations = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFactomdLocations_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFactomdLocations",
	})
	factomdstateStateGetCurrentMinute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetCurrentMinute_Summary",
		Help: "Summary of calls to  factomd_state_State_GetCurrentMinute",
	})
	factomdstateStateIncDBStateAnswerCnt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncDBStateAnswerCnt_Summary",
		Help: "Summary of calls to  factomd_state_State_IncDBStateAnswerCnt",
	})
	factomdstateStateIncFCTSubmits = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncFCTSubmits_Summary",
		Help: "Summary of calls to  factomd_state_State_IncFCTSubmits",
	})
	factomdstateStateIncECCommits = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncECCommits_Summary",
		Help: "Summary of calls to  factomd_state_State_IncECCommits",
	})
	factomdstateStateIncECommits = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncECommits_Summary",
		Help: "Summary of calls to  factomd_state_State_IncECommits",
	})
	factomdstateStateGetAckChange = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAckChange_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAckChange",
	})
	factomdstateStateLoadConfig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadConfig_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadConfig",
	})
	factomdstateStateGetSalt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetSalt_Summary",
		Help: "Summary of calls to  factomd_state_State_GetSalt",
	})
	factomdstateStateInit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Init_Summary",
		Help: "Summary of calls to  factomd_state_State_Init",
	})
	factomdstateStateGetEntryBlockDBHeightComplete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetEntryBlockDBHeightComplete_Summary",
		Help: "Summary of calls to  factomd_state_State_GetEntryBlockDBHeightComplete",
	})
	factomdstateStateSetEntryBlockDBHeightComplete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetEntryBlockDBHeightComplete_Summary",
		Help: "Summary of calls to  factomd_state_State_SetEntryBlockDBHeightComplete",
	})
	factomdstateStateGetEntryBlockDBHeightProcessing = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetEntryBlockDBHeightProcessing_Summary",
		Help: "Summary of calls to  factomd_state_State_GetEntryBlockDBHeightProcessing",
	})
	factomdstateStateSetEntryBlockDBHeightProcessing = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetEntryBlockDBHeightProcessing_Summary",
		Help: "Summary of calls to  factomd_state_State_SetEntryBlockDBHeightProcessing",
	})
	factomdstateStateGetLLeaderHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetLLeaderHeight_Summary",
		Help: "Summary of calls to  factomd_state_State_GetLLeaderHeight",
	})
	factomdstateStateGetFaultTimeout = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFaultTimeout_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFaultTimeout",
	})
	factomdstateStateGetFaultWait = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFaultWait_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFaultWait",
	})
	factomdstateStateGetEntryDBHeightComplete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetEntryDBHeightComplete_Summary",
		Help: "Summary of calls to  factomd_state_State_GetEntryDBHeightComplete",
	})
	factomdstateStateGetMissingEntryCount = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetMissingEntryCount_Summary",
		Help: "Summary of calls to  factomd_state_State_GetMissingEntryCount",
	})
	factomdstateStateGetEBlockKeyMRFromEntryHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetEBlockKeyMRFromEntryHash_Summary",
		Help: "Summary of calls to  factomd_state_State_GetEBlockKeyMRFromEntryHash",
	})
	factomdstateStateGetAndLockDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAndLockDB_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAndLockDB",
	})
	factomdstateStateUnlockDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_UnlockDB_Summary",
		Help: "Summary of calls to  factomd_state_State_UnlockDB",
	})
	factomdstateStateNeeded = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Needed_Summary",
		Help: "Summary of calls to  factomd_state_State_Needed",
	})
	factomdstateStateLoadDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadDBState_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadDBState",
	})
	factomdstateStateLoadDataByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadDataByHash_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadDataByHash",
	})
	factomdstateStateLoadSpecificMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadSpecificMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadSpecificMsg",
	})
	factomdstateStateLoadSpecificMsgAndAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadSpecificMsgAndAck_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadSpecificMsgAndAck",
	})
	factomdstateStateLoadHoldingMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadHoldingMap_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadHoldingMap",
	})
	factomdstateStatefillHoldingMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_fillHoldingMap_Summary",
		Help: "Summary of calls to  factomd_state_State_fillHoldingMap",
	})
	factomdstateStateLoadAcksMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LoadAcksMap_Summary",
		Help: "Summary of calls to  factomd_state_State_LoadAcksMap",
	})
	factomdstateStatefillAcksMap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_fillAcksMap_Summary",
		Help: "Summary of calls to  factomd_state_State_fillAcksMap",
	})
	factomdstateStateGetPendingEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetPendingEntries_Summary",
		Help: "Summary of calls to  factomd_state_State_GetPendingEntries",
	})
	factomdstateStateGetPendingTransactions = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetPendingTransactions_Summary",
		Help: "Summary of calls to  factomd_state_State_GetPendingTransactions",
	})
	factomdstateStateFetchEntryHashFromProcessListsByTxID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FetchEntryHashFromProcessListsByTxID_Summary",
		Help: "Summary of calls to  factomd_state_State_FetchEntryHashFromProcessListsByTxID",
	})
	factomdstateStateIncFactoidTrans = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncFactoidTrans_Summary",
		Help: "Summary of calls to  factomd_state_State_IncFactoidTrans",
	})
	factomdstateStateIncEntryChains = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncEntryChains_Summary",
		Help: "Summary of calls to  factomd_state_State_IncEntryChains",
	})
	factomdstateStateIncEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IncEntries_Summary",
		Help: "Summary of calls to  factomd_state_State_IncEntries",
	})
	factomdstateStateDatabaseContains = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_DatabaseContains_Summary",
		Help: "Summary of calls to  factomd_state_State_DatabaseContains",
	})
	factomdstateStateMessageToLogString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_MessageToLogString_Summary",
		Help: "Summary of calls to  factomd_state_State_MessageToLogString",
	})
	factomdstateStateJournalMessage = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_JournalMessage_Summary",
		Help: "Summary of calls to  factomd_state_State_JournalMessage",
	})
	factomdstateStateGetLeaderVM = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetLeaderVM_Summary",
		Help: "Summary of calls to  factomd_state_State_GetLeaderVM",
	})
	factomdstateStateGetDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDBState_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDBState",
	})
	factomdstateStateGetDirectoryBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDirectoryBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDirectoryBlockByHeight",
	})
	factomdstateStateUpdateState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_UpdateState_Summary",
		Help: "Summary of calls to  factomd_state_State_UpdateState",
	})
	factomdstateStateNoEntryYet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_NoEntryYet_Summary",
		Help: "Summary of calls to  factomd_state_State_NoEntryYet",
	})
	factomdstateStateAddDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddDBSig_Summary",
		Help: "Summary of calls to  factomd_state_State_AddDBSig",
	})
	factomdstateStateAddFedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddFedServer_Summary",
		Help: "Summary of calls to  factomd_state_State_AddFedServer",
	})
	factomdstateStateTrimVMList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_TrimVMList_Summary",
		Help: "Summary of calls to  factomd_state_State_TrimVMList",
	})
	factomdstateStateRemoveFedServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_RemoveFedServer_Summary",
		Help: "Summary of calls to  factomd_state_State_RemoveFedServer",
	})
	factomdstateStateAddAuditServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddAuditServer_Summary",
		Help: "Summary of calls to  factomd_state_State_AddAuditServer",
	})
	factomdstateStateRemoveAuditServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_RemoveAuditServer_Summary",
		Help: "Summary of calls to  factomd_state_State_RemoveAuditServer",
	})
	factomdstateStateGetFedServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFedServers_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFedServers",
	})
	factomdstateStateGetAuditServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuditServers_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuditServers",
	})
	factomdstateStateGetOnlineAuditServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetOnlineAuditServers_Summary",
		Help: "Summary of calls to  factomd_state_State_GetOnlineAuditServers",
	})
	factomdstateStateIsLeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_IsLeader_Summary",
		Help: "Summary of calls to  factomd_state_State_IsLeader",
	})
	factomdstateStateGetVirtualServers = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetVirtualServers_Summary",
		Help: "Summary of calls to  factomd_state_State_GetVirtualServers",
	})
	factomdstateStateGetFactoshisPerEC = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFactoshisPerEC_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFactoshisPerEC",
	})
	factomdstateStateSetFactoshisPerEC = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetFactoshisPerEC_Summary",
		Help: "Summary of calls to  factomd_state_State_SetFactoshisPerEC",
	})
	factomdstateStateGetIdentityChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetIdentityChainID_Summary",
		Help: "Summary of calls to  factomd_state_State_GetIdentityChainID",
	})
	factomdstateStateSetIdentityChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetIdentityChainID_Summary",
		Help: "Summary of calls to  factomd_state_State_SetIdentityChainID",
	})
	factomdstateStateGetDirectoryBlockInSeconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDirectoryBlockInSeconds_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDirectoryBlockInSeconds",
	})
	factomdstateStateSetDirectoryBlockInSeconds = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetDirectoryBlockInSeconds_Summary",
		Help: "Summary of calls to  factomd_state_State_SetDirectoryBlockInSeconds",
	})
	factomdstateStateGetServerPrivateKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetServerPrivateKey_Summary",
		Help: "Summary of calls to  factomd_state_State_GetServerPrivateKey",
	})
	factomdstateStateGetServerPublicKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetServerPublicKey_Summary",
		Help: "Summary of calls to  factomd_state_State_GetServerPublicKey",
	})
	factomdstateStateGetAnchor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAnchor_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAnchor",
	})
	factomdstateStateTallySent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_TallySent_Summary",
		Help: "Summary of calls to  factomd_state_State_TallySent",
	})
	factomdstateStateTallyReceived = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_TallyReceived_Summary",
		Help: "Summary of calls to  factomd_state_State_TallyReceived",
	})
	factomdstateStateGetMessageTalliesSent = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetMessageTalliesSent_Summary",
		Help: "Summary of calls to  factomd_state_State_GetMessageTalliesSent",
	})
	factomdstateStateGetMessageTalliesReceived = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetMessageTalliesReceived_Summary",
		Help: "Summary of calls to  factomd_state_State_GetMessageTalliesReceived",
	})
	factomdstateStateGetFactomdVersion = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFactomdVersion_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFactomdVersion",
	})
	factomdstateStateinitServerKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_initServerKeys_Summary",
		Help: "Summary of calls to  factomd_state_State_initServerKeys",
	})
	factomdstateStateLogInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LogInfo_Summary",
		Help: "Summary of calls to  factomd_state_State_LogInfo",
	})
	factomdstateStateGetAuditHeartBeats = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetAuditHeartBeats_Summary",
		Help: "Summary of calls to  factomd_state_State_GetAuditHeartBeats",
	})
	factomdstateStateSetIsReplaying = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetIsReplaying_Summary",
		Help: "Summary of calls to  factomd_state_State_SetIsReplaying",
	})
	factomdstateStateSetIsDoneReplaying = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetIsDoneReplaying_Summary",
		Help: "Summary of calls to  factomd_state_State_SetIsDoneReplaying",
	})
	factomdstateStateGetTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetTimestamp_Summary",
		Help: "Summary of calls to  factomd_state_State_GetTimestamp",
	})
	factomdstateStateGetTimeOffset = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetTimeOffset_Summary",
		Help: "Summary of calls to  factomd_state_State_GetTimeOffset",
	})
	factomdstateStateSign = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Sign_Summary",
		Help: "Summary of calls to  factomd_state_State_Sign",
	})
	factomdstateStateGetFactoidState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetFactoidState_Summary",
		Help: "Summary of calls to  factomd_state_State_GetFactoidState",
	})
	factomdstateStateSetFactoidState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetFactoidState_Summary",
		Help: "Summary of calls to  factomd_state_State_SetFactoidState",
	})
	factomdstateStateSetPort = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetPort_Summary",
		Help: "Summary of calls to  factomd_state_State_SetPort",
	})
	factomdstateStateGetPort = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetPort_Summary",
		Help: "Summary of calls to  factomd_state_State_GetPort",
	})
	factomdstateStateTickerQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_TickerQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_TickerQueue",
	})
	factomdstateStateTimerMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_TimerMsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_TimerMsgQueue",
	})
	factomdstateStateNetworkInvalidMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_NetworkInvalidMsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_NetworkInvalidMsgQueue",
	})
	factomdstateStateNetworkOutMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_NetworkOutMsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_NetworkOutMsgQueue",
	})
	factomdstateStateInMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_InMsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_InMsgQueue",
	})
	factomdstateStateAPIQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_APIQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_APIQueue",
	})
	factomdstateStateAckQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AckQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_AckQueue",
	})
	factomdstateStateMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_MsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_MsgQueue",
	})
	factomdstateStateGetLeaderTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetLeaderTimestamp_Summary",
		Help: "Summary of calls to  factomd_state_State_GetLeaderTimestamp",
	})
	factomdstateStateSetLeaderTimestamp = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetLeaderTimestamp_Summary",
		Help: "Summary of calls to  factomd_state_State_SetLeaderTimestamp",
	})
	factomdstateStateSetFaultTimeout = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetFaultTimeout_Summary",
		Help: "Summary of calls to  factomd_state_State_SetFaultTimeout",
	})
	factomdstateStateSetFaultWait = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetFaultWait_Summary",
		Help: "Summary of calls to  factomd_state_State_SetFaultWait",
	})
	factomdstateStateGetCfg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetCfg_Summary",
		Help: "Summary of calls to  factomd_state_State_GetCfg",
	})
	factomdstateStateReadCfg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ReadCfg_Summary",
		Help: "Summary of calls to  factomd_state_State_ReadCfg",
	})
	factomdstateStateGetNetworkNumber = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkNumber_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkNumber",
	})
	factomdstateStateGetNetworkID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkID_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkID",
	})
	factomdstateStateGetNetworkBootStrapKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkBootStrapKey_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkBootStrapKey",
	})
	factomdstateStateGetNetworkBootStrapIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkBootStrapIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkBootStrapIdentity",
	})
	factomdstateStateGetNetworkSkeletonIdentity = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkSkeletonIdentity_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkSkeletonIdentity",
	})
	factomdstateStateGetMatryoshka = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetMatryoshka_Summary",
		Help: "Summary of calls to  factomd_state_State_GetMatryoshka",
	})
	factomdstateStateInitLevelDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_InitLevelDB_Summary",
		Help: "Summary of calls to  factomd_state_State_InitLevelDB",
	})
	factomdstateStateInitBoltDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_InitBoltDB_Summary",
		Help: "Summary of calls to  factomd_state_State_InitBoltDB",
	})
	factomdstateStateInitMapDB = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_InitMapDB_Summary",
		Help: "Summary of calls to  factomd_state_State_InitMapDB",
	})
	factomdstateStateString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_String_Summary",
		Help: "Summary of calls to  factomd_state_State_String",
	})
	factomdstateStateShortString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ShortString_Summary",
		Help: "Summary of calls to  factomd_state_State_ShortString",
	})
	factomdstateStateSetString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetString_Summary",
		Help: "Summary of calls to  factomd_state_State_SetString",
	})
	factomdstateStateSummaryHeader = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SummaryHeader_Summary",
		Help: "Summary of calls to  factomd_state_State_SummaryHeader",
	})
	factomdstateStateSetStringConsensus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetStringConsensus_Summary",
		Help: "Summary of calls to  factomd_state_State_SetStringConsensus",
	})
	factomdstateStateSetStringQueues = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetStringQueues_Summary",
		Help: "Summary of calls to  factomd_state_State_SetStringQueues",
	})
	factomdstateStateConstructAuthoritySetString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ConstructAuthoritySetString_Summary",
		Help: "Summary of calls to  factomd_state_State_ConstructAuthoritySetString",
	})
	factomdstateStateGetTrueLeaderHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetTrueLeaderHeight_Summary",
		Help: "Summary of calls to  factomd_state_State_GetTrueLeaderHeight",
	})
	factomdstateStatePrint = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Print_Summary",
		Help: "Summary of calls to  factomd_state_State_Print",
	})
	factomdstateStatePrintln = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Println_Summary",
		Help: "Summary of calls to  factomd_state_State_Println",
	})
	factomdstateStateGetOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetOut_Summary",
		Help: "Summary of calls to  factomd_state_State_GetOut",
	})
	factomdstateStateSetOut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetOut_Summary",
		Help: "Summary of calls to  factomd_state_State_SetOut",
	})
	factomdstateStateGetSystemHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetSystemHeight_Summary",
		Help: "Summary of calls to  factomd_state_State_GetSystemHeight",
	})
	factomdstateStateGetSystemMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetSystemMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_GetSystemMsg",
	})
	factomdstateStateGetInvalidMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetInvalidMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_GetInvalidMsg",
	})
	factomdstateStateProcessInvalidMsgQueue = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessInvalidMsgQueue_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessInvalidMsgQueue",
	})
	factomdstateStateSetPendingSigningKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SetPendingSigningKey_Summary",
		Help: "Summary of calls to  factomd_state_State_SetPendingSigningKey",
	})
	factomdstateStateAddStatus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddStatus_Summary",
		Help: "Summary of calls to  factomd_state_State_AddStatus",
	})
	factomdstateStateGetStatus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetStatus_Summary",
		Help: "Summary of calls to  factomd_state_State_GetStatus",
	})
	factomdstateStateGetLastStatus = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetLastStatus_Summary",
		Help: "Summary of calls to  factomd_state_State_GetLastStatus",
	})
	factomdstateStateexecuteMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_executeMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_executeMsg",
	})
	factomdstateStateProcess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_Process_Summary",
		Help: "Summary of calls to  factomd_state_State_Process",
	})
	factomdstateCheckDBKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_CheckDBKeyMR_Summary",
		Help: "Summary of calls to  factomd_state_CheckDBKeyMR",
	})
	factomdstateStateReviewHolding = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ReviewHolding_Summary",
		Help: "Summary of calls to  factomd_state_State_ReviewHolding",
	})
	factomdstateStateAddDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_AddDBState_Summary",
		Help: "Summary of calls to  factomd_state_State_AddDBState",
	})
	factomdstateStateFollowerExecuteMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteMsg",
	})
	factomdstateStateFollowerExecuteEOM = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteEOM_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteEOM",
	})
	factomdstateStateFollowerExecuteAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteAck_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteAck",
	})
	factomdstateStateFollowerExecuteDBState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteDBState_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteDBState",
	})
	factomdstateStateFollowerExecuteMMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteMMR_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteMMR",
	})
	factomdstateStateFollowerExecuteDataResponse = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteDataResponse_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteDataResponse",
	})
	factomdstateStateFollowerExecuteMissingMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteMissingMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteMissingMsg",
	})
	factomdstateStateFollowerExecuteCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteCommitChain_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteCommitChain",
	})
	factomdstateStateFollowerExecuteCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteCommitEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteCommitEntry",
	})
	factomdstateStateFollowerExecuteRevealEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FollowerExecuteRevealEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_FollowerExecuteRevealEntry",
	})
	factomdstateStateLeaderExecute = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecute_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecute",
	})
	factomdstateStateLeaderExecuteEOM = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecuteEOM_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecuteEOM",
	})
	factomdstateStateLeaderExecuteDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecuteDBSig_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecuteDBSig",
	})
	factomdstateStateLeaderExecuteCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecuteCommitChain_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecuteCommitChain",
	})
	factomdstateStateLeaderExecuteCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecuteCommitEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecuteCommitEntry",
	})
	factomdstateStateLeaderExecuteRevealEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_LeaderExecuteRevealEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_LeaderExecuteRevealEntry",
	})
	factomdstateStateProcessAddServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessAddServer_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessAddServer",
	})
	factomdstateStateProcessRemoveServer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessRemoveServer_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessRemoveServer",
	})
	factomdstateStateProcessChangeServerKey = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessChangeServerKey_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessChangeServerKey",
	})
	factomdstateStateProcessCommitChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessCommitChain_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessCommitChain",
	})
	factomdstateStateProcessCommitEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessCommitEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessCommitEntry",
	})
	factomdstateStateProcessRevealEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessRevealEntry_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessRevealEntry",
	})
	factomdstateStateSendDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SendDBSig_Summary",
		Help: "Summary of calls to  factomd_state_State_SendDBSig",
	})
	factomdstateStateProcessEOM = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessEOM_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessEOM",
	})
	factomdstateStateCheckForIDChange = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_CheckForIDChange_Summary",
		Help: "Summary of calls to  factomd_state_State_CheckForIDChange",
	})
	factomdstateStateProcessDBSig = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessDBSig_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessDBSig",
	})
	factomdstateStateProcessFullServerFault = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessFullServerFault_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessFullServerFault",
	})
	factomdstateStateGetMsg = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetMsg_Summary",
		Help: "Summary of calls to  factomd_state_State_GetMsg",
	})
	factomdstateStateSendHeartBeat = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_SendHeartBeat_Summary",
		Help: "Summary of calls to  factomd_state_State_SendHeartBeat",
	})
	factomdstateStateUpdateECs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_UpdateECs_Summary",
		Help: "Summary of calls to  factomd_state_State_UpdateECs",
	})
	factomdstateStateConsiderSaved = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ConsiderSaved_Summary",
		Help: "Summary of calls to  factomd_state_State_ConsiderSaved",
	})
	factomdstateStateGetNewEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNewEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNewEBlocks",
	})
	factomdstateStatePutNewEBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_PutNewEBlocks_Summary",
		Help: "Summary of calls to  factomd_state_State_PutNewEBlocks",
	})
	factomdstateStatePutNewEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_PutNewEntries_Summary",
		Help: "Summary of calls to  factomd_state_State_PutNewEntries",
	})
	factomdstateStateNextCommit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_NextCommit_Summary",
		Help: "Summary of calls to  factomd_state_State_NextCommit",
	})
	factomdstateStatePutCommit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_PutCommit_Summary",
		Help: "Summary of calls to  factomd_state_State_PutCommit",
	})
	factomdstateStateGetHighestSavedBlk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetHighestSavedBlk_Summary",
		Help: "Summary of calls to  factomd_state_State_GetHighestSavedBlk",
	})
	factomdstateStateGetHighestCompletedBlk = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetHighestCompletedBlk_Summary",
		Help: "Summary of calls to  factomd_state_State_GetHighestCompletedBlk",
	})
	factomdstateStateGetLeaderHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetLeaderHeight_Summary",
		Help: "Summary of calls to  factomd_state_State_GetLeaderHeight",
	})
	factomdstateStateGetHighestKnownBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetHighestKnownBlock_Summary",
		Help: "Summary of calls to  factomd_state_State_GetHighestKnownBlock",
	})
	factomdstateStateGetF = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetF_Summary",
		Help: "Summary of calls to  factomd_state_State_GetF",
	})
	factomdstateStatePutF = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_PutF_Summary",
		Help: "Summary of calls to  factomd_state_State_PutF",
	})
	factomdstateStateGetE = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetE_Summary",
		Help: "Summary of calls to  factomd_state_State_GetE",
	})
	factomdstateStatePutE = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_PutE_Summary",
		Help: "Summary of calls to  factomd_state_State_PutE",
	})
	factomdstateStateComputeVMIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ComputeVMIndex_Summary",
		Help: "Summary of calls to  factomd_state_State_ComputeVMIndex",
	})
	factomdstateStateGetNetworkName = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNetworkName_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNetworkName",
	})
	factomdstateStateGetDBHeightComplete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDBHeightComplete_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDBHeightComplete",
	})
	factomdstateStateGetDirectoryBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetDirectoryBlock_Summary",
		Help: "Summary of calls to  factomd_state_State_GetDirectoryBlock",
	})
	factomdstateStateGetNewHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetNewHash_Summary",
		Help: "Summary of calls to  factomd_state_State_GetNewHash",
	})
	factomdstateStateNewAck = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_NewAck_Summary",
		Help: "Summary of calls to  factomd_state_State_NewAck",
	})
	factomdstateNewDisplayState = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_NewDisplayState_Summary",
		Help: "Summary of calls to  factomd_state_NewDisplayState",
	})
	factomdstateStateCopyStateToControlPanel = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_CopyStateToControlPanel_Summary",
		Help: "Summary of calls to  factomd_state_State_CopyStateToControlPanel",
	})
	factomdstateDeepStateDisplayCopy = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DeepStateDisplayCopy_Summary",
		Help: "Summary of calls to  factomd_state_DeepStateDisplayCopy",
	})
	factomdstateDisplayStateClone = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_DisplayState_Clone_Summary",
		Help: "Summary of calls to  factomd_state_DisplayState_Clone",
	})
	factomdstatemessageLists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_messageLists_Summary",
		Help: "Summary of calls to  factomd_state_messageLists",
	})
	factomdstateStateProcessRecentFERChainEntries = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ProcessRecentFERChainEntries_Summary",
		Help: "Summary of calls to  factomd_state_State_ProcessRecentFERChainEntries",
	})
	factomdstateStateExchangeRateAuthorityIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ExchangeRateAuthorityIsValid_Summary",
		Help: "Summary of calls to  factomd_state_State_ExchangeRateAuthorityIsValid",
	})
	factomdstateStateFerEntryIsValid = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_FerEntryIsValid_Summary",
		Help: "Summary of calls to  factomd_state_State_FerEntryIsValid",
	})
	factomdstateStateGetPredictiveFER = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_GetPredictiveFER_Summary",
		Help: "Summary of calls to  factomd_state_State_GetPredictiveFER",
	})
	factomdstateStateValidatorLoop = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_State_ValidatorLoop_Summary",
		Help: "Summary of calls to  factomd_state_State_ValidatorLoop",
	})
	factomdstateTimertimer = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_state_Timer_timer_Summary",
		Help: "Summary of calls to  factomd_state_Timer_timer",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdstateStateIsStateFullySynced)
	prometheus.MustRegister(factomdstateStateGetACKStatus)
	prometheus.MustRegister(factomdstateStateFetchHoldingMessageByHash)
	prometheus.MustRegister(factomdstateStateFetchECTransactionByHash)
	prometheus.MustRegister(factomdstateStateFetchFactoidTransactionByHash)
	prometheus.MustRegister(factomdstateStateFetchPaidFor)
	prometheus.MustRegister(factomdstateStateFetchEntryByHash)
	prometheus.MustRegister(factomdstateAuthorityType)
	prometheus.MustRegister(factomdstateAuthorityVerifySignature)
	prometheus.MustRegister(factomdstateStateVerifyAuthoritySignature)
	prometheus.MustRegister(factomdstateStateFastVerifyAuthoritySignature)
	prometheus.MustRegister(factomdstatepkEq)
	prometheus.MustRegister(factomdstateStateGetAuthority)
	prometheus.MustRegister(factomdstateStateUpdateAuthSigningKeys)
	prometheus.MustRegister(factomdstateStateUpdateAuthorityFromABEntry)
	prometheus.MustRegister(factomdstateStateGetAuthorityServerType)
	prometheus.MustRegister(factomdstateStateAddAuthorityFromChainID)
	prometheus.MustRegister(factomdstateStateRemoveAuthority)
	prometheus.MustRegister(factomdstateStateisAuthorityChain)
	prometheus.MustRegister(factomdstateStatecreateAuthority)
	prometheus.MustRegister(factomdstateStateRepairAuthorities)
	prometheus.MustRegister(factomdstateregisterAuthAnchor)
	prometheus.MustRegister(factomdstateaddServerSigningKey)
	prometheus.MustRegister(factomdstateDBStateValidNext)
	prometheus.MustRegister(factomdstateDBStateListString)
	prometheus.MustRegister(factomdstateDBStateString)
	prometheus.MustRegister(factomdstateDBStateListGetHighestCompletedBlk)
	prometheus.MustRegister(factomdstateDBStateListGetHighestSignedBlk)
	prometheus.MustRegister(factomdstateDBStateListGetHighestSavedBlk)
	prometheus.MustRegister(factomdstateDBStateListCatchup)
	prometheus.MustRegister(factomdstatecontainsServer)
	prometheus.MustRegister(factomdstateDBStateListFixupLinks)
	prometheus.MustRegister(factomdstateDBStateListProcessBlocks)
	prometheus.MustRegister(factomdstateDBStateListSignDB)
	prometheus.MustRegister(factomdstateDBStateListSaveDBStateToDB)
	prometheus.MustRegister(factomdstateDBStateListUpdateState)
	prometheus.MustRegister(factomdstateDBStateListLast)
	prometheus.MustRegister(factomdstateDBStateListHighest)
	prometheus.MustRegister(factomdstateDBStateListPut)
	prometheus.MustRegister(factomdstateDBStateListGet)
	prometheus.MustRegister(factomdstateDBStateListNewDBState)
	prometheus.MustRegister(factomdstateStatesetTimersMakeRequests)
	prometheus.MustRegister(factomdstateStatesyncEntryBlocks)
	prometheus.MustRegister(factomdstateStatesyncEntries)
	prometheus.MustRegister(factomdstateStatecatchupEBlocks)
	prometheus.MustRegister(factomdstateFactoidStateReset)
	prometheus.MustRegister(factomdstateFactoidStateEndOfPeriod)
	prometheus.MustRegister(factomdstateFactoidStateGetWallet)
	prometheus.MustRegister(factomdstateFactoidStateSetWallet)
	prometheus.MustRegister(factomdstateFactoidStateGetCurrentBlock)
	prometheus.MustRegister(factomdstateFactoidStateAddTransactionBlock)
	prometheus.MustRegister(factomdstateFactoidStateAddECBlock)
	prometheus.MustRegister(factomdstateFactoidStateValidateTransactionAge)
	prometheus.MustRegister(factomdstateFactoidStateAddTransaction)
	prometheus.MustRegister(factomdstateFactoidStateGetFactoidBalance)
	prometheus.MustRegister(factomdstateFactoidStateGetECBalance)
	prometheus.MustRegister(factomdstateFactoidStateResetBalances)
	prometheus.MustRegister(factomdstateFactoidStateUpdateECTransaction)
	prometheus.MustRegister(factomdstateFactoidStateUpdateTransaction)
	prometheus.MustRegister(factomdstateFactoidStateProcessEndOfBlock)
	prometheus.MustRegister(factomdstateFactoidStateValidate)
	prometheus.MustRegister(factomdstateFaultCoreGetHash)
	prometheus.MustRegister(factomdstateFaultCoreMarshalCore)
	prometheus.MustRegister(factomdstatemarkFault)
	prometheus.MustRegister(factomdstatemarkNoFault)
	prometheus.MustRegister(factomdstateNegotiationCheck)
	prometheus.MustRegister(factomdstateFaultCheck)
	prometheus.MustRegister(factomdstateprecedingVMIndex)
	prometheus.MustRegister(factomdstateToggleAuditOffline)
	prometheus.MustRegister(factomdstatecouldIFullFault)
	prometheus.MustRegister(factomdstateCraftFault)
	prometheus.MustRegister(factomdstateCraftFullFault)
	prometheus.MustRegister(factomdstateStateFollowerExecuteSFault)
	prometheus.MustRegister(factomdstateExtractFaultCore)
	prometheus.MustRegister(factomdstateisMyNegotiation)
	prometheus.MustRegister(factomdstateStatematchFault)
	prometheus.MustRegister(factomdstateStateFollowerExecuteFullFault)
	prometheus.MustRegister(factomdstateStatepledgedByAudit)
	prometheus.MustRegister(factomdstateStateReset)
	prometheus.MustRegister(factomdstateStateDoReset)
	prometheus.MustRegister(factomdstateStateGetNetworkSkeletonKey)
	prometheus.MustRegister(factomdstateStateIntiateNetworkSkeletonIdentity)
	prometheus.MustRegister(factomdstateStateAddIdentityFromChainID)
	prometheus.MustRegister(factomdstateStateRemoveIdentity)
	prometheus.MustRegister(factomdstateStateremoveIdentity)
	prometheus.MustRegister(factomdstateStateisIdentityChain)
	prometheus.MustRegister(factomdstateLoadIdentityByEntryBlock)
	prometheus.MustRegister(factomdstateLoadIdentityByEntry)
	prometheus.MustRegister(factomdstateStateCreateBlankFactomIdentity)
	prometheus.MustRegister(factomdstateRegisterFactomIdentity)
	prometheus.MustRegister(factomdstateaddIdentity)
	prometheus.MustRegister(factomdstatecheckIdentityForFull)
	prometheus.MustRegister(factomdstateUpdateManagementKey)
	prometheus.MustRegister(factomdstateregisterIdentityAsServer)
	prometheus.MustRegister(factomdstateRegisterBlockSigningKey)
	prometheus.MustRegister(factomdstateUpdateMatryoshkaHash)
	prometheus.MustRegister(factomdstateRegisterAnchorSigningKey)
	prometheus.MustRegister(factomdstateProcessIdentityToAdminBlock)
	prometheus.MustRegister(factomdstateStateVerifyIsAuthority)
	prometheus.MustRegister(factomdstateUpdateIdentityStatus)
	prometheus.MustRegister(factomdstateIdentityFixMissingKeys)
	prometheus.MustRegister(factomdstateIdentityVerifySignature)
	prometheus.MustRegister(factomdstateIdentityJSONByte)
	prometheus.MustRegister(factomdstateIdentityJSONString)
	prometheus.MustRegister(factomdstateIdentityString)
	prometheus.MustRegister(factomdstateCheckSig)
	prometheus.MustRegister(factomdstateCheckExternalIDsLength)
	prometheus.MustRegister(factomdstateCheckLength)
	prometheus.MustRegister(factomdstateAppendExtIDs)
	prometheus.MustRegister(factomdstateCheckTimestamp)
	prometheus.MustRegister(factomdstatestatusIsFedOrAudit)
	prometheus.MustRegister(factomdstateIdentityIsFull)
	prometheus.MustRegister(factomdstateSetDBFinished)
	prometheus.MustRegister(factomdstateLoadDatabase)
	prometheus.MustRegister(factomdstateGenerateGenesisBlocks)
	prometheus.MustRegister(factomdstatePrintState)
	prometheus.MustRegister(factomdstateRequestkey)
	prometheus.MustRegister(factomdstateProcessListGetAmINegotiator)
	prometheus.MustRegister(factomdstateProcessListSetAmINegotiator)
	prometheus.MustRegister(factomdstateProcessListClear)
	prometheus.MustRegister(factomdstateProcessListGetKeysNewEntries)
	prometheus.MustRegister(factomdstateProcessListGetNewEntry)
	prometheus.MustRegister(factomdstateProcessListLenNewEntries)
	prometheus.MustRegister(factomdstateProcessListComplete)
	prometheus.MustRegister(factomdstateProcessListVMIndexFor)
	prometheus.MustRegister(factomdstateSortServers)
	prometheus.MustRegister(factomdstateProcessListSortFedServers)
	prometheus.MustRegister(factomdstateProcessListSortAuditServers)
	prometheus.MustRegister(factomdstateProcessListSortDBSigs)
	prometheus.MustRegister(factomdstateProcessListFedServerFor)
	prometheus.MustRegister(factomdstateProcessListGetVirtualServers)
	prometheus.MustRegister(factomdstateProcessListGetFedServerIndexHash)
	prometheus.MustRegister(factomdstateProcessListGetAuditServerIndexHash)
	prometheus.MustRegister(factomdstateProcessListMakeMap)
	prometheus.MustRegister(factomdstateProcessListPrintMap)
	prometheus.MustRegister(factomdstateProcessListAddFedServer)
	prometheus.MustRegister(factomdstateProcessListAddAuditServer)
	prometheus.MustRegister(factomdstateProcessListRemoveFedServerHash)
	prometheus.MustRegister(factomdstateProcessListRemoveAuditServerHash)
	prometheus.MustRegister(factomdstateProcessListGetAck)
	prometheus.MustRegister(factomdstateProcessListGetAckAt)
	prometheus.MustRegister(factomdstateProcessListHasMessage)
	prometheus.MustRegister(factomdstateProcessListAddOldMsgs)
	prometheus.MustRegister(factomdstateProcessListDeleteOldMsgs)
	prometheus.MustRegister(factomdstateProcessListGetOldMsgs)
	prometheus.MustRegister(factomdstateProcessListAddNewEBlocks)
	prometheus.MustRegister(factomdstateProcessListGetNewEBlocks)
	prometheus.MustRegister(factomdstateProcessListDeleteEBlocks)
	prometheus.MustRegister(factomdstateProcessListAddNewEntry)
	prometheus.MustRegister(factomdstateProcessListDeleteNewEntry)
	prometheus.MustRegister(factomdstateProcessListCurrentFault)
	prometheus.MustRegister(factomdstateProcessListGetLeaderTimestamp)
	prometheus.MustRegister(factomdstateProcessListResetDiffSigTally)
	prometheus.MustRegister(factomdstateProcessListIncrementDiffSigTally)
	prometheus.MustRegister(factomdstateProcessListCheckDiffSigTally)
	prometheus.MustRegister(factomdstateProcessListGetRequest)
	prometheus.MustRegister(factomdstateProcessListAskDBState)
	prometheus.MustRegister(factomdstateProcessListAsk)
	prometheus.MustRegister(factomdstategetLeaderMin)
	prometheus.MustRegister(factomdstateProcessListTrimVMList)
	prometheus.MustRegister(factomdstateProcessListProcess)
	prometheus.MustRegister(factomdstateProcessListAddToSystemList)
	prometheus.MustRegister(factomdstateProcessListAddToProcessList)
	prometheus.MustRegister(factomdstateProcessListContainsDBSig)
	prometheus.MustRegister(factomdstateProcessListAddDBSig)
	prometheus.MustRegister(factomdstateProcessListString)
	prometheus.MustRegister(factomdstateProcessListReset)
	prometheus.MustRegister(factomdstateNewProcessList)
	prometheus.MustRegister(factomdstateProcessListsLastList)
	prometheus.MustRegister(factomdstateProcessListsUpdateState)
	prometheus.MustRegister(factomdstateProcessListsGet)
	prometheus.MustRegister(factomdstateProcessListsString)
	prometheus.MustRegister(factomdstateNewProcessLists)
	prometheus.MustRegister(factomdstateReplaySave)
	prometheus.MustRegister(factomdstateMinutes)
	prometheus.MustRegister(factomdstateReplayValid)
	prometheus.MustRegister(factomdstateReplayIsTSValid)
	prometheus.MustRegister(factomdstateReplayIsTSValid_)
	prometheus.MustRegister(factomdstateReplayIsHashUnique)
	prometheus.MustRegister(factomdstateReplaySetHashNow)
	prometheus.MustRegister(factomdstateReplayClear)
	prometheus.MustRegister(factomdstateSaveFactomdState)
	prometheus.MustRegister(factomdstateSaveStateTrimBack)
	prometheus.MustRegister(factomdstateSaveStateRestoreFactomdState)
	prometheus.MustRegister(factomdstateStateSimSetNewKeys)
	prometheus.MustRegister(factomdstateStateSimGetSigKey)
	prometheus.MustRegister(factomdstateStateClone)
	prometheus.MustRegister(factomdstateStateAddPrefix)
	prometheus.MustRegister(factomdstateStateGetFactomNodeName)
	prometheus.MustRegister(factomdstateStateGetDBStatesSent)
	prometheus.MustRegister(factomdstateStateSetDBStatesSent)
	prometheus.MustRegister(factomdstateStateGetDropRate)
	prometheus.MustRegister(factomdstateStateSetDropRate)
	prometheus.MustRegister(factomdstateStateSetAuthoritySetString)
	prometheus.MustRegister(factomdstateStateGetAuthoritySetString)
	prometheus.MustRegister(factomdstateStateAddAuthorityDelta)
	prometheus.MustRegister(factomdstateStateGetAuthorityDeltas)
	prometheus.MustRegister(factomdstateStateGetNetStateOff)
	prometheus.MustRegister(factomdstateStateSetNetStateOff)
	prometheus.MustRegister(factomdstateStateGetRpcUser)
	prometheus.MustRegister(factomdstateStateGetRpcPass)
	prometheus.MustRegister(factomdstateStateSetRpcAuthHash)
	prometheus.MustRegister(factomdstateStateGetRpcAuthHash)
	prometheus.MustRegister(factomdstateStateGetTlsInfo)
	prometheus.MustRegister(factomdstateStateGetFactomdLocations)
	prometheus.MustRegister(factomdstateStateGetCurrentMinute)
	prometheus.MustRegister(factomdstateStateIncDBStateAnswerCnt)
	prometheus.MustRegister(factomdstateStateIncFCTSubmits)
	prometheus.MustRegister(factomdstateStateIncECCommits)
	prometheus.MustRegister(factomdstateStateIncECommits)
	prometheus.MustRegister(factomdstateStateGetAckChange)
	prometheus.MustRegister(factomdstateStateLoadConfig)
	prometheus.MustRegister(factomdstateStateGetSalt)
	prometheus.MustRegister(factomdstateStateInit)
	prometheus.MustRegister(factomdstateStateGetEntryBlockDBHeightComplete)
	prometheus.MustRegister(factomdstateStateSetEntryBlockDBHeightComplete)
	prometheus.MustRegister(factomdstateStateGetEntryBlockDBHeightProcessing)
	prometheus.MustRegister(factomdstateStateSetEntryBlockDBHeightProcessing)
	prometheus.MustRegister(factomdstateStateGetLLeaderHeight)
	prometheus.MustRegister(factomdstateStateGetFaultTimeout)
	prometheus.MustRegister(factomdstateStateGetFaultWait)
	prometheus.MustRegister(factomdstateStateGetEntryDBHeightComplete)
	prometheus.MustRegister(factomdstateStateGetMissingEntryCount)
	prometheus.MustRegister(factomdstateStateGetEBlockKeyMRFromEntryHash)
	prometheus.MustRegister(factomdstateStateGetAndLockDB)
	prometheus.MustRegister(factomdstateStateUnlockDB)
	prometheus.MustRegister(factomdstateStateNeeded)
	prometheus.MustRegister(factomdstateStateLoadDBState)
	prometheus.MustRegister(factomdstateStateLoadDataByHash)
	prometheus.MustRegister(factomdstateStateLoadSpecificMsg)
	prometheus.MustRegister(factomdstateStateLoadSpecificMsgAndAck)
	prometheus.MustRegister(factomdstateStateLoadHoldingMap)
	prometheus.MustRegister(factomdstateStatefillHoldingMap)
	prometheus.MustRegister(factomdstateStateLoadAcksMap)
	prometheus.MustRegister(factomdstateStatefillAcksMap)
	prometheus.MustRegister(factomdstateStateGetPendingEntries)
	prometheus.MustRegister(factomdstateStateGetPendingTransactions)
	prometheus.MustRegister(factomdstateStateFetchEntryHashFromProcessListsByTxID)
	prometheus.MustRegister(factomdstateStateIncFactoidTrans)
	prometheus.MustRegister(factomdstateStateIncEntryChains)
	prometheus.MustRegister(factomdstateStateIncEntries)
	prometheus.MustRegister(factomdstateStateDatabaseContains)
	prometheus.MustRegister(factomdstateStateMessageToLogString)
	prometheus.MustRegister(factomdstateStateJournalMessage)
	prometheus.MustRegister(factomdstateStateGetLeaderVM)
	prometheus.MustRegister(factomdstateStateGetDBState)
	prometheus.MustRegister(factomdstateStateGetDirectoryBlockByHeight)
	prometheus.MustRegister(factomdstateStateUpdateState)
	prometheus.MustRegister(factomdstateStateNoEntryYet)
	prometheus.MustRegister(factomdstateStateAddDBSig)
	prometheus.MustRegister(factomdstateStateAddFedServer)
	prometheus.MustRegister(factomdstateStateTrimVMList)
	prometheus.MustRegister(factomdstateStateRemoveFedServer)
	prometheus.MustRegister(factomdstateStateAddAuditServer)
	prometheus.MustRegister(factomdstateStateRemoveAuditServer)
	prometheus.MustRegister(factomdstateStateGetFedServers)
	prometheus.MustRegister(factomdstateStateGetAuditServers)
	prometheus.MustRegister(factomdstateStateGetOnlineAuditServers)
	prometheus.MustRegister(factomdstateStateIsLeader)
	prometheus.MustRegister(factomdstateStateGetVirtualServers)
	prometheus.MustRegister(factomdstateStateGetFactoshisPerEC)
	prometheus.MustRegister(factomdstateStateSetFactoshisPerEC)
	prometheus.MustRegister(factomdstateStateGetIdentityChainID)
	prometheus.MustRegister(factomdstateStateSetIdentityChainID)
	prometheus.MustRegister(factomdstateStateGetDirectoryBlockInSeconds)
	prometheus.MustRegister(factomdstateStateSetDirectoryBlockInSeconds)
	prometheus.MustRegister(factomdstateStateGetServerPrivateKey)
	prometheus.MustRegister(factomdstateStateGetServerPublicKey)
	prometheus.MustRegister(factomdstateStateGetAnchor)
	prometheus.MustRegister(factomdstateStateTallySent)
	prometheus.MustRegister(factomdstateStateTallyReceived)
	prometheus.MustRegister(factomdstateStateGetMessageTalliesSent)
	prometheus.MustRegister(factomdstateStateGetMessageTalliesReceived)
	prometheus.MustRegister(factomdstateStateGetFactomdVersion)
	prometheus.MustRegister(factomdstateStateinitServerKeys)
	prometheus.MustRegister(factomdstateStateLogInfo)
	prometheus.MustRegister(factomdstateStateGetAuditHeartBeats)
	prometheus.MustRegister(factomdstateStateSetIsReplaying)
	prometheus.MustRegister(factomdstateStateSetIsDoneReplaying)
	prometheus.MustRegister(factomdstateStateGetTimestamp)
	prometheus.MustRegister(factomdstateStateGetTimeOffset)
	prometheus.MustRegister(factomdstateStateSign)
	prometheus.MustRegister(factomdstateStateGetFactoidState)
	prometheus.MustRegister(factomdstateStateSetFactoidState)
	prometheus.MustRegister(factomdstateStateSetPort)
	prometheus.MustRegister(factomdstateStateGetPort)
	prometheus.MustRegister(factomdstateStateTickerQueue)
	prometheus.MustRegister(factomdstateStateTimerMsgQueue)
	prometheus.MustRegister(factomdstateStateNetworkInvalidMsgQueue)
	prometheus.MustRegister(factomdstateStateNetworkOutMsgQueue)
	prometheus.MustRegister(factomdstateStateInMsgQueue)
	prometheus.MustRegister(factomdstateStateAPIQueue)
	prometheus.MustRegister(factomdstateStateAckQueue)
	prometheus.MustRegister(factomdstateStateMsgQueue)
	prometheus.MustRegister(factomdstateStateGetLeaderTimestamp)
	prometheus.MustRegister(factomdstateStateSetLeaderTimestamp)
	prometheus.MustRegister(factomdstateStateSetFaultTimeout)
	prometheus.MustRegister(factomdstateStateSetFaultWait)
	prometheus.MustRegister(factomdstateStateGetCfg)
	prometheus.MustRegister(factomdstateStateReadCfg)
	prometheus.MustRegister(factomdstateStateGetNetworkNumber)
	prometheus.MustRegister(factomdstateStateGetNetworkID)
	prometheus.MustRegister(factomdstateStateGetNetworkBootStrapKey)
	prometheus.MustRegister(factomdstateStateGetNetworkBootStrapIdentity)
	prometheus.MustRegister(factomdstateStateGetNetworkSkeletonIdentity)
	prometheus.MustRegister(factomdstateStateGetMatryoshka)
	prometheus.MustRegister(factomdstateStateInitLevelDB)
	prometheus.MustRegister(factomdstateStateInitBoltDB)
	prometheus.MustRegister(factomdstateStateInitMapDB)
	prometheus.MustRegister(factomdstateStateString)
	prometheus.MustRegister(factomdstateStateShortString)
	prometheus.MustRegister(factomdstateStateSetString)
	prometheus.MustRegister(factomdstateStateSummaryHeader)
	prometheus.MustRegister(factomdstateStateSetStringConsensus)
	prometheus.MustRegister(factomdstateStateSetStringQueues)
	prometheus.MustRegister(factomdstateStateConstructAuthoritySetString)
	prometheus.MustRegister(factomdstateStateGetTrueLeaderHeight)
	prometheus.MustRegister(factomdstateStatePrint)
	prometheus.MustRegister(factomdstateStatePrintln)
	prometheus.MustRegister(factomdstateStateGetOut)
	prometheus.MustRegister(factomdstateStateSetOut)
	prometheus.MustRegister(factomdstateStateGetSystemHeight)
	prometheus.MustRegister(factomdstateStateGetSystemMsg)
	prometheus.MustRegister(factomdstateStateGetInvalidMsg)
	prometheus.MustRegister(factomdstateStateProcessInvalidMsgQueue)
	prometheus.MustRegister(factomdstateStateSetPendingSigningKey)
	prometheus.MustRegister(factomdstateStateAddStatus)
	prometheus.MustRegister(factomdstateStateGetStatus)
	prometheus.MustRegister(factomdstateStateGetLastStatus)
	prometheus.MustRegister(factomdstateStateexecuteMsg)
	prometheus.MustRegister(factomdstateStateProcess)
	prometheus.MustRegister(factomdstateCheckDBKeyMR)
	prometheus.MustRegister(factomdstateStateReviewHolding)
	prometheus.MustRegister(factomdstateStateAddDBState)
	prometheus.MustRegister(factomdstateStateFollowerExecuteMsg)
	prometheus.MustRegister(factomdstateStateFollowerExecuteEOM)
	prometheus.MustRegister(factomdstateStateFollowerExecuteAck)
	prometheus.MustRegister(factomdstateStateFollowerExecuteDBState)
	prometheus.MustRegister(factomdstateStateFollowerExecuteMMR)
	prometheus.MustRegister(factomdstateStateFollowerExecuteDataResponse)
	prometheus.MustRegister(factomdstateStateFollowerExecuteMissingMsg)
	prometheus.MustRegister(factomdstateStateFollowerExecuteCommitChain)
	prometheus.MustRegister(factomdstateStateFollowerExecuteCommitEntry)
	prometheus.MustRegister(factomdstateStateFollowerExecuteRevealEntry)
	prometheus.MustRegister(factomdstateStateLeaderExecute)
	prometheus.MustRegister(factomdstateStateLeaderExecuteEOM)
	prometheus.MustRegister(factomdstateStateLeaderExecuteDBSig)
	prometheus.MustRegister(factomdstateStateLeaderExecuteCommitChain)
	prometheus.MustRegister(factomdstateStateLeaderExecuteCommitEntry)
	prometheus.MustRegister(factomdstateStateLeaderExecuteRevealEntry)
	prometheus.MustRegister(factomdstateStateProcessAddServer)
	prometheus.MustRegister(factomdstateStateProcessRemoveServer)
	prometheus.MustRegister(factomdstateStateProcessChangeServerKey)
	prometheus.MustRegister(factomdstateStateProcessCommitChain)
	prometheus.MustRegister(factomdstateStateProcessCommitEntry)
	prometheus.MustRegister(factomdstateStateProcessRevealEntry)
	prometheus.MustRegister(factomdstateStateSendDBSig)
	prometheus.MustRegister(factomdstateStateProcessEOM)
	prometheus.MustRegister(factomdstateStateCheckForIDChange)
	prometheus.MustRegister(factomdstateStateProcessDBSig)
	prometheus.MustRegister(factomdstateStateProcessFullServerFault)
	prometheus.MustRegister(factomdstateStateGetMsg)
	prometheus.MustRegister(factomdstateStateSendHeartBeat)
	prometheus.MustRegister(factomdstateStateUpdateECs)
	prometheus.MustRegister(factomdstateStateConsiderSaved)
	prometheus.MustRegister(factomdstateStateGetNewEBlocks)
	prometheus.MustRegister(factomdstateStatePutNewEBlocks)
	prometheus.MustRegister(factomdstateStatePutNewEntries)
	prometheus.MustRegister(factomdstateStateNextCommit)
	prometheus.MustRegister(factomdstateStatePutCommit)
	prometheus.MustRegister(factomdstateStateGetHighestSavedBlk)
	prometheus.MustRegister(factomdstateStateGetHighestCompletedBlk)
	prometheus.MustRegister(factomdstateStateGetLeaderHeight)
	prometheus.MustRegister(factomdstateStateGetHighestKnownBlock)
	prometheus.MustRegister(factomdstateStateGetF)
	prometheus.MustRegister(factomdstateStatePutF)
	prometheus.MustRegister(factomdstateStateGetE)
	prometheus.MustRegister(factomdstateStatePutE)
	prometheus.MustRegister(factomdstateStateComputeVMIndex)
	prometheus.MustRegister(factomdstateStateGetNetworkName)
	prometheus.MustRegister(factomdstateStateGetDBHeightComplete)
	prometheus.MustRegister(factomdstateStateGetDirectoryBlock)
	prometheus.MustRegister(factomdstateStateGetNewHash)
	prometheus.MustRegister(factomdstateStateNewAck)
	prometheus.MustRegister(factomdstateNewDisplayState)
	prometheus.MustRegister(factomdstateStateCopyStateToControlPanel)
	prometheus.MustRegister(factomdstateDeepStateDisplayCopy)
	prometheus.MustRegister(factomdstateDisplayStateClone)
	prometheus.MustRegister(factomdstatemessageLists)
	prometheus.MustRegister(factomdstateStateProcessRecentFERChainEntries)
	prometheus.MustRegister(factomdstateStateExchangeRateAuthorityIsValid)
	prometheus.MustRegister(factomdstateStateFerEntryIsValid)
	prometheus.MustRegister(factomdstateStateGetPredictiveFER)
	prometheus.MustRegister(factomdstateStateValidatorLoop)
	prometheus.MustRegister(factomdstateTimertimer)
}
