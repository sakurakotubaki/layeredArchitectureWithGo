package handlers

import (
	"net/http"
	"strconv"

	"github.com/example/layeredArchitectureWithGo/internal/domain"
	"github.com/example/layeredArchitectureWithGo/internal/usecase"
	"github.com/labstack/echo/v4"
)

// WorkHandler は作業に関するHTTPハンドラーです
type WorkHandler struct {
	useCase *usecase.WorkUseCase
}

// NewWorkHandler は新しいWorkHandlerを作成します
func NewWorkHandler(u *usecase.WorkUseCase) *WorkHandler {
	return &WorkHandler{
		useCase: u,
	}
}

// GetAll は全ての作業を取得するハンドラーです
func (h *WorkHandler) GetAll(c echo.Context) error {
	works, err := h.useCase.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, works)
}

// GetByID は指定されたIDの作業を取得するハンドラーです
func (h *WorkHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	work, err := h.useCase.GetByID(c.Request().Context(), id)
	if err == domain.ErrWorkNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, work)
}

// Create は新しい作業を作成するハンドラーです
func (h *WorkHandler) Create(c echo.Context) error {
	work := new(domain.Work)
	if err := c.Bind(work); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.useCase.Create(c.Request().Context(), work); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, work)
}

// Update は既存の作業を更新するハンドラーです
func (h *WorkHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	work := new(domain.Work)
	if err := c.Bind(work); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	work.ID = id

	if err := h.useCase.Update(c.Request().Context(), work); err == domain.ErrWorkNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, work)
}

// Delete は指定されたIDの作業を削除するハンドラーです
func (h *WorkHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.useCase.Delete(c.Request().Context(), id); err == domain.ErrWorkNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
