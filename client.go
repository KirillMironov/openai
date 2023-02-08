package openai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.openai.com/v1"

type Client struct {
	apiKey       string
	organization string
	baseURL      string
	httpClient   *http.Client
}

func NewClient(apiKey string, options ...ClientOption) *Client {
	client := &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// Models lists the currently available models,
// and provides basic information about each one such as the owner and availability.
func (c *Client) Models() (ModelsResponse, error) {
	return makeRequest[ModelsResponse](c, http.MethodGet, "/models", nil)
}

// Model retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *Client) Model(model string) (ModelResponse, error) {
	return makeRequest[ModelResponse](c, http.MethodGet, "/models/"+model, nil)
}

// Completion creates a completion for the provided prompt and parameters.
func (c *Client) Completion(request CompletionRequest) (CompletionResponse, error) {
	return makeRequest[CompletionResponse](c, http.MethodPost, "/completions", request)
}

func makeRequest[T any](client *Client, method, path string, body any) (T, error) {
	var target T

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return target, err
	}

	req, err := http.NewRequest(method, client.baseURL+path, buf)
	if err != nil {
		return target, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	if client.organization != "" {
		req.Header.Set("OpenAI-Organization", client.organization)
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return target, err
	}
	defer resp.Body.Close()

	return target, json.NewDecoder(resp.Body).Decode(&target)
}