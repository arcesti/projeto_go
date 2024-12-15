package repository

import (
	"context"
	"log"
	"github.com/jackc/pgx/v4"
)

type DBConfig struct {
	Host	 string
	Port	 string
	User	 string
	Password string
	Dbname	 string
}

var DB *pgx.Conn

func InitDB(config DBConfig) {
	var err error

	connStr := "postgres://"+config.User+":"+config.Password+"@"+config.Host+":"+config.Port+"/"+config.Dbname

	DB, err = pgx.Connect(context.Background(), connStr)
	if err!=nil {
		log.Fatal("Conexão com o banco falhou " + err.Error())
	}
	log.Println("Conexão com o banco de dados realizada com sucesso")
}

func CloseDB() {
	if err := DB.Close(context.Background()); err != nil {
		log.Fatal("Falha ao tentar fechar conexão com o banco de dados")
	}
	log.Println("Conexão com o banco fechada com sucesso")
}