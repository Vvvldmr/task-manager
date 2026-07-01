package main

import (
	"html/template"
	"log"
	"net/http"
)

// ----------------------------------------------------------
// Точка входа в приложение.
//
// Именно отсюда начинается запуск всего нашего Task Manager.
//
// Пока приложение максимально простое:
// - загружаем HTML-шаблоны;
// - настраиваем маршруты;
// - запускаем HTTP-сервер.
//
// По мере развития проекта main.go останется небольшим,
// а вся логика будет переноситься в отдельные пакеты.
// ----------------------------------------------------------

func main() {

	// Загружаем все HTML-шаблоны из каталога templates.
	//
	// template.Must вызывает panic, если шаблоны содержат ошибку.
	// Это удобно, потому что приложение не должно запускаться
	// с поврежденными шаблонами.
	templates := template.Must(template.ParseGlob("web/templates/*.html"))

	// Создаем новый HTTP-мультиплексор.
	//
	// Через него будут регистрироваться все маршруты приложения.
	mux := http.NewServeMux()

	// Главная страница.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := templates.ExecuteTemplate(w, "base", nil)
		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

	})

	// Подключаем статические файлы.
	//
	// Благодаря этому браузер сможет получать:
	//
	// CSS
	// JavaScript
	// изображения
	//
	fs := http.FileServer(http.Dir("web/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("----------------------------------------")
	log.Println("Task Manager")
	log.Println("Server started")
	log.Println("http://localhost:8080")
	log.Println("----------------------------------------")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}