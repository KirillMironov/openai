package openai

type ModelsResponse struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

type DeleteModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int                           `json:"index"`
		Message      ChatCompletionResponseMessage `json:"message"`
		FinishReason string                        `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
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
			TokenLogprobs []float64            `json:"token_logprobs"`
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
			TokenLogprobs []float64            `json:"token_logprobs"`
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

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL     string `json:"url"`
		B64JSON string `json:"b64_json"`
	} `json:"data"`
}

type ImageEditResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL     string `json:"url"`
		B64JSON string `json:"b64_json"`
	} `json:"data"`
}

type ImageVariationResponse struct {
	Created int `json:"created"`
	Data    []struct {
		URL     string `json:"url"`
		B64JSON string `json:"b64_json"`
	} `json:"data"`
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Model  string `json:"model"`
	Data   []struct {
		Index     int       `json:"index"`
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type FilesResponse struct {
	Object string `json:"object"`
	Data   []File `json:"data"`
}

type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FineTunesResponse struct {
	Object string     `json:"object"`
	Data   []FineTune `json:"data"`
}

type FineTuneEventsResponse struct {
	Object string          `json:"object"`
	Data   []FineTuneEvent `json:"data"`
}

type ModerationResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Flagged    bool `json:"flagged"`
		Categories struct {
			Hate            bool `json:"hate"`
			HateThreatening bool `json:"hate/threatening"`
			SelfHarm        bool `json:"self-harm"`
			Sexual          bool `json:"sexual"`
			SexualMinors    bool `json:"sexual/minors"`
			Violence        bool `json:"violence"`
			ViolenceGraphic bool `json:"violence/graphic"`
		} `json:"categories"`
		CategoryScores struct {
			Hate            float64 `json:"hate"`
			HateThreatening float64 `json:"hate/threatening"`
			SelfHarm        float64 `json:"self-harm"`
			Sexual          float64 `json:"sexual"`
			SexualMinors    float64 `json:"sexual/minors"`
			Violence        float64 `json:"violence"`
			ViolenceGraphic float64 `json:"violence/graphic"`
		} `json:"category_scores"`
	} `json:"results"`
}
