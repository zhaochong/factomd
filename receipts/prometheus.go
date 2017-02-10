package receipts

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdreceiptsFileNotExists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_FileNotExists_Summary",
		Help: "Summary of calls to  factomd_receipts_FileNotExists",
	})
	factomdreceiptsSave = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Save_Summary",
		Help: "Summary of calls to  factomd_receipts_Save",
	})
	factomdreceiptsExportEntryReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_ExportEntryReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_ExportEntryReceipt",
	})
	factomdreceiptsExportAllEntryReceipts = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_ExportAllEntryReceipts_Summary",
		Help: "Summary of calls to  factomd_receipts_ExportAllEntryReceipts",
	})
	factomdreceiptsReceiptTrimReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_TrimReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_TrimReceipt",
	})
	factomdreceiptsReceiptValidate = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_Validate_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_Validate",
	})
	factomdreceiptsReceiptIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_IsSameAs",
	})
	factomdreceiptsReceiptJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_JSONByte_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_JSONByte",
	})
	factomdreceiptsReceiptJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_JSONString_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_JSONString",
	})
	factomdreceiptsReceiptCustomMarshalString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_CustomMarshalString_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_CustomMarshalString",
	})
	factomdreceiptsReceiptDecodeString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_Receipt_DecodeString_Summary",
		Help: "Summary of calls to  factomd_receipts_Receipt_DecodeString",
	})
	factomdreceiptsDecodeReceiptString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_DecodeReceiptString_Summary",
		Help: "Summary of calls to  factomd_receipts_DecodeReceiptString",
	})
	factomdreceiptsJSONJSONByte = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_JSON_JSONByte_Summary",
		Help: "Summary of calls to  factomd_receipts_JSON_JSONByte",
	})
	factomdreceiptsJSONJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_JSON_JSONString_Summary",
		Help: "Summary of calls to  factomd_receipts_JSON_JSONString",
	})
	factomdreceiptsJSONString = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_JSON_String_Summary",
		Help: "Summary of calls to  factomd_receipts_JSON_String",
	})
	factomdreceiptsJSONIsSameAs = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_JSON_IsSameAs_Summary",
		Help: "Summary of calls to  factomd_receipts_JSON_IsSameAs",
	})
	factomdreceiptsCreateFullReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_CreateFullReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_CreateFullReceipt",
	})
	factomdreceiptsCreateMinimalReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_CreateMinimalReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_CreateMinimalReceipt",
	})
	factomdreceiptsCreateReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_CreateReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_CreateReceipt",
	})
	factomdreceiptsVerifyFullReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_VerifyFullReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_VerifyFullReceipt",
	})
	factomdreceiptsVerifyMinimalReceipt = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_receipts_VerifyMinimalReceipt_Summary",
		Help: "Summary of calls to  factomd_receipts_VerifyMinimalReceipt",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdreceiptsFileNotExists)
	prometheus.MustRegister(factomdreceiptsSave)
	prometheus.MustRegister(factomdreceiptsExportEntryReceipt)
	prometheus.MustRegister(factomdreceiptsExportAllEntryReceipts)
	prometheus.MustRegister(factomdreceiptsReceiptTrimReceipt)
	prometheus.MustRegister(factomdreceiptsReceiptValidate)
	prometheus.MustRegister(factomdreceiptsReceiptIsSameAs)
	prometheus.MustRegister(factomdreceiptsReceiptJSONByte)
	prometheus.MustRegister(factomdreceiptsReceiptJSONString)
	prometheus.MustRegister(factomdreceiptsReceiptCustomMarshalString)
	prometheus.MustRegister(factomdreceiptsReceiptDecodeString)
	prometheus.MustRegister(factomdreceiptsDecodeReceiptString)
	prometheus.MustRegister(factomdreceiptsJSONJSONByte)
	prometheus.MustRegister(factomdreceiptsJSONJSONString)
	prometheus.MustRegister(factomdreceiptsJSONString)
	prometheus.MustRegister(factomdreceiptsJSONIsSameAs)
	prometheus.MustRegister(factomdreceiptsCreateFullReceipt)
	prometheus.MustRegister(factomdreceiptsCreateMinimalReceipt)
	prometheus.MustRegister(factomdreceiptsCreateReceipt)
	prometheus.MustRegister(factomdreceiptsVerifyFullReceipt)
	prometheus.MustRegister(factomdreceiptsVerifyMinimalReceipt)
}
