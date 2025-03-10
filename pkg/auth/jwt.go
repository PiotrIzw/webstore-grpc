package auth

import (
	"errors"
	"github.com/PiotrIzw/webstore-grcp/config/config"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var jwtSecret []byte

// GenerateToken generates a JWT token for a user.
func GenerateToken(userID string) (string, error) {
	jwtSecret := getSecret()

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func getSecret() []byte {

	if len(jwtSecret) == 0 {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		jwtSecret = []byte(cfg.JWTSecret)
	}

	return jwtSecret
}

// ValidateToken validates a JWT token and returns the user ID.
func ValidateToken(tokenString string) (string, error) {

	jwtSecret := getSecret()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		return userID, nil
	}

	return "", errors.New("invalid token")
}
