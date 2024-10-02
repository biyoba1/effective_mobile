package middleware

import "errors"

func PageValidate(page int) error {
	if page < 1 {
		return errors.New("invalid page")
	}
	return nil
}
