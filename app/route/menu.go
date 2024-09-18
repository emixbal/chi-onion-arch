package route

import (
	"chi-onion-arch/app/handler/menu"

	"github.com/go-chi/chi/v5"
)

func MenuRoute(r chi.Router) {

	h := menu.NewMenuHandler()

	r.Route("/menu", func(r chi.Router) {
		// MiddlewareAuthentication dapat ditambahkan di sini

		r.Get("/list", h.List)
		r.Get("/detail/{id}", h.Detail)
		r.Post("/insert", h.Insert)
		r.Put("/update/{id}", h.Update)
	})
}
