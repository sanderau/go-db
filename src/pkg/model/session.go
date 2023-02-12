package model

type Database struct {
	Name        string `json:"name"`
	Collections []Collection
}

type Collection struct {
	Name string
}
