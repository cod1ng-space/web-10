package api

import (
	"errors"
	"net/http"

	"github.com/cod1ng-space/web-10/pkg/vars"

	"github.com/labstack/echo/v4"
)

// GetQuery возвращает случайное приветствие пользователю
func (srv *Server) GetQuery(e echo.Context) error {
	msg, err := srv.uc.FetchQueryMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// PostQuery Помещает новый вариант приветствия в БД
func (srv *Server) PostQuery(e echo.Context) error {
	input := struct {
		Msg *string `json:"name"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "the 'name' field is missing")
	}

	if *input.Msg == "" {
		return e.String(http.StatusBadRequest, "name is empty")
	}

	if len([]rune(*input.Msg)) > srv.maxSize {
		return e.String(http.StatusBadRequest, "size name too large")
	}

	err = srv.uc.SetQueryMessage(*input.Msg)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusCreated, "OK")
}
