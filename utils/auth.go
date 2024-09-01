package utils

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

func VerificarSenha(senhaHash, senha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
	return err == nil
}

var jwtKey = []byte("sua_chave_secreta")

// Claims estrutura para os dados do token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Autenticado verifica se o usuário está autenticado
func Autenticado(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	tokenStr := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
