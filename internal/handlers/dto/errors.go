package dto

type ErrorDTO struct {
	Message string `json:"message"`
}

func MakeErrorDTO(message string) ErrorDTO {
	return ErrorDTO{Message: message}
}
