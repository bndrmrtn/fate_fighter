package i18n

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/bndrmrtn/fate_fighter/pkg/utils"
)

type I18n struct {
	fallback        string
	languages       []string
	translationData map[string]map[string]string
}

type Config struct {
	FallbackLanguage string
}

func NewI18n(c *Config) *I18n {
	i := new(I18n)
	i.fallback = c.FallbackLanguage
	i.translationData = make(map[string]map[string]string)
	i.mustLoadTranslationData()
	return i
}

func (i *I18n) GetTranslator(lang string) ITranslator {
	var language string
	if slices.Contains(i.languages, lang) {
		language = lang
	} else {
		language = i.fallback
	}

	return &Translator{
		language:  language,
		languages: i.languages,
		data:      i.translationData[language],
	}
}

func (i *I18n) mustLoadTranslationData() {
	f, err := utils.WalkDir("./lang", ".json")
	if err != nil {
		log.Fatalf("Failed to load translation data: %e", err)
	}

	for _, file := range f {
		parts := strings.Split(file, "/")
		// getting the language code from the file names [en, hu, ...]
		languageKey := strings.TrimSuffix(parts[len(parts)-1], ".json")

		b, err := os.ReadFile("./" + file)
		if err != nil {
			fmt.Printf("Failed to load language: %v\n", languageKey)
			continue
		}

		var data = make(map[string]string)
		err = json.Unmarshal(b, &data)
		if err != nil {
			fmt.Printf("Failed to unmarshal json for language: %v\n", languageKey)
			continue
		}

		i.translationData[languageKey] = data
		i.languages = append(i.languages, languageKey)
	}
}
