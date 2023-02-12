package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KirillMironov/openai/internal/formdata"
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
	return makeJSONRequest[ModelsResponse](c, http.MethodGet, "/models", nil)
}

// Model retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *Client) Model(id string) (ModelResponse, error) {
	return makeJSONRequest[ModelResponse](c, http.MethodGet, "/models/"+id, nil)
}

// Completion creates a completion for the provided prompt and parameters.
func (c *Client) Completion(request CompletionRequest) (CompletionResponse, error) {
	return makeJSONRequest[CompletionResponse](c, http.MethodPost, "/completions", request)
}

// Edit creates a new edit for the provided input, instruction, and parameters.
func (c *Client) Edit(request EditRequest) (EditResponse, error) {
	return makeJSONRequest[EditResponse](c, http.MethodPost, "/edits", request)
}

// Image given a prompt and/or an input image, the model will generate a new image.
func (c *Client) Image(request ImageRequest) (ImageResponse, error) {
	return makeJSONRequest[ImageResponse](c, http.MethodPost, "/images/generations", request)
}

// ImageEdit creates an edited or extended image given an original image and a prompt.
func (c *Client) ImageEdit(request ImageEditRequest) (ImageEditResponse, error) {
	return makeFormDataRequest[ImageEditResponse](c, "/images/edits", request)
}

// ImageVariation creates a variation of a given image.
func (c *Client) ImageVariation(request ImageVariationRequest) (ImageVariationResponse, error) {
	return makeFormDataRequest[ImageVariationResponse](c, "/images/variations", request)
}

// Embedding creates an embedding vector representing the input text.
func (c *Client) Embedding(request EmbeddingRequest) (EmbeddingResponse, error) {
	return makeJSONRequest[EmbeddingResponse](c, http.MethodPost, "/embeddings", request)
}

// Files returns a list of files that belong to the user's organization.
func (c *Client) Files() (FilesResponse, error) {
	return makeJSONRequest[FilesResponse](c, http.MethodGet, "/files", nil)
}

// UploadFile upload a file that contains document(s) to be used across various endpoints/features.
// Currently, the size of all the files uploaded by one organization can be up to 1 GB.
func (c *Client) UploadFile(request UploadFileRequest) (UploadFileResponse, error) {
	return makeJSONRequest[UploadFileResponse](c, http.MethodPost, "/files", request)
}

func makeJSONRequest[T any](client *Client, method, path string, payload any) (T, error) {
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

	req, err := http.NewRequest(method, client.baseURL+path, body)
	if err != nil {
		return target, err
	}

	req.Header.Set("Content-Type", "application/json")

	return makeRequest[T](client, req)
}

func makeFormDataRequest[T any](client *Client, path string, payload any) (T, error) {
	var target T

	data, contentType, err := formdata.Marshal(payload)
	if err != nil {
		return target, err
	}

	req, err := http.NewRequest(http.MethodPost, client.baseURL+path, bytes.NewReader(data))
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
