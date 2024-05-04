package httpsvc

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidArgument     = echo.NewHTTPError(http.StatusBadRequest, setErrorMessage("invalid argument"))
	ErrInternal            = echo.NewHTTPError(http.StatusInternalServerError, setErrorMessage("internal system error"))
	ErrUnauthenticated     = echo.NewHTTPError(http.StatusUnauthorized, setErrorMessage("unauthenticated"))
	ErrNotFound            = echo.NewHTTPError(http.StatusNotFound, setErrorMessage("record not found"))
	ErrProductAlreadyExist = echo.NewHTTPError(http.StatusBadRequest, setErrorMessage("product already exist on product"))
	ErrPermissionDenied    = echo.NewHTTPError(http.StatusForbidden, setErrorMessage("permission denied"))
)

// httpValidationOrInternalErr return valdiation or internal error
func httpValidationOrInternalErr(err error) error {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// Jika tidak ada kesalahan validasi, mengembalikan kesalahan internal
		return ErrInternal
	}

	fields := make(map[string]string)
	for _, validationError := range validationErrors {
		tag := validationError.Tag()
		fields[validationError.Field()] = fmt.Sprintf("Failed on the '%s' tag", tag)
	}

	return echo.NewHTTPError(http.StatusBadRequest, setErrorMessage(fields))
}
