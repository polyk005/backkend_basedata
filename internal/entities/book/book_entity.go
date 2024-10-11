package book

import "github.com/google/uuid"

type Entity struct {
	Uuid uuid.UUID
	Name string
}
