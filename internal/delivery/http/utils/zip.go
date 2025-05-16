package utils

import (
	"compress/gzip"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

func ZipAndSendResponse(logger *zerolog.Logger, w http.ResponseWriter, res any) {
	// Set response headers
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// Create gzip writer that wraps the response writer
	gz := gzip.NewWriter(w)
	defer gz.Close()

	// Create JSON encoder that writes to gzip stream
	encoder := json.NewEncoder(gz)

	// Encode and send compressed JSON
	if err := encoder.Encode(res); err != nil {
		logger.Error().Err(err).Msg("fail to encode json")
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
