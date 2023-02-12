package testutil

import (
	"bytes"
	"crypto/rand"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func FilesEqual(t *testing.T, f1, f2 io.ReadSeeker) bool {
	t.Helper()

	_, _ = f1.Seek(0, 0)
	_, _ = f2.Seek(0, 0)

	b1, err := io.ReadAll(f1)
	if err != nil {
		t.Error(err)
	}

	b2, err := io.ReadAll(f2)
	if err != nil {
		t.Error(err)
	}

	return bytes.Equal(b1, b2)
}

// MustCreateTempFile creates a temporary file with random data.
// The file is automatically removed after the test is completed.
func MustCreateTempFile(t *testing.T, bytesSize int) *os.File {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { file.Close() })

	data := make([]byte, bytesSize)

	_, _ = rand.Read(data)

	if _, err = file.Write(data); err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)

	return file
}

// MustCreateImagePNG creates a temporary PNG image with a red line and transparent background.
// The file is automatically removed after the test is completed.
func MustCreateImagePNG(t *testing.T, pixelsSize int) *os.File {
	t.Helper()

	img := image.NewNRGBA(image.Rect(0, 0, pixelsSize, pixelsSize))

	mid := pixelsSize / 2

	for x := 0; x < pixelsSize; x++ {
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
