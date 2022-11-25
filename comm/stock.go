package comm

type ProductDescriptor interface {
	Create()
	Update()
	Delete()
}

type CartDescriptor interface {
	Create()
	Update()
	Delete()
	Pay()
}
