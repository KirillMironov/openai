package formdata

import (
	"bytes"
	"crypto/rand"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

type requestForm struct {
	Name string   `form:"name"`
	Age  int      `form:"age,omitempty"`
	File *os.File `form:"file,omitempty"`
}

func TestMarshal(t *testing.T) {
	t.Parallel()

	file := mustCreateTempFile(t, 100)

	in := requestForm{
		Name: "John",
		Age:  0,
		File: file,
	}

	data, contentType, err := Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		t.Fatal(err)
	}

	r := multipart.NewReader(bytes.NewReader(data), params["boundary"])

	form, err := r.ReadForm(256)
	if err != nil {
		t.Fatal(err)
	}

	if len(form.Value) != 1 {
		t.Errorf("expected 1 value, got %d", len(form.Value))
	}

	if len(form.File) != 1 {
		t.Errorf("expected 1 file, got %d", len(form.File))
	}

	if form.Value["name"][0] != "John" {
		t.Errorf("expected name to be John, got %s", form.Value["name"][0])
	}

	if _, ok := form.Value["age"]; ok {
		t.Error("expected age to be omitted")
	}

	if form.File["file"][0].Size != 100 {
		t.Errorf("expected file size to be 100, got %d", form.File["file"][0].Size)
	}

	if form.File["file"][0].Filename != filepath.Base(file.Name()) {
		t.Errorf("expected file name to be %s, got %s", file.Name(), form.File["file"][0].Filename)
	}

	formFile, err := form.File["file"][0].Open()
	if err != nil {
		t.Fatal(err)
	}
	defer formFile.Close()

	_, _ = file.Seek(0, 0)

	f1, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	f2, err := io.ReadAll(formFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(f1, f2) {
		t.Error("expected file data to be equal")
	}
}

func mustCreateTempFile(t *testing.T, size int) *os.File {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { file.Close() })

	data := make([]byte, size)

	_, _ = rand.Read(data)

	if _, err = file.Write(data); err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)

	return file
}
