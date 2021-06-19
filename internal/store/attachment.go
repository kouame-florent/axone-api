package store

import (
	"bytes"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
)

type AttachentStore struct {
	Folder   string
	fileName uuid.UUID
}

func NewAttachmentStore(folder string, fileName uuid.UUID) *AttachentStore {

	return &AttachentStore{
		Folder:   folder,
		fileName: fileName,
	}

}

func (s *AttachentStore) StoreFile(data bytes.Buffer) error {
	log.Printf("Receive buffer with len: %d", data.Len())
	filePath := path.Join(s.Folder, s.fileName.String())
	err := os.WriteFile(filePath, data.Bytes(), 0755)
	if err != nil {
		return err
	}
	return nil
}
