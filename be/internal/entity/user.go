package entity

import "time"

const (
	ADMIN  = "admin"
	USER   = "user"
	BARBER = "barber"
)

type User struct {
	ID         string    `bson:"_id"`
	Email      string    `bson:"email"`
	Username   string    `bson:"username"`
	Password   string    `bson:"password" json:"-"`
	Type       string    `bson:"type"`
	SignupDate time.Time `bson:"signupDate"`

	CurrentAppointment *Appointment `bson:"currentAppointment,omitempty"`
	OwnedShops         []string     `bson:"ownedShops,omitempty"`
}
