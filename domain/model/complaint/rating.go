package complaint

import "go-complaint/erros"

type Rating struct {
	rate    int
	comment string
}

/*
* <<Value Object>>
Comment is optional and can be empty,
Rate is required and must be between 0 and 5
*/
func NewRating(rate int, comment string) (Rating, error) {
	var r *Rating = &Rating{}
	err := r.setRate(rate)
	if err != nil {
		return *r, err
	}
	err = r.setComment(comment)
	if err != nil {
		return *r, err
	}
	return *r, nil
}

func (r *Rating) setRate(rate int) error {
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
	return nil
}

func (r *Rating) setComment(comment string) error {
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
	return nil
}

func (r Rating) Rate() int {
	return r.rate
}

func (r Rating) Comment() string {
	return r.comment
}
