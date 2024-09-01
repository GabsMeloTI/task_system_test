package subtask_dto

import (
	"awesomeProject/dto/task_dto"
	"time"
)

type ListagemSubtarefasDTO struct {
	ID          uint                              `json:"id"`
	Nome        string                            `json:"ds_nome"`
	Descricao   string                            `json:"ds_descricao"`
	DataCriacao time.Time                         `json:"dt_criacao"`
	Status      string                            `json:"ds_status"`
	TarefaDto   task_dto.ListagemBasicaTarefasDTO `json:"tarefas"`
}
