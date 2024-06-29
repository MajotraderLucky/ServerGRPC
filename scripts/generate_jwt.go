package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// generateJWT создает новый JWT токен.
func generateJWT() string {
	// Секретный ключ, используемый для подписи токена.
	var mySigningKey = []byte("your_secret_key_here")

	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Создаем map для хранения утверждений токена
	claims := token.Claims.(jwt.MapClaims)

	// Устанавливаем утверждения
	claims["authorized"] = true
	claims["user_id"] = "some_user_id"                         // Подставьте нужное значение
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // Токен истекает через 30 дней

	// Подписываем токен нашим секретным ключом
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatalf("Error in JWT token generation: %s", err)
	}

	return tokenString
}

func main() {
	jwtToken := generateJWT()
	log.Println("Generated JWT:", jwtToken)
}
