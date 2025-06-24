package jwt

import (
	"chat-service/configs"
	"errors"

	_jwt "github.com/golang-jwt/jwt/v5"
)

var conf = configs.GetJWTConfig()

type JWTPayload struct {
	UserID   string   `json:"user_id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

func ValidateToken(tokenString string) (*JWTPayload, error) {
	// Define a map to extract custom claims
	claims := _jwt.MapClaims{}

	// Parse and validate the token
	_, err := _jwt.ParseWithClaims(tokenString, claims, func(token *_jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC and specifically HS256
		if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(conf.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract and validate expected fields
	userID, ok1 := claims["user_id"].(string)
	email, ok2 := claims["email"].(string)
	username, _ := claims["username"].(string)
	rolesInterface, ok3 := claims["roles"].([]interface{})

	if !ok1 || !ok2 || !ok3 {
		return nil, errors.New("invalid token payload structure")
	}

	// Convert []interface{} to []string
	roles := make([]string, len(rolesInterface))
	for i, v := range rolesInterface {
		if str, ok := v.(string); ok {
			roles[i] = str
		} else {
			return nil, errors.New("invalid role type in token")
		}
	}

	return &JWTPayload{
		UserID:   userID,
		Email:    email,
		Username: username,
		Roles:    roles,
	}, nil
}
