package repositories

import (
	"errors"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/entities"
)

var (
	ErrEmailScoreNotFound = errors.New("the email score was not found")
)

type EmailScoreRepository interface {
	GetByEmail(email string) (emailScore entities.EmailScore, err error)
	ListByEmails(emails []string) (emailsScore []entities.EmailScore, err error)
	Create(entity *entities.EmailScore) (id string, err error)
	Update(entity *entities.EmailScore) error
}
