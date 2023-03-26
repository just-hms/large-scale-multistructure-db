package repo

import (
	"context"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppointmentRepo struct {
	*mongo.Mongo
}

func NewAppointmentRepo(m *mongo.Mongo) *AppointmentRepo {
	return &AppointmentRepo{m}
}

func (r *AppointmentRepo) Book(ctx context.Context, appointment *entity.Appointment) error {

	filter := bson.M{"_id": appointment.BarbershopID}
	update := bson.M{"$push": bson.M{"appointments": appointment}}

	res, err := r.DB.Collection("barbershops").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount > 0 {
		appointmentID, _ := res.UpsertedID.(primitive.ObjectID)
		appointment.ID = appointmentID.Hex()
	}

	// add the appointment
	userFilter := bson.M{"_id": appointment.UserID}
	userUpdate := bson.M{"$set": bson.M{"currentAppointment": appointment}}
	_, err = r.DB.Collection("users").UpdateOne(ctx, userFilter, userUpdate)
	if err != nil {
		return err
	}

	return err
}

func (r *AppointmentRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {
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
