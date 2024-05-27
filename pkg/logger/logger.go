package logger

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}

func LogRequest(handler http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		requestID := uuid.New().String()
		requestLogger := logger.With().Logger()

		defer func() {
			if rec := recover(); rec != nil {
				requestLogger.Error().Interface("recovered", rec).Msg("recovered from panic")
				http.Error(ww, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		for name, headers := range r.Header {
			for _, h := range headers {
				requestLogger.Debug().Str("header_"+name, h).Msg("")
			}
		}

		if err := r.ParseForm(); err == nil {
			for key, values := range r.Form {
				for _, value := range values {
					requestLogger.Debug().Str("param_"+key, value).Msg("")
				}
			}
		}

		body, err := io.ReadAll(r.Body)
		if err == nil {
			requestLogger.Debug().Str("body", string(body)).Msg("")
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		handler.ServeHTTP(ww, r)

		logMessage := strings.Builder{}
		logMessage.WriteString("request_id: " + requestID + "\n")
		logMessage.WriteString("method: " + r.Method + "\n")
		logMessage.WriteString("uri: " + r.RequestURI + "\n")
		logMessage.WriteString("remote_addr: " + r.RemoteAddr + "\n")
		logMessage.WriteString("status_code: " + strconv.Itoa(ww.statusCode) + "\n")
		logMessage.WriteString("response_time: " + time.Since(start).String() + "\n")

		switch {
		case ww.statusCode >= 500:
			requestLogger.Error().Msg(logMessage.String())
		case ww.statusCode >= 400:
			requestLogger.Warn().Msg(logMessage.String())
		default:
			requestLogger.Info().Msg(logMessage.String())
		}
	})
}
