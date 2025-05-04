package adapters

type AddressData struct {
	PostalCode string
	Street     string
	Block      string
	ExtraData  *string
	City       string
	State      string
	Country    string
}

type Address interface {
	GetAddressFromCep(cep string) (*AddressData, error)
}
