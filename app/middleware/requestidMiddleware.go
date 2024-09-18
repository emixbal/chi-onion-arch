package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// middleware to pass chi request id to trace_id in response API each endpoint
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// dapatkan value requestID dari chi middleware
		requestID := middleware.GetReqID(r.Context())

		// Set the RequestID in the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestID", requestID)
		r = r.WithContext(ctx)

		// Set the RequestID in the response header
		w.Header().Set("X-Request-ID", requestID)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
