s3client
========

Amazon s3 client using the goamz library.

## Features

### Resumable uploads

Uploads larger than 6MB are sent as multipart and can be resumed.

Uploads are resumed if the file being uploaded have the same name and the checksum matches with the parts already on S3.

if the checksum doesn't match the file will be overwritten on S3.

## Usage

    s3client.exe -b bucket_name -f file.zip

## AWS Authentication

Authentication info is read from the enviroment: https://godoc.org/launchpad.net/goamz/aws#EnvAuth

Example:

    SET AWS_ACCESS_KEY_ID=My_key_Id
    SET AWS_SECRET_ACCESS_KEY=My_Secret_Id
