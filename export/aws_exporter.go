package export

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func UploadFile(fileName string, bucketName string, bucketKey string) {
	log.Printf("Uploading file %s into aws on bucket %s", fileName, bucketName)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	uploader := s3manager.NewUploader(sess)

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0662)
	defer f.Close()
	if err != nil {
		log.Panicln(err)
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketKey),
		Body:   f,
	})

	if err != nil {
		log.Panicln(err)
	}

	log.Printf("File uploaded to, %s\n", aws.StringValue(&result.Location))
}
