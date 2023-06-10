package entity

import "time"

type Appointment struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	StartDate time.Time `bson:"startDate"`
	Status    string    `bson:"status"`

	UserID         string `bson:"userId,omitempty"`
	Username       string `bson:"username,omitempty"`
	BarbershopID   string `bson:"shopId,omitempty"`
	BarbershopName string `bson:"shopName,omitempty"`
}
