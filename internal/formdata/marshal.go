package formdata

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const formTag = "form"

type File interface {
	Name() string
	io.Reader
}

func Marshal(value any) (data []byte, contentType string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("formdata: %v", r)
		}
	}()

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

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
			file := field.Interface().(File)

			if file == nil {
				continue
			}

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
