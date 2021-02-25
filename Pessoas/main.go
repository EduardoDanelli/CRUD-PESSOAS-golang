package main

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	
	router.POST("/pessoas", controllers.CriarPessoa)
	router.GET("/pessoas", controllers.BuscarPessoas)
	router.GET("/pessoas/:id", controllers.BuscarPessoa)
	router.PUT("/pessoas/:id", controllers.AtualizarPessoa)
	router.DELETE("/pessoas/:id", controllers.DeletarPessoa)

	router.Run(":9000")
}
