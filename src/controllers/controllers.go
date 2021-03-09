package controllers

import (
	"cadastroSimples/src/banco"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CreateUser cria usuario usando Query params
func CreateUser(c *gin.Context) {
	estados := []string{
		"AC", "AM", "PA", "RR", "RO", "AP", "TO",
		"MA", "PI", "CE", "RN", "PB", "PE", "AL",
		"SE", "BA", "DF", "GO", "MT", "MS", "MG",
		"ES", "RJ", "SP", "PR", "SC", "RS"}

	nome := c.Query("nome")
	uf := c.Query("uf")

	contador := 0
	for i := 0; i < 27; i++ {
		if uf == estados[i] {
			contador++
		}
	}
	if contador == 1 {

		if erro := banco.DBClient.Create(&banco.Cidade{Nome: nome, Uf: uf}); erro != nil {
			fmt.Println(erro.Error)

		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Criado com sucesso",
		})
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"mensagem": "erro ao criar usuário",
		})

	}
}

//GetUsers busca todos usuarios
func GetUsers(c *gin.Context) {

	var cidades []banco.Cidade

	if erro := banco.DBClient.Find(&cidades); erro != nil {
		fmt.Println(erro.Error)

	}
	
	c.JSON(http.StatusOK, cidades)
}

//GetUser busca um usuário
func GetUser(c *gin.Context) {

	id := c.Query("id")

	var cidade banco.Cidade

	if erro := banco.DBClient.Where("id = ?", id).First(&cidade); erro != nil {
		fmt.Println(erro.Error)

	}

	c.JSON(http.StatusOK, cidade)
}

//GetUsersByState pega usarios pelo estado
func GetUsersByState(c *gin.Context) {
	uf := c.Query("uf")

	var cidades []banco.Cidade

	if result := banco.DBClient.Where("uf = ?", uf).Find(&cidades); result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro":     result.Error,
			"mensagem": "erro ao buscar por estado",
		})
	}

	c.JSON(http.StatusOK, cidades)
}

//UpdtUser atualiza dados
func UpdtUser(c *gin.Context) {

	id := c.Query("id")
	nome := c.Query("nome")
	uf := c.Query("uf")

	var cidade banco.Cidade
	if erro := banco.DBClient.Find(&cidade, id); erro.Error != nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"message": "Erro ao encontrar usario",
		})
	}

	if erro := banco.DBClient.Model(&cidade).Updates(&banco.Cidade{Nome: nome, Uf: uf}); erro.Error != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "Erro ao atualizar usuário",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Atualizado com sucesso",
	})
}

//DeleteUser deleta usuário
func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	var cidade banco.Cidade
	if erro := banco.DBClient.Where("id = ?", id).Delete(&cidade); erro.Error != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"message": "Erro ao deletar usuário",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deletado com sucesso",
	})
}
