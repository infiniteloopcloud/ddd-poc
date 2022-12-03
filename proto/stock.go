package proto

type Product struct {
	ID string `json:"id"`
}

func (p Product) Validate() error {
	// TODO implement me
	return nil
}

type Cart struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
}

func (p Cart) Validate() error {
	// TODO implement me
	return nil
}

type PayCart struct {
	Transaction Transaction `json:"transaction"`
}

func (p PayCart) Validate() error {
	// TODO implement me
	return nil
}
