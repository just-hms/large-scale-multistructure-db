package entity

type BarberShop struct {
	ID     string `bson:"_id"`
	Name   string
	Rating float64

	Location  string
	ImageLink string
	Phone     string

	Employees int

	Latitude  string
	Longitude string

	Reviews      []*Review      `bson:",omitempty"`
	Appointments []*Appointment `bson:",omitempty"`
}
