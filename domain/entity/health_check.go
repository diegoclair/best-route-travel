package entity

import "time"

type HealthCheck struct {
	entity
	UUID string
	Exam
	Doctor
	Institution
	Date time.Time
	Note string
}
