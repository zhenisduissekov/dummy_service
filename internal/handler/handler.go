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

// Register
// @Summary Register endpoint
// @Security BasicAuth
// @Description register this request.
// @Tags register
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param input body ABCService.RequestRegisterItems true "Register Body"
// @Success 200 {object} ResponseToFront
// @Failure 400 {object} ErrorResponseToFront
// @Failure 401 {object} ErrorResponseToFront
// @Failure 404 {object} ErrorResponseToFront
// @Failure 500 {object} ErrorResponseToFront
// @Router /api/register [post]
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

func (h *Handler) Schedule(f *fiber.Ctx) error {
	//TODO: needs completion

	return f.Status(fiber.StatusOK).JSON(&ResponseToFront{
		Body: "endpoint not finished",
	})
}
