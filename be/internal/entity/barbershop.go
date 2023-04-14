package entity

type BarberShop struct {
	ID       string `bson:"_id"`
	Name     string
	Rating   float64
	Location Location `json:"location" bson:"location"`

	Address   string
	ImageLink string
	Phone     string

	Employees int

	Reviews      []*Review      `bson:",omitempty"`
	Appointments []*Appointment `bson:",omitempty"`
}
