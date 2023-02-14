package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personalidades/", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}/", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}/", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}/", controllers.AtualizaUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
