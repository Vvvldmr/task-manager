package handlers

import (
	"net/http"

	"task-manager/internal/renderer"
)

// HomeHandler отвечает за работу
// с главной страницей.
type HomeHandler struct {

	renderer *renderer.Renderer

}

// Создает обработчик.
func NewHomeHandler(
	r *renderer.Renderer,
) *HomeHandler {

	return &HomeHandler{
		renderer: r,
	}

}

// Главная страница.
func (h *HomeHandler) Index(
	w http.ResponseWriter,
	r *http.Request,
) {

	err := h.renderer.Render(
		w,
		"base",
		nil,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

	}

}