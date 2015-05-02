package session

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"sync"
	"time"
)

//thread save pool of sessions
type Manager struct {
	sessions map[string]*session //map of all sessions
	l        sync.Mutex
	strength int
	timeout	 time.Duration
}

//Initialize the SessionManager - Strength is the number of bytes from the number generator - Resulting SessionID will be twice as long due to hex encoding
func NewManager(strength int, timeout time.Duration) *Manager {
	m := new(Manager)
	m.sessions = make(map[string]*session)
	m.strength = strength
	m.timeout  = timeout
	return m
}

//creates and returns a new Session. Returns error if generation of sessionID fails, or the generated sessionID is already in the map.
func (m *Manager) Create() (*session, error) {
	b := make([]byte, m.strength)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil, errors.New("could't create sessionID")
	}

	sid := hex.EncodeToString(b)
	now := time.Now()
	valid := now.Add(m.timeout)

	s := session{sessionID: sid, createdAt: now, validUntil: valid, test: "aa"}

	//test if sessionId already exists
	_, err := m.Get(sid)
	if err == nil {
		return nil, errors.New("dublicated sessionID")
	}

	//add session to the pool
	m.l.Lock()
	m.sessions[sid] = &s
	m.l.Unlock()

	return &s, nil
}

//destroy a session - always returns nill
func (m *Manager) Destroy(sid string) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.sessions, sid)
}

//get a existing session - returns error if session pool is empty, or given sid could not be found
func (m *Manager) Get(sid string) (*session, error) {
	m.l.Lock()
	defer m.l.Unlock()

	if len(m.sessions) == 0 {
		return nil, errors.New("session pool is empty")
	}

	if s, ok := m.sessions[sid]; !ok { //session not there!
		return nil, errors.New("could't find session")
	} else if time.Now().After(s.validUntil) { //session time out!
		delete(m.sessions, s.sessionID)
		return nil, errors.New("could't find session")
	} else { //session! :)
		return s, nil	
	}
}

func (m *Manager) String() string {
	out := "{[";
	for k := range m.sessions {
		out += m.sessions[k].String() + ","
	}
	out += "]}"
	return out
}