package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type Atividade struct {
	Id        int    `json:"id"`
	Titulo    string `json:"atividade_titulo"`
	SubTitulo string `json:"atividade_subtitulo"`
	Descricao string `json:"atividade_descricao"`
	Situacao  string `json:"atividade_sistuacao"` // Pendente, Em Andamento, Finalizadas
}

type Atividades struct {
	Atividades []Atividade `json:"atividade"`
}

func main() {
	fmt.Println("Abrindo conexão com o banco")

	// Abre a conexão com o banco de dados.
	db, err := sqlx.Connect("mysql", "witalok2:92526018@(localhost:3306)/atividades")

	// Se houver um erro ao abrir a conexão.
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conecatdo com sucesso", db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	//e.GET("/", home)

	// Lista todas atividades do Usuario
	e.GET("/atividade", func(c echo.Context) error {
		sqlStatement := "SELECT id,	atividade_titulo, atividade_subtitulo, atividade_descricao, atividade_situacao FROM tb_atividades order by id"
		rows, err := db.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
			//return c.JSON(http.StatusCreated, u);
		}
		defer rows.Close()
		result := Atividades{}

		for rows.Next() {
			atividade := Atividade{}
			err2 := rows.Scan(&atividade.Id, &atividade.Titulo, &atividade.SubTitulo, &atividade.Descricao, &atividade.Situacao)
			// Exit if we get an error
			if err2 != nil {
				return err2
			}
			result.Atividades = append(result.Atividades, atividade)
		}
		return c.JSON(http.StatusCreated, result)

		//return c.String(http.StatusOK, "ok")
	})

	// Lista as atividades passado pelo ID
	e.GET("/atividade/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id)
		var id int
		var titulo string
		var subTitulo string
		var descricao string
		var situacao string

		err = db.QueryRow("SELECT id,	atividade_titulo, atividade_subtitulo, atividade_descricao, atividade_situacao FROM tb_atividades WHERE id = ?", requested_id).Scan(&id, &titulo, &subTitulo, &descricao, &situacao)

		if err != nil {
			fmt.Println(err)
		}

		response := Atividade{Id: id, Titulo: titulo, SubTitulo: subTitulo, Descricao: descricao, Situacao: situacao}
		return c.JSON(http.StatusOK, response)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1324"))

	defer e.Close()
	defer db.Close()
}
