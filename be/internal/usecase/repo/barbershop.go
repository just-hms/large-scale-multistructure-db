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

type BarberShopRepo struct {
	*mongo.Mongo
}

func NewBarberShopRepo(m *mongo.Mongo) *BarberShopRepo {
	return &BarberShopRepo{m}
}

func (r *BarberShopRepo) Store(ctx context.Context, shop *entity.BarberShop) error {

	// cannot add a barber without a location if there is an index on it
	if shop.Location == nil {
		shop.Location = entity.FAKE_LOCATION
	}

	if err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"name": shop.Name}).Err(); err == nil {
		return fmt.Errorf("barber shop already exists")
	}

	shop.ID = uuid.NewString()
	_, err := r.DB.Collection("barbershops").InsertOne(ctx, shop)
	if err != nil {
		shop.ID = ""
		return fmt.Errorf("error inserting the barber shop: %s", err.Error())
	}
	return nil
}

func (r *BarberShopRepo) Find(ctx context.Context, lat float64, lon float64, name string, radius float64) ([]*entity.BarberShop, error) {

	filter := bson.D{}

	if radius != 0 {
		filter = append(
			filter,
			bson.E{
				Key: "location",
				Value: bson.D{
					{Key: "$near", Value: bson.D{
						{Key: "$geometry", Value: entity.NewLocation(lat, lon)},
						{Key: "$maxDistance", Value: radius},
					}},
				},
			},
		)
	}

	if name != "" {
		filter = append(
			filter,
			bson.E{Key: "name", Value: primitive.Regex{Pattern: name, Options: "i"}},
		)
	}

	cur, err := r.DB.Collection("barbershops").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	shops := []*entity.BarberShop{}

	for cur.Next(ctx) {
		var shop entity.BarberShop

		if err := cur.Decode(&shop); err != nil {
			return nil, err
		}
		shops = append(shops, &shop)
	}
	return shops, nil
}

func (r *BarberShopRepo) GetByID(ctx context.Context, ID string) (*entity.BarberShop, error) {

	barber := &entity.BarberShop{}

	err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": ID}).Decode(&barber)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return barber, nil
}

func (r *BarberShopRepo) GetOwnedShops(ctx context.Context, user *entity.User) ([]*entity.BarberShop, error) {

	if user.Type != entity.BARBER {
		return nil, fmt.Errorf("user is not a Barber")
	}

	filter := bson.M{"_id": bson.M{"$in": user.OwnedShops}}

	cur, err := r.DB.Collection("barbershops").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	shops := []*entity.BarberShop{}

	for cur.Next(ctx) {
		shop := entity.BarberShop{}
		if err := cur.Decode(&shop); err != nil {
			return nil, err
		}
		shops = append(shops, &shop)
	}
	return shops, nil
}

func (r *BarberShopRepo) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {

	update := bson.M{}

	if shop != nil {
		if shop.Location != nil {
			update["location"] = shop.Location
		}
		if shop.Name != "" {
			update["name"] = shop.Name
		}
		if shop.Description != "" {
			update["description"] = shop.Description
		}
		if shop.Employees != -1 {
			update["employees"] = shop.Employees
		}
	}

	res, err := r.DB.Collection("barbershops").UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": update})
	if res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return err
}

func (r *BarberShopRepo) DeleteByID(ctx context.Context, ID string) error {

	res, err := r.DB.Collection("barbershops").DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
