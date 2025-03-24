package auth

import (
	"bytes"
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService Service
}

func NewAuthHandler(v1 *gin.RouterGroup, authService Service) {

	handler := &authHandler{authService}

	auth := v1.Group("auth")

	//auth.POST("sign_in", handler.SignIn)

	//auth.Use(middlewares.AuthService_Sample())

	//auth.GET("", handler.GetAuth)
	auth.POST("get_token", handler.GetToken)

	user := v1.Group("user")
	user2 := v1.Group("user")
	user.Use(middlewares.AuthService())
	user.GET("/detail_byvendorcode", handler.ApiGetDetailVendor)
	user2.GET("/get_login", handler.ApiGetLogin)
	user.GET("/detail_by_name_application", handler.ApiGetDetailByApplicationName)

	wf_helper := v1.Group("wf_helper")
	wf_helper.Use(middlewares.AuthService())
	wf_helper.GET("/get_flag_by_job", handler.ApiWfGetFlagByJob)
	wf_helper.GET("/get_log_by_doc_id", handler.ApiWfGetLogById)
	wf_helper.GET("/get_attch", handler.ApiWfGetAttch)
	wf_helper.GET("/get_doc_by_activity_owner", handler.ApiWfGetDocByActivityOwner)
	wf_helper.GET("/get_flag_by_id", handler.ApiWfGetFlagById)
	wf_helper.GET("/get_next_flag", handler.ApiWfGetNextFlag)
	wf_helper.POST("/get_doc_by_potential_owner", handler.ApiWfGetDocByPotentialOwner)
	wf_helper.GET("/download_attch", handler.ApiWfDownloadAttch)
	wf_helper.DELETE("/delete_attch", handler.ApiWfDeleteAttch)

	wf_helper.POST("/save_attch", handler.ApiWfSaveAttch)
	wf_helper.POST("/get_app_header_by_doc_id", handler.ApiWfGetAppHeaderByDocId)
	wf_helper.POST("/claim_doc", handler.ApiWfClaimDoc)
	wf_helper.POST("/save_doc", handler.ApiWfSaveDoc)
}

func (h *authHandler) GetAuth(c *gin.Context) {
	auth, err := h.authService.AuthTest()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": auth,
	})
}

func (h *authHandler) SignIn(c *gin.Context) {
	var authrequest domain.AuthRequest

	err := c.ShouldBindJSON(&authrequest)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})

		return
	}

	auth, err := h.authService.SignIn(authrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": auth,
	})
}

func (h *authHandler) GetToken(c *gin.Context) {
	var authtoken domain.AuthToken

	if err := c.ShouldBindJSON(&authtoken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	_, err := middlewares.DecryptAndValidate(authtoken.Token, middlewares.JWTKEY)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    url.QueryEscape(authtoken.Token),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Domain:   os.Getenv("URL_APP"),
	}
	http.SetCookie(c.Writer, &cookie)

	vcookie := http.Cookie{
		Name:     "valid_d",
		Value:    url.QueryEscape(authtoken.ValidDetail),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Domain:   os.Getenv("URL_APP"),
	}
	http.SetCookie(c.Writer, &vcookie)

	if retrievedCookie, err := c.Cookie("session_token"); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token":   true,
			"message": "Cookie berhasil di-set",
			"cookie":  retrievedCookie,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"token":   false,
			"message": err.Error(),
		})
	}
}

// api user
func (h *authHandler) ApiGetDetailVendor(c *gin.Context) {
	vendorCode := c.DefaultQuery("vendor_code", "")
	if vendorCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "vendor_code is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%suser/detail_byvendorcode?vendor_code=%s", os.Getenv("SERVER_URL_UM"), vendorCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiGetLogin(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id_trx := c.DefaultQuery("id_trx", "")
	if id_trx == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_trx is required"})
		return
	}

	url := fmt.Sprintf("%suser/get_login?id=%s&id_trx=%s", os.Getenv("SERVER_URL_UM"), id, id_trx)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiGetDetailByApplicationName(c *gin.Context) {
	name_application := c.DefaultQuery("name_application", "")
	if name_application == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name_application is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%suser/detail_by_name_application?name_application=%s", os.Getenv("SERVER_URL_UM"), name_application)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetFlagByJob(c *gin.Context) {
	job_alias := c.DefaultQuery("job_alias", "")
	if job_alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "job_alias is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_flag_by_job?job_alias=%s", os.Getenv("SERVER_URL_WR"), job_alias)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetLogById(c *gin.Context) {
	document_id := c.DefaultQuery("document_id", "")
	if document_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "document_id is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_log_by_doc_id?document_id=%s", os.Getenv("SERVER_URL_WR"), document_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetAttch(c *gin.Context) {
	app_header_id := c.DefaultQuery("app_header_id", "")
	if app_header_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "app_header_id is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_attch?search=&app_header_id=%s", os.Getenv("SERVER_URL_WR"), app_header_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetDocByActivityOwner(c *gin.Context) {
	user_name := c.DefaultQuery("user_name", "")
	if user_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_name is required"})
		return
	}

	is_administrator := c.DefaultQuery("is_administrator", "")
	if is_administrator == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is_administrator is required"})
		return
	}

	doc_alias := c.DefaultQuery("doc_alias", "")
	if doc_alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doc_alias is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_doc_by_activity_owner?user_name=%s&is_administrator=%s&doc_alias=%s", os.Getenv("SERVER_URL_WR"), user_name, is_administrator, doc_alias)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetFlagById(c *gin.Context) {
	flag_id := c.DefaultQuery("flag_id", "")
	if flag_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "flag_id is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_flag_by_id?flag_id=%s", os.Getenv("SERVER_URL_WR"), flag_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetNextFlag(c *gin.Context) {
	activity_id := c.DefaultQuery("activity_id", "")
	if activity_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activity_id is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_next_flag?activity_id=%s", os.Getenv("SERVER_URL_WR"), activity_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetDocByPotentialOwner(c *gin.Context) {
	var requestBody []interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	role := c.DefaultQuery("role", "")
	if role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role is required"})
		return
	}

	is_administrator := c.DefaultQuery("is_administrator", "")
	if is_administrator == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is_administrator is required"})
		return
	}

	doc_alias := c.DefaultQuery("doc_alias", "")
	if doc_alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "doc_alias is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_doc_by_potential_owner?role=%s&is_administrator=%s&doc_alias=%s", os.Getenv("SERVER_URL_WR"), role, is_administrator, doc_alias)

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request body"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfDeleteAttch(c *gin.Context) {
	attch_id := c.DefaultQuery("attch_id", "")
	if attch_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attch_id is required"})
		return
	}

	deleted_by := c.DefaultQuery("deleted_by", "")
	if deleted_by == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deleted_by is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/delete_attch?attch_id=%s&deleted_by=%s", os.Getenv("SERVER_URL_WR"), attch_id, deleted_by)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfGetAppHeaderByDocId(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/get_app_header_by_doc_id", os.Getenv("SERVER_URL_WR"))

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request body"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfClaimDoc(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/claim_doc", os.Getenv("SERVER_URL_WR"))

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request body"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfSaveDoc(c *gin.Context) {
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/save_doc", os.Getenv("SERVER_URL_WR"))

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request body"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *authHandler) ApiWfSaveAttch(c *gin.Context) {
	authToken, err := c.Cookie("session_token")

	file, fileHeader, err := c.Request.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}
	defer file.Close()

	filename := fileHeader.Filename
	ext := filepath.Ext(filename)
	_json := c.DefaultPostForm("json", "")
	jsonData := []byte(_json)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// part, err := writer.CreateFormFile(ext, filename)
	// _, err = io.Copy(part, file)

	h1 := make(textproto.MIMEHeader)
	h1.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, ext, filename))
	contentType := getMimeType(filename)
	h1.Set("Content-Type", contentType)
	part, err := writer.CreatePart(h1)
	_, err = io.Copy(part, file)

	jsonPart, err := writer.CreateFormField("json")
	_, err = jsonPart.Write(jsonData)
	err = writer.Close()

	url := fmt.Sprintf("%swf_helper/save_attch", os.Getenv("SERVER_URL_WR"))
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(_body, &jsonResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func CreateFormFile(writer *multipart.Writer, fieldname, filename string) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldname, filename))

	// Tentukan Content-Type berdasarkan ekstensi file
	contentType := getMimeType(filename)
	h.Set("Content-Type", contentType)

	// Buat part dalam writer multipart
	part, err := writer.CreatePart(h)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	// Buka file dan salin isinya ke dalam part
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %w", err)
	}

	fmt.Println("File added with Content-Type:", contentType)
	return nil
}

func getMimeType(filename string) string {
	ext := filepath.Ext(filename)
	if mimeType := mime.TypeByExtension(ext); mimeType != "" {
		return mimeType
	}
	return "application/octet-stream"
}

func (h *authHandler) ApiWfDownloadAttch(c *gin.Context) {
	attch_id := c.DefaultQuery("attch_id", "")
	if attch_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "attch_id is required"})
		return
	}

	name_file := c.DefaultQuery("name_file", "")
	if name_file == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name_file is required"})
		return
	}

	authToken, err := c.Cookie("session_token")

	url := fmt.Sprintf("%swf_helper/download_attch?attch_id=%s", os.Getenv("SERVER_URL_WR"), attch_id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("authenticationToken", authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	contentLength := resp.ContentLength
	contentType := resp.Header.Get("Content-Type")
	contentDisposition := resp.Header.Get("Content-Disposition")

	// _files := strings.Split(name_file, ".")
	ext_file := name_file

	if ext_file == "xlsx" {
		contentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	} else if ext_file == "docx" {
		contentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	} else if ext_file == "pptx" {
		contentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	}

	c.Header("Content-Length", fmt.Sprintf("%d", contentLength))
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", contentDisposition)

	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stream file"})
	}
}
