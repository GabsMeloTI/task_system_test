package user_dto

import "time"

type ListagemUsuarioDTO struct {
	ID            uint      `json:"id"`
	Nome          string    `json:"nome"`
	Email         string    `json:"email"`
	Foto          string    `json:"foto"`
	CreatedAt     time.Time `json:"created_at"`
	ProjetosCount int       `json:"projeto_count"`
}
