package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Code     any `json:"code"`
	Status   any `json:"status"`
	Messages any `json:"messages"`
	Data     any `json:"data,omitempty"`
}

type Messages = []any

type Error struct {
	Code     any `json:"code"`
	Status   any `json:"status"`
	Messages any `json:"messages"`
	Data     any `json:"data,omitempty"`
}

func NewError(code int, messages ...any) *Error {
	e := &Error{
		Code:     code,
		Messages: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		e.Messages = messages[0]
	}
	return e
}

var IsProduction bool

const (
	ResponseSuccess = "SUCCESS"
)

func (e *Error) Error() string {
	return fmt.Sprint(e.Messages)
}

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Response{
		Code: fiber.StatusInternalServerError,
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		resp.Code = fiber.StatusForbidden
		resp.Messages = Messages{removeTopStruct(e.Translate(trans))}
	} else if e, ok := err.(*fiber.Error); ok {
		resp.Code = e.Code
		resp.Messages = Messages{e.Message}
	} else if e, ok := err.(*Error); ok {
		resp.Code = e.Code
		resp.Messages = Messages{e.Messages}

		if resp.Messages == nil {
			resp.Messages = Messages{err}
		}
	} else {
		resp.Messages = Messages{err.Error()}
	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error http")
	}

	return Resp(c, resp)
}

func Resp(c *fiber.Ctx, resp Response) error {
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(fiber.StatusOK)
	return c.JSON(resp)
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, msg := range fields {
		stripStruct := field[strings.Index(field, ".")+1:]
		res[stripStruct] = msg
	}
	return res
}
