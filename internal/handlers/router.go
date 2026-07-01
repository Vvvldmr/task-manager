package handlers

import (
	"net/http"

	"task-manager/internal/renderer"
)

// NewRouter регистрирует абсолютно все маршруты приложения.
//
// В дальнейшем именно здесь будут подключаться:
//
// /tasks
// /comments
// /settings
// /api
// и так далее.
func NewRouter(r *renderer.Renderer) *http.ServeMux {

	mux := http.NewServeMux()

	h := NewHomeHandler(r)

	mux.HandleFunc("/", h.Index)

	fs := http.FileServer(http.Dir("web/static"))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", fs),
	)

	return mux

}