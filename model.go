package openai

type Engine struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Ready   bool   `json:"ready"`
}

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type File struct {
	ID            string         `json:"id"`
	Object        string         `json:"object"`
	Bytes         int            `json:"bytes"`
	CreatedAt     int            `json:"created_at"`
	Filename      string         `json:"filename"`
	Purpose       string         `json:"purpose"`
	Status        string         `json:"status"`
	StatusDetails map[string]any `json:"status_details"`
}

type FineTune struct {
	ID              string          `json:"id"`
	Object          string          `json:"object"`
	CreatedAt       int             `json:"created_at"`
	UpdatedAt       int             `json:"updated_at"`
	Model           string          `json:"model"`
	FineTunedModel  string          `json:"fine_tuned_model"`
	OrganizationID  string          `json:"organization_id"`
	Status          string          `json:"status"`
	Hyperparams     map[string]any  `json:"hyperparams"`
	TrainingFiles   []File          `json:"training_files"`
	ValidationFiles []File          `json:"validation_files"`
	ResultFiles     []File          `json:"result_files"`
	Events          []FineTuneEvent `json:"events"`
}

type FineTuneEvent struct {
	Object    string `json:"object"`
	CreatedAt int    `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}
