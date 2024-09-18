package middleware

import (
	"chi-onion-arch/app/helper/util"
	"chi-onion-arch/app/model"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func ClaimHeaderCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		xId := r.Header.Get("X-ID")
		xRole := r.Header.Get("X-Role")

		// Mengecek header X-ID
		if xId == "" {
			apiResponse := model.ApiResponse{
				Message: "Missing X-ID header",
				Code:    http.StatusBadRequest,
			}

			util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
			return
		}

		// Mengecek header X-Role
		if xRole == "" {
			apiResponse := model.ApiResponse{
				Message: "Missing X-Role header",
				Code:    http.StatusBadRequest,
			}

			util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
			return
		}

		// Mengecek header X-ID apakah valid UUID
		_, err := uuid.Parse(xId)
		if err != nil {
			apiResponse := model.ApiResponse{
				Message: "Invalid X-ID format",
				Code:    http.StatusBadRequest,
			}
			util.ResponseJSON(ctx, w, apiResponse.Code, &apiResponse)
			return
		}

		// Menambahkan header ke context untuk penggunaan lebih lanjut
		ctx = context.WithValue(ctx, "X-ID", xId)
		ctx = context.WithValue(ctx, "X-Role", xRole)

		// Lanjut ke handler berikutnya
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
