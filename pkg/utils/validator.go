package utils

import (
	"context"
	"konntent-service-template/pkg/validation"
)

func ValidateRequestWithCtx(c context.Context, r interface{}) map[string]string {
	v := c.Value(Validator).(validation.IValidator)
	if errs := v.Validate(r); len(errs) > 0 {
		return errs
	}

	return nil
}
