package models

// @Description Dados de login do usuário
// @Accept json
// @Produce json
type LoginRequest struct {
	// @Param email body string true "Email do usuário"
	// @Param password body string true "Senha do usuário"
	Email    string `json:"email"`
	Password string `json:"password"`
}
