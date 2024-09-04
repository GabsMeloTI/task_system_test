package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mime/multipart"
	"net/http"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{
		Service: svc,
	}
}

// GetUsers retorna a lista de todos os usuários
// @Summary Lista todos os usuários
// @Description Retorna uma lista de usuários
// @Tags users
// @Produce json
// @Success 200 {array} user_dto.UserListingDTO
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /user [get]
func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Chamando GetUsers")
	users, err := c.Service.GetUsers()
	if err != nil {
		log.Printf("Erro ao obter usuários: %v", err)
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID retorna um usuário por ID
// @Summary Retorna um usuário por ID
// @Description Obtém os detalhes de um usuário específico pelo seu ID
// @Tags users
// @Produce json
// @Param id path string true "ID do Usuário"
// @Success 200 {array} user_dto.UserListingDTO
// @Failure 404 {string} string "Usuário não encontrado"
// @Router /user/{id} [get]
func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := c.Service.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser cria um novo usuário
// @Summary Cria um novo usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags users
// @Accept  multipart/form-data
// @Produce json
// @Param name formData string true "Nome do Usuário"
// @Param email formData string true "Email do Usuário"
// @Param password formData string true "Senha do Usuário"
// @Param avatar formData file false "Imagem do Avatar do Usuário"
// @Success 201 {string} string "Usuário criado com sucesso"
// @Failure 400 {string} string "Erro na solicitação"
// @Router /user [post]
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var file multipart.File
	var fileName string
	if f, _, err := r.FormFile("avatar"); err == nil {
		defer f.Close()
		file = f
		fileName = r.MultipartForm.File["avatar"][0].Filename
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = c.Service.CreateUser(&user, file, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateUser atualiza os dados de um usuário
// @Summary Atualiza os dados de um usuário
// @Description Atualiza os detalhes de um usuário específico pelo seu ID
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID do Usuário"
// @Param name formData string false "Nome do Usuário"
// @Param email formData string false "Email do Usuário"
// @Param password formData string false "Senha do Usuário"
// @Param avatar formData file false "Imagem do Avatar do Usuário"
// @Success 204 "Usuário atualizado com sucesso"
// @Failure 400 {string} string "Erro na solicitação"
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /user/{id} [put]
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	var imageFile multipart.File
	var imageFileName string
	if f, _, err := r.FormFile("avatar"); err == nil {
		defer f.Close() // Ensure the file is closed after processing
		imageFile = f
		imageFileName = r.MultipartForm.File["avatar"][0].Filename
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = c.Service.UpdateUser(id, &user, imageFile, imageFileName)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser exclui um usuário
// @Summary Exclui um usuário
// @Description Exclui um usuário pelo ID
// @Tags users
// @Param id path string true "ID do Usuário"
// @Success 204 "Usuário excluído com sucesso"
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /user/{id} [delete]
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := c.Service.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with an optional avatar image
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "User Name"
// @Param email formData string true "User Email"
// @Param password formData string true "User Password"
// @Param avatar formData file false "Avatar image file"
// @Success 201 {string} string "Created"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to register user"
// @Router /user/register [post]
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	imageFile, _, err := r.FormFile("avatar")
	var imageFileName string
	if err == nil {
		defer imageFile.Close()
		imageFileName = r.MultipartForm.File["avatar"][0].Filename
	} else {
		imageFileName = ""
	}

	if err := c.Service.RegisterUser(&user, imageFile, imageFileName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User Credentials"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {string} string "Invalid request payload"
// @Failure 401 {string} string "Unauthorized"
// @Router /user/login [post]
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := c.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := models.AuthResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateUserImage godoc
// @Summary Update a user's profile image
// @Description Update the profile image of a user by their ID
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID"
// @Param avatar formData file true "Avatar image file"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to update user image"
// @Router /user/{id}/avatar [put]
func (c *UserController) UpdateUserImage(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := r.ParseMultipartForm(10 << 20) // Limit the size to 10 MB
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = c.Service.UpdateUserImage(id, file, "avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
