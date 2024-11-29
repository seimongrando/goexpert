package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"stress-test/stress"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	// Adiciona o comando "test" ao comando raiz
	rootCmd.AddCommand(testCmd)

	// Flags para o comando "test"
	testCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
	testCmd.Flags().IntVar(&requests, "requests", 0, "Número total de requisições (obrigatório)")
	testCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
	testCmd.MarkFlagRequired("url")
	testCmd.MarkFlagRequired("requests")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Realiza um teste de carga",
	Long:  "Executa requisições HTTP para um serviço web em paralelo, gerando um relatório ao final.",
	Run: func(cmd *cobra.Command, args []string) {
		if requests <= 0 || concurrency <= 0 {
			fmt.Println("Parâmetros inválidos. --requests e --concurrency devem ser maiores que 0.")
			return
		}

		report := stress.RunLoadTest(url, requests, concurrency)
		generateReport(report)
	},
}

func generateReport(report stress.Report) {
	fmt.Println("\n--- Teste de Carga Concluído ---")
	fmt.Printf("Tempo total: %v\n", report.TotalTime)
	fmt.Printf("Total de requests: %d\n", report.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", report.Status200)
	if len(report.StatusOthers) > 0 {
		fmt.Println("Distribuição de outros códigos de status:")
		for status, count := range report.StatusOthers {
			fmt.Printf("  Status %d: %d\n", status, count)
		}
	}
}
