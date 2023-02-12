package openai

import "github.com/KirillMironov/openai/internal/formdata"

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
	Image          formdata.File       `form:"image"`
	Mask           formdata.File       `form:"mask,omitempty"`
	Prompt         string              `form:"prompt"`
	N              int                 `form:"n,omitempty"`
	Size           ImageSize           `form:"size,omitempty"`
	ResponseFormat ImageResponseFormat `form:"response_format,omitempty"`
	User           string              `form:"user,omitempty"`
}

type ImageVariationRequest struct {
	Image          formdata.File       `form:"image"`
	N              int                 `form:"n,omitempty"`
	Size           ImageSize           `form:"size,omitempty"`
	ResponseFormat ImageResponseFormat `form:"response_format,omitempty"`
	User           string              `form:"user,omitempty"`
}

type EmbeddingRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
	User  string   `json:"user,omitempty"`
}

type UploadFileRequest struct {
	File    formdata.File `form:"file"`
	Purpose string        `form:"purpose"`
}

type FineTuneRequest struct {
	TrainingFile                 string    `json:"training_file"`
	ValidationFile               string    `json:"validation_file,omitempty"`
	Model                        string    `json:"model,omitempty"`
	NEpochs                      int       `json:"n_epochs,omitempty"`
	BatchSize                    int       `json:"batch_size,omitempty"`
	LearningRateMultiplier       float64   `json:"learning_rate_multiplier,omitempty"`
	PromptLossWeight             float64   `json:"prompt_loss_weight,omitempty"`
	ComputeClassificationMetrics bool      `json:"compute_classification_metrics,omitempty"`
	ClassificationNClasses       int       `json:"classification_n_classes,omitempty"`
	ClassificationPositiveClass  string    `json:"classification_positive_class,omitempty"`
	ClassificationBetas          []float64 `json:"classification_betas,omitempty"`
	Suffix                       string    `json:"suffix,omitempty"`
}

type ModerationRequest struct {
	Input []string `json:"input"`
	Model string   `json:"model,omitempty"`
}
