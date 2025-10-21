package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// trans is the global translator.
var trans ut.Translator

// InitTranslator initializes the validator and its Chinese translator.
func InitTranslator() error {
	// Get the validator instance used by Gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register a custom tag name function to get json tag for field names
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zh := zh.New()
		uni := ut.New(zh, zh)

		var ok bool
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(zh) failed")
		}

		// Register the default Chinese translations
		if err := zh_translations.RegisterDefaultTranslations(v, trans); err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf("failed to get validator engine")
}

// Translate translates a validation error into Chinese.
// It returns the first validation error message.
func Translate(err error) string {
	if err == nil {
		return ""
	}

	// Attempt to cast to validator.ValidationErrors
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		// If it's not a validation error, return the original error message
		return err.Error()
	}

	// Translate the first error and return
	for _, e := range validatorErrs {
		return e.Translate(trans)
	}

	return err.Error()
}
