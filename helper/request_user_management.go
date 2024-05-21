package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDataFromUserManagement() (string, string, string, error) {
	// POST REQUEST UNTUK LOGIN
	loginURL := "http://10.126.20.217:9087/api/user/login"
	loginData := map[string]string{"user_name": "ebapphelper", "password": "ebapphelper"}
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return "", "", "", err
	}

	client := &http.Client{}
	loginReq, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", "", err
	}
	loginReq.Header.Set("Content-Type", "application/json")

	loginResp, err := client.Do(loginReq)
	if err != nil {
		return "", "", "", err
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		return "", "", "", fmt.Errorf("login failed: %s", loginResp.Status)
	}

	var loginResponse struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
		Datas   string `json:"datas"`
	}

	err = json.NewDecoder(loginResp.Body).Decode(&loginResponse)
	if err != nil {
		return "", "", "", err
	}

	if !loginResponse.Result {
		return "", "", "", fmt.Errorf("login failed: %s", loginResponse.Message)
	}

	firstToken := loginResponse.Datas

	// Fungsi untuk melakukan GET request dan mengembalikan nilai dari objek->Value
	getValue := func(url string) (string, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("authenticationToken", firstToken)

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		var response struct {
			Result bool `json:"result"`
			Objek  struct {
				Value string `json:"Value"`
			} `json:"objek"`
			Message string `json:"message"`
		}

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return "", err
		}

		if !response.Result {
			return "", fmt.Errorf("get failed: %s", response.Message)
		}

		return response.Objek.Value, nil
	}

	// Get data URL, Username, and Password
	urlValue, err := getValue("http://10.126.20.217:9087/api/master_datas/get?value_table=SAPHelper&value_column=Url&app_root_name=BAPP")
	if err != nil {
		return "", "", "", err
	}

	usernameValue, err := getValue("http://10.126.20.217:9087/api/master_datas/get?value_table=SAPHelper&value_column=Username&app_root_name=BAPP")
	if err != nil {
		return "", "", "", err
	}

	passwordValue, err := getValue("http://10.126.20.217:9087/api/master_datas/get?value_table=SAPHelper&value_column=Password&app_root_name=BAPP")
	if err != nil {
		return "", "", "", err
	}

	return urlValue, usernameValue, passwordValue, nil
}
