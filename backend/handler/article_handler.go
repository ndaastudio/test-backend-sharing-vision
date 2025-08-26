package handler

import (
	"net/http"
	"strconv"
	"test-backend-sharing-vision/request"
	"test-backend-sharing-vision/service"
	"test-backend-sharing-vision/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	service  service.ArticleService
	validate *validator.Validate
}

func NewArticleHandler(e *echo.Echo, s service.ArticleService) {
	h := &ArticleHandler{
		service:  s,
		validate: validator.New(),
	}

	e.POST("/article", h.Create)
	e.GET("/article/:limit/:offset", h.GetAll)
	e.GET("/article/:id", h.GetByID)
	e.PUT("/article/:id", h.Update)
	e.DELETE("/article/:id", h.Delete)
}

func (h *ArticleHandler) GetAll(c echo.Context) error {
	limitParam := c.Param("limit")
	offsetParam := c.Param("offset")
	statusParam := c.QueryParam("status")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Limit tidak valid")
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Offset tidak valid")
	}

	articles, err := h.service.GetAll(limit, offset, statusParam)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, articles)
}

func (h *ArticleHandler) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.service.GetByID(uint(id))
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, "Artikel tidak ditemukan")
	}
	return util.SuccessResponse(c, http.StatusOK, article)
}

func (h *ArticleHandler) Create(c echo.Context) error {
	var req request.CreateArticleRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Permintaan tidak valid")
	}
	if err := h.validate.Struct(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, e := range errs {
				field := e.Field()

				switch e.Tag() {
				case "required":
					errorMessages[field] = field + " tidak boleh kosong"
				case "min":
					errorMessages[field] = field + " harus memiliki minimal " + e.Param() + " karakter"
				case "oneof":
					errorMessages[field] = field + " harus salah satu dari: " + e.Param()
				default:
					errorMessages[field] = "Permintaan tidak valid untuk " + field
				}
			}

			return util.ErrorResponse(c, http.StatusBadRequest, errorMessages)
		}

		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	article, err := h.service.Create(req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(c, http.StatusCreated, article)
}

func (h *ArticleHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req request.UpdateArticleRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "Permintaan tidak valid")
	}
	if err := h.validate.Struct(req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, e := range errs {
				field := e.Field()

				switch e.Tag() {
				case "min":
					errorMessages[field] = field + " harus memiliki minimal " + e.Param() + " karakter"
				case "oneof":
					errorMessages[field] = field + " harus salah satu dari: " + e.Param()
				default:
					errorMessages[field] = "Permintaan tidak valid untuk " + field
				}
			}

			return util.ErrorResponse(c, http.StatusBadRequest, errorMessages)
		}

		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	article, err := h.service.Update(uint(id), req)
	if err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, article)
}

func (h *ArticleHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		return util.ErrorResponse(c, http.StatusNotFound, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, "Artikel berhasil dihapus")
}
