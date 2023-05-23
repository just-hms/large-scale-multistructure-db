package repo

import (
	"context"
	"errors"

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

	// Store appointment fields before removing unused fields in the BarberShop's Appointment
	userID := appointment.UserID
	shopID := appointment.BarbershopID
	shopName := appointment.BarbershopName
	appointment.BarbershopID = ""
	appointment.BarbershopName = ""

	// Add the Appointment to the Shop
	filter := bson.M{"_id": shopID}
	update := bson.M{"$push": bson.M{"appointments": appointment}}

	_, err = r.DB.Collection("barbershops").UpdateOne(ctx, filter, update)
	if err != nil {
		appointment.ID = ""
		return err
	}

	// Replace previously removed fields and remove unused field in the User's Appointment
	appointment.BarbershopID = shopID
	appointment.BarbershopName = shopName
	appointment.UserID = ""

	// Add the appointment to the User
	userFilter := bson.M{"_id": userID}
	userUpdate := bson.M{"$set": bson.M{"currentAppointment": appointment}}
	_, err = r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)
	if err != nil {
		return err
	}

	return err
}

func (r *AppointmentRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {
	if appointment == nil {
		return errors.New("you must provide an appointment")
	}
	filter := bson.M{"_id": appointment.BarbershopID}
	update := bson.M{"$pull": bson.M{"appointments": bson.M{"_id": appointment.ID}}}

	_, err := r.DB.Collection("barbershops").UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	// remove the appoinment from the user

	userFilter := bson.M{"_id": appointment.UserID}
	userUpdate := bson.M{"$unset": bson.M{"currentAppointment": ""}}
	_, err = r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)
	return err
}

func (r *AppointmentRepo) GetByIDs(ctx context.Context, shopID, ID string) (*entity.Appointment, error) {
	filter := bson.M{"_id": shopID, "appointments._id": ID}

	var barbershop entity.BarberShop
	err := r.DB.Collection("barbershops").FindOne(ctx, filter).Decode(&barbershop)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("appointment not found")
		}
		return nil, err
	}

	for _, appointment := range barbershop.Appointments {
		if appointment.ID == ID {
			return appointment, nil
		}
	}

	return nil, errors.New("appointment not found")
}
