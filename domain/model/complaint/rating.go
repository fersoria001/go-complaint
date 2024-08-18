package complaint

import (
	"go-complaint/domain/model/recipient"
	"go-complaint/erros"
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	id             uuid.UUID
	sentToReviewBy recipient.Recipient
	ratedBy        recipient.Recipient
	rate           int
	comment        string
	createdAt      time.Time
	lastUpdate     time.Time
}

/*
* <<Value Object>>
Comment is optional and can be empty,
Rate is required and must be between 0 and 5
*/
func NewRating(
	id uuid.UUID,
	sentToReviewBy, ratedBy recipient.Recipient,
	rate int,
	comment string,
	createdAt, lastUpdate time.Time,
) *Rating {
	return &Rating{
		id:             id,
		sentToReviewBy: sentToReviewBy,
		ratedBy:        ratedBy,
		rate:           rate,
		comment:        comment,
		createdAt:      createdAt,
		lastUpdate:     lastUpdate,
	}
}

func (r *Rating) SetRate(rate int) error {
	if rate < 0 {
		return &erros.ValidationError{
			Expected: "more than 0",
		}
	}
	if rate > 5 {
		return &erros.ValidationError{

			Expected: "less than 5",
		}
	}
	r.rate = rate
	r.lastUpdate = time.Now()
	return nil
}

func (r *Rating) SetComment(comment string) error {
	var len = len(comment)
	if len > 250 {
		return &erros.InvalidLengthError{
			AttributeName: "comment",
			MinLength:     3,
			MaxLength:     250,
			CurrentLength: len,
		}
	}
	r.comment = comment
	r.lastUpdate = time.Now()
	return nil
}

func (r Rating) Id() uuid.UUID {
	return r.id
}

func (r Rating) Rate() int {
	return r.rate
}

func (r Rating) Comment() string {
	return r.comment
}

func (r Rating) SentToReviewBy() recipient.Recipient {
	return r.sentToReviewBy
}

func (r Rating) RatedBy() recipient.Recipient {
	return r.ratedBy
}

func (r Rating) CreatedAt() time.Time {
	return r.createdAt
}

func (r Rating) LastUpdate() time.Time {
	return r.lastUpdate
}
