package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

func GenerateToken(header string, payload map[string]interface{}, secret string) (string, error) {
	// Encode the header using base64url encoding
	headerBytes := []byte(header)
	header64 := base64.RawURLEncoding.EncodeToString(headerBytes)

	// Encode the payload using base64url encoding
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payload64 := base64.RawURLEncoding.EncodeToString(payloadBytes)

	// Create the message by concatenating header and payload with a period ('.')
	message := header64 + "." + payload64

	// Create a new hash of type sha256 with the secret key
	h := hmac.New(sha256.New, []byte(secret))

	// Write the message to the hash
	h.Write([]byte(message))

	// Get the hash sum and encode it using base64url encoding
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	// Combine the encoded header, payload, and signature with periods to form the token
	tokenStr := header64 + "." + payload64 + "." + signature

	return tokenStr, nil

}

func GenerateRefreshToken(userID uint, secretKey string) (string, error) {
	// Define the header for the refresh token (you can customize this)
	refreshHeader := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerBytes, err := json.Marshal(refreshHeader)

	if err != nil {
		return "Error in Header Bytes Generation", err
	}

	// Define the payload for the refresh token (you can customize this)
	refreshPayload := map[string]interface{}{
		"userID": userID,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(), // Expiration time (e.g., 7 days)
	}

	// Call the GenerateToken function to generate the refresh token
	refreshToken, err := GenerateToken(string(headerBytes), refreshPayload, secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
