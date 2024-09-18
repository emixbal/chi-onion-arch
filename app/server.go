package app

import (
	"chi-onion-arch/app/middleware"
	"chi-onion-arch/app/route"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

// SetupRouter mengatur routing dan middleware untuk aplikasi
func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middleware RequestID
	r.Use(chiMiddleware.RequestID)

	r.Use(chiMiddleware.Recoverer)

	// logger
	r.Use(chiMiddleware.Logger)

	// middleware to pass chi request id to trace_id in response API each endpoint
	r.Use(middleware.RequestID)

	// Grup route /v1 dengan middleware ClaimHeaderCheckMiddleware
	r.Route("/v1", func(v1 chi.Router) {
		v1.Use(middleware.ClaimHeaderCheckMiddleware)

		route.MenuRoute(v1)
	})

	return r
}

func Run() {
	// Memanggil SetupRouter untuk menginisialisasi routing
	router := SetupRouter()

	// Menentukan alamat dan port server
	address := ":" + viper.GetString(`server.address`)

	// Membuat server HTTP dengan timeout dan pengaturan lain
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Menampilkan pesan server berjalan dan alamat
	log.Printf("Server running at %s\n", address)

	// Menjalankan server, log fatal jika gagal
	log.Fatal(srv.ListenAndServe())
}
