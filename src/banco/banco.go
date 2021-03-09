package banco

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // GORM
)

//Cidade a ser usado como modelo
type Cidade struct {
	gorm.Model
	Nome string
	Uf   string
}

//DBClient é o export do bd
var DBClient *gorm.DB

//IniciarMigracaoBD Inicia conexao com bando
func IniciarMigracaoBD() {

	stringConexao := "golang:golang@/cadastroSimples?charset=utf8&parseTime=True&loc=Local"
	db, erro := gorm.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}

	DBClient = db
	fmt.Println(db)

}
