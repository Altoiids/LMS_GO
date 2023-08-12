package jwt_middleware_handler

import (
	"mvc/pkg/models"
	"net/http"
	"strings"
)

func VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathComponents := strings.Split(r.URL.Path, "/")
		firstPartOfURL := pathComponents[1]
		if r.URL.Path == "/" || r.URL.Path == "/userLogout" || r.URL.Path == "/adminLogout" || firstPartOfURL == "static" || r.URL.Path == "/adminHome" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if r.URL.Path == "/" || r.URL.Path == "/userLogout" || r.URL.Path == "/adminLogout" || r.URL.Path == "/adminHome" || r.URL.Path == "/signUp" || r.URL.Path == "/login" || firstPartOfURL == "static" {
				next.ServeHTTP(w, r)
				return
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		tokenString := strings.TrimSpace(cookie.Value)
		claims, err := models.VerifyToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			username := claims.Username
			if firstPartOfURL == "admin" {
				err := models.ValidateUserStatus(username, "admin")
				if err == nil {
					if firstPartOfURL == "admin" {
						next.ServeHTTP(w, r)
						return
					} else {
						http.Redirect(w, r, "/adminHome", http.StatusSeeOther)
					}
				}
			} else {
				err := models.ValidateUserStatus(username, "client")
				if err == nil {
					if firstPartOfURL == "client" {
						next.ServeHTTP(w, r)
						return
					} else {
						http.Redirect(w, r, "/", http.StatusSeeOther)
					}
				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		}
	})
}
