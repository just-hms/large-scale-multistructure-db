package entity

const (
	ADMIN  = "admin"
	USER   = "user"
	BARBER = "barber"
)

type User struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Type     string `bson:"type"`

	CurrentAppointment *Appointment
	OwnedShops         []*BarberShop `bson:"ownedShops"`
}
