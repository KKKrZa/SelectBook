package Sever

import (
	"log"
	"net/http"
)

func Server() {
	server := http.Server{
		Addr: ":8089",
		Handler: &CorsMiddleware{
			Next: http.DefaultServeMux,
		},
	}
	RegisterRoutes()

	err := server.ListenAndServe()
	if err != nil {
		log.Printf(err.Error())
		return
	}
}
