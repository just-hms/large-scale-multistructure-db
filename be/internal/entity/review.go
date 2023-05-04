package entity

import "time"

type Review struct {
	ReviewID  string    `bson:"reviewId"`
	UserID    string    `bson:"userId"`
	Username  string    `bson:"username"`
	Rating    int       `bson:"rating"`
	Reported  bool      `bson:"reported"`
	Content   string    `bson:"content"`
	UpVotes   []string  `bson:"upvotes"`
	DownVotes []string  `bson:"downvotes"`
	CreatedAt time.Time `bson:"createdAt"`
}
