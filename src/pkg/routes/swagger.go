package routes

import (
	_ "github.com/AdilBikeev/simple-go-mysql/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var RegisterSwaggerRoutes = func(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
