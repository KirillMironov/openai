package openai

type ModelsResponse struct {
	Object string `json:"object"`
	Data   []struct {
		ID      string `json:"id"`
		Object  string `json:"object"`
		Created int    `json:"created"`
		OwnedBy string `json:"owned_by"`
	} `json:"data"`
}

type ModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text     string `json:"text"`
		Index    int    `json:"index"`
		Logprobs struct {
			Tokens        []string             `json:"tokens"`
			TokenLogprobs []int                `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int                `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type EditResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text     string `json:"text"`
		Index    int    `json:"index"`
		Logprobs struct {
			Tokens        []string             `json:"tokens"`
			TokenLogprobs []int                `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int                `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
