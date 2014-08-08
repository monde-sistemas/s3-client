s3client
========

Amazon s3 client using the goamz library.

## Features

### Resumable uploads

Uploads larger than 6MB are sent as multipart and can be resumed.

Uploads are resumed if the file being uploaded have the same name and the checksum matches with the parts already on S3.

if the checksum doesn't match the file will be overwritten on S3.

## Usage

    s3client.exe -b bucket_name -f file.zip -d s3/directory

### Flags

- **-b** S3 bucket name to upload to. Ex: my-s3-bucket
- **-f** Full or relative local path to the file to be uploaded. Ex: C:\Path\to\file.exe
- **-d** Directory to put the file on s3. Ex: my_dir (optional)

If a directory is specified using the -d flag it will be created on S3 if not exists.

## AWS Authentication

Authentication info is read from the enviroment: https://godoc.org/launchpad.net/goamz/aws#EnvAuth

Example:

    SET AWS_ACCESS_KEY_ID=My_key_Id
    SET AWS_SECRET_ACCESS_KEY=My_Secret_Id
