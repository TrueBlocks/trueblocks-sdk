package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v5/pkg/types"
)

func reportScrape(logger *slog.Logger, report *scraperReport) {
	msg := fmt.Sprintf("behind (% 10.10s)...", report.Chain)
	if report.Staged < 30 {
		msg = fmt.Sprintf("atHead (% 10.10s)...", report.Chain)
	}
	logger.Info(msg,
		"head", report.Head,
		"unripe", -report.Unripe,
		"staged", -report.Staged,
		"finalized", -report.Finalized,
		"blockCnt", report.BlockCnt)
}

type scraperReport struct {
	Chain     string `json:"chain"`
	BlockCnt  int    `json:"blockCnt"`
	Head      int    `json:"head"`
	Unripe    int    `json:"unripe"`
	Staged    int    `json:"staged"`
	Finalized int    `json:"finalized"`
	Time      string `json:"time"`
}

func (r *scraperReport) String() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

func reportScrapeRun(meta *types.MetaData, chain string, blockCnt int) *scraperReport {
	return &scraperReport{
		Chain:     chain,
		BlockCnt:  blockCnt,
		Head:      int(meta.Latest),
		Unripe:    int(meta.Latest) - int(meta.Unripe),
		Staged:    int(meta.Latest) - int(meta.Staging),
		Finalized: int(meta.Latest) - int(meta.Finalized),
		Time:      time.Now().Format("01-02 15:04:05"),
	}
}
