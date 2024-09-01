package task_dto

import "awesomeProject/models"

type ListagemBasicaTarefasDTO struct {
	ID         uint              `json:"id"`
	Nome       string            `json:"ds_nome"`
	Prioridade models.Prioridade `json:"ds_prioridade"`
}
