package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetTarefa(w http.ResponseWriter, r *http.Request) {
	var tarefas []models.Tarefa
	var tarefasDto []task_dto.ListagemTarefasDTO

	// Carregar todas as associações
	err := db.DB.Preload("Usuario").Preload("Secao").Preload("Etiquetas").Find(&tarefas).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, tarefa := range tarefas {
		log.Printf("Tarefa: %+v\n", tarefa) // Log para verificar os dados carregados

		var etiquetasDTO []label_dto.ListagemEtiquetaDTO
		for _, etiqueta := range tarefa.Etiquetas {
			log.Printf("Etiqueta: %+v\n", etiqueta) // Log para verificar as etiquetas
			etiquetasDTO = append(etiquetasDTO, label_dto.ListagemEtiquetaDTO{
				ID:   etiqueta.ID,
				Nome: etiqueta.Nome,
				Cor:  etiqueta.Cor,
			})
		}

		tarefasDto = append(tarefasDto, task_dto.ListagemTarefasDTO{
			ID:                    tarefa.ID,
			Nome:                  tarefa.Nome,
			Descricao:             tarefa.Descricao,
			DataConclusaoPrevista: tarefa.DataConclusaoPrevista,
			Prioridade:            tarefa.Prioridade,
			CreatedAt:             task.CreatedAt,
			Status:                tarefa.Status,
			Etiquetas:             etiquetasDTO,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   tarefa.Usuario.ID,
				Nome: tarefa.Usuario.Nome,
			},
			SecaoDTO: section_dto.ListagemBasicaSecaoDTO{
				ID:   tarefa.Secao.ID,
				Nome: tarefa.Secao.Nome,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefasDto)
}

func GetTarefasId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tarefas []models.Tarefa
	var tarefasDto []task_dto.ListagemTarefasDTO

	err := db.DB.Preload("Usuario").Preload("Secao").First(&tarefas, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for _, tarefa := range tarefas {
		var etiquetasDTO []label_dto.ListagemEtiquetaDTO
		for _, etiqueta := range tarefa.Etiquetas {
			etiquetasDTO = append(etiquetasDTO, label_dto.ListagemEtiquetaDTO{
				ID:   etiqueta.ID,
				Nome: etiqueta.Nome,
				Cor:  etiqueta.Cor,
			})
		}

		tarefasDto = append(tarefasDto, task_dto.ListagemTarefasDTO{
			ID:                    tarefa.ID,
			Nome:                  tarefa.Nome,
			Descricao:             tarefa.Descricao,
			DataConclusaoPrevista: tarefa.DataConclusaoPrevista,
			Prioridade:            tarefa.Prioridade,
			CreatedAt:             task.CreatedAt,
			Status:                tarefa.Status,
			Etiquetas:             etiquetasDTO,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   tarefa.Usuario.ID,
				Nome: tarefa.Usuario.Nome,
			},
			SecaoDTO: section_dto.ListagemBasicaSecaoDTO{
				ID:   tarefa.Secao.ID,
				Nome: tarefa.Secao.Nome,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefasDto)
}

func CreateTarefa(w http.ResponseWriter, r *http.Request) {
	var tarefa models.Tarefa

	err := json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := models.ValidarPrioridadeError(tarefa.Prioridade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&tarefa).Error
	if err != nil {
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func AddEtiquetaToTarefa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tarefaID, err := strconv.Atoi(vars["tarefa_id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var etiqueta models.Etiqueta
	err = json.NewDecoder(r.Body).Decode(&etiqueta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tarefa models.Tarefa
	err = db.DB.First(&tarefa, tarefaID).Error
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err = db.DB.Model(&tarefa).Association("Etiquetas").Append(&etiqueta).Error
	if err != nil {
		http.Error(w, "Failed to add label to task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefa)
}

func UpdateTarefa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var tarefas models.Tarefa

	err := db.DB.First(&tarefas, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&tarefas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Save(&tarefas)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTarefa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Tarefa{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
