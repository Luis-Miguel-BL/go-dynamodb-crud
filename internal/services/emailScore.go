package services

import (
	"time"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/entities"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/repositories"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/vo"
	"github.com/google/uuid"
)

type EmailScoreService struct {
	repository repositories.EmailScoreRepository
}

type EmailScoreInterface interface {
	GetByEmail(email string) (entity entities.EmailScore, err error)
	ListByEmails(emails []string) (entities []entities.EmailScore, err error)
	Create(entity *entities.EmailScore) (string, error)
	ConsolidateScore(consolidateScore *vo.ConsolidateEmailScore) error
}

func NewEmailScoreService(repository repositories.EmailScoreRepository) EmailScoreInterface {
	return &EmailScoreService{repository: repository}
}

func (e *EmailScoreService) GetByEmail(email string) (emailScore entities.EmailScore, err error) {
	return e.repository.GetByEmail(email)
}

func (e *EmailScoreService) ListByEmails(emails []string) (emailsScore []entities.EmailScore, err error) {

	return e.repository.ListByEmails(emails)

}

func (e *EmailScoreService) Create(entity *entities.EmailScore) (string, error) {
	setDefaultValues(entity)
	return e.repository.Create(entity)
}

func (e *EmailScoreService) ConsolidateScore(consolidateScore *vo.ConsolidateEmailScore) error {

	oldEmailScore, err := e.GetByEmail(consolidateScore.Email)
	if err != nil && oldEmailScore.ID != "" {
		return err
	}

	newEmailScore := &entities.EmailScore{Email: consolidateScore.Email,
		Score:          oldEmailScore.Score,
		SentCount:      oldEmailScore.SentCount,
		DeliveredCount: oldEmailScore.DeliveredCount,
		OpenedCount:    oldEmailScore.OpenedCount,
		Bounced:        oldEmailScore.Bounced,
		MxValidated:    oldEmailScore.MxValidated,
		LastSent:       oldEmailScore.LastSent,
		LastDelivery:   oldEmailScore.LastDelivery,
		LastOpened:     oldEmailScore.LastOpened,
	}
	setDefaultValues(newEmailScore)

	newEmailScore.ConsolidateScore(*consolidateScore)

	err = e.repository.Update(newEmailScore)

	return err
}

func setDefaultValues(emailScore *entities.EmailScore) {
	emailScore.UpdatedAt = time.Now()
	if emailScore.ID == "" {
		emailScore.ID = uuid.New().String()
		emailScore.CreatedAt = time.Now()
	}
}
