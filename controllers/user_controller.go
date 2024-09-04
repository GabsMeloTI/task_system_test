package controllers

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
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
// @Success 200 {array} models.User
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /user [get]
func (c *UserController) GetUsers(ctx echo.Context) error {
	log.Println("Chamando GetUsers")
	users, err := c.Service.GetUsers()
	if err != nil {
		log.Printf("Erro ao obter usuários: %v", err)
		return ctx.JSON(http.StatusInternalServerError, "Failed to get users")
	}

	return ctx.JSON(http.StatusOK, users)
}

// GetUserByID retorna um usuário por ID
// @Summary Retorna um usuário por ID
// @Description Obtém os detalhes de um usuário específico pelo seu ID
// @Tags users
// @Produce json
// @Param id path string true "ID do Usuário"
// @Success 200 {object} models.User
// @Failure 404 {string} string "Usuário não encontrado"
// @Router /user/{id} [get]
func (c *UserController) GetUserByID(ctx echo.Context) error {
	id := ctx.Param("id")

	user, err := c.Service.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	return ctx.JSON(http.StatusOK, user)
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
func (c *UserController) CreateUser(ctx echo.Context) error {
	err := ctx.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	var file multipart.File
	var fileName string
	if f, _, err := ctx.Request().FormFile("avatar"); err == nil {
		defer f.Close()
		file = f
		fileName = ctx.Request().MultipartForm.File["avatar"][0].Filename
	}

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = c.Service.CreateUser(&user, file, fileName)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
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
func (c *UserController) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")

	err := ctx.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Unable to parse form")
	}

	var imageFile multipart.File
	var imageFileName string
	if f, _, err := ctx.Request().FormFile("avatar"); err == nil {
		defer f.Close()
		imageFile = f
		imageFileName = ctx.Request().MultipartForm.File["avatar"][0].Filename
	}

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = c.Service.UpdateUser(id, &user, imageFile, imageFileName)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return ctx.NoContent(http.StatusNoContent)
}

// DeleteUser exclui um usuário
// @Summary Exclui um usuário
// @Description Exclui um usuário pelo ID
// @Tags users
// @Param id path string true "ID do Usuário"
// @Success 204 "Usuário excluído com sucesso"
// @Failure 500 {string} string "Erro interno do servidor"
// @Router /user/{id} [delete]
func (c *UserController) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := c.Service.DeleteUser(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return ctx.NoContent(http.StatusNoContent)
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
func (c *UserController) RegisterUser(ctx echo.Context) error {
	err := ctx.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Unable to parse form")
	}

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	imageFile, _, err := ctx.Request().FormFile("avatar")
	var imageFileName string
	if err == nil {
		defer imageFile.Close()
		imageFileName = ctx.Request().MultipartForm.File["avatar"][0].Filename
	} else {
		imageFileName = ""
	}

	if err := c.Service.RegisterUser(&user, imageFile, imageFileName); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
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
func (c *UserController) Login(ctx echo.Context) error {
	var credentials models.LoginRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&credentials); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	token, err := c.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err.Error())
	}

	response := models.AuthResponse{
		Token: token,
	}

	return ctx.JSON(http.StatusOK, response)
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
func (c *UserController) UpdateUserImage(ctx echo.Context) error {
	id := ctx.Param("id")

	file, fileHeader, err := ctx.Request().FormFile("avatar")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid file")
	}
	defer file.Close()

	imageFileName := fileHeader.Filename

	if err := c.Service.UpdateUserImage(id, file, imageFileName); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update user image")
	}

	return ctx.NoContent(http.StatusNoContent)
}
