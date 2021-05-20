package session

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/syllabix/rollpay/backend/config"
)

// TODO: move to config
const (
	maxAttempts   = 10
	sessionLength = time.Hour
	tokenName     = "r.sesh"
)

type tokenData struct {
	ID int64
}

type Token struct {
	Value      string
	Expiration time.Time
}

type UserSession struct {
	ID         int64
	UserID     int64
	Expiration time.Time
}

type Manager struct {
	mutex    sync.RWMutex
	sessions map[int64]UserSession

	current  *securecookie.SecureCookie
	previous *securecookie.SecureCookie
}

// Create a new session for the provided user id
func (m *Manager) Create(userID int64) (Token, error) {
	id, err := m.newID()
	if err != nil {
		return Token{}, fmt.Errorf("failed to create a session: %w", err)
	}

	sesh := UserSession{
		ID:         id,
		UserID:     userID,
		Expiration: time.Now().Add(sessionLength),
	}
	m.set(sesh)

	token, err := m.current.Encode(tokenName, tokenData{id})
	if err != nil {
		return Token{}, fmt.Errorf("failed encode session id: %w", err)
	}

	return Token{
		Value:      token,
		Expiration: sesh.Expiration,
	}, nil
}

func (m *Manager) Get(token string) (session UserSession, exists bool) {
	sessionID, err := m.decode(token)
	if err != nil {
		return session, false
	}
	m.mutex.RLock()
	session, exists = m.sessions[sessionID]
	m.mutex.RUnlock()

	return session, exists
}

func (m *Manager) get(sessionID int64) (session UserSession, exists bool) {
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

func (m *Manager) decode(token string) (id int64, err error) {
	var data tokenData
	err = securecookie.DecodeMulti(tokenName, token, &data, m.current, m.previous)
	if err != nil {
		return id, err
	}

	return data.ID, nil
}

func (m *Manager) newID() (int64, error) {
	for i := 0; i < maxAttempts; i++ {
		id, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
		if err != nil {
			return 0, fmt.Errorf("unable to create session id: %w", err)
		}

		_, exists := m.get(id.Int64())
		if !exists {
			return id.Int64(), nil
		}
	}
	return 0, errors.New("unable to create a unique session id")
}

func NewManager(settings config.SecuritySettings) *Manager {
	return &Manager{
		sessions: make(map[int64]UserSession),
		current: securecookie.New(
			settings.CurrentTokenKey.HashKey,
			settings.CurrentTokenKey.BlockKey),
		previous: securecookie.New(
			settings.PreviousTokenKey.HashKey,
			settings.PreviousTokenKey.BlockKey),
	}
}
