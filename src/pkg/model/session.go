package model

// constants for error handling
const DbNotFound = "database not found"
const CollectionExists = "collection already exists"
const CollectionNotFound = "collection not found"

type Database struct {
	Name        string       `json:"name"`
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name      string     `json:"name"`
	Documents []Document `json:"documents"`
}

type Document struct {
	ObjectID string
	Data     []byte
}
