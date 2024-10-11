package book

import (
	"fmt"

	"example.com/m/v2/internal/entities/book"
)

type Repository struct {
	database []book.Entity
}

func NewRepository() *Repository {
	return &Repository{
		database: make([]book.Entity, 0),
	}
}

func (repository *Repository) Create(entity book.Entity) {
	repository.database = append(repository.database, entity)
	fmt.Println(repository.database)
}
