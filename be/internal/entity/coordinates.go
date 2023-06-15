package entity

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

var FAKE_LOCATION = NewLocation(0, 0)

func NewLocation(lon, lat float64) *Location {
	return &Location{
		Type:        "Point",
		Coordinates: []float64{lon, lat},
	}
}
