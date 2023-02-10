package entity

import "time"

type ShopView struct {
	CreatedAt    time.Time
	ViewerID     string
	BarberShopID string
}
