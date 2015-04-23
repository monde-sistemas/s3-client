// Package progress implementes ReaderAtSeeker interface: https://godoc.org/launchpad.net/goamz/s3#ReaderAtSeeker
// It is used to displays progress as the data are read from the file.
package progress

import (
	"github.com/monde-sistemas/pb"
	"log"
	"os"
)

type ProgressFileReader struct {
	file     *os.File
	FileInfo os.FileInfo
	bar      *pb.ProgressBar
}

func (pr ProgressFileReader) ReadAt(p []byte, off int64) (n int, err error) {
	pr.bar.Set64(off)
	return pr.file.ReadAt(p, off)
}

func (pr ProgressFileReader) Seek(offset int64, whence int) (int64, error) {
	return pr.file.Seek(offset, whence)
}

func (pr ProgressFileReader) Read(p []byte) (n int, err error) {
	n, err = pr.file.Read(p)
	pr.bar.Add(n)
	return n, err
}

func (pr *ProgressFileReader) Open(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %s: %s", filePath, err)
	}

	pr.file = f
	pr.FileInfo, err = f.Stat()
	if err != nil {
		f.Close()
		log.Fatalf("Error reading file info %s: %s", filePath, err)
	}
	pr.bar = pb.New64(pr.FileInfo.Size())
	pr.bar.SetUnits(pb.U_BYTES)
	pr.bar.Start()
}

func (pr ProgressFileReader) Close() {
	pr.bar.Finish()
	pr.file.Close()
}
