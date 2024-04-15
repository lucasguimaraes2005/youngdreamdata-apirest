package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Aluno representa um aluno na nossa aplicação
type Aluno struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
}

var alunos []Aluno

// ListAlunos retorna todos os alunos
func ListAlunos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

// GetAluno retorna um aluno específico pelo ID
func GetAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, aluno := range alunos {
		if aluno.ID == params["id"] {
			json.NewEncoder(w).Encode(aluno)
			return
		}
	}
	http.NotFound(w, r)
}

// CreateAluno cria um novo aluno
func CreateAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var aluno Aluno
	_ = json.NewDecoder(r.Body).Decode(&aluno)
	alunos = append(alunos, aluno)
	json.NewEncoder(w).Encode(aluno)
}

// UpdateAluno atualiza os dados de um aluno
func UpdateAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, aluno := range alunos {
		if aluno.ID == params["id"] {
			alunos[i].Nome = "Novo Nome" // Atualize os campos conforme necessário
			json.NewEncoder(w).Encode(alunos[i])
			return
		}
	}
	http.NotFound(w, r)
}

// DeleteAluno exclui um aluno
func DeleteAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, aluno := range alunos {
		if aluno.ID == params["id"] {
			alunos = append(alunos[:i], alunos[i+1:]...)
			fmt.Fprintf(w, "Aluno com ID %s foi excluído", params["id"])
			return
		}
	}
	http.NotFound(w, r)
}
