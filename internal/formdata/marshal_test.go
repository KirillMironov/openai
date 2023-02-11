package formdata

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMarshal(t *testing.T) {
	t.Parallel()

	type form struct {
		Name string   `form:"name"`
		Age  int      `form:"age,omitempty"`
		File *os.File `form:"file"`
	}

	path := filepath.Join(t.TempDir(), "test.txt")

	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { file.Close() })

	if _, err = file.WriteString("test"); err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)

	example := form{
		Name: "John",
		Age:  0,
		File: file,
	}

	data, contentType, err := Marshal(example)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
	t.Log(contentType)
}
