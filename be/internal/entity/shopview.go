package entity

import "time"

type ShopView struct {
	ID           string `bson:"_id"`
	CreatedAt    time.Time
	ViewerID     string
	BarberShopID string
}
