package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func ZerologMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", ww.Status()).
			Int("bytes", ww.BytesWritten()).
			Str("remote_addr", r.RemoteAddr).
			Str("duration", time.Since(start).String()).
			Msg("HTTP request")
	})
}
