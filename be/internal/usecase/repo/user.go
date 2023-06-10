package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if err := r.DB.Collection("users").FindOne(ctx, bson.M{"email": user.Email}).Err(); err == nil {
		return fmt.Errorf("user already exists")
	}

	user.ID = uuid.NewString()
	if user.SignupDate.IsZero() {
		user.SignupDate = time.Now()
	}

	_, err := r.DB.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("error inserting the user")
	}

	return nil
}

// GetByID retrieves a user by ID.
func (r *UserRepo) GetByID(ctx context.Context, ID string) (*entity.User, error) {
	user := &entity.User{}

	err := r.DB.Collection("users").FindOne(ctx, bson.M{"_id": ID}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
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
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *UserRepo) List(ctx context.Context, email string) ([]*entity.User, error) {
	filter := bson.M{}
	if email != "" {
		filter["email"] = primitive.Regex{Pattern: email, Options: "i"}
	}

	cur, err := r.DB.Collection("users").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	users := []*entity.User{}

	for cur.Next(ctx) {
		user := entity.User{}
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
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) EditShopsByIDs(ctx context.Context, user *entity.User, ownedShops []string) error {

	userType := entity.USER
	if len(ownedShops) > 0 {
		userType = entity.BARBER
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"ownedShops": ownedShops, "type": userType}}

	_, err := r.DB.Collection("users").UpdateOne(ctx, filter, update)
	return err
}

func (r *UserRepo) ModifyByID(ctx context.Context, ID string, user *entity.User) error {

	update := bson.M{}

	if user != nil {
		if user.Email != "" {
			update["email"] = user.Email
		}

		if user.Password != "" {
			update["password"] = user.Password
		}

		if user.Type != "" {
			update["type"] = user.Type
		}
	}

	res, err := r.DB.Collection("users").UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})

	if res.MatchedCount == 0 {
		return errors.New("barberShop not found")
	}
	return err
}
