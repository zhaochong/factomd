package databaseOverlay

import "github.com/prometheus/client_golang/prometheus"

var (
	factomddatabaseOverlayOverlayProcessABlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessABlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessABlockBatch",
	})
	factomddatabaseOverlayOverlayProcessABlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessABlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessABlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessABlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessABlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessABlockMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchABlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchABlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchABlock",
	})
	factomddatabaseOverlayOverlayFetchABlockBySecondary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchABlockBySecondary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchABlockBySecondary",
	})
	factomddatabaseOverlayOverlayFetchABlockByPrimary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchABlockByPrimary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchABlockByPrimary",
	})
	factomddatabaseOverlayOverlayFetchABlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchABlockByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchABlockByHeight",
	})
	factomddatabaseOverlayOverlayFetchAllABlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllABlocks_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllABlocks",
	})
	factomddatabaseOverlayOverlayFetchAllABlockKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllABlockKeys_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllABlockKeys",
	})
	factomddatabaseOverlaytoABlocksList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toABlocksList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toABlocksList",
	})
	factomddatabaseOverlayOverlaySaveABlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveABlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveABlockHead",
	})
	factomddatabaseOverlayOverlayFetchABlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchABlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchABlockHead",
	})
	factomddatabaseOverlayAnchorinit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_init_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_init",
	})
	factomddatabaseOverlayOverlayRebuildDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_RebuildDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_RebuildDirBlockInfo",
	})
	factomddatabaseOverlayOverlaySaveAnchorInfoFromEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveAnchorInfoFromEntry_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveAnchorInfoFromEntry",
	})
	factomddatabaseOverlayOverlaySaveAnchorInfoFromEntryMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveAnchorInfoFromEntryMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveAnchorInfoFromEntryMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchAllAnchorInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllAnchorInfo_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllAnchorInfo",
	})
	factomddatabaseOverlayOverlaySaveAnchorInfoAsDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveAnchorInfoAsDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveAnchorInfoAsDirBlockInfo",
	})
	factomddatabaseOverlayAnchorRecordToDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_AnchorRecordToDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_AnchorRecordToDirBlockInfo",
	})
	factomddatabaseOverlayByAnchorDBHeightAccendingLen = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_ByAnchorDBHeightAccending_Len_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_ByAnchorDBHeightAccending_Len",
	})
	factomddatabaseOverlayByAnchorDBHeightAccendingLess = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_ByAnchorDBHeightAccending_Less_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_ByAnchorDBHeightAccending_Less",
	})
	factomddatabaseOverlayByAnchorDBHeightAccendingSwap = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_ByAnchorDBHeightAccending_Swap_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_ByAnchorDBHeightAccending_Swap",
	})
	factomddatabaseOverlayOverlayProcessDBlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessDBlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessDBlockBatch",
	})
	factomddatabaseOverlayOverlayProcessDBlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessDBlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessDBlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessDBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessDBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessDBlockMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchDBlockHeightRange = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockHeightRange_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockHeightRange",
	})
	factomddatabaseOverlayOverlayFetchDBlockHeightByKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockHeightByKeyMR_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockHeightByKeyMR",
	})
	factomddatabaseOverlayOverlayFetchDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlock",
	})
	factomddatabaseOverlayOverlayFetchDBlockByPrimary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockByPrimary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockByPrimary",
	})
	factomddatabaseOverlayOverlayFetchDBlockBySecondary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockBySecondary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockBySecondary",
	})
	factomddatabaseOverlayOverlayFetchDBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockByHeight",
	})
	factomddatabaseOverlayOverlayFetchDBKeyMRByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBKeyMRByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBKeyMRByHeight",
	})
	factomddatabaseOverlayOverlayFetchDBKeyMRByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBKeyMRByHash_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBKeyMRByHash",
	})
	factomddatabaseOverlayOverlayFetchAllDBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllDBlocks_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllDBlocks",
	})
	factomddatabaseOverlayOverlayFetchAllDBlockKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllDBlockKeys_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllDBlockKeys",
	})
	factomddatabaseOverlaytoDBlocksList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toDBlocksList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toDBlocksList",
	})
	factomddatabaseOverlayOverlaySaveDirectoryBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveDirectoryBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveDirectoryBlockHead",
	})
	factomddatabaseOverlayOverlayFetchDBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDBlockHead",
	})
	factomddatabaseOverlayOverlayFetchDirectoryBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDirectoryBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDirectoryBlockHead",
	})
	factomddatabaseOverlayOverlayProcessDirBlockInfoBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessDirBlockInfoBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessDirBlockInfoBatch",
	})
	factomddatabaseOverlayOverlayProcessDirBlockInfoMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessDirBlockInfoMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessDirBlockInfoMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchDirBlockInfoByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDirBlockInfoByHash_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDirBlockInfoByHash",
	})
	factomddatabaseOverlayOverlayFetchDirBlockInfoByKeyMR = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchDirBlockInfoByKeyMR_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchDirBlockInfoByKeyMR",
	})
	factomddatabaseOverlayOverlayFetchAllConfirmedDirBlockInfos = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllConfirmedDirBlockInfos_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllConfirmedDirBlockInfos",
	})
	factomddatabaseOverlayOverlayFetchAllUnconfirmedDirBlockInfos = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllUnconfirmedDirBlockInfos_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllUnconfirmedDirBlockInfos",
	})
	factomddatabaseOverlayOverlayFetchAllDirBlockInfos = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllDirBlockInfos_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllDirBlockInfos",
	})
	factomddatabaseOverlaytoDirBlockInfosList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toDirBlockInfosList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toDirBlockInfosList",
	})
	factomddatabaseOverlayOverlaySaveDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveDirBlockInfo",
	})
	factomddatabaseOverlayOverlayProcessEBlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessEBlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessEBlockBatch",
	})
	factomddatabaseOverlayOverlayProcessEBlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessEBlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessEBlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessEBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessEBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessEBlockMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEBlock",
	})
	factomddatabaseOverlayOverlayFetchEBlockBySecondary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEBlockBySecondary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEBlockBySecondary",
	})
	factomddatabaseOverlayOverlayFetchEBlockByPrimary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEBlockByPrimary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEBlockByPrimary",
	})
	factomddatabaseOverlayOverlayFetchEBKeyMRByHash = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEBKeyMRByHash_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEBKeyMRByHash",
	})
	factomddatabaseOverlayOverlayFetchAllEBlocksByChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllEBlocksByChain_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllEBlocksByChain",
	})
	factomddatabaseOverlayOverlaySaveEBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveEBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveEBlockHead",
	})
	factomddatabaseOverlayOverlayFetchEBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEBlockHead",
	})
	factomddatabaseOverlayOverlayFetchAllEBlockChainIDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllEBlockChainIDs_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllEBlockChainIDs",
	})
	factomddatabaseOverlayOverlayProcessECBlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessECBlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessECBlockBatch",
	})
	factomddatabaseOverlayOverlayProcessECBlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessECBlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessECBlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessECBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessECBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessECBlockMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECBlock",
	})
	factomddatabaseOverlayOverlayFetchECBlockBySecondary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECBlockBySecondary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECBlockBySecondary",
	})
	factomddatabaseOverlayOverlayFetchECBlockByPrimary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECBlockByPrimary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECBlockByPrimary",
	})
	factomddatabaseOverlayOverlayFetchECBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECBlockByHeight",
	})
	factomddatabaseOverlayOverlayFetchAllECBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllECBlocks_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllECBlocks",
	})
	factomddatabaseOverlayOverlayFetchAllECBlockKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllECBlockKeys_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllECBlockKeys",
	})
	factomddatabaseOverlaytoECBlocksList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toECBlocksList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toECBlocksList",
	})
	factomddatabaseOverlayOverlaySaveECBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveECBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveECBlockHead",
	})
	factomddatabaseOverlayOverlayFetchECBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECBlockHead",
	})
	factomddatabaseOverlayOverlayInsertEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_InsertEntry_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_InsertEntry",
	})
	factomddatabaseOverlayOverlayInsertEntryMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_InsertEntryMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_InsertEntryMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchEntry_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchEntry",
	})
	factomddatabaseOverlayOverlayFetchAllEntriesByChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllEntriesByChainID_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllEntriesByChainID",
	})
	factomddatabaseOverlayOverlayFetchAllEntryIDsByChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllEntryIDsByChainID_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllEntryIDsByChainID",
	})
	factomddatabaseOverlayOverlayFetchAllEntryIDs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllEntryIDs_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllEntryIDs",
	})
	factomddatabaseOverlaytoEntryList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toEntryList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toEntryList",
	})
	factomddatabaseOverlayOverlayProcessFBlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessFBlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessFBlockBatch",
	})
	factomddatabaseOverlayOverlayProcessFBlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessFBlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessFBlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessFBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessFBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessFBlockMultiBatch",
	})
	factomddatabaseOverlayOverlayFetchFBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFBlock",
	})
	factomddatabaseOverlayOverlayFetchFBlockBySecondary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFBlockBySecondary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFBlockBySecondary",
	})
	factomddatabaseOverlayOverlayFetchFBlockByPrimary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFBlockByPrimary_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFBlockByPrimary",
	})
	factomddatabaseOverlayOverlayFetchFBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFBlockByHeight",
	})
	factomddatabaseOverlayOverlayFetchAllFBlocks = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllFBlocks_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllFBlocks",
	})
	factomddatabaseOverlayOverlayFetchAllFBlockKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllFBlockKeys_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllFBlockKeys",
	})
	factomddatabaseOverlaytoFactoidList = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_toFactoidList_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_toFactoidList",
	})
	factomddatabaseOverlayOverlaySaveFactoidBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveFactoidBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveFactoidBlockHead",
	})
	factomddatabaseOverlayOverlayFetchFBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFBlockHead",
	})
	factomddatabaseOverlayOverlayFetchFactoidBlockHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFactoidBlockHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFactoidBlockHead",
	})
	factomddatabaseOverlayOverlaySaveIncludedIn = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveIncludedIn_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveIncludedIn",
	})
	factomddatabaseOverlayOverlaySaveIncludedInMultiFromBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveIncludedInMultiFromBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveIncludedInMultiFromBlockMultiBatch",
	})
	factomddatabaseOverlayOverlaySaveIncludedInMultiFromBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveIncludedInMultiFromBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveIncludedInMultiFromBlock",
	})
	factomddatabaseOverlayOverlaySaveIncludedInMultiMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveIncludedInMultiMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveIncludedInMultiMultiBatch",
	})
	factomddatabaseOverlayOverlaySaveIncludedInMulti = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SaveIncludedInMulti_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SaveIncludedInMulti",
	})
	factomddatabaseOverlayOverlayFetchIncludedIn = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchIncludedIn_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchIncludedIn",
	})
	factomddatabaseOverlayinit = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_init_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_init",
	})
	factomddatabaseOverlayOverlayListAllBuckets = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ListAllBuckets_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ListAllBuckets",
	})
	factomddatabaseOverlayOverlaySetExportData = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SetExportData_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SetExportData",
	})
	factomddatabaseOverlayOverlayStartMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_StartMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_StartMultiBatch",
	})
	factomddatabaseOverlayOverlayPutInMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_PutInMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_PutInMultiBatch",
	})
	factomddatabaseOverlayOverlayExecuteMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ExecuteMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ExecuteMultiBatch",
	})
	factomddatabaseOverlayOverlayPutInBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_PutInBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_PutInBatch",
	})
	factomddatabaseOverlayOverlayPut = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Put_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Put",
	})
	factomddatabaseOverlayOverlayListAllKeys = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ListAllKeys_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ListAllKeys",
	})
	factomddatabaseOverlayOverlayGetAll = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_GetAll_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_GetAll",
	})
	factomddatabaseOverlayOverlayGet = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Get_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Get",
	})
	factomddatabaseOverlayOverlayClear = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Clear_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Clear",
	})
	factomddatabaseOverlayOverlayClose = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Close_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Close",
	})
	factomddatabaseOverlayOverlayTrim = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Trim_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Trim",
	})
	factomddatabaseOverlayOverlayDelete = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Delete_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Delete",
	})
	factomddatabaseOverlayNewOverlay = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_NewOverlay_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_NewOverlay",
	})
	factomddatabaseOverlayOverlayFetchBlockByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchBlockByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchBlockByHeight",
	})
	factomddatabaseOverlayOverlayFetchBlockIndexByHeight = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchBlockIndexByHeight_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchBlockIndexByHeight",
	})
	factomddatabaseOverlayOverlayFetchPrimaryIndexBySecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchPrimaryIndexBySecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchPrimaryIndexBySecondaryIndex",
	})
	factomddatabaseOverlayOverlayFetchBlockBySecondaryIndex = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchBlockBySecondaryIndex_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchBlockBySecondaryIndex",
	})
	factomddatabaseOverlayOverlayFetchBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchBlock",
	})
	factomddatabaseOverlayOverlayFetchAllBlocksFromBucket = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllBlocksFromBucket_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllBlocksFromBucket",
	})
	factomddatabaseOverlayOverlayFetchAllBlockKeysFromBucket = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchAllBlockKeysFromBucket_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchAllBlockKeysFromBucket",
	})
	factomddatabaseOverlayOverlayInsert = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_Insert_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_Insert",
	})
	factomddatabaseOverlayOverlayProcessBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessBlockMultiBatch",
	})
	factomddatabaseOverlayOverlayProcessBlockBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessBlockBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessBlockBatch",
	})
	factomddatabaseOverlayOverlayProcessBlockBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessBlockBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessBlockBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayProcessBlockMultiBatchWithoutHead = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_ProcessBlockMultiBatchWithoutHead_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_ProcessBlockMultiBatchWithoutHead",
	})
	factomddatabaseOverlayOverlayFetchHeadIndexByChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchHeadIndexByChainID_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchHeadIndexByChainID",
	})
	factomddatabaseOverlayOverlayFetchChainHeadByChainID = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchChainHeadByChainID_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchChainHeadByChainID",
	})
	factomddatabaseOverlayOverlayFetchBlockIndexesInHeightRange = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchBlockIndexesInHeightRange_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchBlockIndexesInHeightRange",
	})
	factomddatabaseOverlayOverlayDoesKeyExist = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_DoesKeyExist_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_DoesKeyExist",
	})
	factomddatabaseOverlayOverlayGetEntryType = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_GetEntryType_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_GetEntryType",
	})
	factomddatabaseOverlayOverlaySavePaidFor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SavePaidFor_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SavePaidFor",
	})
	factomddatabaseOverlayOverlaySavePaidForMultiFromBlockMultiBatch = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SavePaidForMultiFromBlockMultiBatch_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SavePaidForMultiFromBlockMultiBatch",
	})
	factomddatabaseOverlayOverlaySavePaidForMultiFromBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_SavePaidForMultiFromBlock_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_SavePaidForMultiFromBlock",
	})
	factomddatabaseOverlayOverlayFetchPaidFor = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchPaidFor_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchPaidFor",
	})
	factomddatabaseOverlayOverlayFetchFactoidTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchFactoidTransaction_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchFactoidTransaction",
	})
	factomddatabaseOverlayOverlayFetchECTransaction = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_databaseOverlay_Overlay_FetchECTransaction_Summary",
		Help: "Summary of calls to  factomd_databaseOverlay_Overlay_FetchECTransaction",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessABlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessABlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessABlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchABlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchABlockBySecondary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchABlockByPrimary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchABlockByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllABlocks)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllABlockKeys)
	prometheus.MustRegister(factomddatabaseOverlaytoABlocksList)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveABlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchABlockHead)
	prometheus.MustRegister(factomddatabaseOverlayAnchorinit)
	prometheus.MustRegister(factomddatabaseOverlayOverlayRebuildDirBlockInfo)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveAnchorInfoFromEntry)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveAnchorInfoFromEntryMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllAnchorInfo)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveAnchorInfoAsDirBlockInfo)
	prometheus.MustRegister(factomddatabaseOverlayAnchorRecordToDirBlockInfo)
	prometheus.MustRegister(factomddatabaseOverlayByAnchorDBHeightAccendingLen)
	prometheus.MustRegister(factomddatabaseOverlayByAnchorDBHeightAccendingLess)
	prometheus.MustRegister(factomddatabaseOverlayByAnchorDBHeightAccendingSwap)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessDBlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessDBlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessDBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockHeightRange)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockHeightByKeyMR)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockByPrimary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockBySecondary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBKeyMRByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBKeyMRByHash)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllDBlocks)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllDBlockKeys)
	prometheus.MustRegister(factomddatabaseOverlaytoDBlocksList)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveDirectoryBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDirectoryBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessDirBlockInfoBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessDirBlockInfoMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDirBlockInfoByHash)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchDirBlockInfoByKeyMR)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllConfirmedDirBlockInfos)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllUnconfirmedDirBlockInfos)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllDirBlockInfos)
	prometheus.MustRegister(factomddatabaseOverlaytoDirBlockInfosList)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveDirBlockInfo)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessEBlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessEBlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessEBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEBlockBySecondary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEBlockByPrimary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEBKeyMRByHash)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllEBlocksByChain)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveEBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllEBlockChainIDs)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessECBlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessECBlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessECBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECBlockBySecondary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECBlockByPrimary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECBlockByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllECBlocks)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllECBlockKeys)
	prometheus.MustRegister(factomddatabaseOverlaytoECBlocksList)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveECBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayInsertEntry)
	prometheus.MustRegister(factomddatabaseOverlayOverlayInsertEntryMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchEntry)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllEntriesByChainID)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllEntryIDsByChainID)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllEntryIDs)
	prometheus.MustRegister(factomddatabaseOverlaytoEntryList)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessFBlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessFBlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessFBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFBlockBySecondary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFBlockByPrimary)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFBlockByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllFBlocks)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllFBlockKeys)
	prometheus.MustRegister(factomddatabaseOverlaytoFactoidList)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveFactoidBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFactoidBlockHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveIncludedIn)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveIncludedInMultiFromBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveIncludedInMultiFromBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveIncludedInMultiMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySaveIncludedInMulti)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchIncludedIn)
	prometheus.MustRegister(factomddatabaseOverlayinit)
	prometheus.MustRegister(factomddatabaseOverlayOverlayListAllBuckets)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySetExportData)
	prometheus.MustRegister(factomddatabaseOverlayOverlayStartMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayPutInMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayExecuteMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayPutInBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayPut)
	prometheus.MustRegister(factomddatabaseOverlayOverlayListAllKeys)
	prometheus.MustRegister(factomddatabaseOverlayOverlayGetAll)
	prometheus.MustRegister(factomddatabaseOverlayOverlayGet)
	prometheus.MustRegister(factomddatabaseOverlayOverlayClear)
	prometheus.MustRegister(factomddatabaseOverlayOverlayClose)
	prometheus.MustRegister(factomddatabaseOverlayOverlayTrim)
	prometheus.MustRegister(factomddatabaseOverlayOverlayDelete)
	prometheus.MustRegister(factomddatabaseOverlayNewOverlay)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchBlockByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchBlockIndexByHeight)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchPrimaryIndexBySecondaryIndex)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchBlockBySecondaryIndex)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllBlocksFromBucket)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchAllBlockKeysFromBucket)
	prometheus.MustRegister(factomddatabaseOverlayOverlayInsert)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessBlockBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessBlockBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayProcessBlockMultiBatchWithoutHead)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchHeadIndexByChainID)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchChainHeadByChainID)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchBlockIndexesInHeightRange)
	prometheus.MustRegister(factomddatabaseOverlayOverlayDoesKeyExist)
	prometheus.MustRegister(factomddatabaseOverlayOverlayGetEntryType)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySavePaidFor)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySavePaidForMultiFromBlockMultiBatch)
	prometheus.MustRegister(factomddatabaseOverlayOverlaySavePaidForMultiFromBlock)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchPaidFor)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchFactoidTransaction)
	prometheus.MustRegister(factomddatabaseOverlayOverlayFetchECTransaction)
}
