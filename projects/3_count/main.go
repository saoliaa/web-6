package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
	"strconv" // вдруг понадобиться вам ;)
)

var counter int = 0

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(counter)))
	case http.MethodPost:
		err := r.ParseForm()
		if err == nil {
			countStr := r.FormValue("count")
			if countStr == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
				return
			}

			count, err := strconv.Atoi(countStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
				return
			}

			counter += count
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Метод не поддерживается"))
	}
}

func main() {
	http.HandleFunc("/count", countHandler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
