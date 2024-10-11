package book

import (
	bookEntities "src/internal/entities/book"
	bookService "src/internal/models/book"

	"example.com/m/v2/internal/repositories/book"
	"github.com/google/uuid"
)

type Service struct {
	repository *book.Repository
}

func NewService(repository *book.Repository) *Service {
	return &Service{repository: repository}
}
func (service *Service) Create(model *bookService.CreateModel) {
	bookEntity := bookEntities.Entity{
		Uuid: uuid.New(),
		Name: model.Name,
	}

	service.repository.Create(bookEntity)
}
