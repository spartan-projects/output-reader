package export

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spartan-projects/output-reader/common"
	"log"
	"os"
)

func UploadFile(fileName string, bucketName string, bucketKey string) {
	log.Printf("Uploading file %s into AWS on bucket %s", fileName, bucketName)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(common.AwsRegionEnvKey)),
	}))
	uploader := s3manager.NewUploader(sess)

	f, err := os.OpenFile(fileName, os.O_RDONLY, common.FilePermissions)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketKey),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File uploaded to, %s\n", aws.StringValue(&result.Location))
}
