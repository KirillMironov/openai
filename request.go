package openai

type CompletionRequest struct {
	Model            string         `json:"model"`
	Prompt           string         `json:"prompt,omitempty"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      int            `json:"temperature,omitempty"`
	TopP             int            `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Logprobs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             string         `json:"stop,omitempty"`
	PresencePenalty  int            `json:"presence_penalty,omitempty"`
	FrequencyPenalty int            `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

type EditRequest struct {
	Model       string `json:"model"`
	Input       string `json:"input,omitempty"`
	Instruction string `json:"instruction"`
	N           int    `json:"n,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
}
