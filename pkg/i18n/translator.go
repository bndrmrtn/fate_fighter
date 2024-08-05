package i18n

import (
	"fmt"
	"strings"
)

type ITranslator interface {
	Language() string
	Languages() []string
	Translate(key string, values ...string) string
	T(key string, values ...string) string
}

type Translator struct {
	language  string
	languages []string
	data      map[string]string
}

func (t *Translator) Language() string {
	return t.language
}

func (t *Translator) Languages() []string {
	return t.languages
}

func (t *Translator) Translate(key string, values ...string) string {
	val, ok := t.data[key]
	if !ok {
		return key
	}

	if len(values) > 0 {
		for i, v := range values {
			val = strings.ReplaceAll(val, fmt.Sprintf("{%d}", i), v)
		}
	}

	return val
}

func (t *Translator) T(key string, values ...string) string {
	return t.Translate(key, values...)
}