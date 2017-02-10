package blockExtractor

import "github.com/prometheus/client_golang/prometheus"

var (
	factomdblockExtractorFileNotExists = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_FileNotExists_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_FileNotExists",
	})
	factomdblockExtractorBlockExtractorSaveBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_SaveBinary_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_SaveBinary",
	})
	factomdblockExtractorBlockExtractorSaveEntryBinary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_SaveEntryBinary_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_SaveEntryBinary",
	})
	factomdblockExtractorBlockExtractorSaveJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_SaveJSON_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_SaveJSON",
	})
	factomdblockExtractorBlockExtractorSaveEntryJSON = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_SaveEntryJSON_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_SaveEntryJSON",
	})
	factomdblockExtractorBlockExtractorExportEChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportEChain_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportEChain",
	})
	factomdblockExtractorBlockExtractorExportDChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportDChain_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportDChain",
	})
	factomdblockExtractorBlockExtractorExportECChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportECChain_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportECChain",
	})
	factomdblockExtractorBlockExtractorExportAChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportAChain_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportAChain",
	})
	factomdblockExtractorBlockExtractorExportFctChain = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportFctChain_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportFctChain",
	})
	factomdblockExtractorBlockExtractorExportDirBlockInfo = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportDirBlockInfo_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportDirBlockInfo",
	})
	factomdblockExtractorBlockExtractorExportBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportBlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportBlock",
	})
	factomdblockExtractorBlockExtractorExportEntry = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_BlockExtractor_ExportEntry_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_BlockExtractor_ExportEntry",
	})
	factomdblockExtractorexportDBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_exportDBlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_exportDBlock",
	})
	factomdblockExtractorexportEBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_exportEBlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_exportEBlock",
	})
	factomdblockExtractorexportECBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_exportECBlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_exportECBlock",
	})
	factomdblockExtractorexportABlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_exportABlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_exportABlock",
	})
	factomdblockExtractorexportFctBlock = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "factomd_blockExtractor_exportFctBlock_Summary",
		Help: "Summary of calls to  factomd_blockExtractor_exportFctBlock",
	})
)

func InitPrometheus() {
	prometheus.MustRegister(factomdblockExtractorFileNotExists)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorSaveBinary)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorSaveEntryBinary)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorSaveJSON)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorSaveEntryJSON)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportEChain)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportDChain)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportECChain)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportAChain)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportFctChain)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportDirBlockInfo)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportBlock)
	prometheus.MustRegister(factomdblockExtractorBlockExtractorExportEntry)
	prometheus.MustRegister(factomdblockExtractorexportDBlock)
	prometheus.MustRegister(factomdblockExtractorexportEBlock)
	prometheus.MustRegister(factomdblockExtractorexportECBlock)
	prometheus.MustRegister(factomdblockExtractorexportABlock)
	prometheus.MustRegister(factomdblockExtractorexportFctBlock)
}
