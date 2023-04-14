package entity

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

var FAKE_LOCATION = NewLocation(0, 0)

func NewLocation(lat, lon float64) *Location {
	return &Location{
		Type:        "Point",
		Coordinates: []float64{lat, lon},
	}
}
