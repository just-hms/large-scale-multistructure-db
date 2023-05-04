package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
)

type VoteRepo struct {
	*mongo.Mongo
}

func NewVoteRepo(m *mongo.Mongo) *VoteRepo {
	return &VoteRepo{m}
}

func (v *VoteRepo) UpVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return nil
}

func (v *VoteRepo) DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return nil
}

func (v *VoteRepo) RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return nil
}
