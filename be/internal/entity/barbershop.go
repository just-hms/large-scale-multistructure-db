package entity

type BarberShop struct {
	ID              string `bson:"_id"`
	Name            string
	Latitude        string
	Longitude       string
	EmployeesNumber int
	AverageRating   float64
}
