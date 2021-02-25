package controllers

import (
	"api/database"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	criar  models.Pessoa
	pessoa models.Pessoa
)

// CriarPessoa insere um usuario no banco
func CriarPessoa(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	if err := c.ShouldBindJSON(&criar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := criar.Preparar(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pessoa := models.Pessoa{
		ID:             0,
		Nome:           criar.Nome,
		CPF:            criar.CPF,
		DataNascimento: criar.DataNascimento,
		Telefone:       criar.Telefone,
		Email:          criar.Email,
		Rua:            criar.Rua,
		Bairro:         criar.Bairro,
		Complemento:    criar.Complemento,
		Cidade:         criar.Cidade,
	}
	if err := db.Create(&pessoa).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": pessoa})
}

//BuscarPessoas busca todos os pessoas no banco de dados
func BuscarPessoas(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	var pessoas []models.Pessoa
	if err := db.Table("pessoas").Find(&pessoas).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"dados": pessoas})
}

// BuscarPessoa busca uma pessoa no banco de dados
func BuscarPessoa(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	var pessoa models.Pessoa
	if err := db.Table("pessoas").Where("id = ?", c.Param("id")).Take(&pessoa).Error; err != nil {
		c.JSON(404, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dados": pessoa})
}

// AtualizarPessoa atualiza os dados de uma pessoa no banco de dados
func AtualizarPessoa(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	if err := db.Where("id = ?", c.Param("id")).First(&pessoa).Error; err != nil {
		c.JSON(400, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	var atualizar models.Pessoa
	if err := c.ShouldBindJSON(&atualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := atualizar.Preparar(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Model(&pessoa).Update(atualizar)
	c.JSON(http.StatusOK, gin.H{"dados": pessoa})
}

// DeletarPessoa deleta uma pessoa do banco de dados
func DeletarPessoa(c *gin.Context) {
	db := database.SetupDB()
	defer db.Close()

	if err := db.Where("id = ?", c.Param("id")).First(&pessoa).Error; err != nil {
		c.JSON(400, gin.H{"error": "Pessoa não encontrada"})
		return
	}

	db.Delete(&pessoa)

	c.JSON(http.StatusOK, gin.H{"dados": true})
}
