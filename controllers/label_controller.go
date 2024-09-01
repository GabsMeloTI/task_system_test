package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/label_dto"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetEtiqueta(w http.ResponseWriter, r *http.Request) {
	var etiquetas []models.Etiqueta

	db.DB.Preload("Tarefa").Find(&etiquetas)

	var EtiquetasDTO []label_dto.ListagemEtiquetaDTO

	for _, etiqueta := range etiquetas {
		EtiquetaDTO := label_dto.ListagemEtiquetaDTO{
			ID:   etiqueta.ID,
			Nome: etiqueta.Nome,
			Cor:  etiqueta.Cor,
		}
		EtiquetasDTO = append(EtiquetasDTO, EtiquetaDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EtiquetasDTO)
}

func GetEtiquetaId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var etiquetas []models.Etiqueta
	err := db.DB.Preload("Tarefa").First(&etiquetas, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var EtiquetasDTO []label_dto.ListagemEtiquetaDTO

	for _, etiqueta := range etiquetas {
		EtiquetaDTO := label_dto.ListagemEtiquetaDTO{
			ID:   etiqueta.ID,
			Nome: etiqueta.Nome,
			Cor:  etiqueta.Cor,
		}
		EtiquetasDTO = append(EtiquetasDTO, EtiquetaDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EtiquetasDTO)
}

func CreateEtiqueta(w http.ResponseWriter, r *http.Request) {
	var etiqueta models.Etiqueta

	err := json.NewDecoder(r.Body).Decode(&etiqueta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = db.DB.Create(&etiqueta).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateEtiqueta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var etiqueta models.Etiqueta

	err := db.DB.First(&etiqueta, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&etiqueta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.DB.Save(&etiqueta)
	w.WriteHeader(http.StatusCreated)

}

func DeleteEtiqueta(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Etiqueta{}, params["id"]).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
