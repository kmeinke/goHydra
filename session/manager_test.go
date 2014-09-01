package session

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	m := NewManager(32)
	s, err := m.Create()
	if err != nil || len(m.sessions) != 1 || len(s.sessionID) != 32*2 {
		t.Errorf("error: %v, poolLength: %d, sessionStrength: %d, sessionID: %v", err, len(m.sessions), len(s.sessionID), s.sessionID)
	}
}

func TestCreateUniq(t *testing.T) {
	m := NewManager(32)

	for x := 0; x <= 5; x++ {
		_, err := m.Create()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestDestroy(t *testing.T) {
	m := NewManager(32)
	s, err := m.Create()
	m.Destroy(s.SessionID())

	if err != nil || len(m.sessions) != 0 {
		t.Error(err)
	}
}

func TestDestroyEmpty(t *testing.T) {
	m := NewManager(32)
	m.Destroy("i am a session that is not there")
	//delete none existing map key is not an error but an no-op Oo
}

func TestGetOnEmptyList(t *testing.T) {
	m := NewManager(32)
	_, err := m.Get("empty hashmap")

	if err == nil {
		t.Error(err, m)
	} else {
		t.Log(err)
	}
}

func TestGetNotExisting(t *testing.T) {
	m := NewManager(32)
	m.Create()
	_, err := m.Get("2i am a session that is not there")

	if err == nil {
		t.Error("Got a not existing session. doh!")
	} else {
		t.Log(err)
	}
}

func TestGetExisting(t *testing.T) {
	m := NewManager(32)
	s, err := m.Create()
	if err != nil {
		t.Error(err, s, m)
	}

	s2, err2 := m.Get(s.SessionID())

	if err2 != nil || s.SessionID() != s2.SessionID() {
		t.Error(err)
	}
	fmt.Print()
}

/*func BenchmarkCreateAlot(b *testing.B) {
	m := NewManager(32)

	for x := 0; x <= b.N;x++ {
		_, err := m.Create()
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}*/
