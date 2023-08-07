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
)


type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}


func GenerateToken(name string) (string, error) {
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
		Name: name,
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
		if r.URL.Path == "/" ||  r.URL.Path == "/userlogout" ||  r.URL.Path == "/login" || r.URL.Path == "/signup" || r.URL.Path == "/adminhome" || r.URL.Path == "/adminlogout"  {
			next.ServeHTTP(w, r)
			return
		}


		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther) 
			fmt.Println(err)
			return
			
		}

		tokenString := strings.TrimSpace(cookie.Value)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			fmt.Println(claims)
			fmt.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
			
		}
	})
}
