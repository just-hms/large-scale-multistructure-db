package repo_test

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/usecase/repo"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"
)

func (s *RepoSuite) TestUpVote() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)
	voteRepo := repo.NewVoteRepo(s.db)

	user := &entity.User{Email: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	review := &entity.Review{
		Rating:  4,
		Content: "test",
		UserID:  user.ID,
	}

	err = reviewRepo.Store(context.Background(), review, shop.ID)
	s.Require().NoError(err)

	// check that the review was correctly created
	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews, 1)

	err = voteRepo.UpVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)

	// check that a user cannot upvote twice
	err = voteRepo.UpVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().Error(err, mongo.ErrNoDocuments)

	// check that the vote was correctly created
	// in the review
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].UpVotes, 1)

}

func (s *RepoSuite) TestDownVote() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)
	voteRepo := repo.NewVoteRepo(s.db)

	user := &entity.User{Email: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	review := &entity.Review{
		Rating:  4,
		Content: "test",
		UserID:  user.ID,
	}

	err = reviewRepo.Store(context.Background(), review, shop.ID)
	s.Require().NoError(err)

	err = voteRepo.DownVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)

	// check that a user cannot downvote twice
	err = voteRepo.DownVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().Error(err, mongo.ErrNoDocuments)

	// check that the vote was correctly created
	// in the review
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].DownVotes, 1)

	// check that an upvote correctly removes a downvote
	err = voteRepo.UpVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].UpVotes, 1)
	s.Require().Len(shop.Reviews[0].DownVotes, 0)

	// check that a downvote correctly removes an upvote
	err = voteRepo.DownVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].UpVotes, 0)
	s.Require().Len(shop.Reviews[0].DownVotes, 1)

}

func (s *RepoSuite) TestRemoveVote() {

	userRepo := repo.NewUserRepo(s.db)
	shopRepo := repo.NewBarberShopRepo(s.db)
	reviewRepo := repo.NewReviewRepo(s.db)
	voteRepo := repo.NewVoteRepo(s.db)

	user := &entity.User{Email: "giovanni"}
	shop := &entity.BarberShop{Name: "brownies"}

	err := userRepo.Store(context.Background(), user)
	s.Require().NoError(err)

	err = shopRepo.Store(context.Background(), shop)
	s.Require().NoError(err)

	review := &entity.Review{
		Rating:  4,
		Content: "test",
		UserID:  user.ID,
	}

	err = reviewRepo.Store(context.Background(), review, shop.ID)
	s.Require().NoError(err)

	// check that the review was correctly created
	// in the barbershop collection
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews, 1)

	// check that the upvote is correctly created
	err = voteRepo.UpVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].UpVotes, 1)

	// check that the upvote gets correctly removed
	err = voteRepo.RemoveVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].UpVotes, 0)

	// check that the downvote is correctly created
	err = voteRepo.DownVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].DownVotes, 1)

	// check that the downvote gets correctly removed
	err = voteRepo.RemoveVoteByID(context.Background(), user.ID, shop.ID, review.ID)
	s.Require().NoError(err)
	shop, err = shopRepo.GetByID(context.Background(), shop.ID)
	s.Require().NoError(err)
	s.Require().Len(shop.Reviews[0].DownVotes, 0)

}
