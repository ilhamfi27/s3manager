// Package s3manager allows to interact with an S3 compatible storage.
package s3manager

type Configuration struct {
	Endpoint            string
	UseIam              bool
	IamEndpoint         string
	AccessKeyID         string
	SecretAccessKey     string
	Bucket              string
	Region              string
	AllowDelete         bool
	AllowCreateBucket   bool
	ForceDownload       bool
	UseSSL              bool
	SkipSSLVerification bool
	SignatureType       string
	ListRecursive       bool
	Port                string
	Timeout             int32
	SseType             string
	SseKey              string
}
