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
