package entity

import "time"

type entity struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
