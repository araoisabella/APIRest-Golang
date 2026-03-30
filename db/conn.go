package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "go_db" //container onde roda o banco de dados
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// criar conexao com o banco
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo) //abrindo a conexao com o postgres
	if err != nil {
		panic(err)
	}

	err = db.Ping() //ping para ver se a conexao abriu com sucesso
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
