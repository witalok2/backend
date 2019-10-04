package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DBConnection() (db *sqlx.DB) {
	fmt.Println("Abrindo conexão com o banco")

	// Abre a conexão com o banco de dados.
	db, err := sqlx.Connect("mysql", "witalok2:92526018@(localhost:3306)/atividades")

	// Se houver um erro ao abrir a conexão.
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conecatdo com sucesso", db)
	return db
}
