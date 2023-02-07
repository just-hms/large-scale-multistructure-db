package entity

import "time"

type Slot struct {
	Start                time.Time
	BookedAppoIntments   int
	UnavailableEmployees int
}
