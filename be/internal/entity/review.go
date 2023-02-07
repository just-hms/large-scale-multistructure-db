package entity

import "time"

type Review struct {
	Content   string
	CreatedAt time.Time
	Rating    int
	Reported  bool
}
