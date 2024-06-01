package entity

import (
	"time"
)

type RequestRepositoryInterface interface {
	Do(string, int, int) ([]int, time.Duration, error)
	// Do(string, int, int, float64) ([]int, time.Duration, error)
}

type ReportRepositoryInterface interface {
	Generate([]int, time.Duration) (*Report, error)
}
