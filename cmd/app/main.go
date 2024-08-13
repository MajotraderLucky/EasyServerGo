package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Определяем функцию-обработчик HTTP запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Отправляем ответ 'Hello, World!' клиенту
		fmt.Fprintf(w, "Hello, World!")
	})

	// Настраиваем сервер для прослушивания порта 8080
	fmt.Println("Server is running on http://localhost:8085")
	if err := http.ListenAndServe("0.0.0.0:8085", nil); err != nil {
		fmt.Println(err)
	}
}
