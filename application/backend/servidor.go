package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/witalok2/application/backend/controller"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//---------------------------------------------- { Routes } --------------------------------------
	// Criar Atividade
	e.POST("/atividade/criar", controller.PostAtividade)
	// Lista todas atividades do Usuario
	e.GET("/atividade/listar", controller.GetLista)
	// Lista as atividades passado pelo ID
	e.GET("/atividade/listar/:id", controller.GetListaParam)
	// Atualiza Atividade
	e.PUT("/atividade/atualizar/:id", controller.PutAtualiza)
	// Exluir Atividade
	e.DELETE("/atividade/deletar/:id", controller.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":1324"))
	defer e.Close()
}
