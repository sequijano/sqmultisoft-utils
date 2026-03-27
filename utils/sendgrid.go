package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type sgEmail struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type sgPersonalization struct {
	To                  []sgEmail              `json:"to"`
	DynamicTemplateData map[string]interface{} `json:"dynamic_template_data,omitempty"`
}

type sgContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type sgMessage struct {
	From             sgEmail             `json:"from"`
	Personalizations []sgPersonalization `json:"personalizations"`
	Subject          string              `json:"subject,omitempty"`
	Content          []sgContent         `json:"content,omitempty"`
	TemplateID       string              `json:"template_id,omitempty"`
}

func sendGridRequest(apiKey string, msg sgMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid error %d: %s", resp.StatusCode, string(respBody))
	}
	return nil
}

func SendGridTemplate(apiKey, from, to, templateID string, data map[string]interface{}) error {
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY no configurado")
	}
	return sendGridRequest(apiKey, sgMessage{
		From: sgEmail{Email: from, Name: "Multi-Currency POS"},
		Personalizations: []sgPersonalization{
			{To: []sgEmail{{Email: to}}, DynamicTemplateData: data},
		},
		TemplateID: templateID,
	})
}

func SendGridPlain(apiKey, from, to, subject, body string) error {
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY no configurado")
	}
	return sendGridRequest(apiKey, sgMessage{
		From: sgEmail{Email: from, Name: "Multi-Currency POS"},
		Personalizations: []sgPersonalization{
			{To: []sgEmail{{Email: to}}},
		},
		Subject: subject,
		Content: []sgContent{{Type: "text/plain", Value: body}},
	})
}
