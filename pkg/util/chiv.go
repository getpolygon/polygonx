package util

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	o "github.com/go-ozzo/ozzo-validation"
)

// ChivField will contain the value of the validatable field
// and the custom validation that will need to be applied to
// it.
type ChivSection = map[string]o.Rule

// Chiv is a simple validator for chi, which contains request
// parts that will need to be validated by go-ozzo, including
// URL query, params, headers, etc.
type Chiv struct {
	Params  ChivSection
	Query   ChivSection
	Headers ChivSection
}

func implValidationFor(section ChivSection, vFunc func(field string, rule o.Rule) error) error {
	for field := range section {
		rule := section[field]
		err := vFunc(field, rule)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Chiv) Validate(r *http.Request) error {
	err := implValidationFor(c.Params, func(field string, validator o.Rule) error {
		param := chi.URLParam(r, field)
		err := o.Validate(param, validator)
		return err
	})
	if err != nil {
		return err
	}

	err = implValidationFor(c.Query, func(field string, validator o.Rule) error {
		query := r.URL.Query().Get(field)
		err := o.Validate(query, validator)
		return err
	})
	if err != nil {
		return err
	}

	err = implValidationFor(c.Headers, func(field string, validator o.Rule) error {
		header := r.Header.Get(field)
		err := o.Validate(header, validator)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
