package main

import (
    "html/template"
    "net/http"
    "log"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            Title:   "Привет, мир!",
            Message: "Это моя первая тестовая страница на Go.",
        }

        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            log.Fatalf("Ошибка при парсинге шаблона: %v", err)
        }

        err = tmpl.Execute(w, data)
        if err != nil {
            log.Fatalf("Ошибка при выполнении шаблона: %v", err)
        }
    })

    log.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}