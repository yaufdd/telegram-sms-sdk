package telegram_sms_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

type (
	RegisterBotRequest struct {
		TelegramID int64  `json:"telegram_id"`
		Token      string `json:"token"`
	}

	RegisterBotResponse struct {
		UUID string `json:"uuid"`
		Hash string `json:"hash"`
	}
)

func (c *Client) RegisterBot(telegramID int64, token string) (RegisterBotResponse, error) {
	req := RegisterBotRequest{
		TelegramID: telegramID,
		Token:      token,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return RegisterBotResponse{}, err
	}

	resp, err := c.httpClient.Post(
		c.baseURL+"/register-bot",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return RegisterBotResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RegisterBotResponse{}, fmt.Errorf("%d", resp.StatusCode)
	}

	var response RegisterBotResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return RegisterBotResponse{}, err
	}

	return response, nil
}

type (
	SendMessageRequest struct {
		UUID    string `json:"uuid"`
		Hash    string `json:"hash"`
		Message string `json:"message"`
	}

	SendMessageResponse struct {
		Status string `json:"status"`
	}
)

func (c *Client) SendMessage(uuid string, hash string, receiverTgID int64, text string) (SendMessageResponse, error) {
	req := SendMessageRequest{
		UUID:    uuid,
		Hash:    hash,
		Message: text,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshalling request:", err)
		return SendMessageResponse{}, err
	}

	resp, err := c.httpClient.Post(
		c.baseURL+"/send-message",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return SendMessageResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SendMessageResponse{}, fmt.Errorf("%d", resp.StatusCode)
	}

	var response SendMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return SendMessageResponse{}, err
	}

	return response, nil
}
