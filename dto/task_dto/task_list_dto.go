package task_dto

import (
	"awesomeProject/dto/label_dto"
	"awesomeProject/dto/section_dto"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"time"
)

type ListagemTarefasDTO struct {
	ID                    uint                               `json:"id"`
	Nome                  string                             `json:"ds_nome"`
	Descricao             string                             `json:"ds_descricao"`
	DataConclusaoPrevista time.Time                          `json:"dt_conclusao_prevista"`
	Prioridade            models.Prioridade                  `json:"ds_prioridade"`
	CreatedAt             time.Time                          `json:"created_at"`
	Status                string                             `json:"ds_status"`
	Etiquetas             []label_dto.ListagemEtiquetaDTO    `json:"etiquetas"`
	UsuarioDTO            user_dto.ListagemBasicaUsuarioDTO  `json:"id_usuario"`
	SecaoDTO              section_dto.ListagemBasicaSecaoDTO `json:"id_secao"`
}
