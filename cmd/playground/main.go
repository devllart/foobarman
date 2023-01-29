package main

import (
	"devllart/foobarman/internal/scenes"
	"devllart/foobarman/src/funcs"
	"fmt"
)

func main() {
	fmt.Println(funcs.IsFunc(scenes.Bar, "Bar"))
	// Инициализируем роутер, добавляем путь и хендлер для домашней страницы.
	// mux := pat.New()
	// mux.Get("/:locale", http.HandlerFunc(handleHome))

	// Запускаем HTTP-сервер с помощью роутера.
	// log.Print("starting server on :4018...")
	// err := http.ListenAndServe(":4018", mux)
	// log.Fatal(err)

}
