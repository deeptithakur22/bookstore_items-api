package app

import (
	"net/http"

	"github.com/deeptithakur22/bookstore_items-api/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)

}
