package formdata

import (
	"bytes"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/KirillMironov/openai/internal/testutil"
)

func TestMarshal(t *testing.T) {
	t.Parallel()

	type requestForm struct {
		Name string   `form:"name"`
		Age  int      `form:"age,omitempty"`
		File *os.File `form:"file"`
	}

	tests := []struct {
		name    string
		in      any
		wantErr bool
		valid   func(t *testing.T, form *multipart.Form, in any)
	}{
		{
			name: "valid form",
			in: &requestForm{
				Name: "John",
				Age:  20,
				File: testutil.MustCreateTempFile(t, 100),
			},
			wantErr: false,
			valid: func(t *testing.T, form *multipart.Form, in any) {
				if len(form.Value) != 2 {
					t.Fatalf("expected 2 values, got %d", len(form.Value))
				}
				if len(form.File) != 1 {
					t.Fatalf("expected 1 file, got %d", len(form.File))
				}
				if got, want := form.Value["name"][0], "John"; got != want {
					t.Errorf("expected name to be %s, got %s", want, got)
				}
				if got, want := form.Value["age"][0], "20"; got != want {
					t.Errorf("expected age to be %s, got %s", want, got)
				}
				file := in.(*requestForm).File
				if got, want := form.File["file"][0].Filename, filepath.Base(file.Name()); got != want {
					t.Errorf("expected file name to be %s, got %s", want, got)
				}
				formFile, err := form.File["file"][0].Open()
				if err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { formFile.Close() })
				if !testutil.FilesEqual(t, file, formFile) {
					t.Error("expected files data to be equal")
				}
			},
		},
		{
			name: "omit empty age",
			in: requestForm{
				Name: "John",
				Age:  0,
				File: testutil.MustCreateTempFile(t, 100),
			},
			wantErr: false,
			valid: func(t *testing.T, form *multipart.Form, in any) {
				if len(form.Value) != 1 {
					t.Fatalf("expected 1 value, got %d", len(form.Value))
				}
				if len(form.File) != 1 {
					t.Fatalf("expected 1 file, got %d", len(form.File))
				}
				if _, ok := form.Value["age"]; ok {
					t.Errorf("expected age to be omitted")
				}
			},
		},
		{
			name: "omit empty file",
			in: requestForm{
				Name: "John",
				Age:  20,
				File: nil,
			},
			wantErr: false,
			valid: func(t *testing.T, form *multipart.Form, in any) {
				if len(form.Value) != 2 {
					t.Fatalf("expected 2 values, got %d", len(form.Value))
				}
				if len(form.File) != 0 {
					t.Fatal("expected file to be omitted")
				}
			},
		},
		{
			name:    "empty form",
			in:      requestForm{},
			wantErr: false,
			valid: func(t *testing.T, form *multipart.Form, in any) {
				if len(form.Value) != 1 {
					t.Fatalf("expected 1 value, got %d", len(form.Value))
				}
				if len(form.File) != 0 {
					t.Fatal("expected file to be omitted")
				}
				if got := form.Value["name"][0]; got != "" {
					t.Error("expected name to be empty")
				}
			},
		},
		{
			name: "unsupported field type (complex64)",
			in: struct {
				Number complex64 `form:"number"`
			}{
				Number: 1 + 2i,
			},
			wantErr: true,
		},
		{
			name:    "value is not a struct",
			in:      uint16(10),
			wantErr: true,
		},
		{
			name:    "value is nil",
			in:      nil,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, contentType, err := Marshal(tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Marshal() error = %v, wantErr = %v", err, tc.wantErr)
			}

			if tc.wantErr {
				return
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

			tc.valid(t, form, tc.in)
		})
	}
}
