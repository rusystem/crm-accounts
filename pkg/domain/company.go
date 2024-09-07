package domain

import "time"

type Company struct {
	ID         int64     `json:"id"`
	NameRu     string    `json:"name_ru"`
	NameEn     string    `json:"name_en"`
	Country    string    `json:"country"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Website    string    `json:"website"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsApproved bool      `json:"is_approved"`
	Timezone   string    `json:"timezone"`
}
