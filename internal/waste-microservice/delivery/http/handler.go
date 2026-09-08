package http

import (
	"net/http"
	"strconv"

	"waste-management-service/internal/domain"
	"waste-management-service/internal/waste-microservice/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase *usecase.WasteUsecase
}

func NewHandler(
	usecase *usecase.WasteUsecase,
) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// CreateWaste godoc
// @Summary Create waste
// @Description Create a new waste
// @Tags Wastes
// @Accept json
// @Produce json
// @Param waste body domain.Waste true "Waste"
// @Success 201 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Router /api/v1/wastes [post]
func (h *Handler) Create(
	c echo.Context,
) error {

	var waste domain.Waste

	if err := c.Bind(&waste); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	if err := h.usecase.Create(
		c.Request().Context(),
		&waste,
	); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		waste,
	)
}

// GetAll godoc
// @Summary Get all wastes
// @Description Get all wastes
// @Tags Wastes
// @Produce json
// @Success 200 {array} domain.Waste
// @Failure 500 {object} map[string]string
// @Router /api/v1/wastes [get]
func (h *Handler) GetAll(
	c echo.Context,
) error {

	wastes, err := h.usecase.GetAll(
		c.Request().Context(),
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		wastes,
	)
}

// GetByID godoc
// @Summary Get waste by ID
// @Description Get waste by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [get]
func (h *Handler) GetByID(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid waste id",
			},
		)
	}

	waste, err := h.usecase.GetByID(
		c.Request().Context(),
		id,
	)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		waste,
	)
}

// UpdateWaste godoc
// @Summary Update waste
// @Description Update waste information
// @Tags Wastes
// @Accept json
// @Produce json
// @Param id path int true "Waste ID"
// @Param waste body domain.Waste true "Waste"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [put]
func (h *Handler) Update(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid waste id",
			},
		)
	}

	var waste domain.Waste

	if err := c.Bind(&waste); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	waste.ID = id

	if err := h.usecase.Update(
		c.Request().Context(),
		&waste,
	); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	updatedWaste, err := h.usecase.GetByID(
		c.Request().Context(),
		id,
	)

	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		updatedWaste,
	)
}

// DeleteWaste godoc
// @Summary Delete waste
// @Description Delete waste by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [delete]
func (h *Handler) Delete(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid waste id",
			},
		)
	}

	if err := h.usecase.Delete(
		c.Request().Context(),
		id,
	); err != nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.NoContent(http.StatusNoContent)
}
