package service

import (
	"context"
	"strings"

	"github.com/arcesti/projeto_go/internal/config"
	"github.com/arcesti/projeto_go/internal/entity"
	"github.com/arcesti/projeto_go/internal/repository"
)

func CreateUser(user entity.User) {
	repository.InitDB(config.DbConfig)
	sql := repository.InsUser
	sql = strings.Replace(sql, "?", user.Nome, 1)
	sql = strings.Replace(sql, "?", user.Email, 1)
	sql = strings.Replace(sql, "?", string(user.Idade), 1)
	
	return repository.DB.Exec(context.Background(), sql);
}