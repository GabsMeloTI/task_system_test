package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/subtask_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSubtarefa(w http.ResponseWriter, r *http.Request) {
	var subtarefa []models.Subtarefa

	db.DB.Preload("Tarefas").Find(&subtarefa)

	var subtarefasDto []subtask_dto.ListagemSubtarefasDTO

	for _, subtarefa := range subtarefa {
		subtarefasDto = append(subtarefasDto, subtask_dto.ListagemSubtarefasDTO{
			ID:          subtarefa.ID,
			Nome:        subtarefa.Nome,
			Descricao:   subtarefa.Descricao,
			DataCriacao: subtarefa.DataCriacao,
			Status:      subtarefa.Status,
			TarefaDto: task_dto.ListagemBasicaTarefasDTO{
				ID:   subtarefa.Tarefa.ID,
				Nome: subtarefa.Tarefa.Nome,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtarefasDto)
}

func GetSubtarefasId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var subtarefa []models.Subtarefa

	err := db.DB.Preload("Tarefas").First(&subtarefa, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var subtarefasDto []subtask_dto.ListagemSubtarefasDTO

	for _, subtarefa := range subtarefa {
		subtarefasDto = append(subtarefasDto, subtask_dto.ListagemSubtarefasDTO{
			ID:          subtarefa.ID,
			Nome:        subtarefa.Nome,
			Descricao:   subtarefa.Descricao,
			DataCriacao: subtarefa.DataCriacao,
			Status:      subtarefa.Status,
			TarefaDto: task_dto.ListagemBasicaTarefasDTO{
				ID:   subtarefa.Tarefa.ID,
				Nome: subtarefa.Tarefa.Nome,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtarefasDto)
}

func CreateSubtarefa(w http.ResponseWriter, r *http.Request) {
	var subtarefa models.Subtarefa

	err := json.NewDecoder(r.Body).Decode(&subtarefa)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&subtarefa).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateSubtarefa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var subtarefa models.Subtarefa

	err := db.DB.First(&subtarefa, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&subtarefa)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.DB.Save(&subtarefa)
	w.WriteHeader(http.StatusCreated)
}

func DeleteSubtarefa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Subtarefa{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
