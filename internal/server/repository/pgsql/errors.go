package pgsql

import (
	"errors"
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"strings"
)

func checkError(err error) error {
	cause := errors.Unwrap(err)
	if cause == nil {
		cause = err
	}

	if strings.Contains(cause.Error(), "violates foreign key constraint") {
		return fmt.Errorf("%s. %w", err, model.ErrRelations)
	}

	return err
}
