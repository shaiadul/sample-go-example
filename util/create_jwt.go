package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"` // user id
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	Header := Header{
		Algorithm: "HS256",
		Type:      "JWT",
	}

	byteArrHeader, err := json.Marshal(Header)
	if err != nil {
		return "", err
	}
	headerB64 := base64Encode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64Encode(byteArrData)

	message := headerB64 + "." + payloadB64
	byteArrMessage := []byte(message)

	byteArrSecret := []byte(secret)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64Encode(signature)

	jwt := message + "." + signatureB64
	return jwt, nil
}

func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
