package main

import (
	"flag"
	"github.com/cadena-monde/s3client/progress"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"log"
)

const (
	kb                     = 1 * 1000
	mb                     = kb * 1000
	minMultiPartUploadSize = 6 * mb
)

var (
	bucketName string
	filePath   string
)

func init() {
	flag.StringVar(&bucketName, "b", "", "S3 Bucket Name")
	flag.StringVar(&filePath, "f", "", "Path to the file to be uploaded")
}

func main() {
	flag.Parse()

	if filePath == "" || bucketName == "" {
		flag.Usage()
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatalf("Error reading authentication info: %s", err.Error())
	}

	s := s3.New(auth, aws.USEast)
	bucket := s.Bucket(bucketName)

	file := new(progress.ProgressFileReader)
	file.Open(filePath)
	defer file.Close()

	if canUploadAsMultipart(bucket, file) {
		uploadAsMultiPart(bucket, file)
	} else {
		upload(bucket, file)
	}
}

func canUploadAsMultipart(bucket *s3.Bucket, file *progress.ProgressFileReader) bool {
	return file.FileInfo.Size() > minMultiPartUploadSize
}

func upload(bucket *s3.Bucket, file *progress.ProgressFileReader) {
	err := bucket.PutReader(file.FileInfo.Name(), file, file.FileInfo.Size(), "", "")
	if err != nil {
		log.Fatalf("Error uploading to S3: %s", err)
	}
}

func uploadAsMultiPart(bucket *s3.Bucket, file *progress.ProgressFileReader) {
	multi, err := bucket.Multi(file.FileInfo.Name(), "", "")
	if err != nil {
		log.Fatalf("Error starting mutipart upload: %s", err)
	}

	parts, err := multi.PutAll(file, minMultiPartUploadSize)
	if err != nil {
		log.Fatalf("Error sending parts to S3: %s", err)
	}

	err = multi.Complete(parts)
	if err != nil {
		log.Fatalf("Error completing mutipart upload: %s", err)
	}
}
