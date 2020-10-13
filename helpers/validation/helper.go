package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Descriptive(rawErrors error) map[string]string {

	var verr validator.ValidationErrors

	errors.As(rawErrors, &verr)

	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
			errs[f.Field()] = err
		} else {
			errs[f.Field()] = fmt.Sprintf("This field should be %s", err)
		}
	}

	return errs
}
