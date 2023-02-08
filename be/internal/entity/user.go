package entity

type User struct {
	ID       string
	Email    string
	Password string
	IsAdmin  bool

	CurrentAppointment *Appointment
	BarberShopIDs      []string
}
