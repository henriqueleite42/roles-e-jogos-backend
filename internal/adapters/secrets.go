package adapters

import "crypto"

type Secrets struct {
	Port        string
	DatabaseUrl string

	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUri  string

	MediasS3BucketName string

	MediasCloudfrontKeyId      string
	MediasCloudfrontUrl        string
	MediasCloudfrontPrivateKey crypto.Signer

	EmailAddressSystemMessages string
	NameSystemMessages         string
	EmailTemplateSignInOtp     string
}
