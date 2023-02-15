package module

import (
	"errors"
	"go-db/pkg/model"
)

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
