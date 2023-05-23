package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
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

	// TODO: add an ordered index??
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

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

	slot, err := r.get(appointment.BarbershopID, appointment.StartDate)

	if err != nil {
		slot = &entity.Slot{
			Start:                appointment.StartDate,
			BookedAppointments:   1,
			UnavailableEmployees: 0,
		}
	} else {
		slot.BookedAppointments += 1
	}

	return r.set(appointment.BarbershopID, appointment.StartDate, slot)
}

func (r *SlotRepo) Get(ctx context.Context, shopID string, date time.Time) (*entity.Slot, error) {
	slot, err := r.get(shopID, date)
	if err != nil {
		return &entity.Slot{
			Start:                date,
			BookedAppointments:   0,
			UnavailableEmployees: 0,
		}, nil
	}
	return slot, nil
}

func (r *SlotRepo) Cancel(ctx context.Context, appointment *entity.Appointment) error {

	slot, err := r.get(appointment.BarbershopID, appointment.StartDate)

	if err != nil {
		return errors.New("the slot does not exists")
	}

	if slot.BookedAppointments > 0 {
		slot.BookedAppointments -= 1
	}

	return r.set(appointment.BarbershopID, appointment.StartDate, slot)
}

func (r *SlotRepo) SetHoliday(ctx context.Context, shopID string, date time.Time, unavailableEmployees int) error {

	slot, err := r.get(shopID, date)

	newSlot := &entity.Slot{
		Start:                date,
		UnavailableEmployees: unavailableEmployees,
	}

	if err == nil {
		newSlot.BookedAppointments = slot.BookedAppointments
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

	if slot.Start.Before(time.Now()) {
		return errors.New("cannot create a slot prior to now")
	}

	expTime := time.Until(slot.Start.Add(time.Hour * 24))
	content, err := json.Marshal(slot)

	if err != nil {
		return err
	}

	err = r.Client.Set(key, content, expTime).Err()

	return err
}

func key(barberShopID string, date time.Time) string {
	key := fmt.Sprintf("barbershop:%s:slots:%d", barberShopID, date.Unix())
	return key
}
