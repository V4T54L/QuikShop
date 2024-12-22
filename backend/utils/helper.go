package utils

import (
	"backend/internals/models"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

const TOKEN_VALIDITY_DURATION = time.Minute * 15

func GetInt(s string) (*int, error) {
	if len(s) == 0 {
		return nil, nil
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func getSHA256(input string) string {
	// TODO : use some env variable
	input += "secret"
	hash := sha256.New()

	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}

func validateHash(payload, hashed string) bool {
	return getSHA256(payload) == hashed
}

func GetTokenString(userID int, role string) (string, error) {
	token := models.AuthToken{
		UserID: userID,
		Role:   role,
		Exp:    time.Now().Add(TOKEN_VALIDITY_DURATION),
	}
	data, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	payload := string(data)
	return payload + "." + getSHA256(payload), nil
}

func ParseTokenString(tokenStr string) (*models.AuthToken, error) {
	vals := strings.Split(tokenStr, ".")
	// TODO: Make this constant
	invalidTokenError := errors.New("invalid token")

	if len(vals) != 2 || validateHash(vals[0], vals[1]) {
		return nil, invalidTokenError
	}
	token := models.AuthToken{}
	if err := json.Unmarshal([]byte(vals[0]), &token); err != nil {
		return nil, invalidTokenError
	}
	return &token, nil
}
