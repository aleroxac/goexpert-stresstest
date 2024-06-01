package cli

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "goexpert-stresstest",
	Short: "Projeto desenvolvido durante treinamento GoExpert, no Desafio Técnico 'Sistema de Stress Test'",
	RunE:  runStressTest(),
}

func init() {
	RootCmd.Flags().StringP("url", "u", "", "URL do serviço a ser testado")
	RootCmd.Flags().IntP("requests", "r", 0, "Número total de requests")
	RootCmd.Flags().IntP("concurrency", "c", 1, "Número de chamadas simultâneas")
}
