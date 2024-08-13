package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Определяем функцию-обработчик HTTP запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Получаем IP-адрес клиента из запроса
		ip := r.RemoteAddr
		// Получаем текущее время
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		// Логируем IP-адрес и время запроса
		fmt.Printf("Received request from %s at %s\n", ip, currentTime)

		// Отправляем ответ 'Hello, World!' клиенту
		fmt.Fprintf(w, "Hello, World!")
	})

	// Настраиваем сервер для прослушивания порта 8085
	fmt.Println("Server is running on http://localhost:8085")
	if err := http.ListenAndServe("0.0.0.0:8085", nil); err != nil {
		fmt.Println(err)
	}
}
