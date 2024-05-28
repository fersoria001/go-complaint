package infrastructure

import "golang.org/x/crypto/bcrypt"

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
