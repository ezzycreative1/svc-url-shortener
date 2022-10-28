// mvalidator wrapper dari validator, menambahkan translator dan error response dalam bentuk map[string]
package mvalidator

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/pkg/slicer"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validator interface {
	// Real mengembalian validator instance asli
	Real() *validator.Validate

	// SliceStruct menerima payload []struct dan mengembalikan validasi error yang
	// didapatkan dalam bentuk map string, dan error aslinya. inputan valid apabila err nil
	SliceStruct(input interface{}) (map[string]interface{}, error)

	// SliceStruct menerima payload []struct dan mengembalikan validasi error yang
	// didapatkan dalam bentuk map string dan error aslinya. inputan valid apabila err nil
	Struct(input interface{}) (map[string]interface{}, error)

	// Var menerima payload variable dan mengembalikan validasi error yang
	// didapatkan dalam bentuk map string dan error aslinya. inputan valid apabila err nil
	// eg.
	// var i int
	// validate.Var(i, "gt=1,lt=10")
	Var(input interface{}, tag string) (map[string]interface{}, error)
}

type mValidator struct {
	instance   *validator.Validate
	translator ut.Translator
}

// =======================================================================================
// register validator baru pada bagian code New()  <<<<<<

// New menginisiasi validator dan mengembalikan custom validator
// dengan format error map
func New() *mValidator {

	// init instance of 'validate' with sane defaults
	// init default translator
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Panic("translator not found")
	}

	// register default translation
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	// register tag e.Field() use json tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if len(name) == 0 {
			name = strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
		}

		if name == "-" {
			return ""
		}

		return name
	})

	// register additional func validation here << ---- <<

	_ = validate.RegisterTranslation("majoodate", trans, func(ut ut.Translator) error {
		return ut.Add("majoodate", "{0} must be valid date format", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("majoodate", fe.Field())
		return t
	})

	_ = validate.RegisterValidation("majoodate", func(fl validator.FieldLevel) bool {
		str := fl.Field().String()
		layout := "2006-01-02 15:04:05"
		_, err := time.Parse(layout, str)
		return err == nil
	})

	_ = validate.RegisterTranslation("cdate", trans, func(ut ut.Translator) error {
		return ut.Add("cdate", "{0} must be valid date format", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("cdate", fe.Field())
		return t
	})

	_ = validate.RegisterValidation("cdate", func(fl validator.FieldLevel) bool {
		str := fl.Field().String()
		if str == "" {
			return true
		}
		layout := "2006-01-02"
		_, err := time.Parse(layout, str)
		return err == nil
	})

	return &mValidator{
		instance:   validate,
		translator: trans,
	}
}

// =======================================================================================

// Real mengembalian validator instance asli
func (m *mValidator) Real() *validator.Validate {
	return m.instance
}

// SliceStruct menerima payload []struct dan mengembalikan validasi error yang
// didapatkan dalam bentuk map string, dan error aslinya jika valid akan mengembalikan nil
func (m *mValidator) SliceStruct(input interface{}) (map[string]interface{}, error) {

	errMap := make(map[string]interface{})

	if input == nil || !slicer.IsSlice(input) {
		errorMsg := "developer_vault. input cant be nil or other than slice struct"
		errMap["message"] = errorMsg
		return errMap, errors.New(errorMsg)
	}

	err := m.instance.Var(input, "omitempty,dive")
	if err != nil {

		// not accepted error by validator
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMap["message"] = err.Error()
			return errMap, err
		}

		// iterate error message
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errMap[e.Field()] = e.Translate(m.translator)
		}
		return errMap, errs
	}

	return nil, nil
}

// SliceStruct menerima payload []struct dan mengembalikan validasi error yang
// didapatkan dalam bentuk map string dan error aslinya jika valid akan mengembalikan nil
func (m *mValidator) Struct(input interface{}) (map[string]interface{}, error) {

	errMap := make(map[string]interface{})

	if input == nil {
		errorMsg := "developer_vault. input cant be nil"
		errMap["message"] = errorMsg
		return errMap, errors.New(errorMsg)
	}

	err := m.instance.Struct(input)
	if err != nil {
		// not accepted error by validator
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMap["message"] = err.Error()
			return errMap, err
		}

		// iterate error message
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errMap[e.Field()] = e.Translate(m.translator)
		}
		return errMap, errs
	}

	return nil, nil
}

// Var menerima payload variable dan mengembalikan validasi error yang
// didapatkan dalam bentuk map string dan error aslinya jika valid akan mengembalikan nil
func (m *mValidator) Var(input interface{}, tag string) (map[string]interface{}, error) {

	errMap := make(map[string]interface{})

	err := m.instance.Var(input, tag)
	if err != nil {
		// not accepted error by validator
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMap["message"] = err.Error()
			return errMap, err
		}

		// iterate error message
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errMap[e.Field()] = e.Translate(m.translator)
		}
		return errMap, errs
	}

	return nil, nil
}
