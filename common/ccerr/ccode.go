package ccerr

// Common errors
var (
	OK                  = &CCErr{Code: 0, Message: "OK"}
	InternalServerError = &CCErr{Code: 10001, Message: "Internal server error"}
)
