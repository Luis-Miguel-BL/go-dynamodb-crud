package viewmodel

import (
	"encoding/json"
	"net/http"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/vo"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateConsolidateScore(r *http.Request) (bodyParsed *vo.ConsolidateEmailScore, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bodyParsed)
	if err != nil {
		return bodyParsed, err
	}
	err = Validation.ValidateStruct(bodyParsed,
		Validation.Field(&bodyParsed.Email, Validation.Required, is.Email),
		Validation.Field(&bodyParsed.SentDate, Validation.Required),
	)

	if err != nil {
		return bodyParsed, err
	}
	return bodyParsed, err
}

func ValidateFindByEmails(r *http.Request) (emails []string, err error) {
	type tempStruct struct {
		Emails []string
	}
	bodyParsed := &tempStruct{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bodyParsed)

	return bodyParsed.Emails, err
}
