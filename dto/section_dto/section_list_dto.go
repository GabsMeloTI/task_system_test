package section_dto

import (
	"awesomeProject/dto/project_dto"
	"awesomeProject/dto/user_dto"
	"time"
)

type ListagemSecaoDTO struct {
	ID          uint                                 `json:"id"`
	Nome        string                               `json:"ds_nome"`
	Descricao   string                               `json:"ds_descricao"`
	DataCriacao time.Time                            `json:"ds_data_criacao"`
	UsuarioDTO  user_dto.ListagemBasicaUsuarioDTO    `json:"usuario"`
	ProjetoDTO  project_dto.ListagemBasicaProjetoDTO `json:"projeto"`
}
