package module

import (
	"errors"
	myError "go-db/pkg/errors"
	"go-db/pkg/model"
	"log"
)

// add a collection to a database
func (s *SessionClient) AddCollection(nc model.Collection, dbName string) (model.Collection, error) {
	// get the index of the database
	index := s.getDatabaseIndex(dbName)
	if index == -1 {
		return model.Collection{}, errors.New(myError.DbNotFound)
	}

	// check to see if it exists already
	if s.getCollectionIndex(index, nc.Name) != -1 {
		return model.Collection{}, errors.New(myError.CollectionExists)
	}

	// if the database exists add the collection
	if index != -1 {
		s.databases[index].Collections = append(s.databases[index].Collections, nc)
		return nc, nil
	}

	// else return an error that the database does not exist
	return model.Collection{}, errors.New(myError.DbNotFound)
}

// returns all the collections for a given database
func (s *SessionClient) GetCollections(dbName string) ([]model.Collection, error) {
	// get the index of the database
	index := s.getDatabaseIndex(dbName)

	if index == -1 {
		return []model.Collection{}, errors.New(myError.DbNotFound)
	}

	return s.databases[index].Collections, nil
}

// returns a specific collection from a databse
func (s *SessionClient) GetCollection(dbName string, cName string) (model.Collection, error) {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Collection{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return model.Collection{}, errors.New(myError.CollectionNotFound)
	}

	return s.databases[dIndex].Collections[cIndex], nil
}

// modify an existing collection
func (s *SessionClient) PutCollection(dbName string, oldCollectionName string, nc model.Collection) (model.Collection, error) {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Collection{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, oldCollectionName)
	if cIndex == -1 {
		return model.Collection{}, errors.New(myError.CollectionNotFound)
	}

	s.databases[dIndex].Collections[cIndex].Name = nc.Name

	return s.databases[dIndex].Collections[cIndex], nil
}

// delete a collection
func (s *SessionClient) DeleteCollection(dbName string, cName string) error {
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		log.Println(myError.DbNotFound)
		return nil
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		log.Println(myError.CollectionNotFound)
		return nil
	}

	s.databases[dIndex].Collections = append(s.databases[dIndex].Collections[:cIndex], s.databases[dIndex].Collections[cIndex+1:]...)
	return nil
}
