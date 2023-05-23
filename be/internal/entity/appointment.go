package entity

import "time"

type Appointment struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	StartDate time.Time `bson:"startDate"`

	UserID       string `bson:"userID"`
	BarbershopID string `bson:"userID"`
}
