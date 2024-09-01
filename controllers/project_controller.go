package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// lista todos os projetos.
func GetProjetos(w http.ResponseWriter, r *http.Request) {
	var projetos []models.Projeto
	var projetosDTO []project_dto.ListagemProjetoDTO

	db.DB.Preload("Usuario").Find(&projetos)

	for _, projeto := range projetos {
		projetoDTO := project_dto.ListagemProjetoDTO{
			ID:        projeto.ID,
			Nome:      projeto.Nome,
			Descricao: projeto.Descricao,
			Status:    projeto.Status,
			Data:      projeto.DataCriacao,
			Usuario: user_dto.ListagemBasicaUsuarioDTO{
				ID:   projeto.Usuario.ID,
				Nome: projeto.Usuario.Nome,
			},
		}
		projetosDTO = append(projetosDTO, projetoDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projetosDTO)
}

// retorna um projeto específico por id.
func GetProjetoId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var projetos []models.Projeto
	var projetosDTO []project_dto.ListagemProjetoDTO

	err := db.DB.Preload("Usuario").First(&projetos, params["id"]).Error
	if err != nil {
		http.Error(w, "Projeto não encontrado", http.StatusNotFound)
		return
	}

	for _, projeto := range projetos {
		projetoDTO := project_dto.ListagemProjetoDTO{
			ID:        projeto.ID,
			Nome:      projeto.Nome,
			Descricao: projeto.Descricao,
			Status:    projeto.Status,
			Data:      projeto.DataCriacao,
			Usuario: user_dto.ListagemBasicaUsuarioDTO{
				ID:   projeto.Usuario.ID,
				Nome: projeto.Usuario.Nome,
			},
		}
		projetosDTO = append(projetosDTO, projetoDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projetosDTO)
}

// cria um novo projeto.
func CreateProjeto(w http.ResponseWriter, r *http.Request) {
	var projeto models.Projeto

	err := json.NewDecoder(r.Body).Decode(&projeto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&projeto).Error
	if err != nil {
		http.Error(w, "Erro ao criar projeto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateProjeto atualiza os dados de um projeto existente.
func UpdateProjeto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var projeto models.Projeto

	err := db.DB.First(&projeto, params["id"]).Error
	if err != nil {
		http.Error(w, "Projeto não encontrado", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&projeto)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	db.DB.Save(&projeto)

	w.WriteHeader(http.StatusNoContent)
}

// DeleteProjeto exclui um projeto existente.
func DeleteProjeto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := db.DB.Delete(&models.Projeto{}, params["id"]).Error
	if err != nil {
		http.Error(w, "Erro ao deletar projeto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
