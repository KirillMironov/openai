package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KirillMironov/openai/internal/formdata"
)

func makeJSONRequest[T any](ctx context.Context, client *Client, method, path string, payload any) (T, error) {
	var (
		target T
		body   io.Reader
	)

	if payload != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			return target, err
		}

		body = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+path, body)
	if err != nil {
		return target, err
	}

	req.Header.Set("Content-Type", "application/json")

	return makeRequest[T](client, req)
}

func makeFormDataRequest[T any](ctx context.Context, client *Client, path string, payload any) (T, error) {
	var target T

	data, contentType, err := formdata.Marshal(payload)
	if err != nil {
		return target, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, client.baseURL+path, bytes.NewReader(data))
	if err != nil {
		return target, err
	}

	req.Header.Set("Content-Type", contentType)

	return makeRequest[T](client, req)
}

func makeRequest[T any](client *Client, req *http.Request) (T, error) {
	var target T

	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	if client.organization != "" {
		req.Header.Set("OpenAI-Organization", client.organization)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return target, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return target, err
	}

	if resp.StatusCode != http.StatusOK {
		var openaiApiResponse struct {
			Error struct {
				Message string `json:"message"`
				Type    string `json:"type"`
			} `json:"error"`
		}

		if err = json.Unmarshal(respData, &openaiApiResponse); err != nil {
			return target, Error{
				StatusCode: resp.StatusCode,
				Message:    string(respData),
				Type:       "unknown",
			}
		}

		return target, Error{
			StatusCode: resp.StatusCode,
			Message:    openaiApiResponse.Error.Message,
			Type:       openaiApiResponse.Error.Type,
		}
	}

	return target, json.Unmarshal(respData, &target)
}
