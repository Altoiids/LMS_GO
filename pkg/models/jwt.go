package models

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"mvc/pkg/types"
	"errors"
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

	jwtsecretkey := databaseInfo.JWT_Key
			
	jwtKey := []byte(jwtsecretkey)
	
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), 
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token.SignedString(jwtKey))
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

	jwtsecretkey := databaseInfo.JWT_Key
			
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

	return nil, fmt.Errorf("Invalid token")
}

func VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathComponents := strings.Split(r.URL.Path, "/")
		firstPartOfURL := pathComponents[1]
		if r.URL.Path == "/" || r.URL.Path == "/userlogout"|| r.URL.Path == "/adminlogout" || firstPartOfURL == "static" || r.URL.Path == "/adminhome" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("jwt")
		if err != nil {
			if r.URL.Path == "/" || r.URL.Path == "/userlogout"|| r.URL.Path == "/adminlogout" || r.URL.Path == "/adminhome" || r.URL.Path == "/signup" || r.URL.Path == "/login"|| firstPartOfURL == "static"  {
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
			fmt.Println(claims)
			fmt.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			username := claims.Username
			if firstPartOfURL == "admin" {
				err := ValidateUserStatus(username, "admin")
				if err == nil {
					if r.URL.Path == "/signup" || r.URL.Path == "/login" || r.URL.Path == "/" || r.URL.Path == "/adminhome" || r.URL.Path == "/adminlogout" {
						http.Redirect(w, r, "/client/profilepage", http.StatusSeeOther)
						return
					}
					next.ServeHTTP(w, r)
				} else {
					http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
				}
			} else {
				err := ValidateUserStatus(username, "client")
				if err == nil {
					if r.URL.Path == "/signup" || r.URL.Path == "/login" || r.URL.Path == "/" || r.URL.Path == "/adminhome" || r.URL.Path == "/adminlogout" {
						http.Redirect(w, r, "/client/profilepage", http.StatusSeeOther)
						return
					}
				
						next.ServeHTTP(w, r)
					} else {
						http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
					}
				}
			}
	})
}



func ValidateUserStatus(username, Usertype string) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	var admin_id int
	
	if Usertype == "admin"{
	admin_id = 1
}
if Usertype == "client"{
	admin_id = 0
}
	var CorrectUser bool
	err = db.QueryRow(`SELECT EXISTS (SELECT 1 FROM user WHERE name=? and admin_id=?)`, username, admin_id).Scan(&CorrectUser)
	if err != nil {
		return err
	} else if CorrectUser == false {
		newError := errors.New("Wrong User Type")
		return newError
	} else {
		return nil
	}
}
