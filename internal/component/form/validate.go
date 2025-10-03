package form

import (
	"errors"
	"net/url"
)

func validateRepository(repository string) error {
	_, err := url.ParseRequestURI(repository)
	if err != nil {
		return errors.New("url is invalid")
	}

	return nil
}
