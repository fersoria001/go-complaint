package application

type LoginConfirmation struct {
	JWTToken
	isConfirmed bool
	retry       int
	email       string
}

func NewLoginConfirmation(email string, token JWTToken, isConfirmed bool) *LoginConfirmation {
	lc := &LoginConfirmation{
		JWTToken:    token,
		isConfirmed: isConfirmed,
		retry:       3,
		email:       email,
	}
	return lc
}

func (lc *LoginConfirmation) Confirm() {
	lc.isConfirmed = true
}

func (lc *LoginConfirmation) RetryConfirmation() error {
	if lc.retry <= 0 {
		return ErrConfirmationRetryLimit
	} else {
		lc.retry--
	}
	return nil
}

func (lc LoginConfirmation) IsConfirmed() bool {
	return lc.isConfirmed
}

func (lc LoginConfirmation) Email() string {
	return lc.email
}
