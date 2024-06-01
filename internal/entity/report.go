package entity

import (
	"time"
)

type Report struct {
	TimeSpent          time.Duration
	RequestsMade       int
	SuccessfulRequests int
	FailedRequests     map[string]int
}

func NewReport(time_spent time.Duration, requests_made, successful_requests int, failed_requests map[string]int) *Report {
	return &Report{
		TimeSpent:          time_spent,
		RequestsMade:       requests_made,
		SuccessfulRequests: successful_requests,
		FailedRequests:     failed_requests,
	}
}
