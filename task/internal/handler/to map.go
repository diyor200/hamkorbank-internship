package handler

import (
	"absapi/internal/repository"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func (h *Handler) GetUserFromABSMAp(persID int) (map[string]interface{}, error) {
	client := createClient()
	token, err := h.repo.GetToken()

	if errors.Is(err, repository.ErrNotFound) {
		token, err = getToken(client)
		if err != nil {
			return nil, err
		}
		token.ExpiresAt = int(time.Now().Add(time.Minute * 118).Unix())
		h.repo.InsertToken(token)
	}

	if token.ExpiresAt < int(time.Now().Unix()) {
		token, err = getToken(client)
		if err != nil {
			return nil, err
		}
		h.repo.UpdateToken(token)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://absapi.lab.hamkor.local/api/idutils/check-employer/%d", persID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token.TokenName)
	req.Header.Add("requestId", getUUID())
	req.Header.Add("lang", "EN")

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// mapga o'girish
	b, _ := io.ReadAll(r.Body)
	d, err := convertToMap(b)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode == http.StatusOK {
		log.Println("Request successful")
	} else {
		return nil, errors.New(fmt.Sprintf("Request failed with status code: %d", r.StatusCode))
	}

	return d, nil
}
