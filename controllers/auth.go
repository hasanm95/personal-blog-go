package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"personal-blog/types"
)

// Predefined admin credentials
var adminUser = types.User{
	Username: "admin",
	Password: "password123",
}

// Simple session store (in production, use proper session management)
var sessions = make(map[string]bool)

// AuthMiddleware checks if user is authenticated
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || !sessions[cookie.Value] {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

// LoginHandler handles login form submission
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUser.Username && password == adminUser.Password {
			// Create session
			sessionID := generateSessionID()
			sessions[sessionID] = true

			// Set cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    sessionID,
				Path:     "/",
				HttpOnly: true,
			})

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		// Login failed
		http.Redirect(w, r, "/login?error=1", http.StatusSeeOther)
		return
	}
}

// LogoutHandler handles logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		delete(sessions, cookie.Value)
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Generate secure session ID using crypto/rand
func generateSessionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
