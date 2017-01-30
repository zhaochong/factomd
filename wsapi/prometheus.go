// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wsapi

import "github.com/prometheus/client_golang/prometheus"

var (
	//wsapi variables
	activeConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "factomd_wsapi_connections_total",
		Help: "Count of connections that are online",
	})
	v1Requests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_requests_v1",
		Help: "Number of api v1 service requests",
	})
	v1Returned = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_returned_v1",
		Help: "Number of api requests sucessfully returned for v1 requests",
	})
	v1Errors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_errors_v1",
		Help: "Number of api errors returned on v1 calls",
	})
	v1FactoidSubmitSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_submit_v1_summary",
		Help: "Summary of Factoid Submit on v1 calls",
	})
	v1CommitChainSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_commit_chain_v1_summary",
		Help: "Summary of Submit Chain on v1 calls",
	})
	v1RevealChainSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_reveal_chain_v1_summary",
		Help: "Summary of Reveal Chain on v1 calls",
	})
	v1CommitEntrySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_commit_entry_v1_summary",
		Help: "Summary of Commit Entry on v1 calls",
	})
	v1RevealEntrySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_reveal_entry_v1_summary",
		Help: "Summary of Reveal Entry on v1 calls",
	})
	v1DirectoryBlockHeadSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_directory_block_head_v1_summary",
		Help: "Summary of Directory Block Head on v1 calls",
	})
	v1GetRawDataSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_get_raw_data_v1_summary",
		Help: "Summary of Get Raw Data on v1 calls",
	})
	v1GetReceiptSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_get_receipt_v1_summary",
		Help: "Summary of Get Receipt on v1 calls",
	})
	v1DirectoryBlockByKeymrSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_directory_block_by_keymr_v1_summary",
		Help: "Summary of Directory Block By KeyMR on v1 calls",
	})
	v1DirectoryBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_directory_block_by_height_v1_summary",
		Help: "Summary of Directory Block By Height on v1 calls",
	})
	v1EntryBlockByKeyMRSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_block_by_height_v1_summary",
		Help: "Summary of Entry Block By Height on v1 calls",
	})
	v1EntryByHashSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_by_hash_v1_summary",
		Help: "Summary of Entry By Hash on v1 calls",
	})
	v1ChainHeadSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_chain_head_v1_summary",
		Help: "Summary of Chain Head on v1 calls",
	})
	v1EntryCreditBalanceSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_credit_balance_v1_summary",
		Help: "Summary of Entry Credit Balance on v1 calls",
	})
	v1FactoidBalanceSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_balance_v1_summary",
		Help: "Summary of Factoid Balance on v1 calls",
	})
	v1FactoidGetFeeSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_get_fee_v1_summary",
		Help: "Summary of Factoid Get Fee on v1 calls",
	})
	v1PropertiesSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_properties_v1_summary",
		Help: "Summary of properties on v1 calls",
	})
	v1HeightsSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_heights_v1_summary",
		Help: "Summary of heights on v1 calls",
	})
	v1DBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_dblock_by_height_v1_summary",
		Help: "Summary of DBlock By Height on v1 calls",
	})
	v1ECBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_ecblock_by_height_v1_summary",
		Help: "Summary of ECBlock By Height on v1 calls",
	})
	v1FBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_fblock_by_height_v1_summary",
		Help: "Summary of ECBlock By Height on v1 calls",
	})
	v1ABlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_ablock_by_height_v1_summary",
		Help: "Summary of ABlock By Height on v1 calls",
	})

	//wsapiv2 variables

	v2Requests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_requests_v2",
		Help: "Number of api v2 service requests",
	})
	v2Returned = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_returned_v2",
		Help: "Number of api requests sucessfullly returned for v2 requests",
	})
	v2Errors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "factomd_wsapi_errors_v2",
		Help: "Number of api errors returned on v2 calls",
	})
	v2ChainHeadSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_chain_head_v2_summary",
		Help: "Summary of chain-head on v2 calls",
	})
	v2CommitChainSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_commit_chain_v2_summary",
		Help: "Summary of commit-chain on v2 calls",
	})
	v2CommitEntrySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_commit_entry_v2_summary",
		Help: "Summary of commit-entry on v2 calls",
	})
	v2DirectoryBlockSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_directory_block_v2_summary",
		Help: "Summary of directory-block on v2 calls",
	})
	v2DirectoryBlockHeadSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_directory_block_head_v2_summary",
		Help: "Summary of directory-block-head on v2 calls",
	})
	v2EntryBlockSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_block_v2_summary",
		Help: "Summary of entry-block on v2 calls",
	})
	v2EntrySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_v2_summary",
		Help: "Summary of entry on v2 calls",
	})
	v2EntryCreditBalanceSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_credit_balance_v2_summary",
		Help: "Summary of entry-credit-balance on v2 calls",
	})
	v2EntryCreditRateSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_credit_rate_v2_summary",
		Help: "Summary of entry-credit-rate on v2 calls",
	})
	v2FactoidBalanceSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_balance_v2_summary",
		Help: "Summary of factoid-balance on v2 calls",
	})
	v2FactoidSubmitSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_submit_v2_summary",
		Help: "Summary of factoid-submit on v2 calls",
	})
	v2HeightsSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_heights_v2_summary",
		Help: "Summary of heights on v2 calls",
	})
	v2PropertiesSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_properties_v2_summary",
		Help: "Summary of properties on v2 calls",
	})
	v2RawDataSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_raw_data_v2_summary",
		Help: "Summary of raw-data on v2 calls",
	})
	v2ReceiptSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_receipt_v2_summary",
		Help: "Summary of receipt on v2 calls",
	})
	v2RevealEntrySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_reveal_entry_v2_summary",
		Help: "Summary of reveal-entry on v2 calls",
	})
	v2PendingEntriesSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_pending_entries_v2_summary",
		Help: "Summary of pending-entries on v2 calls",
	})
	v2PendingTransactionsSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_pending_transactions_v2_summary",
		Help: "Summary of pending-transactions on v2 calls",
	})
	v2SendRawMessageSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_send_raw_message_v2_summary",
		Help: "Summary of send-raw-message on v2 calls",
	})
	v2transactionSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_transaction_v2_summary",
		Help: "Summary of transaction on v2 calls",
	})
	v2DBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_dblock_by_height_v2_summary",
		Help: "Summary of dblock-by-height on v2 calls",
	})
	v2ECBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_eblock_by_height_v2_summary",
		Help: "Summary of eblock-by-height on v2 calls",
	})
	v2FBlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_fblock_by_height_v2_summary",
		Help: "Summary of fblock-by-height on v2 calls",
	})
	v2ABlockByHeightSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_ablock_by_height_v2_summary",
		Help: "Summary of ablock-by-height on v2 calls",
	})
	v2MapToObjectSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_MapToObject_func_v2_summary",
		Help: "Summary of MapToObject calls in wsapiv2",
	})
	v2ObjectToJStructSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_ObjectToJStruct_func_v2_summary",
		Help: "Summary of ObjectToJStruct calls in wsapiv2",
	})

	// ack.go variables
	v2FactoidAckSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_factoid_ack_v2_summary",
		Help: "Summary of factoid-ack on v2 calls",
	})
	v2EntryAckSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_entry_ack_v2_summary",
		Help: "Summary of entry-ack on v2 calls",
	})
	AckDecodeTransactionToHashesSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_decodetransactiontohashes_func_ack_summary",
		Help: "Summary of DecodeTransactionToHashes calls in ack.go",
	})

	// logs.go variables
	v2InitLogsSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_wsapi_init_logs_func_ack_summary",
		Help: "Summary of InitLogs calls in logs.go",
	})
)

func InitPrometheus() {
	//wsapi variables
	prometheus.MustRegister(activeConnections)
	prometheus.MustRegister(v1Requests)
	prometheus.MustRegister(v1Errors)
	prometheus.MustRegister(v1Returned)
	prometheus.MustRegister(v1FactoidSubmitSummary)
	prometheus.MustRegister(v1CommitChainSummary)
	prometheus.MustRegister(v1RevealChainSummary)
	prometheus.MustRegister(v1CommitEntrySummary)
	prometheus.MustRegister(v1RevealEntrySummary)
	prometheus.MustRegister(v1DirectoryBlockHeadSummary)
	prometheus.MustRegister(v1GetRawDataSummary)
	prometheus.MustRegister(v1GetReceiptSummary)
	prometheus.MustRegister(v1DirectoryBlockByKeymrSummary)
	prometheus.MustRegister(v1DirectoryBlockByHeightSummary)
	prometheus.MustRegister(v1EntryBlockByKeyMRSummary)
	prometheus.MustRegister(v1EntryByHashSummary)
	prometheus.MustRegister(v1ChainHeadSummary)
	prometheus.MustRegister(v1EntryCreditBalanceSummary)
	prometheus.MustRegister(v1FactoidBalanceSummary)
	prometheus.MustRegister(v1FactoidGetFeeSummary)
	prometheus.MustRegister(v1PropertiesSummary)
	prometheus.MustRegister(v1HeightsSummary)
	prometheus.MustRegister(v1DBlockByHeightSummary)
	prometheus.MustRegister(v1ECBlockByHeightSummary)
	prometheus.MustRegister(v1FBlockByHeightSummary)
	prometheus.MustRegister(v1ABlockByHeightSummary)

	/* these are what is being added to each function to measure function time
	       callTime := time.Now().UnixNano()

	   	runTime := time.Now().UnixNano() - callTime
	   	v1DBlockByHeightSummary.Observe(float64(runTime))
	*/

	//wsapiv2 variables
	prometheus.MustRegister(v2Requests)
	prometheus.MustRegister(v2Errors)
	prometheus.MustRegister(v2Returned)
	prometheus.MustRegister(v2ChainHeadSummary)
	prometheus.MustRegister(v2CommitChainSummary)
	prometheus.MustRegister(v2CommitEntrySummary)
	prometheus.MustRegister(v2DirectoryBlockSummary)
	prometheus.MustRegister(v2DirectoryBlockHeadSummary)
	prometheus.MustRegister(v2EntryBlockSummary)
	prometheus.MustRegister(v2EntrySummary)
	prometheus.MustRegister(v2EntryCreditBalanceSummary)
	prometheus.MustRegister(v2EntryCreditRateSummary)
	prometheus.MustRegister(v2FactoidBalanceSummary)
	prometheus.MustRegister(v2FactoidSubmitSummary)
	prometheus.MustRegister(v2HeightsSummary)
	prometheus.MustRegister(v2PropertiesSummary)
	prometheus.MustRegister(v2RawDataSummary)
	prometheus.MustRegister(v2ReceiptSummary)
	prometheus.MustRegister(v2RevealEntrySummary)
	prometheus.MustRegister(v2PendingEntriesSummary)
	prometheus.MustRegister(v2PendingTransactionsSummary)
	prometheus.MustRegister(v2SendRawMessageSummary)
	prometheus.MustRegister(v2transactionSummary)
	prometheus.MustRegister(v2DBlockByHeightSummary)
	prometheus.MustRegister(v2ECBlockByHeightSummary)
	prometheus.MustRegister(v2FBlockByHeightSummary)
	prometheus.MustRegister(v2ABlockByHeightSummary)
	prometheus.MustRegister(v2MapToObjectSummary)
	prometheus.MustRegister(v2ObjectToJStructSummary)

	//ack.go variables
	prometheus.MustRegister(v2FactoidAckSummary)
	prometheus.MustRegister(v2EntryAckSummary)
	prometheus.MustRegister(AckDecodeTransactionToHashesSummary)

	// logs.go variables
	prometheus.MustRegister(v2InitLogsSummary)
}
