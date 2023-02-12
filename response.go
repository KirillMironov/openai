package openai

type ModelsResponse struct {
	Object string `json:"object"`
	Data   []struct {
		ID         string           `json:"id"`
		Object     string           `json:"object"`
		Created    int              `json:"created"`
		OwnedBy    string           `json:"owned_by"`
		Permission []map[string]any `json:"permission"`
		Root       string           `json:"root"`
		Parent     string           `json:"parent"`
	} `json:"data"`
}

type ModelResponse struct {
	ID         string           `json:"id"`
	Object     string           `json:"object"`
	Created    int              `json:"created"`
	OwnedBy    string           `json:"owned_by"`
	Permission []map[string]any `json:"permission"`
	Root       string           `json:"root"`
	Parent     string           `json:"parent"`
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
	Data   []struct {
		ID            string   `json:"id"`
		Object        string   `json:"object"`
		Bytes         int      `json:"bytes"`
		CreatedAt     int      `json:"created_at"`
		Filename      string   `json:"filename"`
		Purpose       string   `json:"purpose"`
		Status        string   `json:"status"`
		StatusDetails []string `json:"status_details"`
	} `json:"data"`
}

type UploadFileResponse struct {
	ID            string   `json:"id"`
	Object        string   `json:"object"`
	Bytes         int      `json:"bytes"`
	CreatedAt     int      `json:"created_at"`
	Filename      string   `json:"filename"`
	Purpose       string   `json:"purpose"`
	Status        string   `json:"status"`
	StatusDetails []string `json:"status_details"`
}

type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FileResponse struct {
	ID            string   `json:"id"`
	Object        string   `json:"object"`
	Bytes         int      `json:"bytes"`
	CreatedAt     int      `json:"created_at"`
	Filename      string   `json:"filename"`
	Purpose       string   `json:"purpose"`
	Status        string   `json:"status"`
	StatusDetails []string `json:"status_details"`
}
