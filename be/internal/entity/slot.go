package entity

import "time"

type Slot struct {
	Start                time.Time
	BookedAppointments   int
	UnavailableEmployees int
}
