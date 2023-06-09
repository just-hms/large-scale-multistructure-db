package entity

type BarberAnalytics struct {
	AppointmentsByMonth                 map[string]int
	ViewsByMonth                        map[string]int
	ReviewsByMonth                      map[string]int
	AppointmentCancellationRatioByMonth map[string]float64
	AppointmentViewRatioByMonth         map[string]float64
	UpDownVoteCountByMonth              map[string]map[string]int
	ReviewWeightedRating                float64
	InactiveUsersList                   []string
}
