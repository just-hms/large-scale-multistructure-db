package entity

import "time"

type Appointment struct {
	CreatedAt time.Time
	Start     time.Time
	Duration  time.Duration
}
