package infrastructure

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var encryptionService *EncryptionService
var encryptionServiceOnce sync.Once

func EncryptionServiceInstance() *EncryptionService {
	encryptionServiceOnce.Do(func() {
		encryptionService = NewEncryptionService()
	})
	return encryptionService
}

type EncryptionService struct {
}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (p *EncryptionService) Encrypt(text string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(text), 14)
}

func (p *EncryptionService) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
