package usecase

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
)

type ReviewUseCase struct {
	reviewRepo ReviewRepo
	voteRepo   VoteRepo
}

// New -.
func NewReviewUseCase(r ReviewRepo, v VoteRepo) *ReviewUseCase {
	return &ReviewUseCase{
		reviewRepo: r,
		voteRepo:   v,
	}
}

func (uc *ReviewUseCase) Store(ctx context.Context, review *entity.Review, shopID string) error {
	return uc.reviewRepo.Store(ctx, review, shopID)
}

func (uc *ReviewUseCase) GetByBarberShopID(ctx context.Context, shopID string) ([]*entity.Review, error) {
	return uc.reviewRepo.GetByBarberShopID(ctx, shopID)
}

func (uc *ReviewUseCase) DeleteByID(ctx context.Context, reviewID string) error {
	return uc.reviewRepo.DeleteByID(ctx, reviewID)
}

func (uc *ReviewUseCase) UpVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return uc.voteRepo.UpVoteByID(ctx, userID, shopID, reviewID)
}

func (uc *ReviewUseCase) DownVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return uc.voteRepo.DownVoteByID(ctx, userID, shopID, reviewID)
}

func (uc *ReviewUseCase) RemoveVoteByID(ctx context.Context, userID, shopID, reviewID string) error {
	return uc.voteRepo.RemoveVoteByID(ctx, userID, shopID, reviewID)
}
