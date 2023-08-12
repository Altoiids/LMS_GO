package models

import (
	"errors"
	"log"
	"mvc/pkg/types"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var databaseInfo types.DBInfo

	if err := yaml.Unmarshal(file, &databaseInfo); err != nil {
		log.Fatal(err)
	}

	jwtsecretkey := databaseInfo.JWT_KEY

	jwtKey := []byte(jwtsecretkey)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)

}

func VerifyToken(tokenString string) (*Claims, error) {

	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var databaseInfo types.DBInfo

	if err := yaml.Unmarshal(file, &databaseInfo); err != nil {
		log.Fatal(err)
	}

	jwtsecretkey := databaseInfo.JWT_KEY

	jwtKey := []byte(jwtsecretkey)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

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
		claims, err := VerifyToken(tokenString)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			username := claims.Username
			if firstPartOfURL == "admin" {
				err := ValidateUserStatus(username, "admin")
				if err == nil {
					if firstPartOfURL == "admin" {
						next.ServeHTTP(w, r)
						return
					} else {
						http.Redirect(w, r, "/adminHome", http.StatusSeeOther)
					}
				}
			} else {
				err := ValidateUserStatus(username, "client")
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

func ValidateUserStatus(username, userType string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	var adminId int

	if userType == "admin" {
		adminId = 1
	}
	if userType == "client" {
		adminId = 0
	}
	var CorrectUser bool
	err = db.QueryRow(`SELECT EXISTS (SELECT 1 FROM user WHERE name=? and adminId=?)`, username, adminId).Scan(&CorrectUser)
	if err != nil {
		return err
	} else if CorrectUser == false {
		newError := errors.New("Wrong User Type")
		return newError
	} else {
		return nil
	}
}
