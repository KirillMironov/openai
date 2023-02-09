package openai

type CompletionRequest struct {
	Model            string         `json:"model"`
	Prompt           []string       `json:"prompt,omitempty"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float64        `json:"temperature,omitempty"`
	TopP             float64        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Logprobs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float64        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64        `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

type EditRequest struct {
	Model       string  `json:"model"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction"`
	N           int     `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type ImageSize string

const (
	ImageSize256x256   ImageSize = "256x256"
	ImageSize512x512   ImageSize = "512x512"
	ImageSize1024x1024 ImageSize = "1024x1024"
)

type ImageResponseFormat string

const (
	ImageResponseFormatURL     ImageResponseFormat = "url"
	ImageResponseFormatB64JSON ImageResponseFormat = "b64_json"
)

type ImageRequest struct {
	Prompt         string              `json:"prompt"`
	N              int                 `json:"n,omitempty"`
	Size           ImageSize           `json:"size,omitempty"`
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	User           string              `json:"user,omitempty"`
}

type ImageEditRequest struct {
	Image          string              `json:"image"`
	Mask           string              `json:"mask,omitempty"`
	Prompt         string              `json:"prompt"`
	N              int                 `json:"n,omitempty"`
	Size           ImageSize           `json:"size,omitempty"`
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	User           string              `json:"user,omitempty"`
}

type ImageVariationRequest struct {
	Image          string              `json:"image"`
	N              int                 `json:"n,omitempty"`
	Size           ImageSize           `json:"size,omitempty"`
	ResponseFormat ImageResponseFormat `json:"response_format,omitempty"`
	User           string              `json:"user,omitempty"`
}

type EmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
	User  string   `json:"user,omitempty"`
}

type UploadFileRequest struct {
	File    string `json:"file"`
	Purpose string `json:"purpose"`
}
