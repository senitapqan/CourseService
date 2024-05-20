package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func SendNotificationRequest(host, port, title string, followers []string) error {
	url := "/send"

	payload := gin.H{
		"emails": followers,
		"title":title,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(host+":"+port+url, binding.MIMEJSON, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("аутентификация не удалась, статус код: " + resp.Status)
	}

	return nil
}