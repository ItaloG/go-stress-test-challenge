/*
Copyright © 2024 ItaloG
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ItaloG/go-stress-test-challenge/internal"
	"github.com/spf13/cobra"
)

var url string
var requests int
var concurrency int

var rootCmd = &cobra.Command{
	Use:   "go-stress-test-challenge",
	Short: "Aplicação para realizar teste de estresse em web servers",
	Long:  `Aplicação para realizar teste em web servers utilizando multi-threading`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.RunStressTest(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado.")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 0, "Número total de requests.")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 0, "Número de chamadas simultâneas.")

	cobra.OnInitialize(validateFlags)
}

func validateFlags() {
	if url == "" || requests <= 0 || concurrency <= 0 {
		fmt.Println("Erro: As flags --url, --requests e --concurrency são obrigatórias.")
		rootCmd.Usage()
		os.Exit(1)
	}
}
