package entity

type BarberShop struct {
	ID       string `bson:"_id"`
	Name     string
	Rating   float64
	Location *Location `json:"location" bson:"location"`

	Address     string
	Description string
	ImageLink   string
	Phone       string

	Employees int

	Views        []string
	Reviews      []*Review
	Appointments []*Appointment
}
