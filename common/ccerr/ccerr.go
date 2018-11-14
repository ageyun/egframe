package ccerr

import "fmt"

type CCErr struct {
	Code    int
	Message string
}

func (err CCErr) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(CCErr *CCErr, err error) *Err {
	return &Err{Code: CCErr.Code, Message: CCErr.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *CCErr:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}
