package routes

import (
	"github.com/gorilla/mux"
	"controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Rota para listar todos os alunos
	router.HandleFunc("/alunos", ListAlunos).Methods("GET")

	// Rota para obter um aluno específico
	router.HandleFunc("/alunos/{id}", GetAluno).Methods("GET")

	// Rota para criar um novo aluno
	router.HandleFunc("/alunos", CreateAluno).Methods("POST")

	// Rota para atualizar um aluno
	router.HandleFunc("/alunos/{id}", UpdateAluno).Methods("PUT")

	// Rota para excluir um aluno
	router.HandleFunc("/alunos/{id}", DeleteAluno).Methods("DELETE")

	return router
}
