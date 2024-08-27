package exception

type ICustomError interface {
	GetErrorCode() int
	GetStatusCode() int
	GetMessage() string
	GetDebugMessage() string
}

type customError struct {
	ErrorCode    ErrorCode
	StatusCode   int
	Message      string
	DebugMessage string
}

func NewCustomError(errorCode ErrorCode, debug ...string) ICustomError {
	var debugMessage string

	if len(debugMessage) > 0 {
		debugMessage = debug[0]
	}

	message, statusCode := errorCode.GetErrorDetail()

	return customError{ErrorCode: errorCode, StatusCode: statusCode, Message: message, DebugMessage: debugMessage}
}

func (error customError) GetErrorCode() int {
	return error.ErrorCode.Index()
}

func (error customError) GetStatusCode() int {
	return error.StatusCode
}

func (error customError) GetMessage() string {
	return error.Message
}

func (error customError) GetDebugMessage() string {
	return error.DebugMessage
}
