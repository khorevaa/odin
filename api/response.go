package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/khorevaa/odin/errors"
	"github.com/khorevaa/ras-client/messages"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} // @Name Response

func (r *Response) Error() error {

	err, ok := r.Data.(error)

	if !ok {
		return nil
	}

	return err

}

func ErrorResponse(ctx *fiber.Ctx, err error, msgAndArgs ...string) error {

	var data interface{}
	message := messageFromMsgAndArgs(msgAndArgs)

	switch typed := err.(type) {
	case *messages.EndpointFailure:
		data = typed
	case *messages.EndpointMessageFailure:
		data = typed
	case *messages.UnknownMessageError:
		data = typed
	default:
		data = err.Error()
	}

	code := fiber.StatusInternalServerError

	switch errors.GetType(err) {

	case errors.BadRequest:
		code = fiber.StatusBadRequest

	case errors.Other:
		data = err.Error()
	}

	return ctx.Status(code).JSON(&Response{
		Code:    code,
		Message: message,
		Data:    data,
	})

}

func NotImplemented(ctx *fiber.Ctx) error {
	return ErrorResponse(ctx, nil)
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}) error {

	return ctx.Status(fiber.StatusOK).JSON(&Response{
		Code:    fiber.StatusOK,
		Message: successMessage,
		Data:    data,
	})

}

const successMessage = "success"

func HttpResponse(ctx *fiber.Ctx, data interface{}, err error, errMsgAndArgs ...string) error {

	if err != nil {
		return ErrorResponse(ctx, err, errMsgAndArgs...)
	}
	return SuccessResponse(ctx, data)
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
