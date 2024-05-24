package handlers

import (
	// "net/http"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	// import middleware
	"github.com/ahmedaliv/goapi/internal/handlers/middleware"
	
)

func Handler (r *chi.Mux){
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func (router chi.Router){
		router.Use(middleware.Authorization)
		// make a route on / to send a response saying successful login
		// router.Get("/", func(w http.ResponseWriter, r *http.Request){
		// 	w.Write([]byte("Login Successful"))
		// })
		
		router.Get("/coins" , GetCoinBalance)
	})

}