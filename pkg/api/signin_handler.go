package api

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/golang-jwt/jwt/v5"
)

type signInRequest struct {
	Password string `json:"password"`
}

var jwtSecret = []byte("todo-secret")

func signInHandler(w http.ResponseWriter, r *http.Request) {
	var req signInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": ErrInvalidJSON.Error(),
		})
		return
	}

	pass := os.Getenv("TODO_PASSWORD")
	if pass == "" || req.Password != pass {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": ErrPassword.Error(),
		})
		return
	}

	// Формируем хэш пароля для JWT payload
	hash := sha256.Sum256([]byte(pass))
	hashStr := hex.EncodeToString(hash[:])

	// Создаём токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"hash": hashStr,
		"exp":  time.Now().Add(hourTokenLifespan * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": ErrTokenGeneration.Error(),
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   hourTokenLifespan * 3600,
		HttpOnly: true,
	})

	writeJSON(w, http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
