package adapters

type Secrets struct {
	Port        string
	DatabaseUrl string

	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUri  string

	MediasS3BucketName string

	MediasCloudfrontUrl string

	EmailAddressSystemMessages string
	NameSystemMessages         string
	EmailTemplateSignInOtp     string
}
