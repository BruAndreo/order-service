package domain

import "github.com/google/uuid"

type Pizzeria struct {
	ID   uuid.UUID
	Name string
}
