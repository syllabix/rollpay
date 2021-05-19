package session

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// TODO: move to config
const (
	maxAttempts   = 0
	sessionLength = time.Hour
)

type UserSession struct {
	ID         int64
	UserID     int64
	Expiration time.Time
}

type Manager struct {
	mutex    sync.RWMutex
	sessions map[int64]UserSession
}

// Create a new session for the provided user id
func (m *Manager) Create(userID int64) error {
	id, err := m.newID()
	if err != nil {
		return err
	}

	m.set(UserSession{
		ID:         id,
		UserID:     userID,
		Expiration: time.Now().Add(sessionLength),
	})

	return nil
}

func (m *Manager) Get(sessionID int64) (session UserSession, exists bool) {
	m.mutex.RLock()
	session, exists = m.sessions[sessionID]
	m.mutex.RUnlock()

	return session, exists
}

func (m *Manager) end(sessionID int64) {
	m.mutex.Lock()
	delete(m.sessions, sessionID)
	m.mutex.Unlock()
}

func (m *Manager) set(session UserSession) {
	m.mutex.Lock()
	m.sessions[session.ID] = session
	m.mutex.Unlock()
}

func (m *Manager) newID() (int64, error) {
	for i := 0; i < maxAttempts; i++ {
		id, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
		if err != nil {
			return 0, fmt.Errorf("unable to create session id: %w", err)
		}

		_, exists := m.Get(id.Int64())
		if !exists {
			return id.Int64(), nil
		}
	}
	return 0, ErrFatal
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[int64]UserSession),
	}
}
