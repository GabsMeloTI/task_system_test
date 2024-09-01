package controllers

import (
	"awesomeProject/db"
	"awesomeProject/dto/user_dto"
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var jwtKey = []byte("sua_chave_secreta")

// lista todos os usuários.
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	db.DB.Preload("Projetos").Find(&usuarios)

	var usuariosDTO []user_dto.ListagemUsuarioDTO

	for _, usuario := range usuarios {
		usuarioDTO := user_dto.ListagemUsuarioDTO{
			ID:            usuario.ID,
			Nome:          usuario.Nome,
			Email:         usuario.Email,
			Foto:          usuario.Foto,
			DataCriacao:   usuario.DataCriacao,
			ProjetosCount: len(usuario.Projetos),
		}
		usuariosDTO = append(usuariosDTO, usuarioDTO)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuariosDTO)
}

// lista usuário puxando pelo id.
func GetUsuarioId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var usuario models.Usuario
	err := db.DB.Preload("Projetos").First(&usuario, params["id"]).Error
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	usuarioDTO := user_dto.ListagemUsuarioDTO{
		ID:            usuario.ID,
		Nome:          usuario.Nome,
		Email:         usuario.Email,
		Foto:          usuario.Foto,
		DataCriacao:   usuario.DataCriacao,
		ProjetosCount: len(usuario.Projetos),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarioDTO)
}

// criar um usuario - ok = 201
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var existingUsuario models.Usuario
	if err := db.DB.Where("email = ?", usuario.Email).First(&existingUsuario).Error; err == nil {
		http.Error(w, "Email já cadastrado", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashSenha(usuario.Senha)
	if err != nil {
		http.Error(w, "Erro ao criptografar senha", http.StatusInternalServerError)
		return
	}
	usuario.Senha = hashedPassword

	err = db.DB.Create(&usuario).Error
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// atualiza dados do usuario, ok = 204
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var usuario models.Usuario

	err := db.DB.First(&usuario, params["id"]).Error
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	db.DB.Save(&usuario)

	w.WriteHeader(http.StatusNoContent)
}

// excluir usuario
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := db.DB.Delete(&models.Usuario{}, params["id"]).Error
	if err != nil {
		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Registro de usuário - ok = 201
func RegisterUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var existingUsuario models.Usuario
	if err := db.DB.Where("email = ?", usuario.Email).First(&existingUsuario).Error; err == nil {
		http.Error(w, "Email já cadastrado", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashSenha(usuario.Senha)
	if err != nil {
		http.Error(w, "Erro ao criptografar senha", http.StatusInternalServerError)
		return
	}
	usuario.Senha = hashedPassword

	err = db.DB.Create(&usuario).Error
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login do usuário - ok = 200
func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	err = db.DB.Where("email = ?", creds.Email).First(&usuario).Error
	if err != nil || !utils.VerificarSenha(usuario.Senha, creds.Senha) {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &utils.Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
