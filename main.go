package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))

	bucketName := "xxx-bucket"
	objectKey := "xxx-key"

	// S3 clientを作成
	svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 最初の10byteだけ読み込んで表示
	rc := obj.Body
	defer rc.Close()
	buf := make([]byte, 10)
	_, err = rc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)
}
