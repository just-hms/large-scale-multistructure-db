package repo

import (
	"context"
	"errors"
	"fmt"

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
		return fmt.Errorf("User already exists")
	}
	user.ID = uuid.NewString()

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
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}
	return user, nil
}

// TODO: set the user type based on the number of shops
func (r *UserRepo) EditShopsByIDs(ctx context.Context, user *entity.User, IDs []string) error {

	ownedShops := []*entity.BarberShop{}

	for _, barbershopID := range IDs {

		barbershop := &entity.BarberShop{}
		err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": barbershopID}).Decode(&barbershop)

		if err != nil {
			return err
		}

		ownedShops = append(
			ownedShops,
			&entity.BarberShop{
				ID:   barbershopID,
				Name: barbershop.Name,
			},
		)
	}

	userType := entity.USER
	if len(ownedShops) > 0 {
		userType = entity.BARBER
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"ownedShops": ownedShops, "type": userType}}

	_, err := r.DB.Collection("users").UpdateOne(ctx, filter, update)

	return err
}

// TODO : make the user an admin
// TODO : block an user

func (r *UserRepo) ModifyByID(ctx context.Context, ID string, user *entity.User) error {
	if user == nil {
		return errors.New("user does not exists")
	}
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
