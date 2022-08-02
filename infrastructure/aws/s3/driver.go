package s3

import (
	"bookstore/config"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func ConnectAws(cfg *config.AppConfig) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.AWS_REGION),
			Credentials: credentials.NewStaticCredentials(
				cfg.AWS_ACCESS_KEY_ID,
				cfg.AWS_SECRET_ACCESS_KEY,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func DoUpload(session *session.Session, file multipart.FileHeader, bucket string, destination string) (string, error) {
	//upload to the s3 bucket
	uploader := s3manager.NewUploader(session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(destination),
		Body:   body,
	})
	return up.Location, err
}
