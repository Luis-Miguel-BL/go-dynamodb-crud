package entities

import (
	"time"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/vo"
)

type EmailScore struct {
	Base
	Email          string    `json:"email"`
	Score          float64   `json:"score"`
	SentCount      int64     `json:"sent_count"`
	DeliveredCount int64     `json:"delivered_count"`
	OpenedCount    int64     `json:"opened_count"`
	Bounced        bool      `json:"bounced"`
	MxValidated    bool      `json:"mx_validated"`
	LastSent       time.Time `json:"last_sent"`
	LastDelivery   time.Time `json:"last_delivery"`
	LastOpened     time.Time `json:"last_opened"`
}

func (e *EmailScore) ConsolidateScore(data vo.ConsolidateEmailScore) (emailScore *EmailScore) {
	emailScore = e
	emailScore.SentCount++
	emailScore.LastSent = data.SentDate

	emailScore.MxValidated = data.MxValidated

	switch {
	case data.Bounced == true:
		emailScore.Bounced = true

	case data.Opened == true:
		emailScore.OpenedCount++
		emailScore.LastOpened = data.SentDate

	case data.Delivered == true:
		emailScore.DeliveredCount++
		emailScore.LastDelivery = data.SentDate
	}

	// if data.Bounced || emailScore.OpenedCount == 1 || emailScore.DeliveredCount == 1 {
	// 	emailScore.sumScore()
	// }
	emailScore.sumScore()

	return emailScore
}

func (entity *EmailScore) sumScore() {
	// switch {
	// case entity.Bounced == true:
	// 	entity.Score = -1
	// case entity.DeliveredCount > 0:
	// 	entity.Score = 50
	// case entity.OpenedCount > 0:
	// 	entity.Score = 100
	// default:
	// 	entity.Score = 0
	// }
	entity.Score = 100
}
