package repository

import (
	"database/sql"
)

type categoryRepo struct {
	db *sql.DB
}

type CategoryRepo interface {
	Create()
	GetAll()
	UpdateById()
	DeleteById()
	GetById()
}
