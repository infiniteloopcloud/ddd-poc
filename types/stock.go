package types

type Product struct {
	ID string
}

func (p Product) Validate() error {
	// TODO implement me
	return nil
}

type Cart struct {
	ID       string
	Products []Product
}

func (p Cart) Validate() error {
	// TODO implement me
	return nil
}

type PayCart struct {
}

func (p PayCart) Validate() error {
	// TODO implement me
	return nil
}
