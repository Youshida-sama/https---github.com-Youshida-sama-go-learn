package validations

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
)

var (
	uni *ut.UniversalTranslator

	validate *validator.Validate
)

type CoreValidator struct {
	Validator   *validator.Validate
	Translation ut.Translator
}

func NewValidator() (v *CoreValidator, err error) {
	v = &CoreValidator{
		Validator: validator.New(),
	}

	ru := ru.New()
	uni = ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

	v.Translation = trans

	err = ru_translations.RegisterDefaultTranslations(v.Validator, trans)

	if err != nil {
		return
	}

	err = v.Validator.RegisterValidation("isoTime", IsISO8601Date)

	if err != nil {
		return
	}

	err = v.Validator.RegisterValidation("req", RequireAnotherField)

	if err != nil {
		return
	}

	err = v.Validator.RegisterTranslation(
		"isoTime",
		v.Translation,
		func(ut ut.Translator) error {
			return ut.Add("isoTime", "{0} содержит некорретный формат даты", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("isoTime", fe.Field())
			return t
		},
	)

	return
}

func (cv *CoreValidator) Validate(i interface{}) (err error) {
	err = cv.Validator.Struct(i)
	if err != nil {
		msg := ""
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			msg = fmt.Sprint(msg, " ", e.Translate(cv.Translation), ";")
		}
		err = errors.New(msg)
	}

	return
}

// "2001-03-24T16:21:21.269Z"
func IsISO8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(\\d{4})(-(0[1-9]|1[0-2])(-([12]\\d|0[1-9]|3[01]))([T\\s]((([01]\\d|2[0-3])((:)[0-5]\\d))([\\:]\\d+)?)?(:[0-5]\\d([\\.]\\d+)?)?([zZ]|([\\+-])([01]\\d|2[0-3]):?([0-5]\\d)?)?)?)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func RequireAnotherField(fl validator.FieldLevel) bool {
	paramField := fl.Param()

	if paramField == `` {
		return true
	}

	var paramFieldValue reflect.Value

	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(paramField)
	}

	selfValue := fl.Field().String()

	value := paramFieldValue.String()

	if selfValue == "" && value == "" {
		return true
	}

	return value != ""
}
