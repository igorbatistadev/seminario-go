package routes

import (
	"backend/controllers"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Cliente)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
