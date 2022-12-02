package base_validator

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"go-boilerplate/src/constants/error_code"
	"go-boilerplate/src/shared/exception"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// parseError takes an error or multiple errors and attempts to determine the best path to convert them into
// human readable strings
func parseError(errs ...error) []string {
	var out []string
	for _, err := range errs {
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			// if the type is validator.ValidationErrors then it's actually an array of validator.FieldError so we'll
			// loop through each of those and convert them one by one
			for _, e := range typedError {
				out = append(out, parseFieldError(e))
			}
		case *json.UnmarshalTypeError:
			// similarly, if the error is an unmarshalling error we'll parse it into another, more readable string format
			out = append(out, parseMarshallingError(*typedError))
		default:
			out = append(out, err.Error())
		}
	}
	return out
}

func parseFieldError(e validator.FieldError) string {
	// workaround to the fact that the `gt|gtfield=Start` gets passed as an entire tag for some reason
	// https://github.com/go-playground/validator/issues/926
	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required_without":
		return fmt.Sprintf("%s is required if %s is not supplied", fieldPrefix, e.Param())
	case "lt", "ltfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be less than %s", fieldPrefix, param)
	case "gt", "gtfield":
		param := e.Param()
		if param == "" {
			param = time.Now().Format(time.RFC3339)
		}
		return fmt.Sprintf("%s must be greater than %s", fieldPrefix, param)
	default:
		// if it's a tag for which we don't have a good format string yet we'll try using the default english translator
		english := en.New()
		translator := ut.New(english, english)
		if translatorInstance, found := translator.GetTranslator("en"); found {
			return e.Translate(translatorInstance)
		} else {
			return fmt.Errorf("%v", e).Error()
		}
	}
}
func parseMarshallingError(e json.UnmarshalTypeError) string {
	return fmt.Sprintf("The field %s must be a %s", e.Field, e.Type.String())
}

func ValidatorsBody(c *gin.Context, dto any) {
	if err := c.ShouldBindJSON(dto); err != nil {
		panic(exception.BadRequestError(error_code.INVALID_PARAMS, parseError(err)))
		return
	}
}

func ValidatorQuery(c *gin.Context, dto any) {

	if err := c.ShouldBindQuery(dto); err != nil {
		panic(exception.BadRequestError(error_code.INVALID_PARAMS, parseError(err)))
		return
	}
}

func ValidatorParams(c *gin.Context, dto any) {
	if err := c.ShouldBindUri(dto); err != nil {
		panic(exception.BadRequestError(error_code.INVALID_PARAMS, parseError(err)))
		return
	}
}
