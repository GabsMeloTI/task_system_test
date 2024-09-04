package models

// AuthResponse representa a resposta de autenticação com um token JWT.
type AuthResponse struct {
	Token string `json:"token"`
}
