package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type AppointmentRepo struct {
	*mongo.Mongo
}

func NewAppointmentRepo(m *mongo.Mongo) *AppointmentRepo {
	return &AppointmentRepo{m}
}

func (r *AppointmentRepo) Book(ctx context.Context, appointment *entity.Appointment) error {

	var barbershop entity.BarberShop
	err := r.DB.Collection("barbershops").FindOne(ctx, bson.M{"_id": appointment.BarbershopID}).Decode(&barbershop)

	if err != nil {
		return errors.New("the specified shop does not exists")
	}

	appointment.ID = uuid.NewString()
	appointment.BarbershopName = barbershop.Name
	if appointment.CreatedAt.IsZero() {
		appointment.CreatedAt = time.Now()
	}
	appointment.Status = "pending"

	// Store appointment fields before removing unused fields in the BarberShop's Appointment
	userID := appointment.UserID
	shopName := appointment.BarbershopName
	appointment.BarbershopName = ""

	// Add the Appointment its collection
	_, err = r.DB.Collection("appointments").InsertOne(ctx, appointment)
	if err != nil {
		appointment.ID = ""
		return fmt.Errorf("error inserting the appointment: %s", err.Error())
	}

	// Replace previously removed fields and remove unused field in the User's Appointment
	appointment.BarbershopName = shopName
	appointment.UserID = ""
	appointment.Username = ""

	// Add the appointment to the User
	userFilter := bson.M{"_id": userID}
	userUpdate := bson.M{"$set": bson.M{"currentAppointment": appointment}}
	_, err = r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)
	if err != nil {
		return err
	}

	// Add back the UserID for debugging purposes
	appointment.UserID = userID

	return err
}

func (r *AppointmentRepo) SetStatusFromUser(ctx context.Context, userID string, appointment *entity.Appointment) error {
	if appointment == nil {
		return errors.New("you must provide an appointment")
	}

	// Remove the appointment from the user
	userFilter := bson.M{"_id": userID}
	userUpdate := bson.M{"$unset": bson.M{"currentAppointment": ""}}
	_, err := r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": appointment.ID}
	update := bson.M{"$set": bson.M{"status": appointment.Status}}

	_, err = r.DB.Collection("appointments").UpdateOne(ctx, filter, update)

	return err
}

func (r *AppointmentRepo) SetStatusFromShop(ctx context.Context, shopID string, appointment *entity.Appointment) error {
	if appointment == nil {
		return errors.New("you must provide an appointment")
	}

	userID := appointment.UserID
	// Remove the appointment from the shop
	filter := bson.M{"_id": appointment.ID}
	update := bson.M{"$set": bson.M{"status": appointment.Status}}

	_, err := r.DB.Collection("appointments").UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	// remove the appoinment from the user

	userFilter := bson.M{"_id": userID}
	userUpdate := bson.M{"$unset": bson.M{"currentAppointment": ""}}
	_, err = r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)

	// Add the BarberShopID field for the slot repo
	appointment.BarbershopID = shopID

	return err
}

func (r *AppointmentRepo) GetByID(ctx context.Context, ID string) (*entity.Appointment, error) {

	appointment := &entity.Appointment{}

	err := r.DB.Collection("appointments").FindOne(ctx, bson.M{"_id": ID}).Decode(&appointment)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("appointment not found")
		}
		return nil, err
	}
	return appointment, nil
}
