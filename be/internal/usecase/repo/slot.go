package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"large-scale-multistructure-db/be/internal/entity"
	"large-scale-multistructure-db/be/pkg/redis"
)

type SlotRepo struct {
	*redis.Redis
}

func NewSlotRepo(r *redis.Redis) *SlotRepo {
	return &SlotRepo{r}
}

func (r *SlotRepo) GetByBarberShopID(ctx context.Context, ID string) ([]*entity.Slot, error) {
	key := fmt.Sprintf("barbershop:%s:slots", ID)

	slotIDs, err := r.Client.SMembers(key).Result()
	if err != nil {
		return nil, err
	}

	slots := make([]*entity.Slot, 0, len(slotIDs))

	for _, slotID := range slotIDs {
		data, err := r.Client.HGetAll(slotID).Result()
		if err != nil {
			return nil, err
		}

		var slot entity.Slot
		if err := json.Unmarshal([]byte(data["slot"]), &slot); err != nil {
			return nil, err
		}

		slots = append(slots, &slot)
	}

	return slots, nil
}
