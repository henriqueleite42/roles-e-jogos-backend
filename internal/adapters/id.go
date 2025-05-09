package adapters

type Id interface {
	GenId() string
	GenReqId() string
	GenSessionId() string
}
