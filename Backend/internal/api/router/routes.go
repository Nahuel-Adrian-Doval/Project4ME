package router

import (
	"net/http"

	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/internal/api/controllers"
)

func Routes(prefix string, mux *http.ServeMux) {

	mux.HandleFunc("GET "+prefix+"hello", controllers.Hello)
}
