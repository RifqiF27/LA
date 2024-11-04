package utils

import (
    "context"
    "time"
    "math/rand"
    "fmt"
)

type Session struct {
    Token     string
    ExpiresAt time.Time
}

var sessions = make(map[string]Session)

func GenerateToken() string {
    return fmt.Sprintf("%d", rand.Int())
}

func CreateSession(username string, duration time.Duration) (string, context.Context, context.CancelFunc) {
    token := GenerateToken()
    ctx, cancel := context.WithTimeout(context.Background(), duration)

    sessions[token] = Session{
        Token:     token,
        ExpiresAt: time.Now().Add(duration),
    }
    return token, ctx, cancel
}

func ValidateSession(token string) bool {
    session, exists := sessions[token]
    if !exists {
        return false
    }
    if time.Now().After(session.ExpiresAt) {
        delete(sessions, token)
        return false
    }
    return true
}

func EndSession(token string) {
    delete(sessions, token)
}
