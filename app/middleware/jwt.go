package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func ValidateToken(token string, secret string) (bool, error) {
	// JWT has 3 parts separated by '.'
	splitToken := strings.Split(token, ".")
	// if length is not 3, we know that the token is corrupt
	if len(splitToken) != 3 {
		return false, nil
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err

	}

	//again create the signature
	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(signature)

	// if both the signature don't match, this means token is wrong
	if signature != splitToken[2] {
		return false, nil
	}
	// This means the token matches
	return true, nil
}
