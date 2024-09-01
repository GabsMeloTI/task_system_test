package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSecao(w http.ResponseWriter, r *http.Request) {
	var secoes []models.Secao
	var secoesDTO []section_dto.ListagemSecaoDTO

	db.DB.Preload("Projeto").Preload("Usuario").Find(&secoes)

	for _, secao := range secoes {
		secaoDTO := section_dto.ListagemSecaoDTO{
			ID:          secao.ID,
			Nome:        secao.Nome,
			Descricao:   secao.Descricao,
			DataCriacao: secao.DataCriacao,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   secao.Usuario.ID,
				Nome: secao.Usuario.Nome,
			},
			ProjetoDTO: project_dto.ListagemBasicaProjetoDTO{
				ID:     secao.Projeto.ID,
				Nome:   secao.Projeto.Nome,
				Status: secao.Projeto.Status,
			},
		}
		secoesDTO = append(secoesDTO, secaoDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(secoesDTO)
}

func GetSecaoId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var secoes []models.Secao
	var secoesDTO []section_dto.ListagemSecaoDTO

	err := db.DB.Preload("Projeto").Preload("Usuario").First(&secoes, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for _, secao := range secoes {
		secaoDTO := section_dto.ListagemSecaoDTO{
			ID:          secao.ID,
			Nome:        secao.Nome,
			Descricao:   secao.Descricao,
			DataCriacao: secao.DataCriacao,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   secao.Usuario.ID,
				Nome: secao.Usuario.Nome,
			},
			ProjetoDTO: project_dto.ListagemBasicaProjetoDTO{
				ID:     secao.Projeto.ID,
				Nome:   secao.Projeto.Nome,
				Status: secao.Projeto.Status,
			},
		}
		secoesDTO = append(secoesDTO, secaoDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(secoesDTO)
}

func CreateSecao(w http.ResponseWriter, r *http.Request) {
	var secao models.Secao

	err := json.NewDecoder(r.Body).Decode(&secao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&secao).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateSecao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var secao models.Secao

	err := db.DB.First(&secao, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&secao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Save(&secao)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteSecao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Secao{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
