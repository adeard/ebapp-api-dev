package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func ValidateJWT(context *gin.Context) error {
	//token, err := getToken(context)

	// if err != nil {
	// 	return err
	// }

	// _, ok := token.Claims.(jwt.MapClaims)
	// if ok && token.Valid {
	// 	return nil
	// }
	// if !ok || !token.Valid {
	// 	return err
	// }

	return errors.New("invalid token provided")
}

func validateToken(context *gin.Context) (jwt.MapClaims, error) {
	//check validasi detail yang ada di frontend
	validDetail := getValidDetailFromRequest(context)
	hash := sha256.New()
	hash.Write([]byte(validDetail))
	hashedData := hash.Sum(nil)
	hashString := hex.EncodeToString(hashedData)
	println(hashString)
	tokenString := getTokenFromRequest(context)
	//dekrip
	tokenDecrypt, err := Decrypt(tokenString, JWTKEY)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(string(tokenDecrypt), func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}

		return []byte(JWTKEY), nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiry := claims["expiry"].(string)
		expiryTime, err := time.Parse(time.RFC3339Nano, expiry)
		if err != nil {
			fmt.Println("Error parsing expiry time:", err)
			return nil, err
		}
		if time.Now().After(expiryTime) {
			fmt.Println("Token has expired")
			return nil, fmt.Errorf("Token has expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}

func getTokenFromRequest(context *gin.Context) string {
	if retrievedCookie, err := context.Cookie("session_token"); err == nil {
		return retrievedCookie
	} else {
		return ""
	}
}

func getValidDetailFromRequest(context *gin.Context) string {
	if retrievedCookie, err := context.Cookie("valid_d"); err == nil {
		return retrievedCookie
	} else {
		return ""
	}
}

func getTokenFromRequestSample(context *gin.Context) (string, error) {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1], nil
	}
	return "", errors.New("invalid token provided")
}

func GenerateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().UTC().Add(time.Hour * 1).Unix() // Token valid for 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(privateKey)
}

const (
	KeySize        = 32
	NonceSize      = 12
	TagSize        = 16
	Version   byte = 0x1
	JWTKEY         = "1YfnFYwdFp09AejiV2gvfU5+WRjn8exPCMP6KJPgssc="
)

func Decrypt(ciphertextB64, keyB64 string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return "", err
	}

	if len(key) != KeySize {
		return "", errors.New("invalid key size")
	}

	plaintext, err := decryptAESGCM(ciphertext, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func DecryptAndValidate(ciphertextB64, keyB64 string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", err
	}

	key, err := base64.StdEncoding.DecodeString(keyB64)
	if err != nil {
		return "", err
	}

	if len(key) != KeySize {
		return "", errors.New("invalid key size")
	}

	plaintext, err := decryptAESGCM(ciphertext, key)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(string(plaintext), func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}

		return []byte(JWTKEY), nil
	})

	// Check for errors
	if err != nil {
		return "", err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expiry := claims["expiry"].(string)
		expiryTime, err := time.Parse(time.RFC3339Nano, expiry)
		if err != nil {
			fmt.Println("Error parsing expiry time:", err)
			return "", err
		}
		if time.Now().After(expiryTime) {
			fmt.Println("Token has expired")
			return "", fmt.Errorf("Token has expired")
		}
		return "Claims extracted successfully", nil
	}
	return "", fmt.Errorf("Invalid token")
}

func decryptAESGCM(ciphertext, key []byte) ([]byte, error) {
	// if ciphertext[0] != Version {
	// 	return nil, errors.New("unsupported encrypted format or version mismatch")
	// }

	offset := 1
	nonce := ciphertext[offset : offset+NonceSize]
	offset += NonceSize
	tag := ciphertext[offset : offset+TagSize]
	offset += TagSize
	encText := ciphertext[offset:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	fullCiphertext := append(encText, tag...)

	plaintext, err := aesgcm.Open(nil, nonce, fullCiphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return plaintext, nil
}
