package util

import (
	"chi-onion-arch/app/model"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// ResponseJSON is a combined helper function that sets the response with a trace ID, timestamp, and writes it as JSON.
func ResponseJSON(ctx context.Context, w http.ResponseWriter, statusCode int, resp *model.ApiResponse) {
	// Extracting trace ID from the context (if available).
	traceIdInf := ctx.Value("requestID")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}

	// Setting the timestamp and trace ID.
	resp.Timestamp = time.Now()
	resp.TraceID = traceId

	// Setting the response headers and status code.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encoding the response as JSON and sending it to the client.
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		// Handling encoding errors by sending an internal server error response.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
