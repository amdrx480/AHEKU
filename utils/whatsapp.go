package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendWhatsAppMessage mengirimkan pesan WhatsApp menggunakan endpoint yang diberikan.
func SendWhatsAppMessage(chatID, text, session string) error {
	// Buat payload JSON berdasarkan struktur yang diharapkan oleh endpoint
	payload := struct {
		ChatID  string `json:"chatId"`
		Text    string `json:"text"`
		Session string `json:"session"`
	}{
		ChatID:  chatID,
		Text:    text,
		Session: session,
	}

	// Marshal payload ke JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %v", err)
	}

	// Buat request HTTP POST
	req, err := http.NewRequest("POST", "http://localhost:3000/api/sendText", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Atur header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Periksa status response
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// log.Printf("WhatsApp message sent to %s: %s\n", chatID, text)

	return nil
}

// type WhatsAppMessage struct {
// 	ChatID  string `json:"chatId"`
// 	Text    string `json:"text"`
// 	Session string `json:"session"`
// }

// func SendWhatsAppMessage(phone, message string) error {
// 	url := "http://localhost:3000/api/sendText"

// 	whatsAppMessage := WhatsAppMessage{
// 		ChatID:  phone + "@c.us",
// 		Text:    message,
// 		Session: "default",
// 	}

// 	jsonData, err := json.Marshal(whatsAppMessage)
// 	if err != nil {
// 		return fmt.Errorf("error marshaling json: %v", err)
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("error creating new request: %v", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("error sending request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to send WhatsApp message, status code: %d", resp.StatusCode)
// 	}

// 	return nil
// }

// type WhatsAppMessage struct {
// 	ChatID  string `json:"chatId"`
// 	Text    string `json:"text"`
// 	Session string `json:"session"`
// }

// func SendWhatsAppMessage(adminPhone, officerPhone, message string) error {
// 	url := "http://localhost:3000/api/sendText"

// 	payload := WhatsAppMessage{
// 		ChatID:  officerPhone + "@c.us",
// 		Text:    message,
// 		Session: "default",
// 	}

// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Printf("Failed to send WhatsApp message: %s", resp.Status)
// 	}

// 	return nil
// }
