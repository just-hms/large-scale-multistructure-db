package entity

import "time"

type ShopView struct {
	ID           string    `bson:"_id"`
	CreatedAt    time.Time `bson:"createdAt"`
	UserID       string    `bson:"userId"`
	BarbershopID string    `bson:"shopId"`
}
