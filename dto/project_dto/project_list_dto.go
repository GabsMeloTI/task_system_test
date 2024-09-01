package project_dto

import (
	"awesomeProject/dto/user_dto"
	"time"
)

type ListagemProjetoDTO struct {
	ID        uint                              `json:"id"`
	Nome      string                            `json:"ds_nome"`
	Descricao string                            `json:"ds_descricao"`
	Status    string                            `json:"ds_status"`
	Data      time.Time                         `json:"data"`
	Usuario   user_dto.ListagemBasicaUsuarioDTO `json:"usuario"`
}
