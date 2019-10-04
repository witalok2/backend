package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/witalok2/application/backend/database"
	"github.com/witalok2/application/backend/model"
)

// Criar Atividade
func PostAtividade(c echo.Context) error {
	db := database.DBConnection()

	strAtividade := model.Atividade{}
	err := c.Bind(&strAtividade)
	if err != nil {
		return c.JSON(400, err)
	}
	sql := "INSERT INTO tb_atividades(atividade_titulo, atividade_subtitulo, atividade_descricao, atividade_situacao) VALUES( ?, ?, ?,?)"
	result, err := db.Exec(sql, strAtividade.Titulo, strAtividade.SubTitulo, strAtividade.Descricao, strAtividade.Situacao)
	if err != nil {
		fmt.Print(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusCreated, echo.Map{"id": id})
}

// Lista todas atividades do Usuario
func GetLista(c echo.Context) error {
	db := database.DBConnection()
	sqlStatement := "SELECT id,	atividade_titulo, atividade_subtitulo, atividade_descricao, atividade_situacao FROM tb_atividades order by id"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()
	result := model.Atividades{}

	for rows.Next() {
		atividade := model.Atividade{}
		err2 := rows.Scan(&atividade.Id, &atividade.Titulo, &atividade.SubTitulo, &atividade.Descricao, &atividade.Situacao)
		// Exit if we get an error
		if err2 != nil {
			return err2
		}
		result.Atividades = append(result.Atividades, atividade)
	}
	return c.JSON(http.StatusCreated, result)
}

// Lista as atividades passado pelo ID
func GetListaParam(c echo.Context) error {
	db := database.DBConnection()
	requested_id := c.Param("id")
	fmt.Println(requested_id)
	var id int
	var titulo string
	var subTitulo string
	var descricao string
	var situacao string

	err := db.QueryRow("SELECT id, atividade_titulo, atividade_subtitulo, atividade_descricao, atividade_situacao FROM tb_atividades WHERE id = ?", requested_id).Scan(&id, &titulo, &subTitulo, &descricao, &situacao)

	if err != nil {
		fmt.Println(err)
	}

	response := model.Atividade{Id: id, Titulo: titulo, SubTitulo: subTitulo, Descricao: descricao, Situacao: situacao}
	return c.JSON(http.StatusOK, response)
}

// Atualiza atividade
func PutAtualiza(c echo.Context) error {
	db := database.DBConnection()
	strAtividade := model.Atividade{}
	err := c.Bind(&strAtividade)
	if err != nil {
		return c.JSON(400, err)
	}

	sql := "UPDATE tb_atividades SET atividade_titulo = ?, atividade_subtitulo = ?, atividade_descricao = ?, atividade_situacao = ? WHERE id = ?"
	result, err := db.Exec(sql, strAtividade.Titulo, strAtividade.SubTitulo, strAtividade.Descricao, strAtividade.Situacao, strAtividade.Id)
	if err != nil {
		fmt.Print(err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// Deleta Atividade
func Delete(c echo.Context) error {
	requested_id := c.Param("id")
	db := database.DBConnection()
	sql := "Delete FROM tb_atividades Where id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}

	result, err2 := stmt.Exec(requested_id)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(result.RowsAffected())
	return c.JSON(http.StatusOK, "Deleted")
}
