package config

import "github.com/arcesti/projeto_go/internal/repository"

var DbConfig = repository.DBConfig{
	Host: "localhost",
	Port: "5432",
	User: "postgres",
	Password: "postgres123",
}