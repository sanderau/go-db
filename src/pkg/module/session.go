package module

import (
	"errors"
	"go-db/pkg/model"
	"log"
)

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
}

// add a database to the current session
func (s *SessionClient) AddDatabase(ndb model.Database) (model.Database, error) {
	// check to see if the database exists already
	_, err := s.GetDatabase(ndb.Name)
	if err == nil {
		return model.Database{}, errors.New("database already exists")
	}

	// add the new database to the session
	s.databases = append(s.databases, ndb)

	// return the newly created db inside of the
	return ndb, nil
}

// get a specific database from the session by name
func (s *SessionClient) GetDatabase(name string) (model.Database, error) {
	// iterate through all databases and look for matching name
	for _, v := range s.databases {
		if v.Name == name {
			return v, nil
		}
	}

	// match could not be found return error
	return model.Database{}, errors.New("dne")
}

// get all the databases for current sesions
func (s *SessionClient) GetDatabases() ([]model.Database, error) {
	return s.databases, nil
}

// modify an existing resource
func (s *SessionClient) PutDatabase(old string, new string) (model.Database, error) {
	// sort through current databases. if db by old name can be found
	// change name, and return newly named db.
	for i, v := range s.databases {
		if v.Name == old {
			s.databases[i].Name = new
			return v, nil
		}
	}

	// if db by old name cannot be found return an error
	return model.Database{}, errors.New("original database could not be found")
}

// delete a database
func (s *SessionClient) DeleteDatabase(name string) {
	index := -1
	for i, v := range s.databases {
		if v.Name == name {
			index = i
		}
	}

	if index != -1 {
		s.databases = append(s.databases[:index], s.databases[index+1:]...)
	}
}

// helper function to get the index of a database.
// returns -1 if database does not exist
func (s SessionClient) getDatabaseIndex(name string) int {
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
	for i, v := range s.databases[dbIdx].Collections {
		if v.Name == collectionName {
			return i
		}
	}

	return -1
}

// add a collection to a database
func (s *SessionClient) AddCollection(nc model.Collection, dbName string) (model.Collection, error) {
	// get the index of the database
	index := s.getDatabaseIndex(dbName)

	// check to see if it exists already
	if s.getCollectionIndex(index, nc.Name) != -1 {
		return model.Collection{}, errors.New(model.CollectionExists)
	}

	// if the database exists add the collection
	if index != -1 {
		s.databases[index].Collections = append(s.databases[index].Collections, nc)
		return nc, nil
	}

	// else return an error that the database does not exist
	return model.Collection{}, errors.New(model.DbNotFound)
}

// returns all the collections for a given database
func (s *SessionClient) GetCollections(dbName string) ([]model.Collection, error) {
	// get the index of the database
	index := s.getDatabaseIndex(dbName)

	if index == -1 {
		return []model.Collection{}, errors.New(model.DbNotFound)
	}

	return s.databases[index].Collections, nil
}

// returns a specific collection from a databse
func (s *SessionClient) GetCollection(dbName string, cName string) (model.Collection, error) {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Collection{}, errors.New(model.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return model.Collection{}, errors.New(model.CollectionNotFound)
	}

	return s.databases[dIndex].Collections[cIndex], nil
}

func (s *SessionClient) PutCollection(dbName string, nc model.Collection) (model.Collection, error) {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Collection{}, errors.New(model.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, nc.Name)
	if cIndex == -1 {
		return model.Collection{}, errors.New(model.CollectionNotFound)
	}

	s.databases[dIndex].Collections[cIndex] = nc

	return s.databases[dIndex].Collections[cIndex], nil
}

func (s *SessionClient) DeleteCollection(dbName string, cName string) error {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		log.Println(model.DbNotFound)
		return nil
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		log.Println(model.CollectionNotFound)
		return nil
	}

	s.databases[dIndex].Collections = append(s.databases[dIndex].Collections[:cIndex], s.databases[dIndex].Collections[cIndex+1:]...)
	return nil
}
