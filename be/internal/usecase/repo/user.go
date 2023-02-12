package repo

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"

	"github.com/google/uuid"
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

	user.ID = uuid.NewString()

	if err := r.DB.Collection("users").FindOne(ctx, bson.M{"email": user.Email}).Err(); err == nil {
		return fmt.Errorf("User already exists")
	}

	_, err := r.DB.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("Error inserting the user")
	}

	return nil
}

// GetByID retrieves a user by ID.
func (r *UserRepo) GetByID(ctx context.Context, ID string) (*entity.User, error) {
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

func (r *UserRepo) DeleteByID(ctx context.Context, ID string) error {

	res, err := r.DB.Collection("users").DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("User not found")
	}
	return nil
}

func (r *UserRepo) ModifyByID(ctx context.Context, ID string, user *entity.User) error {

	// TODO : partial update
	// TODO : barbershop list

	update := bson.M{}

	if user.Email != "" {
		update["email"] = user.Email
	}

	if user.Password != "" {
		update["password"] = user.Password
	}

	_, err := r.DB.Collection("users").UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	return err
}

// TODO : improve this
func (r *UserRepo) List(ctx context.Context, email string) ([]*entity.User, error) {
	filter := bson.M{}
	if email != "" {
		filter["email"] = email
	}

	cur, err := r.DB.Collection("users").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var users []*entity.User

	for cur.Next(ctx) {
		var user entity.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
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
