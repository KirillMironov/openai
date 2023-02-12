package formdata

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const formTag = "form"

// File represents a file to be marshaled into a multipart/form-data request body.
type File interface {
	Name() string
	io.Reader
}

// Marshal encodes the given value into a multipart/form-data request body.
// The value must be a struct or a pointer to a struct.
// If the field implements the File interface, the field is marshaled as a file.
func Marshal(value any) (data []byte, contentType string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("formdata: %v", r)
		}
	}()

	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, "", errors.New("formdata: value must be a struct or a pointer to a struct")
	}

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get(formTag)

		if strings.Contains(tag, "omitempty") && field.IsZero() {
			continue
		}

		tag = strings.Split(tag, ",")[0]

		if tag == "-" {
			continue
		}

		if tag == "" {
			tag = strings.ToLower(t.Field(i).Name)
		}

		if field.Type().Implements(reflect.TypeOf((*File)(nil)).Elem()) {
			if field.IsNil() {
				continue
			}

			file := field.Interface().(File)

			filename := filepath.Base(file.Name())

			formFile, err := writer.CreateFormFile(tag, filename)
			if err != nil {
				return nil, "", err
			}

			if _, err = io.Copy(formFile, file); err != nil {
				return nil, "", err
			}

			continue
		}

		if field.Kind() == reflect.Interface || field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		var fieldValue string

		switch field.Kind() {
		case reflect.String:
			fieldValue = field.String()
		case reflect.Bool:
			fieldValue = strconv.FormatBool(field.Bool())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue = strconv.FormatInt(field.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldValue = strconv.FormatUint(field.Uint(), 10)
		case reflect.Float32:
			fieldValue = strconv.FormatFloat(field.Float(), 'f', -1, 32)
		case reflect.Float64:
			fieldValue = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		default:
			return nil, "", fmt.Errorf("formdata: unsupported type: %s", field.Kind())
		}

		if err = writer.WriteField(tag, fieldValue); err != nil {
			return nil, "", err
		}
	}

	if err = writer.Close(); err != nil {
		return nil, "", err
	}

	return buf.Bytes(), writer.FormDataContentType(), nil
}
