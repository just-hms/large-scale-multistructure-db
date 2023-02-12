package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/redis"
	"time"
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

	slot, err := r.getByDateAndShop(appointment.BarbershopID, appointment.Start)

	if err != nil {
		slot = &entity.Slot{
			Start:                appointment.Start,
			BookedAppoIntments:   1,
			UnavailableEmployees: 0,
		}
	} else {
		slot.BookedAppoIntments += 1
	}

	return r.setByDateAndShop(appointment.BarbershopID, appointment.Start, slot)
}

func (r *SlotRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {

	slot, err := r.getByDateAndShop(appointment.BarbershopID, appointment.Start)

	// TODO i don't like this
	if err != nil {
		return nil
	}

	if slot.BookedAppoIntments > 0 {
		slot.BookedAppoIntments -= 1
	}

	return r.setByDateAndShop(appointment.BarbershopID, appointment.Start, slot)
}

func (r *SlotRepo) SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {

	slot, err := r.getByDateAndShop(shopID, date)

	newSlot := &entity.Slot{
		Start:                date,
		UnavailableEmployees: unavailableEmployees,
	}
	if err != nil {
		newSlot.BookedAppoIntments = slot.BookedAppoIntments
	}

	return r.setByDateAndShop(shopID, date, newSlot)
}

func (r *SlotRepo) getByDateAndShop(barberShopID string, date time.Time) (*entity.Slot, error) {

	key := fmt.Sprintf("barbershop:%s:slots:%s", barberShopID, date)

	data, err := r.Client.Get(key).Result()

	if err != nil {
		return nil, err
	}

	slot := &entity.Slot{}
	if err := json.Unmarshal([]byte(data), &slot); err != nil {
		return nil, err
	}

	return slot, nil
}

func (r *SlotRepo) setByDateAndShop(barberShopID string, date time.Time, slot *entity.Slot) error {

	key := fmt.Sprintf("barbershop:%s:slots:%s", barberShopID, date)

	// TODO fix expiration date based on the booked time

	if slot.Start.Before(time.Now()) {
		return fmt.Errorf("cannot create a slot prior to now")
	}

	// TODO fix this so it expires the next day, or at least all at the same time
	seconds := slot.Start.Sub(time.Now())
	err := r.Client.Set(key, slot, time.Second*seconds).Err()

	return err
}
