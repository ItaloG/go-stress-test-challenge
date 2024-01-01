package internal

import (
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string)

func RunStressTest(url string, requests int, concurrency int) {
	duration, statusCodes := PerformStressTest(url, requests, concurrency)

	GenerateReport(duration, requests, statusCodes)
}
