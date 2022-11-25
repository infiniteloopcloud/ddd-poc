package proto

import "time"

type Transaction struct {
	MerchantID string    `json:"merchant_id"`
	UserID     string    `json:"user_id"`
	GatewayID  string    `json:"gateway_id"`
	Amount     uint      `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (t Transaction) Validate() error {
	// TODO implement me
	return nil
}

type Gateway struct {
	ID string `json:"id"`
}

func (g Gateway) Validate() error {
	// TODO implement me
	return nil
}
