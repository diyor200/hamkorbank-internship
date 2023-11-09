package handler

import (
	"absapi/internal/entity"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"time"
)

// createClient creates new https client
func createClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12, // Set a minimum TLS version
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			},
		},
	}
	return &http.Client{Transport: tr}
}

// getUUID generates new unique value
func getUUID() string {
	uuid := uuid2.New()
	return uuid.String()
}

//func (h *Handler) GetUserFromABS(persID int) (entity.Response, error) {
//	client := createClient()
//	token, err := h.repo.GetToken()
//	log.Println("data, err = ", token, err)
//
//	if errors.Is(err, repository.ErrNotFound) {
//		token, err = getToken(client)
//		if err != nil {
//			return entity.Response{}, err
//		}
//		token.ExpiresAt = int(time.Now().Add(time.Minute * 118).Unix())
//		h.repo.InsertToken(token)
//	}
//
//	if token.ExpiresAt < int(time.Now().Unix()) {
//		log.Println("data.ExpiresAt, int(time.Now().Unix()) = ", token.ExpiresAt, int(time.Now().Unix()))
//		token, err = getToken(client)
//		if err != nil {
//			return entity.Response{}, err
//		}
//		h.repo.UpdateToken(token)
//	}
//
//	req, err := http.NewRequest("GET", fmt.Sprintf("https://absapi.lab.hamkor.local/api/idutils/check-employer/%d", persID), nil)
//	if err != nil {
//		return entity.Response{}, err
//	}
//	req.Header.Add("Authorization", "Bearer "+token.TokenName)
//	req.Header.Add("requestId", getUUID())
//	req.Header.Add("lang", "EN")
//
//	r, err := client.Do(req)
//	if err != nil {
//		return entity.Response{}, err
//	}
//
//	// jsonga o'girish
//	//var responseData entity.Response
//	//if err = json.NewDecoder(r.Body).Decode(&responseData); err != nil {
//	//	fmt.Println("<+++>", err)
//	//	return entity.Response{}, err
//	//}
//
//	// mapga o'girish
//	b, _ := io.ReadAll(r.Body)
//	d, err := convertToMap(b)
//	if err != nil {
//		return entity.Response{}, err
//	}
//	if h.repo.InsertUserFromMap(d) {
//		log.Println("success")
//	}
//
//	defer r.Body.Close()
//
//	if r.StatusCode == http.StatusOK {
//		// Request was successful. You can parse the response body here.
//		log.Println("Request successful")
//	} else {
//		// Request failed. Handle the error or response accordingly.
//		return entity.Response{}, errors.New(fmt.Sprintf("Request failed with status code: %d", r.StatusCode))
//	}
//
//	if len(responseData.ResponseBody) == 0 {
//		return entity.Response{}, errors.New("user not exists")
//	}
//
//	return responseData, nil
//}

// getToken gets new token from API
func getToken(client *http.Client) (*entity.Token, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	body := []byte(`{"filialCode": "09012", "lang": "EN","password": "pdp_hr","username": "pdp_hr"}`)
	response, err := client.Post("https://absapi.lab.hamkor.local/api/auth/get-token", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return &entity.Token{}, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return &entity.Token{}, err
	}
	resp := string(responseBody)
	l := len(resp)

	return &entity.Token{TokenName: resp[10 : l-2], ExpiresAt: int(time.Now().Add(time.Minute * 118).Unix())}, nil
}

// convertToMap converts json data to map
func convertToMap(jsonData []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON:", err)
	}

	// Check if "msg" is "Success" and access the desired data
	msg, msgExists := data["msg"]
	if msgExists && msg.(string) == "Success" {
		responseBody, bodyExists := data["responseBody"].([]interface{})
		if bodyExists && len(responseBody) > 0 {
			// Extract the first element of "responseBody" as a map
			extractedData, ok := responseBody[0].(map[string]interface{})
			if ok {
				return extractedData, nil
			} else {
				return nil, fmt.Errorf("invalid format for responseBody")
			}
		} else {
			return nil, fmt.Errorf("responseBody is empty or missing")
		}
	} else {
		return nil, fmt.Errorf("user not exists")
	}
}
