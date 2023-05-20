package helpers

import "strings"

func MappingError(error string, section string) string {
	var ErrorMessage string
	switch section {
	case "auth":
		ErrorMessage = MappingErrorAuth(error)
	default:
		ErrorMessage = "internal server error"
	}

	return ErrorMessage
}

func MappingErrorAuth(error string) string {
	var ErrorMessage string
	if strings.Contains(error, "duplicate key value violates unique constraint \"user_email_key\"") {
		ErrorMessage = "email is already exist"
	} else if strings.Contains(error, "user not found") {
		ErrorMessage = "email is not found"
	}

	return ErrorMessage
}
