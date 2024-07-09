package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	st, err := os.Stat(fromPath)
	if err != nil {
		return fmt.Errorf("failed to get stat %s with error %s", fromPath, err.Error())
	}

	if st.IsDir() {
		return ErrUnsupportedFile
	}

	if offset > st.Size() {
		return ErrOffsetExceedsFileSize
	}

	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s, %s", fromPath, err.Error())
	}
	defer fileFrom.Close()

	if offset > 0 {
		_, err := fileFrom.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("failed to seek file %s", err.Error())
		}
	}

	fileTo, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("failed to create temp file %s, %s", toPath+"/"+fileFrom.Name(), err.Error())
	}
	defer fileTo.Close()

	toRead := st.Size() - offset
	if limit > 0 && limit < toRead {
		toRead = limit
	}
	reader := io.LimitReader(fileFrom, toRead)
	bar := pb.Full.Start64(toRead)
	barReader := bar.NewProxyReader(reader)
	_, err = io.CopyN(fileTo, barReader, toRead)
	if err != nil {
		return fmt.Errorf("failed to copy %d bytes with error %s", toRead, err.Error())
	}
	bar.Finish()

	return nil
}
