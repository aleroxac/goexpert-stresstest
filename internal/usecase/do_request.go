package usecase

import (
	"log"
	"time"

	"github.com/aleroxac/goexpert-stresstest/internal/entity"
)

type DoRequestInputDTO struct {
	URL         string
	Requests    int
	Concurrency int
}

type DoRequestOutputDTO struct {
	StatusCodes []int
	TimeSpent   time.Duration
}

type DoRequestUseCase struct {
	RequestRepository entity.RequestRepositoryInterface
}

func NewDoRequestUseCase(request_repository entity.RequestRepositoryInterface) *DoRequestUseCase {
	return &DoRequestUseCase{
		RequestRepository: request_repository,
	}
}

func (r *DoRequestUseCase) Execute(input DoRequestInputDTO) (DoRequestOutputDTO, error) {
	staus_codes, time_spent, err := r.RequestRepository.Do(input.URL, input.Requests, input.Concurrency)
	if err != nil {
		log.Println(err)
	}

	dto := DoRequestOutputDTO{
		StatusCodes: staus_codes,
		TimeSpent:   time_spent,
	}

	return dto, nil
}
