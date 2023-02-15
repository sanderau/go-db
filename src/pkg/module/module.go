package module

import "go-db/pkg/model"

// this holds all the data for the current session
type SessionClient struct {
	databases []model.Database
}

// Return reference to session client so we can use the same one globally
// on the router
func NewSessionClient() *SessionClient {
	return &SessionClient{}
}

// interface that exposes all of the functionality of session
type SessionManager interface {
	// database functions
	AddDatabase(model.Database) (model.Database, error)
	GetDatabase(string) (model.Database, error)
	GetDatabases() ([]model.Database, error)
	PutDatabase(string, string) (model.Database, error)
	DeleteDatabase(string) error

	// collection functions
	AddCollection(model.Collection, string) (model.Collection, error)
	GetCollections(string) ([]model.Collection, error)
	GetCollection(string, string) (model.Collection, error)
	PutCollection(string, string) (model.Collection, error)
	DeleteCollection(string, string) error

	// document functions
	PostDocument(string, string, model.Document) (model.Document, error)
	GetDocuments(string, string) ([]model.Document, error)
}

// helper function to get the index of a database.
// returns -1 if database does not exist
func (s SessionClient) getDatabaseIndex(name string) int {
	if len(s.databases) == 0 {
		return -1
	}

	for i, v := range s.databases {
		if v.Name == name {
			return i
		}
	}

	return -1
}

// helper function to get the index of a collection
// returns -1 if database does not exist
func (s SessionClient) getCollectionIndex(dbIdx int, collectionName string) int {
	if len(s.databases[dbIdx].Collections) == 0 {
		return -1
	}

	for i, v := range s.databases[dbIdx].Collections {
		if v.Name == collectionName {
			return i
		}
	}

	return -1
}
