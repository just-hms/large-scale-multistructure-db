package entity

import "go.mongodb.org/mongo-driver/bson"

type AdminAnalytics struct {
	AppointmentsByMonth                   map[string]int
	ViewsByMonth                          map[string]int
	ReviewsByMonth                        map[string]int
	NewUsersByMonth                       map[string]int
	AppointmentCancellationUserRanking    []bson.M
	GetAppointmentCancellationShopRanking []bson.M
	GetEngagementShopRanking              []bson.M
}
