package handler

import (
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	ABCService "github.com/zhenisduissekov/dummy_service/internal/abc/service"
	"github.com/zhenisduissekov/dummy_service/internal/config"
)

var (
	Msg1  = "msg 1"
	Code1 = "Err1"

	Msg2  = "msg 2"
	Code2 = "Err2"

	Msg3  = "msg 3"
	Code3 = "Err3"
)

type Handler struct {
	DI *config.Container
}

type ResponseToFront struct {
	Body interface{} `json:"redirectLink"`
}

type ErrorResponseToFront struct {
	Status  string `json:"status" default:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

// register payment
func (h *Handler) Register(f *fiber.Ctx) error {
	var items ABCService.RequestRegisterItems

	err := f.BodyParser(&items)
	if err != nil {
		log.Err(err).Msg("error at Register")
		return f.Status(fiber.StatusBadRequest).JSON(&ErrorResponseToFront{
			Message: Msg1,
			Code:    Code1,
			Error:   err.Error(),
		})
	}

	v := validator.New()
	err = v.Struct(items)
	if err != nil {
		log.Err(err).Msg("error at Register")
		return f.Status(fiber.StatusBadRequest).JSON(&ErrorResponseToFront{
			Message: Msg2,
			Code:    Code2,
		})
	}

	data, err := h.DI.ABCService.Register(items)
	if err != nil {
		log.Err(err).Msg("error at service.Register")
		return f.Status(fiber.StatusInternalServerError).JSON(&ErrorResponseToFront{
			Message: Msg3,
			Code:    Code3,
		})
	}

	return f.Status(fiber.StatusOK).JSON(&ResponseToFront{
		Body: data,
	})
}

// schedule payment
// @Summary Get comment by id
// @Security BasicAuth
// @Tags comment
// @Description get comment by id
// @ID getCommentById
// @Accept  json
// @Produce  json
// @Param id path int true "comment ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/comments/{id} [get]
func (h *Handler) Schedule(f *fiber.Ctx) error {
	//TODO: needs completion

	return f.Status(fiber.StatusOK).JSON(&ResponseToFront{
		Body: "endpoint not finished",
	})
}
