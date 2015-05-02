package session

import (
	"testing"
	"time"
	_ "unsafe"
)

func TestShowSession(t *testing.T) {
	timeout,_ := time.ParseDuration("1h")
	m := NewManager(32, timeout)
 	m.Create()
	t.Log(m)
}

func TestCreate(t *testing.T) {
	timeout, _ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	s, err := m.Create()
	if err != nil || len(m.sessions) != 1 || len(s.sessionID) != 32*2 {
		t.Errorf("error: %v, poolLength: %d, sessionStrength: %d, sessionID: %v", err, len(m.sessions), len(s.sessionID), s.sessionID)
	}
}

func TestCreateUniq(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)

	for x := 0; x <= 500000; x++ {
		_, err := m.Create()
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestDestroy(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	s, err := m.Create()
	m.Destroy(s.SessionID())

	if err != nil || len(m.sessions) != 0 {
		t.Error(err)
	}
}

func TestDestroyEmpty(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	m.Destroy("i am a session that is not there")
	//delete none existing map key is not an error but an no-op Oo
}

func TestGetOnEmptyList(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	_, err := m.Get("empty hashmap")

	if err == nil {
		t.Error(err, m)
	} else {
		t.Log(err)
	}
}

func TestGetNotExisting(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	m.Create()
	_, err := m.Get("2i am a session that is not there")

	if err == nil {
		t.Error("Got a not existing session. doh!")
	} else {
		t.Log(err)
	}
}

func TestGetExisting(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	s, err := m.Create()
	if err != nil {
		t.Error(err, s, m)
	}

	s2, err2 := m.Get(s.SessionID())

	if err2 != nil || s.SessionID() != s2.SessionID() {
		t.Error(err)
	}
}

func TestTimeOut(t *testing.T) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)
	s, err := m.Create()
	if err != nil {
		t.Error(err, s, m)
	}
	
	time.Sleep(timeout)
	s2, err2 := m.Get(s.SessionID())

	if err2 == nil {
		t.Error("Got a timeout session. doh!: " +s2.String())
	} else {
		t.Log(err2)
	}
}

func BenchmarkCreateAlot(b *testing.B) {
	timeout,_ := time.ParseDuration("500ms")
	m := NewManager(32, timeout)

	for x := 0; x <= b.N;x++ {
		_, err := m.Create()
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}
