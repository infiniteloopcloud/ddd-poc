package types

import "time"

type User struct {
	ID            string    `json:"id"`
	AccountTypeID string    `json:"account_type_id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Merchant struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) Validate() error {
	return nil
}

func (m Merchant) Validate() error {
	return nil
}
