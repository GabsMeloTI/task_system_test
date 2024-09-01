package comment_dto

import (
	"awesomeProject/dto/task_dto"
	"awesomeProject/dto/user_dto"
	"time"
)

type ListagemComentarioDTO struct {
	ID             uint                              `json:"id"`
	Conteudo       string                            `json:"ds_conteudo"`
	DataPublicacao time.Time                         `json:"dt_publicacao"`
	Imagem         string                            `json:"ds_imagem"`
	UsuarioDTO     user_dto.ListagemBasicaUsuarioDTO `json:"id_usuario"`
	TarefasDTO     task_dto.ListagemBasicaTarefasDTO `json:"id_tarefas"`
}
