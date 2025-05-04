package adapters

type Id interface {
	GenId() (string, error)
}
