package main

import (
	"cadastroSimples/src/banco"
	"cadastroSimples/src/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	
	banco.IniciarMigracaoBD();

	r.POST("/usuarios", controllers.CreateUser)
	r.GET("/usuarios/state", controllers.GetUsersByState)
	r.GET("/usuarios/id", controllers.GetUser)
	r.GET("/usuarios", controllers.GetUsers)
	r.PUT("/usuarios", controllers.UpdtUser)
	r.DELETE("/usuarios", controllers.DeleteUser)

	if erro := r.Run(":5000"); erro != nil {
		log.Fatal(erro.Error())
	}

}
