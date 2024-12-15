package repository

import (
	"github.com/arcesti/projeto_go/internal/entity"
)

var InsUser = "INSERT INTO usuario VALUES (?,?,?)"

func AddUser() {
	var usr []entity.User
	usr = append(usr, entity.User{ID: 1, Nome: "Arcesti", Email: "arcesti15@gmail.com", Idade: 20})
}