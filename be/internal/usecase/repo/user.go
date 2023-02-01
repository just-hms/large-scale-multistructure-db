package repo

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// UserRepo represents a repository for User entities.
type UserRepo struct {
	*mongo.Mongo
}

// NewUserRepo creates a new UserRepo.
func NewUserRepo(m *mongo.Mongo) *UserRepo {
	return &UserRepo{m}
}

// Store inserts a new user into the repository.
func (r *UserRepo) Store(ctx context.Context, user *entity.User) error {
	// TODO: check conversion
	_, err := r.DB.Collection("users").InsertOne(ctx, user)
	return err
}

// GetByID retrieves a user by ID.
func (r *UserRepo) GetByID(ctx context.Context, ID uint) (*entity.User, error) {
	user := &entity.User{}
	err := r.DB.Collection("users").FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}
	return user, nil
}

// GetByEmail retrieves a user by email.
func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.DB.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}
	return user, nil
}
