package merrors

type ErrorCodeType string

const (
	Success              ErrorCodeType = "00200"
	InternalServerError  ErrorCodeType = "00404"
	Unauthorized         ErrorCodeType = "00401"
	BadRequets           ErrorCodeType = "00400"
	AccountSuspended     ErrorCodeType = "00601"
	NotAllowed           ErrorCodeType = "00602"
	EmailAlreadyExist    ErrorCodeType = "00603"
	EmailNotExist        ErrorCodeType = "00604"
	IncorrectCredentials ErrorCodeType = "00605"
	DataNotExist         ErrorCodeType = "00606"
	OutOfStock           ErrorCodeType = "00606"
)

type Error struct {
	Message   string
	ErrorCode ErrorCodeType
	Error     error
}
