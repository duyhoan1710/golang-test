package exception

import "net/http"

type ErrorCode int

const (
	INTERNAL_SERVER_ERROR ErrorCode = 0

	USER_NOT_FOUND      ErrorCode = 1
	USER_ALREADY_EXISTS ErrorCode = 2
	PASSWORD_INCORRECT  ErrorCode = 3

	ORDER_NOT_FOUND       ErrorCode = 4
	LAST_ORDER_NOT_FINISH ErrorCode = 5
	CANNOT_CANNEL_ORDER   ErrorCode = 6

	EXTRACT_TOKEN_ERROR = 7
	VALIDATION_ERROR    = 8
)

func (error ErrorCode) GetErrorDetail() (string, int) {
	switch error {
	case INTERNAL_SERVER_ERROR:
		return "Internal server error", http.StatusInternalServerError
	case USER_NOT_FOUND:
		return "User does not found", http.StatusNotFound
	case USER_ALREADY_EXISTS:
		return "User already exists", http.StatusBadRequest
	case PASSWORD_INCORRECT:
		return "Invalid credentials", http.StatusBadRequest
	case ORDER_NOT_FOUND:
		return "Order does not found", http.StatusBadRequest
	case LAST_ORDER_NOT_FINISH:
		return "Please finish previous first", http.StatusBadRequest
	case CANNOT_CANNEL_ORDER:
		return "Cannot cancel this order", http.StatusBadRequest
	case EXTRACT_TOKEN_ERROR:
		return "Extract refresh token error", http.StatusBadRequest
	case VALIDATION_ERROR:
		return "Please fill correct format form", http.StatusBadRequest
	}

	return "Internal server error", http.StatusInternalServerError
}

func (error ErrorCode) Index() int {
	return int(error)
}
