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
// @Router /wastes [post]
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

	err := h.usecase.Create(
		c.Request().Context(),
		&waste,
	)

	if err != nil {
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

// GetAllWastes godoc
// @Summary Get all wastes
// @Description Get all wastes
// @Tags Wastes
// @Produce json
// @Success 200 {array} domain.Waste
// @Failure 500 {object} map[string]string
// @Router /wastes [get]
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

// GetWaste godoc
// @Summary Get waste by ID
// @Description Get waste by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wastes/{id} [get]
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
// @Description Update waste data by ID
// @Tags Wastes
// @Accept json
// @Produce json
// @Param id path int true "Waste ID"
// @Param waste body domain.Waste true "Waste"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wastes/{id} [put]
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

	err = h.usecase.Update(
		c.Request().Context(),
		&waste,
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
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

// DeleteWaste godoc
// @Summary Delete waste
// @Description Delete waste by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /wastes/{id} [delete]
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

	err = h.usecase.Delete(
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

	return c.NoContent(
		http.StatusNoContent,
	)
}
