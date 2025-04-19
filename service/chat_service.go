package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-gpt/config"
	"io/ioutil"
	"net/http"
)

func GetChatGPTResponse(message string) (string, error) {
	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", config.AppConfig.OpenAIURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.AppConfig.OpenAIAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(resBody, &result); err != nil {
		return "", err
	}

	// Check for error response
	if errMsg, exists := result["error"]; exists {
		errInfo := errMsg.(map[string]interface{})
		return "", errors.New("OpenAI API error: " + errInfo["message"].(string))
	}

	choicesRaw, ok := result["choices"].([]interface{})
	if !ok || len(choicesRaw) == 0 {
		return "", errors.New("No choices returned by OpenAI")
	}

	messageObj := choicesRaw[0].(map[string]interface{})["message"].(map[string]interface{})
	content := messageObj["content"].(string)

	return content, nil
}

// func GetChatGPTResponse(message string) (string, error) {
// 	apiKey := os.Getenv("OPENAI_API_KEY")
// 	payload := map[string]interface{}{
// 		"model": "gpt-3.5-turbo",
// 		"messages": []map[string]string{
// 			{"role": "user", "content": message},
// 		},
// 	}
// 	body, _ := json.Marshal(payload)

// 	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+apiKey)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	resBody, _ := ioutil.ReadAll(resp.Body)

// 	var result map[string]interface{}
// 	json.Unmarshal(resBody, &result)

// 	choices := result["choices"].([]interface{})
// 	messageContent := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

// 	return messageContent, nil
// }
