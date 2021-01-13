package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Config aws.Config
var getObjectInput s3.GetObjectInput

func init() {
	s3Config = aws.Config{Region: aws.String("ap-northeast-1")}

	b := flag.String("b", "xxx-bucket", "bucket Name")
	k := flag.String("k", "xxx-key", "object key")
	flag.Parse()

	getObjectInput = s3.GetObjectInput{
		Bucket: aws.String(*b),
		Key:    aws.String(*k),
	}
}

func main() {
	// sessionの作成
	sess := session.Must(session.NewSession(&s3Config))

	// S3 clientを作成
	svc := s3.New(sess)

	obj, err := svc.GetObject(&getObjectInput)
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
