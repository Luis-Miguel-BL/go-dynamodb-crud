package vo

import "time"

type ConsolidateEmailScore struct {
	Email       string    `json:"email"`
	MxValidated bool      `json:"mx_validated"`
	Bounced     bool      `json:"bounced"`
	Delivered   bool      `json:"delivered"`
	Opened      bool      `json:"opened"`
	SentDate    time.Time `json:"sent_date"`
}
