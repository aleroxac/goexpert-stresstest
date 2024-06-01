package usecase

import (
	"log"
	"time"

	"github.com/aleroxac/goexpert-stresstest/internal/entity"
)

type GenerateReportInputDTO struct {
	StatusCodes []int
	Duration    time.Duration
}

type GenerateReportOutputDTO struct {
	TimeSpent          string         `json:"time_spent"`
	RequestsMade       int            `json:"requests_made"`
	SuccessfulRequests int            `json:"successful_requests"`
	FailedRequests     map[string]int `json:"failed_requests"`
}

type GenerateReportUseCase struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewGenerateReportUseCase(report_repository entity.ReportRepositoryInterface) *GenerateReportUseCase {
	return &GenerateReportUseCase{
		ReportRepository: report_repository,
	}
}

func (r *GenerateReportUseCase) Execute(input GenerateReportInputDTO) (GenerateReportOutputDTO, error) {
	report, err := r.ReportRepository.Generate(input.StatusCodes, input.Duration)
	if err != nil {
		log.Println(err)
	}

	dto := GenerateReportOutputDTO{
		TimeSpent:          report.TimeSpent.String(),
		RequestsMade:       report.RequestsMade,
		SuccessfulRequests: report.SuccessfulRequests,
		FailedRequests:     report.FailedRequests,
	}

	return dto, nil
}
