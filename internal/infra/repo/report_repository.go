package repo

import (
	"time"

	"github.com/aleroxac/goexpert-stresstest/internal/entity"
)

type ReportRespository struct{}

func NewReportRepository() *ReportRespository {
	return &ReportRespository{}
}

func (r *ReportRespository) Generate(status_codes []int, duration time.Duration) (*entity.Report, error) {
	requests_made := len(status_codes)
	succesful_requests := 0
	failed_requests := map[string]int{}

	code_3xx := 0
	code_4xx := 0
	code_5xx := 0

	for _, code := range status_codes {
		if code == 200 {
			succesful_requests++
		} else {
			if code >= 300 && code <= 399 {
				code_3xx++
			}
			if code >= 400 && code <= 499 {
				code_4xx++
			}
			if code >= 500 && code <= 599 {
				code_5xx++
			}
		}
	}

	failed_requests["3xx"] = code_3xx
	failed_requests["4xx"] = code_4xx
	failed_requests["5xx"] = code_5xx
	report := entity.NewReport(duration, requests_made, succesful_requests, failed_requests)

	return report, nil
}
