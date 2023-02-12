package entity

import "time"

type Appointment struct {
	ID        string `bson:"_id"`
	CreatedAt time.Time
	Start     time.Time

	UserID       string
	BarbershopID string
}
