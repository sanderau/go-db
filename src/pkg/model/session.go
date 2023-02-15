package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Database struct {
	Name        string       `json:"name"`
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name      string     `json:"name"`
	Documents []Document `json:"documents"`
}

type Document struct {
	ObjectID uuid.UUID
	Data     json.RawMessage
}

type Search struct {
	Keyword string `json:"search"`
}
