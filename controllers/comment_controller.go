package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/comment_dto"
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetComentario(w http.ResponseWriter, r *http.Request) {
	var comentarios []models.Comentario

	db.DB.Preload("Usuario").Preload("Tarefa").Find(&comentarios)

	var ComentariosDTO []comment_dto.ListagemComentarioDTO

	for _, comentario := range comentarios {
		ComentarioDTO := comment_dto.ListagemComentarioDTO{
			ID:             comentario.ID,
			Conteudo:       comentario.Conteudo,
			DataPublicacao: comentario.DataPublicacao,
			Imagem:         comentario.Img,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   comentario.Usuario.ID,
				Nome: comentario.Usuario.Nome,
			},
			TarefasDTO: task_dto.ListagemBasicaTarefasDTO{
				ID:   comentario.Tarefa.ID,
				Nome: comentario.Tarefa.Nome,
			},
		}
		ComentariosDTO = append(ComentariosDTO, ComentarioDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ComentariosDTO)
}

func GetComentariosId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var comentarios []models.Comentario

	err := db.DB.Preload("Usuario").Preload("Tarefa").First(&comentarios, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var ComentariosDTO []comment_dto.ListagemComentarioDTO

	for _, comentario := range comentarios {
		ComentarioDTO := comment_dto.ListagemComentarioDTO{
			ID:             comentario.ID,
			Conteudo:       comentario.Conteudo,
			DataPublicacao: comentario.DataPublicacao,
			Imagem:         comentario.Img,
			UsuarioDTO: user_dto.ListagemBasicaUsuarioDTO{
				ID:   comentario.Usuario.ID,
				Nome: comentario.Usuario.Nome,
			},
			TarefasDTO: task_dto.ListagemBasicaTarefasDTO{
				ID:   comentario.Tarefa.ID,
				Nome: comentario.Tarefa.Nome,
			},
		}
		ComentariosDTO = append(ComentariosDTO, ComentarioDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ComentariosDTO)
}

func CreateComentario(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario

	err := json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&comentario).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateComentario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comentario models.Comentario
	err := db.DB.First(&comentario, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.DB.Save(&comentario)
	w.WriteHeader(http.StatusCreated)
}

func DeleteComentario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Etiqueta{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
