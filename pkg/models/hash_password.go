
package models

import (
"golang.org/x/crypto/bcrypt"
"crypto/rand"
"encoding/base64"
)


func HashPassword(password string) (string, error) {
    const saltRounds = 10
    saltBytes := make([]byte, 0, 16)
    _, err := rand.Read(saltBytes)
    if err != nil {
        return "", err
    }
    salt := base64.URLEncoding.EncodeToString(saltBytes)
    bytes, err := bcrypt.GenerateFromPassword([]byte(password + salt), saltRounds)
	hash := string(bytes)
    return  hash,err
}