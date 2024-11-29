package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "load-tester",
	Short: "CLI para realizar testes de carga em serviços web",
	Long:  "Uma ferramenta CLI para realizar testes de carga HTTP, configurando URL, número de requisições e concorrência.",
}

// Execute inicia o comando raiz
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
