package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
	"os"
)

func main() {
	aws_access_key_id := "Insert Key ID here"
	aws_secret_access_key := "Insert Secret Here"
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

	_, err := creds.Get()
	if err != nil {
		log.Panic(err)
	}

	cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)

	svc := s3.New(session.New(), cfg)

	//r is an *http.Request object
	//file, h, err := r.FormFile(“data”) // or whatever the form key is
	//if err != nil {
	//	// handle error
	//}

	file, err := os.Open("test.jpeg")

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	size := fileInfo.Size()

	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Name()

	params := &s3.PutObjectInput{
		Bucket: aws.String("testBucket"),
		Key: aws.String(path),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}

	resp, err := svc.PutObject(params)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("response %s", awsutil.StringValue(resp))
}
