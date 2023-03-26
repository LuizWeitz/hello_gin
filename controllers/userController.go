package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luizweitz/hello-gin/initializers"
	"github.com/luizweitz/hello-gin/models"
)

func UserCreate(c *gin.Context) {

	// Criando um struct de body, o qual vai ser com o body da requesição, ou seja, o corpo com os dados do user para criar
	var body struct {
		Username string
		Email    string
		Age      int
		City     string
	}

	c.ShouldBindJSON(&body)

	// Recuperando dados da requesição e setando as informaçãoes da requesição no model user
	user := models.User{Username: body.Username, Email: body.Email, Age: body.Age, City: body.City}

	// Criando o registro do user no banco
	result := initializers.DB.Create(&user)

	// Caso tenha obtido sucesso, vai retornar código 201 de sucesso e o corpo do user criado
	c.JSON(200, gin.H{
		"user": user,
	})

	// Caso tenha obtido algum erro, vai retornar código 400 de erro
	if result.Error != nil {
		c.Status(400)
		return

	}
}

func UserListAll(c *gin.Context) {
	// Criando uma var onde será armazenado todos os users
	var users []models.User

	// Buscando banco de dados todos os users e atribuindo a variável que foi criado `anteriormente `
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {

	// Recuperando o id do user passado
	id := c.Param("id")

	// Criando váriavel onde vai ser armazendo as informações do user buscado
	var user models.User

	// Buscando no banco as informações do user de acordo com o id informado
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {

	// Recuperando o id do user da url
	id := c.Param("id")

	// Recuperando as informaçãoes do corpo da requesição
	var body struct {
		Username string
		Email    string
		Age      int
		City     string
	}

	c.Bind(&body)

	// Pesquisando no banco de dados as informações do user
	var user models.User
	initializers.DB.First(&user, id)

	// Realizando alteração dos dados
	initializers.DB.Model(&user).Updates(models.User{Username: body.Username, Email: body.Email, Age: body.Age, City: body.City})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {

	// Recuperando o id do user da url
	id := c.Param("id")

	// Deletando user do banco de dados
	initializers.DB.Delete(&models.User{}, id)

	c.Status(200)
}
