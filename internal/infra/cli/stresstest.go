package cli

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aleroxac/goexpert-stresstest/internal/infra/repo"
	"github.com/aleroxac/goexpert-stresstest/internal/usecase"
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func runStressTest() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return err
		}
		if url == "" {
			fmt.Println("Please, provide a valid URL.")
			return errors.New("invalid URL")
		}

		requests, err := cmd.Flags().GetInt("requests")
		if err != nil {
			return err
		}
		if requests <= 0 {
			fmt.Println("Please, provide a valid requests number.")
			return errors.New("invalid requests")
		}

		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			return err
		}
		if concurrency <= 0 {
			concurrency = 1
		}

		repo_request := repo.NewRequestRepository()
		request_usecase := usecase.NewDoRequestUseCase(repo_request)
		input_request := usecase.DoRequestInputDTO{
			URL:         url,
			Requests:    requests,
			Concurrency: concurrency,
		}

		request_output, err := request_usecase.Execute(input_request)
		if err != nil {
			return err
		}

		repo_report := repo.NewReportRepository()
		report_usecase := usecase.NewGenerateReportUseCase(repo_report)
		input := usecase.GenerateReportInputDTO{
			StatusCodes: request_output.StatusCodes,
			Duration:    request_output.TimeSpent,
		}

		report, err := report_usecase.Execute(input)
		if err != nil {
			return err
		}

		report_output, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(report_output))

		return nil
	}
}
