package utils

import (
	"fmt"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string) error {
	return &CustomError{Message: message}
}

func main() {
	err := NewCustomError("Ocorreu um erro ao processar os dados")

	if err != nil {
		fmt.Println("Erro:", err)
	}
}
