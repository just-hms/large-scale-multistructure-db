package entity

type BarberShop struct {
	ID     string `bson:"_id"`
	Name   string
	Rating float64

	Hours       []*Hour
	Location    string
	ImageLink   string
	Phone       string
	Coordinates Coordinates
	Employees   int
}
