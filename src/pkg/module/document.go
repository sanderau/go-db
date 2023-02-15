package module

import (
	"encoding/json"
	"errors"
	"go-db/pkg/model"
	"strings"

	myError "go-db/pkg/errors"

	"github.com/google/uuid"
)

// create a new document
func (s *SessionClient) PostDocument(dbName string, cName string, doc model.Document) (model.Document, error) {
	// first get the index of the database and collection
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Document{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return model.Document{}, errors.New(myError.CollectionNotFound)
	}

	// with the index of the database and colleciton create a uuid for the objectID and then insert the new document
	doc.ObjectID = uuid.New()
	s.databases[dIndex].Collections[cIndex].Documents = append(s.databases[dIndex].Collections[cIndex].Documents, doc)

	return doc, nil
}

// get all of the documents in this collection
func (s *SessionClient) GetDocuments(dbName string, cName string) ([]model.Document, error) {
	// first get the index of the database and collection
	// if a db/collection of that name cannot be found return error
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return []model.Document{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return []model.Document{}, errors.New(myError.CollectionNotFound)
	}

	// return all of the documents
	return s.databases[dIndex].Collections[cIndex].Documents, nil
}

// get document by search
func (s *SessionClient) GetDocumentsBySearch(dbName string, cName string, search model.Search) ([]model.Document, error) {
	// first get the index of the database and collection
	// if a db/collection of that name cannot be found return error
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return []model.Document{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return []model.Document{}, errors.New(myError.CollectionNotFound)
	}

	// check to make sure there are documents before trying to search through them
	if len(s.databases[dIndex].Collections[cIndex].Documents) == 0 {
		return []model.Document{}, nil
	}

	// go through the documents and look for substrings
	var docs []model.Document
	for _, v := range s.databases[dIndex].Collections[cIndex].Documents {
		data, err := json.Marshal(&v.Data)
		if err != nil {
			return []model.Document{}, err
		}

		if strings.Contains(string(data), search.Keyword) {
			docs = append(docs, v)
		}
	}

	return docs, nil
}

// get a specific document from the collection
func (s *SessionClient) GetDocument(dbName string, cName string, id string) (model.Document, error) {
	// first get the index of the database and collection
	// if a db/collection of that name cannot be found return error
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Document{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return model.Document{}, errors.New(myError.CollectionNotFound)
	}

	// check to make sure there are documents
	if len(s.databases[dIndex].Collections[cIndex].Documents) == 0 {
		return model.Document{}, errors.New(myError.DocumentNotFound)
	}

	// if there are documents search through them
	for _, doc := range s.databases[dIndex].Collections[cIndex].Documents {
		if doc.ObjectID.String() == id {
			// found matching uuid's, return to user
			return doc, nil
		}
	}

	return model.Document{}, errors.New(myError.DocumentNotFound)
}

// modify an existing document
func (s *SessionClient) PutDocument(dbName string, cName string, id string, newDoc model.Document) (model.Document, error) {
	// first get the index of the database and collection
	// if a db/collection of that name cannot be found return error
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return model.Document{}, errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return model.Document{}, errors.New(myError.CollectionNotFound)
	}

	// check to make sure there are documents
	if len(s.databases[dIndex].Collections[cIndex].Documents) == 0 {
		return model.Document{}, errors.New(myError.DocumentNotFound)
	}

	// find a document with that specific error, and change it if possible
	for i, doc := range s.databases[dIndex].Collections[cIndex].Documents {
		if doc.ObjectID.String() == id {
			s.databases[dIndex].Collections[cIndex].Documents[i].Data = newDoc.Data
			return s.databases[dIndex].Collections[cIndex].Documents[i], nil
		}
	}

	return model.Document{}, errors.New(myError.DocumentNotFound)
}

// delete a document
func (s *SessionClient) DeleteDocument(dbName string, cName string, id string) error {
	// first get the index of the database and collection
	// if a db/collection of that name cannot be found return error
	dIndex := s.getDatabaseIndex(dbName)
	if dIndex == -1 {
		return errors.New(myError.DbNotFound)
	}

	cIndex := s.getCollectionIndex(dIndex, cName)
	if cIndex == -1 {
		return errors.New(myError.CollectionNotFound)
	}

	// check to make sure there are documents
	if len(s.databases[dIndex].Collections[cIndex].Documents) == 0 {
		return errors.New(myError.DocumentNotFound)
	}

	// find a document with that specific error, and change it if possible
	idx := -1
	for i, doc := range s.databases[dIndex].Collections[cIndex].Documents {
		if doc.ObjectID.String() == id {
			idx = i
		}
	}

	// if a document with the provided id can be found, then remove it from the list
	if idx != -1 {
		s.databases[dIndex].Collections[cIndex].Documents = append(s.databases[dIndex].Collections[cIndex].Documents[:idx], s.databases[dIndex].Collections[cIndex].Documents[idx+1:]...)
	}

	// regardless if the item was found and removed the item does not exist so return a 204
	return nil
}
