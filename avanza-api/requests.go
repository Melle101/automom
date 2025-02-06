package avanza_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/types"
	"io"
	"net/http"
	"reflect"
	"strings"

	api_models "github.com/Melle101/mom-bot-v3/avanza-api/api-models"
)

func fromBytes[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}

// HTTPGet makes a GET request with optional headers and unmarshals the JSON response into a generic type T.
func HTTPGet[T any](url string, headers map[string][]string) (*T, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header = api_models.DEF_HEADERS
	// Set custom headers
	for key, value := range headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result T

	if types.BasicKind(reflect.TypeOf(result).Kind()) == types.BasicKind(reflect.String) {
		result = any(string(body)).(T)
	} else {
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &result, nil
}

// HTTPPost makes a POST request with a JSON payload and headers, and unmarshals the response into a generic type T.
func HTTPPost[T any, U any](url string, payload U, headers map[string][]string) (*T, error) {
	client := &http.Client{}
	jsonDataPre, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	jsonData, _ := lowercaseJSONKeys(jsonDataPre)
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set default headers
	req.Header = api_models.DEF_HEADERS
	// Set custom headers
	for key, value := range headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

func HTTPPostHeaders[T any, U any](url string, payload U, headers map[string][]string) (*T, http.Header, error) {
	client := &http.Client{}
	jsonDataPre, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	jsonData, _ := lowercaseJSONKeys(jsonDataPre)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Set default headers
	req.Header = api_models.DEF_HEADERS
	// Set custom headers
	for key, value := range headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, resp.Header, nil
}

// HTTPDelete makes a DELETE request with optional headers and unmarshals the JSON response into a generic type T.
func HTTPDelete[T any](url string, headers map[string][]string) (*T, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	req.Header = api_models.DEF_HEADERS
	// Set custom headers
	for key, value := range headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return &result, nil
}

// makeFirstLetterLower converts the first letter of a string to lowercase.
func makeFirstLetterLower(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// lowercaseKeysRecursively processes JSON keys recursively.
func lowercaseKeysRecursively(data any) any {
	switch v := data.(type) {
	case map[string]any:
		newMap := make(map[string]any)
		for key, value := range v {
			newKey := makeFirstLetterLower(key)
			newMap[newKey] = lowercaseKeysRecursively(value) // Recursively process values
		}
		return newMap
	case []any:
		for i, item := range v {
			v[i] = lowercaseKeysRecursively(item) // Recursively process array elements
		}
	}
	return data
}

// lowercaseJSONKeys takes raw JSON and modifies keys recursively.
func lowercaseJSONKeys(input []byte) ([]byte, error) {
	var original any
	if err := json.Unmarshal(input, &original); err != nil {
		return nil, err
	}

	modified := lowercaseKeysRecursively(original)
	return json.Marshal(modified)
}
