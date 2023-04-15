package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/just-hms/large-scale-multistructure-db/be/internal/entity"
	"github.com/just-hms/large-scale-multistructure-db/be/pkg/redis"
)

type SlotRepo struct {
	*redis.Redis
}

func NewSlotRepo(r *redis.Redis) *SlotRepo {
	return &SlotRepo{r}
}

func (r *SlotRepo) GetByBarberShopID(ctx context.Context, ID string) ([]*entity.Slot, error) {

	key := fmt.Sprintf("barbershop:%s:slots:*", ID)

	keys, err := r.Client.Keys(key).Result()

	if err != nil {
		return nil, err
	}

	slots := make([]*entity.Slot, 0, len(keys))

	for _, key := range keys {

		data, err := r.Client.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var slot entity.Slot
		if err := json.Unmarshal([]byte(data), &slot); err != nil {
			return nil, err
		}

		slots = append(slots, &slot)
	}

	return slots, nil
}

// add entry if not exists
func (r *SlotRepo) Book(ctx context.Context, appointment *entity.Appointment) error {

	if appointment.BarbershopID == "" {
		return errors.New("barberShopID not specified")
	}

	slot, err := r.get(appointment.BarbershopID, appointment.Start)

	if err != nil {
		slot = &entity.Slot{
			Start:                appointment.Start,
			BookedAppoIntments:   1,
			UnavailableEmployees: 0,
		}
	} else {
		slot.BookedAppoIntments += 1
	}

	return r.set(appointment.BarbershopID, appointment.Start, slot)
}

func (r *SlotRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {

	slot, err := r.get(appointment.BarbershopID, appointment.Start)

	if err != nil {
		return errors.New("the slot does not exists")
	}

	if slot.BookedAppoIntments > 0 {
		slot.BookedAppoIntments -= 1
	}

	return r.set(appointment.BarbershopID, appointment.Start, slot)
}

func (r *SlotRepo) SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {

	slot, err := r.get(shopID, date)

	newSlot := &entity.Slot{
		Start:                date,
		UnavailableEmployees: unavailableEmployees,
	}
	if err == nil {
		newSlot.BookedAppoIntments = slot.BookedAppoIntments
	}

	return r.set(shopID, date, newSlot)
}

func (r *SlotRepo) get(barberShopID string, date time.Time) (*entity.Slot, error) {

	key := key(barberShopID, date)

	data, err := r.Client.Get(key).Result()

	if err != nil {
		return nil, err
	}

	slot := &entity.Slot{}
	if err := json.Unmarshal([]byte(data), slot); err != nil {
		return nil, err
	}

	return slot, nil
}

func (r *SlotRepo) set(barberShopID string, date time.Time, slot *entity.Slot) error {

	key := key(barberShopID, date)

	// TODO fix expiration date based on the booked time

	if slot.Start.Before(time.Now()) {
		return errors.New("cannot create a slot prior to now")
	}

	// TODO fix this so it expires the next day, or at least all at the same time
	seconds := time.Until(slot.Start)
	bytes, err := json.Marshal(slot)

	if err != nil {
		return err
	}

	err = r.Client.Set(key, bytes, time.Second*seconds).Err()

	return err
}

func key(barberShopID string, date time.Time) string {
	key := fmt.Sprintf("barbershop:%s:slots:%d", barberShopID, date.UnixNano())
	return key
}
