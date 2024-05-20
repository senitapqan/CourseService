package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"goCourseService/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) sendParseTokenRequest(host, port, token string) (dtos.User, error) {
	var result dtos.User
	url := "/internal-api/parse-token"

	payload := gin.H{
		"token": token,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return dtos.User{}, err
	}

	resp, err := http.Post(host+":"+port+url, binding.MIMEJSON, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return dtos.User{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return dtos.User{}, errors.New("аутентификация не удалась, статус код: " + resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return dtos.User{}, err
	}

	return result, nil
}
