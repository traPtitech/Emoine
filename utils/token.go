package utils

import (
	"errors"
	"strings"
)

const TOKEN_PREFIX = "emo:"

// ## Tokenの構造
// - prefix(4文字): "emo:"固定
// - random(40文字)
// 計44文字

func GenerateToken() (string, error) {
	token, err := SecureRandomString(40)
	if err != nil {
		return "", err
	}

	return TOKEN_PREFIX + token, nil
}

func ParseToken(token string) (string, error) {
	if !strings.HasPrefix(token, TOKEN_PREFIX) {
		return "", errors.New("token: Invalid prefix")
	}

	return strings.TrimPrefix(token, TOKEN_PREFIX), nil
}
