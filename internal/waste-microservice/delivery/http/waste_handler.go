package http

import (
	"net/http"
	"strconv"
	"time"

	"waste-management-service/internal/domain"
	"waste-management-service/internal/waste-microservice/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase *usecase.WasteUsecase
}

func NewHandler(usecase *usecase.WasteUsecase) *Handler {
	return &Handler{usecase: usecase}
}

// 1. fungsi helper (meniru format dari transaksi)
func formatWasteDate(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func formatWaste(waste *domain.Waste) map[string]interface{} {
	return map[string]interface{}{
		"id":                 waste.ID,
		"name":               waste.Name,
		"stock_availability": waste.StockAvailability,
		"costs":              waste.Costs,
		"category":           waste.Category,
		"is_separated":       waste.IsSeparated,
		"created_at":         formatWasteDate(waste.CreatedAt),
	}
}

// Create godoc
// @Summary Create waste catalog
// @Description Create a new waste category/service option
// @Tags Wastes
// @Accept json
// @Produce json
// @Param waste body domain.Waste true "Waste Data"
// @Success 201 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Router /api/v1/wastes [post]
func (h *Handler) Create(c echo.Context) error {
	var waste domain.Waste
	if err := c.Bind(&waste); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.usecase.Create(c.Request().Context(), &waste); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Gunakan formatWaste
	return c.JSON(http.StatusCreated, formatWaste(&waste))
}

// GetAll godoc
// @Summary Get all wastes
// @Description Get all wastes with optional filters for category and separation status
// @Tags Wastes
// @Produce json
// @Param category query string false "Filter by Category (e.g., Organik, Anorganik)"
// @Param is_separated query boolean false "Filter by separation status (true / false)"
// @Success 200 {array} domain.Waste
// @Failure 500 {object} map[string]string
// @Router /api/v1/wastes [get]
func (h *Handler) GetAll(c echo.Context) error {
	category := c.QueryParam("category")
	isSeparated := c.QueryParam("is_separated")

	wastes, err := h.usecase.GetAll(c.Request().Context(), category, isSeparated)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Looping untuk memformat seluruh array
	response := make([]map[string]interface{}, 0)
	for i := range wastes {
		response = append(response, formatWaste(&wastes[i]))
	}

	return c.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary Get waste by ID
// @Description Get specific waste service by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [get]
func (h *Handler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid waste id"})
	}

	waste, err := h.usecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// mengembalikan formatWaste biar sama dengan transaction
	return c.JSON(http.StatusOK, formatWaste(waste))
}

// Update godoc
// @Summary Update waste
// @Description Update waste catalog information
// @Tags Wastes
// @Accept json
// @Produce json
// @Param id path int true "Waste ID"
// @Param waste body domain.Waste true "Waste Data"
// @Success 200 {object} domain.Waste
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [put]
func (h *Handler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid waste id"})
	}

	var waste domain.Waste
	if err := c.Bind(&waste); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	waste.ID = id
	if err := h.usecase.Update(c.Request().Context(), &waste); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedWaste, err := h.usecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// Gunakan formatWaste
	return c.JSON(http.StatusOK, formatWaste(updatedWaste))
}

// Delete godoc
// @Summary Delete waste
// @Description Delete waste catalog by ID
// @Tags Wastes
// @Produce json
// @Param id path int true "Waste ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/wastes/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid waste id"})
	}

	if err := h.usecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	// 2. Handling delete Ubah dari NoContent menjadi JSON notifikasi sukses
	return c.JSON(http.StatusOK, map[string]string{
		"message": "waste successfully deleted",
	})
}
