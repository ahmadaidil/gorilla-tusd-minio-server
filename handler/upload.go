package handler

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tus/tusd"
	"github.com/tus/tusd/cmd/tusd/cli"
	"github.com/tus/tusd/s3store"
)

var (
	accessKey = "MINIO_ACCESS"
	secretKey = "MINIO_SECRET"
	endpoint  = "http://localhost:9000"
	bucket    = "my-bucket"
	region    = "eu-west-1"
)

// FileUpload .
func FileUpload() *tusd.Handler {
	if envAccessKey, ok := os.LookupEnv("MINIO_ACCESS_KEY"); ok {
		accessKey = envAccessKey
	}
	if envSecretKey, ok := os.LookupEnv("MINIO_SECRET_KEY"); ok {
		secretKey = envSecretKey
	}
	if envEndpoint, ok := os.LookupEnv("MINIO_ENDPOINT"); ok {
		endpoint = envEndpoint
	}
	if envBucket, ok := os.LookupEnv("MINIO_BUCKET"); ok {
		bucket = envBucket
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))

	s3Config := aws.NewConfig().
		WithEndpoint(endpoint).
		WithS3ForcePathStyle(true)

	svc := s3.New(sess, s3Config)

	store := s3store.New(bucket, svc)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	cli.SetupPreHooks(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/upload/",
		StoreComposer:           composer,
		NotifyCompleteUploads:   true,
		NotifyTerminatedUploads: true,
		NotifyUploadProgress:    true,
		NotifyCreatedUploads:    true,
	})
	if err != nil {
		log.Printf("unable to create handler: %s", err)
	}

	eventHandler(handler)

	cli.SetupPostHooks(handler)

	return handler
}

func eventHandler(handler *tusd.Handler) {
	go func() {
		for {
			select {
			case info := <-handler.CompleteUploads:
				log.Println("file info --> ", info)
			}
		}
	}()
}
