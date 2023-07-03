package auth

import (
	"bytes"
	"ebapp-api-dev/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Repository interface {
	AuthTest() (domain.Auth, error)
	SignIn(auth domain.Auth) (domain.AuthResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AuthTest() (domain.Auth, error) {
	auth := domain.Auth{
		Username: "test username",
		Password: "test password",
	}

	return auth, nil
}

func (r *repository) SignIn(auth domain.Auth) (domain.AuthResponse, error) {
	authInput, err := json.Marshal(auth)
	if err != nil {
		return domain.AuthResponse{}, err
	}

	url := os.Getenv("USER_SERVICE_URL") + "/auth/sign_in"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(authInput)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return domain.AuthResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result domain.AuthResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
