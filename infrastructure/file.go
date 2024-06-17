package infrastructure

import (
	"bytes"
	"fmt"
	projectpath "go-complaint/project_path"
	"io"
	"log"
	"os"
	"path/filepath"
)

const SIZE_LIMIT = 1024 * 1024 * 4

type FileContent struct {
	fileName string
	fileType string
	file     []byte
	path     string
}

func NewFileContent(name, fileType string, slice []byte) (FileContent, error) {
	if len(slice) > SIZE_LIMIT {
		return FileContent{}, fmt.Errorf("file too big")
	}
	fc := FileContent{
		fileName: name,
		file:     slice,
	}
	fc.fileType = filepath.Ext(fc.fileName)[1:]
	fc.path = filepath.Join(projectpath.Root, "files", fc.fileType, fc.fileName)
	return fc, nil
}
func (fc FileContent) Save() error {
	f, err := os.Open(fc.path)
	if err == nil {
		f.Close()
		return ErrFileAlreadyExists
	}
	f, err = os.Create(fc.path)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bytes.NewBuffer(fc.file)
	log.Printf("file size %d", len(fc.file))
	written, err := io.Copy(f, buf)
	if err != nil {
		return err
	}
	log.Printf("written %d", written)
	if written != int64(len(fc.file)) {
		return ErrFileTooBig
	}
	return nil
}
