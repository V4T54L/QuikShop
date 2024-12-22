package utils

import (
	"backend/internals/models"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const JWT_SECRET = "secret"

func GenerateJWT(data *models.TokenPayload) (string, error) {
	data.Exp = time.Now().Add(time.Minute * 5).Unix()
	// Generate JWT token
	dataStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// base64 encode the data
	payload := base64.StdEncoding.EncodeToString(dataStr)

	// Sign the payload
	signature := HashPassword(payload + JWT_SECRET)
	token := payload + "." + signature
	return token, nil
}

func VerifyJWT(data string) (*models.TokenPayload, error) {
	// Split the token
	parts := strings.Split(data, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid token")
	}

	// Verify the signature
	signature := HashPassword(parts[0] + JWT_SECRET)
	if signature != parts[1] {
		return nil, errors.New("invalid token")
	}

	// Decode the payload
	payload, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}

	// Unmarshal the payload
	var result *models.TokenPayload
	err = json.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > result.Exp {
		return nil, errors.New("token expired")
	}

	return result, nil
}
