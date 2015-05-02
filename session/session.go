package session

import (
	_ "crypto/rand"
	_ "encoding/base64"
	_ "errors"
	_ "io"
	"sync"
	"time"
)

//represents on session
type session struct {
	//willbecome list element
	sessionID	string //ident current user session
	//values		map[interface{}]interface{} //map of all session values
	rw sync.RWMutex
	createdAt	time.Time
	validUntil	time.Time
	test		string
}

func (s *session) SessionID() string {
	return s.sessionID
}

func (s *session) GetValue(key interface{}) interface{} {
	return nil
}

func (s *session) SetValue(key interface{}, value interface{}) error {
	return nil
}

func (s *session) RemValue(key interface{}) error {
	return nil
}


func (s *session) String() string {
	return "{sessionID: " + s.sessionID + ", createdAt: " + s.createdAt.String() + ", validUntil: " + s.validUntil.String() + "}"
}
