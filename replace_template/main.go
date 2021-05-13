package main

import (
	"bytes"
	"errors"
	"fmt"
)

func main() {
	templateContent := "Hello {{user_name}}, Chuc mung sinh nhat tuoi {{user_age}}"
	values := map[string]string {
		"user_name": "quang",
		"user_age": "18",
	}
	smsBuilder := newDefaultSmsBuilder()
	content, err := smsBuilder.smsContentFromTemplate(templateContent, values)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}


type smsBuilder struct {
	placeHolderBeginning string
	placeHolderEnding    string
}

func newDefaultSmsBuilder() smsBuilder {
	return smsBuilder{placeHolderBeginning: "{{", placeHolderEnding: "}}"}
}

func (b smsBuilder) smsContentFromTemplate(templateContent string, values map[string]string) (string, error) {
	var (
		buf   bytes.Buffer
		begin = -1
		i     = 0
		ErrInsufficientDynamicValues = errors.New("insufficient dynamic values")

	)
	for i < len(templateContent) {
		// is beginning of place holder
		if i+len(b.placeHolderBeginning)-1 < len(templateContent) && b.placeHolderBeginning == templateContent[i:i+len(b.placeHolderBeginning)] {
			begin = i
		}

		if begin == -1 {
			_, err := buf.WriteRune(rune(templateContent[i]))
			if err != nil {
				return "", err
			}
			i++
			continue
		}
		// is ending of place holder
		if i+len(b.placeHolderEnding)-1 < len(templateContent) && b.placeHolderEnding == templateContent[i:i+len(b.placeHolderEnding)] {
			v, ok := values[templateContent[begin+len(b.placeHolderBeginning):i]]
			if !ok {
				return "", ErrInsufficientDynamicValues
			}
			_, err := buf.WriteString(v)
			if err != nil {
				return "", err
			}
			i += len(b.placeHolderEnding)
			begin = -1
			continue
		}
		i++
	}
	if begin != -1 {
		_, err := buf.WriteString(templateContent[begin:])
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
