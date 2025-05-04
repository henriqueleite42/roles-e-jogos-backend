package adapters

type SendManyInputContent struct {
	RecipientAddress string
	Placeholders     any // Must be convertible to JSON
}

type SendManyInput struct {
	TemplateId    string
	SenderName    string
	SenderAddress string
	Contents      []*SendManyInputContent
}

type Email interface {
	SendMany(i *SendManyInput) error
}
