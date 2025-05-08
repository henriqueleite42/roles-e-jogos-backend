package adapters

type Secrets struct {
	Port string

	WebsiteUrl string

	DatabaseUrl      string
	DatabaseUsername string
	DatabasePassword string

	LudopediaClientId     string
	LudopediaClientSecret string
	LudopediaRedirectUri  string

	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUri  string

	MediasS3BucketName string

	MediasCloudfrontUrl string

	EmailAddressSystemMessages string
	NameSystemMessages         string
	EmailTemplateSignInOtp     string
}
