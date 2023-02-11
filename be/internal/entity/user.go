package entity

type User struct {
	ID       string `bson:"_id"`
	Email    string
	Password string
	IsAdmin  bool

	CurrentAppointment *Appointment
	BarberShopIDs      []string
}
