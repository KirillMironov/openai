package openai

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestClient_Models(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	models, err := client.Models()
	if err != nil {
		t.Fatal(err)
	}

	if len(models.Data) == 0 {
		t.Fatal("expected at least one model")
	}
}

func TestClient_Model(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	model, err := client.Model("davinci")
	if err != nil {
		t.Fatal(err)
	}

	if model.ID != "davinci" {
		t.Fatalf("expected model id to be davinci, got %s", model.ID)
	}
}

func TestClient_Completion(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	completion, err := client.Completion(CompletionRequest{
		Model:       "ada",
		Prompt:      []string{"This is a test"},
		MaxTokens:   5,
		Temperature: 0.9,
		N:           3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(completion.Choices) != 3 {
		t.Fatalf("expected 3 choices, got %d", len(completion.Choices))
	}
}

func TestClient_Edit(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	edit, err := client.Edit(EditRequest{
		Model:       "text-davinci-edit-001",
		Input:       "This is a test",
		Instruction: "This is a test",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(edit.Choices) != 1 {
		t.Fatalf("expected 1 choice, got %d", len(edit.Choices))
	}
}

func TestClient_Image(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	image, err := client.Image(ImageRequest{
		Prompt: "This is a test",
		Size:   ImageSize1024x1024,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(image.Data) != 1 {
		t.Fatalf("expected 1 image, got %d", len(image.Data))
	}
}

func TestClient_ImageEdit(t *testing.T) {
	t.Parallel()

	client := newClient(t)

	file := mustCreateImage(t)

	image, err := client.ImageEdit(ImageEditRequest{
		Image:  file,
		Prompt: "Add a cat in a hat",
		Size:   ImageSize256x256,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(image.Data) != 1 {
		t.Fatalf("expected 1 image, got %d", len(image.Data))
	}
}

func newClient(t *testing.T) *Client {
	t.Helper()

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		t.Skip(`provide "OPENAI_API_KEY" environment variable to run this test`)
	}

	return NewClient(apiKey, WithTimeout(time.Second*30))
}

// mustCreateImage creates a temporary image file with a red line in the middle.
func mustCreateImage(t *testing.T) *os.File {
	t.Helper()

	width, height := 64, 64

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	mid := height / 2

	for x := 0; x < width; x++ {
		for y := mid - 5; y < mid+5; y++ {
			img.Set(x, y, color.RGBA{R: 255, A: 255})
		}
	}

	path := filepath.Join(t.TempDir(), "red-line.png")

	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { file.Close() })

	if err = png.Encode(file, img); err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)

	return file
}
