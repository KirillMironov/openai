package openai

import (
	"bytes"
	"context"
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
func (c *Client) Models(ctx context.Context) (ModelsResponse, error) {
	return makeJSONRequest[ModelsResponse](ctx, c, http.MethodGet, "/models", nil)
}

// Model retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *Client) Model(ctx context.Context, id string) (Model, error) {
	return makeJSONRequest[Model](ctx, c, http.MethodGet, "/models/"+id, nil)
}

// DeleteModel delete a fine-tuned model.
// You must have the Owner role in your organization.
func (c *Client) DeleteModel(ctx context.Context, id string) (DeleteModelResponse, error) {
	return makeJSONRequest[DeleteModelResponse](ctx, c, http.MethodDelete, "/models/"+id, nil)
}

// Completion creates a completion for the provided prompt and parameters.
func (c *Client) Completion(ctx context.Context, request CompletionRequest) (CompletionResponse, error) {
	return makeJSONRequest[CompletionResponse](ctx, c, http.MethodPost, "/completions", request)
}

// Edit creates a new edit for the provided input, instruction, and parameters.
func (c *Client) Edit(ctx context.Context, request EditRequest) (EditResponse, error) {
	return makeJSONRequest[EditResponse](ctx, c, http.MethodPost, "/edits", request)
}

// Image given a prompt and/or an input image, the model will generate a new image.
func (c *Client) Image(ctx context.Context, request ImageRequest) (ImageResponse, error) {
	return makeJSONRequest[ImageResponse](ctx, c, http.MethodPost, "/images/generations", request)
}

// ImageEdit creates an edited or extended image given an original image and a prompt.
func (c *Client) ImageEdit(ctx context.Context, request ImageEditRequest) (ImageEditResponse, error) {
	return makeFormDataRequest[ImageEditResponse](ctx, c, "/images/edits", request)
}

// ImageVariation creates a variation of a given image.
func (c *Client) ImageVariation(ctx context.Context, request ImageVariationRequest) (ImageVariationResponse, error) {
	return makeFormDataRequest[ImageVariationResponse](ctx, c, "/images/variations", request)
}

// Embedding creates an embedding vector representing the input text.
func (c *Client) Embedding(ctx context.Context, request EmbeddingRequest) (EmbeddingResponse, error) {
	return makeJSONRequest[EmbeddingResponse](ctx, c, http.MethodPost, "/embeddings", request)
}

// Files returns a list of files that belong to the user's organization.
func (c *Client) Files(ctx context.Context) (FilesResponse, error) {
	return makeJSONRequest[FilesResponse](ctx, c, http.MethodGet, "/files", nil)
}

// UploadFile upload a file that contains document(s) to be used across various endpoints/features.
// Currently, the size of all the files uploaded by one organization can be up to 1 GB.
func (c *Client) UploadFile(ctx context.Context, request UploadFileRequest) (File, error) {
	return makeFormDataRequest[File](ctx, c, "/files", request)
}

// DeleteFile deletes a file.
func (c *Client) DeleteFile(ctx context.Context, id string) (DeleteFileResponse, error) {
	return makeJSONRequest[DeleteFileResponse](ctx, c, http.MethodDelete, "/files/"+id, nil)
}

// File returns information about a specific file.
func (c *Client) File(ctx context.Context, id string) (File, error) {
	return makeJSONRequest[File](ctx, c, http.MethodGet, "/files/"+id, nil)
}

// FileContent returns the contents of the specified file.
func (c *Client) FileContent(ctx context.Context, id string) (string, error) {
	return makeJSONRequest[string](ctx, c, http.MethodGet, "/files/"+id+"/content", nil)
}

// CreateFineTune creates a job that fine-tunes a specified model from a given dataset.
func (c *Client) CreateFineTune(ctx context.Context, request FineTuneRequest) (FineTune, error) {
	return makeJSONRequest[FineTune](ctx, c, http.MethodPost, "/fine-tunes", request)
}

// FineTunes list your organization's fine-tuning jobs.
func (c *Client) FineTunes(ctx context.Context) (FineTunesResponse, error) {
	return makeJSONRequest[FineTunesResponse](ctx, c, http.MethodGet, "/fine-tunes", nil)
}

// FineTune gets info about the fine-tune job.
func (c *Client) FineTune(ctx context.Context, id string) (FineTune, error) {
	return makeJSONRequest[FineTune](ctx, c, http.MethodGet, "/fine-tunes"+id, nil)
}

// CancelFineTune immediately cancel a fine-tune job.
func (c *Client) CancelFineTune(ctx context.Context, id string) (FineTune, error) {
	return makeJSONRequest[FineTune](ctx, c, http.MethodPost, "/fine-tunes"+id+"/cancel", nil)
}

// FineTuneEvents get fine-grained status updates for a fine-tune job.
func (c *Client) FineTuneEvents(ctx context.Context, id string) (FineTuneEventsResponse, error) {
	return makeJSONRequest[FineTuneEventsResponse](ctx, c, http.MethodGet, "/fine-tunes"+id+"/events", nil)
}

// Moderation classifies if text violates OpenAI's Content Policy
func (c *Client) Moderation(ctx context.Context, request ModerationRequest) (ModerationResponse, error) {
	return makeJSONRequest[ModerationResponse](ctx, c, http.MethodPost, "/moderations", request)
}

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
