package repo

import (
	"context"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/mongo"

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

func (r *BarberShopRepo) Find(ctx context.Context, lat string, lon string, name string, radius string) ([]*entity.BarberShop, error) {

	filter := bson.M{"latitude": bson.M{"$gt": lat, "$lt": lat + radius}}

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

func (r *BarberShopRepo) Store(ctx context.Context, shop *entity.BarberShop) (string, error) {

	shop.ID = uuid.NewString()

	if err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"name": shop.Name}).Err(); err == nil {
		return "", fmt.Errorf("Barber shop already exists")
	}

	res, err := r.DB.Collection("barbershops").InsertOne(ctx, shop)
	if err != nil {
		return "", fmt.Errorf("Error inserting the user")
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("Error retriving the userID")
	}

	return oid.Hex(), err
}

func (r *BarberShopRepo) GetByID(ctx context.Context, ID string) (*entity.BarberShop, error) {
	barber := &entity.BarberShop{}

	oid, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, fmt.Errorf("Error converting the ID")
	}

	err = r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": oid}).Decode(&barber)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}
	return barber, nil
}

func (r *BarberShopRepo) ModifyByID(ctx context.Context, ID string, shop *entity.BarberShop) error {
	oid, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return fmt.Errorf("Error converting the ID")
	}

	_, err = r.DB.Collection("users").UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": shop})
	if err != nil {
		return err
	}
	return nil
}
func (r *BarberShopRepo) DeleteByID(ctx context.Context, ID string) error {
	oid, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return fmt.Errorf("Error converting the ID")
	}

	res, err := r.DB.Collection("barbershops").DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("User not found")
	}
	return nil
}
