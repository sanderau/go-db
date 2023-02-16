package module

import (
	"errors"
	myError "go-db/pkg/errors"
	"go-db/pkg/model"
)

// add a database to the current session
func (s *SessionClient) AddDatabase(ndb model.Database) (model.Database, error) {
	// check to see if the database exists already
	_, err := s.GetDatabase(ndb.Name)
	if err == nil {
		return model.Database{}, errors.New(myError.DbExists)
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
	return model.Database{}, errors.New(myError.DbNotFound)
}

// get all the databases for current sesions
func (s *SessionClient) GetDatabases() ([]model.Database, error) {
	return s.databases, nil
}

// modify an existing resource
func (s *SessionClient) PutDatabase(old string, new string) (model.Database, error) {
	// sort through current databases. if db by old name can be found
	// change name, and return newly named db.
	if len(s.databases) == 0 {
		return model.Database{}, errors.New(myError.DbNotFound)
	}

	// check to see if the new one already exists
	for _, v := range s.databases {
		if v.Name == new {
			// database already exists cannot update
			return model.Database{}, errors.New(myError.DbExists)
		}
	}

	for i, v := range s.databases {
		if v.Name == old {
			s.databases[i].Name = new
			return v, nil
		}
	}

	// if db by old name cannot be found return an error
	return model.Database{}, errors.New(myError.DbNotFound)
}

// delete a database
func (s *SessionClient) DeleteDatabase(name string) error {
	// idempotent operation, so regardless if it exists or not return 204 to user
	dIndex := s.getDatabaseIndex(name)
	if dIndex == -1 {
		return nil
	}

	// check to see if it has children
	if len(s.databases[dIndex].Collections) != 0 {
		return errors.New(myError.DatabaseHasKids)
	}

	s.databases = append(s.databases[:dIndex], s.databases[dIndex+1:]...)
	return nil
}
