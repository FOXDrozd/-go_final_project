package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if pass == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			hash, ok := claims["hash"].(string)
			if !ok {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}

			currentHash := sha256.Sum256([]byte(pass))
			if hash != hex.EncodeToString(currentHash[:]) {
				http.Error(w, "Authentication required", http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
