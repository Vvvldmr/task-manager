package renderer

import (
	"html/template"
	"net/http"
)

// Renderer отвечает исключительно
// за работу с HTML-шаблонами.
type Renderer struct {
	templates *template.Template
}

// New загружает все шаблоны.
func New() *Renderer {

	t := template.Must(
		template.ParseGlob("web/templates/*.html"),
	)

	return &Renderer{
		templates: t,
	}
}

// Render отображает выбранный шаблон.
func (r *Renderer) Render(
	w http.ResponseWriter,
	name string,
	data any,
) error {

	return r.templates.ExecuteTemplate(
		w,
		name,
		data,
	)

}